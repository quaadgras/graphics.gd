package SoftBody3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/MeshInstance3D"
import "graphics.gd/objects/GeometryInstance3D"
import "graphics.gd/objects/VisualInstance3D"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A deformable 3D physics mesh. Used to create elastic or deformable objects such as cloth, rubber, or other flexible materials.
Additionally, [SoftBody3D] is subject to wind forces defined in [Area3D] (see [member Area3D.wind_source_path], [member Area3D.wind_force_magnitude], and [member Area3D.wind_attenuation_factor]).
[b]Note:[/b] There are many known bugs in [SoftBody3D]. Therefore, it's not recommended to use them for things that can affect gameplay (such as trampolines).
*/
type Instance [1]classdb.SoftBody3D
type Any interface {
	gd.IsClass
	AsSoftBody3D() Instance
}

/*
Returns the internal [RID] used by the [PhysicsServer3D] for this body.
*/
func (self Instance) GetPhysicsRid() Resource.ID {
	return Resource.ID(class(self).GetPhysicsRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionLayerValue(layer_number int, value bool) {
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionLayerValue(layer_number int) bool {
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
func (self Instance) GetCollisionExceptions() gd.Array {
	return gd.Array(class(self).GetCollisionExceptions())
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
func (self Instance) AddCollisionExceptionWith(body objects.Node) {
	class(self).AddCollisionExceptionWith(body)
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
func (self Instance) RemoveCollisionExceptionWith(body objects.Node) {
	class(self).RemoveCollisionExceptionWith(body)
}

/*
Returns local translation of a vertex in the surface array.
*/
func (self Instance) GetPointTransform(point_index int) Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPointTransform(gd.Int(point_index)))
}

/*
Sets the pinned state of a surface vertex. When set to [code]true[/code], the optional [param attachment_path] can define a [Node3D] the pinned vertex will be attached to.
*/
func (self Instance) SetPointPinned(point_index int, pinned bool) {
	class(self).SetPointPinned(gd.Int(point_index), pinned, gd.NewString(string("")).NodePath())
}

/*
Returns [code]true[/code] if vertex is set to pinned.
*/
func (self Instance) IsPointPinned(point_index int) bool {
	return bool(class(self).IsPointPinned(gd.Int(point_index)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SoftBody3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SoftBody3D"))
	return Instance{classdb.SoftBody3D(object)}
}

func (self Instance) CollisionLayer() int {
	return int(int(class(self).GetCollisionLayer()))
}

func (self Instance) SetCollisionLayer(value int) {
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) ParentCollisionIgnore() Path.String {
	return Path.String(class(self).GetParentCollisionIgnore().String())
}

func (self Instance) SetParentCollisionIgnore(value Path.String) {
	class(self).SetParentCollisionIgnore(gd.NewString(string(value)).NodePath())
}

func (self Instance) SimulationPrecision() int {
	return int(int(class(self).GetSimulationPrecision()))
}

func (self Instance) SetSimulationPrecision(value int) {
	class(self).SetSimulationPrecision(gd.Int(value))
}

func (self Instance) TotalMass() Float.X {
	return Float.X(Float.X(class(self).GetTotalMass()))
}

func (self Instance) SetTotalMass(value Float.X) {
	class(self).SetTotalMass(gd.Float(value))
}

func (self Instance) LinearStiffness() Float.X {
	return Float.X(Float.X(class(self).GetLinearStiffness()))
}

func (self Instance) SetLinearStiffness(value Float.X) {
	class(self).SetLinearStiffness(gd.Float(value))
}

func (self Instance) PressureCoefficient() Float.X {
	return Float.X(Float.X(class(self).GetPressureCoefficient()))
}

func (self Instance) SetPressureCoefficient(value Float.X) {
	class(self).SetPressureCoefficient(gd.Float(value))
}

func (self Instance) DampingCoefficient() Float.X {
	return Float.X(Float.X(class(self).GetDampingCoefficient()))
}

func (self Instance) SetDampingCoefficient(value Float.X) {
	class(self).SetDampingCoefficient(gd.Float(value))
}

func (self Instance) DragCoefficient() Float.X {
	return Float.X(Float.X(class(self).GetDragCoefficient()))
}

func (self Instance) SetDragCoefficient(value Float.X) {
	class(self).SetDragCoefficient(gd.Float(value))
}

func (self Instance) RayPickable() bool {
	return bool(class(self).IsRayPickable())
}

func (self Instance) SetRayPickable(value bool) {
	class(self).SetRayPickable(value)
}

func (self Instance) DisableMode() classdb.SoftBody3DDisableMode {
	return classdb.SoftBody3DDisableMode(class(self).GetDisableMode())
}

func (self Instance) SetDisableMode(value classdb.SoftBody3DDisableMode) {
	class(self).SetDisableMode(value)
}

/*
Returns the internal [RID] used by the [PhysicsServer3D] for this body.
*/
//go:nosplit
func (self class) GetPhysicsRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_physics_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionLayer(collision_layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, collision_layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParentCollisionIgnore(parent_collision_ignore gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parent_collision_ignore))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_parent_collision_ignore, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetParentCollisionIgnore() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_parent_collision_ignore, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableMode(mode classdb.SoftBody3DDisableMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_disable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisableMode() classdb.SoftBody3DDisableMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.SoftBody3DDisableMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_disable_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of nodes that were added as collision exceptions for this body.
*/
//go:nosplit
func (self class) GetCollisionExceptions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_collision_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a body to the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) AddCollisionExceptionWith(body objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_add_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a body from the list of bodies that this body can't collide with.
*/
//go:nosplit
func (self class) RemoveCollisionExceptionWith(body objects.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_remove_collision_exception_with, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSimulationPrecision(simulation_precision gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, simulation_precision)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_simulation_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimulationPrecision() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_simulation_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTotalMass(mass gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_total_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTotalMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_total_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearStiffness(linear_stiffness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_stiffness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_linear_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearStiffness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_linear_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressureCoefficient(pressure_coefficient gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pressure_coefficient)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_pressure_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPressureCoefficient() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_pressure_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDampingCoefficient(damping_coefficient gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, damping_coefficient)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_damping_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDampingCoefficient() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_damping_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragCoefficient(drag_coefficient gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, drag_coefficient)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_drag_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDragCoefficient() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_drag_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns local translation of a vertex in the surface array.
*/
//go:nosplit
func (self class) GetPointTransform(point_index gd.Int) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, point_index)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_get_point_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the pinned state of a surface vertex. When set to [code]true[/code], the optional [param attachment_path] can define a [Node3D] the pinned vertex will be attached to.
*/
//go:nosplit
func (self class) SetPointPinned(point_index gd.Int, pinned bool, attachment_path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, point_index)
	callframe.Arg(frame, pinned)
	callframe.Arg(frame, pointers.Get(attachment_path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_point_pinned, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if vertex is set to pinned.
*/
//go:nosplit
func (self class) IsPointPinned(point_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_is_point_pinned, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRayPickable(ray_pickable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ray_pickable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_set_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsRayPickable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SoftBody3D.Bind_is_ray_pickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSoftBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSoftBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMeshInstance3D() MeshInstance3D.Advanced {
	return *((*MeshInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMeshInstance3D() MeshInstance3D.Instance {
	return *((*MeshInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMeshInstance3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMeshInstance3D(), name)
	}
}
func init() {
	classdb.Register("SoftBody3D", func(ptr gd.Object) any { return [1]classdb.SoftBody3D{classdb.SoftBody3D(ptr)} })
}

type DisableMode = classdb.SoftBody3DDisableMode

const (
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], remove from the physics simulation to stop all physics interactions with this [SoftBody3D].
	  Automatically re-added to the physics simulation when the [Node] is processed again.*/
	DisableModeRemove DisableMode = 0
	/*When [member Node.process_mode] is set to [constant Node.PROCESS_MODE_DISABLED], do not affect the physics simulation.*/
	DisableModeKeepActive DisableMode = 1
)
