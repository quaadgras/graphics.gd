package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/noescape"
)

// UnsafeReplaceArray writes value into the Array slot pointed to by frame
// after destroying the Array that Godot pre-allocated there. Use this for
// ptrcall return slots typed Array (godotengine/godot#119440).
func UnsafeReplaceArray(frame gdextension.Pointer, value gdextension.Array) {
	ptr := (*gdextension.Array)(*(*unsafe.Pointer)(unsafe.Pointer(&frame)))
	noescape.Free(gdextension.TypeArray, ptr)
	*ptr = value
}

// UnsafeReplaceDictionary writes value into the Dictionary slot pointed to
// by frame after destroying the Dictionary that Godot pre-allocated there.
// Use this for ptrcall return slots typed Dictionary
// (godotengine/godot#119440).
func UnsafeReplaceDictionary(frame gdextension.Pointer, value gdextension.Dictionary) {
	ptr := (*gdextension.Dictionary)(*(*unsafe.Pointer)(unsafe.Pointer(&frame)))
	noescape.Free(gdextension.TypeDictionary, ptr)
	*ptr = value
}
