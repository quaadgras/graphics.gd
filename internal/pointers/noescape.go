package pointers

import "unsafe"

//go:noescape
func NoEscape(ptr unsafe.Pointer) unsafe.Pointer

//go:noescape
func NoEscape3(a, b, c unsafe.Pointer) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)

//go:linkname noescape graphics.gd/internal/pointers.NoEscape
//go:nosplit
func noescape(ptr unsafe.Pointer) unsafe.Pointer {
	return ptr
}

//go:linkname noescape3 graphics.gd/internal/pointers.NoEscape3
//go:nosplit
func noescape3(a, b, c unsafe.Pointer) (unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) {
	return a, b, c
}
