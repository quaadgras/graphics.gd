//go:build go1.26

#include "textflag.h"

// func call(method, obj uintptr, args, ret unsafe.Pointer)
//
// Directly calls the ptrcall function pointer (PtrcallFn) using the
// AAPCS64 calling convention:
//   R0 = method bind pointer
//   R1 = object pointer
//   R2 = args array (const GDExtensionConstTypePtr*)
//   R3 = return pointer (GDExtensionTypePtr)
//
// LR (R30) is saved/restored across BLR since the branch-with-link
// overwrites the link register.
TEXT ·call(SB),NOSPLIT,$0-32
	MOVD	method+0(FP), R0
	MOVD	obj+8(FP), R1
	MOVD	args+16(FP), R2
	MOVD	ret+24(FP), R3
	MOVD	·PtrcallFn(SB), R9
	SUB	$16, RSP, RSP
	MOVD	R30, (RSP)
	BLR	R9
	MOVD	(RSP), R30
	ADD	$16, RSP, RSP
	RET
