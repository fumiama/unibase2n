//go:build amd64
// +build amd64

package unibase2n

// enc16blk1 for bit 1 (actual enc128blk1)
//    len(in)!=0, len(out)==len(in)*16
func enc16blk1(mask uint16, in, out []byte)
