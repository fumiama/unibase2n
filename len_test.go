package unibase2n

import "testing"

func TestEnDecodeLen(t *testing.T) {
	bs := Base{bit: 1}
	for ; bs.bit < 16; bs.bit++ {
		for i := 1; i <= 65536; i++ {
			in := bs.DecodeLen(bs.EncodeLen(i))
			if i != in {
				t.Fatal("bit:", bs.bit, "in:", i, "!= out:", in)
			}
		}
	}
}
