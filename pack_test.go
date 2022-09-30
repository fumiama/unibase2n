package unibase2n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackUnpack(t *testing.T) {
	bs, err := NewBase(0x1234, 0x5678, 8)
	if err != nil {
		t.Fatal(err)
	}
	p := bs.Pack()
	bs1, err := New(p)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, bs, bs1)
	ismele := isLittleEndian()
	if ismele {
		// simulate be pack -> le unpack
		bs2, err := New(0x1234567808000000)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, bs, bs2)
	} else { // simulate le pack -> be unpack
		bs2, err := New(0x0000000878563412)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, bs, bs2)
	}
}
