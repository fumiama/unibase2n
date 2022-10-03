//go:build amd64
// +build amd64

#include "textflag.h"

// dec128blk1(mask uint128be, in, out []byte)
//    len(in)>0, len(in)%16==0, len(out)==len(in)/16
TEXT ·dec128blk1(SB), NOSPLIT, $0-64
    MOVQ ·mask+0(FP), DX
    BSWAPQ DX
    MOVQ ·in+16(FP), SI
    MOVQ ·in+24(FP), CX
    SHRQ $4, CX
    MOVQ ·in+40(FP), DI
    // go forward
    CLD
lop:
    LODSQ
    BSWAPQ AX
    SUBQ DX, AX
    MOVQ AX, BX
    ANDB $1, AX
    RORQ $1, AX
    SHRQ $17, BX
    SETCS AX
    RORQ $1, AX
    SHRQ $16, BX
    SETCS AX
    RORQ $1, AX
    SHRQ $16, BX
    SETCS AX

    ROLQ $7, AX
    MOVQ AX, R8

    LODSQ
    BSWAPQ AX
    SUBQ DX, AX
    MOVQ AX, BX
    ANDB $1, AX
    RORQ $1, AX
    SHRQ $17, BX
    SETCS AX
    RORQ $1, AX
    SHRQ $16, BX
    SETCS AX
    RORQ $1, AX
    SHRQ $16, BX
    SETCS AX

    ROLQ $3, AX
    ORQ  R8, AX

    STOSB
    LOOP lop
    RET
