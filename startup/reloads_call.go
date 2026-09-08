//go:build reloads && !wasip1

package startup

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"

	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/sys"
)

// reloadsTrace enables the GD_RELOADS_TRACE diagnostics.
var reloadsTrace = os.Getenv("GD_RELOADS_TRACE") != ""

// reloadsCallPool holds the api.Function handles available for one guest
// export. wazero forbids calling a Function again before its previous
// call has returned — each handle owns the wasm stack the call runs on,
// so a re-entrant call would overwrite the frames of the outer one and
// corrupt its return — but engine callbacks nest: a virtual method that
// calls back into the engine can have another virtual dispatched to the
// guest before it returns (a ResourceFormatLoader answering _exists
// while the engine is still inside _recognize_path, for example). Every
// level of nesting therefore takes its own handle, resolved from the
// same export and kept for reuse.
type reloadsCallPool struct {
	free []api.Function
}

// reloadsCallPools is keyed by the handle the generated bindings resolved
// for each export (reloadsBindGuest); it is cleared on every guest swap.
var reloadsCallPools = map[api.Function]*reloadsCallPool{}

var reloadsCallDepth int

// reloadsGuestDead is set when the live module exits on its own — a Go
// panic (exit code 2, which is also the swap code, so an exit is told
// apart from a swap by reloadsExitCode still being reloadsKeepRunning)
// or an os.Exit — inside an engine callback, while the host is still
// deep in the engine frame that made the call and the guest's _start is
// parked in the yield below it. wazero does not refuse calls into a
// closed module: it runs them and only reports the closure once they
// return, so every further callback of that frame would execute the
// dead guest's runtime mid-exit (on its crash stack, its tables torn
// down) and hand the engine garbage, and returning from the yield would
// resume _start the same way; the host then segfaulted inside the
// engine. Instead the remaining calls are answered inertly, the yield
// unwinds _start with the guest's own exit error, and the session drops
// the build and waits for the next successful one — a panic in the
// project no longer costs the editor, like a compile error.
var reloadsGuestDead atomic.Bool
var reloadsGuestExit atomic.Pointer[sys.ExitError]

func reloadsGuestDied(exit *sys.ExitError) {
	if reloadsGuestDead.Swap(true) {
		return
	}
	reloadsGuestExit.Store(exit)
	reloadsGuest.Store(nil) // forwarders answer inertly from here on
	fmt.Fprintf(os.Stderr, "graphics.gd: the project exited (%s) inside an engine callback; hot reload resumes with the next build that compiles\n", exit)
}

// reloadsUnwindDeadGuest panics with the dead guest's exit error, which
// wazero passes through unwrapped as the result of the module's
// instantiation — the same way the guest's own proc_exit unwinds — so a
// control host function (ready, yield) never resumes a closed module.
func reloadsUnwindDeadGuest() {
	if exit := reloadsGuestExit.Load(); exit != nil {
		panic(exit)
	}
}

func reloadsCallPooled(fn api.Function, stack []uint64) {
	if reloadsGuestDead.Load() {
		clear(stack) // zero results: the engine reads them as the answer
		return
	}
	pool := reloadsCallPools[fn]
	if pool == nil {
		pool = &reloadsCallPool{free: []api.Function{fn}}
		reloadsCallPools[fn] = pool
	}
	call := fn
	if n := len(pool.free); n > 0 {
		call = pool.free[n-1]
		pool.free = pool.free[:n-1]
	} else if g := reloadsGuest.Load(); g != nil {
		if names := fn.Definition().ExportNames(); len(names) > 0 {
			if spare := g.module.ExportedFunction(names[0]); spare != nil {
				call = spare
			}
		}
	}
	reloadsCallDepth++
	if reloadsTrace {
		fmt.Fprintf(os.Stderr, "graphics.gd: guest call %v depth=%d\n", fn.Definition().ExportNames(), reloadsCallDepth)
	}
	err := call.CallWithStack(reloadsCtx, stack)
	reloadsCallDepth--
	if call != fn || len(pool.free) == 0 {
		pool.free = append(pool.free, call)
	}
	if err != nil {
		var exit *sys.ExitError
		if errors.As(err, &exit) && reloadsExitCode.Load() == reloadsKeepRunning {
			reloadsGuestDied(exit)
			clear(stack)
			return
		}
		os.Stderr.WriteString("graphics.gd/startup: reloads guest call failed: " + err.Error() + "\n")
	}
}
