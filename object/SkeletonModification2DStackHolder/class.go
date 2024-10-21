package SkeletonModification2DStackHolder

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/SkeletonModification2D"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This [SkeletonModification2D] holds a reference to a [SkeletonModificationStack2D], allowing you to use multiple modification stacks on a single [Skeleton2D].
[b]Note:[/b] The modifications in the held [SkeletonModificationStack2D] will only be executed if their execution mode matches the execution mode of the SkeletonModification2DStackHolder.

*/
type Simple [1]classdb.SkeletonModification2DStackHolder
func (self Simple) SetHeldModificationStack(held_modification_stack [1]classdb.SkeletonModificationStack2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeldModificationStack(held_modification_stack)
}
func (self Simple) GetHeldModificationStack() [1]classdb.SkeletonModificationStack2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkeletonModificationStack2D(Expert(self).GetHeldModificationStack(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DStackHolder
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the [SkeletonModificationStack2D] that this modification is holding. This modification stack will then be executed when this modification is executed.
*/
//go:nosplit
func (self class) SetHeldModificationStack(held_modification_stack object.SkeletonModificationStack2D)  {
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
func (self class) GetHeldModificationStack(ctx gd.Lifetime) object.SkeletonModificationStack2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DStackHolder.Bind_get_held_modification_stack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkeletonModificationStack2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModification2DStackHolder() Expert { return self[0].AsSkeletonModification2DStackHolder() }


//go:nosplit
func (self Simple) AsSkeletonModification2DStackHolder() Simple { return self[0].AsSkeletonModification2DStackHolder() }


//go:nosplit
func (self class) AsSkeletonModification2D() SkeletonModification2D.Expert { return self[0].AsSkeletonModification2D() }


//go:nosplit
func (self Simple) AsSkeletonModification2D() SkeletonModification2D.Simple { return self[0].AsSkeletonModification2D() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SkeletonModification2DStackHolder", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
