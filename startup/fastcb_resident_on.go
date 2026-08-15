//go:build musl || archive

package startup

// Resident-callback mode for static (engine-from-Go) builds. Requires the
// fastcb runtime patch (runtime/cgocall.go overlay, bundled with and applied
// automatically by `gd build`); without it fastcb.SetResident reports false
// and the caller falls back to the stock iteration loop.
//
// The main loop drives each engine frame through IterationResident; with a
// frame-entry-aware runtime patch the frame runs as a marked stock cgo call
// (goroutine _Gsyscall while the engine's C code runs, scannable by the GC
// without suspension), residency engages at the frame's first callback, and
// fastcbFrame's per-frame Yield returns the goroutine to _Gsyscall for the
// frame delay and idle gap — one stock-priced callback per frame, exactly
// like the engine-driven (gdextension) flow. Older patches fall back to the
// raw asmcgocall frame (P held, goroutine _Grunning-in-C for the whole
// frame), which under CPU load can starve the GC's suspendG into a
// whole-process wedge — see IterationHoldingP.

import (
	"runtime"
	"sync/atomic"

	"graphics.gd/internal/fastcb"
)

// fastcbFrames counts completed resident main-loop iterations (diagnostic:
// lets a debugger distinguish a cycling frame loop from a single iteration
// that never returned).
var fastcbFrames atomic.Uint64

func fastcbScene(engine *engineAsStaticLibrary) bool {
	runtime.LockOSThread()
	if !fastcb.SetResident() {
		runtime.UnlockOSThread()
		return false
	}
	if fastcb.FrameCgoAvailable() {
		// Frame-entry-aware patch: let fastcbFrame yield once per frame so
		// the goroutine spends the frame delay and idle gap in _Gsyscall
		// (see the file comment). Residency is already engaged, so
		// fastcbFrame reduces to the per-frame fastcb.Yield.
		fastcbFrameDriving = true
		fastcbEngaged = true
	} else {
		// Raw-asmcgocall fallback: the per-frame driver must stay out of
		// the way (a yield's deferred transition needs a stock base entry,
		// which this path does not have — fastcbSavedPC stays 0).
		fastcbFrameDriving = false
	}
	for !engine.Library.IterationResident() {
		fastcbFrames.Add(1)
	}
	fastcbFrameDriving = false
	fastcb.ClearResident()
	runtime.UnlockOSThread()
	return true
}
