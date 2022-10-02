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

// enc64blk2(mask uint64, in, out []byte)
//    len(in)!=0, len(out)==len(in)*8
TEXT ·enc64blk2(SB), NOSPLIT, $0-56
    // load mask
    MOVQ ·mask+0(FP), BX
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
    SHLL $10, AX
    MOVL AX, DX
    ANDL $0x00030000, DX
    RORQ $14, AX
    MOVL AX, R8
    ANDL $0x00000003, R8
    ORL  R8, DX
    SHLQ $32, DX
    ROLQ $4, AX
    MOVL AX, R8
    ANDL $0x00000003, R8
    ORQ  R8, DX
    SHLL $14, AX
    ANDL $0x00030000, AX
    ORQ  AX, DX
    // add mask
    LEAQ 0(DX)(BX*1), AX
    BSWAPQ AX
    STOSQ
    LOOP lop
    RET

// enc32blk4(mask uint32, in, out []byte)
//    len(in)!=0, len(out)==len(in)*4
TEXT ·enc32blk4(SB), NOSPLIT, $0-56
    // load mask
    MOVQ ·mask+0(FP), BX
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
    // 8 -> 32
    MOVL AX, DX
    ANDL $0x0f, DX
    SHLL $12, AX
    ANDL $0x000f0000, AX
    ORL  AX, DX
    // add mask
    LEAL 0(DX)(BX*1), AX
    BSWAPL AX
    STOSL
    LOOP lop
    RET

// func enc16blk8(mask uint16, in, out []byte)
//    len(in)!=0, len(out)==len(in)*2
TEXT ·enc16blk8(SB), NOSPLIT, $0-56
    // load mask
    MOVQ ·mask+0(FP), BX
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
    // add mask
    LEAW 0(AX)(BX*1), AX
    RORW $8, AX
    STOSW
    LOOP lop
    RET
