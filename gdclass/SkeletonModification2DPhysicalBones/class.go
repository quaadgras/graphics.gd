package SkeletonModification2DPhysicalBones

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
This modification takes the transforms of [PhysicalBone2D] nodes and applies them to [Bone2D] nodes. This allows the [Bone2D] nodes to react to physics thanks to the linked [PhysicalBone2D] nodes.

*/
type Go [1]classdb.SkeletonModification2DPhysicalBones

/*
Sets the [PhysicalBone2D] node at [param joint_idx].
[b]Note:[/b] This is just the index used for this modification, not the bone index used in the [Skeleton2D].
*/
func (self Go) SetPhysicalBoneNode(joint_idx int, physicalbone2d_node string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicalBoneNode(gd.Int(joint_idx), gc.String(physicalbone2d_node).NodePath(gc))
}

/*
Returns the [PhysicalBone2D] node at [param joint_idx].
*/
func (self Go) GetPhysicalBoneNode(joint_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetPhysicalBoneNode(gc, gd.Int(joint_idx)).String())
}

/*
Empties the list of [PhysicalBone2D] nodes and populates it with all [PhysicalBone2D] nodes that are children of the [Skeleton2D].
*/
func (self Go) FetchPhysicalBones() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FetchPhysicalBones()
}

/*
Tell the [PhysicalBone2D] nodes to start simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to start simulating.
*/
func (self Go) StartSimulation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).StartSimulation(([1]gd.ArrayOf[gd.StringName]{}[0]))
}

/*
Tell the [PhysicalBone2D] nodes to stop simulating and interacting with the physics world.
Optionally, an array of bone names can be passed to this function, and that will cause only [PhysicalBone2D] nodes with those names to stop simulating.
*/
func (self Go) StopSimulation() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).StopSimulation(([1]gd.ArrayOf[gd.StringName]{}[0]))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DPhysicalBones
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SkeletonModification2DPhysicalBones"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PhysicalBoneChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPhysicalBoneChainLength()))
}

func (self Go) SetPhysicalBoneChainLength(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPhysicalBoneChainLength(gd.Int(value))
}

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
func (self class) AsSkeletonModification2DPhysicalBones() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DPhysicalBones() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SkeletonModification2DPhysicalBones", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
