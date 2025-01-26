// Package NoiseTexture3D provides methods for working with NoiseTexture3D object instances.
package NoiseTexture3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/Texture3D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

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
type Instance [1]gdclass.NoiseTexture3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNoiseTexture3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NoiseTexture3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NoiseTexture3D"))
	casted := Instance{*(*gdclass.NoiseTexture3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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

func (self Instance) ColorRamp() [1]gdclass.Gradient {
	return [1]gdclass.Gradient(class(self).GetColorRamp())
}

func (self Instance) SetColorRamp(value [1]gdclass.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Instance) Noise() [1]gdclass.Noise {
	return [1]gdclass.Noise(class(self).GetNoise())
}

func (self Instance) SetNoise(value [1]gdclass.Noise) {
	class(self).SetNoise(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int) { //gd:NoiseTexture3D.set_width
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetHeight(height gd.Int) { //gd:NoiseTexture3D.set_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDepth(depth gd.Int) { //gd:NoiseTexture3D.set_depth
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetInvert(invert bool) { //gd:NoiseTexture3D.set_invert
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_invert, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInvert() bool { //gd:NoiseTexture3D.get_invert
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_invert, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSeamless(seamless bool) { //gd:NoiseTexture3D.set_seamless
	var frame = callframe.New()
	callframe.Arg(frame, seamless)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_seamless, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeamless() bool { //gd:NoiseTexture3D.get_seamless
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_seamless, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSeamlessBlendSkirt(seamless_blend_skirt gd.Float) { //gd:NoiseTexture3D.set_seamless_blend_skirt
	var frame = callframe.New()
	callframe.Arg(frame, seamless_blend_skirt)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeamlessBlendSkirt() gd.Float { //gd:NoiseTexture3D.get_seamless_blend_skirt
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalize(normalize bool) { //gd:NoiseTexture3D.set_normalize
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_normalize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsNormalized() bool { //gd:NoiseTexture3D.is_normalized
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_is_normalized, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorRamp(gradient [1]gdclass.Gradient) { //gd:NoiseTexture3D.set_color_ramp
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gradient[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRamp() [1]gdclass.Gradient { //gd:NoiseTexture3D.get_color_ramp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNoise(noise [1]gdclass.Noise) { //gd:NoiseTexture3D.set_noise
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(noise[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_set_noise, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNoise() [1]gdclass.Noise { //gd:NoiseTexture3D.get_noise
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture3D.Bind_get_noise, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Noise{gd.PointerWithOwnershipTransferredToGo[gdclass.Noise](r_ret.Get())}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture3D.Advanced(self.AsTexture3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture3D.Instance(self.AsTexture3D()), name)
	}
}
func init() {
	gdclass.Register("NoiseTexture3D", func(ptr gd.Object) any {
		return [1]gdclass.NoiseTexture3D{*(*gdclass.NoiseTexture3D)(unsafe.Pointer(&ptr))}
	})
}
