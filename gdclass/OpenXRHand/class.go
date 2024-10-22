package OpenXRHand

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node enables OpenXR's hand tracking functionality. The node should be a child node of an [XROrigin3D] node, tracking will update its position to the player's tracked hand Palm joint location (the center of the middle finger's metacarpal bone). This node also updates the skeleton of a properly skinned hand or avatar model.
If the skeleton is a hand (one of the hand bones is the root node of the skeleton), then the skeleton will be placed relative to the hand palm location and the hand mesh and skeleton should be children of the OpenXRHand node.
If the hand bones are part of a full skeleton, then the root of the hand will keep its location with the assumption that IK is used to position the hand and arm.
By default the skeleton hand bones are repositioned to match the size of the tracked hand. To preserve the modeled bone sizes change [member bone_update] to apply rotation only.

*/
type Go [1]classdb.OpenXRHand
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRHand
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRHand"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Hand() classdb.OpenXRHandHands {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.OpenXRHandHands(class(self).GetHand())
}

func (self Go) SetHand(value classdb.OpenXRHandHands) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHand(value)
}

func (self Go) MotionRange() classdb.OpenXRHandMotionRange {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.OpenXRHandMotionRange(class(self).GetMotionRange())
}

func (self Go) SetMotionRange(value classdb.OpenXRHandMotionRange) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMotionRange(value)
}

func (self Go) HandSkeleton() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetHandSkeleton(gc).String())
}

func (self Go) SetHandSkeleton(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandSkeleton(gc.String(value).NodePath(gc))
}

func (self Go) SkeletonRig() classdb.OpenXRHandSkeletonRig {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.OpenXRHandSkeletonRig(class(self).GetSkeletonRig())
}

func (self Go) SetSkeletonRig(value classdb.OpenXRHandSkeletonRig) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkeletonRig(value)
}

func (self Go) BoneUpdate() classdb.OpenXRHandBoneUpdate {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.OpenXRHandBoneUpdate(class(self).GetBoneUpdate())
}

func (self Go) SetBoneUpdate(value classdb.OpenXRHandBoneUpdate) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBoneUpdate(value)
}

//go:nosplit
func (self class) SetHand(hand classdb.OpenXRHandHands)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_set_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHand() classdb.OpenXRHandHands {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRHandHands](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_get_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHandSkeleton(hand_skeleton gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(hand_skeleton))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_set_hand_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandSkeleton(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_get_hand_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMotionRange(motion_range classdb.OpenXRHandMotionRange)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motion_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_set_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMotionRange() classdb.OpenXRHandMotionRange {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRHandMotionRange](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_get_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkeletonRig(skeleton_rig classdb.OpenXRHandSkeletonRig)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skeleton_rig)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_set_skeleton_rig, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkeletonRig() classdb.OpenXRHandSkeletonRig {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRHandSkeletonRig](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_get_skeleton_rig, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBoneUpdate(bone_update classdb.OpenXRHandBoneUpdate)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bone_update)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_set_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneUpdate() classdb.OpenXRHandBoneUpdate {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRHandBoneUpdate](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRHand.Bind_get_bone_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRHand() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRHand() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("OpenXRHand", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Hands = classdb.OpenXRHandHands

const (
/*Tracking the player's left hand.*/
	HandLeft Hands = 0
/*Tracking the player's right hand.*/
	HandRight Hands = 1
/*Maximum supported hands.*/
	HandMax Hands = 2
)
type MotionRange = classdb.OpenXRHandMotionRange

const (
/*When player grips, hand skeleton will form a full fist.*/
	MotionRangeUnobstructed MotionRange = 0
/*When player grips, hand skeleton conforms to the controller the player is holding.*/
	MotionRangeConformToController MotionRange = 1
/*Maximum supported motion ranges.*/
	MotionRangeMax MotionRange = 2
)
type SkeletonRig = classdb.OpenXRHandSkeletonRig

const (
/*An OpenXR compliant skeleton.*/
	SkeletonRigOpenxr SkeletonRig = 0
/*A [SkeletonProfileHumanoid] compliant skeleton.*/
	SkeletonRigHumanoid SkeletonRig = 1
/*Maximum supported hands.*/
	SkeletonRigMax SkeletonRig = 2
)
type BoneUpdate = classdb.OpenXRHandBoneUpdate

const (
/*The skeletons bones are fully updated (both position and rotation) to match the tracked bones.*/
	BoneUpdateFull BoneUpdate = 0
/*The skeletons bones are only rotated to align with the tracked bones, preserving bone length.*/
	BoneUpdateRotationOnly BoneUpdate = 1
/*Maximum supported bone update mode.*/
	BoneUpdateMax BoneUpdate = 2
)
