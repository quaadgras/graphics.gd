// Package Area3D provides methods for working with Area3D object instances.
package Area3D

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
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/RID"

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

/*
[Area3D] is a region of 3D space defined by one or multiple [CollisionShape3D] or [CollisionPolygon3D] child nodes. It detects when other [CollisionObject3D]s enter or exit it, and it also keeps track of which collision objects haven't exited it yet (i.e. which one are overlapping it).
This node can also locally alter or override physics parameters (gravity, damping) and route audio to custom audio buses.
[b]Note:[/b] Areas and bodies created with [PhysicsServer3D] might not interact as expected with [Area3D]s, and might not emit signals or track objects correctly.
[b]Warning:[/b] Using a [ConcavePolygonShape3D] inside a [CollisionShape3D] child of this node (created e.g. by using the [b]Create Trimesh Collision Sibling[/b] option in the [b]Mesh[/b] menu that appears when selecting a [MeshInstance3D] node) may give unexpected results, since this collision shape is hollow. If this is not desired, it has to be split into multiple [ConvexPolygonShape3D]s or primitive shapes like [BoxShape3D], or in some cases it may be replaceable by a [CollisionPolygon3D].
*/
type Instance [1]gdclass.Area3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsArea3D() Instance
}

/*
Returns a list of intersecting [PhysicsBody3D]s and [GridMap]s. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) GetOverlappingBodies() [][1]gdclass.Node3D { //gd:Area3D.get_overlapping_bodies
	return [][1]gdclass.Node3D(gd.ArrayAs[[][1]gdclass.Node3D](gd.InternalArray(class(self).GetOverlappingBodies())))
}

/*
Returns a list of intersecting [Area3D]s. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) GetOverlappingAreas() [][1]gdclass.Area3D { //gd:Area3D.get_overlapping_areas
	return [][1]gdclass.Area3D(gd.ArrayAs[[][1]gdclass.Area3D](gd.InternalArray(class(self).GetOverlappingAreas())))
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody3D]s or [GridMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) HasOverlappingBodies() bool { //gd:Area3D.has_overlapping_bodies
	return bool(class(self).HasOverlappingBodies())
}

/*
Returns [code]true[/code] if intersecting any [Area3D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) HasOverlappingAreas() bool { //gd:Area3D.has_overlapping_areas
	return bool(class(self).HasOverlappingAreas())
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody3D] or a [GridMap] instance. While GridMaps are not physics body themselves, they register their tiles with collision shapes as a virtual physics body.
*/
func (self Instance) OverlapsBody(body [1]gdclass.Node) bool { //gd:Area3D.overlaps_body
	return bool(class(self).OverlapsBody(body))
}

