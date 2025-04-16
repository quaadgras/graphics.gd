// Package OpenXRAPIExtension provides methods for working with OpenXRAPIExtension object instances.
package OpenXRAPIExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2i"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector2i"

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
func (self Instance) GetInstance() int { //gd:OpenXRAPIExtension.get_instance
	return int(int(Advanced(self).GetInstance()))
}

/*
Returns the id of the system, which is a [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSystemId.html]XrSystemId[/url] cast to an integer.
*/
func (self Instance) GetSystemId() int { //gd:OpenXRAPIExtension.get_system_id
	return int(int(Advanced(self).GetSystemId()))
}

/*
Returns the OpenXR session, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSession.html]XrSession[/url] cast to an integer.
*/
func (self Instance) GetSession() int { //gd:OpenXRAPIExtension.get_session
	return int(int(Advanced(self).GetSession()))
}

/*
Creates a [Transform3D] from an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrPosef.html]XrPosef[/url].
*/
func (self Instance) TransformFromPose(pose unsafe.Pointer) Transform3D.BasisOrigin { //gd:OpenXRAPIExtension.transform_from_pose
	return Transform3D.BasisOrigin(Advanced(self).TransformFromPose(pose))
}

/*
Returns [code]true[/code] if the provided [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] (cast to an integer) is successful. Otherwise returns [code]false[/code] and prints the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] converted to a string, with the specified additional information.
*/
func (self Instance) XrResult(result int, format string, args []any) bool { //gd:OpenXRAPIExtension.xr_result
	return bool(Advanced(self).XrResult(int64(result), String.New(format), gd.EngineArrayFromSlice(args)))
}

/*
Returns [code]true[/code] if OpenXR is enabled.
*/
func OpenxrIsEnabled(check_run_in_editor bool) bool { //gd:OpenXRAPIExtension.openxr_is_enabled
	self := Instance{}
	return bool(Advanced(self).OpenxrIsEnabled(check_run_in_editor))
}

/*
Returns the function pointer of the OpenXR function with the specified name, cast to an integer. If the function with the given name does not exist, the method returns [code]0[/code].
[b]Note:[/b] [code]openxr/util.h[/code] contains utility macros for acquiring OpenXR functions, e.g. [code]GDEXTENSION_INIT_XR_FUNC_V(xrCreateAction)[/code].
*/
func (self Instance) GetInstanceProcAddr(name string) int { //gd:OpenXRAPIExtension.get_instance_proc_addr
	return int(int(Advanced(self).GetInstanceProcAddr(String.New(name))))
}

/*
Returns an error string for the given [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url].
*/
func (self Instance) GetErrorString(result int) string { //gd:OpenXRAPIExtension.get_error_string
	return string(Advanced(self).GetErrorString(int64(result)).String())
}

/*
Returns the name of the specified swapchain format.
*/
func (self Instance) GetSwapchainFormatName(swapchain_format int) string { //gd:OpenXRAPIExtension.get_swapchain_format_name
	return string(Advanced(self).GetSwapchainFormatName(int64(swapchain_format)).String())
}

/*
Set the object name of an OpenXR object, used for debug output. [param object_type] must be a valid OpenXR [code]XrObjectType[/code] enum and [param object_handle] must be a valid OpenXR object handle.
*/
func (self Instance) SetObjectName(object_type int, object_handle int, object_name string) { //gd:OpenXRAPIExtension.set_object_name
	Advanced(self).SetObjectName(int64(object_type), int64(object_handle), String.New(object_name))
}

/*
Begins a new debug label region, this label will be reported in debug messages for any calls following this until [method end_debug_label_region] is called. Debug labels can be stacked.
*/
func (self Instance) BeginDebugLabelRegion(label_name string) { //gd:OpenXRAPIExtension.begin_debug_label_region
	Advanced(self).BeginDebugLabelRegion(String.New(label_name))
}

/*
Marks the end of a debug label region. Removes the latest debug label region added by calling [method begin_debug_label_region].
*/
func (self Instance) EndDebugLabelRegion() { //gd:OpenXRAPIExtension.end_debug_label_region
	Advanced(self).EndDebugLabelRegion()
}

