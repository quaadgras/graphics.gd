//go:build go1.26 && (amd64 || arm64) && cgo && !O0

package gd

// The C symbols are defined in the root package's gd.c and linked across the
// whole binary; declaring them extern in the cgo preamble below lets this
// package reach them.

// extern int gd_frame_active;
// extern void *gd_sticky_call_virtual;
// extern void *gd_iterate_g0_addr(void);
// extern void gd_iterate_g0(void*);
import "C"

import (
	"unsafe"

	"graphics.gd/internal/fastcb"
	"graphics.gd/internal/pclntab"
)

type iterArgs struct {
	obj    uintptr
	method uintptr
	shape  uint64
	args   unsafe.Pointer
	result uint64
}

// StickyFastPathArmed reports whether the sticky fast path is armed.
func StickyFastPathArmed() bool { return C.gd_sticky_call_virtual != nil }

// SetFrameActiveForTest toggles the frame-active gate. Test-only: lets a test
// exercise the fast-path dispatch for a NON-allocating virtual without driving
// a real held frame (the P is not actually held, so the virtual must not
// allocate).
func SetFrameActiveForTest(v bool) {
	if v {
		C.gd_frame_active = 1
	} else {
		C.gd_frame_active = 0
	}
}

// IterationHoldingP runs one engine main-loop iteration (a bool-returning
// unsafe call, shape passed by the caller). Returns whether the engine is
// quitting. Main thread only.
//
// With a runtime patch that publishes the frame-entry hook, the frame runs
// through a STOCK cgo call marked via fastcb.FrameCgo: the goroutine is
// _Gsyscall while the engine's C code runs — scannable by the GC without
// runtime.suspendG, whose async-preempt storm against a _Grunning-in-C
// goroutine can starve the frame from ever completing under CPU load (GC
// mark then never finishes and every allocating goroutine parks: a total
// wedge). The frame's first callback re-engages residency through the
// stock path and every later callback rides the resident fast path, so a
// frame costs exactly one stock-priced transition — and fastcbFrame's
// per-frame Yield returns the goroutine to _Gsyscall for the frame delay
// and idle gap.
//
// Older patches fall back to runtime.asmcgocall — a C call on the g0 stack
// WITHOUT entersyscall, holding the P across the frame (the asmcgocall
// entry PC comes from the pclntab, see graphics.gd/internal/pclntab); the
// per-node virtual callbacks nested inside take the no-transition fast
// path and can still allocate, at the cost of the GC-starvation hazard
// above.
func IterationHoldingP(obj, method, shape uintptr) bool {
	// The iteration method takes no arguments, so args stays nil: the
	// stock-cgo path's pointer check forbids Go pointers inside a, and the
	// callframe walk never dereferences args for a zero-argument shape.
	a := iterArgs{obj: obj, method: method, shape: uint64(shape)}
	C.gd_frame_active = 1
	if fastcb.FrameCgoAvailable() {
		fastcb.FrameCgo(true)
		C.gd_iterate_g0(unsafe.Pointer(&a))
		fastcb.FrameCgo(false)
	} else {
		pclntab.Asmcgocall(C.gd_iterate_g0_addr(), unsafe.Pointer(&a))
	}
	C.gd_frame_active = 0
	return a.result != 0
}
