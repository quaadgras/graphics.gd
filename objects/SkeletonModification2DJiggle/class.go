package SkeletonModification2DJiggle

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/SkeletonModification2D"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This modification moves a series of bones, typically called a bone chain, towards a target. What makes this modification special is that it calculates the velocity and acceleration for each bone in the bone chain, and runs a very light physics-like calculation using the inputted values. This allows the bones to overshoot the target and "jiggle" around. It can be configured to act more like a spring, or sway around like cloth might.
This modification is useful for adding additional motion to things like hair, the edges of clothing, and more. It has several settings to that allow control over how the joint moves when the target moves.
[b]Note:[/b] The Jiggle modifier has [code]jiggle_joints[/code], which are the data objects that hold the data for each joint in the Jiggle chain. This is different from than [Bone2D] nodes! Jiggle joints hold the data needed for each [Bone2D] in the bone chain used by the Jiggle modification.
*/
type Instance [1]classdb.SkeletonModification2DJiggle
type Any interface {
	gd.IsClass
	AsSkeletonModification2DJiggle() Instance
}

/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
func (self Instance) SetUseColliders(use_colliders bool) {
	class(self).SetUseColliders(use_colliders)
}

/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
func (self Instance) GetUseColliders() bool {
	return bool(class(self).GetUseColliders())
}

/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
func (self Instance) SetCollisionMask(collision_mask int) {
	class(self).SetCollisionMask(gd.Int(collision_mask))
}

/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
func (self Instance) GetCollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointBone2dNode(joint_idx int, bone2d_node Path.String) {
	class(self).SetJiggleJointBone2dNode(gd.Int(joint_idx), gd.NewString(string(bone2d_node)).NodePath())
}

/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointBone2dNode(joint_idx int) Path.String {
	return Path.String(class(self).GetJiggleJointBone2dNode(gd.Int(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
func (self Instance) SetJiggleJointBoneIndex(joint_idx int, bone_idx int) {
	class(self).SetJiggleJointBoneIndex(gd.Int(joint_idx), gd.Int(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointBoneIndex(joint_idx int) int {
	return int(int(class(self).GetJiggleJointBoneIndex(gd.Int(joint_idx))))
}

/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
func (self Instance) SetJiggleJointOverride(joint_idx int, override bool) {
	class(self).SetJiggleJointOverride(gd.Int(joint_idx), override)
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
func (self Instance) GetJiggleJointOverride(joint_idx int) bool {
	return bool(class(self).GetJiggleJointOverride(gd.Int(joint_idx)))
}

/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointStiffness(joint_idx int, stiffness Float.X) {
	class(self).SetJiggleJointStiffness(gd.Int(joint_idx), gd.Float(stiffness))
}

/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointStiffness(joint_idx int) Float.X {
	return Float.X(Float.X(class(self).GetJiggleJointStiffness(gd.Int(joint_idx))))
}

/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointMass(joint_idx int, mass Float.X) {
	class(self).SetJiggleJointMass(gd.Int(joint_idx), gd.Float(mass))
}

/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointMass(joint_idx int) Float.X {
	return Float.X(Float.X(class(self).GetJiggleJointMass(gd.Int(joint_idx))))
}

/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointDamping(joint_idx int, damping Float.X) {
	class(self).SetJiggleJointDamping(gd.Int(joint_idx), gd.Float(damping))
}

/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointDamping(joint_idx int) Float.X {
	return Float.X(Float.X(class(self).GetJiggleJointDamping(gd.Int(joint_idx))))
}

/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
func (self Instance) SetJiggleJointUseGravity(joint_idx int, use_gravity bool) {
	class(self).SetJiggleJointUseGravity(gd.Int(joint_idx), use_gravity)
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
func (self Instance) GetJiggleJointUseGravity(joint_idx int) bool {
	return bool(class(self).GetJiggleJointUseGravity(gd.Int(joint_idx)))
}

/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointGravity(joint_idx int, gravity Vector2.XY) {
	class(self).SetJiggleJointGravity(gd.Int(joint_idx), gd.Vector2(gravity))
}

/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
func (self Instance) GetJiggleJointGravity(joint_idx int) Vector2.XY {
	return Vector2.XY(class(self).GetJiggleJointGravity(gd.Int(joint_idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SkeletonModification2DJiggle

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DJiggle"))
	return Instance{classdb.SkeletonModification2DJiggle(object)}
}

func (self Instance) TargetNodepath() Path.String {
	return Path.String(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNodepath(value Path.String) {
	class(self).SetTargetNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) JiggleDataChainLength() int {
	return int(int(class(self).GetJiggleDataChainLength()))
}

func (self Instance) SetJiggleDataChainLength(value int) {
	class(self).SetJiggleDataChainLength(gd.Int(value))
}

func (self Instance) Stiffness() Float.X {
	return Float.X(Float.X(class(self).GetStiffness()))
}

func (self Instance) SetStiffness(value Float.X) {
	class(self).SetStiffness(gd.Float(value))
}

func (self Instance) Mass() Float.X {
	return Float.X(Float.X(class(self).GetMass()))
}

func (self Instance) SetMass(value Float.X) {
	class(self).SetMass(gd.Float(value))
}

func (self Instance) Damping() Float.X {
	return Float.X(Float.X(class(self).GetDamping()))
}

func (self Instance) SetDamping(value Float.X) {
	class(self).SetDamping(gd.Float(value))
}

func (self Instance) UseGravity() bool {
	return bool(class(self).GetUseGravity())
}

func (self Instance) SetUseGravity(value bool) {
	class(self).SetUseGravity(value)
}

func (self Instance) Gravity() Vector2.XY {
	return Vector2.XY(class(self).GetGravity())
}

func (self Instance) SetGravity(value Vector2.XY) {
	class(self).SetGravity(gd.Vector2(value))
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(target_nodepath))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJiggleDataChainLength(length gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJiggleDataChainLength() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStiffness(stiffness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStiffness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMass(mass gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDamping(damping gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDamping() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseGravity(use_gravity bool) {
	var frame = callframe.New()
	callframe.Arg(frame, use_gravity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseGravity() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(gravity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
//go:nosplit
func (self class) SetUseColliders(use_colliders bool) {
	var frame = callframe.New()
	callframe.Arg(frame, use_colliders)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_use_colliders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
//go:nosplit
func (self class) GetUseColliders() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_use_colliders, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointBone2dNode(joint_idx gd.Int, bone2d_node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, pointers.Get(bone2d_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBone2dNode(joint_idx gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetJiggleJointBoneIndex(joint_idx gd.Int, bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBoneIndex(joint_idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
//go:nosplit
func (self class) SetJiggleJointOverride(joint_idx gd.Int, override bool) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, override)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
//go:nosplit
func (self class) GetJiggleJointOverride(joint_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointStiffness(joint_idx gd.Int, stiffness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointStiffness(joint_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointMass(joint_idx gd.Int, mass gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointMass(joint_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointDamping(joint_idx gd.Int, damping gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointDamping(joint_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
//go:nosplit
func (self class) SetJiggleJointUseGravity(joint_idx gd.Int, use_gravity bool) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_gravity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
//go:nosplit
func (self class) GetJiggleJointUseGravity(joint_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointGravity(joint_idx gd.Int, gravity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
//go:nosplit
func (self class) GetJiggleJointGravity(joint_idx gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DJiggle() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DJiggle() Instance {
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsSkeletonModification2D(), name)
	}
}
func init() {
	classdb.Register("SkeletonModification2DJiggle", func(ptr gd.Object) any { return classdb.SkeletonModification2DJiggle(ptr) })
}
