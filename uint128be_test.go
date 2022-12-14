package unibase2n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUint128be(t *testing.T) {
	buf := [16]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	zero := readuint128be(make([]byte, 16))
	one := readuint128be([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	num := readuint128be(buf[:])
	assert.Equal(t, zero.sub(one), num)
	num.addeq(one)
	assert.Equal(t, num, zero)
	num.addeq(readuint128be([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}))
	num.shreq(8)
	assert.Equal(t, num, uint128be{0x001234567890abcd, 0xef1234567890abcd})
	assert.Equal(t, num.String(), "0x1234567890abcdef1234567890abcd")
	num.shreq(16)
	assert.Equal(t, num, uint128be{0x0000001234567890, 0xabcdef1234567890})
	num.shreq(32)
	assert.Equal(t, num, uint128be{0x0000000000000012, 0x34567890abcdef12})
	assert.Equal(t, num.String(), "0x1234567890abcdef12")
	num.shleq(8)
	assert.Equal(t, num, uint128be{0x00000000000001234, 0x567890abcdef1200})
	num.shreq(9)
	// 因为 num.a 低 1 位为 0 所以可以成功
	assert.Equal(t, num, uint128be{0x0000000000000012 >> 1, 0x34567890abcdef12 >> 1})
	num.shreq(7)
	num.write(buf[:])
	assert.Equal(t, buf, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef})
	assert.Equal(t, num.String(), "0x1234567890abcdef")
	num.shreq(8)
	assert.Equal(t, num.String(), "0x1234567890abcd")
}
