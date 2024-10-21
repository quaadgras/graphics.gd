package HeightMapShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Shape3D"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 3D heightmap shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D]. This is useful for terrain, but it is limited as overhangs (such as caves) cannot be stored. Holes in a [HeightMapShape3D] are created by assigning very low values to points in the desired area.
[b]Performance:[/b] [HeightMapShape3D] is faster to check collisions against than [ConcavePolygonShape3D], but it is significantly slower than primitive shapes like [BoxShape3D].
A heightmap collision shape can also be build by using an [Image] reference:
[codeblocks]
[gdscript]
var heightmap_texture: Texture = ResourceLoader.load("res://heightmap_image.exr")
var heightmap_image: Image = heightmap_texture.get_image()
heightmap_image.convert(Image.FORMAT_RF)

var height_min: float = 0.0
var height_max: float = 10.0

update_map_data_from_image(heightmap_image, height_min, height_max)
[/gdscript]
[/codeblocks]

*/
type Simple [1]classdb.HeightMapShape3D
func (self Simple) SetMapWidth(width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMapWidth(gd.Int(width))
}
func (self Simple) GetMapWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMapWidth()))
}
func (self Simple) SetMapDepth(height int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMapDepth(gd.Int(height))
}
func (self Simple) GetMapDepth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMapDepth()))
}
func (self Simple) SetMapData(data gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMapData(data)
}
func (self Simple) GetMapData() gd.PackedFloat32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat32Array(Expert(self).GetMapData(gc))
}
func (self Simple) GetMinHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinHeight()))
}
func (self Simple) GetMaxHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMaxHeight()))
}
func (self Simple) UpdateMapDataFromImage(image [1]classdb.Image, height_min float64, height_max float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateMapDataFromImage(image, gd.Float(height_min), gd.Float(height_max))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.HeightMapShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetMapWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_set_map_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMapWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_get_map_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMapDepth(height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_set_map_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMapDepth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_get_map_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMapData(data gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_set_map_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMapData(ctx gd.Lifetime) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_get_map_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the smallest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
//go:nosplit
func (self class) GetMinHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_get_min_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the largest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
//go:nosplit
func (self class) GetMaxHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_get_max_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates [member map_data] with data read from an [Image] reference. Automatically resizes heightmap [member map_width] and [member map_depth] to fit the full image width and height.
The image needs to be in either [constant Image.FORMAT_RF] (32 bit), [constant Image.FORMAT_RH] (16 bit), or [constant Image.FORMAT_R8] (8 bit).
Each image pixel is read in as a float on the range from [code]0.0[/code] (black pixel) to [code]1.0[/code] (white pixel). This range value gets remapped to [param height_min] and [param height_max] to form the final height value.
*/
//go:nosplit
func (self class) UpdateMapDataFromImage(image object.Image, height_min gd.Float, height_max gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, height_min)
	callframe.Arg(frame, height_max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HeightMapShape3D.Bind_update_map_data_from_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsHeightMapShape3D() Expert { return self[0].AsHeightMapShape3D() }


//go:nosplit
func (self Simple) AsHeightMapShape3D() Simple { return self[0].AsHeightMapShape3D() }


//go:nosplit
func (self class) AsShape3D() Shape3D.Expert { return self[0].AsShape3D() }


//go:nosplit
func (self Simple) AsShape3D() Shape3D.Simple { return self[0].AsShape3D() }


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
func init() {classdb.Register("HeightMapShape3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
