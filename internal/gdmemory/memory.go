package gdmemory

import (
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

var arguments gdextension.Pointer
var results [2]gdextension.Pointer
var current int

var setup = sync.OnceFunc(func() {
	arguments = gdextension.Host.Memory.Malloc(64 * 64)
	for i := range results {
		results[i] = gdextension.Host.Memory.Malloc(64 * 64)
	}
})

func CopyArguments[T any](shape gdextension.Shape, args gdextension.CallAccepts[T]) gdextension.Pointer {
	setup()
	size := shape.SizeArguments()
	if size == 0 {
		return gdextension.Pointer(0)
	}
	if size >= 64*64 {
		panic("too many arguments")
	}
	done := 0
	data := unsafe.Pointer(args)
	for {
		switch {
		case size >= 8*8:
			gdextension.Host.Memory.Edit.Bits512(arguments+gdextension.Pointer(done), *(*[8]uint64)(data))
			done += 8 * 8
			size -= 8 * 8
		case size >= 4*8:
			gdextension.Host.Memory.Edit.Bits256(arguments+gdextension.Pointer(done), *(*[4]uint64)(data))
			done += 4 * 8
			size -= 4 * 8
		case size >= 2*8:
			gdextension.Host.Memory.Edit.Bits128(arguments+gdextension.Pointer(done), *(*[2]uint64)(data))
			done += 2 * 8
			size -= 2 * 8
		case size >= 8:
			gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(done), *(*uint64)(data))
			done += 8
			size -= 8
		case size >= 4:
			gdextension.Host.Memory.Edit.Uint32(arguments+gdextension.Pointer(done), *(*uint32)(data))
			done += 4
			size -= 4
		case size >= 2:
			gdextension.Host.Memory.Edit.Uint16(arguments+gdextension.Pointer(done), *(*uint16)(data))
			done += 2
			size -= 2
		case size >= 1:
			gdextension.Host.Memory.Edit.Byte(arguments+gdextension.Pointer(done), *(*uint8)(data))
			done += 1
			size -= 1
		default:
			return arguments
		}
	}
}

func MakeResult(shape gdextension.Shape) gdextension.Pointer {
	setup()
	current++
	return results[current-1]
}

func LoadResult[T any](shape gdextension.Shape, result gdextension.CallReturns[T], from gdextension.Pointer) {
	setup()
	current--
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
			*(*uint32)(unsafe.Pointer(result)) = gdextension.Host.Memory.Load.Uint32(from + gdextension.Pointer(done))
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Pointer(result)) = gdextension.Host.Memory.Load.Uint16(from + gdextension.Pointer(done))
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Pointer(result)) = gdextension.Host.Memory.Load.Byte(from + gdextension.Pointer(done))
			done += 1
			size -= 1
		default:
			return
		}
	}
}

func CopyVariants[T any](args gdextension.CallAccepts[T], n int) gdextension.Pointer {
	setup()
	return 0
}
