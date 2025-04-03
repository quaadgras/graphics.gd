// Package XRBodyTracker provides methods for working with XRBodyTracker object instances.
package XRBodyTracker

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/XRPositionalTracker"
import "graphics.gd/classdb/XRTracker"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A body tracking system will create an instance of this object and add it to the [XRServer]. This tracking system will then obtain skeleton data, convert it to the Godot Humanoid skeleton and store this data on the [XRBodyTracker] object.
Use [XRBodyModifier3D] to animate a body mesh using body tracking data.
*/
type Instance [1]gdclass.XRBodyTracker

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRBodyTracker() Instance
}

/*
Sets flags about the validity of the tracking data for the given body joint.
*/
func (self Instance) SetJointFlags(joint gdclass.XRBodyTrackerJoint, flags gdclass.XRBodyTrackerJointFlags) { //gd:XRBodyTracker.set_joint_flags
	Advanced(self).SetJointFlags(joint, flags)
}

/*
Returns flags about the validity of the tracking data for the given body joint (see [enum XRBodyTracker.JointFlags]).
*/
func (self Instance) GetJointFlags(joint gdclass.XRBodyTrackerJoint) gdclass.XRBodyTrackerJointFlags { //gd:XRBodyTracker.get_joint_flags
	return gdclass.XRBodyTrackerJointFlags(Advanced(self).GetJointFlags(joint))
}

/*
Sets the transform for the given body joint.
*/
func (self Instance) SetJointTransform(joint gdclass.XRBodyTrackerJoint, transform Transform3D.BasisOrigin) { //gd:XRBodyTracker.set_joint_transform
	Advanced(self).SetJointTransform(joint, Transform3D.BasisOrigin(transform))
}

/*
Returns the transform for the given body joint.
*/
func (self Instance) GetJointTransform(joint gdclass.XRBodyTrackerJoint) Transform3D.BasisOrigin { //gd:XRBodyTracker.get_joint_transform
	return Transform3D.BasisOrigin(Advanced(self).GetJointTransform(joint))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRBodyTracker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRBodyTracker"))
	casted := Instance{*(*gdclass.XRBodyTracker)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) HasTrackingData() bool {
	return bool(class(self).GetHasTrackingData())
}

func (self Instance) SetHasTrackingData(value bool) {
	class(self).SetHasTrackingData(value)
}

func (self Instance) BodyFlags() gdclass.XRBodyTrackerBodyFlags {
	return gdclass.XRBodyTrackerBodyFlags(class(self).GetBodyFlags())
}

func (self Instance) SetBodyFlags(value gdclass.XRBodyTrackerBodyFlags) {
	class(self).SetBodyFlags(value)
}

//go:nosplit
func (self class) SetHasTrackingData(has_data bool) { //gd:XRBodyTracker.set_has_tracking_data
	var frame = callframe.New()
	callframe.Arg(frame, has_data)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_set_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHasTrackingData() bool { //gd:XRBodyTracker.get_has_tracking_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBodyFlags(flags gdclass.XRBodyTrackerBodyFlags) { //gd:XRBodyTracker.set_body_flags
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_set_body_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBodyFlags() gdclass.XRBodyTrackerBodyFlags { //gd:XRBodyTracker.get_body_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRBodyTrackerBodyFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_get_body_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets flags about the validity of the tracking data for the given body joint.
*/
//go:nosplit
func (self class) SetJointFlags(joint gdclass.XRBodyTrackerJoint, flags gdclass.XRBodyTrackerJointFlags) { //gd:XRBodyTracker.set_joint_flags
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_set_joint_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns flags about the validity of the tracking data for the given body joint (see [enum XRBodyTracker.JointFlags]).
*/
//go:nosplit
func (self class) GetJointFlags(joint gdclass.XRBodyTrackerJoint) gdclass.XRBodyTrackerJointFlags { //gd:XRBodyTracker.get_joint_flags
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gdclass.XRBodyTrackerJointFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_get_joint_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the transform for the given body joint.
*/
//go:nosplit
func (self class) SetJointTransform(joint gdclass.XRBodyTrackerJoint, transform Transform3D.BasisOrigin) { //gd:XRBodyTracker.set_joint_transform
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_set_joint_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the transform for the given body joint.
*/
//go:nosplit
func (self class) GetJointTransform(joint gdclass.XRBodyTrackerJoint) Transform3D.BasisOrigin { //gd:XRBodyTracker.get_joint_transform
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRBodyTracker.Bind_get_joint_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRBodyTracker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRBodyTracker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRPositionalTracker() XRPositionalTracker.Advanced {
	return *((*XRPositionalTracker.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRPositionalTracker() XRPositionalTracker.Instance {
	return *((*XRPositionalTracker.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsXRTracker() XRTracker.Advanced {
	return *((*XRTracker.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRTracker() XRTracker.Instance {
	return *((*XRTracker.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(XRPositionalTracker.Advanced(self.AsXRPositionalTracker()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRPositionalTracker.Instance(self.AsXRPositionalTracker()), name)
	}
}
func init() {
	gdclass.Register("XRBodyTracker", func(ptr gd.Object) any {
		return [1]gdclass.XRBodyTracker{*(*gdclass.XRBodyTracker)(unsafe.Pointer(&ptr))}
	})
}

type BodyFlags = gdclass.XRBodyTrackerBodyFlags //gd:XRBodyTracker.BodyFlags

const (
	/*Upper body tracking supported.*/
	BodyFlagUpperBodySupported BodyFlags = 1
	/*Lower body tracking supported.*/
	BodyFlagLowerBodySupported BodyFlags = 2
	/*Hand tracking supported.*/
	BodyFlagHandsSupported BodyFlags = 4
)

type Joint = gdclass.XRBodyTrackerJoint //gd:XRBodyTracker.Joint

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

type JointFlags = gdclass.XRBodyTrackerJointFlags //gd:XRBodyTracker.JointFlags

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
