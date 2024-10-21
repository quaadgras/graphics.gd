package SkeletonModification2DPhysicalBones

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
This modification takes the transforms of [PhysicalBone2D] nodes and applies them to [Bone2D] nodes. This allows the [Bone2D] nodes to react to physics thanks to the linked [PhysicalBone2D] nodes.

*/
type Simple [1]classdb.SkeletonModification2DPhysicalBones
func (self Simple) SetPhysicalBoneChainLength(length int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicalBoneChainLength(gd.Int(length))
}
func (self Simple) GetPhysicalBoneChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPhysicalBoneChainLength()))
}
func (self Simple) SetPhysicalBoneNode(joint_idx int, physicalbone2d_node gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPhysicalBoneNode(gd.Int(joint_idx), physicalbone2d_node)
}
func (self Simple) GetPhysicalBoneNode(joint_idx int) gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetPhysicalBoneNode(gc, gd.Int(joint_idx)))
}
func (self Simple) FetchPhysicalBones() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FetchPhysicalBones()
}
func (self Simple) StartSimulation(bones gd.ArrayOf[gd.StringName]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StartSimulation(bones)
}
func (self Simple) StopSimulation(bones gd.ArrayOf[gd.StringName]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StopSimulation(bones)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DPhysicalBones
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPhysicalBoneChainLength(length gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_set_physical_bone_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicalBoneChainLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_get_physical_bone_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [PhysicalBone2D] node at [param joint_idx].
[b]Note:[/b] This is just the index used for this modification, not the bone index used in the [Skeleton2D].
*/
//go:nosplit
func (self class) SetPhysicalBoneNode(joint_idx gd.Int, physicalbone2d_node gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mmm.Get(physicalbone2d_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_set_physical_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [PhysicalBone2D] node at [param joint_idx].
*/
//go:nosplit
func (self class) GetPhysicalBoneNode(ctx gd.Lifetime, joint_idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_get_physical_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Empties the list of [PhysicalBone2D] nodes and populates it with all [PhysicalBone2D] nodes that are children of the [Skeleton2D].
*/
//go:nosplit
func (self class) FetchPhysicalBones()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_fetch_physical_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
//go:nosplit
func (self class) StartSimulation(bones gd.ArrayOf[gd.StringName])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bones))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_start_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
//go:nosplit
func (self class) StopSimulation(bones gd.ArrayOf[gd.StringName])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bones))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DPhysicalBones.Bind_stop_simulation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsSkeletonModification2DPhysicalBones() Expert { return self[0].AsSkeletonModification2DPhysicalBones() }


//go:nosplit
func (self Simple) AsSkeletonModification2DPhysicalBones() Simple { return self[0].AsSkeletonModification2DPhysicalBones() }


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
func init() {classdb.Register("SkeletonModification2DPhysicalBones", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
