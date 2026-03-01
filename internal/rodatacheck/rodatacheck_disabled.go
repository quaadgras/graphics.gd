//go:build O0 || (!cgo && !wasm)

package rodatacheck

// String always returns false when rodatacheck is disabled.
func String(s string) bool { return false }
