// Package CompositorEffect provides methods for working with CompositorEffect object instances.
package CompositorEffect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This resource defines a custom rendering effect that can be applied to [Viewport]s through the viewports' [Environment]. You can implement a callback that is called during rendering at a given stage of the rendering pipeline and allows you to insert additional passes. Note that this callback happens on the rendering thread. CompositorEffect is an abstract base class and must be extended to implement specific rendering logic.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=CompositorEffect)
*/
type Instance [1]gdclass.CompositorEffect

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCompositorEffect() Instance
}
type Interface interface {
	//Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
	RenderCallback(effect_callback_type int, render_data [1]gdclass.RenderData)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) RenderCallback(effect_callback_type int, render_data [1]gdclass.RenderData) {
	return
}

/*
Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
*/
func (Instance) _render_callback(impl func(ptr unsafe.Pointer, effect_callback_type int, render_data [1]gdclass.RenderData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var effect_callback_type = gd.UnsafeGet[gd.Int](p_args, 0)
		var render_data = [1]gdclass.RenderData{pointers.New[gdclass.RenderData]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 1))})}
		defer pointers.End(render_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(effect_callback_type), render_data)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CompositorEffect

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CompositorEffect"))
	casted := Instance{*(*gdclass.CompositorEffect)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Enabled() bool {
	return bool(class(self).GetEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) EffectCallbackType() gdclass.CompositorEffectEffectCallbackType {
	return gdclass.CompositorEffectEffectCallbackType(class(self).GetEffectCallbackType())
}

func (self Instance) SetEffectCallbackType(value gdclass.CompositorEffectEffectCallbackType) {
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
func (class) _render_callback(impl func(ptr unsafe.Pointer, effect_callback_type gd.Int, render_data [1]gdclass.RenderData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var effect_callback_type = gd.UnsafeGet[gd.Int](p_args, 0)
		var render_data = [1]gdclass.RenderData{pointers.New[gdclass.RenderData]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(render_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, effect_callback_type, render_data)
	}
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEffectCallbackType(effect_callback_type gdclass.CompositorEffectEffectCallbackType) {
	var frame = callframe.New()
	callframe.Arg(frame, effect_callback_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_effect_callback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEffectCallbackType() gdclass.CompositorEffectEffectCallbackType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CompositorEffectEffectCallbackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_effect_callback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessResolvedColor(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_access_resolved_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessResolvedColor() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_access_resolved_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessResolvedDepth(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_access_resolved_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessResolvedDepth() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_access_resolved_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeedsMotionVectors(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_needs_motion_vectors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeedsMotionVectors() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_needs_motion_vectors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeedsNormalRoughness(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_needs_normal_roughness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeedsNormalRoughness() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_needs_normal_roughness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeedsSeparateSpecular(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_set_needs_separate_specular, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeedsSeparateSpecular() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CompositorEffect.Bind_get_needs_separate_specular, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_render_callback":
		return reflect.ValueOf(self._render_callback)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_render_callback":
		return reflect.ValueOf(self._render_callback)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("CompositorEffect", func(ptr gd.Object) any {
		return [1]gdclass.CompositorEffect{*(*gdclass.CompositorEffect)(unsafe.Pointer(&ptr))}
	})
}

type EffectCallbackType = gdclass.CompositorEffectEffectCallbackType

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
