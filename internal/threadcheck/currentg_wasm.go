//go:build wasm && !(gd && wasm.gc)

package threadcheck

// currentg returns the current goroutine's g pointer (the wasm g register).
func currentg() uintptr
