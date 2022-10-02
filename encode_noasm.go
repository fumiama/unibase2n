//go:build !amd64
// +build !amd64

package unibase2n

import "encoding/binary"

// enc16blk1 for bit 1 (actual enc128blk1)
//    len(in)!=0, len(out)==len(in)*16
//go:nosplit
func enc16blk1(mask uint16, in, out []byte) {
	for i, n := range in {
		c := i * 16
		binary.BigEndian.PutUint16(out[c:c+2], uint16(n>>7)+mask)
		binary.BigEndian.PutUint16(out[c+2:c+4], uint16(n>>6&1)+mask)
		binary.BigEndian.PutUint16(out[c+4:c+6], uint16(n>>5&1)+mask)
		binary.BigEndian.PutUint16(out[c+6:c+8], uint16(n>>4&1)+mask)
		binary.BigEndian.PutUint16(out[c+8:c+10], uint16(n>>3&1)+mask)
		binary.BigEndian.PutUint16(out[c+10:c+12], uint16(n>>2&1)+mask)
		binary.BigEndian.PutUint16(out[c+12:c+14], uint16(n>>1&1)+mask)
		binary.BigEndian.PutUint16(out[c+14:c+16], uint16(n&1)+mask)
	}
}

// enc64blk2 for bit 2
//    len(in)!=0, len(out)==len(in)*8
//go:nosplit
func enc64blk2(mask uint64, in, out []byte) {
	for i, n := range in {
		c := i * 8
		x := (uint64(n)<<42 | uint64(n)<<28 | uint64(n)<<14 | uint64(n)) & 0x00030003_00030003
		binary.BigEndian.PutUint64(out[c:c+8], x+mask)
	}
}

// enc32blk4 for bit 4
//    len(in)!=0, len(out)==len(in)*4
//go:nosplit
func enc32blk4(mask uint32, in, out []byte) {
	for i, n := range in {
		c := i * 4
		x := (uint32(n)<<12 | uint32(n)) & 0x000f000f
		binary.BigEndian.PutUint32(out[c:c+4], x+mask)
	}
}

// enc16blk8 for bit 8
//    len(in)!=0, len(out)==len(in)*2
//go:nosplit
func enc16blk8(mask uint16, in, out []byte) {
	for i, n := range in {
		c := i * 2
		binary.BigEndian.PutUint16(out[c:c+2], uint16(n)+mask)
	}
}
