package CanvasItemMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Material"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[CanvasItemMaterial]s provide a means of modifying the textures associated with a CanvasItem. They specialize in describing blend and lighting behaviors for textures. Use a [ShaderMaterial] to more fully customize a material's interactions with a [CanvasItem].

*/
type Simple [1]classdb.CanvasItemMaterial
func (self Simple) SetBlendMode(blend_mode classdb.CanvasItemMaterialBlendMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBlendMode(blend_mode)
}
func (self Simple) GetBlendMode() classdb.CanvasItemMaterialBlendMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CanvasItemMaterialBlendMode(Expert(self).GetBlendMode())
}
func (self Simple) SetLightMode(light_mode classdb.CanvasItemMaterialLightMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLightMode(light_mode)
}
func (self Simple) GetLightMode() classdb.CanvasItemMaterialLightMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CanvasItemMaterialLightMode(Expert(self).GetLightMode())
}
func (self Simple) SetParticlesAnimation(particles_anim bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimation(particles_anim)
}
func (self Simple) GetParticlesAnimation() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetParticlesAnimation())
}
func (self Simple) SetParticlesAnimHFrames(frames int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimHFrames(gd.Int(frames))
}
func (self Simple) GetParticlesAnimHFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetParticlesAnimHFrames()))
}
func (self Simple) SetParticlesAnimVFrames(frames int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimVFrames(gd.Int(frames))
}
func (self Simple) GetParticlesAnimVFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetParticlesAnimVFrames()))
}
func (self Simple) SetParticlesAnimLoop(loop bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParticlesAnimLoop(loop)
}
func (self Simple) GetParticlesAnimLoop() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetParticlesAnimLoop())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CanvasItemMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetBlendMode(blend_mode classdb.CanvasItemMaterialBlendMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendMode() classdb.CanvasItemMaterialBlendMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemMaterialBlendMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLightMode(light_mode classdb.CanvasItemMaterialLightMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, light_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_set_light_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightMode() classdb.CanvasItemMaterialLightMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemMaterialLightMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_get_light_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimation(particles_anim bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, particles_anim)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_set_particles_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimation() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_get_particles_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimHFrames(frames gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_set_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimHFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_get_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimVFrames(frames gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_set_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimVFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_get_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimLoop(loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_set_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CanvasItemMaterial.Bind_get_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCanvasItemMaterial() Expert { return self[0].AsCanvasItemMaterial() }


//go:nosplit
func (self Simple) AsCanvasItemMaterial() Simple { return self[0].AsCanvasItemMaterial() }


//go:nosplit
func (self class) AsMaterial() Material.Expert { return self[0].AsMaterial() }


//go:nosplit
func (self Simple) AsMaterial() Material.Simple { return self[0].AsMaterial() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CanvasItemMaterial", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BlendMode = classdb.CanvasItemMaterialBlendMode

const (
/*Mix blending mode. Colors are assumed to be independent of the alpha (opacity) value.*/
	BlendModeMix BlendMode = 0
/*Additive blending mode.*/
	BlendModeAdd BlendMode = 1
/*Subtractive blending mode.*/
	BlendModeSub BlendMode = 2
/*Multiplicative blending mode.*/
	BlendModeMul BlendMode = 3
/*Mix blending mode. Colors are assumed to be premultiplied by the alpha (opacity) value.*/
	BlendModePremultAlpha BlendMode = 4
)
type LightMode = classdb.CanvasItemMaterialLightMode

const (
/*Render the material using both light and non-light sensitive material properties.*/
	LightModeNormal LightMode = 0
/*Render the material as if there were no light.*/
	LightModeUnshaded LightMode = 1
/*Render the material as if there were only light.*/
	LightModeLightOnly LightMode = 2
)
