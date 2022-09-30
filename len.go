package unibase2n

import "math/bits"

// EncodeLen calculate encode length
//    offset will be appended as til+offset at last
func (bs Base) EncodeLen(in int) (out, offset int) {
	if in <= 0 {
		return 0, 0
	}
	if bits.OnesCount8(bs.bit) == 1 { // 2的幂永无offset (1/2/4/8)
		return in * (16 >> "\x00\x00\x01\x00\x02\x00\x00\x00\x03"[bs.bit]), 0
	}
	if bs.bit == 6 { // 特判 (6) 3*8=4*6
		out = in / 3 * 4 * 2
		offset = in % 3
		if offset == 0 {
			return
		}
		// 8 = 6+2, 16 = 6+6+4, 再加结尾指示字符
		out += (offset + 1 + 1) * 2
		return
	}
	if bs.bit < 8 { // (3/5/7)
		out = in / int(bs.bit) * 8 * 2
		offset = in % int(bs.bit)
		if offset == 0 {
			return
		}
		remain := offset * 8
		out += (remain/int(bs.bit) + 1 + 1) * 2
		return
	}
	if bs.bit%2 == 0 { // bit 大于8且可被2整除 (10/12/14)
		hb := int(bs.bit) >> 1
		out = in / hb * 8 * 2
		offset = in % hb
		if offset == 0 {
			return
		}
		remain := offset * 8
		out += (remain/int(bs.bit) + 1 + 1) * 2
		return
	}
	// (9/11/13/15)
	out = in / int(bs.bit) * 16 * 2
	offset = in % int(bs.bit)
	if offset == 0 {
		return
	}
	remain := offset * 8
	out += (remain/int(bs.bit) + 1 + 1) * 2
	return
}
