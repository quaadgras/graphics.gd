// Package OpenXRHand provides methods for working with OpenXRHand object instances.
package OpenXRHand

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Path"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node enables OpenXR's hand tracking functionality. The node should be a child node of an [XROrigin3D] node, tracking will update its position to the player's tracked hand Palm joint location (the center of the middle finger's metacarpal bone). This node also updates the skeleton of a properly skinned hand or avatar model.
If the skeleton is a hand (one of the hand bones is the root node of the skeleton), then the skeleton will be placed relative to the hand palm location and the hand mesh and skeleton should be children of the OpenXRHand node.
If the hand bones are part of a full skeleton, then the root of the hand will keep its location with the assumption that IK is used to position the hand and arm.
By default the skeleton hand bones are repositioned to match the size of the tracked hand. To preserve the modeled bone sizes change [member bone_update] to apply rotation only.
*/
type Instance [1]gdclass.OpenXRHand
type Any interface {
	gd.IsClass
	AsOpenXRHand() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRHand

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRHand"))
	return Instance{*(*gdclass.OpenXRHand)(unsafe.Pointer(&object))}
}

func (self Instance) Hand() gdclass.OpenXRHandHands {
	return gdclass.OpenXRHandHands(class(self).GetHand())
}

func (self Instance) SetHand(value gdclass.OpenXRHandHands) {
	class(self).SetHand(value)
}

func (self Instance) MotionRange() gdclass.OpenXRHandMotionRange {
	return gdclass.OpenXRHandMotionRange(class(self).GetMotionRange())
}

func (self Instance) SetMotionRange(value gdclass.OpenXRHandMotionRange) {
	class(self).SetMotionRange(value)
}

func (self Instance) HandSkeleton() Path.String {
	return Path.String(class(self).GetHandSkeleton().String())
}

func (self Instance) SetHandSkeleton(value Path.String) {
	class(self).SetHandSkeleton(gd.NewString(string(value)).NodePath())
}

func (self Instance) SkeletonRig() gdclass.OpenXRHandSkeletonRig {
	return gdclass.OpenXRHandSkeletonRig(class(self).GetSkeletonRig())
}

func (self Instance) SetSkeletonRig(value gdclass.OpenXRHandSkeletonRig) {
	class(self).SetSkeletonRig(value)
}

func (self Instance) BoneUpdate() gdclass.OpenXRHandBoneUpdate {
	return gdclass.OpenXRHandBoneUpdate(class(self).GetBoneUpdate())
}

func (self Instance) SetBoneUpdate(value gdclass.OpenXRHandBoneUpdate) {
	class(self).SetBoneUpdate(value)
}

//go:nosplit
func (self class) SetHand(hand gdclass.OpenXRHandHands) {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_set_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHand() gdclass.OpenXRHandHands {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OpenXRHandHands](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_get_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandSkeleton(hand_skeleton gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(hand_skeleton))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_set_hand_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandSkeleton() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_get_hand_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMotionRange(motion_range gdclass.OpenXRHandMotionRange) {
	var frame = callframe.New()
	callframe.Arg(frame, motion_range)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_set_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotionRange() gdclass.OpenXRHandMotionRange {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OpenXRHandMotionRange](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_get_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeletonRig(skeleton_rig gdclass.OpenXRHandSkeletonRig) {
	var frame = callframe.New()
	callframe.Arg(frame, skeleton_rig)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_set_skeleton_rig, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeletonRig() gdclass.OpenXRHandSkeletonRig {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OpenXRHandSkeletonRig](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_get_skeleton_rig, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBoneUpdate(bone_update gdclass.OpenXRHandBoneUpdate) {
	var frame = callframe.New()
	callframe.Arg(frame, bone_update)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_set_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneUpdate() gdclass.OpenXRHandBoneUpdate {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OpenXRHandBoneUpdate](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRHand.Bind_get_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRHand() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRHand() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {
	gdclass.Register("OpenXRHand", func(ptr gd.Object) any { return [1]gdclass.OpenXRHand{*(*gdclass.OpenXRHand)(unsafe.Pointer(&ptr))} })
}

type Hands = gdclass.OpenXRHandHands

const (
	/*Tracking the player's left hand.*/
	HandLeft Hands = 0
	/*Tracking the player's right hand.*/
	HandRight Hands = 1
	/*Maximum supported hands.*/
	HandMax Hands = 2
)

type MotionRange = gdclass.OpenXRHandMotionRange

const (
	/*When player grips, hand skeleton will form a full fist.*/
	MotionRangeUnobstructed MotionRange = 0
	/*When player grips, hand skeleton conforms to the controller the player is holding.*/
	MotionRangeConformToController MotionRange = 1
	/*Maximum supported motion ranges.*/
	MotionRangeMax MotionRange = 2
)

type SkeletonRig = gdclass.OpenXRHandSkeletonRig

const (
	/*An OpenXR compliant skeleton.*/
	SkeletonRigOpenxr SkeletonRig = 0
	/*A [SkeletonProfileHumanoid] compliant skeleton.*/
	SkeletonRigHumanoid SkeletonRig = 1
	/*Maximum supported hands.*/
	SkeletonRigMax SkeletonRig = 2
)

type BoneUpdate = gdclass.OpenXRHandBoneUpdate

const (
	/*The skeletons bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
	/*The skeletons bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
	/*Maximum supported bone update mode.*/
	BoneUpdateMax BoneUpdate = 2
)
