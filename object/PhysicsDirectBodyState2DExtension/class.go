package PhysicsDirectBodyState2DExtension

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PhysicsDirectBodyState2D"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class extends [PhysicsDirectBodyState2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectBodyState2D].
	// PhysicsDirectBodyState2DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsDirectBodyState2DExtension interface {
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
		GetTotalGravity() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
		GetTotalLinearDamp() gd.Float
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
		GetTotalAngularDamp() gd.Float
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
		GetCenterOfMass() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
		GetCenterOfMassLocal() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
		GetInverseMass() gd.Float
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
		GetInverseInertia() gd.Float
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
		SetLinearVelocity(velocity gd.Vector2) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
		GetLinearVelocity() gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
		SetAngularVelocity(velocity gd.Float) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
		GetAngularVelocity() gd.Float
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
		ApplyTorqueImpulse(impulse gd.Float) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
		ApplyCentralForce(force gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_force].
		ApplyForce(force gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
		ApplyTorque(torque gd.Float) 
		//Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
		AddConstantCentralForce(force gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
		AddConstantForce(force gd.Vector2, position gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
		AddConstantTorque(torque gd.Float) 
		//Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
		SetConstantForce(force gd.Vector2) 
		//Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
		GetConstantForce() gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
		SetConstantTorque(torque gd.Float) 
		//Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
		GetConstantTorque() gd.Float
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
		SetSleepState(enabled bool) 
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
		IsSleeping() bool
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
		GetContactCount() gd.Int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
		GetContactLocalPosition(contact_idx gd.Int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
		GetContactLocalNormal(contact_idx gd.Int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
		GetContactLocalShape(contact_idx gd.Int) gd.Int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
		GetContactLocalVelocityAtPosition(contact_idx gd.Int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
		GetContactCollider(contact_idx gd.Int) gd.RID
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
		GetContactColliderPosition(contact_idx gd.Int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
		GetContactColliderId(contact_idx gd.Int) gd.Int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
		GetContactColliderObject(contact_idx gd.Int) gd.Object
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
		GetContactColliderShape(contact_idx gd.Int) gd.Int
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
		GetContactColliderVelocityAtPosition(contact_idx gd.Int) gd.Vector2
		//Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
		GetContactImpulse(contact_idx gd.Int) gd.Vector2
		//Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
		GetStep() gd.Float
		//Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
		IntegrateForces() 
		//Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
		GetSpaceState() object.PhysicsDirectSpaceState2D
	}

*/
type Simple [1]classdb.PhysicsDirectBodyState2DExtension
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsDirectBodyState2DExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_gravity] and its respective getter.
*/
func (class) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_linear_damp] and its respective getter.
*/
func (class) _get_total_linear_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.total_angular_damp] and its respective getter.
*/
func (class) _get_total_angular_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass] and its respective getter.
*/
func (class) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.center_of_mass_local] and its respective getter.
*/
func (class) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_mass] and its respective getter.
*/
func (class) _get_inverse_mass(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.inverse_inertia] and its respective getter.
*/
func (class) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective setter.
*/
func (class) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var velocity = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.linear_velocity] and its respective getter.
*/
func (class) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective setter.
*/
func (class) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var velocity = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.angular_velocity] and its respective getter.
*/
func (class) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective setter.
*/
func (class) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, transform)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.transform] and its respective getter.
*/
func (class) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_velocity_at_local_position].
*/
func (class) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector2) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var local_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_impulse].
*/
func (class) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_impulse].
*/
func (class) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse, position)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque_impulse].
*/
func (class) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_central_force].
*/
func (class) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_force].
*/
func (class) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.apply_torque].
*/
func (class) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_central_force].
*/
func (class) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_force].
*/
func (class) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2, position gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		var position = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.add_constant_torque].
*/
func (class) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_force].
*/
func (class) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_force].
*/
func (class) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.set_constant_torque].
*/
func (class) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_constant_torque].
*/
func (class) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective setter.
*/
func (class) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.sleeping] and its respective getter.
*/
func (class) _is_sleeping(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_count].
*/
func (class) _get_contact_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_position].
*/
func (class) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_normal].
*/
func (class) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_shape].
*/
func (class) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_local_velocity_at_position].
*/
func (class) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider].
*/
func (class) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_position].
*/
func (class) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_id].
*/
func (class) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_object].
*/
func (class) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_shape].
*/
func (class) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_collider_velocity_at_position].
*/
func (class) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_contact_impulse].
*/
func (class) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Implement to override the behavior of [member PhysicsDirectBodyState2D.step] and its respective getter.
*/
func (class) _get_step(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.integrate_forces].
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Overridable version of [method PhysicsDirectBodyState2D.get_space_state].
*/
func (class) _get_space_state(impl func(ptr unsafe.Pointer) object.PhysicsDirectSpaceState2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}


//go:nosplit
func (self class) AsPhysicsDirectBodyState2DExtension() Expert { return self[0].AsPhysicsDirectBodyState2DExtension() }


//go:nosplit
func (self Simple) AsPhysicsDirectBodyState2DExtension() Simple { return self[0].AsPhysicsDirectBodyState2DExtension() }


//go:nosplit
func (self class) AsPhysicsDirectBodyState2D() PhysicsDirectBodyState2D.Expert { return self[0].AsPhysicsDirectBodyState2D() }


//go:nosplit
func (self Simple) AsPhysicsDirectBodyState2D() PhysicsDirectBodyState2D.Simple { return self[0].AsPhysicsDirectBodyState2D() }

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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicsDirectBodyState2DExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
