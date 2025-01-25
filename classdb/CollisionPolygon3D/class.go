// Package CollisionPolygon3D provides methods for working with CollisionPolygon3D object instances.
package CollisionPolygon3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A node that provides a thickened polygon shape (a prism) to a [CollisionObject3D] parent and allows to edit it. The polygon can be concave or convex. This can give a detection shape to an [Area3D] or turn [PhysicsBody3D] into a solid object.
[b]Warning:[/b] A non-uniformly scaled [CollisionShape3D] will likely not behave as expected. Make sure to keep its scale the same on all axes and adjust its shape resource instead.
*/
type Instance [1]gdclass.CollisionPolygon3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCollisionPolygon3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CollisionPolygon3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CollisionPolygon3D"))
	casted := Instance{*(*gdclass.CollisionPolygon3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Depth() Float.X {
	return Float.X(Float.X(class(self).GetDepth()))
}

func (self Instance) SetDepth(value Float.X) {
	class(self).SetDepth(gd.Float(value))
}

func (self Instance) Disabled() bool {
	return bool(class(self).IsDisabled())
}

func (self Instance) SetDisabled(value bool) {
	class(self).SetDisabled(value)
}

func (self Instance) Polygon() []Vector2.XY {
	return []Vector2.XY(class(self).GetPolygon().AsSlice())
}

func (self Instance) SetPolygon(value []Vector2.XY) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

func (self Instance) Margin() Float.X {
	return Float.X(Float.X(class(self).GetMargin()))
}

func (self Instance) SetMargin(value Float.X) {
	class(self).SetMargin(gd.Float(value))
}

//go:nosplit
func (self class) SetDepth(depth gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisabled(disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCollisionPolygon3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCollisionPolygon3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("CollisionPolygon3D", func(ptr gd.Object) any {
		return [1]gdclass.CollisionPolygon3D{*(*gdclass.CollisionPolygon3D)(unsafe.Pointer(&ptr))}
	})
}
