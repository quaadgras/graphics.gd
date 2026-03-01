//go:build wasm && !O0

package rodatacheck

import "unsafe"

// threshold is set to a stack address during init; in WASM linear memory
// the layout is: data segments (low) -> stack -> heap (high), so any
// pointer below the stack is in the static data region.
var threshold uintptr

func init() {
	var x byte
	threshold = uintptr(unsafe.Pointer(&x))
}

// String reports whether s is backed by the binary's read-only data section.
func String(s string) bool {
	p := unsafe.StringData(s)
	if p == nil {
		return false
	}
	return uintptr(unsafe.Pointer(p)) < threshold
}
