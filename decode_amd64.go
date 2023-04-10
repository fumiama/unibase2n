//go:build amd64
// +build amd64

package unibase2n

// dec128blk1 for bit 1
//
//	len(in)>0, len(in)%16==0, len(out)==len(in)/16
//
//go:noescape
func dec128blk1(mask uint128be, in, out []byte)
