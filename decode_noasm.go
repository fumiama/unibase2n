//go:build !amd64
// +build !amd64

package unibase2n

// dec128blk1 for bit 1
//    len(in)>0, len(in)%16==0, len(out)==len(in)/16
//go:nosplit
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
