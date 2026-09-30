package startup

// The wasip1 build of graphics.gd is the hot-reloadable "guest" half of
// the reloads runtime (see reloads.go, built with -tags reloads): the
// native host process owns the engine and re-instantiates this module
// whenever the project's source changes. All engine calls tunnel through
// the "gd" wasm import module and all engine callbacks arrive through
// the on_* wasm exports (startup_wasip1_v2.go).
//
// Everything runs on the host's engine main thread: the guest never
// blocks. Instead of parking in Scene(), it repeatedly calls the host's
// yield import, and the host runs one engine iteration inside each call.
// Engine callbacks therefore re-enter this module only as nested calls
// while it is suspended inside a host import, which is the supported
// wasmexport reentrancy pattern — no cross-thread entry ever happens.

import (
	"iter"
	"slices"

	EngineClass "graphics.gd/classdb/Engine"
	gd "graphics.gd/internal"
	internal "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

// reloads_adopt hands an engine object that outlived the previous
// module over to this one: a fresh Go instance of the given class (this
// module's class token) is bound to the existing object. Called by the
// host after a swap, once this module's registrations have flushed.
//
//go:wasmexport reloads_adopt
func reloads_adopt(class uint64, obj uint64) uint64 {
	defer threadcheck.LeaveFrame(threadcheck.EnterFrame())
	return uint64(gd.ExtensionInstanceAdopt(gdextension.ExtensionClassID(class), gdextension.Object(obj)))
}

// reloads_ready signals that the guest's registrations are queued: the
// host binds this module's exports and replays the engine initialization
// levels into on_engine_init before returning, so the call comes back
// with the guest fully linked and its classes registered.
//
//go:wasmimport reloads ready
func reloads_ready()

// reloads_yield runs one engine iteration on the host. It reports 0 to
// keep going, 1 when the engine is shutting down and 2 when the host
// wants to swap this module out for a newer build.
//
//go:wasmimport reloads yield
func reloads_yield() uint32

const (
	reloadsKeepRunning = 0
	reloadsShutdown    = 1
	reloadsSwap        = 2
)

func (engine *engineAsSharedLibrary) Start() {
	reloads_ready()
}

func (engine *engineAsSharedLibrary) Scene() {
	code := uint32(reloadsKeepRunning)
	for code == reloadsKeepRunning {
		code = reloads_yield()
	}
	if code == reloadsSwap {
		// The module is being swapped out while the engine keeps
		// running: run NO cleanups — they would unregister classes and
		// free engine objects that live engine state still references
		// (Godot frees extension method binds with no live-instance
		// protection). The host marks everything this module owned as
		// stale instead, and the replacement module takes over.
		return
	}
	reloadsCleanup()
}

// reloadsCleanedUp records that the cleanups already ran from the scene
// level exit callback (the static host tears the engine down while this
// module is still inside reloads_yield), so Scene must not repeat them
// against an engine that is gone.
var reloadsCleanedUp bool

func reloadsCleanup() {
	if reloadsCleanedUp {
		return
	}
	reloadsCleanedUp = true
	for _, cleanup := range slices.Backward(gd.Cleanups()) {
		cleanup()
	}
	pointers.Cycle()
	pointers.Cycle()
	internal.Linked = false
}

func (engine *engineAsSharedLibrary) Rendering() iter.Seq[Float.X] {
	panic("startup.Rendering is not yet supported by the wasip1 reloads guest")
}

func init() {
	gdextension.On.Engine = gdextension.CallbacksForEngine{
		Init: func(level gdextension.InitializationLevel) {
			gd.Init(level)
			if level == 2 {
				for _, fn := range gd.StartupFunctions {
					fn()
				}
				for _, fn := range gd.PostStartupFunctions {
					fn()
				}
			}
		},
		Exit: func(level gdextension.InitializationLevel) {
			// On a module swap the engine keeps running and never fires
			// exit callbacks, so cleanup normally runs at the end of
			// Scene. When the host does forward the scene level exit,
			// the engine is shutting down underneath this module: clean
			// up now, after the scene tree is gone and before the engine
			// is, exactly where a native extension would.
			if level == 2 {
				reloadsCleanup()
			}
		},
	}
	gdextension.On.MainLoop.FirstFrame = func() {
		Callable.Cycle()
		if EngineClass.IsEditorHint() {
			editorSetup()
		}
	}
	gdextension.On.MainLoop.EveryFrame = func() {
		Callable.Cycle()
		// Nothing drains the guest's ring at callback returns (natively
		// gd.c does), so flush buffered calls once a frame, as the native
		// EveryFrame does: otherwise a QueueRedraw from Process waits for
		// the next value-returning call or a full ring (256 calls).
		ring.Main.Flush()
		pointers.Cycle()
	}
	gdextension.On.MainLoop.FinalFrame = func() {}
}
