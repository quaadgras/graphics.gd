package OpenXRInterface

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRInterface"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The OpenXR interface allows Godot to interact with OpenXR runtimes and make it possible to create XR experiences and games.
Due to the needs of OpenXR this interface works slightly different than other plugin based XR interfaces. It needs to be initialized when Godot starts. You need to enable OpenXR, settings for this can be found in your games project settings under the XR heading. You do need to mark a viewport for use with XR in order for Godot to know which render result should be output to the headset.

*/
type Go [1]classdb.OpenXRInterface

/*
Returns [code]true[/code] if OpenXR's foveation extension is supported, the interface must be initialized before this returns a valid value.
[b]Note:[/b] This feature is only available on the compatibility renderer and currently only available on some stand alone headsets. For Vulkan set [member Viewport.vrs_mode] to [code]VRS_XR[/code] on desktop.
*/
func (self Go) IsFoveationSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsFoveationSupported())
}

/*
Returns [code]true[/code] if the given action set is active.
*/
func (self Go) IsActionSetActive(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsActionSetActive(gc.String(name)))
}

/*
Sets the given action set as active or inactive.
*/
func (self Go) SetActionSetActive(name string, active bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetActionSetActive(gc.String(name), active)
}

/*
Returns a list of action sets registered with Godot (loaded from the action map at runtime).
*/
func (self Go) GetActionSets() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetActionSets(gc))
}

/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the OpenXR runtime and after the interface has been initialized.
*/
func (self Go) GetAvailableDisplayRefreshRates() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetAvailableDisplayRefreshRates(gc))
}

/*
If handtracking is enabled and motion range is supported, sets the currently configured motion range for [param hand] to [param motion_range].
*/
func (self Go) SetMotionRange(hand classdb.OpenXRInterfaceHand, motion_range classdb.OpenXRInterfaceHandMotionRange) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMotionRange(hand, motion_range)
}

/*
If handtracking is enabled and motion range is supported, gets the currently configured motion range for [param hand].
*/
func (self Go) GetMotionRange(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandMotionRange {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRInterfaceHandMotionRange(class(self).GetMotionRange(hand))
}

/*
If handtracking is enabled and hand tracking source is supported, gets the source of the hand tracking data for [param hand].
*/
func (self Go) GetHandTrackingSource(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandTrackedSource {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRInterfaceHandTrackedSource(class(self).GetHandTrackingSource(hand))
}

/*
If handtracking is enabled, returns flags that inform us of the validity of the tracking data.
*/
func (self Go) GetHandJointFlags(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) classdb.OpenXRInterfaceHandJointFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRInterfaceHandJointFlags(class(self).GetHandJointFlags(hand, joint))
}

/*
If handtracking is enabled, returns the rotation of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR.
*/
func (self Go) GetHandJointRotation(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Quaternion {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Quaternion(class(self).GetHandJointRotation(hand, joint))
}

/*
If handtracking is enabled, returns the position of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
func (self Go) GetHandJointPosition(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetHandJointPosition(hand, joint))
}

/*
If handtracking is enabled, returns the radius of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is without worldscale applied!
*/
func (self Go) GetHandJointRadius(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetHandJointRadius(hand, joint)))
}

/*
If handtracking is enabled, returns the linear velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
func (self Go) GetHandJointLinearVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetHandJointLinearVelocity(hand, joint))
}

/*
If handtracking is enabled, returns the angular velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D]!
*/
func (self Go) GetHandJointAngularVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetHandJointAngularVelocity(hand, joint))
}

