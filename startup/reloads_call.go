//go:build reloads && !wasip1

package startup

import (
	"fmt"
	"os"

	"github.com/tetratelabs/wazero/api"
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

func reloadsCallPooled(fn api.Function, stack []uint64) {
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
		os.Stderr.WriteString("graphics.gd/startup: reloads guest call failed: " + err.Error() + "\n")
	}
}
