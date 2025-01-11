// Package VehicleWheel3D provides methods for working with VehicleWheel3D object instances.
package VehicleWheel3D

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
A node used as a child of a [VehicleBody3D] parent to simulate the behavior of one of its wheels. This node also acts as a collider to detect if the wheel is touching a surface.
[b]Note:[/b] This class has known issues and isn't designed to provide realistic 3D vehicle physics. If you want advanced vehicle physics, you may need to write your own physics integration using another [PhysicsBody3D] class.
*/
type Instance [1]gdclass.VehicleWheel3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVehicleWheel3D() Instance
}

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
func (self Instance) GetContactBody() [1]gdclass.Node3D {
	return [1]gdclass.Node3D(class(self).GetContactBody())
}

/*
Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is skidding (the wheel has lost grip, e.g. icy terrain), 1.0 means not skidding (the wheel has full grip, e.g. dry asphalt road).
*/
func (self Instance) GetSkidinfo() Float.X {
	return Float.X(Float.X(class(self).GetSkidinfo()))
}

/*
Returns the rotational speed of the wheel in revolutions per minute.
*/
func (self Instance) GetRpm() Float.X {
	return Float.X(Float.X(class(self).GetRpm()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VehicleWheel3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VehicleWheel3D"))
	casted := Instance{*(*gdclass.VehicleWheel3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) EngineForce() Float.X {
	return Float.X(Float.X(class(self).GetEngineForce()))
}

func (self Instance) SetEngineForce(value Float.X) {
	class(self).SetEngineForce(gd.Float(value))
}

func (self Instance) Brake() Float.X {
	return Float.X(Float.X(class(self).GetBrake()))
}

func (self Instance) SetBrake(value Float.X) {
	class(self).SetBrake(gd.Float(value))
}

func (self Instance) Steering() Float.X {
	return Float.X(Float.X(class(self).GetSteering()))
}

func (self Instance) SetSteering(value Float.X) {
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

func (self Instance) WheelRollInfluence() Float.X {
	return Float.X(Float.X(class(self).GetRollInfluence()))
}

func (self Instance) SetWheelRollInfluence(value Float.X) {
	class(self).SetRollInfluence(gd.Float(value))
}

func (self Instance) WheelRadius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetWheelRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) WheelRestLength() Float.X {
	return Float.X(Float.X(class(self).GetSuspensionRestLength()))
}

func (self Instance) SetWheelRestLength(value Float.X) {
	class(self).SetSuspensionRestLength(gd.Float(value))
}

func (self Instance) WheelFrictionSlip() Float.X {
	return Float.X(Float.X(class(self).GetFrictionSlip()))
}

func (self Instance) SetWheelFrictionSlip(value Float.X) {
	class(self).SetFrictionSlip(gd.Float(value))
}

func (self Instance) SuspensionTravel() Float.X {
	return Float.X(Float.X(class(self).GetSuspensionTravel()))
}

func (self Instance) SetSuspensionTravel(value Float.X) {
	class(self).SetSuspensionTravel(gd.Float(value))
}

func (self Instance) SuspensionStiffness() Float.X {
	return Float.X(Float.X(class(self).GetSuspensionStiffness()))
}

func (self Instance) SetSuspensionStiffness(value Float.X) {
	class(self).SetSuspensionStiffness(gd.Float(value))
}

func (self Instance) SuspensionMaxForce() Float.X {
	return Float.X(Float.X(class(self).GetSuspensionMaxForce()))
}

func (self Instance) SetSuspensionMaxForce(value Float.X) {
	class(self).SetSuspensionMaxForce(gd.Float(value))
}

func (self Instance) DampingCompression() Float.X {
	return Float.X(Float.X(class(self).GetDampingCompression()))
}

func (self Instance) SetDampingCompression(value Float.X) {
	class(self).SetDampingCompression(gd.Float(value))
}

func (self Instance) DampingRelaxation() Float.X {
	return Float.X(Float.X(class(self).GetDampingRelaxation()))
}

func (self Instance) SetDampingRelaxation(value Float.X) {
	class(self).SetDampingRelaxation(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionRestLength(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_rest_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionRestLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_rest_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionTravel(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_travel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionTravel() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_travel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionStiffness(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionStiffness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSuspensionMaxForce(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_suspension_max_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSuspensionMaxForce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_suspension_max_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDampingCompression(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_damping_compression, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDampingCompression() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_damping_compression, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDampingRelaxation(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_damping_relaxation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDampingRelaxation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_damping_relaxation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAsTraction(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_use_as_traction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsedAsTraction() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_is_used_as_traction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAsSteering(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_use_as_steering, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsedAsSteering() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_is_used_as_steering, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrictionSlip(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_friction_slip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrictionSlip() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_friction_slip, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_is_in_contact, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the contacting body node if valid in the tree, as [Node3D]. At the moment, [GridMap] is not supported so the node will be always of type [PhysicsBody3D].
Returns [code]null[/code] if the wheel is not in contact with a surface, or the contact body is not a [PhysicsBody3D].
*/
//go:nosplit
func (self class) GetContactBody() [1]gdclass.Node3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_contact_body, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node3D{gd.PointerMustAssertInstanceID[gdclass.Node3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRollInfluence(roll_influence gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, roll_influence)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_roll_influence, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRollInfluence() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_roll_influence, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_skidinfo, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_rpm, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEngineForce(engine_force gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, engine_force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_engine_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEngineForce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_engine_force, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBrake(brake gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, brake)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_brake, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBrake() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_brake, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSteering(steering gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, steering)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_set_steering, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSteering() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleWheel3D.Bind_get_steering, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("VehicleWheel3D", func(ptr gd.Object) any {
		return [1]gdclass.VehicleWheel3D{*(*gdclass.VehicleWheel3D)(unsafe.Pointer(&ptr))}
	})
}
