// Package OpenXRInterface provides methods for working with OpenXRInterface object instances.
package OpenXRInterface

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/XRInterface"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The OpenXR interface allows Godot to interact with OpenXR runtimes and make it possible to create XR experiences and games.
Due to the needs of OpenXR this interface works slightly different than other plugin based XR interfaces. It needs to be initialized when Godot starts. You need to enable OpenXR, settings for this can be found in your games project settings under the XR heading. You do need to mark a viewport for use with XR in order for Godot to know which render result should be output to the headset.
*/
type Instance [1]gdclass.OpenXRInterface

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRInterface() Instance
}

/*
Returns [code]true[/code] if OpenXR's foveation extension is supported, the interface must be initialized before this returns a valid value.
[b]Note:[/b] This feature is only available on the compatibility renderer and currently only available on some stand alone headsets. For Vulkan set [member Viewport.vrs_mode] to [code]VRS_XR[/code] on desktop.
*/
func (self Instance) IsFoveationSupported() bool {
	return bool(class(self).IsFoveationSupported())
}

/*
Returns [code]true[/code] if the given action set is active.
*/
func (self Instance) IsActionSetActive(name string) bool {
	return bool(class(self).IsActionSetActive(gd.NewString(name)))
}

/*
Sets the given action set as active or inactive.
*/
func (self Instance) SetActionSetActive(name string, active bool) {
	class(self).SetActionSetActive(gd.NewString(name), active)
}

/*
Returns a list of action sets registered with Godot (loaded from the action map at runtime).
*/
func (self Instance) GetActionSets() []any {
	return []any(gd.ArrayAs[[]any](class(self).GetActionSets()))
}

/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the OpenXR runtime and after the interface has been initialized.
*/
func (self Instance) GetAvailableDisplayRefreshRates() []any {
	return []any(gd.ArrayAs[[]any](class(self).GetAvailableDisplayRefreshRates()))
}

/*
If handtracking is enabled and motion range is supported, sets the currently configured motion range for [param hand] to [param motion_range].
*/
func (self Instance) SetMotionRange(hand gdclass.OpenXRInterfaceHand, motion_range gdclass.OpenXRInterfaceHandMotionRange) {
	class(self).SetMotionRange(hand, motion_range)
}

/*
If handtracking is enabled and motion range is supported, gets the currently configured motion range for [param hand].
*/
func (self Instance) GetMotionRange(hand gdclass.OpenXRInterfaceHand) gdclass.OpenXRInterfaceHandMotionRange {
	return gdclass.OpenXRInterfaceHandMotionRange(class(self).GetMotionRange(hand))
}

/*
If handtracking is enabled and hand tracking source is supported, gets the source of the hand tracking data for [param hand].
*/
func (self Instance) GetHandTrackingSource(hand gdclass.OpenXRInterfaceHand) gdclass.OpenXRInterfaceHandTrackedSource {
	return gdclass.OpenXRInterfaceHandTrackedSource(class(self).GetHandTrackingSource(hand))
}

/*
If handtracking is enabled, returns flags that inform us of the validity of the tracking data.
*/
func (self Instance) GetHandJointFlags(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gdclass.OpenXRInterfaceHandJointFlags {
	return gdclass.OpenXRInterfaceHandJointFlags(class(self).GetHandJointFlags(hand, joint))
}

/*
If handtracking is enabled, returns the rotation of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR.
*/
func (self Instance) GetHandJointRotation(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetHandJointRotation(hand, joint))
}

/*
If handtracking is enabled, returns the position of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
func (self Instance) GetHandJointPosition(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetHandJointPosition(hand, joint))
}

/*
If handtracking is enabled, returns the radius of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is without worldscale applied!
*/
func (self Instance) GetHandJointRadius(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) Float.X {
	return Float.X(Float.X(class(self).GetHandJointRadius(hand, joint)))
}

