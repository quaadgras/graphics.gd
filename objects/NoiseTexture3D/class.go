package NoiseTexture3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture3D"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
type Instance [1]classdb.NoiseTexture3D
type Any interface {
	gd.IsClass
	AsNoiseTexture3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.NoiseTexture3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NoiseTexture3D"))
	return Instance{*(*classdb.NoiseTexture3D)(unsafe.Pointer(&object))}
}

func (self Instance) SetWidth(value int) {
	class(self).SetWidth(gd.Int(value))
}

func (self Instance) SetHeight(value int) {
	class(self).SetHeight(gd.Int(value))
}

func (self Instance) SetDepth(value int) {
	class(self).SetDepth(gd.Int(value))
}

func (self Instance) Invert() bool {
	return bool(class(self).GetInvert())
}

func (self Instance) SetInvert(value bool) {
	class(self).SetInvert(value)
}

func (self Instance) Seamless() bool {
	return bool(class(self).GetSeamless())
}

func (self Instance) SetSeamless(value bool) {
	class(self).SetSeamless(value)
}

func (self Instance) SeamlessBlendSkirt() Float.X {
	return Float.X(Float.X(class(self).GetSeamlessBlendSkirt()))
}

func (self Instance) SetSeamlessBlendSkirt(value Float.X) {
	class(self).SetSeamlessBlendSkirt(gd.Float(value))
}

func (self Instance) Normalize() bool {
	return bool(class(self).IsNormalized())
}

func (self Instance) SetNormalize(value bool) {
	class(self).SetNormalize(value)
}

func (self Instance) ColorRamp() objects.Gradient {
	return objects.Gradient(class(self).GetColorRamp())
}

func (self Instance) SetColorRamp(value objects.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Instance) Noise() objects.Noise {
	return objects.Noise(class(self).GetNoise())
}

func (self Instance) SetNoise(value objects.Noise) {
	class(self).SetNoise(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetHeight(height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDepth(depth gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetInvert(invert bool) {
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
func (self class) SetSeamless(seamless bool) {
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
func (self class) SetSeamlessBlendSkirt(seamless_blend_skirt gd.Float) {
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
func (self class) SetNormalize(normalize bool) {
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
func (self class) SetColorRamp(gradient objects.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gradient[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRamp() objects.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Gradient{gd.PointerWithOwnershipTransferredToGo[classdb.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNoise(noise objects.Noise) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(noise[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_noise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNoise() objects.Noise {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_noise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Noise{gd.PointerWithOwnershipTransferredToGo[classdb.Noise](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsNoiseTexture3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNoiseTexture3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.Advanced {
	return *((*Texture3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture3D() Texture3D.Instance {
	return *((*Texture3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsTexture3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture3D(), name)
	}
}
func init() {
	classdb.Register("NoiseTexture3D", func(ptr gd.Object) any {
		return [1]classdb.NoiseTexture3D{*(*classdb.NoiseTexture3D)(unsafe.Pointer(&ptr))}
	})
}
