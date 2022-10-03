package unibase2n

import (
	"bytes"
	"math/bits"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnDecode(t *testing.T) {
	var buf [4096]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, bits.OnesCount8(8))
	bs := Base{off: 0, til: 32768, bit: 1}
	for ; bs.bit < 16; bs.bit++ {
		for i := 1; i <= 4096; i++ {
			e := bs.Encode(buf[:i])
			d := bs.Decode(e)
			if !bytes.Equal(d, buf[:i]) {
				t.Fatal(bs.bit, i, e, d, buf[:i])
			}
		}
	}
}

func TestDec128blk1(t *testing.T) {
	var in, tmp [32]byte
	_, err := rand.Read(in[:])
	if err != nil {
		t.Fatal(err)
	}
	n, _ := Base{bit: 1}.EncodeLen(32)
	out := make([]byte, n)
	enc16blk1(0x2333, in[:], out)
	t.Log(out)
	dec128blk1(uint128be{0x2333233323332333, 0x2333233323332333}, out[:], tmp[:])
	assert.Equal(t, in, tmp)
}

func TestDec64blk2(t *testing.T) {
	var in, tmp [32]byte
	_, err := rand.Read(in[:])
	if err != nil {
		t.Fatal(err)
	}
	n, _ := Base{bit: 2}.EncodeLen(32)
	out := make([]byte, n)
	enc64blk2(0x2333233323332333, in[:], out)
	t.Log(out)
	dec64blk2(0x2333233323332333, out[:], tmp[:])
	assert.Equal(t, in, tmp)
}

func TestDec32blk4(t *testing.T) {
	var in, tmp [32]byte
	_, err := rand.Read(in[:])
	if err != nil {
		t.Fatal(err)
	}
	n, _ := Base{bit: 4}.EncodeLen(32)
	out := make([]byte, n)
	enc32blk4(0x23332333, in[:], out)
	t.Log(out)
	dec32blk4(0x23332333, out[:], tmp[:])
	assert.Equal(t, in, tmp)
}

func TestDec16blk8(t *testing.T) {
	var in, tmp [32]byte
	_, err := rand.Read(in[:])
	if err != nil {
		t.Fatal(err)
	}
	n, _ := Base{bit: 8}.EncodeLen(32)
	out := make([]byte, n)
	enc16blk8(0x1100, in[:], out)
	t.Log(out)
	dec16blk8(0x1100, out[:], tmp[:])
	assert.Equal(t, in, tmp)
}

func TestDec128blk(t *testing.T) {
	var in, tmp [32]byte
	_, err := rand.Read(in[:])
	if err != nil {
		t.Fatal(err)
	}
	mask := uint128be{0x2333233323332333, 0x2333233323332333}
	bs := Base{}
	for _, bit := range [...]byte{3, 5, 6, 7, 9, 11, 13, 15} {
		bs.bit = bit
		n, _ := bs.EncodeLen(32)
		out := make([]byte, n)
		enc128blk(mask, bit, in[:], out)
		t.Log(out)
		dec128blk(mask, bit, out[:], tmp[:])
		assert.Equal(t, in, tmp)
	}
}

func TestDec64blk(t *testing.T) {
	var in, tmp [32]byte
	_, err := rand.Read(in[:])
	if err != nil {
		t.Fatal(err)
	}
	// 6
	n, _ := Base{bit: 6}.EncodeLen(32)
	out := make([]byte, n)
	enc64blk(0x2333233323332333, 6, in[:], out)
	t.Log(out)
	dec64blk(0x2333233323332333, 6, out[:], tmp[:])
	assert.Equal(t, in, tmp)
	// 10
	n, _ = Base{bit: 10}.EncodeLen(32)
	out = make([]byte, n)
	enc64blk(0x2333233323332333, 10, in[:], out)
	t.Log(out)
	dec64blk(0x2333233323332333, 10, out[:], tmp[:])
	assert.Equal(t, in, tmp)
	// 12
	n, _ = Base{bit: 12}.EncodeLen(32)
	out = make([]byte, n)
	enc64blk(0x2333233323332333, 12, in[:], out)
	t.Log(out)
	dec64blk(0x2333233323332333, 12, out[:], tmp[:])
	assert.Equal(t, in, tmp)
	// 14
	n, _ = Base{bit: 14}.EncodeLen(32)
	out = make([]byte, n)
	enc64blk(0x2333233323332333, 14, in[:], out)
	t.Log(out)
	dec64blk(0x2333233323332333, 14, out[:], tmp[:])
	assert.Equal(t, in, tmp)
}
