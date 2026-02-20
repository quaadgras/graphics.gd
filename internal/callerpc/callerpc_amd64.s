#include "textflag.h"

// func Callerpc() uintptr
//
// Returns the PC of Call[T]'s caller by walking the frame pointer chain.
// Chain: Callerpc → [ABI wrapper frame] → Call[T] → generated code.
// 0(BP) = Call[T]'s BP, 8(0(BP)) = Call[T]'s return address.
// Requires frame pointers (Go 1.21+ default on amd64).
TEXT ·Callerpc(SB),NOSPLIT,$0-8
	MOVQ 0(BP), BX
	MOVQ 8(BX), AX
	MOVQ AX, ret+0(FP)
	RET