/*
Inserts a debug label, this label is reported in any debug message resulting from the OpenXR calls that follows, until any of [method begin_debug_label_region], [method end_debug_label_region], or [method insert_debug_label] is called.
*/
func (self Instance) InsertDebugLabel(label_name string) { //gd:OpenXRAPIExtension.insert_debug_label
	Advanced(self).InsertDebugLabel(String.New(label_name))
}

/*
Returns [code]true[/code] if OpenXR is initialized.
*/
func (self Instance) IsInitialized() bool { //gd:OpenXRAPIExtension.is_initialized
	return bool(Advanced(self).IsInitialized())
}

/*
Returns [code]true[/code] if OpenXR is running ([url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/xrBeginSession.html]xrBeginSession[/url] was successfully called and the swapchains were created).
*/
func (self Instance) IsRunning() bool { //gd:OpenXRAPIExtension.is_running
	return bool(Advanced(self).IsRunning())
}

/*
Returns the play space, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSpace.html]XrSpace[/url] cast to an integer.
*/
func (self Instance) GetPlaySpace() int { //gd:OpenXRAPIExtension.get_play_space
	return int(int(Advanced(self).GetPlaySpace()))
}

/*
Returns the predicted display timing for the current frame.
*/
func (self Instance) GetPredictedDisplayTime() int { //gd:OpenXRAPIExtension.get_predicted_display_time
	return int(int(Advanced(self).GetPredictedDisplayTime()))
}

/*
Returns the predicted display timing for the next frame.
*/
func (self Instance) GetNextFrameTime() int { //gd:OpenXRAPIExtension.get_next_frame_time
	return int(int(Advanced(self).GetNextFrameTime()))
}

/*
Returns [code]true[/code] if OpenXR is initialized for rendering with an XR viewport.
*/
func (self Instance) CanRender() bool { //gd:OpenXRAPIExtension.can_render
	return bool(Advanced(self).CanRender())
}

/*
Returns the [RID] corresponding to an [code]Action[/code] of a matching name, optionally limited to a specified action set.
*/
func (self Instance) FindAction(name string, action_set RID.ActionSet) RID.Action { //gd:OpenXRAPIExtension.find_action
	return RID.Action(Advanced(self).FindAction(String.New(name), RID.Any(action_set)))
}

/*
Returns the corresponding [code]XrAction[/code] OpenXR handle for the given action RID.
*/
func (self Instance) ActionGetHandle(action RID.Action) int { //gd:OpenXRAPIExtension.action_get_handle
	return int(int(Advanced(self).ActionGetHandle(RID.Any(action))))
}

/*
Returns the corresponding [code]XRHandTrackerEXT[/code] handle for the given hand index value.
*/
func (self Instance) GetHandTracker(hand_index int) int { //gd:OpenXRAPIExtension.get_hand_tracker
	return int(int(Advanced(self).GetHandTracker(int64(hand_index))))
}

/*
Registers the given extension as a composition layer provider.
*/
func (self Instance) RegisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.register_composition_layer_provider
	Advanced(self).RegisterCompositionLayerProvider(extension)
}

/*
Unregisters the given extension as a composition layer provider.
*/
func (self Instance) UnregisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.unregister_composition_layer_provider
	Advanced(self).UnregisterCompositionLayerProvider(extension)
}

/*
Registers the given extension as a provider of additional data structures to projections views.
*/
func (self Instance) RegisterProjectionViewsExtension(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.register_projection_views_extension
	Advanced(self).RegisterProjectionViewsExtension(extension)
}

/*
Unregisters the given extension as a provider of additional data structures to projections views.
*/
func (self Instance) UnregisterProjectionViewsExtension(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.unregister_projection_views_extension
	Advanced(self).UnregisterProjectionViewsExtension(extension)
}

/*
Returns the near boundary value of the camera frustum.
[b]Note:[/b] This is only accessible in the render thread.
*/
func (self Instance) GetRenderStateZNear() Float.X { //gd:OpenXRAPIExtension.get_render_state_z_near
	return Float.X(Float.X(Advanced(self).GetRenderStateZNear()))
}

/*
Returns the far boundary value of the camera frustum.
[b]Note:[/b] This is only accessible in the render thread.
*/
func (self Instance) GetRenderStateZFar() Float.X { //gd:OpenXRAPIExtension.get_render_state_z_far
	return Float.X(Float.X(Advanced(self).GetRenderStateZFar()))
}

