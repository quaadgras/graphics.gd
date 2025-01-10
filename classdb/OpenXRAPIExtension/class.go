// Package OpenXRAPIExtension provides methods for working with OpenXRAPIExtension object instances.
package OpenXRAPIExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
[OpenXRAPIExtension] makes OpenXR available for GDExtension. It provides the OpenXR API to GDExtension through the [method get_instance_proc_addr] method, and the OpenXR instance through [method get_instance].
It also provides methods for querying the status of OpenXR initialization, and helper methods for ease of use of the API with GDExtension.
*/
type Instance [1]gdclass.OpenXRAPIExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRAPIExtension() Instance
}

/*
Returns the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrInstance.html]XrInstance[/url] created during the initialization of the OpenXR API.
*/
func (self Instance) GetInstance() int {
	return int(int(class(self).GetInstance()))
}

/*
Returns the id of the system, which is a [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSystemId.html]XrSystemId[/url] cast to an integer.
*/
func (self Instance) GetSystemId() int {
	return int(int(class(self).GetSystemId()))
}

/*
Returns the OpenXR session, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSession.html]XrSession[/url] cast to an integer.
*/
func (self Instance) GetSession() int {
	return int(int(class(self).GetSession()))
}

/*
Creates a [Transform3D] from an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrPosef.html]XrPosef[/url].
*/
func (self Instance) TransformFromPose(pose unsafe.Pointer) Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).TransformFromPose(pose))
}

/*
Returns [code]true[/code] if the provided [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] (cast to an integer) is successful. Otherwise returns [code]false[/code] and prints the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] converted to a string, with the specified additional information.
*/
func (self Instance) XrResult(result int, format string, args Array.Any) bool {
	return bool(class(self).XrResult(gd.Int(result), gd.NewString(format), args))
}

/*
Returns [code]true[/code] if OpenXR is enabled.
*/
func OpenxrIsEnabled(check_run_in_editor bool) bool {
	self := Instance{}
	return bool(class(self).OpenxrIsEnabled(check_run_in_editor))
}

/*
Returns the function pointer of the OpenXR function with the specified name, cast to an integer. If the function with the given name does not exist, the method returns [code]0[/code].
[b]Note:[/b] [code]openxr/util.h[/code] contains utility macros for acquiring OpenXR functions, e.g. [code]GDEXTENSION_INIT_XR_FUNC_V(xrCreateAction)[/code].
*/
func (self Instance) GetInstanceProcAddr(name string) int {
	return int(int(class(self).GetInstanceProcAddr(gd.NewString(name))))
}

/*
Returns an error string for the given [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url].
*/
func (self Instance) GetErrorString(result int) string {
	return string(class(self).GetErrorString(gd.Int(result)).String())
}

/*
Returns the name of the specified swapchain format.
*/
func (self Instance) GetSwapchainFormatName(swapchain_format int) string {
	return string(class(self).GetSwapchainFormatName(gd.Int(swapchain_format)).String())
}

/*
Returns [code]true[/code] if OpenXR is initialized.
*/
func (self Instance) IsInitialized() bool {
	return bool(class(self).IsInitialized())
}

/*
Returns [code]true[/code] if OpenXR is running ([url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/xrBeginSession.html]xrBeginSession[/url] was successfully called and the swapchains were created).
*/
func (self Instance) IsRunning() bool {
	return bool(class(self).IsRunning())
}

/*
Returns the play space, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSpace.html]XrSpace[/url] cast to an integer.
*/
func (self Instance) GetPlaySpace() int {
	return int(int(class(self).GetPlaySpace()))
}

/*
Returns the predicted display timing for the current frame.
*/
func (self Instance) GetPredictedDisplayTime() int {
	return int(int(class(self).GetPredictedDisplayTime()))
}

/*
Returns the predicted display timing for the next frame.
*/
func (self Instance) GetNextFrameTime() int {
	return int(int(class(self).GetNextFrameTime()))
}

/*
Returns [code]true[/code] if OpenXR is initialized for rendering with an XR viewport.
*/
func (self Instance) CanRender() bool {
	return bool(class(self).CanRender())
}

/*
Returns the corresponding [code]XRHandTrackerEXT[/code] handle for the given hand index value.
*/
func (self Instance) GetHandTracker(hand_index int) int {
	return int(int(class(self).GetHandTracker(gd.Int(hand_index))))
}

/*
Registers the given extension as a composition layer provider.
*/
func (self Instance) RegisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) {
	class(self).RegisterCompositionLayerProvider(extension)
}

/*
Unregisters the given extension as a composition layer provider.
*/
func (self Instance) UnregisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) {
	class(self).UnregisterCompositionLayerProvider(extension)
}

