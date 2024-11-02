package PolygonPathFinder

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

type Instance [1]classdb.PolygonPathFinder

func (self Instance) Setup(points []gd.Vector2, connections []int32) {
	class(self).Setup(gd.NewPackedVector2Slice(points), gd.NewPackedInt32Slice(connections))
}
func (self Instance) FindPath(from gd.Vector2, to gd.Vector2) []gd.Vector2 {
	return []gd.Vector2(class(self).FindPath(from, to).AsSlice())
}
func (self Instance) GetIntersections(from gd.Vector2, to gd.Vector2) []gd.Vector2 {
	return []gd.Vector2(class(self).GetIntersections(from, to).AsSlice())
}
func (self Instance) GetClosestPoint(point gd.Vector2) gd.Vector2 {
	return gd.Vector2(class(self).GetClosestPoint(point))
}
func (self Instance) IsPointInside(point gd.Vector2) bool {
	return bool(class(self).IsPointInside(point))
}
func (self Instance) SetPointPenalty(idx int, penalty float64) {
	class(self).SetPointPenalty(gd.Int(idx), gd.Float(penalty))
}
func (self Instance) GetPointPenalty(idx int) float64 {
	return float64(float64(class(self).GetPointPenalty(gd.Int(idx))))
}
func (self Instance) GetBounds() gd.Rect2 {
	return gd.Rect2(class(self).GetBounds())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PolygonPathFinder

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PolygonPathFinder"))
	return Instance{classdb.PolygonPathFinder(object)}
}

//go:nosplit
func (self class) Setup(points gd.PackedVector2Array, connections gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, pointers.Get(connections))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) FindPath(from gd.Vector2, to gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_find_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetIntersections(from gd.Vector2, to gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_intersections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetClosestPoint(point gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsPointInside(point gd.Vector2) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_is_point_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointPenalty(idx gd.Int, penalty gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, penalty)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_set_point_penalty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPointPenalty(idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_point_penalty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBounds() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPolygonPathFinder() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPolygonPathFinder() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("PolygonPathFinder", func(ptr gd.Object) any { return classdb.PolygonPathFinder(ptr) })
}
