package unibase2n

import (
	"encoding/binary"
)

// Encode data to base2n
func (bs Base) Encode(data []byte) (out []byte) {
	outlen, offset := bs.EncodeLen(len(data))
	if outlen <= 0 || offset < 0 {
		return nil
	}
	out = make([]byte, outlen)
	bs.encode(data, out, outlen, offset)
	return
}

// EncodeTo data to base2n
func (bs Base) EncodeTo(data, out []byte) {
	outlen, offset := bs.EncodeLen(len(data))
	if outlen <= 0 || offset < 0 {
		return
	}
	bs.encode(data, out, outlen, offset)
}

func (bs Base) encode(data, out []byte, outlen, offset int) {
	switch bs.bit {
	case 3, 5, 7, 9, 11, 13, 15:
		mask64 := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
		mask := uint128be{a: mask64, b: mask64}
		enc128blk(mask, bs.bit, data, out)
	case 6, 10, 12, 14:
		mask := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
		enc64blk(mask, bs.bit, data, out)
	case 8:
		enc16blk8(bs.off, data, out)
	case 4:
		mask := uint32(bs.off) | uint32(bs.off)<<16
		enc32blk4(mask, data, out)
	case 2:
		mask := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
		enc64blk2(mask, data, out)
	case 1:
		enc16blk1(bs.off, data, out)
	}
	if offset > 0 {
		binary.BigEndian.PutUint16(out[outlen-2:outlen], bs.til+uint16(offset))
	}
}

// enc128blk for bit 3 5 6 7 9 11 13 15
//
//	len(in)>0, len(out)==len(in)/bit*16
//
//go:nosplit
func enc128blk(mask uint128be, bit uint8, in, out []byte) {
	// 由于最后一次处理有读取越界, 因此作扩展
	var buf [16]byte
	last := 0
	if len(in)%int(bit) == 0 {
		last = len(in) - int(bit)
	} else {
		last = len(in) / int(bit) * int(bit)
	}
	copy(buf[:], in[last:])
	n := 0
	c := 16 - bit
	for i := 0; i < last; i += int(bit) {
		m := uint128be{a: (uint64(1)<<bit - 1) << 48}
		shift := readuint128be(in[i:])
		shift.shreq(c)
		sum := shift.and(m)
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		shift.shreq(c)
		m.shreq(16)
		sum.oreq(shift.and(m))
		sum.addeq(mask)
		sum.write(out[n : n+16])
		n += 16
	}
	// 最后一次
	m := uint128be{a: (uint64(1)<<bit - 1) << 48}
	shift := readuint128be(buf[:])
	shift.shreq(c)
	sum := shift.and(m)
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	shift.shreq(c)
	m.shreq(16)
	sum.oreq(shift.and(m))
	sum.addeq(mask)
	sum.write(buf[:])
	copy(out[n:], buf[:])
}

// enc64blk for bit 6 10 12 14
//
//	len(in)>0, len(out)==len(in)/bit*16
//
//go:nosplit
func enc64blk(mask uint64, bit uint8, in, out []byte) {
	// 由于最后一次处理有读取越界, 因此作扩展
	var buf [8]byte
	s := int(bit) >> 1
	last := 0
	if len(in)%s == 0 {
		last = len(in) - s
	} else {
		last = len(in) / s * s
	}
	copy(buf[:], in[last:])
	n := 0
	c := 16 - bit
	for i := 0; i < last; i += s {
		m := (uint64(1)<<bit - 1) << 48
		b := in[i:]
		if len(b) < 8 {
			b = append(b, make([]byte, 8-len(b))...)
		}
		shift := binary.BigEndian.Uint64(b) >> c
		sum := shift
		sum &= m
		m >>= 16
		shift >>= c
		sum |= shift & m
		m >>= 16
		shift >>= c
		sum |= shift & m
		m >>= 16
		shift >>= c
		sum |= shift & m
		sum += mask
		binary.BigEndian.PutUint64(out[n:], sum)
		n += 8
	}
	m := (uint64(1)<<bit - 1) << 48
	shift := binary.BigEndian.Uint64(buf[:]) >> c
	sum := shift
	sum &= m
	m >>= 16
	shift >>= c
	sum |= shift & m
	m >>= 16
	shift >>= c
	sum |= shift & m
	m >>= 16
	shift >>= c
	sum |= shift & m
	sum += mask
	binary.BigEndian.PutUint64(buf[:], sum)
	copy(out[n:], buf[:])
}
