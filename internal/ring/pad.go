//go:build !(gd && wasm.gc)

package ring

// pad follows each pointer-sized word of an [Entry] (see pad_wasmgc.go).
type pad = [0]byte
