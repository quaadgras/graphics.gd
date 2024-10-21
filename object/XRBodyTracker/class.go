package XRBodyTracker

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/XRPositionalTracker"
import "grow.graphics/gd/object/XRTracker"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A body tracking system will create an instance of this object and add it to the [XRServer]. This tracking system will then obtain skeleton data, convert it to the Godot Humanoid skeleton and store this data on the [XRBodyTracker] object.
Use [XRBodyModifier3D] to animate a body mesh using body tracking data.

*/
type Simple [1]classdb.XRBodyTracker
func (self Simple) SetHasTrackingData(has_data bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHasTrackingData(has_data)
}
func (self Simple) GetHasTrackingData() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetHasTrackingData())
}
func (self Simple) SetBodyFlags(flags classdb.XRBodyTrackerBodyFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBodyFlags(flags)
}
func (self Simple) GetBodyFlags() classdb.XRBodyTrackerBodyFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRBodyTrackerBodyFlags(Expert(self).GetBodyFlags())
}
func (self Simple) SetJointFlags(joint classdb.XRBodyTrackerJoint, flags classdb.XRBodyTrackerJointFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointFlags(joint, flags)
}
func (self Simple) GetJointFlags(joint classdb.XRBodyTrackerJoint) classdb.XRBodyTrackerJointFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRBodyTrackerJointFlags(Expert(self).GetJointFlags(joint))
}
func (self Simple) SetJointTransform(joint classdb.XRBodyTrackerJoint, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetJointTransform(joint, transform)
}
func (self Simple) GetJointTransform(joint classdb.XRBodyTrackerJoint) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).GetJointTransform(joint))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.XRBodyTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetHasTrackingData(has_data bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, has_data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_set_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHasTrackingData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBodyFlags(flags classdb.XRBodyTrackerBodyFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_set_body_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBodyFlags() classdb.XRBodyTrackerBodyFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRBodyTrackerBodyFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_get_body_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets flags about the validity of the tracking data for the given body joint.
*/
//go:nosplit
func (self class) SetJointFlags(joint classdb.XRBodyTrackerJoint, flags classdb.XRBodyTrackerJointFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_set_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns flags about the validity of the tracking data for the given body joint (see [enum XRBodyTracker.JointFlags]).
*/
//go:nosplit
func (self class) GetJointFlags(joint classdb.XRBodyTrackerJoint) classdb.XRBodyTrackerJointFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[classdb.XRBodyTrackerJointFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_get_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the transform for the given body joint.
*/
//go:nosplit
func (self class) SetJointTransform(joint classdb.XRBodyTrackerJoint, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_set_joint_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the transform for the given body joint.
*/
//go:nosplit
func (self class) GetJointTransform(joint classdb.XRBodyTrackerJoint) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRBodyTracker.Bind_get_joint_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsXRBodyTracker() Expert { return self[0].AsXRBodyTracker() }


//go:nosplit
func (self Simple) AsXRBodyTracker() Simple { return self[0].AsXRBodyTracker() }


//go:nosplit
func (self class) AsXRPositionalTracker() XRPositionalTracker.Expert { return self[0].AsXRPositionalTracker() }


//go:nosplit
func (self Simple) AsXRPositionalTracker() XRPositionalTracker.Simple { return self[0].AsXRPositionalTracker() }


//go:nosplit
func (self class) AsXRTracker() XRTracker.Expert { return self[0].AsXRTracker() }


//go:nosplit
func (self Simple) AsXRTracker() XRTracker.Simple { return self[0].AsXRTracker() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("XRBodyTracker", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BodyFlags = classdb.XRBodyTrackerBodyFlags

const (
/*Upper body tracking supported.*/
	BodyFlagUpperBodySupported BodyFlags = 1
/*Lower body tracking supported.*/
	BodyFlagLowerBodySupported BodyFlags = 2
/*Hand tracking supported.*/
	BodyFlagHandsSupported BodyFlags = 4
)
type Joint = classdb.XRBodyTrackerJoint

const (
/*Root joint.*/
	JointRoot Joint = 0
/*Hips joint.*/
	JointHips Joint = 1
/*Spine joint.*/
	JointSpine Joint = 2
/*Chest joint.*/
	JointChest Joint = 3
/*Upper chest joint.*/
	JointUpperChest Joint = 4
/*Neck joint.*/
	JointNeck Joint = 5
/*Head joint.*/
	JointHead Joint = 6
/*Head tip joint.*/
	JointHeadTip Joint = 7
/*Left shoulder joint.*/
	JointLeftShoulder Joint = 8
/*Left upper arm joint.*/
	JointLeftUpperArm Joint = 9
/*Left lower arm joint.*/
	JointLeftLowerArm Joint = 10
/*Right shoulder joint.*/
	JointRightShoulder Joint = 11
/*Right upper arm joint.*/
	JointRightUpperArm Joint = 12
/*Right lower arm joint.*/
	JointRightLowerArm Joint = 13
/*Left upper leg joint.*/
	JointLeftUpperLeg Joint = 14
/*Left lower leg joint.*/
	JointLeftLowerLeg Joint = 15
/*Left foot joint.*/
	JointLeftFoot Joint = 16
/*Left toes joint.*/
	JointLeftToes Joint = 17
/*Right upper leg joint.*/
	JointRightUpperLeg Joint = 18
/*Right lower leg joint.*/
	JointRightLowerLeg Joint = 19
/*Right foot joint.*/
	JointRightFoot Joint = 20
/*Right toes joint.*/
	JointRightToes Joint = 21
/*Left hand joint.*/
	JointLeftHand Joint = 22
/*Left palm joint.*/
	JointLeftPalm Joint = 23
/*Left wrist joint.*/
	JointLeftWrist Joint = 24
/*Left thumb metacarpal joint.*/
	JointLeftThumbMetacarpal Joint = 25
/*Left thumb phalanx proximal joint.*/
	JointLeftThumbPhalanxProximal Joint = 26
/*Left thumb phalanx distal joint.*/
	JointLeftThumbPhalanxDistal Joint = 27
/*Left thumb tip joint.*/
	JointLeftThumbTip Joint = 28
/*Left index finger metacarpal joint.*/
	JointLeftIndexFingerMetacarpal Joint = 29
/*Left index finger phalanx proximal joint.*/
	JointLeftIndexFingerPhalanxProximal Joint = 30
/*Left index finger phalanx intermediate joint.*/
	JointLeftIndexFingerPhalanxIntermediate Joint = 31
/*Left index finger phalanx distal joint.*/
	JointLeftIndexFingerPhalanxDistal Joint = 32
/*Left index finger tip joint.*/
	JointLeftIndexFingerTip Joint = 33
/*Left middle finger metacarpal joint.*/
	JointLeftMiddleFingerMetacarpal Joint = 34
/*Left middle finger phalanx proximal joint.*/
	JointLeftMiddleFingerPhalanxProximal Joint = 35
/*Left middle finger phalanx intermediate joint.*/
	JointLeftMiddleFingerPhalanxIntermediate Joint = 36
/*Left middle finger phalanx distal joint.*/
	JointLeftMiddleFingerPhalanxDistal Joint = 37
/*Left middle finger tip joint.*/
	JointLeftMiddleFingerTip Joint = 38
/*Left ring finger metacarpal joint.*/
	JointLeftRingFingerMetacarpal Joint = 39
/*Left ring finger phalanx proximal joint.*/
	JointLeftRingFingerPhalanxProximal Joint = 40
/*Left ring finger phalanx intermediate joint.*/
	JointLeftRingFingerPhalanxIntermediate Joint = 41
/*Left ring finger phalanx distal joint.*/
	JointLeftRingFingerPhalanxDistal Joint = 42
/*Left ring finger tip joint.*/
	JointLeftRingFingerTip Joint = 43
/*Left pinky finger metacarpal joint.*/
	JointLeftPinkyFingerMetacarpal Joint = 44
/*Left pinky finger phalanx proximal joint.*/
	JointLeftPinkyFingerPhalanxProximal Joint = 45
/*Left pinky finger phalanx intermediate joint.*/
	JointLeftPinkyFingerPhalanxIntermediate Joint = 46
/*Left pinky finger phalanx distal joint.*/
	JointLeftPinkyFingerPhalanxDistal Joint = 47
/*Left pinky finger tip joint.*/
	JointLeftPinkyFingerTip Joint = 48
/*Right hand joint.*/
	JointRightHand Joint = 49
/*Right palm joint.*/
	JointRightPalm Joint = 50
/*Right wrist joint.*/
	JointRightWrist Joint = 51
/*Right thumb metacarpal joint.*/
	JointRightThumbMetacarpal Joint = 52
/*Right thumb phalanx proximal joint.*/
	JointRightThumbPhalanxProximal Joint = 53
/*Right thumb phalanx distal joint.*/
	JointRightThumbPhalanxDistal Joint = 54
/*Right thumb tip joint.*/
	JointRightThumbTip Joint = 55
/*Right index finger metacarpal joint.*/
	JointRightIndexFingerMetacarpal Joint = 56
/*Right index finger phalanx proximal joint.*/
	JointRightIndexFingerPhalanxProximal Joint = 57
/*Right index finger phalanx intermediate joint.*/
	JointRightIndexFingerPhalanxIntermediate Joint = 58
/*Right index finger phalanx distal joint.*/
	JointRightIndexFingerPhalanxDistal Joint = 59
/*Right index finger tip joint.*/
	JointRightIndexFingerTip Joint = 60
/*Right middle finger metacarpal joint.*/
	JointRightMiddleFingerMetacarpal Joint = 61
/*Right middle finger phalanx proximal joint.*/
	JointRightMiddleFingerPhalanxProximal Joint = 62
/*Right middle finger phalanx intermediate joint.*/
	JointRightMiddleFingerPhalanxIntermediate Joint = 63
/*Right middle finger phalanx distal joint.*/
	JointRightMiddleFingerPhalanxDistal Joint = 64
/*Right middle finger tip joint.*/
	JointRightMiddleFingerTip Joint = 65
/*Right ring finger metacarpal joint.*/
	JointRightRingFingerMetacarpal Joint = 66
/*Right ring finger phalanx proximal joint.*/
	JointRightRingFingerPhalanxProximal Joint = 67
/*Right ring finger phalanx intermediate joint.*/
	JointRightRingFingerPhalanxIntermediate Joint = 68
/*Right ring finger phalanx distal joint.*/
	JointRightRingFingerPhalanxDistal Joint = 69
/*Right ring finger tip joint.*/
	JointRightRingFingerTip Joint = 70
/*Right pinky finger metacarpal joint.*/
	JointRightPinkyFingerMetacarpal Joint = 71
/*Right pinky finger phalanx proximal joint.*/
	JointRightPinkyFingerPhalanxProximal Joint = 72
/*Right pinky finger phalanx intermediate joint.*/
	JointRightPinkyFingerPhalanxIntermediate Joint = 73
/*Right pinky finger phalanx distal joint.*/
	JointRightPinkyFingerPhalanxDistal Joint = 74
/*Right pinky finger tip joint.*/
	JointRightPinkyFingerTip Joint = 75
/*Represents the size of the [enum Joint] enum.*/
	JointMax Joint = 76
)
type JointFlags = classdb.XRBodyTrackerJointFlags

const (
/*The joint's orientation data is valid.*/
	JointFlagOrientationValid JointFlags = 1
/*The joint's orientation is actively tracked. May not be set if tracking has been temporarily lost.*/
	JointFlagOrientationTracked JointFlags = 2
/*The joint's position data is valid.*/
	JointFlagPositionValid JointFlags = 4
/*The joint's position is actively tracked. May not be set if tracking has been temporarily lost.*/
	JointFlagPositionTracked JointFlags = 8
)
