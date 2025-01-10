// Package SpringArm3D provides methods for working with SpringArm3D object instances.
package SpringArm3D

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
import "graphics.gd/variant/Float"
import "graphics.gd/classdb/Resource"

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
[SpringArm3D] casts a ray or a shape along its Z axis and moves all its direct children to the collision point, with an optional margin. This is useful for 3rd person cameras that move closer to the player when inside a tight space (you may need to exclude the player's collider from the [SpringArm3D]'s collision check).
*/
type Instance [1]gdclass.SpringArm3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSpringArm3D() Instance
}

/*
Returns the spring arm's current length.
*/
func (self Instance) GetHitLength() Float.X {
	return Float.X(Float.X(class(self).GetHitLength()))
}

/*
Adds the [PhysicsBody3D] object with the given [RID] to the list of [PhysicsBody3D] objects excluded from the collision check.
*/
func (self Instance) AddExcludedObject(rid Resource.ID) {
	class(self).AddExcludedObject(rid)
}

/*
Removes the given [RID] from the list of [PhysicsBody3D] objects excluded from the collision check.
*/
func (self Instance) RemoveExcludedObject(rid Resource.ID) bool {
	return bool(class(self).RemoveExcludedObject(rid))
}

/*
Clears the list of [PhysicsBody3D] objects excluded from the collision check.
*/
func (self Instance) ClearExcludedObjects() {
	class(self).ClearExcludedObjects()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SpringArm3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SpringArm3D"))
	casted := Instance{*(*gdclass.SpringArm3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) Shape() [1]gdclass.Shape3D {
	return [1]gdclass.Shape3D(class(self).GetShape())
}

func (self Instance) SetShape(value [1]gdclass.Shape3D) {
	class(self).SetShape(value)
}

func (self Instance) SpringLength() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

func (self Instance) SetSpringLength(value Float.X) {
	class(self).SetLength(gd.Float(value))
}

func (self Instance) Margin() Float.X {
	return Float.X(Float.X(class(self).GetMargin()))
}

func (self Instance) SetMargin(value Float.X) {
	class(self).SetMargin(gd.Float(value))
}

/*
Returns the spring arm's current length.
*/
//go:nosplit
func (self class) GetHitLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_get_hit_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLength(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShape(shape [1]gdclass.Shape3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() [1]gdclass.Shape3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Shape3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Shape3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds the [PhysicsBody3D] object with the given [RID] to the list of [PhysicsBody3D] objects excluded from the collision check.
*/
//go:nosplit
func (self class) AddExcludedObject(rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_add_excluded_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the given [RID] from the list of [PhysicsBody3D] objects excluded from the collision check.
*/
//go:nosplit
func (self class) RemoveExcludedObject(rid gd.RID) bool {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_remove_excluded_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears the list of [PhysicsBody3D] objects excluded from the collision check.
*/
//go:nosplit
func (self class) ClearExcludedObjects() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_clear_excluded_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SpringArm3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSpringArm3D() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSpringArm3D() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
	gdclass.Register("SpringArm3D", func(ptr gd.Object) any { return [1]gdclass.SpringArm3D{*(*gdclass.SpringArm3D)(unsafe.Pointer(&ptr))} })
}
