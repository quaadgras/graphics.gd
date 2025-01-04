// Package XRHandTracker provides methods for working with XRHandTracker object instances.
package XRHandTracker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/XRPositionalTracker"
import "graphics.gd/classdb/XRTracker"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A hand tracking system will create an instance of this object and add it to the [XRServer]. This tracking system will then obtain skeleton data, convert it to the Godot Humanoid hand skeleton and store this data on the [XRHandTracker] object.
Use [XRHandModifier3D] to animate a hand mesh using hand tracking data.
*/
type Instance [1]gdclass.XRHandTracker
type Any interface {
	gd.IsClass
	AsXRHandTracker() Instance
}

/*
Sets flags about the validity of the tracking data for the given hand joint.
*/
func (self Instance) SetHandJointFlags(joint gdclass.XRHandTrackerHandJoint, flags gdclass.XRHandTrackerHandJointFlags) {
	class(self).SetHandJointFlags(joint, flags)
}

/*
Returns flags about the validity of the tracking data for the given hand joint (see [enum XRHandTracker.HandJointFlags]).
*/
func (self Instance) GetHandJointFlags(joint gdclass.XRHandTrackerHandJoint) gdclass.XRHandTrackerHandJointFlags {
	return gdclass.XRHandTrackerHandJointFlags(class(self).GetHandJointFlags(joint))
}

/*
Sets the transform for the given hand joint.
*/
func (self Instance) SetHandJointTransform(joint gdclass.XRHandTrackerHandJoint, transform Transform3D.BasisOrigin) {
	class(self).SetHandJointTransform(joint, gd.Transform3D(transform))
}

/*
Returns the transform for the given hand joint.
*/
func (self Instance) GetHandJointTransform(joint gdclass.XRHandTrackerHandJoint) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetHandJointTransform(joint))
}

/*
Sets the radius of the given hand joint.
*/
func (self Instance) SetHandJointRadius(joint gdclass.XRHandTrackerHandJoint, radius Float.X) {
	class(self).SetHandJointRadius(joint, gd.Float(radius))
}

/*
Returns the radius of the given hand joint.
*/
func (self Instance) GetHandJointRadius(joint gdclass.XRHandTrackerHandJoint) Float.X {
	return Float.X(Float.X(class(self).GetHandJointRadius(joint)))
}

/*
Sets the linear velocity for the given hand joint.
*/
func (self Instance) SetHandJointLinearVelocity(joint gdclass.XRHandTrackerHandJoint, linear_velocity Vector3.XYZ) {
	class(self).SetHandJointLinearVelocity(joint, gd.Vector3(linear_velocity))
}

/*
Returns the linear velocity for the given hand joint.
*/
func (self Instance) GetHandJointLinearVelocity(joint gdclass.XRHandTrackerHandJoint) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetHandJointLinearVelocity(joint))
}

/*
Sets the angular velocity for the given hand joint.
*/
func (self Instance) SetHandJointAngularVelocity(joint gdclass.XRHandTrackerHandJoint, angular_velocity Vector3.XYZ) {
	class(self).SetHandJointAngularVelocity(joint, gd.Vector3(angular_velocity))
}

/*
Returns the angular velocity for the given hand joint.
*/
func (self Instance) GetHandJointAngularVelocity(joint gdclass.XRHandTrackerHandJoint) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetHandJointAngularVelocity(joint))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRHandTracker

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRHandTracker"))
	return Instance{*(*gdclass.XRHandTracker)(unsafe.Pointer(&object))}
}

func (self Instance) HasTrackingData() bool {
	return bool(class(self).GetHasTrackingData())
}

func (self Instance) SetHasTrackingData(value bool) {
	class(self).SetHasTrackingData(value)
}

func (self Instance) HandTrackingSource() gdclass.XRHandTrackerHandTrackingSource {
	return gdclass.XRHandTrackerHandTrackingSource(class(self).GetHandTrackingSource())
}

func (self Instance) SetHandTrackingSource(value gdclass.XRHandTrackerHandTrackingSource) {
	class(self).SetHandTrackingSource(value)
}

