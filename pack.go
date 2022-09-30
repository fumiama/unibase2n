package unibase2n

import (
	"encoding/binary"
	"unsafe"
)

type Pack uint64

var (
	// Base16384 CJK Unified Ideographs
	//    see https://github.com/fumiama/base16384
	Base16384 = newbasepack(0x4e00, 0x3d00, 14)
	// Base8192 谚文音節 Hangul Syllables
	Base8192 = newbasepack(0xac00, 0xcc00, 13)
	// Base256 谚文字母 Hangul Jamo
	Base256 = newbasepack(0x1100, 0, 8)
	// BaseMath (256) 數學運算符 Mathematical Operators
	BaseMath = newbasepack(0x2200, 0, 8)
	// Base128 帶圈或括號的字母數字 Enclosed Alphanumerics
	Base128 = newbasepack(0x2460, 0x24f4, 7)
	// Base64 箭頭 Arrows
	Base64 = newbasepack(0x2190, 0x21d0, 6)
	// Base64Gua 六十四卦 YiJing Hexagram Symbols
	Base64Gua = newbasepack(0x3400, 0x262f, 6)
	// Base32 方塊元素 Block Elements
	Base32 = newbasepack(0x2580, 0x259f, 5)
	// Base16 漢文訓讀點 Kanbun Kundoku Den
	Base16 = newbasepack(0x3190, 0, 4)
	// Base8 八卦 YiJing Hexagram Symbols
	Base8 = newbasepack(0x2630, 0x2689, 3)
)

func newbasepack(off, til uint16, bit uint8) Pack {
	b, err := NewBase(off, til, bit)
	if err != nil {
		panic(err)
	}
	return b.Pack()
}

// New base2n from a packed config
func New(pack Pack) *Base {
	b := &Base{}
	ismele := isLittleEndian()
	isitle := pack&0xffffff != 0
	if ismele == isitle { // same endian
		*(*Pack)(unsafe.Pointer(b)) = pack
		return b
	}
	field := (*[8]byte)(unsafe.Pointer(&pack))
	if isitle { // packed in little endian but I am big
		b.off = binary.BigEndian.Uint16(field[6:8])
		b.til = binary.BigEndian.Uint16(field[4:6])
	} else { // packed in big endian but I am little
		b.off = binary.LittleEndian.Uint16(field[6:8])
		b.til = binary.LittleEndian.Uint16(field[4:6])
	}
	b.bit = field[3]
	return b
}

// Pack this config into an integer
func (bs *Base) Pack() Pack {
	return *(*Pack)(unsafe.Pointer(bs))
}

func isLittleEndian() bool {
	n := 0x1234
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
}
