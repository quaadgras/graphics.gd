// Package GradientTexture2D provides methods for working with GradientTexture2D object instances.
package GradientTexture2D

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
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
A 2D texture that obtains colors from a [Gradient] to fill the texture data. This texture is able to transform a color transition into different patterns such as a linear or a radial gradient. The gradient is sampled individually for each pixel so it does not necessarily represent an exact copy of the gradient(see [member width] and [member height]). See also [GradientTexture1D], [CurveTexture] and [CurveXYZTexture].
*/
type Instance [1]gdclass.GradientTexture2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGradientTexture2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GradientTexture2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GradientTexture2D"))
	casted := Instance{*(*gdclass.GradientTexture2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Gradient() [1]gdclass.Gradient {
	return [1]gdclass.Gradient(class(self).GetGradient())
}

func (self Instance) SetGradient(value [1]gdclass.Gradient) {
	class(self).SetGradient(value)
}

func (self Instance) SetWidth(value int) {
	class(self).SetWidth(gd.Int(value))
}

func (self Instance) SetHeight(value int) {
	class(self).SetHeight(gd.Int(value))
}

func (self Instance) UseHdr() bool {
	return bool(class(self).IsUsingHdr())
}

func (self Instance) SetUseHdr(value bool) {
	class(self).SetUseHdr(value)
}

func (self Instance) Fill() gdclass.GradientTexture2DFill {
	return gdclass.GradientTexture2DFill(class(self).GetFill())
}

func (self Instance) SetFill(value gdclass.GradientTexture2DFill) {
	class(self).SetFill(value)
}

func (self Instance) FillFrom() Vector2.XY {
	return Vector2.XY(class(self).GetFillFrom())
}

func (self Instance) SetFillFrom(value Vector2.XY) {
	class(self).SetFillFrom(gd.Vector2(value))
}

func (self Instance) FillTo() Vector2.XY {
	return Vector2.XY(class(self).GetFillTo())
}

func (self Instance) SetFillTo(value Vector2.XY) {
	class(self).SetFillTo(gd.Vector2(value))
}

func (self Instance) Repeat() gdclass.GradientTexture2DRepeat {
	return gdclass.GradientTexture2DRepeat(class(self).GetRepeat())
}

func (self Instance) SetRepeat(value gdclass.GradientTexture2DRepeat) {
	class(self).SetRepeat(value)
}

//go:nosplit
func (self class) SetGradient(gradient [1]gdclass.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gradient[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGradient() [1]gdclass.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetHeight(height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseHdr(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_use_hdr, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingHdr() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_is_using_hdr, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFill(fill gdclass.GradientTexture2DFill) {
	var frame = callframe.New()
	callframe.Arg(frame, fill)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFill() gdclass.GradientTexture2DFill {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GradientTexture2DFill](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFillFrom(fill_from gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, fill_from)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillFrom() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFillTo(fill_to gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, fill_to)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillTo() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeat(repeat gdclass.GradientTexture2DRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeat() gdclass.GradientTexture2DRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GradientTexture2DRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGradientTexture2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGradientTexture2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GradientTexture2D", func(ptr gd.Object) any {
		return [1]gdclass.GradientTexture2D{*(*gdclass.GradientTexture2D)(unsafe.Pointer(&ptr))}
	})
}

type Fill = gdclass.GradientTexture2DFill

const (
	/*The colors are linearly interpolated in a straight line.*/
	FillLinear Fill = 0
	/*The colors are linearly interpolated in a circular pattern.*/
	FillRadial Fill = 1
	/*The colors are linearly interpolated in a square pattern.*/
	FillSquare Fill = 2
)

type Repeat = gdclass.GradientTexture2DRepeat

const (
	/*The gradient fill is restricted to the range defined by [member fill_from] to [member fill_to] offsets.*/
	RepeatNone Repeat = 0
	/*The texture is filled starting from [member fill_from] to [member fill_to] offsets, repeating the same pattern in both directions.*/
	RepeatDefault Repeat = 1
	/*The texture is filled starting from [member fill_from] to [member fill_to] offsets, mirroring the pattern in both directions.*/
	RepeatMirror Repeat = 2
)
