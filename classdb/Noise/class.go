// Package Noise provides methods for working with Noise object instances.
package Noise

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class defines the interface for noise generation libraries to inherit from.
A default [method get_seamless_image] implementation is provided for libraries that do not provide seamless noise. This function requests a larger image from the [method get_image] method, reverses the quadrants of the image, then uses the strips of extra width to blend over the seams.
Inheriting noise classes can optionally override this function to provide a more optimal algorithm.
*/
type Instance [1]gdclass.Noise
type Any interface {
	gd.IsClass
	AsNoise() Instance
}

/*
Returns the 1D noise value at the given (x) coordinate.
*/
func (self Instance) GetNoise1d(x Float.X) Float.X {
	return Float.X(Float.X(class(self).GetNoise1d(gd.Float(x))))
}

/*
Returns the 2D noise value at the given position.
*/
func (self Instance) GetNoise2d(x Float.X, y Float.X) Float.X {
	return Float.X(Float.X(class(self).GetNoise2d(gd.Float(x), gd.Float(y))))
}

/*
Returns the 2D noise value at the given position.
*/
func (self Instance) GetNoise2dv(v Vector2.XY) Float.X {
	return Float.X(Float.X(class(self).GetNoise2dv(gd.Vector2(v))))
}

/*
Returns the 3D noise value at the given position.
*/
func (self Instance) GetNoise3d(x Float.X, y Float.X, z Float.X) Float.X {
	return Float.X(Float.X(class(self).GetNoise3d(gd.Float(x), gd.Float(y), gd.Float(z))))
}

/*
Returns the 3D noise value at the given position.
*/
func (self Instance) GetNoise3dv(v Vector3.XYZ) Float.X {
	return Float.X(Float.X(class(self).GetNoise3dv(gd.Vector3(v))))
}

/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetImage(width int, height int) [1]gdclass.Image {
	return [1]gdclass.Image(class(self).GetImage(gd.Int(width), gd.Int(height), false, false, true))
}

/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetSeamlessImage(width int, height int) [1]gdclass.Image {
	return [1]gdclass.Image(class(self).GetSeamlessImage(gd.Int(width), gd.Int(height), false, false, gd.Float(0.1), true))
}

/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetImage3d(width int, height int, depth int) gd.Array {
	return gd.Array(class(self).GetImage3d(gd.Int(width), gd.Int(height), gd.Int(depth), false, true))
}

/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetSeamlessImage3d(width int, height int, depth int) gd.Array {
	return gd.Array(class(self).GetSeamlessImage3d(gd.Int(width), gd.Int(height), gd.Int(depth), false, gd.Float(0.1), true))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Noise

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Noise"))
	return Instance{*(*gdclass.Noise)(unsafe.Pointer(&object))}
}

/*
Returns the 1D noise value at the given (x) coordinate.
*/
//go:nosplit
func (self class) GetNoise1d(x gd.Float) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_1d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 2D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise2d(x gd.Float, y gd.Float) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 2D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise2dv(v gd.Vector2) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, v)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_2dv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise3d(x gd.Float, y gd.Float, z gd.Float) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, z)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise3dv(v gd.Vector3) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, v)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_3dv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetImage(width gd.Int, height gd.Int, invert bool, in_3d_space bool, normalize bool) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, in_3d_space)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetSeamlessImage(width gd.Int, height gd.Int, invert bool, in_3d_space bool, skirt gd.Float, normalize bool) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, in_3d_space)
	callframe.Arg(frame, skirt)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_seamless_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetImage3d(width gd.Int, height gd.Int, depth gd.Int, invert bool, normalize bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_image_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetSeamlessImage3d(width gd.Int, height gd.Int, depth gd.Int, invert bool, skirt gd.Float, normalize bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, skirt)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_seamless_image_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsNoise() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNoise() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	gdclass.Register("Noise", func(ptr gd.Object) any { return [1]gdclass.Noise{*(*gdclass.Noise)(unsafe.Pointer(&ptr))} })
}
