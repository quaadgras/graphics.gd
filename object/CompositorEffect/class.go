package CompositorEffect

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource defines a custom rendering effect that can be applied to [Viewport]s through the viewports' [Environment]. You can implement a callback that is called during rendering at a given stage of the rendering pipeline and allows you to insert additional passes. Note that this callback happens on the rendering thread. CompositorEffect is an abstract base class and must be extended to implement specific rendering logic.
	// CompositorEffect methods that can be overridden by a [Class] that extends it.
	type CompositorEffect interface {
		//Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
		RenderCallback(effect_callback_type gd.Int, render_data object.RenderData) 
	}

*/
type Simple [1]classdb.CompositorEffect
func (Simple) _render_callback(impl func(ptr unsafe.Pointer, effect_callback_type int, render_data [1]classdb.RenderData) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var effect_callback_type = gd.UnsafeGet[gd.Int](p_args,0)
		var render_data [1]classdb.RenderData
		render_data[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(effect_callback_type), render_data)
		gc.End()
	}
}
func (self Simple) SetEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnabled(enabled)
}
func (self Simple) GetEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetEnabled())
}
func (self Simple) SetEffectCallbackType(effect_callback_type classdb.CompositorEffectEffectCallbackType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEffectCallbackType(effect_callback_type)
}
func (self Simple) GetEffectCallbackType() classdb.CompositorEffectEffectCallbackType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CompositorEffectEffectCallbackType(Expert(self).GetEffectCallbackType())
}
func (self Simple) SetAccessResolvedColor(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAccessResolvedColor(enable)
}
func (self Simple) GetAccessResolvedColor() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAccessResolvedColor())
}
func (self Simple) SetAccessResolvedDepth(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAccessResolvedDepth(enable)
}
func (self Simple) GetAccessResolvedDepth() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAccessResolvedDepth())
}
func (self Simple) SetNeedsMotionVectors(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNeedsMotionVectors(enable)
}
func (self Simple) GetNeedsMotionVectors() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNeedsMotionVectors())
}
func (self Simple) SetNeedsNormalRoughness(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNeedsNormalRoughness(enable)
}
func (self Simple) GetNeedsNormalRoughness() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNeedsNormalRoughness())
}
func (self Simple) SetNeedsSeparateSpecular(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNeedsSeparateSpecular(enable)
}
func (self Simple) GetNeedsSeparateSpecular() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNeedsSeparateSpecular())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CompositorEffect
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Implement this function with your custom rendering code. [param effect_callback_type] should always match the effect callback type you've specified in [member effect_callback_type]. [param render_data] provides access to the rendering state, it is only valid during rendering and should not be stored.
*/
func (class) _render_callback(impl func(ptr unsafe.Pointer, effect_callback_type gd.Int, render_data object.RenderData) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var effect_callback_type = gd.UnsafeGet[gd.Int](p_args,0)
		var render_data object.RenderData
		render_data[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,1)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, effect_callback_type, render_data)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEffectCallbackType(effect_callback_type classdb.CompositorEffectEffectCallbackType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, effect_callback_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_effect_callback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEffectCallbackType() classdb.CompositorEffectEffectCallbackType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CompositorEffectEffectCallbackType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_effect_callback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAccessResolvedColor(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_access_resolved_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccessResolvedColor() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_access_resolved_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAccessResolvedDepth(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_access_resolved_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccessResolvedDepth() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_access_resolved_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNeedsMotionVectors(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_needs_motion_vectors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNeedsMotionVectors() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_needs_motion_vectors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNeedsNormalRoughness(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_needs_normal_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNeedsNormalRoughness() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_needs_normal_roughness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNeedsSeparateSpecular(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_set_needs_separate_specular, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNeedsSeparateSpecular() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CompositorEffect.Bind_get_needs_separate_specular, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCompositorEffect() Expert { return self[0].AsCompositorEffect() }


//go:nosplit
func (self Simple) AsCompositorEffect() Simple { return self[0].AsCompositorEffect() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_render_callback": return reflect.ValueOf(self._render_callback);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_render_callback": return reflect.ValueOf(self._render_callback);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CompositorEffect", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
