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
