package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

type gdptr uint32

type Variant pointers.Trio[Variant]
type Signal pointers.Pair[Signal]
type Callable pointers.Pair[Callable]

type Dictionary pointers.Half[Dictionary]
type Array pointers.Half[Array]
type String pointers.Half[String]
type StringName pointers.Half[StringName]
type NodePath pointers.Half[NodePath]

type PackedByteArray pointers.Solo[PackedByteArray]
type PackedInt32Array pointers.Solo[PackedInt32Array]
type PackedInt64Array pointers.Solo[PackedInt64Array]
type PackedFloat32Array pointers.Solo[PackedFloat32Array]
type PackedFloat64Array pointers.Solo[PackedFloat64Array]
type PackedStringArray pointers.Solo[PackedStringArray]
type PackedVector2Array pointers.Solo[PackedVector2Array]
type PackedVector3Array pointers.Solo[PackedVector3Array]
type PackedVector4Array pointers.Solo[PackedVector4Array]
type PackedColorArray pointers.Solo[PackedColorArray]

type EnginePointer = uint32
type PackedPointers = [1]uint64

func UnsafeGet[T any](frame Address, index int) T {
	var addr = gdextension.Pointer(frame)
	var zero T
	var done = 0
	var size = unsafe.Sizeof([1]T{})
	for size > 0 {
		switch {
		case size >= 8:
			*(*uint64)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Uint64(addr)
			addr += 8
			done += 8
			size -= 8
		case size >= 4:
			*(*uint32)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Uint32(addr)
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Uint16(addr)
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Byte(addr)
			addr += 1
			done += 1
			size -= 1
		}
	}
	return zero
}

func UnsafeSet[T any](frame Address, value T) {
	var addr = gdextension.Pointer(frame)
	var size = unsafe.Sizeof([1]T{})
	var done = 0
	for size > 0 {
		switch {
		case size >= 8:
			gdextension.Host.Memory.Edit.Uint64(addr, *(*uint64)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 8
			done += 8
			size -= 8
		case size >= 4:
			gdextension.Host.Memory.Edit.Uint32(addr, *(*uint32)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			gdextension.Host.Memory.Edit.Uint16(addr, *(*uint16)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			gdextension.Host.Memory.Edit.Byte(addr, *(*uint8)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 1
			done += 1
			size -= 1
		}
	}
}
