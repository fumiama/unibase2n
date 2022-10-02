package unibase2n

import (
	"math/rand"
	"testing"

	base14 "github.com/fumiama/go-base16384"
	"github.com/stretchr/testify/assert"
)

func TestEncodeBase16384(t *testing.T) {
	var buf [4096]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i <= 4096; i++ {
		if !assert.Equal(t, base14.Encode(buf[:i]), Base16384.Encode(buf[:i])) {
			t.Fatal("expect:", base14.EncodeToString(buf[:i]), "actual:", Base16384.EncodeToString(buf[:i]))
		}
	}
}

func TestEncodeBase8192(t *testing.T) {
	assert.Equal(t, "눠찁", Base8192.EncodeString("1"))
	assert.Equal(t, "눦됀찂", Base8192.EncodeString("12"))
	assert.Equal(t, "눦듌찃", Base8192.EncodeString("123"))
	assert.Equal(t, "눦듌였찄", Base8192.EncodeString("1234"))
	assert.Equal(t, "눦듌옚밀찅", Base8192.EncodeString("12345"))
	assert.Equal(t, "눦듌옚뽠찆", Base8192.EncodeString("123456"))
	assert.Equal(t, "눦듌옚뽣먀찇", Base8192.EncodeString("1234567"))
	assert.Equal(t, "눦듌옚뽣며찈", Base8192.EncodeString("12345678"))
	assert.Equal(t, "눦듌옚뽣며멀찉", Base8192.EncodeString("123456789"))
	assert.Equal(t, "눦듌옚뽣며멌가찊", Base8192.EncodeString("1234567890"))
	assert.Equal(t, "눦듌옚뽣며멌궈찋", Base8192.EncodeString("12345678901"))
	if !assert.Equal(t, "눦듌옚뽣며멌궉븀찌", Base8192.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "눦듌옚뽣며멌궉븳", Base8192.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "눦듌옚뽣며멌궉븳늀찁", Base8192.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBase4096(t *testing.T) {
	assert.Equal(t, "㜐䐁", Base4096.EncodeString("1"))
	assert.Equal(t, "㜓㘀䐂", Base4096.EncodeString("12"))
	assert.Equal(t, "㜓㘳㐀䐃", Base4096.EncodeString("123"))
	assert.Equal(t, "㜓㘳㝀䐄", Base4096.EncodeString("1234"))
	assert.Equal(t, "㜓㘳㝃㤀䐅", Base4096.EncodeString("12345"))
	assert.Equal(t, "㜓㘳㝃㤶", Base4096.EncodeString("123456"))
	assert.Equal(t, "㜓㘳㝃㤶㝰䐁", Base4096.EncodeString("1234567"))
	assert.Equal(t, "㜓㘳㝃㤶㝳㰀䐂", Base4096.EncodeString("12345678"))
	assert.Equal(t, "㜓㘳㝃㤶㝳㰹㐀䐃", Base4096.EncodeString("123456789"))
	assert.Equal(t, "㜓㘳㝃㤶㝳㰹㜀䐄", Base4096.EncodeString("1234567890"))
	assert.Equal(t, "㜓㘳㝃㤶㝳㰹㜃㔀䐅", Base4096.EncodeString("12345678901"))
	if !assert.Equal(t, "㜓㘳㝃㤶㝳㰹㜃㔲", Base4096.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
}

func TestEncodeBase512(t *testing.T) {
	assert.Equal(t, "ᑢᘁ", Base512.EncodeString("1"))
	assert.Equal(t, "ᑢᓈᘂ", Base512.EncodeString("12"))
	assert.Equal(t, "ᑢᓈᖘᘃ", Base512.EncodeString("123"))
	assert.Equal(t, "ᑢᓈᖙᕀᘄ", Base512.EncodeString("1234"))
	assert.Equal(t, "ᑢᓈᖙᕃᒠᘅ", Base512.EncodeString("12345"))
	assert.Equal(t, "ᑢᓈᖙᕃᒦᖀᘆ", Base512.EncodeString("123456"))
	assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖀᘇ", Base512.EncodeString("1234567"))
	assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜ᐀ᘈ", Base512.EncodeString("12345678"))
	assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜᐹ", Base512.EncodeString("123456789"))
	assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜᐹᑠᘁ", Base512.EncodeString("1234567890"))
	assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜᐹᑠᓄᘂ", Base512.EncodeString("12345678901"))
	if !assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜᐹᑠᓄᖐᘃ", Base512.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜᐹᑠᓄᖑᔰᘄ", Base512.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "ᑢᓈᖙᕃᒦᖍᖜᐹᑠᓄᖑᔳᒀᘅ", Base512.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBase256(t *testing.T) {
	assert.Equal(t, "ᄱ", Base256.EncodeString("1"))
	assert.Equal(t, "ᄱᄲ", Base256.EncodeString("12"))
	assert.Equal(t, "ᄱᄲᄳ", Base256.EncodeString("123"))
	assert.Equal(t, "ᄱᄲᄳᄴ", Base256.EncodeString("1234"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵ", Base256.EncodeString("12345"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵᄶ", Base256.EncodeString("123456"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷ", Base256.EncodeString("1234567"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸ", Base256.EncodeString("12345678"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸᄹ", Base256.EncodeString("123456789"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸᄹᄰ", Base256.EncodeString("1234567890"))
	assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸᄹᄰᄱ", Base256.EncodeString("12345678901"))
	if !assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸᄹᄰᄱᄲ", Base256.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸᄹᄰᄱᄲᄳ", Base256.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "ᄱᄲᄳᄴᄵᄶᄷᄸᄹᄰᄱᄲᄳᄴ", Base256.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBaseTanWi(t *testing.T) {
	assert.Equal(t, "㌱", BaseTanWi.EncodeString("1"))
	assert.Equal(t, "㌱㌲", BaseTanWi.EncodeString("12"))
	assert.Equal(t, "㌱㌲㌳", BaseTanWi.EncodeString("123"))
	assert.Equal(t, "㌱㌲㌳㌴", BaseTanWi.EncodeString("1234"))
	assert.Equal(t, "㌱㌲㌳㌴㌵", BaseTanWi.EncodeString("12345"))
	assert.Equal(t, "㌱㌲㌳㌴㌵㌶", BaseTanWi.EncodeString("123456"))
	assert.Equal(t, "㌱㌲㌳㌴㌵㌶㌷", BaseTanWi.EncodeString("1234567"))
	assert.Equal(t, "㌱㌲㌳㌴㌵㌶㌷㌸", BaseTanWi.EncodeString("12345678"))
	assert.Equal(t, "㌱㌲㌳㌴㌵㌶㌷㌸㌹", BaseTanWi.EncodeString("123456789"))
	assert.Equal(t, "㌱㌲㌳㌴㌵㌶㌷㌸㌹㌰", BaseTanWi.EncodeString("1234567890"))
	assert.Equal(t, "㌱㌲㌳㌴㌵㌶㌷㌸㌹㌰㌱", BaseTanWi.EncodeString("12345678901"))
}

func TestEncodeBase128(t *testing.T) {
	assert.Equal(t, "⑸⒠⓵", Base128.EncodeString("1"))
	assert.Equal(t, "⑸⒬⒠⓶", Base128.EncodeString("12"))
	assert.Equal(t, "⑸⒬⒦⒐⓷", Base128.EncodeString("123"))
	assert.Equal(t, "⑸⒬⒦⒓⒀⓸", Base128.EncodeString("1234"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴⓹", Base128.EncodeString("12345"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⓺", Base128.EncodeString("123456"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗", Base128.EncodeString("1234567"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼①⓵", Base128.EncodeString("12345678"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼⑮⒀⓶", Base128.EncodeString("123456789"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼⑮⒆①⓷", Base128.EncodeString("1234567890"))
	assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼⑮⒆④⑨⓸", Base128.EncodeString("12345678901"))
	if !assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼⑮⒆④⑩⒨⓹", Base128.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼⑮⒆④⑩⒨Ⓠ⓺", Base128.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "⑸⒬⒦⒓⒁⒴Ⓦ⒗⑼⑮⒆④⑩⒨Ⓠ⒔", Base128.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBaseDevanagari(t *testing.T) {
	assert.Equal(t, "घी০", BaseDevanagari.EncodeString("1"))
	assert.Equal(t, "घौी১", BaseDevanagari.EncodeString("12"))
	assert.Equal(t, "घौॆर২", BaseDevanagari.EncodeString("123"))
	assert.Equal(t, "घौॆळठ৩", BaseDevanagari.EncodeString("1234"))
	assert.Equal(t, "घौॆळड॔৪", BaseDevanagari.EncodeString("12345"))
	assert.Equal(t, "घौॆळड॔६৫", BaseDevanagari.EncodeString("123456"))
	assert.Equal(t, "घौॆळड॔६ष", BaseDevanagari.EncodeString("1234567"))
	assert.Equal(t, "घौॆळड॔६षजऀ০", BaseDevanagari.EncodeString("12345678"))
	assert.Equal(t, "घौॆळड॔६षजऎठ১", BaseDevanagari.EncodeString("123456789"))
	assert.Equal(t, "घौॆळड॔६षजऎदऀ২", BaseDevanagari.EncodeString("1234567890"))
	assert.Equal(t, "घौॆळड॔६षजऎदःई৩", BaseDevanagari.EncodeString("12345678901"))
	if !assert.Equal(t, "घौॆळड॔६षजऎदःउै৪", BaseDevanagari.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "घौॆळड॔६षजऎदःउै०৫", BaseDevanagari.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "घौॆळड॔६षजऎदःउै०ऴ", BaseDevanagari.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBase64Gua(t *testing.T) {
	assert.Equal(t, "䷌䷐☰", Base64Gua.EncodeString("1"))
	assert.Equal(t, "䷌䷓䷈☱", Base64Gua.EncodeString("12"))
	assert.Equal(t, "䷌䷓䷈䷳", Base64Gua.EncodeString("123"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷀☰", Base64Gua.EncodeString("1234"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔☱", Base64Gua.EncodeString("12345"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶", Base64Gua.EncodeString("123456"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷰☰", Base64Gua.EncodeString("1234567"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠☱", Base64Gua.EncodeString("12345678"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠䷹", Base64Gua.EncodeString("123456789"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠䷹䷌䷀☰", Base64Gua.EncodeString("1234567890"))
	assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠䷹䷌䷃䷄☱", Base64Gua.EncodeString("12345678901"))
	if !assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠䷹䷌䷃䷄䷲", Base64Gua.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠䷹䷌䷃䷄䷲䷌䷰☰", Base64Gua.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "䷌䷓䷈䷳䷍䷃䷔䷶䷍䷳䷠䷹䷌䷃䷄䷲䷌䷳䷐☱", Base64Gua.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBaseRune(t *testing.T) {
	assert.Equal(t, "ᚬᚰᛡ", BaseRune.EncodeString("1"))
	assert.Equal(t, "ᚬᚳᚨᛢ", BaseRune.EncodeString("12"))
	assert.Equal(t, "ᚬᚳᚨᛓ", BaseRune.EncodeString("123"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚠᛡ", BaseRune.EncodeString("1234"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛢ", BaseRune.EncodeString("12345"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖ", BaseRune.EncodeString("123456"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛐᛡ", BaseRune.EncodeString("1234567"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛢ", BaseRune.EncodeString("12345678"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛙ", BaseRune.EncodeString("123456789"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛙᚬᚠᛡ", BaseRune.EncodeString("1234567890"))
	assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛙᚬᚣᚤᛢ", BaseRune.EncodeString("12345678901"))
	if !assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛙᚬᚣᚤᛒ", BaseRune.EncodeString("123456789012")) {
		t.Fatal("123456789012")
	}
	if !assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛙᚬᚣᚤᛒᚬᛐᛡ", BaseRune.EncodeString("1234567890123")) {
		t.Fatal("1234567890123")
	}
	if !assert.Equal(t, "ᚬᚳᚨᛓᚭᚣᚴᛖᚭᛓᛀᛙᚬᚣᚤᛒᚬᛓᚰᛢ", BaseRune.EncodeString("12345678901234")) {
		t.Fatal("12345678901234")
	}
}

func TestEncodeBaseMongolian(t *testing.T) {
	assert.Equal(t, "ᠬᠰᡡ", BaseMongolian.EncodeString("1"))
	assert.Equal(t, "ᠬᠳᠨᡢ", BaseMongolian.EncodeString("12"))
	assert.Equal(t, "ᠬᠳᠨᡓ", BaseMongolian.EncodeString("123"))
	assert.Equal(t, "ᠬᠳᠨᡓᠭᠠᡡ", BaseMongolian.EncodeString("1234"))
	assert.Equal(t, "ᠬᠳᠨᡓᠭᠣᠴᡢ", BaseMongolian.EncodeString("12345"))
	assert.Equal(t, "ᠬᠳᠨᡓᠭᠣᠴᡖ", BaseMongolian.EncodeString("123456"))
	assert.Equal(t, "ᠬᠳᠨᡓᠭᠣᠴᡖᠭᡐᡡ", BaseMongolian.EncodeString("1234567"))
	assert.Equal(t, "ᠬᠳᠨᡓᠭᠣᠴᡖᠭᡓᡀᡢ", BaseMongolian.EncodeString("12345678"))
	assert.Equal(t, "ᠬᠳᠨᡓᠭᠣᠴᡖᠭᡓᡀᡙ", BaseMongolian.EncodeString("123456789"))
}

func TestEncodeBase32(t *testing.T) {
	assert.Equal(t, "▆▄■", Base32.EncodeString("1"))
	assert.Equal(t, "▆▄▙▀□", Base32.EncodeString("12"))
	assert.Equal(t, "▆▄▙▃▆▢", Base32.EncodeString("123"))
	assert.Equal(t, "▆▄▙▃▆▍▀▣", Base32.EncodeString("1234"))
	assert.Equal(t, "▆▄▙▃▆▍▁▕", Base32.EncodeString("12345"))
	assert.Equal(t, "▆▄▙▃▆▍▁▕▆▘■", Base32.EncodeString("123456"))
	assert.Equal(t, "▆▄▙▃▆▍▁▕▆▘▛▐□", Base32.EncodeString("1234567"))
	assert.Equal(t, "▆▄▙▃▆▍▁▕▆▘▛▓▐▢", Base32.EncodeString("12345678"))
	assert.Equal(t, "▆▄▙▃▆▍▁▕▆▘▛▓▐▎█▣", Base32.EncodeString("123456789"))
}

func TestEncodeBaseTibetan(t *testing.T) {
	assert.Equal(t, "ཏཌྷ༠", BaseTibetan.EncodeString("1"))
	assert.Equal(t, "ཏཌྷརཉ༡", BaseTibetan.EncodeString("12"))
	assert.Equal(t, "ཏཌྷརཌཏ༢", BaseTibetan.EncodeString("123"))
	assert.Equal(t, "ཏཌྷརཌཏབཉ༣", BaseTibetan.EncodeString("1234"))
	assert.Equal(t, "ཏཌྷརཌཏབཊཞ", BaseTibetan.EncodeString("12345"))
	assert.Equal(t, "ཏཌྷརཌཏབཊཞཏཡ༠", BaseTibetan.EncodeString("123456"))
	assert.Equal(t, "ཏཌྷརཌཏབཊཞཏཡཤཙ༡", BaseTibetan.EncodeString("1234567"))
	assert.Equal(t, "ཏཌྷརཌཏབཊཞཏཡཤཛྷཙ༢", BaseTibetan.EncodeString("12345678"))
	assert.Equal(t, "ཏཌྷརཌཏབཊཞཏཡཤཛྷཙབྷད༣", BaseTibetan.EncodeString("123456789"))
}

func TestEncodeBase16(t *testing.T) {
	assert.Equal(t, "㆓㆑", Base16.EncodeString("1"))
	assert.Equal(t, "㆓㆑㆓㆒", Base16.EncodeString("12"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓", Base16.EncodeString("123"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓㆓㆔", Base16.EncodeString("1234"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓㆓㆔㆓㆕", Base16.EncodeString("12345"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓㆓㆔㆓㆕㆓㆖", Base16.EncodeString("123456"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓㆓㆔㆓㆕㆓㆖㆓㆗", Base16.EncodeString("1234567"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓㆓㆔㆓㆕㆓㆖㆓㆗㆓㆘", Base16.EncodeString("12345678"))
	assert.Equal(t, "㆓㆑㆓㆒㆓㆓㆓㆔㆓㆕㆓㆖㆓㆗㆓㆘㆓㆙", Base16.EncodeString("123456789"))
}

func TestEncodeBaseBuginese(t *testing.T) {
	assert.Equal(t, "ᨃᨁ", BaseBuginese.EncodeString("1"))
	assert.Equal(t, "ᨃᨁᨃᨂ", BaseBuginese.EncodeString("12"))
	assert.Equal(t, "ᨃᨁᨃᨂᨃᨃ", BaseBuginese.EncodeString("123"))
	assert.Equal(t, "ᨃᨁᨃᨂᨃᨃᨃᨄ", BaseBuginese.EncodeString("1234"))
}

func TestEncodeBase8Gua(t *testing.T) {
	assert.Equal(t, "☱☴☲⚊", Base8Gua.EncodeString("1"))
	assert.Equal(t, "☱☴☲☳☱☰⚋", Base8Gua.EncodeString("12"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳", Base8Gua.EncodeString("123"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳☱☵☰⚊", Base8Gua.EncodeString("1234"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳☱☵☰☳☲☴⚋", Base8Gua.EncodeString("12345"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳☱☵☰☳☲☴☶☶", Base8Gua.EncodeString("123456"))
}
