//go:build go1.26

#include "textflag.h"

// func currentm() uintptr
//
// Returns the m (OS thread) pointer by dereferencing g.m at offset 48.
// The g register (R28) holds the current goroutine pointer on arm64.
TEXT ·currentm(SB),NOSPLIT,$0-8
	MOVD g, R0
	MOVD 48(R0), R0
	MOVD R0, ret+0(FP)
	RET
