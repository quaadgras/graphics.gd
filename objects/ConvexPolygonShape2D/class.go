package ConvexPolygonShape2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Shape2D"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 2D convex polygon shape, intended for use in physics. Used internally in [CollisionPolygon2D] when it's in [constant CollisionPolygon2D.BUILD_SOLIDS] mode.
[ConvexPolygonShape2D] is [i]solid[/i], which means it detects collisions from objects that are fully inside it, unlike [ConcavePolygonShape2D] which is hollow. This makes it more suitable for both detection and physics.
[b]Convex decomposition:[/b] A concave polygon can be split up into several convex polygons. This allows dynamic physics bodies to have complex concave collisions (at a performance cost) and can be achieved by using several [ConvexPolygonShape2D] nodes or by using the [CollisionPolygon2D] node in [constant CollisionPolygon2D.BUILD_SOLIDS] mode. To generate a collision polygon from a sprite, select the [Sprite2D] node, go to the [b]Sprite2D[/b] menu that appears above the viewport, and choose [b]Create Polygon2D Sibling[/b].
[b]Performance:[/b] [ConvexPolygonShape2D] is faster to check collisions against compared to [ConcavePolygonShape2D], but it is slower than primitive collision shapes such as [CircleShape2D] and [RectangleShape2D]. Its use should generally be limited to medium-sized objects that cannot have their collision accurately represented by primitive shapes.
*/
type Instance [1]classdb.ConvexPolygonShape2D
type Any interface {
	gd.IsClass
	AsConvexPolygonShape2D() Instance
}

/*
Based on the set of points provided, this assigns the [member points] property using the convex hull algorithm, removing all unneeded points. See [method Geometry2D.convex_hull] for details.
*/
func (self Instance) SetPointCloud(point_cloud []Vector2.XY) {
	class(self).SetPointCloud(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&point_cloud))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ConvexPolygonShape2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConvexPolygonShape2D"))
	return Instance{classdb.ConvexPolygonShape2D(object)}
}

func (self Instance) Points() []Vector2.XY {
	return []Vector2.XY(class(self).GetPoints().AsSlice())
}

func (self Instance) SetPoints(value []Vector2.XY) {
	class(self).SetPoints(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

/*
Based on the set of points provided, this assigns the [member points] property using the convex hull algorithm, removing all unneeded points. See [method Geometry2D.convex_hull] for details.
*/
//go:nosplit
func (self class) SetPointCloud(point_cloud gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(point_cloud))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConvexPolygonShape2D.Bind_set_point_cloud, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetPoints(points gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConvexPolygonShape2D.Bind_set_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPoints() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConvexPolygonShape2D.Bind_get_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsConvexPolygonShape2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsConvexPolygonShape2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.Advanced         { return *((*Shape2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape2D() Shape2D.Instance {
	return *((*Shape2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {
	classdb.Register("ConvexPolygonShape2D", func(ptr gd.Object) any { return [1]classdb.ConvexPolygonShape2D{classdb.ConvexPolygonShape2D(ptr)} })
}
