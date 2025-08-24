// Package gdmemory provides functions for transferring data between Go and the graphics engine.
//
// This package is primarily used on platforms where the extension is running in a different
// address space, ie. web/wasm.
package gdmemory

import (
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

var arguments gdextension.Pointer
var receiver gdextension.Pointer
var results [2]gdextension.Pointer
var current int

var setup = sync.OnceFunc(func() {
	arguments = gdextension.Host.Memory.Malloc(64 * 64)
	receiver = gdextension.Host.Memory.Malloc(64 * 64 / 16)
	for i := range results {
		results[i] = gdextension.Host.Memory.Malloc(64 * 64)
	}
})

// CopyArguments copies arguments from args to the arguments buffer, respecting Go's alignment rules.
func CopyArguments[T any](shape gdextension.Shape, args gdextension.CallAccepts[T]) gdextension.Pointer {
	setup()
	if args == nil {
		return 0
	}
	return copyIntoEngine(shape, unsafe.Pointer(args), arguments)
}

// CopyReceiver copies the receiver into the receiver buffer, respecting Go's alignment rules.
func CopyReceiver[T any](shape gdextension.Shape, self gdextension.CallMutates[T]) gdextension.Pointer {
	setup()
	if self == nil {
		return 0
	}
	return copyIntoEngine(shape&0xff00, unsafe.Pointer(self), receiver)
}

func copyIntoEngine(shape gdextension.Shape, args unsafe.Pointer, into gdextension.Pointer) gdextension.Pointer {
	if into == 0 {
		panic("nil pointer dereference")
	}
	buf := unsafe.Slice((*byte)(args), shape.SizeArguments())
	ptr := into
	off := gdextension.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdextension.Host.Memory.Edit.Uint64(ptr+off, *(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdextension.Host.Memory.Edit.Uint32(ptr+off, *(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdextension.Host.Memory.Edit.Uint16(ptr+off, *(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdextension.Host.Memory.Edit.Byte(ptr+off, *(*uint8)(unsafe.Pointer(&buf[0])))
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
	gdextension.Host.Memory.Clear(results[current], shape.SizeResult())
	return results[current]
}

func LoadResult[T ~unsafe.Pointer](shape gdextension.Shape, result T, from gdextension.Pointer) {
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
	for size > 0 {
		switch {
		case size >= 8:
			*(*uint64)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Uint64(from + gdextension.Pointer(done))
			done += 8
			size -= 8
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Uint32(from + gdextension.Pointer(done))
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Uint16(from + gdextension.Pointer(done))
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(data, done)) = gdextension.Host.Memory.Load.Byte(from + gdextension.Pointer(done))
			done += 1
			size -= 1
		default:
			return
		}
	}
}

func CopyVariants[T any](args gdextension.CallAccepts[T], n int) gdextension.Pointer {
	setup()
	var offset int
	var data = unsafe.Pointer(args)
	for i := range n {
		gdextension.Host.Memory.Edit.Bits128(arguments+gdextension.Pointer(offset), *(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset+16), *(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
	}
	return arguments
}

func Int64frombits(bits uint64) int64 {
	return *(*int64)(unsafe.Pointer(&bits))
}

func CopyBufferToEngine(buf []byte) gdextension.Pointer {
	ptr := gdextension.Host.Memory.Malloc(len(buf))
	off := gdextension.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gdextension.Host.Memory.Edit.Uint64(ptr+off, *(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gdextension.Host.Memory.Edit.Uint32(ptr+off, *(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gdextension.Host.Memory.Edit.Uint16(ptr+off, *(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gdextension.Host.Memory.Edit.Byte(ptr+off, *(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func CopyBufferToGo(ptr gdextension.Pointer, buf []byte) {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	off := 0
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			*(*uint64)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Uint64(ptr + gdextension.Pointer(off))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Uint32(ptr + gdextension.Pointer(off))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Uint16(ptr + gdextension.Pointer(off))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			*(*uint8)(unsafe.Pointer(&buf[0])) = gdextension.Host.Memory.Load.Byte(ptr + gdextension.Pointer(off))
			buf = buf[1:]
			off += 1
		}
	}
	gdextension.Host.Memory.Free(ptr)
}
