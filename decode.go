package unibase2n

import (
	"encoding/binary"
	"math/bits"
)

func (bs Base) Decode(data []byte) (out []byte) {
	if len(data) == 0 || len(data)%2 != 0 {
		return nil
	}
	if bits.OnesCount8(bs.bit) == 1 { // 2的幂永无offset, 无需管til (1/2/4/8)
		out = make([]byte, bs.DecodeLen(len(data), 0))
		switch bs.bit {
		case 8:
			dec16blk8(bs.off, data, out)
		case 4:
			mask := uint32(bs.off) | uint32(bs.off)<<16
			dec32blk4(mask, data, out)
		case 2:
			mask := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
			dec64blk2(mask, data, out)
		case 1:
			mask64 := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
			mask := uint128be{a: mask64, b: mask64}
			dec128blk1(mask, data, out)
		}
		return
	}
	offset := 0
	tile := uint32(bs.til) // [til, tile)
	if bs.bit > 4 && bs.bit%2 == 0 {
		tile += uint32(bs.bit / 2) // 6 10 12 14
	} else {
		tile += uint32(bs.bit) // 3 5 7 9 11 13 15
	}
	realtail := binary.BigEndian.Uint16(data[len(data)-2:])
	if realtail >= bs.til && uint32(realtail) < tile {
		offset = int(realtail - bs.til)
	}
	out = make([]byte, bs.DecodeLen(len(data), offset))
	switch bs.bit {
	case 3, 5, 7, 9, 11, 13, 15:
		mask64 := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
		mask := uint128be{a: mask64, b: mask64}
		dec128blk(mask, bs.bit, data, out)
	case 6, 10, 12, 14:
		mask := uint64(bs.off) | uint64(bs.off)<<16 | uint64(bs.off)<<32 | uint64(bs.off)<<48
		dec64blk(mask, bs.bit, data, out)
	}
	return
}

// dec64blk2 for bit 2
//    len(in)>0, len(in)%8==0, len(out)==len(in)/8
//go:nosplit
func dec64blk2(mask uint64, in, out []byte) {
	for i := range out {
		c := i * 8
		n := binary.BigEndian.Uint64(in[c:c+8]) - mask
		sum := n & 0x03
		n >>= 16 - 2
		sum |= n & 0x0c
		n >>= 16 - 2
		sum |= n & 0x30
		n >>= 16 - 2
		sum |= n & 0xc0
		out[i] = uint8(sum)
	}
}

// dec32blk4 for bit 4
//    len(in)>0, len(in)%4==0, len(out)==len(in)/4
//go:nosplit
func dec32blk4(mask uint32, in, out []byte) {
	for i := range out {
		c := i * 4
		n := binary.BigEndian.Uint32(in[c:c+4]) - mask
		sum := n & 0x0f
		n >>= 16 - 4
		sum |= n & 0xf0
		out[i] = uint8(sum)
	}
}

// dec16blk8 for bit 8
//    len(in)>0, len(in)%2==0, len(out)==len(in)/2
//go:nosplit
func dec16blk8(mask uint16, in, out []byte) {
	for i := range out {
		c := i * 2
		n := binary.BigEndian.Uint16(in[c:c+2]) - mask
		out[i] = uint8(n)
	}
}

// dec128blk for bit 3 5 6 7 9 11 13 15
//    len(in)>0, len(in)%2==0, len(out)==len(in)/16*bit+offset
//go:nosplit
func dec128blk(mask uint128be, bit uint8, in, out []byte) {
	var tmp [16]byte
	i := 0
	n := 0
	c := 16 - bit
	for ; i <= len(out)-int(bit); n += 16 {
		m := uint128be{a: (uint64(1)<<bit - 1) << (64 - bit)}
		shift := readuint128be(in[n:]).sub(mask)
		shift.shleq(c)
		sum := shift.and(m)
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		if len(out[i:]) < 16 {
			sum.write(tmp[:])
			copy(out[i:], tmp[:])
		} else {
			sum.write(out[i:])
		}
		i += int(bit)
	}
	if n < len(in)-2 {
		m := uint128be{a: (uint64(1)<<bit - 1) << (64 - bit)}
		copy(tmp[:], in[n:])
		msk := uint16(mask.a & 0xffff)
		for j := 0; j < 8; j++ {
			x := binary.BigEndian.Uint16(tmp[j*2:])
			if x < msk {
				binary.BigEndian.PutUint16(tmp[j*2:], msk)
			}
		}
		shift := readuint128be(tmp[:]).sub(mask)
		shift.shleq(c)
		sum := shift.and(m)
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		shift.shleq(c)
		m.shreq(bit)
		sum.oreq(shift.and(m))
		sum.write(tmp[:])
		copy(out[i:], tmp[:])
	}
}

// dec64blk for bit 6 10 12 14
//    len(in)>0, len(in)%2==0, len(out)==len(in)/16*bit+offset
//go:nosplit
func dec64blk(mask uint64, bit uint8, in, out []byte) {
	var tmp [8]byte
	i := 0
	s := int(bit) >> 1
	n := 0
	c := 16 - bit
	for ; i <= len(out)-s; n += 8 {
		m := (uint64(1)<<bit - 1) << (64 - bit)
		shift := binary.BigEndian.Uint64(in[n:]) - mask
		shift <<= c
		sum := shift & m
		shift <<= c
		m >>= bit
		sum |= shift & m
		shift <<= c
		m >>= bit
		sum |= shift & m
		shift <<= c
		m >>= bit
		sum |= shift & m
		if len(out[i:]) < 8 {
			binary.BigEndian.PutUint64(tmp[:], sum)
			copy(out[i:], tmp[:])
		} else {
			binary.BigEndian.PutUint64(out[i:], sum)
		}
		i += s
	}
	if n < len(in)-2 {
		m := (uint64(1)<<bit - 1) << (64 - bit)
		copy(tmp[:], in[n:])
		msk := uint16(mask & 0xffff)
		for j := 0; j < 4; j++ {
			x := binary.BigEndian.Uint16(tmp[j*2:])
			if x < msk {
				binary.BigEndian.PutUint16(tmp[j*2:], msk)
			}
		}
		shift := binary.BigEndian.Uint64(tmp[:]) - mask
		shift <<= c
		sum := shift & m
		shift <<= c
		m >>= bit
		sum |= shift & m
		shift <<= c
		m >>= bit
		sum |= shift & m
		shift <<= c
		m >>= bit
		sum |= shift & m
		binary.BigEndian.PutUint64(tmp[:], sum)
		copy(out[i:], tmp[:])
	}
}
