//go:build !cgo

// noescape.s is intentionally empty. Its presence marks the package as
// containing assembly, which lets the bodyless //go:noescape declarations in
// noescape.go be defined via //go:linkname. Non-cgo and wasm builds assemble
// this file with `go tool asm`, so the symbols end up in go.o, which the Go
// linker already tags with a non-executable .note.GNU-stack.
//
// The cgo build excludes this file and uses noescape_cgo.S instead: under cgo,
// package .s/.S files are handed to the C assembler, and an object without an
// explicit .note.GNU-stack makes the linker mark the whole shared object's
// stack executable (RWE), which glibc refuses to dlopen. See discussion #310.
