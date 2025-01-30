// Package BitMap provides methods for working with BitMap object instances.
package BitMap

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2i"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

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
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A two-dimensional array of boolean values, can be used to efficiently store a binary matrix (every matrix element takes only one bit) and query the values using natural cartesian coordinates.
*/
type Instance [1]gdclass.BitMap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBitMap() Instance
}

/*
Creates a bitmap with the specified size, filled with [code]false[/code].
*/
func (self Instance) Create(size Vector2i.XY) { //gd:BitMap.create
	class(self).Create(Vector2i.XY(size))
}

/*
Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to [code]false[/code] if the alpha value of the image at that position is equal to [param threshold] or less, and [code]true[/code] in other case.
*/
func (self Instance) CreateFromImageAlpha(image [1]gdclass.Image) { //gd:BitMap.create_from_image_alpha
	class(self).CreateFromImageAlpha(image, float64(0.1))
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
func (self Instance) SetBitv(position Vector2i.XY, bit bool) { //gd:BitMap.set_bitv
	class(self).SetBitv(Vector2i.XY(position), bit)
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
func (self Instance) SetBit(x int, y int, bit bool) { //gd:BitMap.set_bit
	class(self).SetBit(int64(x), int64(y), bit)
}

/*
Returns bitmap's value at the specified position.
*/
func (self Instance) GetBitv(position Vector2i.XY) bool { //gd:BitMap.get_bitv
	return bool(class(self).GetBitv(Vector2i.XY(position)))
}

/*
Returns bitmap's value at the specified position.
*/
func (self Instance) GetBit(x int, y int) bool { //gd:BitMap.get_bit
	return bool(class(self).GetBit(int64(x), int64(y)))
}

/*
Sets a rectangular portion of the bitmap to the specified value.
*/
func (self Instance) SetBitRect(rect Rect2i.PositionSize, bit bool) { //gd:BitMap.set_bit_rect
	class(self).SetBitRect(Rect2i.PositionSize(rect), bit)
}

/*
Returns the number of bitmap elements that are set to [code]true[/code].
*/
func (self Instance) GetTrueBitCount() int { //gd:BitMap.get_true_bit_count
	return int(int(class(self).GetTrueBitCount()))
}

/*
Returns bitmap's dimensions.
*/
func (self Instance) GetSize() Vector2i.XY { //gd:BitMap.get_size
	return Vector2i.XY(class(self).GetSize())
}

/*
Resizes the image to [param new_size].
*/
func (self Instance) Resize(new_size Vector2i.XY) { //gd:BitMap.resize
	class(self).Resize(Vector2i.XY(new_size))
}

/*
Applies morphological dilation or erosion to the bitmap. If [param pixels] is positive, dilation is applied to the bitmap. If [param pixels] is negative, erosion is applied to the bitmap. [param rect] defines the area where the morphological operation is applied. Pixels located outside the [param rect] are unaffected by [method grow_mask].
*/
func (self Instance) GrowMask(pixels int, rect Rect2i.PositionSize) { //gd:BitMap.grow_mask
	class(self).GrowMask(int64(pixels), Rect2i.PositionSize(rect))
}

/*
Returns an image of the same size as the bitmap and with a [enum Image.Format] of type [constant Image.FORMAT_L8]. [code]true[/code] bits of the bitmap are being converted into white pixels, and [code]false[/code] bits into black.
*/
func (self Instance) ConvertToImage() [1]gdclass.Image { //gd:BitMap.convert_to_image
	return [1]gdclass.Image(class(self).ConvertToImage())
}

/*
Creates an [Array] of polygons covering a rectangular portion of the bitmap. It uses a marching squares algorithm, followed by Ramer-Douglas-Peucker (RDP) reduction of the number of vertices. Each polygon is described as a [PackedVector2Array] of its vertices.
To get polygons covering the whole bitmap, pass:
[codeblock]
Rect2(Vector2(), get_size())
[/codeblock]
[param epsilon] is passed to RDP to control how accurately the polygons cover the bitmap: a lower [param epsilon] corresponds to more points in the polygons.
*/
func (self Instance) OpaqueToPolygons(rect Rect2i.PositionSize) [][]Vector2.XY { //gd:BitMap.opaque_to_polygons
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).OpaqueToPolygons(Rect2i.PositionSize(rect), float64(2.0)))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.BitMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BitMap"))
	casted := Instance{*(*gdclass.BitMap)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a bitmap with the specified size, filled with [code]false[/code].
*/
//go:nosplit
func (self class) Create(size Vector2i.XY) { //gd:BitMap.create
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to [code]false[/code] if the alpha value of the image at that position is equal to [param threshold] or less, and [code]true[/code] in other case.
*/
//go:nosplit
func (self class) CreateFromImageAlpha(image [1]gdclass.Image, threshold float64) { //gd:BitMap.create_from_image_alpha
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_create_from_image_alpha, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
//go:nosplit
func (self class) SetBitv(position Vector2i.XY, bit bool) { //gd:BitMap.set_bitv
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, bit)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_set_bitv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
//go:nosplit
func (self class) SetBit(x int64, y int64, bit bool) { //gd:BitMap.set_bit
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, bit)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_set_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns bitmap's value at the specified position.
*/
//go:nosplit
func (self class) GetBitv(position Vector2i.XY) bool { //gd:BitMap.get_bitv
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_bitv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns bitmap's value at the specified position.
*/
//go:nosplit
func (self class) GetBit(x int64, y int64) bool { //gd:BitMap.get_bit
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a rectangular portion of the bitmap to the specified value.
*/
//go:nosplit
func (self class) SetBitRect(rect Rect2i.PositionSize, bit bool) { //gd:BitMap.set_bit_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, bit)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_set_bit_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of bitmap elements that are set to [code]true[/code].
*/
//go:nosplit
func (self class) GetTrueBitCount() int64 { //gd:BitMap.get_true_bit_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_true_bit_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns bitmap's dimensions.
*/
//go:nosplit
func (self class) GetSize() Vector2i.XY { //gd:BitMap.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resizes the image to [param new_size].
*/
//go:nosplit
func (self class) Resize(new_size Vector2i.XY) { //gd:BitMap.resize
	var frame = callframe.New()
	callframe.Arg(frame, new_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies morphological dilation or erosion to the bitmap. If [param pixels] is positive, dilation is applied to the bitmap. If [param pixels] is negative, erosion is applied to the bitmap. [param rect] defines the area where the morphological operation is applied. Pixels located outside the [param rect] are unaffected by [method grow_mask].
*/
//go:nosplit
func (self class) GrowMask(pixels int64, rect Rect2i.PositionSize) { //gd:BitMap.grow_mask
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_grow_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an image of the same size as the bitmap and with a [enum Image.Format] of type [constant Image.FORMAT_L8]. [code]true[/code] bits of the bitmap are being converted into white pixels, and [code]false[/code] bits into black.
*/
//go:nosplit
func (self class) ConvertToImage() [1]gdclass.Image { //gd:BitMap.convert_to_image
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_convert_to_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates an [Array] of polygons covering a rectangular portion of the bitmap. It uses a marching squares algorithm, followed by Ramer-Douglas-Peucker (RDP) reduction of the number of vertices. Each polygon is described as a [PackedVector2Array] of its vertices.
To get polygons covering the whole bitmap, pass:
[codeblock]
Rect2(Vector2(), get_size())
[/codeblock]
[param epsilon] is passed to RDP to control how accurately the polygons cover the bitmap: a lower [param epsilon] corresponds to more points in the polygons.
*/
//go:nosplit
func (self class) OpaqueToPolygons(rect Rect2i.PositionSize, epsilon float64) Array.Contains[Packed.Array[Vector2.XY]] { //gd:BitMap.opaque_to_polygons
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_opaque_to_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsBitMap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBitMap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("BitMap", func(ptr gd.Object) any { return [1]gdclass.BitMap{*(*gdclass.BitMap)(unsafe.Pointer(&ptr))} })
}
