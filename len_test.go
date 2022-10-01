package unibase2n

import (
	"testing"

	base14 "github.com/fumiama/go-base16384"
)

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

func TestBase16384EnDecodeLen(t *testing.T) {
	bs := Base{bit: 14}
	for i := 1; i <= 65536; i++ {
		myi, _ := bs.EncodeLen(i)
		aci := base14.EncodeLen(i)
		if myi != aci {
			t.Fatal("bit:", bs.bit, "input:", i, "me:", myi, "!= actual:", aci)
		}
	}
	for i := 1; i <= 4096; i++ {
		myi, off := bs.EncodeLen(i)
		myo := bs.DecodeLen(myi, off)
		aco := base14.DecodeLen(myi, off)
		if myo != aco {
			t.Fatal("bit:", bs.bit, "input:", i, "me:", myo, "!= actual:", aco)
		}
	}
}
