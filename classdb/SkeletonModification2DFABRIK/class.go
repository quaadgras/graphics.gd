// Package SkeletonModification2DFABRIK provides methods for working with SkeletonModification2DFABRIK object instances.
package SkeletonModification2DFABRIK

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
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This [SkeletonModification2D] uses an algorithm called Forward And Backward Reaching Inverse Kinematics, or FABRIK, to rotate a bone chain so that it reaches a target.
FABRIK works by knowing the positions and lengths of a series of bones, typically called a "bone chain". It first starts by running a forward pass, which places the final bone at the target's position. Then all other bones are moved towards the tip bone, so they stay at the defined bone length away. Then a backwards pass is performed, where the root/first bone in the FABRIK chain is placed back at the origin. Then all other bones are moved so they stay at the defined bone length away. This positions the bone chain so that it reaches the target when possible, but all of the bones stay the correct length away from each other.
Because of how FABRIK works, it often gives more natural results than those seen in [SkeletonModification2DCCDIK]. FABRIK also supports angle constraints, which are fully taken into account when solving.
[b]Note:[/b] The FABRIK modifier has [code]fabrik_joints[/code], which are the data objects that hold the data for each joint in the FABRIK chain. This is different from [Bone2D] nodes! FABRIK joints hold the data needed for each [Bone2D] in the bone chain used by FABRIK.
To help control how the FABRIK joints move, a magnet vector can be passed, which can nudge the bones in a certain direction prior to solving, giving a level of control over the final result.
*/
type Instance [1]gdclass.SkeletonModification2DFABRIK

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModification2DFABRIK() Instance
}

/*
Sets the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
func (self Instance) SetFabrikJointBone2dNode(joint_idx int, bone2d_nodepath NodePath.String) {
	class(self).SetFabrikJointBone2dNode(gd.Int(joint_idx), gd.NewString(string(bone2d_nodepath)).NodePath())
}

/*
Returns the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
func (self Instance) GetFabrikJointBone2dNode(joint_idx int) NodePath.String {
	return NodePath.String(class(self).GetFabrikJointBone2dNode(gd.Int(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the FABRIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the FABRIK joint based on data provided by the linked skeleton.
*/
func (self Instance) SetFabrikJointBoneIndex(joint_idx int, bone_idx int) {
	class(self).SetFabrikJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
func (self Instance) GetFabrikJointBoneIndex(joint_idx int) int {
	return int(int(class(self).GetFabrikJointBoneIndex(gd.Int(joint_idx))))
}

/*
Sets the magnet position vector for the joint at [param joint_idx].
*/
func (self Instance) SetFabrikJointMagnetPosition(joint_idx int, magnet_position Vector2.XY) {
	class(self).SetFabrikJointMagnetPosition(gd.Int(joint_idx), gd.Vector2(magnet_position))
}

/*
Returns the magnet position vector for the joint at [param joint_idx].
*/
func (self Instance) GetFabrikJointMagnetPosition(joint_idx int) Vector2.XY {
	return Vector2.XY(class(self).GetFabrikJointMagnetPosition(gd.Int(joint_idx)))
}

/*
Sets whether the joint at [param joint_idx] will use the target node's rotation rather than letting FABRIK rotate the node.
[b]Note:[/b] This option only works for the tip/final joint in the chain. For all other nodes, this option will be ignored.
*/
func (self Instance) SetFabrikJointUseTargetRotation(joint_idx int, use_target_rotation bool) {
	class(self).SetFabrikJointUseTargetRotation(gd.Int(joint_idx), use_target_rotation)
}

/*
Returns whether the joint is using the target's rotation rather than allowing FABRIK to rotate the joint. This option only applies to the tip/final joint in the chain.
*/
func (self Instance) GetFabrikJointUseTargetRotation(joint_idx int) bool {
	return bool(class(self).GetFabrikJointUseTargetRotation(gd.Int(joint_idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModification2DFABRIK

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DFABRIK"))
	casted := Instance{*(*gdclass.SkeletonModification2DFABRIK)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TargetNodepath() NodePath.String {
	return NodePath.String(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNodepath(value NodePath.String) {
	class(self).SetTargetNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) FabrikDataChainLength() int {
	return int(int(class(self).GetFabrikDataChainLength()))
}

func (self Instance) SetFabrikDataChainLength(value int) {
	class(self).SetFabrikDataChainLength(gd.Int(value))
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(target_nodepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFabrikDataChainLength(length gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFabrikDataChainLength() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetFabrikJointBone2dNode(joint_idx gd.Int, bone2d_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, pointers.Get(bone2d_nodepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Bone2D] node assigned to the FABRIK joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetFabrikJointBone2dNode(joint_idx gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the bone index, [param bone_idx], of the FABRIK joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the FABRIK joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetFabrikJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the magnet position vector for the joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetFabrikJointMagnetPosition(joint_idx gd.Int, magnet_position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, magnet_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_magnet_position, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_magnet_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the joint at [param joint_idx] will use the target node's rotation rather than letting FABRIK rotate the node.
[b]Note:[/b] This option only works for the tip/final joint in the chain. For all other nodes, this option will be ignored.
*/
//go:nosplit
func (self class) SetFabrikJointUseTargetRotation(joint_idx gd.Int, use_target_rotation bool) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_target_rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_set_fabrik_joint_use_target_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DFABRIK.Bind_get_fabrik_joint_use_target_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DFABRIK() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DFABRIK() Instance {
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
	gdclass.Register("SkeletonModification2DFABRIK", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModification2DFABRIK{*(*gdclass.SkeletonModification2DFABRIK)(unsafe.Pointer(&ptr))}
	})
}
