//go:build js || wasip1

package gdmemory

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

func Get[T gdextension.AnyVariant | gdextension.CallError](frame gdextension.Pointer) T {
	if frame == 0 {
		panic("nil pointer dereference")
	}
	var addr = frame
	var zero T
	var done = 0
	var size = unsafe.Sizeof([1]T{})
	for size > 0 {
		switch {
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

func Set[T gdextension.AnyVariant | gdextension.CallError](frame gdextension.Pointer, value T) {
	if frame == 0 {
		panic("nil pointer dereference")
	}
	var addr = gdextension.Pointer(uintptr(frame))
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

func IntoSlice[T gdextension.Packable](ptr gdextension.Pointer, len int) []T {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	var slice = make([]T, len)
	for i := range slice {
		slice[i] = Get[T](ptr + gdextension.Pointer(i)*gdextension.Pointer(unsafe.Sizeof([1]T{}[0])))
	}
	return slice
}

func LoadSlice[T gdextension.Packable](ptr gdextension.Pointer, slice []T) {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	if len(slice) == 0 {
		return
	}
	off := gdextension.Pointer(0)
	buf := unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*int(unsafe.Sizeof([1]T{}[0])))
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
}

func IndexVariants(addr gdextension.Accepts[gdextension.Variant], len, idx int) gdextension.Variant {
	if addr == [1]gdextension.Accepts[gdextension.Variant]{}[0] {
		panic("nil pointer dereference")
	}
	ptr := Get[gdextension.Object](gdextension.Pointer(addr) + gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(gdextension.Pointer(0))))
	if ptr == 0 {
		return gdextension.Variant{}
	}
	return Get[gdextension.Variant](gdextension.Pointer(ptr))
}