//go:nosplit
func (self class) SetHasTrackingData(has_data bool) {
	var frame = callframe.New()
	callframe.Arg(frame, has_data)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHasTrackingData() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandTrackingSource(source gdclass.XRHandTrackerHandTrackingSource) {
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHandTrackingSource() gdclass.XRHandTrackerHandTrackingSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRHandTrackerHandTrackingSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets flags about the validity of the tracking data for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointFlags(joint gdclass.XRHandTrackerHandJoint, flags gdclass.XRHandTrackerHandJointFlags) {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns flags about the validity of the tracking data for the given hand joint (see [enum XRHandTracker.HandJointFlags]).
*/
//go:nosplit
func (self class) GetHandJointFlags(joint gdclass.XRHandTrackerHandJoint) gdclass.XRHandTrackerHandJointFlags {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gdclass.XRHandTrackerHandJointFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the transform for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointTransform(joint gdclass.XRHandTrackerHandJoint, transform gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_hand_joint_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the transform for the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointTransform(joint gdclass.XRHandTrackerHandJoint) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_hand_joint_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the radius of the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointRadius(joint gdclass.XRHandTrackerHandJoint, radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the radius of the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointRadius(joint gdclass.XRHandTrackerHandJoint) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the linear velocity for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointLinearVelocity(joint gdclass.XRHandTrackerHandJoint, linear_velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the linear velocity for the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointLinearVelocity(joint gdclass.XRHandTrackerHandJoint) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the angular velocity for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointAngularVelocity(joint gdclass.XRHandTrackerHandJoint, angular_velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_set_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the angular velocity for the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointAngularVelocity(joint gdclass.XRHandTrackerHandJoint) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRHandTracker.Bind_get_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRHandTracker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRHandTracker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}
func init() {
	gdclass.Register("XRHandTracker", func(ptr gd.Object) any {
		return [1]gdclass.XRHandTracker{*(*gdclass.XRHandTracker)(unsafe.Pointer(&ptr))}
	})
}

type HandTrackingSource = gdclass.XRHandTrackerHandTrackingSource

const (
	/*The source of hand tracking data is unknown.*/
	HandTrackingSourceUnknown HandTrackingSource = 0
	/*The source of hand tracking data is unobstructed, meaning that an accurate method of hand tracking is used. These include optical hand tracking, data gloves, etc.*/
	HandTrackingSourceUnobstructed HandTrackingSource = 1
	/*The source of hand tracking data is a controller, meaning that joint positions are inferred from controller inputs.*/
	HandTrackingSourceController HandTrackingSource = 2
	/*Represents the size of the [enum HandTrackingSource] enum.*/
	HandTrackingSourceMax HandTrackingSource = 3
)

type HandJoint = gdclass.XRHandTrackerHandJoint

const (
	/*Palm joint.*/
	HandJointPalm HandJoint = 0
	/*Wrist joint.*/
	HandJointWrist HandJoint = 1
	/*Thumb metacarpal joint.*/
	HandJointThumbMetacarpal HandJoint = 2
	/*Thumb phalanx proximal joint.*/
	HandJointThumbPhalanxProximal HandJoint = 3
	/*Thumb phalanx distal joint.*/
	HandJointThumbPhalanxDistal HandJoint = 4
	/*Thumb tip joint.*/
	HandJointThumbTip HandJoint = 5
	/*Index finger metacarpal joint.*/
	HandJointIndexFingerMetacarpal HandJoint = 6
	/*Index finger phalanx proximal joint.*/
	HandJointIndexFingerPhalanxProximal HandJoint = 7
	/*Index finger phalanx intermediate joint.*/
	HandJointIndexFingerPhalanxIntermediate HandJoint = 8
	/*Index finger phalanx distal joint.*/
	HandJointIndexFingerPhalanxDistal HandJoint = 9
	/*Index finger tip joint.*/
	HandJointIndexFingerTip HandJoint = 10
	/*Middle finger metacarpal joint.*/
	HandJointMiddleFingerMetacarpal HandJoint = 11
	/*Middle finger phalanx proximal joint.*/
	HandJointMiddleFingerPhalanxProximal HandJoint = 12
	/*Middle finger phalanx intermediate joint.*/
	HandJointMiddleFingerPhalanxIntermediate HandJoint = 13
	/*Middle finger phalanx distal joint.*/
	HandJointMiddleFingerPhalanxDistal HandJoint = 14
	/*Middle finger tip joint.*/
	HandJointMiddleFingerTip HandJoint = 15
	/*Ring finger metacarpal joint.*/
	HandJointRingFingerMetacarpal HandJoint = 16
	/*Ring finger phalanx proximal joint.*/
	HandJointRingFingerPhalanxProximal HandJoint = 17
	/*Ring finger phalanx intermediate joint.*/
	HandJointRingFingerPhalanxIntermediate HandJoint = 18
	/*Ring finger phalanx distal joint.*/
	HandJointRingFingerPhalanxDistal HandJoint = 19
	/*Ring finger tip joint.*/
	HandJointRingFingerTip HandJoint = 20
	/*Pinky finger metacarpal joint.*/
	HandJointPinkyFingerMetacarpal HandJoint = 21
	/*Pinky finger phalanx proximal joint.*/
	HandJointPinkyFingerPhalanxProximal HandJoint = 22
	/*Pinky finger phalanx intermediate joint.*/
	HandJointPinkyFingerPhalanxIntermediate HandJoint = 23
	/*Pinky finger phalanx distal joint.*/
	HandJointPinkyFingerPhalanxDistal HandJoint = 24
	/*Pinky finger tip joint.*/
	HandJointPinkyFingerTip HandJoint = 25
	/*Represents the size of the [enum HandJoint] enum.*/
	HandJointMax HandJoint = 26
)

type HandJointFlags = gdclass.XRHandTrackerHandJointFlags

const (
	/*The hand joint's orientation data is valid.*/
	HandJointFlagOrientationValid HandJointFlags = 1
	/*The hand joint's orientation is actively tracked. May not be set if tracking has been temporarily lost.*/
	HandJointFlagOrientationTracked HandJointFlags = 2
	/*The hand joint's position data is valid.*/
	HandJointFlagPositionValid HandJointFlags = 4
	/*The hand joint's position is actively tracked. May not be set if tracking has been temporarily lost.*/
	HandJointFlagPositionTracked HandJointFlags = 8
	/*The hand joint's linear velocity data is valid.*/
	HandJointFlagLinearVelocityValid HandJointFlags = 16
	/*The hand joint's angular velocity data is valid.*/
	HandJointFlagAngularVelocityValid HandJointFlags = 32
)
