//go:build go1.26 && windows

#include "textflag.h"

// func call(method, obj uintptr, args, ret unsafe.Pointer)
//
// Directly CALLs the ptrcall function pointer (PtrcallFn) using the
// Microsoft x64 ABI calling convention:
//   RCX = method bind pointer
//   RDX = object pointer
//   R8  = args array (const GDExtensionConstTypePtr*)
//   R9  = return pointer (GDExtensionTypePtr)
//
// 32 bytes of shadow space are allocated before the CALL as required
// by the Microsoft x64 ABI. PUSHQ BP ensures 16-byte stack alignment.
TEXT ·call(SB),NOSPLIT,$0-32
	MOVQ	method+0(FP), CX
	MOVQ	obj+8(FP), DX
	MOVQ	args+16(FP), R8
	MOVQ	ret+24(FP), R9
	MOVQ	·PtrcallFn(SB), AX
	PUSHQ	BP
	SUBQ	$32, SP
	CALL	AX
	ADDQ	$32, SP
	POPQ	BP
	RET
