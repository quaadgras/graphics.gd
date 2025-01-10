// Package CollisionShape3D provides methods for working with CollisionShape3D object instances.
package CollisionShape3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
A node that provides a [Shape3D] to a [CollisionObject3D] parent and allows to edit it. This can give a detection shape to an [Area3D] or turn a [PhysicsBody3D] into a solid object.
[b]Warning:[/b] A non-uniformly scaled [CollisionShape3D] will likely not behave as expected. Make sure to keep its scale the same on all axes and adjust its [member shape] resource instead.
*/
type Instance [1]gdclass.CollisionShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCollisionShape3D() Instance
}

/*
This method does nothing.
*/
func (self Instance) ResourceChanged(resource [1]gdclass.Resource) {
	class(self).ResourceChanged(resource)
}

/*
Sets the collision shape's shape to the addition of all its convexed [MeshInstance3D] siblings geometry.
*/
func (self Instance) MakeConvexFromSiblings() {
	class(self).MakeConvexFromSiblings()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CollisionShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CollisionShape3D"))
	casted := Instance{*(*gdclass.CollisionShape3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Shape() [1]gdclass.Shape3D {
	return [1]gdclass.Shape3D(class(self).GetShape())
}

func (self Instance) SetShape(value [1]gdclass.Shape3D) {
	class(self).SetShape(value)
}

func (self Instance) Disabled() bool {
	return bool(class(self).IsDisabled())
}

func (self Instance) SetDisabled(value bool) {
	class(self).SetDisabled(value)
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) ResourceChanged(resource [1]gdclass.Resource) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape3D.Bind_resource_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetShape(shape [1]gdclass.Shape3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape3D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() [1]gdclass.Shape3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape3D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Shape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape3D.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape3D.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the collision shape's shape to the addition of all its convexed [MeshInstance3D] siblings geometry.
*/
//go:nosplit
func (self class) MakeConvexFromSiblings() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionShape3D.Bind_make_convex_from_siblings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsCollisionShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCollisionShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced       { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance    { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
	gdclass.Register("CollisionShape3D", func(ptr gd.Object) any {
		return [1]gdclass.CollisionShape3D{*(*gdclass.CollisionShape3D)(unsafe.Pointer(&ptr))}
	})
}
