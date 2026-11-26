//go:build gd && wasm.gc

package pointers

import "unsafe"

// With compiler.gd's WebAssembly GC port a pointer cannot be made of an
// integer, so the function values that [tables] holds as words are kept
// here, by their address. (The port has one thread: no lock.)
var funcs = map[uintptr]unsafe.Pointer{}

// wordOf returns the word that stands for a function value (see malloc).
func wordOf(fn unsafe.Pointer) uintptr {
	word := uintptr(fn)
	if word != 0 {
		if _, ok := funcs[word]; !ok {
			funcs[word] = fn
		}
	}
	return word
}

// funcOf returns a pointer to the function value that the word at jump
// stands for.
func funcOf(jump *uintptr) unsafe.Pointer {
	fn := new(unsafe.Pointer)
	*fn = funcs[*jump]
	return unsafe.Pointer(fn)
}
