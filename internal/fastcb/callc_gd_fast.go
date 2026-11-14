//go:build gd && ((linux && amd64) || ((darwin || ios) && arm64))

package fastcb

import "unsafe" // also for go:linkname

// The fork carries a fused resident outbound crossing on the platforms
// where the direct entry thunk exists: runtime.fastcbCallCFast is
// fastcbCallC + asmcgocall specialized for the resident goroutine calling
// a single-pointer-argument C function with the errno dropped. It skips
// the osPreemptExt bookkeeping, which is a no-op exactly here (not Windows), so
// selecting it is gated on the same platforms as the entry thunk.
//
//go:linkname runtimeCallCFast runtime.fastcbCallCFast
func runtimeCallCFast(fn, arg unsafe.Pointer)

func callC(fn, arg unsafe.Pointer) int32 {
	runtimeCallCFast(fn, arg)
	return 0
}
