package BitMap

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
A two-dimensional array of boolean values, can be used to efficiently store a binary matrix (every matrix element takes only one bit) and query the values using natural cartesian coordinates.

*/
type Go [1]classdb.BitMap

/*
Creates a bitmap with the specified size, filled with [code]false[/code].
*/
func (self Go) Create(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Create(size)
}

/*
Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to [code]false[/code] if the alpha value of the image at that position is equal to [param threshold] or less, and [code]true[/code] in other case.
*/
func (self Go) CreateFromImageAlpha(image gdclass.Image) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CreateFromImageAlpha(image, gd.Float(0.1))
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
func (self Go) SetBitv(position gd.Vector2i, bit bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBitv(position, bit)
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
func (self Go) SetBit(x int, y int, bit bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBit(gd.Int(x), gd.Int(y), bit)
}

/*
Returns bitmap's value at the specified position.
*/
func (self Go) GetBitv(position gd.Vector2i) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetBitv(position))
}

/*
Returns bitmap's value at the specified position.
*/
func (self Go) GetBit(x int, y int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetBit(gd.Int(x), gd.Int(y)))
}

/*
Sets a rectangular portion of the bitmap to the specified value.
*/
func (self Go) SetBitRect(rect gd.Rect2i, bit bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBitRect(rect, bit)
}

/*
Returns the number of bitmap elements that are set to [code]true[/code].
*/
func (self Go) GetTrueBitCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTrueBitCount()))
}

/*
Returns bitmap's dimensions.
*/
func (self Go) GetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(class(self).GetSize())
}

/*
Resizes the image to [param new_size].
*/
func (self Go) Resize(new_size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Resize(new_size)
}

/*
Applies morphological dilation or erosion to the bitmap. If [param pixels] is positive, dilation is applied to the bitmap. If [param pixels] is negative, erosion is applied to the bitmap. [param rect] defines the area where the morphological operation is applied. Pixels located outside the [param rect] are unaffected by [method grow_mask].
*/
func (self Go) GrowMask(pixels int, rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).GrowMask(gd.Int(pixels), rect)
}

/*
Returns an image of the same size as the bitmap and with a [enum Image.Format] of type [constant Image.FORMAT_L8]. [code]true[/code] bits of the bitmap are being converted into white pixels, and [code]false[/code] bits into black.
*/
func (self Go) ConvertToImage() gdclass.Image {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Image(class(self).ConvertToImage(gc))
}

/*
Creates an [Array] of polygons covering a rectangular portion of the bitmap. It uses a marching squares algorithm, followed by Ramer-Douglas-Peucker (RDP) reduction of the number of vertices. Each polygon is described as a [PackedVector2Array] of its vertices.
To get polygons covering the whole bitmap, pass:
[codeblock]
Rect2(Vector2(), get_size())
[/codeblock]
[param epsilon] is passed to RDP to control how accurately the polygons cover the bitmap: a lower [param epsilon] corresponds to more points in the polygons.
*/
func (self Go) OpaqueToPolygons(rect gd.Rect2i) gd.ArrayOf[gd.PackedVector2Array] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.PackedVector2Array](class(self).OpaqueToPolygons(gc, rect, gd.Float(2.0)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BitMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("BitMap"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates a bitmap with the specified size, filled with [code]false[/code].
*/
//go:nosplit
func (self class) Create(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to [code]false[/code] if the alpha value of the image at that position is equal to [param threshold] or less, and [code]true[/code] in other case.
*/
//go:nosplit
func (self class) CreateFromImageAlpha(image gdclass.Image, threshold gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_create_from_image_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the bitmap's element at the specified position, to the specified value.
*/
//go:nosplit
func (self class) SetBitv(position gd.Vector2i, bit bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, bit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_set_bitv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the bitmap's element at the specified position, to the specified value.
*/
//go:nosplit
func (self class) SetBit(x gd.Int, y gd.Int, bit bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, bit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_set_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns bitmap's value at the specified position.
*/
//go:nosplit
func (self class) GetBitv(position gd.Vector2i) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_get_bitv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns bitmap's value at the specified position.
*/
//go:nosplit
func (self class) GetBit(x gd.Int, y gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_get_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a rectangular portion of the bitmap to the specified value.
*/
//go:nosplit
func (self class) SetBitRect(rect gd.Rect2i, bit bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, bit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_set_bit_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of bitmap elements that are set to [code]true[/code].
*/
//go:nosplit
func (self class) GetTrueBitCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_get_true_bit_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns bitmap's dimensions.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resizes the image to [param new_size].
*/
//go:nosplit
func (self class) Resize(new_size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, new_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies morphological dilation or erosion to the bitmap. If [param pixels] is positive, dilation is applied to the bitmap. If [param pixels] is negative, erosion is applied to the bitmap. [param rect] defines the area where the morphological operation is applied. Pixels located outside the [param rect] are unaffected by [method grow_mask].
*/
//go:nosplit
func (self class) GrowMask(pixels gd.Int, rect gd.Rect2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_grow_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an image of the same size as the bitmap and with a [enum Image.Format] of type [constant Image.FORMAT_L8]. [code]true[/code] bits of the bitmap are being converted into white pixels, and [code]false[/code] bits into black.
*/
//go:nosplit
func (self class) ConvertToImage(ctx gd.Lifetime) gdclass.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_convert_to_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
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
func (self class) OpaqueToPolygons(ctx gd.Lifetime, rect gd.Rect2i, epsilon gd.Float) gd.ArrayOf[gd.PackedVector2Array] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_opaque_to_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.PackedVector2Array](ret)
}
func (self class) AsBitMap() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBitMap() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("BitMap", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
