package PhysicsDirectBodyState3DExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PhysicsDirectBodyState3D"
import "grow.graphics/gd/variant/Vector3"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Basis"
import "grow.graphics/gd/variant/Transform3D"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This class extends [PhysicsDirectBodyState3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectBodyState3D].

	// PhysicsDirectBodyState3DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsDirectBodyState3DExtension interface {
		GetTotalGravity() Vector3.XYZ
		GetTotalLinearDamp() Float.X
		GetTotalAngularDamp() Float.X
		GetCenterOfMass() Vector3.XYZ
		GetCenterOfMassLocal() Vector3.XYZ
		GetPrincipalInertiaAxes() Basis.XYZ
		GetInverseMass() Float.X
		GetInverseInertia() Vector3.XYZ
		GetInverseInertiaTensor() Basis.XYZ
		SetLinearVelocity(velocity Vector3.XYZ)
		GetLinearVelocity() Vector3.XYZ
		SetAngularVelocity(velocity Vector3.XYZ)
		GetAngularVelocity() Vector3.XYZ
		SetTransform(transform Transform3D.BasisOrigin)
		GetTransform() Transform3D.BasisOrigin
		GetVelocityAtLocalPosition(local_position Vector3.XYZ) Vector3.XYZ
		ApplyCentralImpulse(impulse Vector3.XYZ)
		ApplyImpulse(impulse Vector3.XYZ, position Vector3.XYZ)
		ApplyTorqueImpulse(impulse Vector3.XYZ)
		ApplyCentralForce(force Vector3.XYZ)
		ApplyForce(force Vector3.XYZ, position Vector3.XYZ)
		ApplyTorque(torque Vector3.XYZ)
		AddConstantCentralForce(force Vector3.XYZ)
		AddConstantForce(force Vector3.XYZ, position Vector3.XYZ)
		AddConstantTorque(torque Vector3.XYZ)
		SetConstantForce(force Vector3.XYZ)
		GetConstantForce() Vector3.XYZ
		SetConstantTorque(torque Vector3.XYZ)
		GetConstantTorque() Vector3.XYZ
		SetSleepState(enabled bool)
		IsSleeping() bool
		GetContactCount() int
		GetContactLocalPosition(contact_idx int) Vector3.XYZ
		GetContactLocalNormal(contact_idx int) Vector3.XYZ
		GetContactImpulse(contact_idx int) Vector3.XYZ
		GetContactLocalShape(contact_idx int) int
		GetContactLocalVelocityAtPosition(contact_idx int) Vector3.XYZ
		GetContactCollider(contact_idx int) Resource.ID
		GetContactColliderPosition(contact_idx int) Vector3.XYZ
		GetContactColliderId(contact_idx int) int
		GetContactColliderObject(contact_idx int) gd.Object
		GetContactColliderShape(contact_idx int) int
		GetContactColliderVelocityAtPosition(contact_idx int) Vector3.XYZ
		GetStep() Float.X
		IntegrateForces()
		GetSpaceState() objects.PhysicsDirectSpaceState3D
	}
*/
type Instance [1]classdb.PhysicsDirectBodyState3DExtension

