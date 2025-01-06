// Package PhysicsDirectBodyState2DExtension provides methods for working with PhysicsDirectBodyState2DExtension object instances.
package PhysicsDirectBodyState2DExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/PhysicsDirectBodyState2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class extends [PhysicsDirectBodyState2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectBodyState2D].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicsDirectBodyState2DExtension)
*/
type Instance [1]gdclass.PhysicsDirectBodyState2DExtension
type Any interface {
	gd.IsClass
	AsPhysicsDirectBodyState2DExtension() Instance
}
type Interface interface {
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
	GetTotalGravity() Vector2.XY
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
	GetTotalLinearDamp() Float.X
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
	GetTotalAngularDamp() Float.X
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
	GetCenterOfMass() Vector2.XY
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
	GetCenterOfMassLocal() Vector2.XY
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
	GetInverseMass() Float.X
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
	GetInverseInertia() Float.X
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
	SetLinearVelocity(velocity Vector2.XY)
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
	GetLinearVelocity() Vector2.XY
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
	SetAngularVelocity(velocity Float.X)
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
	GetAngularVelocity() Float.X
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
	SetTransform(transform Transform2D.OriginXY)
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
	GetTransform() Transform2D.OriginXY
	//Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
	GetVelocityAtLocalPosition(local_position Vector2.XY) Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
	ApplyCentralImpulse(impulse Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
	ApplyImpulse(impulse Vector2.XY, position Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
	ApplyTorqueImpulse(impulse Float.X)
	//Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
	ApplyCentralForce(force Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.apply_force].
	ApplyForce(force Vector2.XY, position Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
	ApplyTorque(torque Float.X)
	//Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
	AddConstantCentralForce(force Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
	AddConstantForce(force Vector2.XY, position Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
	AddConstantTorque(torque Float.X)
	//Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
	SetConstantForce(force Vector2.XY)
	//Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
	GetConstantForce() Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
	SetConstantTorque(torque Float.X)
	//Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
	GetConstantTorque() Float.X
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
	SetSleepState(enabled bool)
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
	IsSleeping() bool
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
	GetContactCount() int
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
	GetContactLocalPosition(contact_idx int) Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
	GetContactLocalNormal(contact_idx int) Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
	GetContactLocalShape(contact_idx int) int
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
	GetContactLocalVelocityAtPosition(contact_idx int) Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
	GetContactCollider(contact_idx int) Resource.ID
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
	GetContactColliderPosition(contact_idx int) Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
	GetContactColliderId(contact_idx int) int
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
	GetContactColliderObject(contact_idx int) Object.Instance
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
	GetContactColliderShape(contact_idx int) int
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
	GetContactColliderVelocityAtPosition(contact_idx int) Vector2.XY
	//Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
	GetContactImpulse(contact_idx int) Vector2.XY
	//Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
	GetStep() Float.X
	//Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
	IntegrateForces()
	//Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
	GetSpaceState() [1]gdclass.PhysicsDirectSpaceState2D
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) GetTotalGravity() (_ Vector2.XY)             { return }
func (self Implementation) GetTotalLinearDamp() (_ Float.X)             { return }
func (self Implementation) GetTotalAngularDamp() (_ Float.X)            { return }
func (self Implementation) GetCenterOfMass() (_ Vector2.XY)             { return }
func (self Implementation) GetCenterOfMassLocal() (_ Vector2.XY)        { return }
func (self Implementation) GetInverseMass() (_ Float.X)                 { return }
func (self Implementation) GetInverseInertia() (_ Float.X)              { return }
func (self Implementation) SetLinearVelocity(velocity Vector2.XY)       { return }
func (self Implementation) GetLinearVelocity() (_ Vector2.XY)           { return }
func (self Implementation) SetAngularVelocity(velocity Float.X)         { return }
func (self Implementation) GetAngularVelocity() (_ Float.X)             { return }
func (self Implementation) SetTransform(transform Transform2D.OriginXY) { return }
func (self Implementation) GetTransform() (_ Transform2D.OriginXY)      { return }
func (self Implementation) GetVelocityAtLocalPosition(local_position Vector2.XY) (_ Vector2.XY) {
	return
}
func (self Implementation) ApplyCentralImpulse(impulse Vector2.XY)                           { return }
func (self Implementation) ApplyImpulse(impulse Vector2.XY, position Vector2.XY)             { return }
func (self Implementation) ApplyTorqueImpulse(impulse Float.X)                               { return }
func (self Implementation) ApplyCentralForce(force Vector2.XY)                               { return }
func (self Implementation) ApplyForce(force Vector2.XY, position Vector2.XY)                 { return }
func (self Implementation) ApplyTorque(torque Float.X)                                       { return }
func (self Implementation) AddConstantCentralForce(force Vector2.XY)                         { return }
func (self Implementation) AddConstantForce(force Vector2.XY, position Vector2.XY)           { return }
func (self Implementation) AddConstantTorque(torque Float.X)                                 { return }
func (self Implementation) SetConstantForce(force Vector2.XY)                                { return }
func (self Implementation) GetConstantForce() (_ Vector2.XY)                                 { return }
func (self Implementation) SetConstantTorque(torque Float.X)                                 { return }
func (self Implementation) GetConstantTorque() (_ Float.X)                                   { return }
func (self Implementation) SetSleepState(enabled bool)                                       { return }
func (self Implementation) IsSleeping() (_ bool)                                             { return }
func (self Implementation) GetContactCount() (_ int)                                         { return }
func (self Implementation) GetContactLocalPosition(contact_idx int) (_ Vector2.XY)           { return }
func (self Implementation) GetContactLocalNormal(contact_idx int) (_ Vector2.XY)             { return }
func (self Implementation) GetContactLocalShape(contact_idx int) (_ int)                     { return }
func (self Implementation) GetContactLocalVelocityAtPosition(contact_idx int) (_ Vector2.XY) { return }
func (self Implementation) GetContactCollider(contact_idx int) (_ Resource.ID)               { return }
func (self Implementation) GetContactColliderPosition(contact_idx int) (_ Vector2.XY)        { return }
func (self Implementation) GetContactColliderId(contact_idx int) (_ int)                     { return }
func (self Implementation) GetContactColliderObject(contact_idx int) (_ Object.Instance)     { return }
func (self Implementation) GetContactColliderShape(contact_idx int) (_ int)                  { return }
func (self Implementation) GetContactColliderVelocityAtPosition(contact_idx int) (_ Vector2.XY) {
	return
}
func (self Implementation) GetContactImpulse(contact_idx int) (_ Vector2.XY)        { return }
func (self Implementation) GetStep() (_ Float.X)                                    { return }
func (self Implementation) IntegrateForces()                                        { return }
func (self Implementation) GetSpaceState() (_ [1]gdclass.PhysicsDirectSpaceState2D) { return }

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
*/
func (Instance) _get_total_gravity(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
*/
func (Instance) _get_total_linear_damp(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
*/
func (Instance) _get_total_angular_damp(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
*/
func (Instance) _get_center_of_mass(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
*/
func (Instance) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
*/
func (Instance) _get_inverse_mass(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
*/
func (Instance) _get_inverse_inertia(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
*/
func (Instance) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
*/
func (Instance) _get_linear_velocity(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
*/
func (Instance) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(velocity))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
*/
func (Instance) _get_angular_velocity(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
*/
func (Instance) _set_transform(impl func(ptr unsafe.Pointer, transform Transform2D.OriginXY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, transform)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
*/
func (Instance) _get_transform(impl func(ptr unsafe.Pointer) Transform2D.OriginXY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Transform2D(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
*/
func (Instance) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position Vector2.XY) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var local_position = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
*/
func (Instance) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
*/
func (Instance) _apply_impulse(impl func(ptr unsafe.Pointer, impulse Vector2.XY, position Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
*/
func (Instance) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(impulse))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
*/
func (Instance) _apply_central_force(impl func(ptr unsafe.Pointer, force Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_force].
*/
func (Instance) _apply_force(impl func(ptr unsafe.Pointer, force Vector2.XY, position Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
*/
func (Instance) _apply_torque(impl func(ptr unsafe.Pointer, torque Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(torque))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
*/
func (Instance) _add_constant_central_force(impl func(ptr unsafe.Pointer, force Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
*/
func (Instance) _add_constant_force(impl func(ptr unsafe.Pointer, force Vector2.XY, position Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
*/
func (Instance) _add_constant_torque(impl func(ptr unsafe.Pointer, torque Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(torque))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
*/
func (Instance) _set_constant_force(impl func(ptr unsafe.Pointer, force Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
*/
func (Instance) _get_constant_force(impl func(ptr unsafe.Pointer) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
*/
func (Instance) _set_constant_torque(impl func(ptr unsafe.Pointer, torque Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(torque))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
*/
func (Instance) _get_constant_torque(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
*/
func (Instance) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enabled)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
*/
func (Instance) _is_sleeping(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
*/
func (Instance) _get_contact_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
*/
func (Instance) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
*/
func (Instance) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
*/
func (Instance) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
*/
func (Instance) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
*/
func (Instance) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx int) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
*/
func (Instance) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
*/
func (Instance) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
*/
func (Instance) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx int) Object.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
*/
func (Instance) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
*/
func (Instance) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
*/
func (Instance) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
*/
func (Instance) _get_step(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
*/
func (Instance) _integrate_forces(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
*/
func (Instance) _get_space_state(impl func(ptr unsafe.Pointer) [1]gdclass.PhysicsDirectSpaceState2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsDirectBodyState2DExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectBodyState2DExtension"))
	casted := Instance{*(*gdclass.PhysicsDirectBodyState2DExtension)(unsafe.Pointer(&object))}
	return casted
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
*/
func (class) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
*/
func (class) _get_total_linear_damp(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
*/
func (class) _get_total_angular_damp(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
*/
func (class) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
*/
func (class) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
*/
func (class) _get_inverse_mass(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
*/
func (class) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
*/
func (class) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
*/
func (class) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
*/
func (class) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
*/
func (class) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
*/
func (class) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, transform)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
*/
func (class) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
*/
func (class) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector2) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var local_position = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
*/
func (class) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
*/
func (class) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2, position gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
*/
func (class) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
*/
func (class) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_force].
*/
func (class) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
*/
func (class) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
*/
func (class) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
*/
func (class) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector2](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
*/
func (class) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
*/
func (class) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector2](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
*/
func (class) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
*/
func (class) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
*/
func (class) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
*/
func (class) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enabled)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
*/
func (class) _is_sleeping(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
*/
func (class) _get_contact_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
*/
func (class) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
*/
func (class) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
*/
func (class) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
*/
func (class) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
*/
func (class) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
*/
func (class) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
*/
func (class) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
*/
func (class) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx gd.Int) [1]gd.Object) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
*/
func (class) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
*/
func (class) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
*/
func (class) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
*/
func (class) _get_step(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
*/
func (class) _get_space_state(impl func(ptr unsafe.Pointer) [1]gdclass.PhysicsDirectSpaceState2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsPhysicsDirectBodyState2DExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectBodyState2DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPhysicsDirectBodyState2D() PhysicsDirectBodyState2D.Advanced {
	return *((*PhysicsDirectBodyState2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectBodyState2D() PhysicsDirectBodyState2D.Instance {
	return *((*PhysicsDirectBodyState2D.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity":
		return reflect.ValueOf(self._get_total_gravity)
	case "_get_total_linear_damp":
		return reflect.ValueOf(self._get_total_linear_damp)
	case "_get_total_angular_damp":
		return reflect.ValueOf(self._get_total_angular_damp)
	case "_get_center_of_mass":
		return reflect.ValueOf(self._get_center_of_mass)
	case "_get_center_of_mass_local":
		return reflect.ValueOf(self._get_center_of_mass_local)
	case "_get_inverse_mass":
		return reflect.ValueOf(self._get_inverse_mass)
	case "_get_inverse_inertia":
		return reflect.ValueOf(self._get_inverse_inertia)
	case "_set_linear_velocity":
		return reflect.ValueOf(self._set_linear_velocity)
	case "_get_linear_velocity":
		return reflect.ValueOf(self._get_linear_velocity)
	case "_set_angular_velocity":
		return reflect.ValueOf(self._set_angular_velocity)
	case "_get_angular_velocity":
		return reflect.ValueOf(self._get_angular_velocity)
	case "_set_transform":
		return reflect.ValueOf(self._set_transform)
	case "_get_transform":
		return reflect.ValueOf(self._get_transform)
	case "_get_velocity_at_local_position":
		return reflect.ValueOf(self._get_velocity_at_local_position)
	case "_apply_central_impulse":
		return reflect.ValueOf(self._apply_central_impulse)
	case "_apply_impulse":
		return reflect.ValueOf(self._apply_impulse)
	case "_apply_torque_impulse":
		return reflect.ValueOf(self._apply_torque_impulse)
	case "_apply_central_force":
		return reflect.ValueOf(self._apply_central_force)
	case "_apply_force":
		return reflect.ValueOf(self._apply_force)
	case "_apply_torque":
		return reflect.ValueOf(self._apply_torque)
	case "_add_constant_central_force":
		return reflect.ValueOf(self._add_constant_central_force)
	case "_add_constant_force":
		return reflect.ValueOf(self._add_constant_force)
	case "_add_constant_torque":
		return reflect.ValueOf(self._add_constant_torque)
	case "_set_constant_force":
		return reflect.ValueOf(self._set_constant_force)
	case "_get_constant_force":
		return reflect.ValueOf(self._get_constant_force)
	case "_set_constant_torque":
		return reflect.ValueOf(self._set_constant_torque)
	case "_get_constant_torque":
		return reflect.ValueOf(self._get_constant_torque)
	case "_set_sleep_state":
		return reflect.ValueOf(self._set_sleep_state)
	case "_is_sleeping":
		return reflect.ValueOf(self._is_sleeping)
	case "_get_contact_count":
		return reflect.ValueOf(self._get_contact_count)
	case "_get_contact_local_position":
		return reflect.ValueOf(self._get_contact_local_position)
	case "_get_contact_local_normal":
		return reflect.ValueOf(self._get_contact_local_normal)
	case "_get_contact_local_shape":
		return reflect.ValueOf(self._get_contact_local_shape)
	case "_get_contact_local_velocity_at_position":
		return reflect.ValueOf(self._get_contact_local_velocity_at_position)
	case "_get_contact_collider":
		return reflect.ValueOf(self._get_contact_collider)
	case "_get_contact_collider_position":
		return reflect.ValueOf(self._get_contact_collider_position)
	case "_get_contact_collider_id":
		return reflect.ValueOf(self._get_contact_collider_id)
	case "_get_contact_collider_object":
		return reflect.ValueOf(self._get_contact_collider_object)
	case "_get_contact_collider_shape":
		return reflect.ValueOf(self._get_contact_collider_shape)
	case "_get_contact_collider_velocity_at_position":
		return reflect.ValueOf(self._get_contact_collider_velocity_at_position)
	case "_get_contact_impulse":
		return reflect.ValueOf(self._get_contact_impulse)
	case "_get_step":
		return reflect.ValueOf(self._get_step)
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	case "_get_space_state":
		return reflect.ValueOf(self._get_space_state)
	default:
		return gd.VirtualByName(PhysicsDirectBodyState2D.Advanced(self.AsPhysicsDirectBodyState2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity":
		return reflect.ValueOf(self._get_total_gravity)
	case "_get_total_linear_damp":
		return reflect.ValueOf(self._get_total_linear_damp)
	case "_get_total_angular_damp":
		return reflect.ValueOf(self._get_total_angular_damp)
	case "_get_center_of_mass":
		return reflect.ValueOf(self._get_center_of_mass)
	case "_get_center_of_mass_local":
		return reflect.ValueOf(self._get_center_of_mass_local)
	case "_get_inverse_mass":
		return reflect.ValueOf(self._get_inverse_mass)
	case "_get_inverse_inertia":
		return reflect.ValueOf(self._get_inverse_inertia)
	case "_set_linear_velocity":
		return reflect.ValueOf(self._set_linear_velocity)
	case "_get_linear_velocity":
		return reflect.ValueOf(self._get_linear_velocity)
	case "_set_angular_velocity":
		return reflect.ValueOf(self._set_angular_velocity)
	case "_get_angular_velocity":
		return reflect.ValueOf(self._get_angular_velocity)
	case "_set_transform":
		return reflect.ValueOf(self._set_transform)
	case "_get_transform":
		return reflect.ValueOf(self._get_transform)
	case "_get_velocity_at_local_position":
		return reflect.ValueOf(self._get_velocity_at_local_position)
	case "_apply_central_impulse":
		return reflect.ValueOf(self._apply_central_impulse)
	case "_apply_impulse":
		return reflect.ValueOf(self._apply_impulse)
	case "_apply_torque_impulse":
		return reflect.ValueOf(self._apply_torque_impulse)
	case "_apply_central_force":
		return reflect.ValueOf(self._apply_central_force)
	case "_apply_force":
		return reflect.ValueOf(self._apply_force)
	case "_apply_torque":
		return reflect.ValueOf(self._apply_torque)
	case "_add_constant_central_force":
		return reflect.ValueOf(self._add_constant_central_force)
	case "_add_constant_force":
		return reflect.ValueOf(self._add_constant_force)
	case "_add_constant_torque":
		return reflect.ValueOf(self._add_constant_torque)
	case "_set_constant_force":
		return reflect.ValueOf(self._set_constant_force)
	case "_get_constant_force":
		return reflect.ValueOf(self._get_constant_force)
	case "_set_constant_torque":
		return reflect.ValueOf(self._set_constant_torque)
	case "_get_constant_torque":
		return reflect.ValueOf(self._get_constant_torque)
	case "_set_sleep_state":
		return reflect.ValueOf(self._set_sleep_state)
	case "_is_sleeping":
		return reflect.ValueOf(self._is_sleeping)
	case "_get_contact_count":
		return reflect.ValueOf(self._get_contact_count)
	case "_get_contact_local_position":
		return reflect.ValueOf(self._get_contact_local_position)
	case "_get_contact_local_normal":
		return reflect.ValueOf(self._get_contact_local_normal)
	case "_get_contact_local_shape":
		return reflect.ValueOf(self._get_contact_local_shape)
	case "_get_contact_local_velocity_at_position":
		return reflect.ValueOf(self._get_contact_local_velocity_at_position)
	case "_get_contact_collider":
		return reflect.ValueOf(self._get_contact_collider)
	case "_get_contact_collider_position":
		return reflect.ValueOf(self._get_contact_collider_position)
	case "_get_contact_collider_id":
		return reflect.ValueOf(self._get_contact_collider_id)
	case "_get_contact_collider_object":
		return reflect.ValueOf(self._get_contact_collider_object)
	case "_get_contact_collider_shape":
		return reflect.ValueOf(self._get_contact_collider_shape)
	case "_get_contact_collider_velocity_at_position":
		return reflect.ValueOf(self._get_contact_collider_velocity_at_position)
	case "_get_contact_impulse":
		return reflect.ValueOf(self._get_contact_impulse)
	case "_get_step":
		return reflect.ValueOf(self._get_step)
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	case "_get_space_state":
		return reflect.ValueOf(self._get_space_state)
	default:
		return gd.VirtualByName(PhysicsDirectBodyState2D.Instance(self.AsPhysicsDirectBodyState2D()), name)
	}
}
func init() {
	gdclass.Register("PhysicsDirectBodyState2DExtension", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsDirectBodyState2DExtension{*(*gdclass.PhysicsDirectBodyState2DExtension)(unsafe.Pointer(&ptr))}
	})
}
