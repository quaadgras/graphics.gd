package classdb

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/internal/threadsafe"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type pinnedVirtualFunc struct {
	fn gd.ExtensionClassCallVirtualFunc

	// tab is the receiver class's static itab (classImplementation.tab):
	// with it, dispatch rebuilds the receiver interface directly from the
	// instance word (fastInterface) instead of resolving the instance
	// record. Nil on portable builds, which fall back to the table.
	tab unsafe.Pointer

	// tick is the per-frame virtuals' shortcut: _process and
	// _physics_process are called once per node per frame, so a scene runs
	// them tens of thousands of times where it runs any other virtual once,
	// and they are the only ones whose dispatch is worth specialising. It
	// calls the user's method with the instance word directly, skipping the
	// receiver interface, the conversion into the generic wrapper's any
	// parameter and the assertion back out of it — none of which recover
	// anything the instance word did not already carry. Nil unless the
	// method has exactly the func(*T, Float.X) shape [classImplementation.tickOf]
	// looks for; the generic path handles everything else.
	tick func(unsafe.Pointer, Float.X)
}

var (
	virtualPinner  runtime.Pinner
	pinnedVirtuals []*pinnedVirtualFunc
	// pinnedVirtualCache reuses a pinned dispatch record per (class,
	// interned method name): the engine resolves virtuals per instance,
	// so without it every node entering the tree pins a fresh record for
	// the same answer.
	pinnedVirtualCache threadsafe.Map[[2]uintptr, uintptr]
)

var debugOwnership = strings.Contains(os.Getenv("GDDEBUG"), "ownership")

// tickDisabled sends _process and _physics_process down the generic virtual
// path instead of [pinnedVirtualFunc.tick]. The shortcut is worth a few
// nanoseconds per node per frame, which is small enough that only an A/B in
// the same binary can measure it honestly — comparing two builds moves the
// answer by more than the effect (see BenchmarkVirtualProcess in
// internal/virtual_process_test.go). This makes that A/B possible, and gives
// anyone who suspects the shortcut a way to switch it off.
var tickDisabled = strings.Contains(os.Getenv("GDDEBUG"), "notick")

