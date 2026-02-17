//go:build wasm

package threadcheck

// Main always returns true on wasm, which is single-threaded.
func Main() bool { return true }