/*
If set to [code]true[/code], an OpenXR extension is loaded which is capable of emulating the [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] blend mode.
*/
func (self Instance) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool) {
	class(self).SetEmulateEnvironmentBlendModeAlphaBlend(enabled)
}

/*
Returns [enum OpenXRAPIExtension.OpenXRAlphaBlendModeSupport] denoting if [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported, emulated or not supported at all.
*/
func (self Instance) IsEnvironmentBlendModeAlphaSupported() gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport {
	return gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport(class(self).IsEnvironmentBlendModeAlphaSupported())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRAPIExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRAPIExtension"))
	casted := Instance{*(*gdclass.OpenXRAPIExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrInstance.html]XrInstance[/url] created during the initialization of the OpenXR API.
*/
//go:nosplit
func (self class) GetInstance() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the id of the system, which is a [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSystemId.html]XrSystemId[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetSystemId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_system_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the OpenXR session, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSession.html]XrSession[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetSession() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_session, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a [Transform3D] from an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrPosef.html]XrPosef[/url].
*/
//go:nosplit
func (self class) TransformFromPose(pose unsafe.Pointer) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, pose)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_transform_from_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the provided [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] (cast to an integer) is successful. Otherwise returns [code]false[/code] and prints the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] converted to a string, with the specified additional information.
*/
//go:nosplit
func (self class) XrResult(result gd.Int, format gd.String, args gd.Array) bool {
	var frame = callframe.New()
	callframe.Arg(frame, result)
	callframe.Arg(frame, pointers.Get(format))
	callframe.Arg(frame, pointers.Get(args))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_xr_result, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is enabled.
*/
//go:nosplit
func (self class) OpenxrIsEnabled(check_run_in_editor bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, check_run_in_editor)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the function pointer of the OpenXR function with the specified name, cast to an integer. If the function with the given name does not exist, the method returns [code]0[/code].
[b]Note:[/b] [code]openxr/util.h[/code] contains utility macros for acquiring OpenXR functions, e.g. [code]GDEXTENSION_INIT_XR_FUNC_V(xrCreateAction)[/code].
*/
//go:nosplit
func (self class) GetInstanceProcAddr(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_instance_proc_addr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an error string for the given [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url].
*/
//go:nosplit
func (self class) GetErrorString(result gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, result)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_error_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the name of the specified swapchain format.
*/
//go:nosplit
func (self class) GetSwapchainFormatName(swapchain_format gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, swapchain_format)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_swapchain_format_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is initialized.
*/
//go:nosplit
func (self class) IsInitialized() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_is_initialized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is running ([url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/xrBeginSession.html]xrBeginSession[/url] was successfully called and the swapchains were created).
*/
//go:nosplit
func (self class) IsRunning() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_is_running, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the play space, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSpace.html]XrSpace[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetPlaySpace() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_play_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the predicted display timing for the current frame.
*/
//go:nosplit
func (self class) GetPredictedDisplayTime() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_predicted_display_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the predicted display timing for the next frame.
*/
//go:nosplit
func (self class) GetNextFrameTime() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_next_frame_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is initialized for rendering with an XR viewport.
*/
//go:nosplit
func (self class) CanRender() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_can_render, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the corresponding [code]XRHandTrackerEXT[/code] handle for the given hand index value.
*/
//go:nosplit
func (self class) GetHandTracker(hand_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, hand_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Registers the given extension as a composition layer provider.
*/
//go:nosplit
func (self class) RegisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(extension[0].AsObject()[0]))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_register_composition_layer_provider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Unregisters the given extension as a composition layer provider.
*/
//go:nosplit
func (self class) UnregisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_unregister_composition_layer_provider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If set to [code]true[/code], an OpenXR extension is loaded which is capable of emulating the [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] blend mode.
*/
//go:nosplit
func (self class) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_emulate_environment_blend_mode_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [enum OpenXRAPIExtension.OpenXRAlphaBlendModeSupport] denoting if [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported, emulated or not supported at all.
*/
//go:nosplit
func (self class) IsEnvironmentBlendModeAlphaSupported() gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_is_environment_blend_mode_alpha_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRAPIExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRAPIExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("OpenXRAPIExtension", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRAPIExtension{*(*gdclass.OpenXRAPIExtension)(unsafe.Pointer(&ptr))}
	})
}

type OpenXRAlphaBlendModeSupport = gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport

const (
	/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] isn't supported at all.*/
	OpenxrAlphaBlendModeSupportNone OpenXRAlphaBlendModeSupport = 0
	/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported.*/
	OpenxrAlphaBlendModeSupportReal OpenXRAlphaBlendModeSupport = 1
	/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is emulated.*/
	OpenxrAlphaBlendModeSupportEmulating OpenXRAlphaBlendModeSupport = 2
)
