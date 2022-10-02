package unibase2n

var (
	// Base16384 CJK Unified Ideographs
	//    see https://github.com/fumiama/base16384
	Base16384, _ = NewBase(0x4e00, 0x3d00, 14)
	// Base8192 谚文音節 Hangul Syllables
	Base8192, _ = NewBase(0xac00, 0xcc00, 13)
	// Base4096 CJK Unified Ideographs Extension A
	Base4096, _ = NewBase(0x3400, 0x4400, 12)
	// Base512 統一加拿大原住民音節文字 Unified Canadian Aboriginal Syllabics
	Base512, _ = NewBase(0x1400, 0x1600, 9)
	// Base256 谚文字母 Hangul Jamo
	Base256, _ = NewBase(0x1100, 0, 8)
	// BaseMath (256) 數學運算符 Mathematical Operators
	BaseMath, _ = NewBase(0x2200, 0, 8)
	// BaseTanWi 機種依存單位字符
	BaseTanWi, _ = NewBase(0x3300, 0, 8)
	// Base128 帶圈或括號的字母數字 Enclosed Alphanumerics
	Base128, _ = NewBase(0x2460, 0x24f4, 7)
	// BaseDevanagari 天城文 Devanagari
	BaseDevanagari, _ = NewBase(0x0900, 0x09e5, 7)
	// Base64 箭頭 Arrows
	Base64, _ = NewBase(0x2190, 0x21d0, 6)
	// Base64Gua 六十四卦 YiJing Hexagram Symbols
	Base64Gua, _ = NewBase(0x4dc0, 0x262f, 6)
	// BaseRune 盧恩字母 Runic
	BaseRune, _ = NewBase(0x16a0, 0x16e0, 6)
	// BaseMongolian 蒙古文 Mongolian
	BaseMongolian, _ = NewBase(0x1820, 0x1860, 6)
	// BaseBox 製表符 Box Drawing
	BaseBox, _ = NewBase(0x2500, 0x25f0, 7)
	// Base32 方塊元素 Block Elements
	Base32, _ = NewBase(0x2580, 0x259f, 5)
	// BaseTibetan 藏文 Tibetan
	BaseTibetan, _ = NewBase(0x0f49, 0x0f1f, 5)
	// Base16 漢文訓讀點 Kanbun Kundoku Den
	Base16, _ = NewBase(0x3190, 0, 4)
	// BaseBuginese 布吉文 Buginese
	BaseBuginese, _ = NewBase(0x1a00, 0, 4)
	// Base8Gua 八卦 YiJing Hexagram Symbols
	Base8Gua, _ = NewBase(0x2630, 0x2689, 3)
)
