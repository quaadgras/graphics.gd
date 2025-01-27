// Package CurveTexture provides methods for working with CurveTexture object instances.
package CurveTexture

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
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

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

/*
A 1D texture where pixel brightness corresponds to points on a [Curve] resource, either in grayscale or in red. This visual representation simplifies the task of saving curves as image files.
If you need to store up to 3 curves within a single texture, use [CurveXYZTexture] instead. See also [GradientTexture1D] and [GradientTexture2D].
*/
type Instance [1]gdclass.CurveTexture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCurveTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CurveTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CurveTexture"))
	casted := Instance{*(*gdclass.CurveTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SetWidth(value int) {
	class(self).SetWidth(gd.Int(value))
}

func (self Instance) TextureMode() gdclass.CurveTextureTextureMode {
	return gdclass.CurveTextureTextureMode(class(self).GetTextureMode())
}

func (self Instance) SetTextureMode(value gdclass.CurveTextureTextureMode) {
	class(self).SetTextureMode(value)
}

func (self Instance) Curve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetCurve())
}

func (self Instance) SetCurve(value [1]gdclass.Curve) {
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int) { //gd:CurveTexture.set_width
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve) { //gd:CurveTexture.set_curve
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve { //gd:CurveTexture.get_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMode(texture_mode gdclass.CurveTextureTextureMode) { //gd:CurveTexture.set_texture_mode
	var frame = callframe.New()
	callframe.Arg(frame, texture_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_set_texture_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMode() gdclass.CurveTextureTextureMode { //gd:CurveTexture.get_texture_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CurveTextureTextureMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_get_texture_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCurveTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCurveTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Texture2D.Advanced(self.AsTexture2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Instance(self.AsTexture2D()), name)
	}
}
func init() {
	gdclass.Register("CurveTexture", func(ptr gd.Object) any {
		return [1]gdclass.CurveTexture{*(*gdclass.CurveTexture)(unsafe.Pointer(&ptr))}
	})
}

type TextureMode = gdclass.CurveTextureTextureMode //gd:CurveTexture.TextureMode

const (
	/*Store the curve equally across the red, green and blue channels. This uses more video memory, but is more compatible with shaders that only read the green and blue values.*/
	TextureModeRgb TextureMode = 0
	/*Store the curve only in the red channel. This saves video memory, but some custom shaders may not be able to work with this.*/
	TextureModeRed TextureMode = 1
)
