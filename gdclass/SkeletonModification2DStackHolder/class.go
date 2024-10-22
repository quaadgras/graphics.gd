package SkeletonModification2DStackHolder

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SkeletonModification2D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This [SkeletonModification2D] holds a reference to a [SkeletonModificationStack2D], allowing you to use multiple modification stacks on a single [Skeleton2D].
[b]Note:[/b] The modifications in the held [SkeletonModificationStack2D] will only be executed if their execution mode matches the execution mode of the SkeletonModification2DStackHolder.

*/
type Go [1]classdb.SkeletonModification2DStackHolder

/*
Sets the [SkeletonModificationStack2D] that this modification is holding. This modification stack will then be executed when this modification is executed.
*/
func (self Go) SetHeldModificationStack(held_modification_stack gdclass.SkeletonModificationStack2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHeldModificationStack(held_modification_stack)
}

/*
Returns the [SkeletonModificationStack2D] that this modification is holding.
*/
func (self Go) GetHeldModificationStack() gdclass.SkeletonModificationStack2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.SkeletonModificationStack2D(class(self).GetHeldModificationStack(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DStackHolder
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2DStackHolder"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Sets the [SkeletonModificationStack2D] that this modification is holding. This modification stack will then be executed when this modification is executed.
*/
//go:nosplit
func (self class) SetHeldModificationStack(held_modification_stack gdclass.SkeletonModificationStack2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(held_modification_stack[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DStackHolder.Bind_set_held_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [SkeletonModificationStack2D] that this modification is holding.
*/
//go:nosplit
func (self class) GetHeldModificationStack(ctx gd.Lifetime) gdclass.SkeletonModificationStack2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DStackHolder.Bind_get_held_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SkeletonModificationStack2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DStackHolder() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DStackHolder() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSkeletonModification2D() SkeletonModification2D.GD { return *((*SkeletonModification2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2D() SkeletonModification2D.Go { return *((*SkeletonModification2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}
func init() {classdb.Register("SkeletonModification2DStackHolder", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
