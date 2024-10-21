package BitMap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A two-dimensional array of boolean values, can be used to efficiently store a binary matrix (every matrix element takes only one bit) and query the values using natural cartesian coordinates.

*/
type Simple [1]classdb.BitMap
func (self Simple) Create(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Create(size)
}
func (self Simple) CreateFromImageAlpha(image [1]classdb.Image, threshold float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateFromImageAlpha(image, gd.Float(threshold))
}
func (self Simple) SetBitv(position gd.Vector2i, bit bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBitv(position, bit)
}
func (self Simple) SetBit(x int, y int, bit bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBit(gd.Int(x), gd.Int(y), bit)
}
func (self Simple) GetBitv(position gd.Vector2i) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetBitv(position))
}
func (self Simple) GetBit(x int, y int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetBit(gd.Int(x), gd.Int(y)))
}
func (self Simple) SetBitRect(rect gd.Rect2i, bit bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBitRect(rect, bit)
}
func (self Simple) GetTrueBitCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTrueBitCount()))
}
func (self Simple) GetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSize())
}
func (self Simple) Resize(new_size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Resize(new_size)
}
func (self Simple) GrowMask(pixels int, rect gd.Rect2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GrowMask(gd.Int(pixels), rect)
}
func (self Simple) ConvertToImage() [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).ConvertToImage(gc))
}
func (self Simple) OpaqueToPolygons(rect gd.Rect2i, epsilon float64) gd.ArrayOf[gd.PackedVector2Array] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.PackedVector2Array](Expert(self).OpaqueToPolygons(gc, rect, gd.Float(epsilon)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.BitMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) CreateFromImageAlpha(image object.Image, threshold gd.Float)  {
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
func (self class) ConvertToImage(ctx gd.Lifetime) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BitMap.Bind_convert_to_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
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

//go:nosplit
func (self class) AsBitMap() Expert { return self[0].AsBitMap() }


//go:nosplit
func (self Simple) AsBitMap() Simple { return self[0].AsBitMap() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("BitMap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
