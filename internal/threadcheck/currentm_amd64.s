//go:build go1.26

#include "textflag.h"

// func currentm() uintptr
//
// Returns the m (OS thread) pointer by dereferencing g.m at offset 48.
// R14 holds the current g pointer on amd64 (Go 1.17+).
TEXT ·currentm(SB),NOSPLIT,$0-8
	MOVQ 48(R14), AX
	MOVQ AX, ret+0(FP)
	RET
