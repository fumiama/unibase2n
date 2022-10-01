package unibase2n

import "math/bits"

const powtab = "\x00\x04\x03\x00\x02\x00\x00\x00\x01"

// EncodeLen calculate encode length
//    offset will be appended as til+offset at last
func (bs Base) EncodeLen(in int) (out, offset int) {
	if in <= 0 {
		return 0, 0
	}
	if bits.OnesCount8(bs.bit) == 1 { // 2的幂永无offset (1/2/4/8)
		return in << powtab[bs.bit], 0
	}
	if bs.bit%2 == 0 { // bit 可被2整除, 64blk (6/10/12/14)
		hb := int(bs.bit) >> 1
		out = in / hb * 8
		offset = in % hb
		if offset == 0 {
			return
		}
		remain := offset * 8
		out += (remain/int(bs.bit) + 1 + 1) * 2
		return
	}
	// 其他奇数, 128blk (3/5/7/9/11/13/15)
	out = in / int(bs.bit) * 16
	offset = in % int(bs.bit)
	if offset == 0 {
		return
	}
	remain := offset * 8
	out += (remain/int(bs.bit) + 1 + 1) * 2
	return
}

// DecodeLen calculate decode length
//    offset must be legal or out will <= 0
func (bs Base) DecodeLen(in, offset int) (out int) {
	if in <= 0 || offset < 0 {
		return 0
	}
	if bits.OnesCount8(bs.bit) == 1 { // 2的幂永无offset (1/2/4/8)
		return in >> powtab[bs.bit]
	}
	if bs.bit%2 == 0 { // bit 可被2整除, 64blk (6/10/12/14)
		hb := int(bs.bit) >> 1
		if offset != 0 {
			remain := offset * 8
			in -= (remain/int(bs.bit) + 1) * 2
		}
		return (in>>3)*hb + offset
	}
	// 其他奇数, 128blk (3/5/7/9/11/13/15)
	if offset != 0 {
		remain := offset * 8
		in -= (remain/int(bs.bit) + 1) * 2
	}
	return (in>>4)*int(bs.bit) + offset
}