/*
If handtracking is enabled, returns the linear velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
func (self Instance) GetHandJointLinearVelocity(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetHandJointLinearVelocity(hand, joint))
}

/*
If handtracking is enabled, returns the angular velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D]!
*/
func (self Instance) GetHandJointAngularVelocity(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetHandJointAngularVelocity(hand, joint))
}

/*
Returns [code]true[/code] if OpenXR's hand tracking is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
func (self Instance) IsHandTrackingSupported() bool {
	return bool(class(self).IsHandTrackingSupported())
}

/*
Returns [code]true[/code] if OpenXR's hand interaction profile is supported and enabled.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
func (self Instance) IsHandInteractionSupported() bool {
	return bool(class(self).IsHandInteractionSupported())
}

/*
Returns the capabilities of the eye gaze interaction extension.
[b]Note:[/b] This only returns a valid value after OpenXR has been initialized.
*/
func (self Instance) IsEyeGazeInteractionSupported() bool {
	return bool(class(self).IsEyeGazeInteractionSupported())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRInterface

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRInterface"))
	casted := Instance{*(*gdclass.OpenXRInterface)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DisplayRefreshRate() Float.X {
	return Float.X(Float.X(class(self).GetDisplayRefreshRate()))
}

func (self Instance) SetDisplayRefreshRate(value Float.X) {
	class(self).SetDisplayRefreshRate(gd.Float(value))
}

func (self Instance) RenderTargetSizeMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetRenderTargetSizeMultiplier()))
}

func (self Instance) SetRenderTargetSizeMultiplier(value Float.X) {
	class(self).SetRenderTargetSizeMultiplier(gd.Float(value))
}

func (self Instance) FoveationLevel() int {
	return int(int(class(self).GetFoveationLevel()))
}

func (self Instance) SetFoveationLevel(value int) {
	class(self).SetFoveationLevel(gd.Int(value))
}

func (self Instance) FoveationDynamic() bool {
	return bool(class(self).GetFoveationDynamic())
}

func (self Instance) SetFoveationDynamic(value bool) {
	class(self).SetFoveationDynamic(value)
}

func (self Instance) VrsMinRadius() Float.X {
	return Float.X(Float.X(class(self).GetVrsMinRadius()))
}

func (self Instance) SetVrsMinRadius(value Float.X) {
	class(self).SetVrsMinRadius(gd.Float(value))
}

func (self Instance) VrsStrength() Float.X {
	return Float.X(Float.X(class(self).GetVrsStrength()))
}

func (self Instance) SetVrsStrength(value Float.X) {
	class(self).SetVrsStrength(gd.Float(value))
}