/*
Sets the render target of the velocity texture.
*/
func (self Instance) SetVelocityTexture(render_target RID.Framebuffer) { //gd:OpenXRAPIExtension.set_velocity_texture
	Advanced(self).SetVelocityTexture(RID.Any(render_target))
}

/*
Sets the render target of the velocity depth texture.
*/
func (self Instance) SetVelocityDepthTexture(render_target RID.Framebuffer) { //gd:OpenXRAPIExtension.set_velocity_depth_texture
	Advanced(self).SetVelocityDepthTexture(RID.Any(render_target))
}

/*
Sets the target size of the velocity and velocity depth textures.
*/
func (self Instance) SetVelocityTargetSize(target_size Vector2i.XY) { //gd:OpenXRAPIExtension.set_velocity_target_size
	Advanced(self).SetVelocityTargetSize(Vector2i.XY(target_size))
}

/*
Returns an array of supported swapchain formats.
*/
func (self Instance) GetSupportedSwapchainFormats() []int64 { //gd:OpenXRAPIExtension.get_supported_swapchain_formats
	return []int64(slices.Collect(Advanced(self).GetSupportedSwapchainFormats().Values()))
}

/*
Returns a pointer to a new swapchain created using the provided parameters.
*/
func (self Instance) OpenxrSwapchainCreate(create_flags int, usage_flags int, swapchain_format int, width int, height int, sample_count int, array_size int) int { //gd:OpenXRAPIExtension.openxr_swapchain_create
	return int(int(Advanced(self).OpenxrSwapchainCreate(int64(create_flags), int64(usage_flags), int64(swapchain_format), int64(width), int64(height), int64(sample_count), int64(array_size))))
}

/*
Destroys the provided swapchain and frees it from memory.
*/
func (self Instance) OpenxrSwapchainFree(swapchain int) { //gd:OpenXRAPIExtension.openxr_swapchain_free
	Advanced(self).OpenxrSwapchainFree(int64(swapchain))
}

/*
Returns the [code]XrSwapchain[/code] handle of the provided swapchain.
*/
func (self Instance) OpenxrSwapchainGetSwapchain(swapchain int) int { //gd:OpenXRAPIExtension.openxr_swapchain_get_swapchain
	return int(int(Advanced(self).OpenxrSwapchainGetSwapchain(int64(swapchain))))
}

/*
Acquires the image of the provided swapchain.
*/
func (self Instance) OpenxrSwapchainAcquire(swapchain int) { //gd:OpenXRAPIExtension.openxr_swapchain_acquire
	Advanced(self).OpenxrSwapchainAcquire(int64(swapchain))
}

/*
Returns the RID of the provided swapchain's image.
*/
func (self Instance) OpenxrSwapchainGetImage(swapchain int) RID.Texture { //gd:OpenXRAPIExtension.openxr_swapchain_get_image
	return RID.Texture(Advanced(self).OpenxrSwapchainGetImage(int64(swapchain)))
}

/*
Releases the image of the provided swapchain.
*/
func (self Instance) OpenxrSwapchainRelease(swapchain int) { //gd:OpenXRAPIExtension.openxr_swapchain_release
	Advanced(self).OpenxrSwapchainRelease(int64(swapchain))
}

/*
Returns a pointer to the render state's [code]XrCompositionLayerProjection[/code] struct.
[b]Note:[/b] This method should only be called from the rendering thread.
*/
func (self Instance) GetProjectionLayer() int { //gd:OpenXRAPIExtension.get_projection_layer
	return int(int(Advanced(self).GetProjectionLayer()))
}

/*
Sets the render region to [param render_region], overriding the normal render target's rect.
*/
func (self Instance) SetRenderRegion(render_region Rect2i.PositionSize) { //gd:OpenXRAPIExtension.set_render_region
	Advanced(self).SetRenderRegion(Rect2i.PositionSize(render_region))
}

/*
If set to [code]true[/code], an OpenXR extension is loaded which is capable of emulating the [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] blend mode.
*/
func (self Instance) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool) { //gd:OpenXRAPIExtension.set_emulate_environment_blend_mode_alpha_blend
	Advanced(self).SetEmulateEnvironmentBlendModeAlphaBlend(enabled)
}

