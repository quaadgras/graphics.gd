package classdb

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
)

var debugOwnership = strings.Contains(os.Getenv("GDDEBUG"), "ownership")

func init() {
	gd.ExtensionInstanceLookup = func(obj gdextension.Object) any {
		inst := gdunsafe.Object(obj).ExtensionFetch()
		if inst == nil {
			return nil
		}
		w, ok := inst.(*instanceWrapper)
		if !ok {
			return nil
		}
		ptr, _ := w.impl.Interface()
		return ptr
	}
	gd.ExtensionInstanceGoOnly = func(obj gdextension.Object, goOnly bool) (gdreference.Object, bool) {
		inst := gdunsafe.Object(obj).ExtensionFetch()
		if inst == nil {
			return gdreference.Object{}, false
		}
		w, ok := inst.(*instanceWrapper)
		if !ok {
			return gdreference.Object{}, false
		}
		impl := w.impl
		if debugOwnership {
			var owner string = "Engine"
			if goOnly {
				owner = "Go"
			}
			_, file, line, _ := runtime.Caller(2)
			fmt.Fprintf(os.Stderr, "%s now owned by %s (%s:%d)\n", gd.ObjectGetClass(gdreference.RawObject(obj)).String(), owner, file, line)
		}
		if goOnly {
			impl.strong = nil
		} else {
			impl.strong, _ = impl.Interface()
		}
		val, ok := impl.Interface()
		if !ok {
			return gdreference.Object{}, false
		}
		return val.AsObject()[0], true
	}
	gd.RegisterCleanup(func() {
		for inst := range gdunsafe.ExtensionInstances {
			if w, ok := inst.(*instanceWrapper); ok {
				w.impl.Free()
			}
		}
	})
}

// classWrapper implements gdunsafe.ExtensionClass by wrapping *classImplementation.
type classWrapper struct {
	impl *classImplementation
}

func (w *classWrapper) Create(notify bool) gdunsafe.Object {
	return gdunsafe.Object(gdreference.GetObject(w.impl.CreateInstance(notify)[0]))
}

func (w *classWrapper) Method(name gdunsafe.StringName, hash uint32) gdunsafe.ExtensionFunction {
	sn := gdextension.StringName{gdextension.Pointer(name)}
	virtual, ok := w.impl.GetVirtual(pointers.Let[gd.StringName](sn)).(gd.ExtensionClassCallVirtualFunc)
	if !ok || virtual == nil {
		return nil
	}
	return &virtualFunctionWrapper{fn: virtual}
}

// instanceWrapper implements gdunsafe.ExtensionInstance by wrapping *instanceImplementation.
type instanceWrapper struct {
	impl *instanceImplementation
}

func (w *instanceWrapper) Set(name gdunsafe.StringName, value gdunsafe.Variant) bool {
	sn := gdextension.StringName{gdextension.Pointer(name)}
	return w.impl.Set(pointers.Let[gd.StringName](sn), pointers.Let[gd.Variant](value).Copy())
}

func (w *instanceWrapper) Get(name gdunsafe.StringName) (gdunsafe.Variant, bool) {
	sn := gdextension.StringName{gdextension.Pointer(name)}
	v, ok := w.impl.Get(pointers.Let[gd.StringName](sn))
	if !ok {
		return gdunsafe.Variant{}, false
	}
	raw, ok := pointers.End(v)
	if ok {
		return raw, true
	}
	return pointers.Get(v), true
}

func (w *instanceWrapper) HasDefault(name gdunsafe.StringName) bool {
	sn := gdextension.StringName{gdextension.Pointer(name)}
	return w.impl.PropertyCanRevert(pointers.Let[gd.StringName](sn))
}

func (w *instanceWrapper) GetDefault(name gdunsafe.StringName) (gdunsafe.Variant, bool) {
	sn := gdextension.StringName{gdextension.Pointer(name)}
	v, ok := w.impl.PropertyGetRevert(pointers.Let[gd.StringName](sn))
	if !ok {
		return gdunsafe.Variant{}, false
	}
	raw, ok := pointers.End(v)
	if ok {
		return raw, true
	}
	return pointers.Get(v), true
}

func (w *instanceWrapper) PropertyList() gdunsafe.PropertyList {
	return gdunsafe.PropertyList(w.impl.GetPropertyList())
}