func (Instance) _get_total_gravity(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_total_linear_damp(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _get_total_angular_damp(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _get_center_of_mass(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_principal_inertia_axes(impl func(ptr unsafe.Pointer) Basis.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Basis(ret))
	}
}
func (Instance) _get_inverse_mass(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _get_inverse_inertia(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_inverse_inertia_tensor(impl func(ptr unsafe.Pointer) Basis.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Basis(ret))
	}
}
func (Instance) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}
func (Instance) _get_linear_velocity(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}
func (Instance) _get_angular_velocity(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _set_transform(impl func(ptr unsafe.Pointer, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, transform)
	}
}
func (Instance) _get_transform(impl func(ptr unsafe.Pointer) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}
func (Instance) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position Vector3.XYZ) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var local_position = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}
func (Instance) _apply_impulse(impl func(ptr unsafe.Pointer, impulse Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse, position)
	}
}
func (Instance) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}
func (Instance) _apply_central_force(impl func(ptr unsafe.Pointer, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}
func (Instance) _apply_force(impl func(ptr unsafe.Pointer, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}
func (Instance) _apply_torque(impl func(ptr unsafe.Pointer, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}
func (Instance) _add_constant_central_force(impl func(ptr unsafe.Pointer, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}
func (Instance) _add_constant_force(impl func(ptr unsafe.Pointer, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}
func (Instance) _add_constant_torque(impl func(ptr unsafe.Pointer, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}
func (Instance) _set_constant_force(impl func(ptr unsafe.Pointer, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}
func (Instance) _get_constant_force(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _set_constant_torque(impl func(ptr unsafe.Pointer, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}
func (Instance) _get_constant_torque(impl func(ptr unsafe.Pointer) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enabled)
	}
}
func (Instance) _is_sleeping(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_contact_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx int) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx int) gd.Object) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(contact_idx))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (Instance) _get_step(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _integrate_forces(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _get_space_state(impl func(ptr unsafe.Pointer) objects.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
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
type class [1]classdb.PhysicsDirectBodyState3DExtension

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectBodyState3DExtension"))
	return Instance{classdb.PhysicsDirectBodyState3DExtension(object)}
}

func (class) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_total_linear_damp(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_total_angular_damp(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_principal_inertia_axes(impl func(ptr unsafe.Pointer) gd.Basis) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_inverse_mass(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_inverse_inertia_tensor(impl func(ptr unsafe.Pointer) gd.Basis) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}

func (class) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var velocity = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, velocity)
	}
}

func (class) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, transform)
	}
}

func (class) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector3) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var local_position = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}

func (class) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse, position)
	}
}

func (class) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, impulse)
	}
}

func (class) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

func (class) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}

func (class) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}

func (class) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

func (class) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force, position)
	}
}

func (class) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}

func (class) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var force = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

func (class) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, torque)
	}
}

func (class) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var enabled = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enabled)
	}
}

func (class) _is_sleeping(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Object) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var contact_idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_step(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _integrate_forces(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _get_space_state(impl func(ptr unsafe.Pointer) objects.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (self class) AsPhysicsDirectBodyState3DExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectBodyState3DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPhysicsDirectBodyState3D() PhysicsDirectBodyState3D.Advanced {
	return *((*PhysicsDirectBodyState3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectBodyState3D() PhysicsDirectBodyState3D.Instance {
	return *((*PhysicsDirectBodyState3D.Instance)(unsafe.Pointer(&self)))
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
	case "_get_principal_inertia_axes":
		return reflect.ValueOf(self._get_principal_inertia_axes)
	case "_get_inverse_mass":
		return reflect.ValueOf(self._get_inverse_mass)
	case "_get_inverse_inertia":
		return reflect.ValueOf(self._get_inverse_inertia)
	case "_get_inverse_inertia_tensor":
		return reflect.ValueOf(self._get_inverse_inertia_tensor)
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
	case "_get_contact_impulse":
		return reflect.ValueOf(self._get_contact_impulse)
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
	case "_get_step":
		return reflect.ValueOf(self._get_step)
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	case "_get_space_state":
		return reflect.ValueOf(self._get_space_state)
	default:
		return gd.VirtualByName(self.AsPhysicsDirectBodyState3D(), name)
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
	case "_get_principal_inertia_axes":
		return reflect.ValueOf(self._get_principal_inertia_axes)
	case "_get_inverse_mass":
		return reflect.ValueOf(self._get_inverse_mass)
	case "_get_inverse_inertia":
		return reflect.ValueOf(self._get_inverse_inertia)
	case "_get_inverse_inertia_tensor":
		return reflect.ValueOf(self._get_inverse_inertia_tensor)
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
	case "_get_contact_impulse":
		return reflect.ValueOf(self._get_contact_impulse)
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
	case "_get_step":
		return reflect.ValueOf(self._get_step)
	case "_integrate_forces":
		return reflect.ValueOf(self._integrate_forces)
	case "_get_space_state":
		return reflect.ValueOf(self._get_space_state)
	default:
		return gd.VirtualByName(self.AsPhysicsDirectBodyState3D(), name)
	}
}
func init() {
	classdb.Register("PhysicsDirectBodyState3DExtension", func(ptr gd.Object) any { return classdb.PhysicsDirectBodyState3DExtension(ptr) })
}
