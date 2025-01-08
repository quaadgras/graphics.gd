// Package SkeletonModification2DStackHolder provides methods for working with SkeletonModification2DStackHolder object instances.
package SkeletonModification2DStackHolder

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/SkeletonModification2D"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This [SkeletonModification2D] holds a reference to a [SkeletonModificationStack2D], allowing you to use multiple modification stacks on a single [Skeleton2D].
[b]Note:[/b] The modifications in the held [SkeletonModificationStack2D] will only be executed if their execution mode matches the execution mode of the SkeletonModification2DStackHolder.
*/
type Instance [1]gdclass.SkeletonModification2DStackHolder

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModification2DStackHolder() Instance
}

/*
Sets the [SkeletonModificationStack2D] that this modification is holding. This modification stack will then be executed when this modification is executed.
*/
func (self Instance) SetHeldModificationStack(held_modification_stack [1]gdclass.SkeletonModificationStack2D) {
	class(self).SetHeldModificationStack(held_modification_stack)
}

/*
Returns the [SkeletonModificationStack2D] that this modification is holding.
*/
func (self Instance) GetHeldModificationStack() [1]gdclass.SkeletonModificationStack2D {
	return [1]gdclass.SkeletonModificationStack2D(class(self).GetHeldModificationStack())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModification2DStackHolder

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DStackHolder"))
	casted := Instance{*(*gdclass.SkeletonModification2DStackHolder)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sets the [SkeletonModificationStack2D] that this modification is holding. This modification stack will then be executed when this modification is executed.
*/
//go:nosplit
func (self class) SetHeldModificationStack(held_modification_stack [1]gdclass.SkeletonModificationStack2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(held_modification_stack[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DStackHolder.Bind_set_held_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [SkeletonModificationStack2D] that this modification is holding.
*/
//go:nosplit
func (self class) GetHeldModificationStack() [1]gdclass.SkeletonModificationStack2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DStackHolder.Bind_get_held_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.SkeletonModificationStack2D{gd.PointerWithOwnershipTransferredToGo[gdclass.SkeletonModificationStack2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DStackHolder() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DStackHolder() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSkeletonModification2D() SkeletonModification2D.Advanced {
	return *((*SkeletonModification2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2D() SkeletonModification2D.Instance {
	return *((*SkeletonModification2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModification2D.Advanced(self.AsSkeletonModification2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModification2D.Instance(self.AsSkeletonModification2D()), name)
	}
}
func init() {
	gdclass.Register("SkeletonModification2DStackHolder", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModification2DStackHolder{*(*gdclass.SkeletonModification2DStackHolder)(unsafe.Pointer(&ptr))}
	})
}
