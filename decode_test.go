package unibase2n

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