/*
Returns [code]true[/code] if OpenXR's hand tracking is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
func (self Go) IsHandTrackingSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsHandTrackingSupported())
}

/*
Returns [code]true[/code] if OpenXR's hand interaction profile is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
func (self Go) IsHandInteractionSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsHandInteractionSupported())
}

/*
Returns the capabilities of the eye gaze interaction extension.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
func (self Go) IsEyeGazeInteractionSupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsEyeGazeInteractionSupported())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRInterface
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRInterface"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DisplayRefreshRate() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDisplayRefreshRate()))
}

func (self Go) SetDisplayRefreshRate(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisplayRefreshRate(gd.Float(value))
}

func (self Go) RenderTargetSizeMultiplier() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRenderTargetSizeMultiplier()))
}

func (self Go) SetRenderTargetSizeMultiplier(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRenderTargetSizeMultiplier(gd.Float(value))
}

func (self Go) FoveationLevel() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetFoveationLevel()))
}

func (self Go) SetFoveationLevel(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFoveationLevel(gd.Int(value))
}

func (self Go) FoveationDynamic() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFoveationDynamic())
}

func (self Go) SetFoveationDynamic(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFoveationDynamic(value)
}

func (self Go) VrsMinRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVrsMinRadius()))
}

func (self Go) SetVrsMinRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVrsMinRadius(gd.Float(value))
}

func (self Go) VrsStrength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVrsStrength()))
}

func (self Go) SetVrsStrength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVrsStrength(gd.Float(value))
}

//go:nosplit
func (self class) GetDisplayRefreshRate() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisplayRefreshRate(refresh_rate gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, refresh_rate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderTargetSizeMultiplier() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_render_target_size_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderTargetSizeMultiplier(multiplier gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_render_target_size_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if OpenXR's foveation extension is supported, the interface must be initialized before this returns a valid value.
[b]Note:[/b] This feature is only available on the compatibility renderer and currently only available on some stand alone headsets. For Vulkan set [member Viewport.vrs_mode] to [code]VRS_XR[/code] on desktop.
*/
//go:nosplit
func (self class) IsFoveationSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_foveation_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFoveationLevel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_foveation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFoveationLevel(foveation_level gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, foveation_level)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_foveation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFoveationDynamic() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_foveation_dynamic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFoveationDynamic(foveation_dynamic bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, foveation_dynamic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_foveation_dynamic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given action set is active.
*/
//go:nosplit
func (self class) IsActionSetActive(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_action_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given action set as active or inactive.
*/
//go:nosplit
func (self class) SetActionSetActive(name gd.String, active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_action_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of action sets registered with Godot (loaded from the action map at runtime).
*/
//go:nosplit
func (self class) GetActionSets(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the OpenXR runtime and after the interface has been initialized.
*/
//go:nosplit
func (self class) GetAvailableDisplayRefreshRates(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_available_display_refresh_rates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If handtracking is enabled and motion range is supported, sets the currently configured motion range for [param hand] to [param motion_range].
*/
//go:nosplit
func (self class) SetMotionRange(hand classdb.OpenXRInterfaceHand, motion_range classdb.OpenXRInterfaceHandMotionRange)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, motion_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If handtracking is enabled and motion range is supported, gets the currently configured motion range for [param hand].
*/
//go:nosplit
func (self class) GetMotionRange(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandMotionRange {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret = callframe.Ret[classdb.OpenXRInterfaceHandMotionRange](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_motion_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled and hand tracking source is supported, gets the source of the hand tracking data for [param hand].
*/
//go:nosplit
func (self class) GetHandTrackingSource(hand classdb.OpenXRInterfaceHand) classdb.OpenXRInterfaceHandTrackedSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret = callframe.Ret[classdb.OpenXRInterfaceHandTrackedSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns flags that inform us of the validity of the tracking data.
*/
//go:nosplit
func (self class) GetHandJointFlags(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) classdb.OpenXRInterfaceHandJointFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[classdb.OpenXRInterfaceHandJointFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the rotation of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR.
*/
//go:nosplit
func (self class) GetHandJointRotation(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Quaternion {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the position of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointPosition(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the radius of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointRadius(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the linear velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointLinearVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If handtracking is enabled, returns the angular velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D]!
*/
//go:nosplit
func (self class) GetHandJointAngularVelocity(hand classdb.OpenXRInterfaceHand, joint classdb.OpenXRInterfaceHandJoints) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR's hand tracking is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
//go:nosplit
func (self class) IsHandTrackingSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_hand_tracking_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR's hand interaction profile is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
//go:nosplit
func (self class) IsHandInteractionSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_hand_interaction_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the capabilities of the eye gaze interaction extension.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
//go:nosplit
func (self class) IsEyeGazeInteractionSupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_is_eye_gaze_interaction_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetVrsMinRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsMinRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVrsStrength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsStrength(strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRInterface.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnSessionBegun(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("session_begun"), gc.Callable(cb), 0)
}


func (self Go) OnSessionStopping(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("session_stopping"), gc.Callable(cb), 0)
}


func (self Go) OnSessionFocussed(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("session_focussed"), gc.Callable(cb), 0)
}


func (self Go) OnSessionVisible(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("session_visible"), gc.Callable(cb), 0)
}


func (self Go) OnSessionLossPending(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("session_loss_pending"), gc.Callable(cb), 0)
}


func (self Go) OnInstanceExiting(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("instance_exiting"), gc.Callable(cb), 0)
}


func (self Go) OnPoseRecentered(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pose_recentered"), gc.Callable(cb), 0)
}


func (self Go) OnRefreshRateChanged(cb func(refresh_rate float64)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("refresh_rate_changed"), gc.Callable(cb), 0)
}


func (self class) AsOpenXRInterface() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRInterface() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRInterface() XRInterface.GD { return *((*XRInterface.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRInterface() XRInterface.Go { return *((*XRInterface.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRInterface(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRInterface(), name)
	}
}
func init() {classdb.Register("OpenXRInterface", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Hand = classdb.OpenXRInterfaceHand

const (
/*Left hand.*/
	HandLeft Hand = 0
/*Right hand.*/
	HandRight Hand = 1
/*Maximum value for the hand enum.*/
	HandMax Hand = 2
)
type HandMotionRange = classdb.OpenXRInterfaceHandMotionRange

const (
/*Full hand range, if user closes their hands, we make a full fist.*/
	HandMotionRangeUnobstructed HandMotionRange = 0
/*Conform to controller, if user closes their hands, the tracked data conforms to the shape of the controller.*/
	HandMotionRangeConformToController HandMotionRange = 1
/*Maximum value for the motion range enum.*/
	HandMotionRangeMax HandMotionRange = 2
)
type HandTrackedSource = classdb.OpenXRInterfaceHandTrackedSource

const (
/*The source of hand tracking data is unknown (the extension is likely unsupported).*/
	HandTrackedSourceUnknown HandTrackedSource = 0
/*The source of hand tracking is unobstructed, this means that an accurate method of hand tracking is used, e.g. optical hand tracking, data gloves, etc.*/
	HandTrackedSourceUnobstructed HandTrackedSource = 1
/*The source of hand tracking is a controller, bone positions are inferred from controller inputs.*/
	HandTrackedSourceController HandTrackedSource = 2
/*Maximum value for the hand tracked source enum.*/
	HandTrackedSourceMax HandTrackedSource = 3
)
type HandJoints = classdb.OpenXRInterfaceHandJoints

const (
/*Palm joint.*/
	HandJointPalm HandJoints = 0
/*Wrist joint.*/
	HandJointWrist HandJoints = 1
/*Thumb metacarpal joint.*/
	HandJointThumbMetacarpal HandJoints = 2
/*Thumb proximal joint.*/
	HandJointThumbProximal HandJoints = 3
/*Thumb distal joint.*/
	HandJointThumbDistal HandJoints = 4
/*Thumb tip joint.*/
	HandJointThumbTip HandJoints = 5
/*Index metacarpal joint.*/
	HandJointIndexMetacarpal HandJoints = 6
/*Index proximal joint.*/
	HandJointIndexProximal HandJoints = 7
/*Index intermediate joint.*/
	HandJointIndexIntermediate HandJoints = 8
/*Index distal joint.*/
	HandJointIndexDistal HandJoints = 9
/*Index tip joint.*/
	HandJointIndexTip HandJoints = 10
/*Middle metacarpal joint.*/
	HandJointMiddleMetacarpal HandJoints = 11
/*Middle proximal joint.*/
	HandJointMiddleProximal HandJoints = 12
/*Middle intermediate joint.*/
	HandJointMiddleIntermediate HandJoints = 13
/*Middle distal joint.*/
	HandJointMiddleDistal HandJoints = 14
/*Middle tip joint.*/
	HandJointMiddleTip HandJoints = 15
/*Ring metacarpal joint.*/
	HandJointRingMetacarpal HandJoints = 16
/*Ring proximal joint.*/
	HandJointRingProximal HandJoints = 17
/*Ring intermediate joint.*/
	HandJointRingIntermediate HandJoints = 18
/*Ring distal joint.*/
	HandJointRingDistal HandJoints = 19
/*Ring tip joint.*/
	HandJointRingTip HandJoints = 20
/*Little metacarpal joint.*/
	HandJointLittleMetacarpal HandJoints = 21
/*Little proximal joint.*/
	HandJointLittleProximal HandJoints = 22
/*Little intermediate joint.*/
	HandJointLittleIntermediate HandJoints = 23
/*Little distal joint.*/
	HandJointLittleDistal HandJoints = 24
/*Little tip joint.*/
	HandJointLittleTip HandJoints = 25
/*Maximum value for the hand joint enum.*/
	HandJointMax HandJoints = 26
)
type HandJointFlags = classdb.OpenXRInterfaceHandJointFlags

const (
/*No flags are set.*/
	HandJointNone HandJointFlags = 0
/*If set, the orientation data is valid, otherwise, the orientation data is unreliable and should not be used.*/
	HandJointOrientationValid HandJointFlags = 1
/*If set, the orientation data comes from tracking data, otherwise, the orientation data contains predicted data.*/
	HandJointOrientationTracked HandJointFlags = 2
/*If set, the positional data is valid, otherwise, the positional data is unreliable and should not be used.*/
	HandJointPositionValid HandJointFlags = 4
/*If set, the positional data comes from tracking data, otherwise, the positional data contains predicted data.*/
	HandJointPositionTracked HandJointFlags = 8
/*If set, our linear velocity data is valid, otherwise, the linear velocity data is unreliable and should not be used.*/
	HandJointLinearVelocityValid HandJointFlags = 16
/*If set, our angular velocity data is valid, otherwise, the angular velocity data is unreliable and should not be used.*/
	HandJointAngularVelocityValid HandJointFlags = 32
)
