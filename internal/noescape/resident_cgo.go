//go:build cgo && !(gd && wasm)

package noescape

// The C functions live in the root package's gd.c and are linked across the
// whole binary; the addr getters exist because fastcb.CallC (asmcgocall)
// needs a raw C function pointer, not a cgo wrapper.

// extern void *gd_resident_call_addr(void);
// extern void *gd_resident_variant_call_addr(void);
// extern void *gd_resident_script_call_addr(void);
import "C"

import (
	"structs"
	"unsafe"

	"graphics.gd/internal/fastcb"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/threadcheck"
)

// Resident outbound calls: while resident-callback mode is engaged (see
// graphics.gd/internal/fastcb), main-thread engine crossings go through
// fastcb.CallC (runtime.asmcgocall) instead of cgocall, so the goroutine
// stays _Grunning with its P wired for the duration of the engine call and
// any engine->Go callbacks nested inside it (script calls into Go callables,
// re-entrant virtuals, signal handlers) take the resident fast path instead
// of a full cgocallback transition.
//
// Because the goroutine stays _Grunning, a nested callback's Go code can grow
// — and therefore MOVE — this goroutine's stack while the engine holds raw
// pointers into the call's argument and result memory (the engine encodes the
// return value only after the method body ran). cgo-generated shims survive
// this via their _cgo_topofstack pointer adjustment, but that cannot fix
// pointers the engine itself holds across the callback. So the calls are
// staged through package-level packs: Go globals never move. One pack per
// nesting depth, since a nested callback can itself make resident calls.

// residentPack mirrors gd.c's struct gd_resident_pack.
type residentPack struct {
	_      structs.HostLayout
	object uintptr
	method uintptr
	shape  uint64
	args   [256]byte
	result [64]byte
}

// residentVariantPack mirrors gd.c's struct gd_resident_variant_pack.
type residentVariantPack struct {
	_      structs.HostLayout
	object uintptr
	method uintptr
	argc   int64
	args   [240]byte // up to 10 Variants, 24 bytes each
	result gdextension.Variant
	err    gdextension.CallError
	_      [4]byte
}

var (
	residentPacks        [8]residentPack
	residentDepth        int
	residentVariantPacks [4]residentVariantPack
	residentVariantDepth int

	// C function pointers for fastcb.CallC (asmcgocall needs a raw pointer, not
	// a cgo wrapper). Resolved once at init; the getters return constant
	// addresses, so there is nothing to defer to the hot path.
	residentCallAddr        = C.gd_resident_call_addr()
	residentVariantCallAddr = C.gd_resident_variant_call_addr()
	residentScriptCallAddr  = C.gd_resident_script_call_addr()
)

// residentAvailable reports whether the resident outbound path can carry this
// ptrcall and, when it can, the decoded argument size (so residentCall need not
// decode the shape a second time). Main thread only (the caller must check
// threadcheck.Main first so the fastcb.Resident read never happens off-main).
func residentAvailable(shape gdextension.Shape) (argSize int, ok bool) {
	if !fastcb.Resident() || residentDepth >= len(residentPacks) {
		return 0, false
	}
	n := shape.SizeArguments()
	if n > len(residentPacks[0].args) {
		return 0, false
	}
	return n, true
}

// residentCall performs the engine crossing on the calling (main) thread via
// fastcb.CallC, staying _Grunning. See the package comment above for why the
// call is staged through a global pack. argSize is the decoded argument size
// from residentAvailable.
func residentCall[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, argSize int, argptr unsafe.Pointer) T {
	p := &residentPacks[residentDepth]
	residentDepth++
	p.object = uintptr(object)
	p.method = uintptr(method)
	p.shape = uint64(shape)
	if argSize > 0 && argptr != nil {
		copy(p.args[:argSize], unsafe.Slice((*byte)(argptr), argSize))
	}
	fastcb.CallC(residentCallAddr, unsafe.Pointer(p))
	residentDepth--
	var result T
	if unsafe.Sizeof(result) > 0 {
		result = *(*T)(unsafe.Pointer(&p.result))
	}
	return result
}

// residentVariantDispatch stages a variant-shaped call (method_bind_call or
// script call, selected by fn) through the next free variant pack.
func residentVariantDispatch(fn unsafe.Pointer, object, method uintptr, args []gdextension.Variant) (gdextension.Variant, gdextension.CallError) {
	p := &residentVariantPacks[residentVariantDepth]
	residentVariantDepth++
	p.object = object
	p.method = method
	p.argc = int64(len(args))
	if len(args) > 0 {
		copy(p.args[:], unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(args))), len(args)*int(unsafe.Sizeof(gdextension.Variant{}))))
	}
	p.err = gdextension.CallError{}
	fastcb.CallC(fn, unsafe.Pointer(p))
	residentVariantDepth--
	return p.result, p.err
}

// residentVariantAvailable reports whether a variant-shaped call with these
// arguments can take the resident path. Main thread only (callers must check
// threadcheck.Main first).
func residentVariantAvailable(args []gdextension.Variant) bool {
	return fastcb.Resident() &&
		residentVariantDepth < len(residentVariantPacks) &&
		len(args)*int(unsafe.Sizeof(gdextension.Variant{})) <= len(residentVariantPacks[0].args)
}

// residentVariantCall is the variant-call (method_bind_call) analogue of
// residentCall, for script/vararg method calls.
func residentVariantCall(object gdextension.Object, method MethodForClass, args []gdextension.Variant) (gdextension.Variant, error) {
	result, err := residentVariantDispatch(residentVariantCallAddr, uintptr(object), uintptr(method), args)
	return result, err.Err()
}

// ScriptCallResident performs an object script-method call on the resident
// path, reporting false (without calling) when the path is unavailable so the
// caller falls back to the stock Host crossing. See the package comment above
// for the staging rationale — the engine writes both the result and the call
// error through raw pointers after the script ran, and scripts re-enter Go.
func ScriptCallResident(object gdextension.Object, name gdextension.StringName, args []gdextension.Variant) (gdextension.Variant, gdextension.CallError, bool) {
	if !threadcheck.Main() || !residentVariantAvailable(args) {
		return gdextension.Variant{}, gdextension.CallError{}, false
	}
	result, err := residentVariantDispatch(residentScriptCallAddr, uintptr(object), uintptr(name[0]), args)
	return result, err, true
}
