package VehicleBody3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/RigidBody3D"
import "grow.graphics/gd/object/PhysicsBody3D"
import "grow.graphics/gd/object/CollisionObject3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This physics body implements all the physics logic needed to simulate a car. It is based on the raycast vehicle system commonly found in physics engines. Aside from a [CollisionShape3D] for the main body of the vehicle, you must also add a [VehicleWheel3D] node for each wheel. You should also add a [MeshInstance3D] to this node for the 3D model of the vehicle, but this model should generally not include meshes for the wheels. You can control the vehicle by using the [member brake], [member engine_force], and [member steering] properties. The position or orientation of this node shouldn't be changed directly.
[b]Note:[/b] The origin point of your VehicleBody3D will determine the center of gravity of your vehicle. To make the vehicle more grounded, the origin point is usually kept low, moving the [CollisionShape3D] and [MeshInstance3D] upwards.
[b]Note:[/b] This class has known issues and isn't designed to provide realistic 3D vehicle physics. If you want advanced vehicle physics, you may have to write your own physics integration using [CharacterBody3D] or [RigidBody3D].

*/
type Simple [1]classdb.VehicleBody3D
func (self Simple) SetEngineForce(engine_force float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEngineForce(gd.Float(engine_force))
}
func (self Simple) GetEngineForce() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEngineForce()))
}
func (self Simple) SetBrake(brake float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBrake(gd.Float(brake))
}
func (self Simple) GetBrake() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBrake()))
}
func (self Simple) SetSteering(steering float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSteering(gd.Float(steering))
}
func (self Simple) GetSteering() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSteering()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VehicleBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetEngineForce(engine_force gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, engine_force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleBody3D.Bind_set_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEngineForce() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleBody3D.Bind_get_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBrake(brake gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, brake)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleBody3D.Bind_set_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBrake() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleBody3D.Bind_get_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSteering(steering gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, steering)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleBody3D.Bind_set_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSteering() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleBody3D.Bind_get_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVehicleBody3D() Expert { return self[0].AsVehicleBody3D() }


//go:nosplit
func (self Simple) AsVehicleBody3D() Simple { return self[0].AsVehicleBody3D() }


//go:nosplit
func (self class) AsRigidBody3D() RigidBody3D.Expert { return self[0].AsRigidBody3D() }


//go:nosplit
func (self Simple) AsRigidBody3D() RigidBody3D.Simple { return self[0].AsRigidBody3D() }


//go:nosplit
func (self class) AsPhysicsBody3D() PhysicsBody3D.Expert { return self[0].AsPhysicsBody3D() }


//go:nosplit
func (self Simple) AsPhysicsBody3D() PhysicsBody3D.Simple { return self[0].AsPhysicsBody3D() }


//go:nosplit
func (self class) AsCollisionObject3D() CollisionObject3D.Expert { return self[0].AsCollisionObject3D() }


//go:nosplit
func (self Simple) AsCollisionObject3D() CollisionObject3D.Simple { return self[0].AsCollisionObject3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VehicleBody3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
