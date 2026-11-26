//go:build !(gd && wasm.gc)

package gdreference

// cleanupIsID reports whether a [runtime.Cleanup] is its id alone (see
// cleanup_wasmgc.go); it is an id and a pointer, the size of an object.
const cleanupIsID = false
