package unibase2n

import (
	"encoding/binary"
	"errors"
	"unsafe"
)

type Pack uint64

var (
	ErrInvalidPack = errors.New("invalid pack")
)

// New base2n from a packed config
func New(pack Pack) (bs Base, err error) {
	if pack&0xff == 0 && pack&0xff000000_00000000 == 0 {
		err = ErrInvalidPack
		return
	}
	ismele := isLittleEndian()
	isitle := pack&0xffffff != 0
	if ismele == isitle { // same endian
		*(*Pack)(unsafe.Pointer(&bs)) = pack
	} else {
		field := (*[8]byte)(unsafe.Pointer(&pack))
		if isitle { // packed in little endian but I am big
			bs.off = binary.BigEndian.Uint16(field[6:8])
			bs.til = binary.BigEndian.Uint16(field[4:6])
		} else { // packed in big endian but I am little
			bs.off = binary.LittleEndian.Uint16(field[6:8])
			bs.til = binary.LittleEndian.Uint16(field[4:6])
		}
		bs.bit = field[3]
	}
	err = bs.check()
	return
}

// Pack this config into an integer
func (bs Base) Pack() Pack {
	return *(*Pack)(unsafe.Pointer(&bs))
}

func isLittleEndian() bool {
	n := 0x1234
	f := *((*byte)(unsafe.Pointer(&n)))
	return (f ^ 0x34) == 0
}
