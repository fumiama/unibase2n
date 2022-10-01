package unibase2n

import (
	"encoding/binary"
)

func (bs Base) Decode(data []byte) []byte {
	return nil
}

// dec128blk1 for bit 1
//    len(in)>0, len(in)%16==0, len(out)==len(in)/16
func dec128blk1(mask uint128be, in, out []byte) {
	for i := range out {
		c := i * 16
		n := readuint128be(in[c : c+16])
		one := u128one
		n.subeq(mask)
		sum := n.and(one)
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		n.shreq(16 - 1)
		one.shleq(1)
		sum.oreq(n.and(one))
		out[i] = uint8(sum.b)
	}
}

// dec64blk2 for bit 2
//    len(in)!=0, len(out)==len(in)/8
func dec64blk2(mask uint64, in, out []byte) {
	for i, n := range in {
		c := i * 8
		x := (uint64(n)<<42 | uint64(n)<<28 | uint64(n)<<14 | uint64(n)) & 0x00030003_00030003
		binary.BigEndian.PutUint64(out[c:c+8], x+mask)
	}
}

// dec32blk4 for bit 4
//    len(in)!=0, len(out)==len(in)/4
func dec32blk4(mask uint32, in, out []byte) {
	for i, n := range in {
		c := i * 4
		x := (uint32(n)<<12 | uint32(n)) & 0x000f000f
		binary.BigEndian.PutUint32(out[c:c+4], x+mask)
	}
}
