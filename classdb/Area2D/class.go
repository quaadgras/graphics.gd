// Package Area2D provides methods for working with Area2D object instances.
package Area2D

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/CollisionObject2D"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode

/*
[Area2D] is a region of 2D space defined by one or multiple [CollisionShape2D] or [CollisionPolygon2D] child nodes. It detects when other [CollisionObject2D]s enter or exit it, and it also keeps track of which collision objects haven't exited it yet (i.e. which one are overlapping it).
This node can also locally alter or override physics parameters (gravity, damping) and route audio to custom audio buses.
[b]Note:[/b] Areas and bodies created with [PhysicsServer2D] might not interact as expected with [Area2D]s, and might not emit signals or track objects correctly.
*/
type Instance [1]gdclass.Area2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsArea2D() Instance
}

/*
Returns a list of intersecting [PhysicsBody2D]s and [TileMap]s. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) GetOverlappingBodies() [][1]gdclass.Node2D { //gd:Area2D.get_overlapping_bodies
	return [][1]gdclass.Node2D(gd.ArrayAs[[][1]gdclass.Node2D](gd.InternalArray(class(self).GetOverlappingBodies())))
}

/*
Returns a list of intersecting [Area2D]s. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) GetOverlappingAreas() [][1]gdclass.Area2D { //gd:Area2D.get_overlapping_areas
	return [][1]gdclass.Area2D(gd.ArrayAs[[][1]gdclass.Area2D](gd.InternalArray(class(self).GetOverlappingAreas())))
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody2D]s or [TileMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) HasOverlappingBodies() bool { //gd:Area2D.has_overlapping_bodies
	return bool(class(self).HasOverlappingBodies())
}

/*
Returns [code]true[/code] if intersecting any [Area2D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) HasOverlappingAreas() bool { //gd:Area2D.has_overlapping_areas
	return bool(class(self).HasOverlappingAreas())
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody2D] or a [TileMap] instance. While TileMaps are not physics bodies themselves, they register their tiles with collision shapes as a virtual physics body.
*/
func (self Instance) OverlapsBody(body [1]gdclass.Node) bool { //gd:Area2D.overlaps_body
	return bool(class(self).OverlapsBody(body))
}

