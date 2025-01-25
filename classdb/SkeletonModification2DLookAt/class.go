// Package SkeletonModification2DLookAt provides methods for working with SkeletonModification2DLookAt object instances.
package SkeletonModification2DLookAt

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/classdb/SkeletonModification2D"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
This [SkeletonModification2D] rotates a bone to look a target. This is extremely helpful for moving character's head to look at the player, rotating a turret to look at a target, or any other case where you want to make a bone rotate towards something quickly and easily.
*/
type Instance [1]gdclass.SkeletonModification2DLookAt

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModification2DLookAt() Instance
}

/*
Sets the amount of additional rotation that is to be applied after executing the modification. This allows for offsetting the results by the inputted rotation amount.
*/
func (self Instance) SetAdditionalRotation(rotation Float.X) {
	class(self).SetAdditionalRotation(gd.Float(rotation))
}

/*
Returns the amount of additional rotation that is applied after the LookAt modification executes.
*/
func (self Instance) GetAdditionalRotation() Float.X {
	return Float.X(Float.X(class(self).GetAdditionalRotation()))
}

/*
Sets whether this modification will use constraints or not. When [code]true[/code], constraints will be applied when solving the LookAt modification.
*/
func (self Instance) SetEnableConstraint(enable_constraint bool) {
	class(self).SetEnableConstraint(enable_constraint)
}

/*
Returns [code]true[/code] if the LookAt modification is using constraints.
*/
func (self Instance) GetEnableConstraint() bool {
	return bool(class(self).GetEnableConstraint())
}

/*
Sets the constraint's minimum allowed angle.
*/
func (self Instance) SetConstraintAngleMin(angle_min Float.X) {
	class(self).SetConstraintAngleMin(gd.Float(angle_min))
}

/*
Returns the constraint's minimum allowed angle.
*/
func (self Instance) GetConstraintAngleMin() Float.X {
	return Float.X(Float.X(class(self).GetConstraintAngleMin()))
}

/*
Sets the constraint's maximum allowed angle.
*/
func (self Instance) SetConstraintAngleMax(angle_max Float.X) {
	class(self).SetConstraintAngleMax(gd.Float(angle_max))
}

/*
Returns the constraint's maximum allowed angle.
*/
func (self Instance) GetConstraintAngleMax() Float.X {
	return Float.X(Float.X(class(self).GetConstraintAngleMax()))
}

/*
When [code]true[/code], the modification will use an inverted joint constraint.
An inverted joint constraint only constraints the [Bone2D] to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
func (self Instance) SetConstraintAngleInvert(invert bool) {
	class(self).SetConstraintAngleInvert(invert)
}

/*
Returns whether the constraints to this modification are inverted or not.
*/
func (self Instance) GetConstraintAngleInvert() bool {
	return bool(class(self).GetConstraintAngleInvert())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModification2DLookAt

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DLookAt"))
	casted := Instance{*(*gdclass.SkeletonModification2DLookAt)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BoneIndex() int {
	return int(int(class(self).GetBoneIndex()))
}

func (self Instance) SetBoneIndex(value int) {
	class(self).SetBoneIndex(gd.Int(value))
}

func (self Instance) Bone2dNode() NodePath.String {
	return NodePath.String(class(self).GetBone2dNode().String())
}

func (self Instance) SetBone2dNode(value NodePath.String) {
	class(self).SetBone2dNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) TargetNodepath() NodePath.String {
	return NodePath.String(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNodepath(value NodePath.String) {
	class(self).SetTargetNode(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetBone2dNode(bone2d_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bone2d_nodepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBone2dNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBoneIndex(bone_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_bone_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_bone_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(target_nodepath))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the amount of additional rotation that is to be applied after executing the modification. This allows for offsetting the results by the inputted rotation amount.
*/
//go:nosplit
func (self class) SetAdditionalRotation(rotation gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_additional_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the amount of additional rotation that is applied after the LookAt modification executes.
*/
//go:nosplit
func (self class) GetAdditionalRotation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_additional_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether this modification will use constraints or not. When [code]true[/code], constraints will be applied when solving the LookAt modification.
*/
//go:nosplit
func (self class) SetEnableConstraint(enable_constraint bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable_constraint)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the LookAt modification is using constraints.
*/
//go:nosplit
func (self class) GetEnableConstraint() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_enable_constraint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the constraint's minimum allowed angle.
*/
//go:nosplit
func (self class) SetConstraintAngleMin(angle_min gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angle_min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the constraint's minimum allowed angle.
*/
//go:nosplit
func (self class) GetConstraintAngleMin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_constraint_angle_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the constraint's maximum allowed angle.
*/
//go:nosplit
func (self class) SetConstraintAngleMax(angle_max gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angle_max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the constraint's maximum allowed angle.
*/
//go:nosplit
func (self class) GetConstraintAngleMax() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_constraint_angle_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
When [code]true[/code], the modification will use an inverted joint constraint.
An inverted joint constraint only constraints the [Bone2D] to the angles [i]outside of[/i] the inputted minimum and maximum angles. For this reason, it is referred to as an inverted joint constraint, as it constraints the joint to the outside of the inputted values.
*/
//go:nosplit
func (self class) SetConstraintAngleInvert(invert bool) {
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_set_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the constraints to this modification are inverted or not.
*/
//go:nosplit
func (self class) GetConstraintAngleInvert() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DLookAt.Bind_get_constraint_angle_invert, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DLookAt() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DLookAt() Instance {
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
	gdclass.Register("SkeletonModification2DLookAt", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModification2DLookAt{*(*gdclass.SkeletonModification2DLookAt)(unsafe.Pointer(&ptr))}
	})
}
