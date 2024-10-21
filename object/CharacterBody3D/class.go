package CharacterBody3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
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
[CharacterBody3D] is a specialized class for physics bodies that are meant to be user-controlled. They are not affected by physics at all, but they affect other physics bodies in their path. They are mainly used to provide high-level API to move objects with wall and slope detection ([method move_and_slide] method) in addition to the general collision detection provided by [method PhysicsBody3D.move_and_collide]. This makes it useful for highly configurable physics bodies that must move in specific ways and collide with the world, as is often the case with user-controlled characters.
For game objects that don't require complex movement or collision detection, such as moving platforms, [AnimatableBody3D] is simpler to configure.

*/
type Simple [1]classdb.CharacterBody3D
func (self Simple) MoveAndSlide() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).MoveAndSlide())
}
func (self Simple) ApplyFloorSnap() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ApplyFloorSnap()
}
func (self Simple) SetVelocity(velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVelocity(velocity)
}
func (self Simple) GetVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetVelocity())
}
func (self Simple) SetSafeMargin(margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSafeMargin(gd.Float(margin))
}
func (self Simple) GetSafeMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSafeMargin()))
}
func (self Simple) IsFloorStopOnSlopeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFloorStopOnSlopeEnabled())
}
func (self Simple) SetFloorStopOnSlopeEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFloorStopOnSlopeEnabled(enabled)
}
func (self Simple) SetFloorConstantSpeedEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFloorConstantSpeedEnabled(enabled)
}
func (self Simple) IsFloorConstantSpeedEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFloorConstantSpeedEnabled())
}
func (self Simple) SetFloorBlockOnWallEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFloorBlockOnWallEnabled(enabled)
}
func (self Simple) IsFloorBlockOnWallEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFloorBlockOnWallEnabled())
}
func (self Simple) SetSlideOnCeilingEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSlideOnCeilingEnabled(enabled)
}
func (self Simple) IsSlideOnCeilingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSlideOnCeilingEnabled())
}
func (self Simple) SetPlatformFloorLayers(exclude_layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlatformFloorLayers(gd.Int(exclude_layer))
}
func (self Simple) GetPlatformFloorLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPlatformFloorLayers()))
}
func (self Simple) SetPlatformWallLayers(exclude_layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlatformWallLayers(gd.Int(exclude_layer))
}
func (self Simple) GetPlatformWallLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPlatformWallLayers()))
}
func (self Simple) GetMaxSlides() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxSlides()))
}
func (self Simple) SetMaxSlides(max_slides int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxSlides(gd.Int(max_slides))
}
func (self Simple) GetFloorMaxAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFloorMaxAngle()))
}
func (self Simple) SetFloorMaxAngle(radians float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFloorMaxAngle(gd.Float(radians))
}
func (self Simple) GetFloorSnapLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFloorSnapLength()))
}
func (self Simple) SetFloorSnapLength(floor_snap_length float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFloorSnapLength(gd.Float(floor_snap_length))
}
func (self Simple) GetWallMinSlideAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetWallMinSlideAngle()))
}
func (self Simple) SetWallMinSlideAngle(radians float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWallMinSlideAngle(gd.Float(radians))
}
func (self Simple) GetUpDirection() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetUpDirection())
}
func (self Simple) SetUpDirection(up_direction gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUpDirection(up_direction)
}
func (self Simple) SetMotionMode(mode classdb.CharacterBody3DMotionMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMotionMode(mode)
}
func (self Simple) GetMotionMode() classdb.CharacterBody3DMotionMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CharacterBody3DMotionMode(Expert(self).GetMotionMode())
}
func (self Simple) SetPlatformOnLeave(on_leave_apply_velocity classdb.CharacterBody3DPlatformOnLeave) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlatformOnLeave(on_leave_apply_velocity)
}
func (self Simple) GetPlatformOnLeave() classdb.CharacterBody3DPlatformOnLeave {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CharacterBody3DPlatformOnLeave(Expert(self).GetPlatformOnLeave())
}
func (self Simple) IsOnFloor() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnFloor())
}
func (self Simple) IsOnFloorOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnFloorOnly())
}
func (self Simple) IsOnCeiling() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnCeiling())
}
func (self Simple) IsOnCeilingOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnCeilingOnly())
}
func (self Simple) IsOnWall() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnWall())
}
func (self Simple) IsOnWallOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOnWallOnly())
}
func (self Simple) GetFloorNormal() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetFloorNormal())
}
func (self Simple) GetWallNormal() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetWallNormal())
}
func (self Simple) GetLastMotion() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetLastMotion())
}
func (self Simple) GetPositionDelta() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetPositionDelta())
}
func (self Simple) GetRealVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetRealVelocity())
}
func (self Simple) GetFloorAngle(up_direction gd.Vector3) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFloorAngle(up_direction)))
}
func (self Simple) GetPlatformVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetPlatformVelocity())
}
func (self Simple) GetPlatformAngularVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetPlatformAngularVelocity())
}
func (self Simple) GetSlideCollisionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSlideCollisionCount()))
}
func (self Simple) GetSlideCollision(slide_idx int) [1]classdb.KinematicCollision3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.KinematicCollision3D(Expert(self).GetSlideCollision(gc, gd.Int(slide_idx)))
}
func (self Simple) GetLastSlideCollision() [1]classdb.KinematicCollision3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.KinematicCollision3D(Expert(self).GetLastSlideCollision(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CharacterBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body rather than stop immediately. If the other body is a [CharacterBody3D] or [RigidBody3D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for more detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
//go:nosplit
func (self class) MoveAndSlide() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_move_and_slide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Allows to manually apply a snap to the floor regardless of the body's velocity. This function does nothing when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) ApplyFloorSnap()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_apply_floor_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVelocity(velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSafeMargin(margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_safe_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSafeMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_safe_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsFloorStopOnSlopeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFloorStopOnSlopeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFloorConstantSpeedEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFloorConstantSpeedEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFloorBlockOnWallEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFloorBlockOnWallEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSlideOnCeilingEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSlideOnCeilingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlatformFloorLayers(exclude_layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exclude_layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlatformFloorLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlatformWallLayers(exclude_layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, exclude_layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlatformWallLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMaxSlides() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_max_slides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSlides(max_slides gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_slides)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_max_slides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFloorMaxAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFloorMaxAngle(radians gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFloorSnapLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFloorSnapLength(floor_snap_length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, floor_snap_length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWallMinSlideAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWallMinSlideAngle(radians gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpDirection() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_up_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpDirection(up_direction gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_up_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMotionMode(mode classdb.CharacterBody3DMotionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_motion_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMotionMode() classdb.CharacterBody3DMotionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CharacterBody3DMotionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_motion_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlatformOnLeave(on_leave_apply_velocity classdb.CharacterBody3DPlatformOnLeave)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, on_leave_apply_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_set_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlatformOnLeave() classdb.CharacterBody3DPlatformOnLeave {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CharacterBody3DPlatformOnLeave](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the body collided with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
//go:nosplit
func (self class) IsOnFloor() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_on_floor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the body collided only with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
//go:nosplit
func (self class) IsOnFloorOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_on_floor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the body collided with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
//go:nosplit
func (self class) IsOnCeiling() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_on_ceiling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the body collided only with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
//go:nosplit
func (self class) IsOnCeilingOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_on_ceiling_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the body collided with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
//go:nosplit
func (self class) IsOnWall() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_on_wall, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the body collided only with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
//go:nosplit
func (self class) IsOnWallOnly() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_is_on_wall_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetFloorNormal() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_floor_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetWallNormal() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_wall_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the last motion applied to the [CharacterBody3D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
//go:nosplit
func (self class) GetLastMotion() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_last_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetPositionDelta() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_position_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
//go:nosplit
func (self class) GetRealVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_real_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) GetFloorAngle(up_direction gd.Vector3) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_floor_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
//go:nosplit
func (self class) GetPlatformVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_platform_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the angular velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
//go:nosplit
func (self class) GetPlatformAngularVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_platform_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetSlideCollisionCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_slide_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [KinematicCollision3D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
*/
//go:nosplit
func (self class) GetSlideCollision(ctx gd.Lifetime, slide_idx gd.Int) object.KinematicCollision3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slide_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_slide_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.KinematicCollision3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [KinematicCollision3D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetLastSlideCollision(ctx gd.Lifetime) object.KinematicCollision3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody3D.Bind_get_last_slide_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.KinematicCollision3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCharacterBody3D() Expert { return self[0].AsCharacterBody3D() }


//go:nosplit
func (self Simple) AsCharacterBody3D() Simple { return self[0].AsCharacterBody3D() }


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
func init() {classdb.Register("CharacterBody3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type MotionMode = classdb.CharacterBody3DMotionMode

const (
/*Apply when notions of walls, ceiling and floor are relevant. In this mode the body motion will react to slopes (acceleration/slowdown). This mode is suitable for grounded games like platformers.*/
	MotionModeGrounded MotionMode = 0
/*Apply when there is no notion of floor or ceiling. All collisions will be reported as [code]on_wall[/code]. In this mode, when you slide, the speed will always be constant. This mode is suitable for games without ground like space games.*/
	MotionModeFloating MotionMode = 1
)
type PlatformOnLeave = classdb.CharacterBody3DPlatformOnLeave

const (
/*Add the last platform velocity to the [member velocity] when you leave a moving platform.*/
	PlatformOnLeaveAddVelocity PlatformOnLeave = 0
/*Add the last platform velocity to the [member velocity] when you leave a moving platform, but any downward motion is ignored. It's useful to keep full jump height even when the platform is moving down.*/
	PlatformOnLeaveAddUpwardVelocity PlatformOnLeave = 1
/*Do nothing when leaving a platform.*/
	PlatformOnLeaveDoNothing PlatformOnLeave = 2
)
