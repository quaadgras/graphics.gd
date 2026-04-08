// Package gdmemory provides functions for transferring data between Go and the graphics engine.
//
// This package is primarily used on platforms where the extension is running in a different
// address space, ie. web/wasm.
package gdmemory

import (
	"sync"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

var arguments gdextension.Pointer
var receiver gdextension.Pointer
var results [2]gdextension.Pointer
var current int

var setup = sync.OnceFunc(func() {
	arguments = gdextension.Pointer(gdunsafe.Malloc(64 * 64))
	receiver = gdextension.Pointer(gdunsafe.Malloc(64 * 64 / 16))
	for i := range results {
		results[i] = gdextension.Pointer(gdunsafe.Malloc(64 * 64))
		gdunsafe.Memset(gdunsafe.MutablePointer(results[i]), 64*64, 0)
	}
})

// CopyArguments copies arguments from args to the arguments buffer, respecting Go's alignment rules.
func CopyArguments[T any](shape gdextension.Shape, args gdextension.CallAccepts[T]) gdextension.Pointer {
	setup()
	if args == nil {
		return 0
	}
	return copyIntoEngine(shape.SizeArguments(), unsafe.Pointer(args), arguments)
}

// CopyReceiver copies the receiver into the receiver buffer, respecting Go's alignment rules.
func CopyReceiver[T any](shape gdextension.Shape, self gdextension.CallMutates[T]) gdextension.Pointer {
	setup()
	if self == nil {
		return 0
	}
	return copyIntoEngine((shape & 0b11110000).SizeArguments(), unsafe.Pointer(self), receiver)
}

func copyIntoEngine(bytes int, args unsafe.Pointer, into gdextension.Pointer) gdextension.Pointer {
	if into == 0 {
		panic("nil pointer dereference")
	}
	buf := unsafe.Slice((*byte)(args), bytes)
	ptr := into
	off := gdextension.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdunsafe.MutablePointerTo[uint64](ptr + off).Set(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdunsafe.MutablePointerTo[uint32](ptr + off).Set(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdunsafe.MutablePointerTo[uint16](ptr + off).Set(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdunsafe.MutablePointerTo[byte](ptr + off).Set(*(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func MakeResult(shape gdextension.Shape) gdextension.Pointer {
	setup()
	// alternating between two buffers for results
	if current == 0 {
		current = 1
	} else {
		current = 0
	}
	return results[current]
}

func LoadResult[T ~unsafe.Pointer | *gdextension.Variant](shape gdextension.Shape, result T, from gdextension.Pointer) {
	setup()
	if from == 0 {
		panic("nil pointer dereference")
	}
	data := unsafe.Pointer(result)
	done := 0
	size := shape.SizeResult()
	if size == 0 {
		return
	}
	defer gdunsafe.Memset(gdunsafe.MutablePointer(from), uintptr(shape.SizeResult()), 0)
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = gdunsafe.PointerTo[uint32](from + gdextension.Pointer(done)).Get()
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = gdunsafe.PointerTo[uint16](from + gdextension.Pointer(done)).Get()
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(data, done)) = gdunsafe.PointerTo[byte](from + gdextension.Pointer(done)).Get()
			done += 1
			size -= 1
		default:
			return
		}
	}
}

func CopyVariants[T ~unsafe.Pointer | *gdextension.Variant](args T, n int) gdextension.Pointer {
	setup()
	var offset int
	var data = unsafe.Pointer(args)
	for i := range n {
		gdunsafe.MutablePointerTo[[2]uint64](arguments + gdextension.Pointer(offset)).Set(*(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		gdunsafe.MutablePointerTo[uint64](arguments + gdextension.Pointer(offset+16)).Set(*(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return arguments
}

func Int64frombits(bits uint64) int64 {
	return *(*int64)(unsafe.Pointer(&bits))
}

func CopyBufferToEngine(buf []byte) gdextension.Pointer {
	ptr := gdunsafe.Malloc(uintptr(len(buf)))
	off := gdunsafe.MutablePointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdunsafe.MutablePointerTo[uint64](ptr + off).Set(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdunsafe.MutablePointerTo[uint32](ptr + off).Set(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdunsafe.MutablePointerTo[uint16](ptr + off).Set(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdunsafe.MutablePointerTo[byte](ptr + off).Set(*(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return gdextension.Pointer(ptr)
}

func CopyBufferToGo(ptr gdextension.Pointer, buf []byte) {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	off := 0
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = gdunsafe.PointerTo[uint32](ptr + gdextension.Pointer(off)).Get()
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = gdunsafe.PointerTo[uint16](ptr + gdextension.Pointer(off)).Get()
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			*(*uint8)(unsafe.Pointer(&buf[0])) = gdunsafe.PointerTo[byte](ptr + gdextension.Pointer(off)).Get()
			buf = buf[1:]
			off += 1
		}
	}
	gdunsafe.MutablePointer(ptr).Free()
}