/*
Returns [code]true[/code] if the given [Area2D] intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, the list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Instance) OverlapsArea(area [1]gdclass.Node) bool { //gd:Area2D.overlaps_area
	return bool(class(self).OverlapsArea(area))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Area2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Area2D"))
	casted := Instance{*(*gdclass.Area2D)(unsafe.Pointer(&object))}
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

func (self Instance) GravitySpaceOverride() gdclass.Area2DSpaceOverride {
	return gdclass.Area2DSpaceOverride(class(self).GetGravitySpaceOverrideMode())
}

func (self Instance) SetGravitySpaceOverride(value gdclass.Area2DSpaceOverride) {
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

func (self Instance) GravityPointCenter() Vector2.XY {
	return Vector2.XY(class(self).GetGravityPointCenter())
}

func (self Instance) SetGravityPointCenter(value Vector2.XY) {
	class(self).SetGravityPointCenter(gd.Vector2(value))
}

func (self Instance) GravityDirection() Vector2.XY {
	return Vector2.XY(class(self).GetGravityDirection())
}

func (self Instance) SetGravityDirection(value Vector2.XY) {
	class(self).SetGravityDirection(gd.Vector2(value))
}

func (self Instance) Gravity() Float.X {
	return Float.X(Float.X(class(self).GetGravity()))
}

func (self Instance) SetGravity(value Float.X) {
	class(self).SetGravity(gd.Float(value))
}

func (self Instance) LinearDampSpaceOverride() gdclass.Area2DSpaceOverride {
	return gdclass.Area2DSpaceOverride(class(self).GetLinearDampSpaceOverrideMode())
}

func (self Instance) SetLinearDampSpaceOverride(value gdclass.Area2DSpaceOverride) {
	class(self).SetLinearDampSpaceOverrideMode(value)
}

func (self Instance) LinearDamp() Float.X {
	return Float.X(Float.X(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value Float.X) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Instance) AngularDampSpaceOverride() gdclass.Area2DSpaceOverride {
	return gdclass.Area2DSpaceOverride(class(self).GetAngularDampSpaceOverrideMode())
}

func (self Instance) SetAngularDampSpaceOverride(value gdclass.Area2DSpaceOverride) {
	class(self).SetAngularDampSpaceOverrideMode(value)
}

func (self Instance) AngularDamp() Float.X {
	return Float.X(Float.X(class(self).GetAngularDamp()))
}

func (self Instance) SetAngularDamp(value Float.X) {
	class(self).SetAngularDamp(gd.Float(value))
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
	class(self).SetAudioBusName(String.Name(String.New(value)))
}

//go:nosplit
func (self class) SetGravitySpaceOverrideMode(space_override_mode gdclass.Area2DSpaceOverride) { //gd:Area2D.set_gravity_space_override_mode
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravitySpaceOverrideMode() gdclass.Area2DSpaceOverride { //gd:Area2D.get_gravity_space_override_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Area2DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityIsPoint(enable bool) { //gd:Area2D.set_gravity_is_point
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_is_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGravityAPoint() bool { //gd:Area2D.is_gravity_a_point
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_gravity_a_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityPointUnitDistance(distance_scale gd.Float) { //gd:Area2D.set_gravity_point_unit_distance
	var frame = callframe.New()
	callframe.Arg(frame, distance_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityPointUnitDistance() gd.Float { //gd:Area2D.get_gravity_point_unit_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityPointCenter(center gd.Vector2) { //gd:Area2D.set_gravity_point_center
	var frame = callframe.New()
	callframe.Arg(frame, center)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityPointCenter() gd.Vector2 { //gd:Area2D.get_gravity_point_center
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityDirection(direction gd.Vector2) { //gd:Area2D.set_gravity_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityDirection() gd.Vector2 { //gd:Area2D.get_gravity_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(gravity gd.Float) { //gd:Area2D.set_gravity
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravity() gd.Float { //gd:Area2D.get_gravity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampSpaceOverrideMode(space_override_mode gdclass.Area2DSpaceOverride) { //gd:Area2D.set_linear_damp_space_override_mode
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampSpaceOverrideMode() gdclass.Area2DSpaceOverride { //gd:Area2D.get_linear_damp_space_override_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Area2DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampSpaceOverrideMode(space_override_mode gdclass.Area2DSpaceOverride) { //gd:Area2D.set_angular_damp_space_override_mode
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampSpaceOverrideMode() gdclass.Area2DSpaceOverride { //gd:Area2D.get_angular_damp_space_override_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Area2DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float) { //gd:Area2D.set_linear_damp
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() gd.Float { //gd:Area2D.get_linear_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float) { //gd:Area2D.set_angular_damp
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() gd.Float { //gd:Area2D.get_angular_damp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPriority(priority gd.Int) { //gd:Area2D.set_priority
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPriority() gd.Int { //gd:Area2D.get_priority
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMonitoring(enable bool) { //gd:Area2D.set_monitoring
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_monitoring, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMonitoring() bool { //gd:Area2D.is_monitoring
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_monitoring, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMonitorable(enable bool) { //gd:Area2D.set_monitorable
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMonitorable() bool { //gd:Area2D.is_monitorable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_monitorable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of intersecting [PhysicsBody2D]s and [TileMap]s. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingBodies() Array.Contains[[1]gdclass.Node2D] { //gd:Area2D.get_overlapping_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Node2D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a list of intersecting [Area2D]s. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingAreas() Array.Contains[[1]gdclass.Area2D] { //gd:Area2D.get_overlapping_areas
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Area2D]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody2D]s or [TileMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingBodies() bool { //gd:Area2D.has_overlapping_bodies
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_has_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if intersecting any [Area2D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingAreas() bool { //gd:Area2D.has_overlapping_areas
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_has_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody2D] or a [TileMap] instance. While TileMaps are not physics bodies themselves, they register their tiles with collision shapes as a virtual physics body.
*/
//go:nosplit
func (self class) OverlapsBody(body [1]gdclass.Node) bool { //gd:Area2D.overlaps_body
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_overlaps_body, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [Area2D] intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, the list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) OverlapsArea(area [1]gdclass.Node) bool { //gd:Area2D.overlaps_area
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(area[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_overlaps_area, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioBusName(name String.Name) { //gd:Area2D.set_audio_bus_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAudioBusName() String.Name { //gd:Area2D.get_audio_bus_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioBusOverride(enable bool) { //gd:Area2D.set_audio_bus_override
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_audio_bus_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOverridingAudioBus() bool { //gd:Area2D.is_overriding_audio_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_overriding_audio_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnBodyShapeEntered(cb func(body_rid RID.Any, body [1]gdclass.Node2D, body_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyShapeExited(cb func(body_rid RID.Any, body [1]gdclass.Node2D, body_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyEntered(cb func(body [1]gdclass.Node2D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyExited(cb func(body [1]gdclass.Node2D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("body_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaShapeEntered(cb func(area_rid RID.Any, area [1]gdclass.Area2D, area_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaShapeExited(cb func(area_rid RID.Any, area [1]gdclass.Area2D, area_shape_index int, local_shape_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaEntered(cb func(area [1]gdclass.Area2D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaExited(cb func(area [1]gdclass.Area2D)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("area_exited"), gd.NewCallable(cb), 0)
}

func (self class) AsArea2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsArea2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject2D() CollisionObject2D.Advanced {
	return *((*CollisionObject2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject2D() CollisionObject2D.Instance {
	return *((*CollisionObject2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CollisionObject2D.Advanced(self.AsCollisionObject2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CollisionObject2D.Instance(self.AsCollisionObject2D()), name)
	}
}
func init() {
	gdclass.Register("Area2D", func(ptr gd.Object) any { return [1]gdclass.Area2D{*(*gdclass.Area2D)(unsafe.Pointer(&ptr))} })
}

type SpaceOverride = gdclass.Area2DSpaceOverride //gd:Area2D.SpaceOverride

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
