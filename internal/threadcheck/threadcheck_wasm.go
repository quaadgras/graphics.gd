//go:build wasm

package threadcheck

func Init() {}

// Main always returns true on wasm, which is single-threaded.
func Main() bool { return true }