/*
Returns [enum OpenXRAPIExtension.OpenXRAlphaBlendModeSupport] denoting if [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported, emulated or not supported at all.
*/
func (self Instance) IsEnvironmentBlendModeAlphaSupported() gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport { //gd:OpenXRAPIExtension.is_environment_blend_mode_alpha_supported
	return gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport(Advanced(self).IsEnvironmentBlendModeAlphaSupported())
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
func (self class) GetInstance() int64 { //gd:OpenXRAPIExtension.get_instance
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_instance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the id of the system, which is a [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSystemId.html]XrSystemId[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetSystemId() int64 { //gd:OpenXRAPIExtension.get_system_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_system_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the OpenXR session, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSession.html]XrSession[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetSession() int64 { //gd:OpenXRAPIExtension.get_session
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_session, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a [Transform3D] from an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrPosef.html]XrPosef[/url].
*/
//go:nosplit
func (self class) TransformFromPose(pose unsafe.Pointer) Transform3D.BasisOrigin { //gd:OpenXRAPIExtension.transform_from_pose
	var frame = callframe.New()
	callframe.Arg(frame, pose)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_transform_from_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = gd.Transposed(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the provided [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] (cast to an integer) is successful. Otherwise returns [code]false[/code] and prints the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] converted to a string, with the specified additional information.
*/
//go:nosplit
func (self class) XrResult(result int64, format String.Readable, args Array.Any) bool { //gd:OpenXRAPIExtension.xr_result
	var frame = callframe.New()
	callframe.Arg(frame, result)
	callframe.Arg(frame, pointers.Get(gd.InternalString(format)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_xr_result, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is enabled.
*/
//go:nosplit
func (self class) OpenxrIsEnabled(check_run_in_editor bool) bool { //gd:OpenXRAPIExtension.openxr_is_enabled
	var frame = callframe.New()
	callframe.Arg(frame, check_run_in_editor)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the function pointer of the OpenXR function with the specified name, cast to an integer. If the function with the given name does not exist, the method returns [code]0[/code].
[b]Note:[/b] [code]openxr/util.h[/code] contains utility macros for acquiring OpenXR functions, e.g. [code]GDEXTENSION_INIT_XR_FUNC_V(xrCreateAction)[/code].
*/
//go:nosplit
func (self class) GetInstanceProcAddr(name String.Readable) int64 { //gd:OpenXRAPIExtension.get_instance_proc_addr
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_instance_proc_addr, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an error string for the given [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url].
*/
//go:nosplit
func (self class) GetErrorString(result int64) String.Readable { //gd:OpenXRAPIExtension.get_error_string
	var frame = callframe.New()
	callframe.Arg(frame, result)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_error_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the name of the specified swapchain format.
*/
//go:nosplit
func (self class) GetSwapchainFormatName(swapchain_format int64) String.Readable { //gd:OpenXRAPIExtension.get_swapchain_format_name
	var frame = callframe.New()
	callframe.Arg(frame, swapchain_format)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_swapchain_format_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Set the object name of an OpenXR object, used for debug output. [param object_type] must be a valid OpenXR [code]XrObjectType[/code] enum and [param object_handle] must be a valid OpenXR object handle.
*/
//go:nosplit
func (self class) SetObjectName(object_type int64, object_handle int64, object_name String.Readable) { //gd:OpenXRAPIExtension.set_object_name
	var frame = callframe.New()
	callframe.Arg(frame, object_type)
	callframe.Arg(frame, object_handle)
	callframe.Arg(frame, pointers.Get(gd.InternalString(object_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_object_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Begins a new debug label region, this label will be reported in debug messages for any calls following this until [method end_debug_label_region] is called. Debug labels can be stacked.
*/
//go:nosplit
func (self class) BeginDebugLabelRegion(label_name String.Readable) { //gd:OpenXRAPIExtension.begin_debug_label_region
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_begin_debug_label_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Marks the end of a debug label region. Removes the latest debug label region added by calling [method begin_debug_label_region].
*/
//go:nosplit
func (self class) EndDebugLabelRegion() { //gd:OpenXRAPIExtension.end_debug_label_region
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_end_debug_label_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Inserts a debug label, this label is reported in any debug message resulting from the OpenXR calls that follows, until any of [method begin_debug_label_region], [method end_debug_label_region], or [method insert_debug_label] is called.
*/
//go:nosplit
func (self class) InsertDebugLabel(label_name String.Readable) { //gd:OpenXRAPIExtension.insert_debug_label
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(label_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_insert_debug_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if OpenXR is initialized.
*/
//go:nosplit
func (self class) IsInitialized() bool { //gd:OpenXRAPIExtension.is_initialized
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_is_initialized, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is running ([url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/xrBeginSession.html]xrBeginSession[/url] was successfully called and the swapchains were created).
*/
//go:nosplit
func (self class) IsRunning() bool { //gd:OpenXRAPIExtension.is_running
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_is_running, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the play space, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSpace.html]XrSpace[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetPlaySpace() int64 { //gd:OpenXRAPIExtension.get_play_space
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_play_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the predicted display timing for the current frame.
*/
//go:nosplit
func (self class) GetPredictedDisplayTime() int64 { //gd:OpenXRAPIExtension.get_predicted_display_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_predicted_display_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the predicted display timing for the next frame.
*/
//go:nosplit
func (self class) GetNextFrameTime() int64 { //gd:OpenXRAPIExtension.get_next_frame_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_next_frame_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OpenXR is initialized for rendering with an XR viewport.
*/
//go:nosplit
func (self class) CanRender() bool { //gd:OpenXRAPIExtension.can_render
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_can_render, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [RID] corresponding to an [code]Action[/code] of a matching name, optionally limited to a specified action set.
*/
//go:nosplit
func (self class) FindAction(name String.Readable, action_set RID.Any) RID.Any { //gd:OpenXRAPIExtension.find_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, action_set)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_find_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the corresponding [code]XrAction[/code] OpenXR handle for the given action RID.
*/
//go:nosplit
func (self class) ActionGetHandle(action RID.Any) int64 { //gd:OpenXRAPIExtension.action_get_handle
	var frame = callframe.New()
	callframe.Arg(frame, action)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_action_get_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the corresponding [code]XRHandTrackerEXT[/code] handle for the given hand index value.
*/
//go:nosplit
func (self class) GetHandTracker(hand_index int64) int64 { //gd:OpenXRAPIExtension.get_hand_tracker
	var frame = callframe.New()
	callframe.Arg(frame, hand_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Registers the given extension as a composition layer provider.
*/
//go:nosplit
func (self class) RegisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.register_composition_layer_provider
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(extension[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_register_composition_layer_provider, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the given extension as a composition layer provider.
*/
//go:nosplit
func (self class) UnregisterCompositionLayerProvider(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.unregister_composition_layer_provider
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_unregister_composition_layer_provider, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers the given extension as a provider of additional data structures to projections views.
*/
//go:nosplit
func (self class) RegisterProjectionViewsExtension(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.register_projection_views_extension
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(extension[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_register_projection_views_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the given extension as a provider of additional data structures to projections views.
*/
//go:nosplit
func (self class) UnregisterProjectionViewsExtension(extension [1]gdclass.OpenXRExtensionWrapperExtension) { //gd:OpenXRAPIExtension.unregister_projection_views_extension
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(extension[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_unregister_projection_views_extension, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the near boundary value of the camera frustum.
[b]Note:[/b] This is only accessible in the render thread.
*/
//go:nosplit
func (self class) GetRenderStateZNear() float64 { //gd:OpenXRAPIExtension.get_render_state_z_near
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_render_state_z_near, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the far boundary value of the camera frustum.
[b]Note:[/b] This is only accessible in the render thread.
*/
//go:nosplit
func (self class) GetRenderStateZFar() float64 { //gd:OpenXRAPIExtension.get_render_state_z_far
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_render_state_z_far, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the render target of the velocity texture.
*/
//go:nosplit
func (self class) SetVelocityTexture(render_target RID.Any) { //gd:OpenXRAPIExtension.set_velocity_texture
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_velocity_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the render target of the velocity depth texture.
*/
//go:nosplit
func (self class) SetVelocityDepthTexture(render_target RID.Any) { //gd:OpenXRAPIExtension.set_velocity_depth_texture
	var frame = callframe.New()
	callframe.Arg(frame, render_target)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_velocity_depth_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the target size of the velocity and velocity depth textures.
*/
//go:nosplit
func (self class) SetVelocityTargetSize(target_size Vector2i.XY) { //gd:OpenXRAPIExtension.set_velocity_target_size
	var frame = callframe.New()
	callframe.Arg(frame, target_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_velocity_target_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of supported swapchain formats.
*/
//go:nosplit
func (self class) GetSupportedSwapchainFormats() Packed.Array[int64] { //gd:OpenXRAPIExtension.get_supported_swapchain_formats
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_supported_swapchain_formats, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int64](Array.Through(gd.PackedProxy[gd.PackedInt64Array, int64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a pointer to a new swapchain created using the provided parameters.
*/
//go:nosplit
func (self class) OpenxrSwapchainCreate(create_flags int64, usage_flags int64, swapchain_format int64, width int64, height int64, sample_count int64, array_size int64) int64 { //gd:OpenXRAPIExtension.openxr_swapchain_create
	var frame = callframe.New()
	callframe.Arg(frame, create_flags)
	callframe.Arg(frame, usage_flags)
	callframe.Arg(frame, swapchain_format)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, sample_count)
	callframe.Arg(frame, array_size)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_swapchain_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Destroys the provided swapchain and frees it from memory.
*/
//go:nosplit
func (self class) OpenxrSwapchainFree(swapchain int64) { //gd:OpenXRAPIExtension.openxr_swapchain_free
	var frame = callframe.New()
	callframe.Arg(frame, swapchain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_swapchain_free, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]XrSwapchain[/code] handle of the provided swapchain.
*/
//go:nosplit
func (self class) OpenxrSwapchainGetSwapchain(swapchain int64) int64 { //gd:OpenXRAPIExtension.openxr_swapchain_get_swapchain
	var frame = callframe.New()
	callframe.Arg(frame, swapchain)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_swapchain_get_swapchain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Acquires the image of the provided swapchain.
*/
//go:nosplit
func (self class) OpenxrSwapchainAcquire(swapchain int64) { //gd:OpenXRAPIExtension.openxr_swapchain_acquire
	var frame = callframe.New()
	callframe.Arg(frame, swapchain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_swapchain_acquire, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the RID of the provided swapchain's image.
*/
//go:nosplit
func (self class) OpenxrSwapchainGetImage(swapchain int64) RID.Any { //gd:OpenXRAPIExtension.openxr_swapchain_get_image
	var frame = callframe.New()
	callframe.Arg(frame, swapchain)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_swapchain_get_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Releases the image of the provided swapchain.
*/
//go:nosplit
func (self class) OpenxrSwapchainRelease(swapchain int64) { //gd:OpenXRAPIExtension.openxr_swapchain_release
	var frame = callframe.New()
	callframe.Arg(frame, swapchain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_openxr_swapchain_release, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a pointer to the render state's [code]XrCompositionLayerProjection[/code] struct.
[b]Note:[/b] This method should only be called from the rendering thread.
*/
//go:nosplit
func (self class) GetProjectionLayer() int64 { //gd:OpenXRAPIExtension.get_projection_layer
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_get_projection_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the render region to [param render_region], overriding the normal render target's rect.
*/
//go:nosplit
func (self class) SetRenderRegion(render_region Rect2i.PositionSize) { //gd:OpenXRAPIExtension.set_render_region
	var frame = callframe.New()
	callframe.Arg(frame, render_region)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_render_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If set to [code]true[/code], an OpenXR extension is loaded which is capable of emulating the [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] blend mode.
*/
//go:nosplit
func (self class) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool) { //gd:OpenXRAPIExtension.set_emulate_environment_blend_mode_alpha_blend
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_set_emulate_environment_blend_mode_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [enum OpenXRAPIExtension.OpenXRAlphaBlendModeSupport] denoting if [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported, emulated or not supported at all.
*/
//go:nosplit
func (self class) IsEnvironmentBlendModeAlphaSupported() gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport { //gd:OpenXRAPIExtension.is_environment_blend_mode_alpha_supported
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRAPIExtension.Bind_is_environment_blend_mode_alpha_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
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

type OpenXRAlphaBlendModeSupport = gdclass.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport //gd:OpenXRAPIExtension.OpenXRAlphaBlendModeSupport

const (
	/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] isn't supported at all.*/
	OpenxrAlphaBlendModeSupportNone OpenXRAlphaBlendModeSupport = 0
	/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported.*/
	OpenxrAlphaBlendModeSupportReal OpenXRAlphaBlendModeSupport = 1
	/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is emulated.*/
	OpenxrAlphaBlendModeSupportEmulating OpenXRAlphaBlendModeSupport = 2
)
