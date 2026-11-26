//go:build wasm && !(gd && wasm.gc)

#include "textflag.h"

// currentg returns the g register: the current goroutine's g pointer.
TEXT ·currentg(SB), NOSPLIT, $0-8
	MOVD g, ret+0(FP)
	RET
