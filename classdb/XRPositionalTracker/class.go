// Package XRPositionalTracker provides methods for working with XRPositionalTracker object instances.
package XRPositionalTracker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/XRTracker"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
An instance of this object represents a device that is tracked, such as a controller or anchor point. HMDs aren't represented here as they are handled internally.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRNode3D] and [XRAnchor3D] both consume objects of this type and should be used in your project. The positional trackers are just under-the-hood objects that make this all work. These are mostly exposed so that GDExtension-based interfaces can interact with them.
*/
type Instance [1]gdclass.XRPositionalTracker

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRPositionalTracker() Instance
}

/*
Returns [code]true[/code] if the tracker is available and is currently tracking the bound [param name] pose.
*/
func (self Instance) HasPose(name string) bool {
	return bool(class(self).HasPose(gd.NewStringName(name)))
}

/*
Returns the current [XRPose] state object for the bound [param name] pose.
*/
func (self Instance) GetPose(name string) [1]gdclass.XRPose {
	return [1]gdclass.XRPose(class(self).GetPose(gd.NewStringName(name)))
}

/*
Marks this pose as invalid, we don't clear the last reported state but it allows users to decide if trackers need to be hidden if we lose tracking or just remain at their last known position.
*/
func (self Instance) InvalidatePose(name string) {
	class(self).InvalidatePose(gd.NewStringName(name))
}

/*
Sets the transform, linear velocity, angular velocity and tracking confidence for the given pose. This method is called by a [XRInterface] implementation and should not be used directly.
*/
func (self Instance) SetPose(name string, transform Transform3D.BasisOrigin, linear_velocity Vector3.XYZ, angular_velocity Vector3.XYZ, tracking_confidence gdclass.XRPoseTrackingConfidence) {
	class(self).SetPose(gd.NewStringName(name), gd.Transform3D(transform), gd.Vector3(linear_velocity), gd.Vector3(angular_velocity), tracking_confidence)
}

/*
Returns an input for this tracker. It can return a boolean, float or [Vector2] value depending on whether the input is a button, trigger or thumbstick/thumbpad.
*/
func (self Instance) GetInput(name string) any {
	return any(class(self).GetInput(gd.NewStringName(name)).Interface())
}

/*
Changes the value for the given input. This method is called by a [XRInterface] implementation and should not be used directly.
*/
func (self Instance) SetInput(name string, value any) {
	class(self).SetInput(gd.NewStringName(name), gd.NewVariant(value))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRPositionalTracker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRPositionalTracker"))
	casted := Instance{*(*gdclass.XRPositionalTracker)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Profile() string {
	return string(class(self).GetTrackerProfile().String())
}

func (self Instance) SetProfile(value string) {
	class(self).SetTrackerProfile(gd.NewString(value))
}

func (self Instance) Hand() gdclass.XRPositionalTrackerTrackerHand {
	return gdclass.XRPositionalTrackerTrackerHand(class(self).GetTrackerHand())
}

func (self Instance) SetHand(value gdclass.XRPositionalTrackerTrackerHand) {
	class(self).SetTrackerHand(value)
}

//go:nosplit
func (self class) GetTrackerProfile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_tracker_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerProfile(profile gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_tracker_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrackerHand() gdclass.XRPositionalTrackerTrackerHand {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRPositionalTrackerTrackerHand](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerHand(hand gdclass.XRPositionalTrackerTrackerHand) {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tracker is available and is currently tracking the bound [param name] pose.
*/
//go:nosplit
func (self class) HasPose(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_has_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current [XRPose] state object for the bound [param name] pose.
*/
//go:nosplit
func (self class) GetPose(name gd.StringName) [1]gdclass.XRPose {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.XRPose{gd.PointerWithOwnershipTransferredToGo[gdclass.XRPose](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Marks this pose as invalid, we don't clear the last reported state but it allows users to decide if trackers need to be hidden if we lose tracking or just remain at their last known position.
*/
//go:nosplit
func (self class) InvalidatePose(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_invalidate_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the transform, linear velocity, angular velocity and tracking confidence for the given pose. This method is called by a [XRInterface] implementation and should not be used directly.
*/
//go:nosplit
func (self class) SetPose(name gd.StringName, transform gd.Transform3D, linear_velocity gd.Vector3, angular_velocity gd.Vector3, tracking_confidence gdclass.XRPoseTrackingConfidence) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, transform)
	callframe.Arg(frame, linear_velocity)
	callframe.Arg(frame, angular_velocity)
	callframe.Arg(frame, tracking_confidence)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an input for this tracker. It can return a boolean, float or [Vector2] value depending on whether the input is a button, trigger or thumbstick/thumbpad.
*/
//go:nosplit
func (self class) GetInput(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Changes the value for the given input. This method is called by a [XRInterface] implementation and should not be used directly.
*/
//go:nosplit
func (self class) SetInput(name gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnPoseChanged(cb func(pose [1]gdclass.XRPose)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pose_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPoseLostTracking(cb func(pose [1]gdclass.XRPose)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pose_lost_tracking"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonPressed(cb func(name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("button_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonReleased(cb func(name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("button_released"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInputFloatChanged(cb func(name string, value Float.X)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("input_float_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInputVector2Changed(cb func(name string, vector Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("input_vector2_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProfileChanged(cb func(role string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("profile_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsXRPositionalTracker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRPositionalTracker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(XRTracker.Advanced(self.AsXRTracker()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRTracker.Instance(self.AsXRTracker()), name)
	}
}
func init() {
	gdclass.Register("XRPositionalTracker", func(ptr gd.Object) any {
		return [1]gdclass.XRPositionalTracker{*(*gdclass.XRPositionalTracker)(unsafe.Pointer(&ptr))}
	})
}

type TrackerHand = gdclass.XRPositionalTrackerTrackerHand

const (
	/*The hand this tracker is held in is unknown or not applicable.*/
	TrackerHandUnknown TrackerHand = 0
	/*This tracker is the left hand controller.*/
	TrackerHandLeft TrackerHand = 1
	/*This tracker is the right hand controller.*/
	TrackerHandRight TrackerHand = 2
	/*Represents the size of the [enum TrackerHand] enum.*/
	TrackerHandMax TrackerHand = 3
)
