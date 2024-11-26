package CompositorEffect

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This resource defines a custom rendering effect that can be applied to [Viewport]s through the viewports' [Environment]. You can implement a callback that is called during rendering at a given stage of the rendering pipeline and allows you to insert additional passes. Note that this callback happens on the rendering thread. CompositorEffect is an abstract base class and must be extended to implement specific rendering logic.

	// CompositorEffect methods that can be overridden by a [Class] that extends it.
	type CompositorEffect interface {
		//Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
		RenderCallback(effect_callback_type int, render_data objects.RenderData)
	}
*/
type Instance [1]classdb.CompositorEffect

/*
Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
*/
func (Instance) _render_callback(impl func(ptr unsafe.Pointer, effect_callback_type int, render_data objects.RenderData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var effect_callback_type = gd.UnsafeGet[gd.Int](p_args, 0)
		var render_data = objects.RenderData{pointers.New[classdb.RenderData]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(render_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(effect_callback_type), render_data)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CompositorEffect

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CompositorEffect"))
	return Instance{classdb.CompositorEffect(object)}
}

func (self Instance) Enabled() bool {
	return bool(class(self).GetEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) EffectCallbackType() classdb.CompositorEffectEffectCallbackType {
	return classdb.CompositorEffectEffectCallbackType(class(self).GetEffectCallbackType())
}

func (self Instance) SetEffectCallbackType(value classdb.CompositorEffectEffectCallbackType) {
	class(self).SetEffectCallbackType(value)
}

func (self Instance) AccessResolvedColor() bool {
	return bool(class(self).GetAccessResolvedColor())
}

func (self Instance) SetAccessResolvedColor(value bool) {
	class(self).SetAccessResolvedColor(value)
}

func (self Instance) AccessResolvedDepth() bool {
	return bool(class(self).GetAccessResolvedDepth())
}

func (self Instance) SetAccessResolvedDepth(value bool) {
	class(self).SetAccessResolvedDepth(value)
}

func (self Instance) NeedsMotionVectors() bool {
	return bool(class(self).GetNeedsMotionVectors())
}

func (self Instance) SetNeedsMotionVectors(value bool) {
	class(self).SetNeedsMotionVectors(value)
}

func (self Instance) NeedsNormalRoughness() bool {
	return bool(class(self).GetNeedsNormalRoughness())
}

func (self Instance) SetNeedsNormalRoughness(value bool) {
	class(self).SetNeedsNormalRoughness(value)
}

func (self Instance) NeedsSeparateSpecular() bool {
	return bool(class(self).GetNeedsSeparateSpecular())
}

func (self Instance) SetNeedsSeparateSpecular(value bool) {
	class(self).SetNeedsSeparateSpecular(value)
}

/*
Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
*/
func (class) _render_callback(impl func(ptr unsafe.Pointer, effect_callback_type gd.Int, render_data objects.RenderData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var effect_callback_type = gd.UnsafeGet[gd.Int](p_args, 0)
		var render_data = objects.RenderData{pointers.New[classdb.RenderData]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(render_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, effect_callback_type, render_data)
	}
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEffectCallbackType(effect_callback_type classdb.CompositorEffectEffectCallbackType) {
	var frame = callframe.New()
	callframe.Arg(frame, effect_callback_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_effect_callback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEffectCallbackType() classdb.CompositorEffectEffectCallbackType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CompositorEffectEffectCallbackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_effect_callback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessResolvedColor(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_access_resolved_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessResolvedColor() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_access_resolved_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessResolvedDepth(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_access_resolved_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessResolvedDepth() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_access_resolved_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeedsMotionVectors(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_needs_motion_vectors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeedsMotionVectors() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_needs_motion_vectors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeedsNormalRoughness(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_needs_normal_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeedsNormalRoughness() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_needs_normal_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeedsSeparateSpecular(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_needs_separate_specular, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeedsSeparateSpecular() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_needs_separate_specular, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCompositorEffect() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCompositorEffect() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_render_callback":
		return reflect.ValueOf(self._render_callback)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_render_callback":
		return reflect.ValueOf(self._render_callback)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("CompositorEffect", func(ptr gd.Object) any { return classdb.CompositorEffect(ptr) })
}

type EffectCallbackType = classdb.CompositorEffectEffectCallbackType

const (
	/*The callback is called before our opaque rendering pass, but after depth prepass (if applicable).*/
	EffectCallbackTypePreOpaque EffectCallbackType = 0
	/*The callback is called after our opaque rendering pass, but before our sky is rendered.*/
	EffectCallbackTypePostOpaque EffectCallbackType = 1
	/*The callback is called after our sky is rendered, but before our back buffers are created (and if enabled, before subsurface scattering and/or screen space reflections).*/
	EffectCallbackTypePostSky EffectCallbackType = 2
	/*The callback is called before our transparent rendering pass, but after our sky is rendered and we've created our back buffers.*/
	EffectCallbackTypePreTransparent EffectCallbackType = 3
	/*The callback is called after our transparent rendering pass, but before any build in post effects and output to our render target.*/
	EffectCallbackTypePostTransparent EffectCallbackType = 4
	/*Represents the size of the [enum EffectCallbackType] enum.*/
	EffectCallbackTypeMax EffectCallbackType = 5
)
