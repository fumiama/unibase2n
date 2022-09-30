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
	num.add(one)
	assert.Equal(t, num, zero)
	num.add(readuint128be([]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}))
	num.shr(8)
	assert.Equal(t, num, &uint128be{0x001234567890abcd, 0xef1234567890abcd})
	num.shr(16)
	assert.Equal(t, num, &uint128be{0x0000001234567890, 0xabcdef1234567890})
	num.shr(32)
	assert.Equal(t, num, &uint128be{0x0000000000000012, 0x34567890abcdef12})
	num.shr(1)
	// 因为 num.a 低 1 位为 0 所以可以成功
	assert.Equal(t, num, &uint128be{0x0000000000000012 >> 1, 0x34567890abcdef12 >> 1})
	num.shr(7)
	num.write(buf[:])
	assert.Equal(t, buf, [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef})
}
