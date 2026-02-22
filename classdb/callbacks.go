package classdb

import (
	"reflect"
	"runtime"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Object"
)

type pinnedVirtualFunc struct {
	fn gd.ExtensionClassCallVirtualFunc
}

var (
	virtualPinner  runtime.Pinner
	pinnedVirtuals []*pinnedVirtualFunc
)

func init() {
	gd.ExtensionInstanceLookup = func(obj gdextension.Object) any {
		val := instances.Get(gdextension.Host.Objects.Extension.Fetch(obj))
		if val == nil {
			return nil
		}
		return val.Value
	}
	gd.ExtensionInstanceGoOnly = func(obj gdextension.Object, goOnly bool) (gdreference.Object, bool) {
		impl := instances.Get(gdextension.Host.Objects.Extension.Fetch(obj))
		if impl == nil {
			return gdreference.Object{}, false
		}
		key := reflect.ValueOf(impl.Value)
		if goOnly {
			roots.Remove(key)
			local.Insert(key, struct{}{})
		} else {
			gdreference.PinObject((*gdreference.Object)(reflect.ValueOf(impl.Value).UnsafePointer()), obj)
			local.Remove(key)
			if keepalive := compile_keepalive(reflect.TypeOf(impl.Value)); keepalive != nil {
				roots.Insert(key, keepalive)
			}
		}
		return impl.Value.AsObject()[0], true
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
				pv := (*pinnedVirtualFunc)(unsafe.Pointer(uintptr(callData)))
				receiver := instances.Get(instance)
				pv.fn(receiver.Value, gdextension.Pointer(args), gdextension.Pointer(result))
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
				defer gd.Recover()
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
				instances.Get(instance).Free()
				instances.Del(instance)
			},
		},
		Class: gdextension.CallbacksForExtensionClass{
			Create: func(class gdextension.ExtensionClassID, notify_postinitialize bool) gdextension.Object {
				return gdreference.GetObject(classes.Get(class).CreateInstance(notify_postinitialize)[0])
			},
			Method: func(class gdextension.ExtensionClassID, method gdextension.StringName, hash uint32) gdextension.FunctionID {
				virtual, ok := classes.Get(class).GetVirtual(pointers.Let[gd.StringName](method)).(gd.ExtensionClassCallVirtualFunc)
				if !ok || virtual == nil {
					return 0
				}
				return methods.New(&methodImplementation{
					checked: func(instance *instanceImplementation, args, ret gdextension.Pointer) {
						virtual(instance.Value, args, ret)
					},
				})
			},
			Caller: func(class gdextension.ExtensionClassID, method gdextension.StringName, hash uint32) uintptr {
				virtual, ok := classes.Get(class).GetVirtual(pointers.Let[gd.StringName](method)).(gd.ExtensionClassCallVirtualFunc)
				if !ok || virtual == nil {
					return 0
				}
				pv := &pinnedVirtualFunc{fn: virtual}
				virtualPinner.Pin(pv)
				pinnedVirtuals = append(pinnedVirtuals, pv)
				return uintptr(unsafe.Pointer(pv))
			},
		},
	}
}
