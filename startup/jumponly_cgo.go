//go:build cgo

package startup

// #include <stdint.h>
// extern uintptr_t gd_ptrcall_fn_addr();
import "C"

import "graphics.gd/internal/jumponly"

// initJumponly sets the ptrcall function pointer for the jumponly package.
// Must be called after cgo_extension_init has loaded proc addresses
// (i.e., during engine CORE-level init or later).
func initJumponly() {
	jumponly.PtrcallFn = uintptr(C.gd_ptrcall_fn_addr())
}
