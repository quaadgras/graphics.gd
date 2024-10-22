package VehicleWheel3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node used as a child of a [VehicleBody3D] parent to simulate the behavior of one of its wheels. This node also acts as a collider to detect if the wheel is touching a surface.
[b]Note:[/b] This class has known issues and isn't designed to provide realistic 3D vehicle physics. If you want advanced vehicle physics, you may need to write your own physics integration using another [PhysicsBody3D] class.

*/
type Go [1]classdb.VehicleWheel3D

/*
Returns [code]true[/code] if this wheel is in contact with a surface.
*/
func (self Go) IsInContact() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsInContact())
}

/*
Returns the contacting body node if valid in the tree, as [Node3D]. At the moment, [GridMap] is not supported so the node will be always of type [PhysicsBody3D].
Returns [code]null[/code] if the wheel is not in contact with a surface, or the contact body is not a [PhysicsBody3D].
*/
func (self Go) GetContactBody() gdclass.Node3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node3D(class(self).GetContactBody(gc))
}

/*
Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is skidding (the wheel has lost grip, e.g. icy terrain), 1.0 means not skidding (the wheel has full grip, e.g. dry asphalt road).
*/
func (self Go) GetSkidinfo() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetSkidinfo()))
}

/*
Returns the rotational speed of the wheel in revolutions per minute.
*/
func (self Go) GetRpm() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetRpm()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VehicleWheel3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VehicleWheel3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) EngineForce() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetEngineForce()))
}

func (self Go) SetEngineForce(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEngineForce(gd.Float(value))
}

func (self Go) Brake() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBrake()))
}

func (self Go) SetBrake(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBrake(gd.Float(value))
}

func (self Go) Steering() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSteering()))
}

func (self Go) SetSteering(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSteering(gd.Float(value))
}

func (self Go) UseAsTraction() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsedAsTraction())
}

func (self Go) SetUseAsTraction(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseAsTraction(value)
}

func (self Go) UseAsSteering() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsedAsSteering())
}

func (self Go) SetUseAsSteering(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseAsSteering(value)
}

func (self Go) WheelRollInfluence() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRollInfluence()))
}

func (self Go) SetWheelRollInfluence(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRollInfluence(gd.Float(value))
}

func (self Go) WheelRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetWheelRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadius(gd.Float(value))
}

func (self Go) WheelRestLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSuspensionRestLength()))
}

func (self Go) SetWheelRestLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSuspensionRestLength(gd.Float(value))
}

func (self Go) WheelFrictionSlip() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFrictionSlip()))
}

func (self Go) SetWheelFrictionSlip(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrictionSlip(gd.Float(value))
}

func (self Go) SuspensionTravel() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSuspensionTravel()))
}

func (self Go) SetSuspensionTravel(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSuspensionTravel(gd.Float(value))
}

func (self Go) SuspensionStiffness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSuspensionStiffness()))
}

func (self Go) SetSuspensionStiffness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSuspensionStiffness(gd.Float(value))
}

func (self Go) SuspensionMaxForce() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSuspensionMaxForce()))
}

func (self Go) SetSuspensionMaxForce(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSuspensionMaxForce(gd.Float(value))
}

func (self Go) DampingCompression() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDampingCompression()))
}

func (self Go) SetDampingCompression(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDampingCompression(gd.Float(value))
}

func (self Go) DampingRelaxation() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDampingRelaxation()))
}

func (self Go) SetDampingRelaxation(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDampingRelaxation(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSuspensionRestLength(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_suspension_rest_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSuspensionRestLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_suspension_rest_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSuspensionTravel(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_suspension_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSuspensionTravel() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_suspension_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSuspensionStiffness(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_suspension_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSuspensionStiffness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_suspension_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSuspensionMaxForce(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_suspension_max_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSuspensionMaxForce() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_suspension_max_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDampingCompression(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_damping_compression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDampingCompression() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_damping_compression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDampingRelaxation(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_damping_relaxation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDampingRelaxation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_damping_relaxation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseAsTraction(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_use_as_traction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsedAsTraction() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_is_used_as_traction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseAsSteering(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_use_as_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsedAsSteering() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_is_used_as_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrictionSlip(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_friction_slip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrictionSlip() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_friction_slip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this wheel is in contact with a surface.
*/
//go:nosplit
func (self class) IsInContact() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_is_in_contact, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the contacting body node if valid in the tree, as [Node3D]. At the moment, [GridMap] is not supported so the node will be always of type [PhysicsBody3D].
Returns [code]null[/code] if the wheel is not in contact with a surface, or the contact body is not a [PhysicsBody3D].
*/
//go:nosplit
func (self class) GetContactBody(ctx gd.Lifetime) gdclass.Node3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_contact_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node3D
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRollInfluence(roll_influence gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, roll_influence)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_roll_influence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRollInfluence() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_roll_influence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is skidding (the wheel has lost grip, e.g. icy terrain), 1.0 means not skidding (the wheel has full grip, e.g. dry asphalt road).
*/
//go:nosplit
func (self class) GetSkidinfo() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_skidinfo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the rotational speed of the wheel in revolutions per minute.
*/
//go:nosplit
func (self class) GetRpm() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_rpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEngineForce(engine_force gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, engine_force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEngineForce() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBrake() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_set_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSteering() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VehicleWheel3D.Bind_get_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVehicleWheel3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVehicleWheel3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("VehicleWheel3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
