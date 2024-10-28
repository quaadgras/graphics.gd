package SkeletonModification2DFABRIK

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
This [SkeletonModification2D] uses an algorithm called Forward And Backward Reaching Inverse Kinematics, or FABRIK, to rotate a bone chain so that it reaches a target.
FABRIK works by knowing the positions and lengths of a series of bones, typically called a "bone chain". It first starts by running a forward pass, which places the final bone at the target's position. Then all other bones are moved towards the tip bone, so they stay at the defined bone length away. Then a backwards pass is performed, where the root/first bone in the FABRIK chain is placed back at the origin. Then all other bones are moved so they stay at the defined bone length away. This positions the bone chain so that it reaches the target when possible, but all of the bones stay the correct length away from each other.
Because of how FABRIK works, it often gives more natural results than those seen in [SkeletonModification2DCCDIK]. FABRIK also supports angle constraints, which are fully taken into account when solving.
[b]Note:[/b] The FABRIK modifier has [code]fabrik_joints[/code], which are the data objects that hold the data for each joint in the FABRIK chain. This is different from [Bone2D] nodes! FABRIK joints hold the data needed for each [Bone2D] in the bone chain used by FABRIK.
To help control how the FABRIK joints move, a magnet vector can be passed, which can nudge the bones in a certain direction prior to solving, giving a level of control over the final result.

*/
type Go [1]classdb.SkeletonModification2DFABRIK

/*
Sets the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
func (self Go) SetFabrikJointBone2dNode(joint_idx int, bone2d_nodepath string) {
	class(self).SetFabrikJointBone2dNode(gd.Int(joint_idx), gd.NewString(bone2d_nodepath).NodePath())
}

/*
Returns the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
func (self Go) GetFabrikJointBone2dNode(joint_idx int) string {
	return string(class(self).GetFabrikJointBone2dNode(gd.Int(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the FABRIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the FABRIK joint based on data provided by the linked skeleton.
*/
func (self Go) SetFabrikJointBoneIndex(joint_idx int, bone_idx int) {
	class(self).SetFabrikJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
func (self Go) GetFabrikJointBoneIndex(joint_idx int) int {
	return int(int(class(self).GetFabrikJointBoneIndex(gd.Int(joint_idx))))
}

/*
Sets the magnet position vector for the joint at [param joint_idx].
*/
func (self Go) SetFabrikJointMagnetPosition(joint_idx int, magnet_position gd.Vector2) {
	class(self).SetFabrikJointMagnetPosition(gd.Int(joint_idx), magnet_position)
}

/*
Returns the magnet position vector for the joint at [param joint_idx].
*/
func (self Go) GetFabrikJointMagnetPosition(joint_idx int) gd.Vector2 {
	return gd.Vector2(class(self).GetFabrikJointMagnetPosition(gd.Int(joint_idx)))
}

/*
Sets whether the joint at [param joint_idx] will use the target node's rotation rather than letting FABRIK rotate the node.
[b]Note:[/b] This option only works for the tip/final joint in the chain. For all other nodes, this option will be ignored.
*/
func (self Go) SetFabrikJointUseTargetRotation(joint_idx int, use_target_rotation bool) {
	class(self).SetFabrikJointUseTargetRotation(gd.Int(joint_idx), use_target_rotation)
}

/*
Returns whether the joint is using the target's rotation rather than allowing FABRIK to rotate the joint. This option only applies to the tip/final joint in the chain.
*/
func (self Go) GetFabrikJointUseTargetRotation(joint_idx int) bool {
	return bool(class(self).GetFabrikJointUseTargetRotation(gd.Int(joint_idx)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModification2DFABRIK
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DFABRIK"))
	return Go{classdb.SkeletonModification2DFABRIK(object)}
}

func (self Go) TargetNodepath() string {
		return string(class(self).GetTargetNode().String())
}

func (self Go) SetTargetNodepath(value string) {
	class(self).SetTargetNode(gd.NewString(value).NodePath())
}

func (self Go) FabrikDataChainLength() int {
		return int(int(class(self).GetFabrikDataChainLength()))
}

func (self Go) SetFabrikDataChainLength(value int) {
	class(self).SetFabrikDataChainLength(gd.Int(value))
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(target_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFabrikDataChainLength(length gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFabrikDataChainLength() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetFabrikJointBone2dNode(joint_idx gd.Int, bone2d_nodepath gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, discreet.Get(bone2d_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointBone2dNode(joint_idx gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone index, [param bone_idx], of the FABRIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the FABRIK joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetFabrikJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointBoneIndex(joint_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the magnet position vector for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetFabrikJointMagnetPosition(joint_idx gd.Int, magnet_position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, magnet_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_magnet_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the magnet position vector for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointMagnetPosition(joint_idx gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_magnet_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_target_rotation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_use_target_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the joint is using the target's rotation rather than allowing FABRIK to rotate the joint. This option only applies to the tip/final joint in the chain.
*/
//go:nosplit
func (self class) GetFabrikJointUseTargetRotation(joint_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_use_target_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DFABRIK() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModification2DFABRIK() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SkeletonModification2DFABRIK", func(ptr gd.Object) any { return classdb.SkeletonModification2DFABRIK(ptr) })}
