package CanvasItemMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[CanvasItemMaterial]s provide a means of modifying the textures associated with a CanvasItem. They specialize in describing blend and lighting behaviors for textures. Use a [ShaderMaterial] to more fully customize a material's interactions with a [CanvasItem].

*/
type Go [1]classdb.CanvasItemMaterial
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CanvasItemMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasItemMaterial"))
	return Go{classdb.CanvasItemMaterial(object)}
}

func (self Go) BlendMode() classdb.CanvasItemMaterialBlendMode {
		return classdb.CanvasItemMaterialBlendMode(class(self).GetBlendMode())
}

func (self Go) SetBlendMode(value classdb.CanvasItemMaterialBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Go) LightMode() classdb.CanvasItemMaterialLightMode {
		return classdb.CanvasItemMaterialLightMode(class(self).GetLightMode())
}

func (self Go) SetLightMode(value classdb.CanvasItemMaterialLightMode) {
	class(self).SetLightMode(value)
}

func (self Go) ParticlesAnimation() bool {
		return bool(class(self).GetParticlesAnimation())
}

func (self Go) SetParticlesAnimation(value bool) {
	class(self).SetParticlesAnimation(value)
}

func (self Go) ParticlesAnimHFrames() int {
		return int(int(class(self).GetParticlesAnimHFrames()))
}

func (self Go) SetParticlesAnimHFrames(value int) {
	class(self).SetParticlesAnimHFrames(gd.Int(value))
}

func (self Go) ParticlesAnimVFrames() int {
		return int(int(class(self).GetParticlesAnimVFrames()))
}

func (self Go) SetParticlesAnimVFrames(value int) {
	class(self).SetParticlesAnimVFrames(gd.Int(value))
}

func (self Go) ParticlesAnimLoop() bool {
		return bool(class(self).GetParticlesAnimLoop())
}

func (self Go) SetParticlesAnimLoop(value bool) {
	class(self).SetParticlesAnimLoop(value)
}

//go:nosplit
func (self class) SetBlendMode(blend_mode classdb.CanvasItemMaterialBlendMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, blend_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBlendMode() classdb.CanvasItemMaterialBlendMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemMaterialBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLightMode(light_mode classdb.CanvasItemMaterialLightMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, light_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_light_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightMode() classdb.CanvasItemMaterialLightMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CanvasItemMaterialLightMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_light_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimation(particles_anim bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, particles_anim)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimation() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimHFrames(frames gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimHFrames() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_anim_h_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimVFrames(frames gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimVFrames() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_anim_v_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParticlesAnimLoop(loop bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_set_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParticlesAnimLoop() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasItemMaterial.Bind_get_particles_anim_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCanvasItemMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItemMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.GD { return *((*Material.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMaterial() Material.Go { return *((*Material.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {classdb.Register("CanvasItemMaterial", func(ptr gd.Object) any { return classdb.CanvasItemMaterial(ptr) })}
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
