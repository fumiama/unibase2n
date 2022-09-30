package unibase2n

import (
	"encoding/binary"
	"math/bits"
)

type uint128be struct {
	a uint64
	b uint64
}

func readuint128be(b []byte) *uint128be {
	return &uint128be{
		a: binary.BigEndian.Uint64(b[:8]),
		b: binary.BigEndian.Uint64(b[8:16]),
	}
}

func (num *uint128be) add(n *uint128be) {
	var c uint64
	num.b, c = bits.Add64(num.b, n.b, 0)
	num.a, _ = bits.Add64(num.a, n.a, c)
	return
}

// shr only supports shifting 1 ~ 63 bits
func (num *uint128be) shr(c uint8) {
	mask := uint64(1)<<c - 1
	aout := (num.a & mask) << (64 - c)
	num.a = num.a >> c
	num.b = aout | (num.b >> c)
}

func (num *uint128be) and(n *uint128be) {
	num.a &= n.a
	num.b &= n.b
}

func (num *uint128be) write(b []byte) {
	binary.BigEndian.PutUint64(b[:8], num.a)
	binary.BigEndian.PutUint64(b[8:16], num.b)
	return
}
