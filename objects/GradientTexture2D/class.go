package GradientTexture2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Texture2D"
import "grow.graphics/gd/objects/Texture"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A 2D texture that obtains colors from a [Gradient] to fill the texture data. This texture is able to transform a color transition into different patterns such as a linear or a radial gradient. The gradient is sampled individually for each pixel so it does not necessarily represent an exact copy of the gradient(see [member width] and [member height]). See also [GradientTexture1D], [CurveTexture] and [CurveXYZTexture].
*/
type Instance [1]classdb.GradientTexture2D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GradientTexture2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GradientTexture2D"))
	return Instance{classdb.GradientTexture2D(object)}
}

func (self Instance) Gradient() objects.Gradient {
	return objects.Gradient(class(self).GetGradient())
}

func (self Instance) SetGradient(value objects.Gradient) {
	class(self).SetGradient(value)
}

func (self Instance) UseHdr() bool {
	return bool(class(self).IsUsingHdr())
}

func (self Instance) SetUseHdr(value bool) {
	class(self).SetUseHdr(value)
}

func (self Instance) Fill() classdb.GradientTexture2DFill {
	return classdb.GradientTexture2DFill(class(self).GetFill())
}

func (self Instance) SetFill(value classdb.GradientTexture2DFill) {
	class(self).SetFill(value)
}

func (self Instance) FillFrom() gd.Vector2 {
	return gd.Vector2(class(self).GetFillFrom())
}

func (self Instance) SetFillFrom(value gd.Vector2) {
	class(self).SetFillFrom(value)
}

func (self Instance) FillTo() gd.Vector2 {
	return gd.Vector2(class(self).GetFillTo())
}

func (self Instance) SetFillTo(value gd.Vector2) {
	class(self).SetFillTo(value)
}

func (self Instance) Repeat() classdb.GradientTexture2DRepeat {
	return classdb.GradientTexture2DRepeat(class(self).GetRepeat())
}

func (self Instance) SetRepeat(value classdb.GradientTexture2DRepeat) {
	class(self).SetRepeat(value)
}

//go:nosplit
func (self class) SetGradient(gradient objects.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gradient[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGradient() objects.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetHeight(height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseHdr(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_use_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingHdr() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_is_using_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFill(fill classdb.GradientTexture2DFill) {
	var frame = callframe.New()
	callframe.Arg(frame, fill)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFill() classdb.GradientTexture2DFill {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientTexture2DFill](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFillFrom(fill_from gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, fill_from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillFrom() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFillTo(fill_to gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, fill_to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillTo() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRepeat(repeat classdb.GradientTexture2DRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRepeat() classdb.GradientTexture2DRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientTexture2DRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	classdb.Register("GradientTexture2D", func(ptr gd.Object) any { return classdb.GradientTexture2D(ptr) })
}

type Fill = classdb.GradientTexture2DFill

const (
	/*The colors are linearly interpolated in a straight line.*/
	FillLinear Fill = 0
	/*The colors are linearly interpolated in a circular pattern.*/
	FillRadial Fill = 1
	/*The colors are linearly interpolated in a square pattern.*/
	FillSquare Fill = 2
)

type Repeat = classdb.GradientTexture2DRepeat

const (
	/*The gradient fill is restricted to the range defined by [member fill_from] to [member fill_to] offsets.*/
	RepeatNone Repeat = 0
	/*The texture is filled starting from [member fill_from] to [member fill_to] offsets, repeating the same pattern in both directions.*/
	RepeatDefault Repeat = 1
	/*The texture is filled starting from [member fill_from] to [member fill_to] offsets, mirroring the pattern in both directions.*/
	RepeatMirror Repeat = 2
)
