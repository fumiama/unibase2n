package unibase2n

import (
	"golang.org/x/text/encoding/unicode"
)

var format = unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)

// UTF16BE2UTF8 to display the result as string
func UTF16BE2UTF8(b []byte) ([]byte, error) {
	return format.NewDecoder().Bytes(b)
}

// UTF82UTF16BE to decode from string
func UTF82UTF16BE(b []byte) ([]byte, error) {
	return format.NewEncoder().Bytes(b)
}

func (bs *Base) EncodeToString(b []byte) string {
	out, err := UTF16BE2UTF8(bs.Encode(b))
	if err != nil {
		return ""
	}
	return BytesToString(out)
}

func (bs *Base) EncodeFromString(s string) []byte {
	return bs.Encode(StringToBytes(s))
}

func (bs *Base) EncodeString(s string) string {
	out, err := UTF16BE2UTF8(bs.Encode(StringToBytes(s)))
	if err != nil {
		return ""
	}
	return BytesToString(out)
}

func (bs *Base) DecodeToString(d []byte) string {
	return BytesToString(bs.Decode(d))
}

func (bs *Base) DecodeFromString(s string) []byte {
	d, err := UTF82UTF16BE(StringToBytes(s))
	if err != nil {
		return nil
	}
	return bs.Decode(d)
}

func (bs *Base) DecodeString(s string) string {
	d, err := UTF82UTF16BE(StringToBytes(s))
	if err != nil {
		return ""
	}
	return BytesToString(bs.Decode(d))
}
