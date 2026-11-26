//go:build gd && wasm.gc

package threadcheck

import "unsafe"

// The WebAssembly GC runtime of compiler.gd has no assembly and no g
// register: the running goroutine is the runtime's curg.
//
//go:linkname curg runtime.curg
var curg unsafe.Pointer

// currentg returns the current goroutine's g pointer.
func currentg() uintptr { return uintptr(curg) }
