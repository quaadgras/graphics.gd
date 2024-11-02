package WeakRef

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A weakref can hold a [RefCounted] without contributing to the reference counter. A weakref can be created from an [Object] using [method @GlobalScope.weakref]. If this object is not a reference, weakref still works, however, it does not have any effect on the object. Weakrefs are useful in cases where multiple classes have variables that refer to each other. Without weakrefs, using these classes could lead to memory leaks, since both references keep each other from being released. Making part of the variables a weakref can prevent this cyclic dependency, and allows the references to be released.
*/
type Instance [1]classdb.WeakRef

/*
Returns the [Object] this weakref is referring to. Returns [code]null[/code] if that object no longer exists.
*/
func (self Instance) GetRef() gd.Variant {
	return gd.Variant(class(self).GetRef())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.WeakRef

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WeakRef"))
	return Instance{classdb.WeakRef(object)}
}

/*
Returns the [Object] this weakref is referring to. Returns [code]null[/code] if that object no longer exists.
*/
//go:nosplit
func (self class) GetRef() gd.Variant {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WeakRef.Bind_get_ref, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsWeakRef() Advanced            { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWeakRef() Instance         { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() { classdb.Register("WeakRef", func(ptr gd.Object) any { return classdb.WeakRef(ptr) }) }
