// Package StaticBody3D provides methods for working with StaticBody3D object instances.
package StaticBody3D

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
A static 3D physics body. It can't be moved by external forces or contacts, but can be moved manually by other means such as code, [AnimationMixer]s (with [member AnimationMixer.callback_mode_process] set to [constant AnimationMixer.ANIMATION_CALLBACK_MODE_PROCESS_PHYSICS]), and [RemoteTransform3D].
When [StaticBody3D] is moved, it is teleported to its new position without affecting other physics bodies in its path. If this is not desired, use [AnimatableBody3D] instead.
[StaticBody3D] is useful for completely static objects like floors and walls, as well as moving surfaces like conveyor belts and circular revolving platforms (by using [member constant_linear_velocity] and [member constant_angular_velocity]).
*/
type Instance [1]gdclass.StaticBody3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStaticBody3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StaticBody3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StaticBody3D"))
	casted := Instance{*(*gdclass.StaticBody3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) PhysicsMaterialOverride() [1]gdclass.PhysicsMaterial {
	return [1]gdclass.PhysicsMaterial(class(self).GetPhysicsMaterialOverride())
}

func (self Instance) SetPhysicsMaterialOverride(value [1]gdclass.PhysicsMaterial) {
	class(self).SetPhysicsMaterialOverride(value)
}

func (self Instance) ConstantLinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetConstantLinearVelocity())
}

func (self Instance) SetConstantLinearVelocity(value Vector3.XYZ) {
	class(self).SetConstantLinearVelocity(gd.Vector3(value))
}

func (self Instance) ConstantAngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetConstantAngularVelocity())
}

func (self Instance) SetConstantAngularVelocity(value Vector3.XYZ) {
	class(self).SetConstantAngularVelocity(gd.Vector3(value))
}

//go:nosplit
func (self class) SetConstantLinearVelocity(vel gd.Vector3) { //gd:StaticBody3D.set_constant_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody3D.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetConstantAngularVelocity(vel gd.Vector3) { //gd:StaticBody3D.set_constant_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, vel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody3D.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantLinearVelocity() gd.Vector3 { //gd:StaticBody3D.get_constant_linear_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody3D.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetConstantAngularVelocity() gd.Vector3 { //gd:StaticBody3D.get_constant_angular_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody3D.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsMaterialOverride(physics_material_override [1]gdclass.PhysicsMaterial) { //gd:StaticBody3D.set_physics_material_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(physics_material_override[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody3D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsMaterialOverride() [1]gdclass.PhysicsMaterial { //gd:StaticBody3D.get_physics_material_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StaticBody3D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PhysicsMaterial{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsMaterial](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsStaticBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStaticBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody3D() PhysicsBody3D.Advanced {
	return *((*PhysicsBody3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody3D() PhysicsBody3D.Instance {
	return *((*PhysicsBody3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCollisionObject3D() CollisionObject3D.Advanced {
	return *((*CollisionObject3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject3D() CollisionObject3D.Instance {
	return *((*CollisionObject3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody3D.Advanced(self.AsPhysicsBody3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody3D.Instance(self.AsPhysicsBody3D()), name)
	}
}
func init() {
	gdclass.Register("StaticBody3D", func(ptr gd.Object) any {
		return [1]gdclass.StaticBody3D{*(*gdclass.StaticBody3D)(unsafe.Pointer(&ptr))}
	})
}
