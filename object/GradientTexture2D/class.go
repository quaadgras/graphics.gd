package GradientTexture2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture2D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 2D texture that obtains colors from a [Gradient] to fill the texture data. This texture is able to transform a color transition into different patterns such as a linear or a radial gradient. The gradient is sampled individually for each pixel so it does not necessarily represent an exact copy of the gradient(see [member width] and [member height]). See also [GradientTexture1D], [CurveTexture] and [CurveXYZTexture].

*/
type Simple [1]classdb.GradientTexture2D
func (self Simple) SetGradient(gradient [1]classdb.Gradient) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGradient(gradient)
}
func (self Simple) GetGradient() [1]classdb.Gradient {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Gradient(Expert(self).GetGradient(gc))
}
func (self Simple) SetWidth(width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWidth(gd.Int(width))
}
func (self Simple) SetHeight(height int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeight(gd.Int(height))
}
func (self Simple) SetUseHdr(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseHdr(enabled)
}
func (self Simple) IsUsingHdr() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingHdr())
}
func (self Simple) SetFill(fill classdb.GradientTexture2DFill) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFill(fill)
}
func (self Simple) GetFill() classdb.GradientTexture2DFill {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GradientTexture2DFill(Expert(self).GetFill())
}
func (self Simple) SetFillFrom(fill_from gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFillFrom(fill_from)
}
func (self Simple) GetFillFrom() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetFillFrom())
}
func (self Simple) SetFillTo(fill_to gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFillTo(fill_to)
}
func (self Simple) GetFillTo() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetFillTo())
}
func (self Simple) SetRepeat(repeat classdb.GradientTexture2DRepeat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRepeat(repeat)
}
func (self Simple) GetRepeat() classdb.GradientTexture2DRepeat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GradientTexture2DRepeat(Expert(self).GetRepeat())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GradientTexture2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetGradient(gradient object.Gradient)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(gradient[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGradient(ctx gd.Lifetime) object.Gradient {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Gradient
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHeight(height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseHdr(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_use_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingHdr() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_is_using_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFill(fill classdb.GradientTexture2DFill)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fill)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFill() classdb.GradientTexture2DFill {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientTexture2DFill](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_get_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFillFrom(fill_from gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fill_from)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_fill_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillFrom() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_get_fill_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFillTo(fill_to gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fill_to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_fill_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillTo() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_get_fill_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRepeat(repeat classdb.GradientTexture2DRepeat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_set_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRepeat() classdb.GradientTexture2DRepeat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientTexture2DRepeat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GradientTexture2D.Bind_get_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGradientTexture2D() Expert { return self[0].AsGradientTexture2D() }


//go:nosplit
func (self Simple) AsGradientTexture2D() Simple { return self[0].AsGradientTexture2D() }


//go:nosplit
func (self class) AsTexture2D() Texture2D.Expert { return self[0].AsTexture2D() }


//go:nosplit
func (self Simple) AsTexture2D() Texture2D.Simple { return self[0].AsTexture2D() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


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
func init() {classdb.Register("GradientTexture2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
