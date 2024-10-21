package SkeletonModification2DFABRIK

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
This [SkeletonModification2D] uses an algorithm called Forward And Backward Reaching Inverse Kinematics, or FABRIK, to rotate a bone chain so that it reaches a target.
FABRIK works by knowing the positions and lengths of a series of bones, typically called a "bone chain". It first starts by running a forward pass, which places the final bone at the target's position. Then all other bones are moved towards the tip bone, so they stay at the defined bone length away. Then a backwards pass is performed, where the root/first bone in the FABRIK chain is placed back at the origin. Then all other bones are moved so they stay at the defined bone length away. This positions the bone chain so that it reaches the target when possible, but all of the bones stay the correct length away from each other.
Because of how FABRIK works, it often gives more natural results than those seen in [SkeletonModification2DCCDIK]. FABRIK also supports angle constraints, which are fully taken into account when solving.
[b]Note:[/b] The FABRIK modifier has [code]fabrik_joints[/code], which are the data objects that hold the data for each joint in the FABRIK chain. This is different from [Bone2D] nodes! FABRIK joints hold the data needed for each [Bone2D] in the bone chain used by FABRIK.
To help control how the FABRIK joints move, a magnet vector can be passed, which can nudge the bones in a certain direction prior to solving, giving a level of control over the final result.

*/
type Simple [1]classdb.SkeletonModification2DFABRIK
func (self Simple) SetTargetNode(target_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetNode(target_nodepath)
}
func (self Simple) GetTargetNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetTargetNode(gc))
}
func (self Simple) SetFabrikDataChainLength(length int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFabrikDataChainLength(gd.Int(length))
}
func (self Simple) GetFabrikDataChainLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFabrikDataChainLength()))
}
func (self Simple) SetFabrikJointBone2dNode(joint_idx int, bone2d_nodepath gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFabrikJointBone2dNode(gd.Int(joint_idx), bone2d_nodepath)
}
func (self Simple) GetFabrikJointBone2dNode(joint_idx int) gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetFabrikJointBone2dNode(gc, gd.Int(joint_idx)))
}
func (self Simple) SetFabrikJointBoneIndex(joint_idx int, bone_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFabrikJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}
func (self Simple) GetFabrikJointBoneIndex(joint_idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFabrikJointBoneIndex(gd.Int(joint_idx))))
}
func (self Simple) SetFabrikJointMagnetPosition(joint_idx int, magnet_position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFabrikJointMagnetPosition(gd.Int(joint_idx), magnet_position)
}
func (self Simple) GetFabrikJointMagnetPosition(joint_idx int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetFabrikJointMagnetPosition(gd.Int(joint_idx)))
}
func (self Simple) SetFabrikJointUseTargetRotation(joint_idx int, use_target_rotation bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFabrikJointUseTargetRotation(gd.Int(joint_idx), use_target_rotation)
}
func (self Simple) GetFabrikJointUseTargetRotation(joint_idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFabrikJointUseTargetRotation(gd.Int(joint_idx)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SkeletonModification2DFABRIK
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(target_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFabrikDataChainLength(length gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFabrikDataChainLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetFabrikJointBone2dNode(joint_idx gd.Int, bone2d_nodepath gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mmm.Get(bone2d_nodepath))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointBone2dNode(ctx gd.Lifetime, joint_idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone index, [param bone_idx], of the FABRIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the FABRIK joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetFabrikJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointBoneIndex(joint_idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the magnet position vector for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetFabrikJointMagnetPosition(joint_idx gd.Int, magnet_position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, magnet_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_magnet_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the magnet position vector for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointMagnetPosition(joint_idx gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_magnet_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the joint at [param joint_idx] will use the target node's rotation rather than letting FABRIK rotate the node.
[b]Note:[/b] This option only works for the tip/final joint in the chain. For all other nodes, this option will be ignored.
*/
//go:nosplit
func (self class) SetFabrikJointUseTargetRotation(joint_idx gd.Int, use_target_rotation bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_target_rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_use_target_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the joint is using the target's rotation rather than allowing FABRIK to rotate the joint. This option only applies to the tip/final joint in the chain.
*/
//go:nosplit
func (self class) GetFabrikJointUseTargetRotation(joint_idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_use_target_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSkeletonModification2DFABRIK() Expert { return self[0].AsSkeletonModification2DFABRIK() }


//go:nosplit
func (self Simple) AsSkeletonModification2DFABRIK() Simple { return self[0].AsSkeletonModification2DFABRIK() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SkeletonModification2DFABRIK", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
