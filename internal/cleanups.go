package gd

import (
	"sync"

	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
)

var cleanups []func()
var mutex sync.Mutex

// RegisterCleanup registers a function to be called when the engine shuts down.
func RegisterCleanup(f func()) {
	mutex.Lock()
	defer mutex.Unlock()
	cleanups = append(cleanups, f)
}

// Cleanups returns a slice of all registered cleanup functions.
func Cleanups() []func() {
	mutex.Lock()
	defer mutex.Unlock()
	return cleanups
}

var StartupFunctions []func()
var PostStartupFunctions []func()

var EditorStartupFunctions []func()

// Flush the ring buffer. ring.Main is a main-thread-only batch buffer, so this
// is a no-op off the main thread (e.g. when a generated binding is reached from
// a ResourceFormatLoader callback running on a dedicated resource-loading
// thread) — flushing it there would race the main thread and corrupt the ring.
func Flush() {
	if threadcheck.Main() {
		ring.Main.Flush()
	}
}

var CrashHandlers []func()
