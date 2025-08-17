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

// ALIGN_UP aligns a value to the next multiple of align.
func alignUp(value, align uint32) uint32 {
	return (value + (align - 1)) & ^(align - 1)
}

// CopyArguments copies arguments from args to the arguments buffer, respecting Go's alignment rules.
func CopyArguments[T any](shape gdextension.Shape, args gdextension.CallAccepts[T]) gdextension.Pointer {
	setup()

	offset := uint32(0)          // Track current offset in the arguments buffer
	data := unsafe.Pointer(args) // Base pointer to args
	totalSize := uint32(0)       // Track total size for limit check
	shapeVal := shape            // Assuming shape is convertible to uint64

	for i := 1; i < 16; i++ {
		// Extract 4-bit code from shape
		code := (shapeVal >> (i * 4)) & 0xF
		var size = uint32(code.SizeResult())
		if size == 0 {
			return arguments
		}

		// Check total size limit
		totalSize += size
		if totalSize >= 64*64 {
			panic("too many arguments")
		}

		// Align offset to min(size, 8) as per Go's 64-bit alignment rules
		alignment := size
		if alignment > 8 {
			alignment = 8
		}
		offset = alignUp(offset, alignment)

		// Copy argument to the aligned offset in the arguments buffer
		switch size {
		case 1:
			gdextension.Host.Memory.Edit.Byte(arguments+gdextension.Pointer(offset), *(*uint8)(data))
		case 2:
			gdextension.Host.Memory.Edit.Uint16(arguments+gdextension.Pointer(offset), *(*uint16)(data))
		case 4:
			gdextension.Host.Memory.Edit.Uint32(arguments+gdextension.Pointer(offset), *(*uint32)(data))
		case 8:
			gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset), *(*uint64)(data))
		case 12:
			// Copy as 8 bytes + 4 bytes
			gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset), *(*uint64)(data))
			gdextension.Host.Memory.Edit.Uint32(arguments+gdextension.Pointer(offset+8), *(*uint32)(unsafe.Pointer(uintptr(data) + 8)))
		case 16:
			gdextension.Host.Memory.Edit.Bits128(arguments+gdextension.Pointer(offset), *(*[2]uint64)(data))
		case 24:
			// Copy as 8 + 8 + 8 bytes
			gdextension.Host.Memory.Edit.Bits128(arguments+gdextension.Pointer(offset), *(*[2]uint64)(data))
			gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset+16), *(*uint64)(unsafe.Pointer(uintptr(data) + 16)))
		case 32:
			gdextension.Host.Memory.Edit.Bits256(arguments+gdextension.Pointer(offset), *(*[4]uint64)(data))
		case 36:
			// Copy as 32 + 4 bytes
			gdextension.Host.Memory.Edit.Bits256(arguments+gdextension.Pointer(offset), *(*[4]uint64)(data))
			gdextension.Host.Memory.Edit.Uint32(arguments+gdextension.Pointer(offset+32), *(*uint32)(unsafe.Pointer(uintptr(data) + 32)))
		case 40:
			// Copy as 32 + 8 bytes
			gdextension.Host.Memory.Edit.Bits256(arguments+gdextension.Pointer(offset), *(*[4]uint64)(data))
			gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset+32), *(*uint64)(unsafe.Pointer(uintptr(data) + 32)))
		case 48:
			// Copy as 32 + 16 bytes
			gdextension.Host.Memory.Edit.Bits256(arguments+gdextension.Pointer(offset), *(*[4]uint64)(data))
			gdextension.Host.Memory.Edit.Bits128(arguments+gdextension.Pointer(offset+32), *(*[2]uint64)(unsafe.Pointer(uintptr(data) + 32)))
		case 64:
			gdextension.Host.Memory.Edit.Bits512(arguments+gdextension.Pointer(offset), *(*[8]uint64)(data))
		case 72:
			// Copy as 64 + 8 bytes
			gdextension.Host.Memory.Edit.Bits512(arguments+gdextension.Pointer(offset), *(*[8]uint64)(data))
			gdextension.Host.Memory.Edit.Uint64(arguments+gdextension.Pointer(offset+64), *(*uint64)(unsafe.Pointer(uintptr(data) + 64)))
		case 96:
			// Copy as 64 + 32 bytes
			gdextension.Host.Memory.Edit.Bits512(arguments+gdextension.Pointer(offset), *(*[8]uint64)(data))
			gdextension.Host.Memory.Edit.Bits256(arguments+gdextension.Pointer(offset+64), *(*[4]uint64)(unsafe.Pointer(uintptr(data) + 64)))
		case 128:
			// Copy as 64 + 64 bytes
			gdextension.Host.Memory.Edit.Bits512(arguments+gdextension.Pointer(offset), *(*[8]uint64)(data))
			gdextension.Host.Memory.Edit.Bits512(arguments+gdextension.Pointer(offset+64), *(*[8]uint64)(unsafe.Pointer(uintptr(data) + 64)))
		default:
			panic("unsupported argument size")
		}

		// Advance data pointer and offset
		data = unsafe.Pointer(uintptr(data) + uintptr(size))
		offset += size
	}

	return arguments
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
