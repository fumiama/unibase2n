package unibase2n

import (
	"errors"
)

// Base has an encoding buffer thus should not be copied.
//    total size: 8 bytes
type Base struct {
	off uint16  // starting offset (0 is not permitted)
	til uint16  // remianing indicator starting offset
	bit uint8   // 2^bit, max is 15 (32768)
	_   [3]byte // always 0, indicates the byte order
}

var (
	ErrInvalidBitSize   = errors.New("bit size >= 16 or == 0")
	ErrZeroOffsetStart  = errors.New("zero offset start")
	ErrOffsetOverflow   = errors.New("offset overflow")
	ErrTailOverflow     = errors.New("tail overflow")
	ErrTailInCodingArea = errors.New("tail in coding area")
)

// NewBase generates a new base2n config
func NewBase(off, til uint16, bit uint8) (bs Base, err error) {
	bs = Base{
		off: off,
		til: til,
		bit: bit,
	}
	err = bs.check()
	return
}

func (bs Base) check() error {
	off := bs.off
	til := bs.til
	bit := bs.bit
	if off == 0 {
		return ErrZeroOffsetStart
	}
	if bit >= 16 || bit == 0 {
		return ErrInvalidBitSize
	}
	offe := uint32(off) + 1<<bit // [off, offe)
	if offe > 0x10000 {
		return ErrOffsetOverflow
	}
	tile := uint32(til) // [til, tile)
	if bit > 8 && bit%2 == 0 {
		tile += uint32(bit / 2)
	} else {
		tile += uint32(bit)
	}
	if tile > 0x10000 {
		return ErrTailOverflow
	}
	if tile > uint32(off) && tile <= offe {
		return ErrTailInCodingArea
	}
	return nil
}
