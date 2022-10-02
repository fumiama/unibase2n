//go:build amd64
// +build amd64

#include "textflag.h"

// enc16blk1(mask uint16, in, out []byte)
//    for bit 1 (actual enc128blk1)
TEXT ·enc16blk1(SB), NOSPLIT, $0-56
    // create mask
    MOVQ ·mask+0(FP), BX
    ANDQ $0xffff, BX
    BSWAPQ BX
    MOVQ BX, AX
    SHRQ $16, AX
    ORQ  AX, BX
    SHRQ $16, AX
    ORQ  AX, BX
    SHRQ $16, AX
    ORQ  AX, BX
    // load source addr
    MOVQ ·in+8(FP), SI
    // load source len
    MOVQ ·inlen+16(FP), CX
    // load dest addr
    MOVQ ·out+32(FP), DI
    // go forward
    CLD
lop:
    LODSB
    // 8 -> 64
    XORQ DX, DX
    SHLB $1, AX
    SETCS DX
    SHLQ $16, DX
    SHLB $1, AX
    SETCS DX
    SHLQ $16, DX
    SHLB $1, AX
    SETCS DX
    SHLQ $16, DX
    SHLB $1, AX
    SETCS DX
    // add mask
    MOVQ AX, R8
    LEAQ 0(DX)(BX*1), AX
    BSWAPQ AX
    STOSQ
    // 8 -> 64
    XORQ AX, AX
    SHLB $1, R8
    SETCS AX
    SHLQ $16, AX
    SHLB $1, R8
    SETCS AX
    SHLQ $16, AX
    SHLB $1, R8
    SETCS AX
    SHLQ $16, AX
    SHLB $1, R8
    SETCS AX
    // add mask
    ADDQ BX, AX
    BSWAPQ AX
    STOSQ
    LOOP lop
    RET
