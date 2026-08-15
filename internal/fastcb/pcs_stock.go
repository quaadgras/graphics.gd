//go:build !gd

package fastcb

import "unsafe" // also for go:linkname

// On stock toolchains the resident-callback machinery arrives as the
// gd CLI's runtime/cgocall.go overlay, which pushes the hook PCs into
// this declaration via its own linkname. Without the overlay nothing
// writes it and the slots stay zero.
//
//go:linkname runtimePCs
var runtimePCs [4]uintptr

// runtimeCallCFastPC receives (from the overlay, same push mechanism as
// runtimePCs) the ABIInternal PC of the fused resident outbound crossing
// (runtime.fastcbCallCFast), or stays 0 when the overlay is absent, older,
// or the platform keeps the generic crossing. A separate variable rather
// than a fifth runtimePCs slot so overlay and graphics.gd versions can skew
// without an array-size mismatch.
//
//go:linkname runtimeCallCFastPC
var runtimeCallCFastPC uintptr

// runtimeArmEntryPC receives (same push mechanism as runtimePCs) the
// ABIInternal PC of the overlay's fastcbArmEntry, through which the direct
// virtual-call entry thunk is armed. 0 when the overlay is absent, older,
// or the platform lacks the thunk.
//
//go:linkname runtimeArmEntryPC
var runtimeArmEntryPC uintptr

// runtimeFrameCgoPC receives (same push mechanism as runtimePCs) the
// ABIInternal PC of the overlay's fastcbFrameCgo, which marks the static
// Scene loop's per-frame stock cgocall as a frame entry (residency may
// engage under it and the return path rebalances). 0 when the overlay is
// absent or older.
//
//go:linkname runtimeFrameCgoPC
var runtimeFrameCgoPC uintptr

// Arming state: the overlay publishes its hook PCs only at the first C->Go
// callback, which is AFTER package init (in a c-shared build, package inits
// run while the library loads, before the engine makes any call). armEntry
// therefore stashes the registration and tryArmEntry completes it from
// initCallC (SetResident), by which point the PCs are published. The
// engine consults the installed pointer per virtual call, so late
// installation simply means earlier calls took the stock path.
var (
	armPendingDispatch func(instance, userdata, result, args uintptr)
	armPendingFallback unsafe.Pointer
	armPendingInstall  func(pc uintptr)
	armFV              funcval
)

func armEntry(dispatch func(instance, userdata, result, args uintptr), fallback unsafe.Pointer, install func(pc uintptr)) {
	armPendingDispatch, armPendingFallback, armPendingInstall = dispatch, fallback, install
	tryArmEntry()
}

func tryArmEntry() {
	if armPendingInstall == nil || runtimeArmEntryPC == 0 {
		return
	}
	armFV = funcval{fn: runtimeArmEntryPC}
	fp := unsafe.Pointer(&armFV)
	arm := *(*func(dispatch func(instance, userdata, result, args uintptr), fallback unsafe.Pointer) uintptr)(unsafe.Pointer(&fp))
	if pc := arm(armPendingDispatch, armPendingFallback); pc != 0 {
		armPendingInstall(pc)
	}
	armPendingDispatch, armPendingFallback, armPendingInstall = nil, nil, nil
}

// callCFV/callCFn: the overlay publishes fastcbCallC only as a PC, so the
// func value is materialised ONCE (initCallC, from SetResident) instead of
// per call — the resident outbound crossing is on the per-frame hot path.
// When the overlay also published the fused crossing (runtimeCallCFastPC),
// callCFastFn is materialised instead and callC prefers it.
var (
	callCFV     funcval
	callCFn     func(fn, arg unsafe.Pointer) int32
	callCFastFV funcval
	callCFastFn func(fn, arg unsafe.Pointer)
)

func initCallC() {
	tryArmEntry()
	if runtimeCallCFastPC != 0 && callCFastFn == nil {
		callCFastFV = funcval{fn: runtimeCallCFastPC}
		fp := unsafe.Pointer(&callCFastFV)
		callCFastFn = *(*func(fn, arg unsafe.Pointer))(unsafe.Pointer(&fp))
	}
	if callCFn != nil || runtimePCs[3] == 0 {
		return
	}
	callCFV = funcval{fn: runtimePCs[3]}
	fp := unsafe.Pointer(&callCFV)
	callCFn = *(*func(fn, arg unsafe.Pointer) int32)(unsafe.Pointer(&fp))
}

func callC(fn, arg unsafe.Pointer) int32 {
	if callCFastFn != nil {
		callCFastFn(fn, arg)
		return 0
	}
	return callCFn(fn, arg)
}
