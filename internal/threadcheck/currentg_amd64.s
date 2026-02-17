#include "textflag.h"

// func currentg() uintptr
//
// Returns the current goroutine pointer from R14,
// which the Go runtime (1.17+) dedicates as the g register on amd64.
TEXT ·currentg(SB),NOSPLIT,$0-8
	MOVQ R14, ret+0(FP)
	RET
