//go:build gd && !((linux && amd64) || ((darwin || ios) && arm64))

package fastcb

import "unsafe"

// Generic fork crossing for platforms without the fused variant: the
// full fastcbCallC (osPreemptExt bookkeeping + asmcgocall).
func callC(fn, arg unsafe.Pointer) int32 { return runtimeCallC(fn, arg) }
