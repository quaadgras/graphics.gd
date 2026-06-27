package startup

import gd "graphics.gd/internal"

// OnCrash registers a function to run on a best-effort basis when a crash occurs within the engine.
// The function will not be called when the project crashes within Go code, use [debug.SetCrashOutput]
// instead for catching those scenarios. Not all engine crashes can be observed in this way and on
// supported platforms, launching a supervisor child process is recommended. It may be useful for the
// handler to use runtime functionality dump the Go stack.
func OnCrash(fn func()) {
	gd.CrashHandlers = append(gd.CrashHandlers, fn)
}
