//go:build go1.26 && (amd64 || arm64)

// Package threadcheck provides a fast check for whether the current
// goroutine is running on the main OS thread.
//
// This works by reading the m (machine/OS-thread) pointer from the
// current g struct via the dedicated g register (R14 on amd64, R28
// on arm64). The m pointer is stable across different goroutines
// running on the same OS thread, unlike the g pointer which changes
// when Go assigns a different goroutine to a cgo callback.
//
// Guarded behind go1.26 because it depends on the offset of g.m
// (48 bytes on 64-bit) which is stable but internal to the runtime.
package threadcheck

// currentm returns the m pointer (OS thread) from the current g struct.
// The g register holds the goroutine pointer, and g.m at offset 48
// points to the OS thread struct.
func currentm() uintptr

var mainM = currentm()

func Init() {
	mainM = currentm()
}

// Main reports whether the caller is running on the main OS thread.
func Main() bool {
	return currentm() == mainM
}
