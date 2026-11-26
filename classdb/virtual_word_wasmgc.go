//go:build gd && wasm.gc

package classdb

import "graphics.gd/internal/gdextension"

// virtualWord returns the word the engine holds for a virtual function and
// hands back on every call of it. With compiler.gd's WebAssembly GC port a
// pointer cannot be made of an integer, so the word is where the function
// (the last one made) is in pinnedVirtuals, counted from 1.
func virtualWord(pv *pinnedVirtualFunc) uintptr { return uintptr(len(pinnedVirtuals)) }

// virtualOf returns the virtual function that the engine's word is for.
func virtualOf(callData gdextension.Pointer) *pinnedVirtualFunc {
	return pinnedVirtuals[callData-1]
}
