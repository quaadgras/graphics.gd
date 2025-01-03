package CanvasItemMaterial

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Material"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[CanvasItemMaterial]s provide a means of modifying the textures associated with a CanvasItem. They specialize in describing blend and lighting behaviors for textures. Use a [ShaderMaterial] to more fully customize a material's interactions with a [CanvasItem].
*/
type Instance [1]classdb.CanvasItemMaterial
type Any interface {
	gd.IsClass
	AsCanvasItemMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CanvasItemMaterial

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasItemMaterial"))
	return Instance{*(*classdb.CanvasItemMaterial)(unsafe.Pointer(&object))}
}

func (self Instance) BlendMode() classdb.CanvasItemMaterialBlendMode {
	return classdb.CanvasItemMaterialBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value classdb.CanvasItemMaterialBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Instance) LightMode() classdb.CanvasItemMaterialLightMode {
	return classdb.CanvasItemMaterialLightMode(class(self).GetLightMode())
}

func (self Instance) SetLightMode(value classdb.CanvasItemMaterialLightMode) {
	class(self).SetLightMode(value)
}

func (self Instance) ParticlesAnimation() bool {
	return bool(class(self).GetParticlesAnimation())
}

func (self Instance) SetParticlesAnimation(value bool) {
	class(self).SetParticlesAnimation(value)
}

func (self Instance) ParticlesAnimHFrames() int {
	return int(int(class(self).GetParticlesAnimHFrames()))
}

func (self Instance) SetParticlesAnimHFrames(value int) {
	class(self).SetParticlesAnimHFrames(gd.Int(value))
}

func (self Instance) ParticlesAnimVFrames() int {
	return int(int(class(self).GetParticlesAnimVFrames()))
}

func (self Instance) SetParticlesAnimVFrames(value int) {
	class(self).SetParticlesAnimVFrames(gd.Int(value))
}

func (self Instance) ParticlesAnimLoop() bool {
	return bool(class(self).GetParticlesAnimLoop())
}

func (self Instance) SetParticlesAnimLoop(value bool) {
	class(self).SetParticlesAnimLoop(value)
}

//go:nosplit
func (self class) SetBlendMode(blend_mode classdb.CanvasItemMaterialBlendMode) {
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
func (self class) SetLightMode(light_mode classdb.CanvasItemMaterialLightMode) {
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
func (self class) SetParticlesAnimation(particles_anim bool) {
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
func (self class) SetParticlesAnimHFrames(frames gd.Int) {
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
func (self class) SetParticlesAnimVFrames(frames gd.Int) {
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
func (self class) SetParticlesAnimLoop(loop bool) {
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
func (self class) AsCanvasItemMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasItemMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
}
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
	default:
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {
	classdb.Register("CanvasItemMaterial", func(ptr gd.Object) any {
		return [1]classdb.CanvasItemMaterial{*(*classdb.CanvasItemMaterial)(unsafe.Pointer(&ptr))}
	})
}

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
