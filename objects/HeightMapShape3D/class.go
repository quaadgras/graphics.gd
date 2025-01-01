package HeightMapShape3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Shape3D"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
type Instance [1]classdb.HeightMapShape3D
type Any interface {
	gd.IsClass
	AsHeightMapShape3D() Instance
}

/*
Returns the smallest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
func (self Instance) GetMinHeight() Float.X {
	return Float.X(Float.X(class(self).GetMinHeight()))
}

/*
Returns the largest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
func (self Instance) GetMaxHeight() Float.X {
	return Float.X(Float.X(class(self).GetMaxHeight()))
}

/*
Updates [member map_data] with data read from an [Image] reference. Automatically resizes heightmap [member map_width] and [member map_depth] to fit the full image width and height.
The image needs to be in either [constant Image.FORMAT_RF] (32 bit), [constant Image.FORMAT_RH] (16 bit), or [constant Image.FORMAT_R8] (8 bit).
Each image pixel is read in as a float on the range from [code]0.0[/code] (black pixel) to [code]1.0[/code] (white pixel). This range value gets remapped to [param height_min] and [param height_max] to form the final height value.
*/
func (self Instance) UpdateMapDataFromImage(image objects.Image, height_min Float.X, height_max Float.X) {
	class(self).UpdateMapDataFromImage(image, gd.Float(height_min), gd.Float(height_max))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.HeightMapShape3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HeightMapShape3D"))
	return Instance{classdb.HeightMapShape3D(object)}
}

func (self Instance) MapWidth() int {
	return int(int(class(self).GetMapWidth()))
}

func (self Instance) SetMapWidth(value int) {
	class(self).SetMapWidth(gd.Int(value))
}

func (self Instance) MapDepth() int {
	return int(int(class(self).GetMapDepth()))
}

func (self Instance) SetMapDepth(value int) {
	class(self).SetMapDepth(gd.Int(value))
}

func (self Instance) MapData() []float32 {
	return []float32(class(self).GetMapData().AsSlice())
}

func (self Instance) SetMapData(value []float32) {
	class(self).SetMapData(gd.NewPackedFloat32Slice(value))
}

//go:nosplit
func (self class) SetMapWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_set_map_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMapWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_map_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMapDepth(height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_set_map_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMapDepth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_map_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMapData(data gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_set_map_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMapData() gd.PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_map_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the smallest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
//go:nosplit
func (self class) GetMinHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_min_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the largest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
//go:nosplit
func (self class) GetMaxHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_max_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) UpdateMapDataFromImage(image objects.Image, height_min gd.Float, height_max gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, height_min)
	callframe.Arg(frame, height_max)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_update_map_data_from_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsHeightMapShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHeightMapShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.Advanced     { return *((*Shape3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape3D() Shape3D.Instance {
	return *((*Shape3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsShape3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape3D(), name)
	}
}
func init() {
	classdb.Register("HeightMapShape3D", func(ptr gd.Object) any { return [1]classdb.HeightMapShape3D{classdb.HeightMapShape3D(ptr)} })
}
