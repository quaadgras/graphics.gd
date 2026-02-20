package gd

import (
	"sync"

	"graphics.gd/internal/ring"
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

// Flush the ring buffer.
func Flush() {
	ring.Main.Flush()
}
