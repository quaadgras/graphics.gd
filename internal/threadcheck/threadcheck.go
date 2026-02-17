//go:build amd64 || arm64

// Package threadcheck provides a fast check for whether the current
// goroutine is the main goroutine locked to the main OS thread.
//
// This works by comparing the current g pointer (kept in a dedicated
// register by the Go runtime) against the g pointer captured at init
// time. Since the main goroutine is always locked to the main OS
// thread via runtime.LockOSThread, its g pointer is stable for the
// lifetime of the process.
package threadcheck

// currentg returns the current goroutine pointer from the dedicated
// g register (R14 on amd64, R28 on arm64).
func currentg() uintptr

var mainG = currentg()

func Init() {
	mainG = currentg()
}

// Main reports whether the caller is running on the main goroutine
// (which is locked to the main OS thread).
func Main() bool {
	return currentg() == mainG
}
