package unibase2n

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"strconv"
)

type uint128be struct {
	a uint64
	b uint64
}

func readuint128be(b []byte) uint128be {
	if len(b) < 16 {
		b = append(b, make([]byte, 16-len(b))...)
	}
	return uint128be{
		a: binary.BigEndian.Uint64(b[:8]),
		b: binary.BigEndian.Uint64(b[8:16]),
	}
}

func (num *uint128be) addeq(n uint128be) {
	var c uint64
	num.b, c = bits.Add64(num.b, n.b, 0)
	num.a, _ = bits.Add64(num.a, n.a, c)
	return
}

// shreq only supports shifting 1 ~ 63 bits
func (num *uint128be) shreq(c uint8) {
	mask := uint64(1)<<c - 1
	aout := (num.a & mask) << (64 - c)
	num.a = num.a >> c
	num.b = aout | (num.b >> c)
}

func (num *uint128be) andeq(n uint128be) {
	num.a &= n.a
	num.b &= n.b
}

func (num uint128be) and(n uint128be) (r uint128be) {
	r.a = num.a & n.a
	r.b = num.b & n.b
	return
}

func (num *uint128be) oreq(n uint128be) {
	num.a |= n.a
	num.b |= n.b
}

func (num uint128be) or(n uint128be) (r uint128be) {
	r.a = num.a | n.a
	r.b = num.b | n.b
	return
}

func (num *uint128be) write(b []byte) {
	binary.BigEndian.PutUint64(b[:8], num.a)
	binary.BigEndian.PutUint64(b[8:16], num.b)
	return
}

func (num uint128be) String() string {
	if num.a == 0 {
		return "0x" + strconv.FormatUint(num.b, 16)
	}
	return fmt.Sprintf("0x%x%016x", num.a, num.b)
}