func init() {
	gd.ExtensionInstanceLookup = func(obj gdextension.Object) any {
		val := instances.Get(gdextension.Host.Objects.Extension.Fetch(obj))
		if val == nil {
			return nil
		}
		ptr, _ := val.Interface()
		return ptr
	}
	gd.ExtensionInstanceGoOnly = func(obj gdextension.Object, goOnly bool) (gdreference.Object, bool) {
		impl := instances.Get(gdextension.Host.Objects.Extension.Fetch(obj))
		if impl == nil {
			return gdreference.Object{}, false
		}
		if debugOwnership {
			var owner string = "Engine"
			if goOnly {
				owner = "Go"
			}
			_, file, line, _ := runtime.Caller(2)
			fmt.Fprintf(os.Stderr, "%s now owned by %s (%s:%d)\n", gd.ObjectGetClass(gdreference.RawObject(obj)).String(), owner, file, line)
		}
		if goOnly {
			// Go owns the object now: drop the strong root AND the
			// dispatch-word pin so the GC may collect the wrapper (its
			// cleanup then frees the engine object). The engine no
			// longer holds meaningful references, so it will not
			// dispatch on the (now unpinned) instance word.
			impl.strong = nil
			impl.pinner.Unpin()
		} else {
			impl.strong, _ = impl.Interface()
			repinInstance(impl)
		}
		val, ok := impl.Interface()
		if !ok {
			return gdreference.Object{}, false
		}
		return val.AsObject()[0], true
	}
	gd.RegisterCleanup(func() {
		for instance := range instances.All {
			if instance != nil {
				instance.Free()
			}
		}
		virtualPinner.Unpin()
		virtualPinner = runtime.Pinner{}
		pinnedVirtuals = nil
		pinnedVirtualCache = threadsafe.Map[[2]uintptr, uintptr]{}
	})

	gdextension.On.Extension = gdextension.CallbacksForExtension{
		Binding: gdextension.CallbacksForExtensionBinding{
			Created: func(instance gdextension.ExtensionInstanceID) gdextension.ExtensionBindingID {
				return 0
			},
			Removed: func(instance gdextension.ExtensionInstanceID, binding gdextension.ExtensionBindingID) {

			},
			Reference: func(instance gdextension.ExtensionInstanceID, increment bool) bool {
				return false
			},
		},
		Instance: gdextension.CallbacksForExtensionInstance{
			Set: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName, value gdextension.Variant) bool {
				return instances.Get(instance).Set(pointers.Let[gd.StringName](field), pointers.Let[gd.Variant](value).Copy())
			},
			Get: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName, result gdextension.Returns[gdextension.Variant]) bool {
				v, ok := instances.Get(instance).Get(pointers.Let[gd.StringName](field))
				if !ok {
					return false
				}
				raw, ok := pointers.End(v)
				if ok {
					gdmemory.Set(gdextension.Pointer(result), raw)
				} else {
					gdmemory.Set(gdextension.Pointer(result), pointers.Get(v))
				}
				return true
			},
			PropertyList: func(instance gdextension.ExtensionInstanceID) gdextension.PropertyList {
				return instances.Get(instance).GetPropertyList()
			},
			PropertyValidation: func(instance gdextension.ExtensionInstanceID, list gdextension.PropertyList) bool {
				return instances.Get(instance).ValidateProperty(list)
			},
			PropertyHasDefault: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName) bool {
				return instances.Get(instance).PropertyCanRevert(pointers.Let[gd.StringName](field))
			},
			PropertyGetDefault: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName, result gdextension.Returns[gdextension.Variant]) bool {
				v, ok := instances.Get(instance).PropertyGetRevert(pointers.Let[gd.StringName](field))
				if ok {
					raw, ok := pointers.End(v)
					if ok {
						gdmemory.Set(gdextension.Pointer(result), raw)
					} else {
						gdmemory.Set(gdextension.Pointer(result), pointers.Get(v))
					}
				}
				return ok
			},
			Stringify: func(instance gdextension.ExtensionInstanceID) gdextension.String {
				s, ok := instances.Get(instance).ToString()
				if ok {
					raw, ok := pointers.End(s)
					if ok {
						return raw
					} else {
						return pointers.Get(s)
					}
				}
				return gdextension.String{}
			},
			Reference: func(instance gdextension.ExtensionInstanceID, increment bool) bool {
				if increment {
					instances.Get(instance).Reference()
					return true
				}
				return instances.Get(instance).Unreference()
			},
			RID: func(instance gdextension.ExtensionInstanceID, rid gdextension.Returns[uint64]) {
				gdmemory.Set(gdextension.Pointer(rid), uint64(0))
			},
			Notification: func(instance gdextension.ExtensionInstanceID, what int32, reverse bool) {
				instances.Get(instance).Notification(Object.Notification(what), reverse)
				gdreference.Barrier()
			},
			CheckedCall: func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[any], args gdextension.Accepts[any]) {
				//defer gd.Recover()
				var receiver *instanceImplementation
				if instance != 0 {
					receiver = instances.Get(instance)
				}
				methods.Get(fn).checked(receiver, gdextension.Pointer(args), gdextension.Pointer(result))
				gdreference.Barrier()
			},
			Called: func(instance gdextension.ExtensionInstanceID, callData gdextension.Pointer, result gdextension.Returns[any], args gdextension.Accepts[any]) {
				pv := (*pinnedVirtualFunc)(*(*unsafe.Pointer)(unsafe.Pointer(&callData))) // runtime.Pinned, so this is ok.
				if pv.tick != nil && instance != 0 {
					pv.tick(unsafe.Pointer(uintptr(instance)), Float.X(gd.UnsafeGet[float64](gdextension.Pointer(args), 0)))
					gdreference.Barrier()
					return
				}
				if ptr, ok := fastInterface(pv.tab, instance); ok {
					pv.fn(ptr, gdextension.Pointer(args), gdextension.Pointer(result))
					gdreference.Barrier()
					return
				}
				receiver := instances.Get(instance)
				if receiver == nil {
					return
				}
				ptr, ok := receiver.Interface()
				if !ok {
					return
				}
				pv.fn(ptr, gdextension.Pointer(args), gdextension.Pointer(result))
				gdreference.Barrier()
			},
			VariantCall: func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], args gdextension.Accepts[gdextension.Variant]) {
				defer gd.Recover()
				var receiver *instanceImplementation
				if instance != 0 {
					receiver = instances.Get(instance)
				}
				method := methods.Get(fn)
				var variants = make([]gd.Variant, method.arg_count)
				for i := range method.arg_count {
					variants[i] = pointers.Let[gd.Variant](gdmemory.IndexVariants(args, method.arg_count, i))
				}
				v := method.variant(receiver, variants...)
				raw, ok := pointers.End(v)
				if ok {
					gdmemory.Set(gdextension.Pointer(result), raw)
				} else {
					gdmemory.Set(gdextension.Pointer(result), pointers.Get(v))
				}
				gdreference.Barrier()
			},
			DynamicCall: func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], arg_count int, args gdextension.Accepts[gdextension.Variant], call_err gdextension.Returns[gdextension.CallError]) {
				defer gd.RecoverCall(call_err)
				var receiver *instanceImplementation
				if instance != 0 {
					receiver = instances.Get(instance)
				}
				var variants = make([]gd.Variant, arg_count)
				for i := range arg_count {
					variants[i] = pointers.Let[gd.Variant](gdmemory.IndexVariants(args, arg_count, i))
				}
				v, err := methods.Get(fn).dynamic(receiver, variants...)
				if err != nil {
					gdmemory.Set(gdextension.Pointer(call_err), gdextension.CallError{
						Type: gdextension.CallInvalidMethod,
					})
					gdreference.Barrier()
					return
				}
				raw, ok := pointers.End(v)
				if ok {
					gdmemory.Set(gdextension.Pointer(result), raw)
				} else {
					gdmemory.Set(gdextension.Pointer(result), pointers.Get(v))
				}
				gdreference.Barrier()
			},
			Free: func(instance gdextension.ExtensionInstanceID) {
				if impl := instances.Get(instance); impl != nil {
					impl.Free()
				}
				instances.Del(instance)
			},
		},
		Class: gdextension.CallbacksForExtensionClass{
			Create: func(class gdextension.ExtensionClassID, notify_postinitialize bool) gdextension.Object {
				obj := gdreference.GetObject(classes.Get(class).CreateInstance(notify_postinitialize)[0])
				if threadcheck.Main() && ring.Main.Pending() {
					ring.Main.Flush()
				}
				return obj
			},
			Method: func(class gdextension.ExtensionClassID, method gdextension.StringName, hash uint32) gdextension.FunctionID {
				virtual, ok := classes.Get(class).GetVirtual(pointers.Let[gd.StringName](method)).(gd.ExtensionClassCallVirtualFunc)
				if !ok || virtual == nil {
					return 0
				}
				return methods.New(&methodImplementation{
					checked: func(instance *instanceImplementation, args, ret gdextension.Pointer) {
						ptr, ok := instance.Interface()
						if !ok {
							return
						}
						virtual(ptr, args, ret)
					},
				})
			},
			Caller: func(class gdextension.ExtensionClassID, method gdextension.StringName, hash uint32) uintptr {
				key := [2]uintptr{uintptr(class), uintptr(method[0])}
				if cached, ok := pinnedVirtualCache.Lookup(key); ok {
					return cached
				}
				classImpl := classes.Get(class)
				name := pointers.Let[gd.StringName](method)
				virtual, ok := classImpl.GetVirtual(name).(gd.ExtensionClassCallVirtualFunc)
				if !ok || virtual == nil {
					pinnedVirtualCache.Insert(key, 0)
					return 0
				}
				pv := &pinnedVirtualFunc{fn: virtual, tab: classImpl.tab, tick: classImpl.tickOf(name.String())}
				virtualPinner.Pin(pv)
				pinnedVirtuals = append(pinnedVirtuals, pv)
				pinnedVirtualCache.Insert(key, uintptr(unsafe.Pointer(pv)))
				return uintptr(unsafe.Pointer(pv))
			},
		},
	}
}
