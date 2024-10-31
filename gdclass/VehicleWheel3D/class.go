package VehicleWheel3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A node used as a child of a [VehicleBody3D] parent to simulate the behavior of one of its wheels. This node also acts as a collider to detect if the wheel is touching a surface.
[b]Note:[/b] This class has known issues and isn't designed to provide realistic 3D vehicle physics. If you want advanced vehicle physics, you may need to write your own physics integration using another [PhysicsBody3D] class.
*/
type Instance [1]classdb.VehicleWheel3D

/*
Returns [code]true[/code] if this wheel is in contact with a surface.
*/
func (self Instance) IsInContact() bool {
	return bool(class(self).IsInContact())
}

/*
Returns the contacting body node if valid in the tree, as [Node3D]. At the moment, [GridMap] is not supported so the node will be always of type [PhysicsBody3D].
Returns [code]null[/code] if the wheel is not in contact with a surface, or the contact body is not a [PhysicsBody3D].
*/
func (self Instance) GetContactBody() gdclass.Node3D {
	return gdclass.Node3D(class(self).GetContactBody())
}

/*
Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is skidding (the wheel has lost grip, e.g. icy terrain), 1.0 means not skidding (the wheel has full grip, e.g. dry asphalt road).
*/
func (self Instance) GetSkidinfo() float64 {
	return float64(float64(class(self).GetSkidinfo()))
}

/*
Returns the rotational speed of the wheel in revolutions per minute.
*/
func (self Instance) GetRpm() float64 {
	return float64(float64(class(self).GetRpm()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VehicleWheel3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VehicleWheel3D"))
	return Instance{classdb.VehicleWheel3D(object)}
}

func (self Instance) EngineForce() float64 {
	return float64(float64(class(self).GetEngineForce()))
}

func (self Instance) SetEngineForce(value float64) {
	class(self).SetEngineForce(gd.Float(value))
}

func (self Instance) Brake() float64 {
	return float64(float64(class(self).GetBrake()))
}

func (self Instance) SetBrake(value float64) {
	class(self).SetBrake(gd.Float(value))
}

func (self Instance) Steering() float64 {
	return float64(float64(class(self).GetSteering()))
}

func (self Instance) SetSteering(value float64) {
	class(self).SetSteering(gd.Float(value))
}

func (self Instance) UseAsTraction() bool {
	return bool(class(self).IsUsedAsTraction())
}

func (self Instance) SetUseAsTraction(value bool) {
	class(self).SetUseAsTraction(value)
}

func (self Instance) UseAsSteering() bool {
	return bool(class(self).IsUsedAsSteering())
}

func (self Instance) SetUseAsSteering(value bool) {
	class(self).SetUseAsSteering(value)
}

func (self Instance) WheelRollInfluence() float64 {
	return float64(float64(class(self).GetRollInfluence()))
}

func (self Instance) SetWheelRollInfluence(value float64) {
	class(self).SetRollInfluence(gd.Float(value))
}

func (self Instance) WheelRadius() float64 {
	return float64(float64(class(self).GetRadius()))
}

func (self Instance) SetWheelRadius(value float64) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) WheelRestLength() float64 {
	return float64(float64(class(self).GetSuspensionRestLength()))
}

func (self Instance) SetWheelRestLength(value float64) {
	class(self).SetSuspensionRestLength(gd.Float(value))
}

func (self Instance) WheelFrictionSlip() float64 {
	return float64(float64(class(self).GetFrictionSlip()))
}

func (self Instance) SetWheelFrictionSlip(value float64) {
	class(self).SetFrictionSlip(gd.Float(value))
}

func (self Instance) SuspensionTravel() float64 {
	return float64(float64(class(self).GetSuspensionTravel()))
}

func (self Instance) SetSuspensionTravel(value float64) {
	class(self).SetSuspensionTravel(gd.Float(value))
}

func (self Instance) SuspensionStiffness() float64 {
	return float64(float64(class(self).GetSuspensionStiffness()))
}

func (self Instance) SetSuspensionStiffness(value float64) {
	class(self).SetSuspensionStiffness(gd.Float(value))
}

func (self Instance) SuspensionMaxForce() float64 {
	return float64(float64(class(self).GetSuspensionMaxForce()))
}

func (self Instance) SetSuspensionMaxForce(value float64) {
	class(self).SetSuspensionMaxForce(gd.Float(value))
}

func (self Instance) DampingCompression() float64 {
	return float64(float64(class(self).GetDampingCompression()))
}

func (self Instance) SetDampingCompression(value float64) {
	class(self).SetDampingCompression(gd.Float(value))
}

func (self Instance) DampingRelaxation() float64 {
	return float64(float64(class(self).GetDampingRelaxation()))
}

func (self Instance) SetDampingRelaxation(value float64) {
	class(self).SetDampingRelaxation(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionRestLength(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_rest_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionRestLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_rest_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionTravel(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionTravel() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionStiffness(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionStiffness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionMaxForce(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_max_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionMaxForce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_max_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDampingCompression(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_damping_compression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDampingCompression() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_damping_compression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDampingRelaxation(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_damping_relaxation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDampingRelaxation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_damping_relaxation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAsTraction(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_use_as_traction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsedAsTraction() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_is_used_as_traction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAsSteering(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_use_as_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsedAsSteering() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_is_used_as_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrictionSlip(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_friction_slip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrictionSlip() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_friction_slip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this wheel is in contact with a surface.
*/
//go:nosplit
func (self class) IsInContact() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_is_in_contact, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the contacting body node if valid in the tree, as [Node3D]. At the moment, [GridMap] is not supported so the node will be always of type [PhysicsBody3D].
Returns [code]null[/code] if the wheel is not in contact with a surface, or the contact body is not a [PhysicsBody3D].
*/
//go:nosplit
func (self class) GetContactBody() gdclass.Node3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_contact_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Node3D{classdb.Node3D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRollInfluence(roll_influence gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, roll_influence)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_roll_influence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRollInfluence() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_roll_influence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is skidding (the wheel has lost grip, e.g. icy terrain), 1.0 means not skidding (the wheel has full grip, e.g. dry asphalt road).
*/
//go:nosplit
func (self class) GetSkidinfo() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_skidinfo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the rotational speed of the wheel in revolutions per minute.
*/
//go:nosplit
func (self class) GetRpm() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_rpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEngineForce(engine_force gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, engine_force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEngineForce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBrake(brake gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, brake)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBrake() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSteering(steering gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, steering)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSteering() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVehicleWheel3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVehicleWheel3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced     { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance  { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {
	classdb.Register("VehicleWheel3D", func(ptr gd.Object) any { return classdb.VehicleWheel3D(ptr) })
}