func (w *instanceWrapper) ValidateProperty(name gdunsafe.StringName) bool {
	// The C callback actually passes a PropertyList pointer despite the StringName type.
	return w.impl.ValidateProperty(gdextension.PropertyList(name))
}

func (w *instanceWrapper) Notification(what int32, reverse bool) {
	w.impl.Notification(Object.Notification(what), reverse)
	gdreference.Barrier()
}

func (w *instanceWrapper) UnsafeString() gdunsafe.String {
	s, ok := w.impl.ToString()
	if ok {
		raw, ok := pointers.End(s)
		if ok {
			return gdunsafe.String(raw[0])
		}
		return gdunsafe.String(pointers.Get(s)[0])
	}
	return 0
}

func (w *instanceWrapper) Reference(increment bool) bool {
	if increment {
		w.impl.Reference()
		return true
	}
	return w.impl.Unreference()
}

func (w *instanceWrapper) RID() RID.Any {
	return w.impl.GetRID()
}

func (w *instanceWrapper) Free() {
	w.impl.Free()
}

// virtualFunctionWrapper implements gdunsafe.ExtensionFunction for virtual methods.
type virtualFunctionWrapper struct {
	fn gd.ExtensionClassCallVirtualFunc
}

func (w *virtualFunctionWrapper) PointerCall(inst gdunsafe.ExtensionInstance, args, result gdunsafe.Pointer) {
	receiver := inst.(*instanceWrapper).impl
	ptr, ok := receiver.Interface()
	if !ok {
		return
	}
	w.fn(ptr, gdextension.Pointer(args), gdextension.Pointer(result))
	gdreference.Barrier()
}

func (w *virtualFunctionWrapper) CheckedCall(gdunsafe.ExtensionInstance, gdunsafe.VariadicVariants) gdunsafe.Variant {
	return gdunsafe.Variant{}
}

func (w *virtualFunctionWrapper) DynamicCall(gdunsafe.ExtensionInstance, gdunsafe.VariadicVariants) (gdunsafe.Variant, gdunsafe.CallError) {
	return gdunsafe.Variant{}, gdunsafe.CallError{Type: gdunsafe.CallInvalidMethod}
}

// methodFunctionWrapper implements gdunsafe.ExtensionFunction for registered methods.
type methodFunctionWrapper struct {
	impl *methodImplementation
}

func (w *methodFunctionWrapper) PointerCall(inst gdunsafe.ExtensionInstance, args, result gdunsafe.Pointer) {
	var receiver *instanceImplementation
	if inst != nil {
		receiver = inst.(*instanceWrapper).impl
	}
	w.impl.checked(receiver, gdextension.Pointer(args), gdextension.Pointer(result))
	gdreference.Barrier()
}

func (w *methodFunctionWrapper) CheckedCall(inst gdunsafe.ExtensionInstance, vargs gdunsafe.VariadicVariants) gdunsafe.Variant {
	defer gd.Recover()
	defer gdreference.Barrier()
	var receiver *instanceImplementation
	if inst != nil {
		receiver = inst.(*instanceWrapper).impl
	}
	vargs.Count = w.impl.arg_count
	var variants = make([]gd.Variant, w.impl.arg_count)
	for i := range w.impl.arg_count {
		variants[i] = pointers.Let[gd.Variant](vargs.Index(i))
	}
	v := w.impl.variant(receiver, variants...)
	raw, ok := pointers.End(v)
	if ok {
		return raw
	}
	return pointers.Get(v)
}

func (w *methodFunctionWrapper) DynamicCall(inst gdunsafe.ExtensionInstance, vargs gdunsafe.VariadicVariants) (gdunsafe.Variant, gdunsafe.CallError) {
	defer gd.Recover()
	defer gdreference.Barrier()
	var receiver *instanceImplementation
	if inst != nil {
		receiver = inst.(*instanceWrapper).impl
	}
	var variants = make([]gd.Variant, vargs.Count)
	for i := range vargs.Count {
		variants[i] = pointers.Let[gd.Variant](vargs.Index(i))
	}
	v, err := w.impl.dynamic(receiver, variants...)
	if err != nil {
		return gdunsafe.Variant{}, gdunsafe.CallError{Type: gdunsafe.CallInvalidMethod}
	}
	raw, ok := pointers.End(v)
	if ok {
		return raw, gdunsafe.CallError{}
	}
	return pointers.Get(v), gdunsafe.CallError{}
}
