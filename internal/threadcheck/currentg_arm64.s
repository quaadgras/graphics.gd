#include "textflag.h"

// func currentg() uintptr
//
// Returns the current goroutine pointer from the g register,
// which the Go runtime dedicates as R28 on arm64.
TEXT ·currentg(SB),NOSPLIT,$0-8
	MOVD g, ret+0(FP)
	RET
