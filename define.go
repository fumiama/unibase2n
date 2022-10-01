package unibase2n

var (
	// Base16384 CJK Unified Ideographs
	//    see https://github.com/fumiama/base16384
	Base16384, _ = NewBase(0x4e00, 0x3d00, 14)
	// Base8192 谚文音節 Hangul Syllables
	Base8192, _ = NewBase(0xac00, 0xcc00, 13)
	// Base256 谚文字母 Hangul Jamo
	Base256, _ = NewBase(0x1100, 0, 8)
	// BaseMath (256) 數學運算符 Mathematical Operators
	BaseMath, _ = NewBase(0x2200, 0, 8)
	// Base128 帶圈或括號的字母數字 Enclosed Alphanumerics
	Base128, _ = NewBase(0x2460, 0x24f4, 7)
	// Base64 箭頭 Arrows
	Base64, _ = NewBase(0x2190, 0x21d0, 6)
	// Base64Gua 六十四卦 YiJing Hexagram Symbols
	Base64Gua, _ = NewBase(0x4dc0, 0x262f, 6)
	// Base32 方塊元素 Block Elements
	Base32, _ = NewBase(0x2580, 0x259f, 5)
	// Base16 漢文訓讀點 Kanbun Kundoku Den
	Base16, _ = NewBase(0x3190, 0, 4)
	// Base8Gua 八卦 YiJing Hexagram Symbols
	Base8Gua, _ = NewBase(0x2630, 0x2689, 3)
)
