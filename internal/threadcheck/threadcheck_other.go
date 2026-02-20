//go:build (!go1.26 || (!amd64 && !arm64)) && !wasm

package threadcheck

import "graphics.gd/internal/gdextension"

func Init() {}

// Main falls back to the cgo-based thread check on platforms
// without assembly support or unsupported Go versions.
func Main() bool {
	return gdextension.Host.Threads.Main()
}
