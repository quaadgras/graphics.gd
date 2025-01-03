package CurveTexture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture2D"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A 1D texture where pixel brightness corresponds to points on a [Curve] resource, either in grayscale or in red. This visual representation simplifies the task of saving curves as image files.
If you need to store up to 3 curves within a single texture, use [CurveXYZTexture] instead. See also [GradientTexture1D] and [GradientTexture2D].
*/
type Instance [1]classdb.CurveTexture
type Any interface {
	gd.IsClass
	AsCurveTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CurveTexture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CurveTexture"))
	return Instance{*(*classdb.CurveTexture)(unsafe.Pointer(&object))}
}

func (self Instance) SetWidth(value int) {
	class(self).SetWidth(gd.Int(value))
}

func (self Instance) TextureMode() classdb.CurveTextureTextureMode {
	return classdb.CurveTextureTextureMode(class(self).GetTextureMode())
}

func (self Instance) SetTextureMode(value classdb.CurveTextureTextureMode) {
	class(self).SetTextureMode(value)
}

func (self Instance) Curve() objects.Curve {
	return objects.Curve(class(self).GetCurve())
}

func (self Instance) SetCurve(value objects.Curve) {
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurve(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{gd.PointerWithOwnershipTransferredToGo[classdb.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMode(texture_mode classdb.CurveTextureTextureMode) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_set_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMode() classdb.CurveTextureTextureMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CurveTextureTextureMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CurveTexture.Bind_get_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("CurveTexture", func(ptr gd.Object) any {
		return [1]classdb.CurveTexture{*(*classdb.CurveTexture)(unsafe.Pointer(&ptr))}
	})
}

type TextureMode = classdb.CurveTextureTextureMode

const (
	/*Store the curve equally across the red, green and blue channels. This uses more video memory, but is more compatible with shaders that only read the green and blue values.*/
	TextureModeRgb TextureMode = 0
	/*Store the curve only in the red channel. This saves video memory, but some custom shaders may not be able to work with this.*/
	TextureModeRed TextureMode = 1
)