/*
Returns [code]true[/code] if the given [Area3D] intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Instance) OverlapsArea(area [1]gdclass.Node) bool { //gd:Area3D.overlaps_area
	return bool(class(self).OverlapsArea(area))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Area3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Area3D"))
	casted := Instance{*(*gdclass.Area3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Monitoring() bool {
	return bool(class(self).IsMonitoring())
}

func (self Instance) SetMonitoring(value bool) {
	class(self).SetMonitoring(value)
}

func (self Instance) Monitorable() bool {
	return bool(class(self).IsMonitorable())
}

func (self Instance) SetMonitorable(value bool) {
	class(self).SetMonitorable(value)
}

func (self Instance) Priority() int {
	return int(int(class(self).GetPriority()))
}

func (self Instance) SetPriority(value int) {
	class(self).SetPriority(gd.Int(value))
}

func (self Instance) GravitySpaceOverride() gdclass.Area3DSpaceOverride {
	return gdclass.Area3DSpaceOverride(class(self).GetGravitySpaceOverrideMode())
}

func (self Instance) SetGravitySpaceOverride(value gdclass.Area3DSpaceOverride) {
	class(self).SetGravitySpaceOverrideMode(value)
}

func (self Instance) GravityPoint() bool {
	return bool(class(self).IsGravityAPoint())
}

func (self Instance) SetGravityPoint(value bool) {
	class(self).SetGravityIsPoint(value)
}

func (self Instance) GravityPointUnitDistance() Float.X {
	return Float.X(Float.X(class(self).GetGravityPointUnitDistance()))
}

func (self Instance) SetGravityPointUnitDistance(value Float.X) {
	class(self).SetGravityPointUnitDistance(gd.Float(value))
}

func (self Instance) GravityPointCenter() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetGravityPointCenter())
}

func (self Instance) SetGravityPointCenter(value Vector3.XYZ) {
	class(self).SetGravityPointCenter(gd.Vector3(value))
}

func (self Instance) GravityDirection() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetGravityDirection())
}

func (self Instance) SetGravityDirection(value Vector3.XYZ) {
	class(self).SetGravityDirection(gd.Vector3(value))
}

func (self Instance) Gravity() Float.X {
	return Float.X(Float.X(class(self).GetGravity()))
}

func (self Instance) SetGravity(value Float.X) {
	class(self).SetGravity(gd.Float(value))
}

func (self Instance) LinearDampSpaceOverride() gdclass.Area3DSpaceOverride {
	return gdclass.Area3DSpaceOverride(class(self).GetLinearDampSpaceOverrideMode())
}

func (self Instance) SetLinearDampSpaceOverride(value gdclass.Area3DSpaceOverride) {
	class(self).SetLinearDampSpaceOverrideMode(value)
}

func (self Instance) LinearDamp() Float.X {
	return Float.X(Float.X(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value Float.X) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Instance) AngularDampSpaceOverride() gdclass.Area3DSpaceOverride {
	return gdclass.Area3DSpaceOverride(class(self).GetAngularDampSpaceOverrideMode())
}

func (self Instance) SetAngularDampSpaceOverride(value gdclass.Area3DSpaceOverride) {
	class(self).SetAngularDampSpaceOverrideMode(value)
}

func (self Instance) AngularDamp() Float.X {
	return Float.X(Float.X(class(self).GetAngularDamp()))
}

func (self Instance) SetAngularDamp(value Float.X) {
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Instance) WindForceMagnitude() Float.X {
	return Float.X(Float.X(class(self).GetWindForceMagnitude()))
}

func (self Instance) SetWindForceMagnitude(value Float.X) {
	class(self).SetWindForceMagnitude(gd.Float(value))
}

func (self Instance) WindAttenuationFactor() Float.X {
	return Float.X(Float.X(class(self).GetWindAttenuationFactor()))
}

func (self Instance) SetWindAttenuationFactor(value Float.X) {
	class(self).SetWindAttenuationFactor(gd.Float(value))
}

func (self Instance) WindSourcePath() NodePath.String {
	return NodePath.String(class(self).GetWindSourcePath().String())
}

func (self Instance) SetWindSourcePath(value NodePath.String) {
	class(self).SetWindSourcePath(gd.NewString(string(value)).NodePath())
}

func (self Instance) AudioBusOverride() bool {
	return bool(class(self).IsOverridingAudioBus())
}

func (self Instance) SetAudioBusOverride(value bool) {
	class(self).SetAudioBusOverride(value)
}

func (self Instance) AudioBusName() string {
	return string(class(self).GetAudioBusName().String())
}

func (self Instance) SetAudioBusName(value string) {
	class(self).SetAudioBusName(gd.NewStringName(value))
}

func (self Instance) ReverbBusEnabled() bool {
	return bool(class(self).IsUsingReverbBus())
}

func (self Instance) SetReverbBusEnabled(value bool) {
	class(self).SetUseReverbBus(value)
}

func (self Instance) ReverbBusName() string {
	return string(class(self).GetReverbBusName().String())
}

func (self Instance) SetReverbBusName(value string) {
	class(self).SetReverbBusName(gd.NewStringName(value))
}

func (self Instance) ReverbBusAmount() Float.X {
	return Float.X(Float.X(class(self).GetReverbAmount()))
}

func (self Instance) SetReverbBusAmount(value Float.X) {
	class(self).SetReverbAmount(gd.Float(value))
}

func (self Instance) ReverbBusUniformity() Float.X {
	return Float.X(Float.X(class(self).GetReverbUniformity()))
}

func (self Instance) SetReverbBusUniformity(value Float.X) {
	class(self).SetReverbUniformity(gd.Float(value))
}

//go:nosplit
func (self class) SetGravitySpaceOverrideMode(space_override_mode gdclass.Area3DSpaceOverride) { //gd:Area3D.set_gravity_space_override_mode
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravitySpaceOverrideMode() gdclass.Area3DSpaceOverride { //gd:Area3D.get_gravity_space_override_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Area3DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityIsPoint(enable bool) { //gd:Area3D.set_gravity_is_point
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_gravity_is_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGravityAPoint() bool { //gd:Area3D.is_gravity_a_point
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_is_gravity_a_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityPointUnitDistance(distance_scale gd.Float) { //gd:Area3D.set_gravity_point_unit_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityPointUnitDistance() gd.Float { //gd:Area3D.get_gravity_point_unit_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityPointCenter(center gd.Vector3) { //gd:Area3D.set_gravity_point_center
	var frame = callframe.New()
	callframe.Arg(frame, center)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityPointCenter() gd.Vector3 { //gd:Area3D.get_gravity_point_center
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityDirection(direction gd.Vector3) { //gd:Area3D.set_gravity_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityDirection() gd.Vector3 { //gd:Area3D.get_gravity_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(gravity gd.Float) { //gd:Area3D.set_gravity
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravity() gd.Float { //gd:Area3D.get_gravity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampSpaceOverrideMode(space_override_mode gdclass.Area3DSpaceOverride) { //gd:Area3D.set_linear_damp_space_override_mode
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampSpaceOverrideMode() gdclass.Area3DSpaceOverride { //gd:Area3D.get_linear_damp_space_override_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Area3DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampSpaceOverrideMode(space_override_mode gdclass.Area3DSpaceOverride) { //gd:Area3D.set_angular_damp_space_override_mode
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampSpaceOverrideMode() gdclass.Area3DSpaceOverride { //gd:Area3D.get_angular_damp_space_override_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Area3DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float) { //gd:Area3D.set_angular_damp
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() gd.Float { //gd:Area3D.get_angular_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float) { //gd:Area3D.set_linear_damp
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() gd.Float { //gd:Area3D.get_linear_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPriority(priority gd.Int) { //gd:Area3D.set_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPriority() gd.Int { //gd:Area3D.get_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWindForceMagnitude(wind_force_magnitude gd.Float) { //gd:Area3D.set_wind_force_magnitude
	var frame = callframe.New()
	callframe.Arg(frame, wind_force_magnitude)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_wind_force_magnitude, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWindForceMagnitude() gd.Float { //gd:Area3D.get_wind_force_magnitude
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_wind_force_magnitude, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWindAttenuationFactor(wind_attenuation_factor gd.Float) { //gd:Area3D.set_wind_attenuation_factor
	var frame = callframe.New()
	callframe.Arg(frame, wind_attenuation_factor)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_wind_attenuation_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWindAttenuationFactor() gd.Float { //gd:Area3D.get_wind_attenuation_factor
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_wind_attenuation_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWindSourcePath(wind_source_path gd.NodePath) { //gd:Area3D.set_wind_source_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(wind_source_path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_wind_source_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWindSourcePath() gd.NodePath { //gd:Area3D.get_wind_source_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_wind_source_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMonitorable(enable bool) { //gd:Area3D.set_monitorable
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMonitorable() bool { //gd:Area3D.is_monitorable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_is_monitorable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMonitoring(enable bool) { //gd:Area3D.set_monitoring
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_monitoring, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMonitoring() bool { //gd:Area3D.is_monitoring
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_is_monitoring, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of intersecting [PhysicsBody3D]s and [GridMap]s. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingBodies() Array.Contains[[1]gdclass.Node3D] { //gd:Area3D.get_overlapping_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Node3D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a list of intersecting [Area3D]s. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingAreas() Array.Contains[[1]gdclass.Area3D] { //gd:Area3D.get_overlapping_areas
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Area3D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody3D]s or [GridMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingBodies() bool { //gd:Area3D.has_overlapping_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_has_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if intersecting any [Area3D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject3D.collision_layer] must be part of this area's [member CollisionObject3D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingAreas() bool { //gd:Area3D.has_overlapping_areas
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_has_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody3D] or a [GridMap] instance. While GridMaps are not physics body themselves, they register their tiles with collision shapes as a virtual physics body.
*/
//go:nosplit
func (self class) OverlapsBody(body [1]gdclass.Node) bool { //gd:Area3D.overlaps_body
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_overlaps_body, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [Area3D] intersects or overlaps this [Area3D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) OverlapsArea(area [1]gdclass.Node) bool { //gd:Area3D.overlaps_area
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(area[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_overlaps_area, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioBusOverride(enable bool) { //gd:Area3D.set_audio_bus_override
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_audio_bus_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOverridingAudioBus() bool { //gd:Area3D.is_overriding_audio_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_is_overriding_audio_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioBusName(name gd.StringName) { //gd:Area3D.set_audio_bus_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAudioBusName() gd.StringName { //gd:Area3D.get_audio_bus_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseReverbBus(enable bool) { //gd:Area3D.set_use_reverb_bus
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_use_reverb_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingReverbBus() bool { //gd:Area3D.is_using_reverb_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_is_using_reverb_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReverbBusName(name gd.StringName) { //gd:Area3D.set_reverb_bus_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_reverb_bus_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReverbBusName() gd.StringName { //gd:Area3D.get_reverb_bus_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_reverb_bus_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReverbAmount(amount gd.Float) { //gd:Area3D.set_reverb_amount
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_reverb_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReverbAmount() gd.Float { //gd:Area3D.get_reverb_amount
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_reverb_amount, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReverbUniformity(amount gd.Float) { //gd:Area3D.set_reverb_uniformity
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_set_reverb_uniformity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReverbUniformity() gd.Float { //gd:Area3D.get_reverb_uniformity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area3D.Bind_get_reverb_uniformity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnBodyShapeEntered(cb func(body_rid RID.Any, body [1]gdclass.Node3D, body_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyShapeExited(cb func(body_rid RID.Any, body [1]gdclass.Node3D, body_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyEntered(cb func(body [1]gdclass.Node3D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyExited(cb func(body [1]gdclass.Node3D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaShapeEntered(cb func(area_rid RID.Any, area [1]gdclass.Area3D, area_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaShapeExited(cb func(area_rid RID.Any, area [1]gdclass.Area3D, area_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaEntered(cb func(area [1]gdclass.Area3D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaExited(cb func(area [1]gdclass.Area3D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsArea3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsArea3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(CollisionObject3D.Advanced(self.AsCollisionObject3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CollisionObject3D.Instance(self.AsCollisionObject3D()), name)
	}
}
func init() {
	gdclass.Register("Area3D", func(ptr gd.Object) any { return [1]gdclass.Area3D{*(*gdclass.Area3D)(unsafe.Pointer(&ptr))} })
}

type SpaceOverride = gdclass.Area3DSpaceOverride //gd:Area3D.SpaceOverride

const (
	/*This area does not affect gravity/damping.*/
	SpaceOverrideDisabled SpaceOverride = 0
	/*This area adds its gravity/damping values to whatever has been calculated so far (in [member priority] order).*/
	SpaceOverrideCombine SpaceOverride = 1
	/*This area adds its gravity/damping values to whatever has been calculated so far (in [member priority] order), ignoring any lower priority areas.*/
	SpaceOverrideCombineReplace SpaceOverride = 2
	/*This area replaces any gravity/damping, even the defaults, ignoring any lower priority areas.*/
	SpaceOverrideReplace SpaceOverride = 3
	/*This area replaces any gravity/damping calculated so far (in [member priority] order), but keeps calculating the rest of the areas.*/
	SpaceOverrideReplaceCombine SpaceOverride = 4
)
