//go:build go1.26 && !windows

#include "textflag.h"

// func call(method, obj uintptr, args, ret unsafe.Pointer)
//
// Directly CALLs the ptrcall function pointer (PtrcallFn) using the
// System V AMD64 ABI calling convention:
//   RDI = method bind pointer
//   RSI = object pointer
//   RDX = args array (const GDExtensionConstTypePtr*)
//   RCX = return pointer (GDExtensionTypePtr)
//
// PUSHQ BP ensures 16-byte stack alignment for the C callee.
TEXT ·call(SB),NOSPLIT,$0-32
	MOVQ	method+0(FP), DI
	MOVQ	obj+8(FP), SI
	MOVQ	args+16(FP), DX
	MOVQ	ret+24(FP), CX
	MOVQ	·PtrcallFn(SB), AX
	PUSHQ	BP
	CALL	AX
	POPQ	BP
	RET