//go:nosplit
func (self class) GetDisplayRefreshRate() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisplayRefreshRate(refresh_rate gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, refresh_rate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_display_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRenderTargetSizeMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_render_target_size_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderTargetSizeMultiplier(multiplier gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_render_target_size_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if OpenXR's foveation extension is supported, the interface must be initialized before this returns a valid value.
[b]Note:[/b] This feature is only available on the compatibility renderer and currently only available on some stand alone headsets. For Vulkan set [member Viewport.vrs_mode] to [code]VRS_XR[/code] on desktop.
*/
//go:nosplit
func (self class) IsFoveationSupported() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_is_foveation_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFoveationLevel() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_foveation_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFoveationLevel(foveation_level gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, foveation_level)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_foveation_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFoveationDynamic() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_foveation_dynamic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFoveationDynamic(foveation_dynamic bool) {
	var frame = callframe.New()
	callframe.Arg(frame, foveation_dynamic)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_foveation_dynamic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given action set is active.
*/
//go:nosplit
func (self class) IsActionSetActive(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_is_action_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given action set as active or inactive.
*/
//go:nosplit
func (self class) SetActionSetActive(name gd.String, active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_action_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of action sets registered with Godot (loaded from the action map at runtime).
*/
//go:nosplit
func (self class) GetActionSets() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns display refresh rates supported by the current HMD. Only returned if this feature is supported by the OpenXR runtime and after the interface has been initialized.
*/
//go:nosplit
func (self class) GetAvailableDisplayRefreshRates() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_available_display_refresh_rates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
If handtracking is enabled and motion range is supported, sets the currently configured motion range for [param hand] to [param motion_range].
*/
//go:nosplit
func (self class) SetMotionRange(hand gdclass.OpenXRInterfaceHand, motion_range gdclass.OpenXRInterfaceHandMotionRange) {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, motion_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_motion_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If handtracking is enabled and motion range is supported, gets the currently configured motion range for [param hand].
*/
//go:nosplit
func (self class) GetMotionRange(hand gdclass.OpenXRInterfaceHand) gdclass.OpenXRInterfaceHandMotionRange {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret = callframe.Ret[gdclass.OpenXRInterfaceHandMotionRange](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_motion_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled and hand tracking source is supported, gets the source of the hand tracking data for [param hand].
*/
//go:nosplit
func (self class) GetHandTrackingSource(hand gdclass.OpenXRInterfaceHand) gdclass.OpenXRInterfaceHandTrackedSource {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret = callframe.Ret[gdclass.OpenXRInterfaceHandTrackedSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled, returns flags that inform us of the validity of the tracking data.
*/
//go:nosplit
func (self class) GetHandJointFlags(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gdclass.OpenXRInterfaceHandJointFlags {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gdclass.OpenXRInterfaceHandJointFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled, returns the rotation of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR.
*/
//go:nosplit
func (self class) GetHandJointRotation(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gd.Quaternion {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_joint_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled, returns the position of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointPosition(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_joint_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled, returns the radius of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointRadius(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled, returns the linear velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D] without worldscale applied!
*/
//go:nosplit
func (self class) GetHandJointLinearVelocity(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If handtracking is enabled, returns the angular velocity of a joint ([param joint]) of a hand ([param hand]) as provided by OpenXR. This is relative to [XROrigin3D]!
*/
//go:nosplit
func (self class) GetHandJointAngularVelocity(hand gdclass.OpenXRInterfaceHand, joint gdclass.OpenXRInterfaceHandJoints) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_is_hand_tracking_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_is_hand_interaction_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_is_eye_gaze_interaction_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVrsMinRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsMinRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsStrength(strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInterface.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnSessionBegun(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("session_begun"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSessionStopping(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("session_stopping"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSessionFocussed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("session_focussed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSessionVisible(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("session_visible"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSessionLossPending(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("session_loss_pending"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInstanceExiting(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("instance_exiting"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPoseRecentered(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pose_recentered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRefreshRateChanged(cb func(refresh_rate Float.X)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("refresh_rate_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsOpenXRInterface() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRInterface() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRInterface() XRInterface.Advanced {
	return *((*XRInterface.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRInterface() XRInterface.Instance {
	return *((*XRInterface.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(XRInterface.Advanced(self.AsXRInterface()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(XRInterface.Instance(self.AsXRInterface()), name)
	}
}
func init() {
	gdclass.Register("OpenXRInterface", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRInterface{*(*gdclass.OpenXRInterface)(unsafe.Pointer(&ptr))}
	})
}

type Hand = gdclass.OpenXRInterfaceHand

const (
	/*Left hand.*/
	HandLeft Hand = 0
	/*Right hand.*/
	HandRight Hand = 1
	/*Maximum value for the hand enum.*/
	HandMax Hand = 2
)

type HandMotionRange = gdclass.OpenXRInterfaceHandMotionRange

const (
	/*Full hand range, if user closes their hands, we make a full fist.*/
	HandMotionRangeUnobstructed HandMotionRange = 0
	/*Conform to controller, if user closes their hands, the tracked data conforms to the shape of the controller.*/
	HandMotionRangeConformToController HandMotionRange = 1
	/*Maximum value for the motion range enum.*/
	HandMotionRangeMax HandMotionRange = 2
)

type HandTrackedSource = gdclass.OpenXRInterfaceHandTrackedSource

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

type HandJoints = gdclass.OpenXRInterfaceHandJoints

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

type HandJointFlags = gdclass.OpenXRInterfaceHandJointFlags

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
