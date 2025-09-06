package classdb

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/pointers"
)

func init() {
	gd.ExtensionInstanceLookup = func(obj gdextension.Object) any {
		val, ok := handles.Load(uintptr(gdextension.Host.Objects.Extension.Fetch(obj)))
		if !ok {
			return nil
		}
		return val.(*instanceImplementation).Value
	}
	gd.RegisterCleanup(func() {
		handles.Range(func(key any, value any) bool {
			if instance, ok := value.(*instanceImplementation); ok {
				instance.Free()
			}
			return true
		})
	})

	gdextension.On.Extension = gdextension.CallbacksForExtension{
		Binding: gdextension.CallbacksForExtensionBinding{
			Created: func(instance gdextension.ExtensionInstanceID) gdextension.ExtensionBindingID {
				return gdextension.ExtensionBindingID(instance)
			},
			Removed: func(instance gdextension.ExtensionInstanceID, binding gdextension.ExtensionBindingID) {

			},
			Reference: func(instance gdextension.ExtensionInstanceID, increment bool) bool {
				return false
			},
		},
		Instance: gdextension.CallbacksForExtensionInstance{
			Set: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName, value gdextension.Variant) bool {
				return cgoHandle(instance).Value().(*instanceImplementation).Set(pointers.Let[gd.StringName](field), pointers.Let[gd.Variant](value).Copy())
			},
			Get: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName, result gdextension.Returns[gdextension.Variant]) bool {
				v, ok := cgoHandle(instance).Value().(*instanceImplementation).Get(pointers.Let[gd.StringName](field))
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
				return cgoHandle(instance).Value().(*instanceImplementation).GetPropertyList()
			},
			PropertyValidation: func(instance gdextension.ExtensionInstanceID, list gdextension.PropertyList) bool {
				return cgoHandle(instance).Value().(*instanceImplementation).ValidateProperty(list)
			},
			PropertyHasDefault: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName) bool {
				return cgoHandle(instance).Value().(*instanceImplementation).PropertyCanRevert(pointers.Let[gd.StringName](field))
			},
			PropertyGetDefault: func(instance gdextension.ExtensionInstanceID, field gdextension.StringName, result gdextension.Returns[gdextension.Variant]) bool {
				v, ok := cgoHandle(instance).Value().(*instanceImplementation).PropertyGetRevert(pointers.Let[gd.StringName](field))
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
				s, ok := cgoHandle(instance).Value().(*instanceImplementation).ToString()
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
					cgoHandle(instance).Value().(*instanceImplementation).Reference()
					return true
				}
				return cgoHandle(instance).Value().(*instanceImplementation).Unreference()
			},
			RID: func(instance gdextension.ExtensionInstanceID, rid gdextension.Returns[uint64]) {
				gdmemory.Set(gdextension.Pointer(rid), uint64(0))
			},
			Notification: func(instance gdextension.ExtensionInstanceID, what int32, reverse bool) {
				cgoHandle(instance).Value().(*instanceImplementation).Notification(what, reverse)
			},
			CheckedCall: func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[any], args gdextension.Accepts[any]) {
				defer gd.Recover()
				cgoHandle(fn).Value().(*methodImplementation).checked(cgoHandle(instance).Value(), gdextension.Pointer(args), gdextension.Pointer(result))
			},
			VariantCall: func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], args gdextension.Accepts[gdextension.Variant]) {
				defer gd.Recover()
				method := cgoHandle(fn).Value().(*methodImplementation)
				var variants = make([]gd.Variant, method.arg_count)
				for i := range method.arg_count {
					variants[i] = pointers.Let[gd.Variant](gdmemory.IndexVariants(args, method.arg_count, i))
				}
				v := method.variant(cgoHandle(instance).Value(), variants...)
				raw, ok := pointers.End(v)
				if ok {
					gdmemory.Set(gdextension.Pointer(result), raw)
				} else {
					gdmemory.Set(gdextension.Pointer(result), pointers.Get(v))
				}
			},
			DynamicCall: func(instance gdextension.ExtensionInstanceID, fn gdextension.FunctionID, result gdextension.Returns[gdextension.Variant], arg_count int, args gdextension.Accepts[gdextension.Variant], call_err gdextension.Returns[gdextension.CallError]) {
				defer gd.Recover()
				var variants = make([]gd.Variant, arg_count)
				for i := range arg_count {
					variants[i] = pointers.Let[gd.Variant](gdmemory.IndexVariants(args, arg_count, i))
				}
				v, err := cgoHandle(fn).Value().(*methodImplementation).dynamic(cgoHandle(instance).Value(), variants...)
				if err != nil {
					gdmemory.Set(gdextension.Pointer(call_err), gdextension.CallError{
						Type: gdextension.CallInvalidMethod,
					})
					return
				}
				raw, ok := pointers.End(v)
				if ok {
					gdmemory.Set(gdextension.Pointer(result), raw)
				} else {
					gdmemory.Set(gdextension.Pointer(result), pointers.Get(v))
				}
			},
			Free: func(instance gdextension.ExtensionInstanceID) {
				cgoHandle(instance).Value().(*instanceImplementation).Free()
				cgoHandle(instance).Delete()
			},
		},
		Class: gdextension.CallbacksForExtensionClass{
			Create: func(class gdextension.ExtensionClassID, notify_postinitialize bool) gdextension.Object {
				return gdextension.Object(pointers.Get(cgoHandle(class).Value().(*classImplementation).CreateInstance(notify_postinitialize)[0])[0])
			},
			Method: func(class gdextension.ExtensionClassID, method gdextension.StringName, hash uint32) gdextension.FunctionID {
				virtual := cgoHandle(class).Value().(*classImplementation).GetVirtual(pointers.Let[gd.StringName](method))
				if virtual == nil {
					return 0
				}
				return gdextension.FunctionID(cgoNewHandle(&methodImplementation{
					checked: func(instance any, args, ret gdextension.Pointer) {
						instance.(*instanceImplementation).CallVirtual(virtual, args, ret)
					},
				}))
			},
		},
	}
}
