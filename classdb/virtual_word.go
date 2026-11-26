//go:build !(gd && wasm.gc)

package classdb

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// virtualWord returns the word the engine holds for a virtual function and
// hands back on every call of it: its address, as it is pinned.
func virtualWord(pv *pinnedVirtualFunc) uintptr { return uintptr(unsafe.Pointer(pv)) }

// virtualOf returns the virtual function that the engine's word is for.
func virtualOf(callData gdextension.Pointer) *pinnedVirtualFunc {
	return (*pinnedVirtualFunc)(*(*unsafe.Pointer)(unsafe.Pointer(&callData))) // runtime.Pinned, so this is ok.
}
