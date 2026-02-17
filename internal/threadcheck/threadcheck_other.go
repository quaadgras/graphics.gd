//go:build !amd64 && !arm64 && !wasm

package threadcheck

import "graphics.gd/internal/gdextension"

// Main falls back to the cgo-based thread check on platforms
// without assembly support.
func Main() bool {
	return gdextension.Host.Threads.Main()
}
