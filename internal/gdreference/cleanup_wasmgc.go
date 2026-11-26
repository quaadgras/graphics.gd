//go:build gd && wasm.gc

package gdreference

// cleanupIsID reports whether a [runtime.Cleanup] is its id alone, an int64,
// as it is in the WebAssembly GC runtime of compiler.gd: it fits in an
// object, with its id first, which is all that is asked of it.
const cleanupIsID = true
