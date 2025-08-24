//go:build amd64 || arm64

package gdmemory

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

func Get[T gdextension.AnyVariant](frame gdextension.Pointer) T {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&frame))
	return *(*T)(ptr)
}

func Set[T gdextension.AnyVariant](frame gdextension.Pointer, value T) {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&frame))
	*(*T)(ptr) = value
}

func IntoSlice[T gdextension.Packable](ptr gdextension.Pointer, len int) []T {
	var slice = make([]T, len)
	copy(slice, unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), len))
	return slice
}

func LoadSlice[T gdextension.Packable](ptr gdextension.Pointer, slice []T) {
	if len(slice) == 0 {
		return
	}
	copy(unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), len(slice)), slice)
}

func IndexVariants(ptr gdextension.Accepts[gdextension.Variant], len, idx int) gdextension.Variant {
	return *(unsafe.Slice((**gdextension.Variant)(ptr), len)[idx])
}
