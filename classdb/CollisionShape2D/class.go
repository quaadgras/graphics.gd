// Package CollisionShape2D provides methods for working with CollisionShape2D object instances.
package CollisionShape2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A node that provides a [Shape2D] to a [CollisionObject2D] parent and allows to edit it. This can give a detection shape to an [Area2D] or turn a [PhysicsBody2D] into a solid object.
*/
type Instance [1]gdclass.CollisionShape2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCollisionShape2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CollisionShape2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CollisionShape2D"))
	casted := Instance{*(*gdclass.CollisionShape2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Shape() [1]gdclass.Shape2D {
	return [1]gdclass.Shape2D(class(self).GetShape())
}

func (self Instance) SetShape(value [1]gdclass.Shape2D) {
	class(self).SetShape(value)
}

func (self Instance) Disabled() bool {
	return bool(class(self).IsDisabled())
}

func (self Instance) SetDisabled(value bool) {
	class(self).SetDisabled(value)
}

func (self Instance) OneWayCollision() bool {
	return bool(class(self).IsOneWayCollisionEnabled())
}

func (self Instance) SetOneWayCollision(value bool) {
	class(self).SetOneWayCollision(value)
}

func (self Instance) OneWayCollisionMargin() Float.X {
	return Float.X(Float.X(class(self).GetOneWayCollisionMargin()))
}

func (self Instance) SetOneWayCollisionMargin(value Float.X) {
	class(self).SetOneWayCollisionMargin(gd.Float(value))
}

func (self Instance) DebugColor() Color.RGBA {
	return Color.RGBA(class(self).GetDebugColor())
}

func (self Instance) SetDebugColor(value Color.RGBA) {
	class(self).SetDebugColor(gd.Color(value))
}

//go:nosplit
func (self class) SetShape(shape [1]gdclass.Shape2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() [1]gdclass.Shape2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Shape2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisabled(disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOneWayCollision(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_set_one_way_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOneWayCollisionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_is_one_way_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOneWayCollisionMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_set_one_way_collision_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOneWayCollisionMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_get_one_way_collision_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_set_debug_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape2D.Bind_get_debug_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCollisionShape2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCollisionShape2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced       { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance    { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("CollisionShape2D", func(ptr gd.Object) any {
		return [1]gdclass.CollisionShape2D{*(*gdclass.CollisionShape2D)(unsafe.Pointer(&ptr))}
	})
}
