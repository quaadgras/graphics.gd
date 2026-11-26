//go:build gd && wasm.gc

package ring

// pad follows each pointer-sized word of an [Entry]. The engine's web build
// reads an entry as Go's linear WebAssembly port lays it out, with words of
// 8 bytes; compiler.gd's WebAssembly GC port has words of 4, so each is
// followed by 4 bytes that stay zero (the target is little-endian).
type pad = [4]byte
