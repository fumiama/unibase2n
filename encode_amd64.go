//go:build amd64
// +build amd64

package unibase2n

// enc16blk1 for bit 1 (actual enc128blk1)
//
//	len(in)!=0, len(out)==len(in)*16
//
//go:noescape
func enc16blk1(mask uint16, in, out []byte)

// enc64blk2 for bit 2
//
//	len(in)!=0, len(out)==len(in)*8
//
//go:noescape
func enc64blk2(mask uint64, in, out []byte)

// enc32blk4 for bit 4
//
//	len(in)!=0, len(out)==len(in)*4
//
//go:noescape
func enc32blk4(mask uint32, in, out []byte)

// enc16blk8 for bit 8
//
//	len(in)!=0, len(out)==len(in)*2
//
//go:noescape
func enc16blk8(mask uint16, in, out []byte)
