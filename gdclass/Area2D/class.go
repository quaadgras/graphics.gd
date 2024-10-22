package Area2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CollisionObject2D"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[Area2D] is a region of 2D space defined by one or multiple [CollisionShape2D] or [CollisionPolygon2D] child nodes. It detects when other [CollisionObject2D]s enter or exit it, and it also keeps track of which collision objects haven't exited it yet (i.e. which one are overlapping it).
This node can also locally alter or override physics parameters (gravity, damping) and route audio to custom audio buses.
[b]Note:[/b] Areas and bodies created with [PhysicsServer2D] might not interact as expected with [Area2D]s, and might not emit signals or track objects correctly.

*/
type Go [1]classdb.Area2D

/*
Returns a list of intersecting [PhysicsBody2D]s and [TileMap]s. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) GetOverlappingBodies() gd.ArrayOf[gdclass.Node2D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Node2D](class(self).GetOverlappingBodies(gc))
}

/*
Returns a list of intersecting [Area2D]s. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) GetOverlappingAreas() gd.ArrayOf[gdclass.Area2D] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Area2D](class(self).GetOverlappingAreas(gc))
}

/*
Returns [code]true[/code] if intersecting any [PhysicsBody2D]s or [TileMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) HasOverlappingBodies() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasOverlappingBodies())
}

/*
Returns [code]true[/code] if intersecting any [Area2D]s, otherwise returns [code]false[/code]. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping areas is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
func (self Go) HasOverlappingAreas() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasOverlappingAreas())
}

/*
Returns [code]true[/code] if the given physics body intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
The [param body] argument can either be a [PhysicsBody2D] or a [TileMap] instance. While TileMaps are not physics bodies themselves, they register their tiles with collision shapes as a virtual physics body.
*/
func (self Go) OverlapsBody(body gdclass.Node) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).OverlapsBody(body))
}

