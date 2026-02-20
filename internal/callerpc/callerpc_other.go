//go:build !amd64 && !arm64

package callerpc

import "runtime"

// Callerpc returns the PC of the caller's caller via runtime.Callers.
func Callerpc() uintptr {
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	return pcs[0]
}
