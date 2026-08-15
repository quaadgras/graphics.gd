//go:build gd

package fastcb

import "unsafe" // also for go:linkname

// The compiler.gd fork carries the resident-callback machinery in its
// own runtime (no overlay involved) and publishes the hook PCs into
// runtime.graphicsFastcbPCs, which has a handshake linkname for this
// alias. The fork sets the `gd` build tag, so this pull only ever
// resolves against a runtime that defines the symbol.
//
//go:linkname runtimePCs runtime.graphicsFastcbPCs
var runtimePCs [4]uintptr

// The fork's runtime also carries the direct C-ABI virtual-call entry
// (runtime.fastcbentry): armEntry registers the dispatch target plus the
// stock fallback entry and returns the thunk's C-callable PC (0 when the
// platform lacks the thunk). See runtime.fastcbArmEntry in the fork.
//
//go:linkname runtimeArmEntry runtime.fastcbArmEntry
func runtimeArmEntry(dispatch func(instance, userdata, result, args uintptr), fallback unsafe.Pointer) uintptr

func armEntry(dispatch func(instance, userdata, result, args uintptr), fallback unsafe.Pointer, install func(pc uintptr)) {
	if pc := runtimeArmEntry(dispatch, fallback); pc != 0 {
		install(pc)
	}
}

// The fork exports fastcbCallC by linkname, so the resident outbound
// crossing is a direct call — no per-call func-value construction (the
// stock overlay path can only reach the hook through its published PC).
// callC itself is platform-split: linux/amd64 selects the fused
// runtime.fastcbCallCFast (callc_gd_fast.go), everything else falls back
// to this generic hook (callc_gd_generic.go).
//
//go:linkname runtimeCallC runtime.fastcbCallC
func runtimeCallC(fn, arg unsafe.Pointer) int32

func initCallC() {}

// The fork's runtime does not carry the frame-entry cgocall hook yet:
// the PC stays zero and FrameCgo no-ops (the static Scene loop then runs
// its frames through the raw asmcgocall path as before).
var runtimeFrameCgoPC uintptr
