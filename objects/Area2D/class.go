package Area2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/CollisionObject2D"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[Area2D] is a region of 2D space defined by one or multiple [CollisionShape2D] or [CollisionPolygon2D] child nodes. It detects when other [CollisionObject2D]s enter or exit it, and it also keeps track of which collision objects haven't exited it yet (i.e. which one are overlapping it).
This node can also locally alter or override physics parameters (gravity, damping) and route audio to custom audio buses.
[b]Note:[/b] Areas and bodies created with [PhysicsServer2D] might not interact as expected with [Area2D]s, and might not emit signals or track objects correctly.
*/
type Instance [1]classdb.Area2D

/*
Returns a list of intersecting [PhysicsBody2D]s and [TileMap]s. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) GetOverlappingBodies() gd.Array {
	return gd.Array(class(self).GetOverlappingBodies())
}

/*
Returns a list of intersecting [Area2D]s. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) GetOverlappingAreas() gd.Array {
	return gd.Array(class(self).GetOverlappingAreas())
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody2D]s or [TileMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) HasOverlappingBodies() bool {
	return bool(class(self).HasOverlappingBodies())
}

/*
Returns [code]true[/code] if intersecting any [Area2D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Instance) HasOverlappingAreas() bool {
	return bool(class(self).HasOverlappingAreas())
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody2D] or a [TileMap] instance. While TileMaps are not physics bodies themselves, they register their tiles with collision shapes as a virtual physics body.
*/
func (self Instance) OverlapsBody(body objects.Node) bool {
	return bool(class(self).OverlapsBody(body))
}

/*
Returns [code]true[/code] if the given [Area2D] intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, the list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Instance) OverlapsArea(area objects.Node) bool {
	return bool(class(self).OverlapsArea(area))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Area2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Area2D"))
	return Instance{classdb.Area2D(object)}
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

func (self Instance) GravitySpaceOverride() classdb.Area2DSpaceOverride {
	return classdb.Area2DSpaceOverride(class(self).GetGravitySpaceOverrideMode())
}

func (self Instance) SetGravitySpaceOverride(value classdb.Area2DSpaceOverride) {
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

func (self Instance) LinearDampSpaceOverride() classdb.Area2DSpaceOverride {
	return classdb.Area2DSpaceOverride(class(self).GetLinearDampSpaceOverrideMode())
}

func (self Instance) SetLinearDampSpaceOverride(value classdb.Area2DSpaceOverride) {
	class(self).SetLinearDampSpaceOverrideMode(value)
}

func (self Instance) LinearDamp() Float.X {
	return Float.X(Float.X(class(self).GetLinearDamp()))
}

func (self Instance) SetLinearDamp(value Float.X) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Instance) AngularDampSpaceOverride() classdb.Area2DSpaceOverride {
	return classdb.Area2DSpaceOverride(class(self).GetAngularDampSpaceOverrideMode())
}

func (self Instance) SetAngularDampSpaceOverride(value classdb.Area2DSpaceOverride) {
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
	class(self).SetAudioBusName(gd.NewStringName(value))
}

//go:nosplit
func (self class) SetGravitySpaceOverrideMode(space_override_mode classdb.Area2DSpaceOverride) {
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravitySpaceOverrideMode() classdb.Area2DSpaceOverride {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area2DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityIsPoint(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_is_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsGravityAPoint() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_gravity_a_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityPointUnitDistance(distance_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityPointUnitDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityPointCenter(center gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, center)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityPointCenter() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravityDirection(direction gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravityDirection() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(gravity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDampSpaceOverrideMode(space_override_mode classdb.Area2DSpaceOverride) {
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDampSpaceOverrideMode() classdb.Area2DSpaceOverride {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area2DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDampSpaceOverrideMode(space_override_mode classdb.Area2DSpaceOverride) {
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDampSpaceOverrideMode() classdb.Area2DSpaceOverride {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area2DSpaceOverride](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMonitoring(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_monitoring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMonitoring() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_monitoring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMonitorable(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMonitorable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of intersecting [PhysicsBody2D]s and [TileMap]s. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingBodies() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of intersecting [Area2D]s. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingAreas() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody2D]s or [TileMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingBodies() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_has_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if intersecting any [Area2D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingAreas() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_has_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) OverlapsBody(body objects.Node) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_overlaps_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [Area2D] intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, the list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) OverlapsArea(area objects.Node) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(area[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_overlaps_area, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioBusName(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAudioBusName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_get_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioBusOverride(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_set_audio_bus_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsOverridingAudioBus() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Area2D.Bind_is_overriding_audio_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnBodyShapeEntered(cb func(body_rid Resource.ID, body objects.Node2D, body_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("body_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyShapeExited(cb func(body_rid Resource.ID, body objects.Node2D, body_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("body_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyEntered(cb func(body objects.Node2D)) {
	self[0].AsObject().Connect(gd.NewStringName("body_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBodyExited(cb func(body objects.Node2D)) {
	self[0].AsObject().Connect(gd.NewStringName("body_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaShapeEntered(cb func(area_rid Resource.ID, area objects.Area2D, area_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("area_shape_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaShapeExited(cb func(area_rid Resource.ID, area objects.Area2D, area_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("area_shape_exited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaEntered(cb func(area objects.Area2D)) {
	self[0].AsObject().Connect(gd.NewStringName("area_entered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAreaExited(cb func(area objects.Area2D)) {
	self[0].AsObject().Connect(gd.NewStringName("area_exited"), gd.NewCallable(cb), 0)
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
		return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}
func init() { classdb.Register("Area2D", func(ptr gd.Object) any { return classdb.Area2D(ptr) }) }

type SpaceOverride = classdb.Area2DSpaceOverride

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
