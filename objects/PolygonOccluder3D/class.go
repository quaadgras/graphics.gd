package PolygonOccluder3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Occluder3D"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[PolygonOccluder3D] stores a polygon shape that can be used by the engine's occlusion culling system. When an [OccluderInstance3D] with a [PolygonOccluder3D] is selected in the editor, an editor will appear at the top of the 3D viewport so you can add/remove points. All points must be placed on the same 2D plane, which means it is not possible to create arbitrary 3D shapes with a single [PolygonOccluder3D]. To use arbitrary 3D shapes as occluders, use [ArrayOccluder3D] or [OccluderInstance3D]'s baking feature instead.
See [OccluderInstance3D]'s documentation for instructions on setting up occlusion culling.
*/
type Instance [1]classdb.PolygonOccluder3D
type Any interface {
	gd.IsClass
	AsPolygonOccluder3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PolygonOccluder3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PolygonOccluder3D"))
	return Instance{classdb.PolygonOccluder3D(object)}
}

func (self Instance) Polygon() []Vector2.XY {
	return []Vector2.XY(class(self).GetPolygon().AsSlice())
}

func (self Instance) SetPolygon(value []Vector2.XY) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonOccluder3D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonOccluder3D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsPolygonOccluder3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPolygonOccluder3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsOccluder3D() Occluder3D.Advanced {
	return *((*Occluder3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOccluder3D() Occluder3D.Instance {
	return *((*Occluder3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}
func init() {
	classdb.Register("PolygonOccluder3D", func(ptr gd.Object) any { return [1]classdb.PolygonOccluder3D{classdb.PolygonOccluder3D(ptr)} })
}
