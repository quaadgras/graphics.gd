#include "textflag.h"

// func Callerpc() uintptr
//
// arm64 frame record: [saved_FP, saved_LR] at FP (R29).
// 0(R29) = Call[T]'s FP, 8(0(R29)) = Call[T]'s saved LR.
TEXT ·Callerpc(SB),NOSPLIT,$0-8
	MOVD 0(R29), R0
	MOVD 8(R0), R0
	MOVD R0, ret+0(FP)
	RET
