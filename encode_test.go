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

func TestEncodeBase8Gua(t *testing.T) {
	assert.Equal(t, "☱☴☲⚊", Base8Gua.EncodeString("1"))
	assert.Equal(t, "☱☴☲☳☱☰⚋", Base8Gua.EncodeString("12"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳", Base8Gua.EncodeString("123"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳☱☵☰⚊", Base8Gua.EncodeString("1234"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳☱☵☰☳☲☴⚋", Base8Gua.EncodeString("12345"))
	assert.Equal(t, "☱☴☲☳☱☰☶☳☱☵☰☳☲☴☶☶", Base8Gua.EncodeString("123456"))
}