/*
Returns [code]true[/code] if the given [Area2D] intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, the list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Go) OverlapsArea(area gdclass.Node) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).OverlapsArea(area))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Area2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Area2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Monitoring() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMonitoring())
}

func (self Go) SetMonitoring(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMonitoring(value)
}

func (self Go) Monitorable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMonitorable())
}

func (self Go) SetMonitorable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMonitorable(value)
}

func (self Go) Priority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPriority()))
}

func (self Go) SetPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPriority(gd.Int(value))
}

func (self Go) GravitySpaceOverride() classdb.Area2DSpaceOverride {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Area2DSpaceOverride(class(self).GetGravitySpaceOverrideMode())
}

func (self Go) SetGravitySpaceOverride(value classdb.Area2DSpaceOverride) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravitySpaceOverrideMode(value)
}

func (self Go) GravityPoint() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsGravityAPoint())
}

func (self Go) SetGravityPoint(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityIsPoint(value)
}

func (self Go) GravityPointUnitDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGravityPointUnitDistance()))
}

func (self Go) SetGravityPointUnitDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityPointUnitDistance(gd.Float(value))
}

func (self Go) GravityPointCenter() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetGravityPointCenter())
}

func (self Go) SetGravityPointCenter(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityPointCenter(value)
}

func (self Go) GravityDirection() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetGravityDirection())
}

func (self Go) SetGravityDirection(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravityDirection(value)
}

func (self Go) Gravity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetGravity()))
}

func (self Go) SetGravity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGravity(gd.Float(value))
}

func (self Go) LinearDampSpaceOverride() classdb.Area2DSpaceOverride {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Area2DSpaceOverride(class(self).GetLinearDampSpaceOverrideMode())
}

func (self Go) SetLinearDampSpaceOverride(value classdb.Area2DSpaceOverride) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearDampSpaceOverrideMode(value)
}

func (self Go) LinearDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLinearDamp()))
}

func (self Go) SetLinearDamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Go) AngularDampSpaceOverride() classdb.Area2DSpaceOverride {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Area2DSpaceOverride(class(self).GetAngularDampSpaceOverrideMode())
}

func (self Go) SetAngularDampSpaceOverride(value classdb.Area2DSpaceOverride) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularDampSpaceOverrideMode(value)
}

func (self Go) AngularDamp() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAngularDamp()))
}

func (self Go) SetAngularDamp(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Go) AudioBusOverride() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsOverridingAudioBus())
}

func (self Go) SetAudioBusOverride(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAudioBusOverride(value)
}

func (self Go) AudioBusName() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAudioBusName(gc).String())
}

func (self Go) SetAudioBusName(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAudioBusName(gc.StringName(value))
}

//go:nosplit
func (self class) SetGravitySpaceOverrideMode(space_override_mode classdb.Area2DSpaceOverride)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravitySpaceOverrideMode() classdb.Area2DSpaceOverride {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area2DSpaceOverride](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_gravity_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityIsPoint(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_gravity_is_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGravityAPoint() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_is_gravity_a_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityPointUnitDistance(distance_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, distance_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityPointUnitDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_gravity_point_unit_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityPointCenter(center gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, center)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityPointCenter() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_gravity_point_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityDirection(direction gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityDirection() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_gravity_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravity(gravity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDampSpaceOverrideMode(space_override_mode classdb.Area2DSpaceOverride)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDampSpaceOverrideMode() classdb.Area2DSpaceOverride {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area2DSpaceOverride](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_linear_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDampSpaceOverrideMode(space_override_mode classdb.Area2DSpaceOverride)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, space_override_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDampSpaceOverrideMode() classdb.Area2DSpaceOverride {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Area2DSpaceOverride](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_angular_damp_space_override_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMonitoring(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_monitoring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMonitoring() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_is_monitoring, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMonitorable(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMonitorable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_is_monitorable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of intersecting [PhysicsBody2D]s and [TileMap]s. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingBodies(ctx gd.Lifetime) gd.ArrayOf[gdclass.Node2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Node2D](ret)
}
/*
Returns a list of intersecting [Area2D]s. The overlapping area's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) this list is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) GetOverlappingAreas(ctx gd.Lifetime) gd.ArrayOf[gdclass.Area2D] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Area2D](ret)
}
/*
Returns [code]true[/code] if intersecting any [PhysicsBody2D]s or [TileMap]s, otherwise returns [code]false[/code]. The overlapping body's [member CollisionObject2D.collision_layer] must be part of this area's [member CollisionObject2D.collision_mask] in order to be detected.
For performance reasons (collisions are all processed at the same time) the list of overlapping bodies is modified once during the physics step, not immediately after objects are moved. Consider using signals instead.
*/
//go:nosplit
func (self class) HasOverlappingBodies() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_has_overlapping_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_has_overlapping_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) OverlapsBody(body gdclass.Node) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(body[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_overlaps_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [Area2D] intersects or overlaps this [Area2D], [code]false[/code] otherwise.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, the list of overlaps is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) OverlapsArea(area gdclass.Node) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(area[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_overlaps_area, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAudioBusName(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAudioBusName(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_get_audio_bus_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAudioBusOverride(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_set_audio_bus_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOverridingAudioBus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Area2D.Bind_is_overriding_audio_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnBodyShapeEntered(cb func(body_rid gd.RID, body gdclass.Node2D, body_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_shape_entered"), gc.Callable(cb), 0)
}


func (self Go) OnBodyShapeExited(cb func(body_rid gd.RID, body gdclass.Node2D, body_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_shape_exited"), gc.Callable(cb), 0)
}


func (self Go) OnBodyEntered(cb func(body gdclass.Node2D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_entered"), gc.Callable(cb), 0)
}


func (self Go) OnBodyExited(cb func(body gdclass.Node2D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("body_exited"), gc.Callable(cb), 0)
}


func (self Go) OnAreaShapeEntered(cb func(area_rid gd.RID, area gdclass.Area2D, area_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_shape_entered"), gc.Callable(cb), 0)
}


func (self Go) OnAreaShapeExited(cb func(area_rid gd.RID, area gdclass.Area2D, area_shape_index int, local_shape_index int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_shape_exited"), gc.Callable(cb), 0)
}


func (self Go) OnAreaEntered(cb func(area gdclass.Area2D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_entered"), gc.Callable(cb), 0)
}


func (self Go) OnAreaExited(cb func(area gdclass.Area2D)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("area_exited"), gc.Callable(cb), 0)
}


func (self class) AsArea2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsArea2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject2D() CollisionObject2D.GD { return *((*CollisionObject2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject2D() CollisionObject2D.Go { return *((*CollisionObject2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCollisionObject2D(), name)
	}
}
func init() {classdb.Register("Area2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
