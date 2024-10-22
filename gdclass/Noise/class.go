package Noise

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class defines the interface for noise generation libraries to inherit from.
A default [method get_seamless_image] implementation is provided for libraries that do not provide seamless noise. This function requests a larger image from the [method get_image] method, reverses the quadrants of the image, then uses the strips of extra width to blend over the seams.
Inheriting noise classes can optionally override this function to provide a more optimal algorithm.

*/
type Go [1]classdb.Noise

/*
Returns the 1D noise value at the given (x) coordinate.
*/
func (self Go) GetNoise1d(x float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetNoise1d(gd.Float(x))))
}

/*
Returns the 2D noise value at the given position.
*/
func (self Go) GetNoise2d(x float64, y float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetNoise2d(gd.Float(x), gd.Float(y))))
}

/*
Returns the 2D noise value at the given position.
*/
func (self Go) GetNoise2dv(v gd.Vector2) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetNoise2dv(v)))
}

/*
Returns the 3D noise value at the given position.
*/
func (self Go) GetNoise3d(x float64, y float64, z float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetNoise3d(gd.Float(x), gd.Float(y), gd.Float(z))))
}

/*
Returns the 3D noise value at the given position.
*/
func (self Go) GetNoise3dv(v gd.Vector3) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetNoise3dv(v)))
}

/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Go) GetImage(width int, height int) gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Image(class(self).GetImage(gc, gd.Int(width), gd.Int(height), false, false, true))
}

/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Go) GetSeamlessImage(width int, height int) gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Image(class(self).GetSeamlessImage(gc, gd.Int(width), gd.Int(height), false, false, gd.Float(0.1), true))
}

/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Go) GetImage3d(width int, height int, depth int) gd.ArrayOf[gdclass.Image] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Image](class(self).GetImage3d(gc, gd.Int(width), gd.Int(height), gd.Int(depth), false, true))
}

/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Go) GetSeamlessImage3d(width int, height int, depth int) gd.ArrayOf[gdclass.Image] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Image](class(self).GetSeamlessImage3d(gc, gd.Int(width), gd.Int(height), gd.Int(depth), false, gd.Float(0.1), true))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Noise
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Noise"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the 1D noise value at the given (x) coordinate.
*/
//go:nosplit
func (self class) GetNoise1d(x gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_noise_1d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 2D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise2d(x gd.Float, y gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_noise_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 2D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise2dv(v gd.Vector2) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, v)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_noise_2dv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 3D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise3d(x gd.Float, y gd.Float, z gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, z)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_noise_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 3D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise3dv(v gd.Vector3) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, v)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_noise_3dv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetImage(ctx gd.Lifetime, width gd.Int, height gd.Int, invert bool, in_3d_space bool, normalize bool) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, in_3d_space)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetSeamlessImage(ctx gd.Lifetime, width gd.Int, height gd.Int, invert bool, in_3d_space bool, skirt gd.Float, normalize bool) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, in_3d_space)
	callframe.Arg(frame, skirt)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_seamless_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetImage3d(ctx gd.Lifetime, width gd.Int, height gd.Int, depth gd.Int, invert bool, normalize bool) gd.ArrayOf[gdclass.Image] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_image_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Image](ret)
}
/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetSeamlessImage3d(ctx gd.Lifetime, width gd.Int, height gd.Int, depth gd.Int, invert bool, skirt gd.Float, normalize bool) gd.ArrayOf[gdclass.Image] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, skirt)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Noise.Bind_get_seamless_image_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Image](ret)
}
func (self class) AsNoise() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNoise() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Noise", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
