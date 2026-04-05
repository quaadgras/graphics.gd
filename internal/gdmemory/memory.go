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
		gdunsafe.Clear(gdunsafe.Pointer(results[i]), 64*64)
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
			gdunsafe.Pointer(ptr + off).SetUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdunsafe.Pointer(ptr + off).SetUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdunsafe.Pointer(ptr + off).SetUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdunsafe.Pointer(ptr + off).SetByte(*(*uint8)(unsafe.Pointer(&buf[0])))
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
	defer gdunsafe.Clear(gdunsafe.Pointer(from), gdunsafe.Int(shape.SizeResult()))
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = gdunsafe.Pointer(from + gdextension.Pointer(done)).Uint32()
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = gdunsafe.Pointer(from + gdextension.Pointer(done)).Uint16()
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(data, done)) = gdunsafe.Pointer(from + gdextension.Pointer(done)).Byte()
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
		gdunsafe.Pointer(arguments + gdextension.Pointer(offset)).SetBits128(*(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		gdunsafe.Pointer(arguments + gdextension.Pointer(offset+16)).SetUint64(*(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return arguments
}

func Int64frombits(bits uint64) int64 {
	return *(*int64)(unsafe.Pointer(&bits))
}

func CopyBufferToEngine(buf []byte) gdextension.Pointer {
	ptr := gdunsafe.Malloc(gdunsafe.Int(len(buf)))
	off := gdunsafe.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdunsafe.Pointer(ptr + off).SetUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdunsafe.Pointer(ptr + off).SetUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdunsafe.Pointer(ptr + off).SetUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdunsafe.Pointer(ptr + off).SetByte(*(*uint8)(unsafe.Pointer(&buf[0])))
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
			*(*uint32)(unsafe.Pointer(&buf[0])) = gdunsafe.Pointer(ptr + gdextension.Pointer(off)).Uint32()
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = gdunsafe.Pointer(ptr + gdextension.Pointer(off)).Uint16()
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			*(*uint8)(unsafe.Pointer(&buf[0])) = gdunsafe.Pointer(ptr + gdextension.Pointer(off)).Byte()
			buf = buf[1:]
			off += 1
		}
	}
	gdunsafe.Pointer(ptr).Free()
}
