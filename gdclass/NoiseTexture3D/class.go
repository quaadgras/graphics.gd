package NoiseTexture3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture3D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Uses the [FastNoiseLite] library or other noise generators to fill the texture data of your desired size.
The class uses [Thread]s to generate the texture data internally, so [method Texture3D.get_data] may return [code]null[/code] if the generation process has not completed yet. In that case, you need to wait for the texture to be generated before accessing the image:
[codeblock]
var texture = NoiseTexture3D.new()
texture.noise = FastNoiseLite.new()
await texture.changed
var data = texture.get_data()
[/codeblock]

*/
type Go [1]classdb.NoiseTexture3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NoiseTexture3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NoiseTexture3D"))
	return Go{classdb.NoiseTexture3D(object)}
}

func (self Go) Invert() bool {
		return bool(class(self).GetInvert())
}

func (self Go) SetInvert(value bool) {
	class(self).SetInvert(value)
}

func (self Go) Seamless() bool {
		return bool(class(self).GetSeamless())
}

func (self Go) SetSeamless(value bool) {
	class(self).SetSeamless(value)
}

func (self Go) SeamlessBlendSkirt() float64 {
		return float64(float64(class(self).GetSeamlessBlendSkirt()))
}

func (self Go) SetSeamlessBlendSkirt(value float64) {
	class(self).SetSeamlessBlendSkirt(gd.Float(value))
}

func (self Go) Normalize() bool {
		return bool(class(self).IsNormalized())
}

func (self Go) SetNormalize(value bool) {
	class(self).SetNormalize(value)
}

func (self Go) ColorRamp() gdclass.Gradient {
		return gdclass.Gradient(class(self).GetColorRamp())
}

func (self Go) SetColorRamp(value gdclass.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Go) Noise() gdclass.Noise {
		return gdclass.Noise(class(self).GetNoise())
}

func (self Go) SetNoise(value gdclass.Noise) {
	class(self).SetNoise(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHeight(height gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDepth(depth gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInvert(invert bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInvert() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSeamless(seamless bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, seamless)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_seamless, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeamless() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_seamless, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSeamlessBlendSkirt(seamless_blend_skirt gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, seamless_blend_skirt)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeamlessBlendSkirt() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNormalize(normalize bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_normalize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsNormalized() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_is_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorRamp(gradient gdclass.Gradient)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(gradient[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorRamp() gdclass.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNoise(noise gdclass.Noise)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(noise[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_noise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNoise() gdclass.Noise {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_noise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Noise{classdb.Noise(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsNoiseTexture3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNoiseTexture3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.GD { return *((*Texture3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture3D() Texture3D.Go { return *((*Texture3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture3D(), name)
	}
}
func init() {classdb.Register("NoiseTexture3D", func(ptr gd.Object) any { return classdb.NoiseTexture3D(ptr) })}
