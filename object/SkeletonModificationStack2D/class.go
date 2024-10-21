package SkeletonModificationStack2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource is used by the Skeleton and holds a stack of [SkeletonModification2D]s.
This controls the order of the modifications and how they are applied. Modification order is especially important for full-body IK setups, as you need to execute the modifications in the correct order to get the desired results. For example, you want to execute a modification on the spine [i]before[/i] the arms on a humanoid skeleton.
This resource also controls how strongly all of the modifications are applied to the [Skeleton2D].

*/
type Simple [1]classdb.SkeletonModificationStack2D
func (self Simple) Setup() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Setup()
}
func (self Simple) Execute(delta float64, execution_mode int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Execute(gd.Float(delta), gd.Int(execution_mode))
}
func (self Simple) EnableAllModifications(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnableAllModifications(enabled)
}
func (self Simple) GetModification(mod_idx int) [1]classdb.SkeletonModification2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SkeletonModification2D(Expert(self).GetModification(gc, gd.Int(mod_idx)))
}
func (self Simple) AddModification(modification [1]classdb.SkeletonModification2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddModification(modification)
}
func (self Simple) DeleteModification(mod_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeleteModification(gd.Int(mod_idx))
}
func (self Simple) SetModification(mod_idx int, modification [1]classdb.SkeletonModification2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetModification(gd.Int(mod_idx), modification)
}
func (self Simple) SetModificationCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetModificationCount(gd.Int(count))
}
func (self Simple) GetModificationCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetModificationCount()))
}
func (self Simple) GetIsSetup() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetIsSetup())
}
func (self Simple) SetEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnabled(enabled)
}
func (self Simple) GetEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnabled())
}
func (self Simple) SetStrength(strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStrength(gd.Float(strength))
}
func (self Simple) GetStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetStrength()))
}
func (self Simple) GetSkeleton() [1]classdb.Skeleton2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Skeleton2D(Expert(self).GetSkeleton(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModificationStack2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets up the modification stack so it can execute. This function should be called by [Skeleton2D] and shouldn't be manually called unless you know what you are doing.
*/
//go:nosplit
func (self class) Setup()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Executes all of the [SkeletonModification2D]s in the stack that use the same execution mode as the passed-in [param execution_mode], starting from index [code]0[/code] to [member modification_count].
[b]Note:[/b] The order of the modifications can matter depending on the modifications. For example, modifications on a spine should operate before modifications on the arms in order to get proper results.
*/
//go:nosplit
func (self class) Execute(delta gd.Float, execution_mode gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, execution_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables all [SkeletonModification2D]s in the stack.
*/
//go:nosplit
func (self class) EnableAllModifications(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_enable_all_modifications, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [SkeletonModification2D] at the passed-in index, [param mod_idx].
*/
//go:nosplit
func (self class) GetModification(ctx gd.Lifetime, mod_idx gd.Int) object.SkeletonModification2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_get_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SkeletonModification2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds the passed-in [SkeletonModification2D] to the stack.
*/
//go:nosplit
func (self class) AddModification(modification object.SkeletonModification2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(modification[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_add_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes the [SkeletonModification2D] at the index position [param mod_idx], if it exists.
*/
//go:nosplit
func (self class) DeleteModification(mod_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_delete_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the modification at [param mod_idx] to the passed-in modification, [param modification].
*/
//go:nosplit
func (self class) SetModification(mod_idx gd.Int, modification object.SkeletonModification2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	callframe.Arg(frame, mmm.Get(modification[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_set_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetModificationCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_set_modification_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModificationCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_get_modification_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a boolean that indicates whether the modification stack is setup and can execute.
*/
//go:nosplit
func (self class) GetIsSetup() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_get_is_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_set_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_get_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Skeleton2D] node that the SkeletonModificationStack2D is bound to.
*/
//go:nosplit
func (self class) GetSkeleton(ctx gd.Lifetime) object.Skeleton2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModificationStack2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Skeleton2D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModificationStack2D() Expert { return self[0].AsSkeletonModificationStack2D() }


//go:nosplit
func (self Simple) AsSkeletonModificationStack2D() Simple { return self[0].AsSkeletonModificationStack2D() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SkeletonModificationStack2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
