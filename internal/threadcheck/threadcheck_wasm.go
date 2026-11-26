//go:build wasm

package threadcheck

// On wasm there is a single OS thread, so thread identity cannot tell the
// frame-driving code apart from free-running goroutines. Goroutine identity
// can: the wasm g register is read by currentg (see currentg_wasm.s) at a
// cost comparable to the native g-register thread check, and with no other
// threads there is no need for atomics on the variables below.

// mainG identifies the program's main goroutine. Package initialisers run on
// it, before any other goroutine can exist. The main goroutine is
// frame-aligned by construction: in Rendering()/Scene() mode the frame
// handshake lock-steps the engine to it (a frame cannot advance, and so the
// per-frame collection cannot run, while user main-loop code holds the frame
// slot), so its wrappers may safely be frame-temporaries.
var mainG = currentg()

// frameG identifies the goroutine currently executing an engine→Go export
// (see EnterFrame): while an engine callback runs, its goroutine is the
// frame-driving goroutine. Zero outside any callback. If callback code
// blocks, goroutines scheduled in the meantime do not match frameG (their g
// differs, and a parked g cannot be reused) and are treated as free-running.
var frameG uintptr

func Init() {}

// Main always returns true on wasm, which is single-threaded: every caller
// may dispatch directly into the engine.
func Main() bool { return true }

// EnterFrame records the calling goroutine as the frame-driving goroutine,
// returning the previous value for LeaveFrame. Called on entry to every
// engine→Go wasm export; enter/leave pairs nest like a stack, so re-entrant
// callbacks restore correctly.
func EnterFrame() uintptr {
	prev := frameG
	frameG = currentg()
	frameDepth++
	return prev
}

// LeaveFrame restores the frame-driving goroutine recorded by the matching
// EnterFrame.
func LeaveFrame(prev uintptr) {
	frameG = prev
	frameDepth--
}

// frameDepth counts the engine->Go exports that are running (see EnterFrame).
var frameDepth int

// FrameDepth reports how many engine->Go callbacks are running: 0 outside any
// callback, 1 inside a callback, 2 inside a callback made by an engine call
// from a callback, and so on. internal/gdmemory uses it to give the engine
// calls of each depth their own argument buffers.
func FrameDepth() int { return frameDepth }

// FrameTemporaries reports whether the calling goroutine is frame-aligned:
// either the goroutine executing the current engine callback, or the
// program's main goroutine. Wrappers created anywhere else are anchored to
// the Go garbage collector (see gd.Wrap*), because the per-frame collection
// is not aware of goroutine timelines.
func FrameTemporaries() bool {
	g := currentg()
	return g == frameG || g == mainG
}

// Mark is a no-op on wasm, which is single-threaded.
func Mark() {}

// EnterCall is a no-op on wasm, which is single-threaded.
func EnterCall() {}

// LeaveCall is a no-op on wasm, which is single-threaded.
func LeaveCall() {}

// Engine always returns true on wasm, which is single-threaded.
func Engine() bool { return true }
