//go:build js || wasip1

package gdmemory

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

func Get[T gdextension.AnyVariant | gdextension.CallError | gdextension.Pointer](frame gdextension.Pointer) T {
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
			*(*uint32)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.PointerTo[uint32](addr).Get()
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.PointerTo[uint16](addr).Get()
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.PointerTo[uint8](addr).Get()
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
			gdunsafe.MutablePointerTo[uint64](addr).Set(*(*uint64)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 8
			done += 8
			size -= 8
		case size >= 4:
			gdunsafe.MutablePointerTo[uint32](addr).Set(*(*uint32)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			gdunsafe.MutablePointerTo[uint16](addr).Set(*(*uint16)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			gdunsafe.MutablePointerTo[uint8](addr).Set(*(*uint8)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 1
			done += 1
			size -= 1
		}
	}
}

func IntoSlice[T gdunsafe.Packable](ptr gdunsafe.PointerTo[T], len int) []T {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	var slice = make([]T, len)
	for i := range slice {
		slice[i] = (gdunsafe.PointerTo[T](ptr) + gdunsafe.PointerTo[T](i)*gdunsafe.PointerTo[T](unsafe.Sizeof([1]T{}[0]))).Get()
	}
	return slice
}

func LoadSlice[T gdunsafe.Packable](ptr gdunsafe.MutablePointerTo[T], slice []T) {
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	if len(slice) == 0 {
		return
	}
	off := gdunsafe.MutablePointerTo[T](0)
	buf := unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*int(unsafe.Sizeof([1]T{}[0])))
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
			gdunsafe.MutablePointerTo[uint8](ptr + off).Set(*(*uint8)(unsafe.Pointer(&buf[0])))
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

// UnsafePointer returns the underlying [unsafe.Pointer] for
// the Pointer[T] and true if the result is valid.
func (ptr *Pointer[T]) UnsafePointer() (unsafe.Pointer, bool) {
	return nil, false
}
