//go:build !(gd && wasm.gc)

package pointers

import "unsafe"

// wordOf returns the word that is a static function value (see malloc).
func wordOf(fn unsafe.Pointer) uintptr { return uintptr(fn) }

// funcOf returns a pointer to the function value that the word at jump is.
func funcOf(jump *uintptr) unsafe.Pointer { return unsafe.Pointer(jump) }
