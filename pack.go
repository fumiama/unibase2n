package unibase2n

import "unsafe"

type Pack uint64

var (
	Base16384 = newbasepack(0x4e00, 0x3d00, 14)
	Base8192  = newbasepack(0xac00, 0xcc00, 13)
	Base256   = newbasepack(0x1100, 0x0000, 4)
)

func newbasepack(off, til uint16, bit uint8) Pack {
	b, err := NewBase(off, til, bit)
	if err != nil {
		panic(err)
	}
	return b.Pack()
}

func New(pack Pack) *Base {
	b := &Base{}
	*(*Pack)(unsafe.Pointer(b)) = pack
	return b
}

func (b *Base) Pack() Pack {
	return *(*Pack)(unsafe.Pointer(b))
}
