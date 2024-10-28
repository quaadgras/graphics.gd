package PhysicsDirectBodyState2DExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PhysicsDirectBodyState2D"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class extends [PhysicsDirectBodyState2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectBodyState2D].
	// PhysicsDirectBodyState2DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsDirectBodyState2DExtension interface {
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
		GetTotalGravity() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
		GetTotalLinearDamp() float64
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
		GetTotalAngularDamp() float64
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
		GetCenterOfMass() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
		GetCenterOfMassLocal() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
		GetInverseMass() float64
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
		GetInverseInertia() float64
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
		SetLinearVelocity(velocity gd.Vector2) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
		GetLinearVelocity() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
		SetAngularVelocity(velocity float64) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
		GetAngularVelocity() float64
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
		SetTransform(transform gd.Transform2D) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
		GetTransform() gd.Transform2D
		//Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
		GetVelocityAtLocalPosition(local_position gd.Vector2) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
		ApplyCentralImpulse(impulse gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
		ApplyImpulse(impulse gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
		ApplyTorqueImpulse(impulse float64) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
		ApplyCentralForce(force gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_force].
		ApplyForce(force gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
		ApplyTorque(torque float64) 
		//Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
		AddConstantCentralForce(force gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
		AddConstantForce(force gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
		AddConstantTorque(torque float64) 
		//Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
		SetConstantForce(force gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
		GetConstantForce() gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
		SetConstantTorque(torque float64) 
		//Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
		GetConstantTorque() float64
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
		SetSleepState(enabled bool) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
		IsSleeping() bool
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
		GetContactCount() int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
		GetContactLocalPosition(contact_idx int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
		GetContactLocalNormal(contact_idx int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
		GetContactLocalShape(contact_idx int) int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
		GetContactLocalVelocityAtPosition(contact_idx int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
		GetContactCollider(contact_idx int) gd.RID
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
		GetContactColliderPosition(contact_idx int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
		GetContactColliderId(contact_idx int) int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
		GetContactColliderObject(contact_idx int) gd.Object
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
		GetContactColliderShape(contact_idx int) int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
		GetContactColliderVelocityAtPosition(contact_idx int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
		GetContactImpulse(contact_idx int) gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
		GetStep() float64
		//Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
		IntegrateForces() 
		//Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
		GetSpaceState() gdclass.PhysicsDirectSpaceState2D
	}

*/
type Go [1]classdb.PhysicsDirectBodyState2DExtension

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
*/
func (Go) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
*/
func (Go) _get_total_linear_damp(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
*/
func (Go) _get_total_angular_damp(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
*/
func (Go) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
*/
func (Go) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
*/
func (Go) _get_inverse_mass(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
*/
func (Go) _get_inverse_inertia(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
*/
func (Go) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
*/
func (Go) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
*/
func (Go) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(velocity))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
*/
func (Go) _get_angular_velocity(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
*/
func (Go) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, transform)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
*/
func (Go) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
*/
func (Go) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector2) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var local_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
*/
func (Go) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
*/
func (Go) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
*/
func (Go) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(impulse))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
*/
func (Go) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_force].
*/
func (Go) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
*/
func (Go) _apply_torque(impl func(ptr unsafe.Pointer, torque float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(torque))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
*/
func (Go) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
*/
func (Go) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
*/
func (Go) _add_constant_torque(impl func(ptr unsafe.Pointer, torque float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(torque))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
*/
func (Go) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
*/
func (Go) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
*/
func (Go) _set_constant_torque(impl func(ptr unsafe.Pointer, torque float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(torque))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
*/
func (Go) _get_constant_torque(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
*/
func (Go) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
*/
func (Go) _is_sleeping(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
*/
func (Go) _get_contact_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
*/
func (Go) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
*/
func (Go) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
*/
func (Go) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
*/
func (Go) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
*/
func (Go) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
*/
func (Go) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
*/
func (Go) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
*/
func (Go) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx int) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
*/
func (Go) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
*/
func (Go) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
*/
func (Go) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
*/
func (Go) _get_step(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
*/
func (Go) _integrate_forces(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
*/
func (Go) _get_space_state(impl func(ptr unsafe.Pointer) gdclass.PhysicsDirectSpaceState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicsDirectBodyState2DExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectBodyState2DExtension"))
	return Go{classdb.PhysicsDirectBodyState2DExtension(object)}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
*/
func (class) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
*/
func (class) _get_total_linear_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
*/
func (class) _get_total_angular_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
*/
func (class) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
*/
func (class) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
*/
func (class) _get_inverse_mass(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
*/
func (class) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
*/
func (class) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
*/
func (class) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
*/
func (class) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
*/
func (class) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
*/
func (class) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, transform)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
*/
func (class) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
*/
func (class) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector2) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var local_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
*/
func (class) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
*/
func (class) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
*/
func (class) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
*/
func (class) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_force].
*/
func (class) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
*/
func (class) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
*/
func (class) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
*/
func (class) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
*/
func (class) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
*/
func (class) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
*/
func (class) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
*/
func (class) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
*/
func (class) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
*/
func (class) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
*/
func (class) _is_sleeping(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
*/
func (class) _get_contact_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
*/
func (class) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
*/
func (class) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
*/
func (class) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
*/
func (class) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
*/
func (class) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
*/
func (class) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
*/
func (class) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
*/
func (class) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
*/
func (class) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
*/
func (class) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
*/
func (class) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
*/
func (class) _get_step(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
*/
func (class) _get_space_state(impl func(ptr unsafe.Pointer) gdclass.PhysicsDirectSpaceState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsPhysicsDirectBodyState2DExtension() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsDirectBodyState2DExtension() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsDirectBodyState2D() PhysicsDirectBodyState2D.GD { return *((*PhysicsDirectBodyState2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsDirectBodyState2D() PhysicsDirectBodyState2D.Go { return *((*PhysicsDirectBodyState2D.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity": return reflect.ValueOf(self._get_total_gravity);
	case "_get_total_linear_damp": return reflect.ValueOf(self._get_total_linear_damp);
	case "_get_total_angular_damp": return reflect.ValueOf(self._get_total_angular_damp);
	case "_get_center_of_mass": return reflect.ValueOf(self._get_center_of_mass);
	case "_get_center_of_mass_local": return reflect.ValueOf(self._get_center_of_mass_local);
	case "_get_inverse_mass": return reflect.ValueOf(self._get_inverse_mass);
	case "_get_inverse_inertia": return reflect.ValueOf(self._get_inverse_inertia);
	case "_set_linear_velocity": return reflect.ValueOf(self._set_linear_velocity);
	case "_get_linear_velocity": return reflect.ValueOf(self._get_linear_velocity);
	case "_set_angular_velocity": return reflect.ValueOf(self._set_angular_velocity);
	case "_get_angular_velocity": return reflect.ValueOf(self._get_angular_velocity);
	case "_set_transform": return reflect.ValueOf(self._set_transform);
	case "_get_transform": return reflect.ValueOf(self._get_transform);
	case "_get_velocity_at_local_position": return reflect.ValueOf(self._get_velocity_at_local_position);
	case "_apply_central_impulse": return reflect.ValueOf(self._apply_central_impulse);
	case "_apply_impulse": return reflect.ValueOf(self._apply_impulse);
	case "_apply_torque_impulse": return reflect.ValueOf(self._apply_torque_impulse);
	case "_apply_central_force": return reflect.ValueOf(self._apply_central_force);
	case "_apply_force": return reflect.ValueOf(self._apply_force);
	case "_apply_torque": return reflect.ValueOf(self._apply_torque);
	case "_add_constant_central_force": return reflect.ValueOf(self._add_constant_central_force);
	case "_add_constant_force": return reflect.ValueOf(self._add_constant_force);
	case "_add_constant_torque": return reflect.ValueOf(self._add_constant_torque);
	case "_set_constant_force": return reflect.ValueOf(self._set_constant_force);
	case "_get_constant_force": return reflect.ValueOf(self._get_constant_force);
	case "_set_constant_torque": return reflect.ValueOf(self._set_constant_torque);
	case "_get_constant_torque": return reflect.ValueOf(self._get_constant_torque);
	case "_set_sleep_state": return reflect.ValueOf(self._set_sleep_state);
	case "_is_sleeping": return reflect.ValueOf(self._is_sleeping);
	case "_get_contact_count": return reflect.ValueOf(self._get_contact_count);
	case "_get_contact_local_position": return reflect.ValueOf(self._get_contact_local_position);
	case "_get_contact_local_normal": return reflect.ValueOf(self._get_contact_local_normal);
	case "_get_contact_local_shape": return reflect.ValueOf(self._get_contact_local_shape);
	case "_get_contact_local_velocity_at_position": return reflect.ValueOf(self._get_contact_local_velocity_at_position);
	case "_get_contact_collider": return reflect.ValueOf(self._get_contact_collider);
	case "_get_contact_collider_position": return reflect.ValueOf(self._get_contact_collider_position);
	case "_get_contact_collider_id": return reflect.ValueOf(self._get_contact_collider_id);
	case "_get_contact_collider_object": return reflect.ValueOf(self._get_contact_collider_object);
	case "_get_contact_collider_shape": return reflect.ValueOf(self._get_contact_collider_shape);
	case "_get_contact_collider_velocity_at_position": return reflect.ValueOf(self._get_contact_collider_velocity_at_position);
	case "_get_contact_impulse": return reflect.ValueOf(self._get_contact_impulse);
	case "_get_step": return reflect.ValueOf(self._get_step);
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	case "_get_space_state": return reflect.ValueOf(self._get_space_state);
	default: return gd.VirtualByName(self.AsPhysicsDirectBodyState2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity": return reflect.ValueOf(self._get_total_gravity);
	case "_get_total_linear_damp": return reflect.ValueOf(self._get_total_linear_damp);
	case "_get_total_angular_damp": return reflect.ValueOf(self._get_total_angular_damp);
	case "_get_center_of_mass": return reflect.ValueOf(self._get_center_of_mass);
	case "_get_center_of_mass_local": return reflect.ValueOf(self._get_center_of_mass_local);
	case "_get_inverse_mass": return reflect.ValueOf(self._get_inverse_mass);
	case "_get_inverse_inertia": return reflect.ValueOf(self._get_inverse_inertia);
	case "_set_linear_velocity": return reflect.ValueOf(self._set_linear_velocity);
	case "_get_linear_velocity": return reflect.ValueOf(self._get_linear_velocity);
	case "_set_angular_velocity": return reflect.ValueOf(self._set_angular_velocity);
	case "_get_angular_velocity": return reflect.ValueOf(self._get_angular_velocity);
	case "_set_transform": return reflect.ValueOf(self._set_transform);
	case "_get_transform": return reflect.ValueOf(self._get_transform);
	case "_get_velocity_at_local_position": return reflect.ValueOf(self._get_velocity_at_local_position);
	case "_apply_central_impulse": return reflect.ValueOf(self._apply_central_impulse);
	case "_apply_impulse": return reflect.ValueOf(self._apply_impulse);
	case "_apply_torque_impulse": return reflect.ValueOf(self._apply_torque_impulse);
	case "_apply_central_force": return reflect.ValueOf(self._apply_central_force);
	case "_apply_force": return reflect.ValueOf(self._apply_force);
	case "_apply_torque": return reflect.ValueOf(self._apply_torque);
	case "_add_constant_central_force": return reflect.ValueOf(self._add_constant_central_force);
	case "_add_constant_force": return reflect.ValueOf(self._add_constant_force);
	case "_add_constant_torque": return reflect.ValueOf(self._add_constant_torque);
	case "_set_constant_force": return reflect.ValueOf(self._set_constant_force);
	case "_get_constant_force": return reflect.ValueOf(self._get_constant_force);
	case "_set_constant_torque": return reflect.ValueOf(self._set_constant_torque);
	case "_get_constant_torque": return reflect.ValueOf(self._get_constant_torque);
	case "_set_sleep_state": return reflect.ValueOf(self._set_sleep_state);
	case "_is_sleeping": return reflect.ValueOf(self._is_sleeping);
	case "_get_contact_count": return reflect.ValueOf(self._get_contact_count);
	case "_get_contact_local_position": return reflect.ValueOf(self._get_contact_local_position);
	case "_get_contact_local_normal": return reflect.ValueOf(self._get_contact_local_normal);
	case "_get_contact_local_shape": return reflect.ValueOf(self._get_contact_local_shape);
	case "_get_contact_local_velocity_at_position": return reflect.ValueOf(self._get_contact_local_velocity_at_position);
	case "_get_contact_collider": return reflect.ValueOf(self._get_contact_collider);
	case "_get_contact_collider_position": return reflect.ValueOf(self._get_contact_collider_position);
	case "_get_contact_collider_id": return reflect.ValueOf(self._get_contact_collider_id);
	case "_get_contact_collider_object": return reflect.ValueOf(self._get_contact_collider_object);
	case "_get_contact_collider_shape": return reflect.ValueOf(self._get_contact_collider_shape);
	case "_get_contact_collider_velocity_at_position": return reflect.ValueOf(self._get_contact_collider_velocity_at_position);
	case "_get_contact_impulse": return reflect.ValueOf(self._get_contact_impulse);
	case "_get_step": return reflect.ValueOf(self._get_step);
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	case "_get_space_state": return reflect.ValueOf(self._get_space_state);
	default: return gd.VirtualByName(self.AsPhysicsDirectBodyState2D(), name)
	}
}
func init() {classdb.Register("PhysicsDirectBodyState2DExtension", func(ptr gd.Object) any { return classdb.PhysicsDirectBodyState2DExtension(ptr) })}
