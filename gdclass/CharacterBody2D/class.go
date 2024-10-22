package CharacterBody2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PhysicsBody2D"
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
[CharacterBody2D] is a specialized class for physics bodies that are meant to be user-controlled. They are not affected by physics at all, but they affect other physics bodies in their path. They are mainly used to provide high-level API to move objects with wall and slope detection ([method move_and_slide] method) in addition to the general collision detection provided by [method PhysicsBody2D.move_and_collide]. This makes it useful for highly configurable physics bodies that must move in specific ways and collide with the world, as is often the case with user-controlled characters.
For game objects that don't require complex movement or collision detection, such as moving platforms, [AnimatableBody2D] is simpler to configure.

*/
type Go [1]classdb.CharacterBody2D

/*
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body (by default only on floor) rather than stop immediately. If the other body is a [CharacterBody2D] or [RigidBody2D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
The general behavior and available properties change according to the [member motion_mode].
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
func (self Go) MoveAndSlide() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).MoveAndSlide())
}

/*
Allows to manually apply a snap to the floor regardless of the body's velocity. This function does nothing when [method is_on_floor] returns [code]true[/code].
*/
func (self Go) ApplyFloorSnap() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ApplyFloorSnap()
}

/*
Returns [code]true[/code] if the body collided with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
func (self Go) IsOnFloor() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnFloor())
}

/*
Returns [code]true[/code] if the body collided only with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
func (self Go) IsOnFloorOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnFloorOnly())
}

/*
Returns [code]true[/code] if the body collided with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
func (self Go) IsOnCeiling() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnCeiling())
}

/*
Returns [code]true[/code] if the body collided only with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
func (self Go) IsOnCeilingOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnCeilingOnly())
}

/*
Returns [code]true[/code] if the body collided with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
func (self Go) IsOnWall() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnWall())
}

/*
Returns [code]true[/code] if the body collided only with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
func (self Go) IsOnWallOnly() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsOnWallOnly())
}

/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
func (self Go) GetFloorNormal() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetFloorNormal())
}

/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
func (self Go) GetWallNormal() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetWallNormal())
}

/*
Returns the last motion applied to the [CharacterBody2D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
func (self Go) GetLastMotion() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetLastMotion())
}

/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
func (self Go) GetPositionDelta() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetPositionDelta())
}

/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
func (self Go) GetRealVelocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetRealVelocity())
}

/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
func (self Go) GetFloorAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetFloorAngle(gd.Vector2{0, -1})))
}

/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
func (self Go) GetPlatformVelocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetPlatformVelocity())
}

/*
Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
func (self Go) GetSlideCollisionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSlideCollisionCount()))
}

/*
Returns a [KinematicCollision2D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
[b]Example usage:[/b]
[codeblocks]
[gdscript]
for i in get_slide_collision_count():
    var collision = get_slide_collision(i)
    print("Collided with: ", collision.get_collider().name)
[/gdscript]
[csharp]
for (int i = 0; i < GetSlideCollisionCount(); i++)
{
    KinematicCollision2D collision = GetSlideCollision(i);
    GD.Print("Collided with: ", (collision.GetCollider() as Node).Name);
}
[/csharp]
[/codeblocks]
*/
func (self Go) GetSlideCollision(slide_idx int) gdclass.KinematicCollision2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.KinematicCollision2D(class(self).GetSlideCollision(gc, gd.Int(slide_idx)))
}

/*
Returns a [KinematicCollision2D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
func (self Go) GetLastSlideCollision() gdclass.KinematicCollision2D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.KinematicCollision2D(class(self).GetLastSlideCollision(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CharacterBody2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CharacterBody2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MotionMode() classdb.CharacterBody2DMotionMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CharacterBody2DMotionMode(class(self).GetMotionMode())
}

func (self Go) SetMotionMode(value classdb.CharacterBody2DMotionMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMotionMode(value)
}

func (self Go) UpDirection() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetUpDirection())
}

func (self Go) SetUpDirection(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUpDirection(value)
}

func (self Go) Velocity() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetVelocity())
}

func (self Go) SetVelocity(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVelocity(value)
}

func (self Go) SlideOnCeiling() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSlideOnCeilingEnabled())
}

func (self Go) SetSlideOnCeiling(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSlideOnCeilingEnabled(value)
}

func (self Go) MaxSlides() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxSlides()))
}

func (self Go) SetMaxSlides(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxSlides(gd.Int(value))
}

func (self Go) WallMinSlideAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetWallMinSlideAngle()))
}

func (self Go) SetWallMinSlideAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetWallMinSlideAngle(gd.Float(value))
}

func (self Go) FloorStopOnSlope() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFloorStopOnSlopeEnabled())
}

func (self Go) SetFloorStopOnSlope(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFloorStopOnSlopeEnabled(value)
}

func (self Go) FloorConstantSpeed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFloorConstantSpeedEnabled())
}

func (self Go) SetFloorConstantSpeed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFloorConstantSpeedEnabled(value)
}

func (self Go) FloorBlockOnWall() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFloorBlockOnWallEnabled())
}

func (self Go) SetFloorBlockOnWall(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFloorBlockOnWallEnabled(value)
}

func (self Go) FloorMaxAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFloorMaxAngle()))
}

func (self Go) SetFloorMaxAngle(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFloorMaxAngle(gd.Float(value))
}

func (self Go) FloorSnapLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFloorSnapLength()))
}

func (self Go) SetFloorSnapLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFloorSnapLength(gd.Float(value))
}

func (self Go) PlatformOnLeave() classdb.CharacterBody2DPlatformOnLeave {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.CharacterBody2DPlatformOnLeave(class(self).GetPlatformOnLeave())
}

func (self Go) SetPlatformOnLeave(value classdb.CharacterBody2DPlatformOnLeave) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlatformOnLeave(value)
}

func (self Go) PlatformFloorLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPlatformFloorLayers()))
}

func (self Go) SetPlatformFloorLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlatformFloorLayers(gd.Int(value))
}

func (self Go) PlatformWallLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPlatformWallLayers()))
}

func (self Go) SetPlatformWallLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlatformWallLayers(gd.Int(value))
}

func (self Go) SafeMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSafeMargin()))
}

func (self Go) SetSafeMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSafeMargin(gd.Float(value))
}

/*
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body (by default only on floor) rather than stop immediately. If the other body is a [CharacterBody2D] or [RigidBody2D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
The general behavior and available properties change according to the [member motion_mode].
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
//go:nosplit
func (self class) MoveAndSlide() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_move_and_slide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_apply_floor_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVelocity(velocity gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_safe_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSafeMargin() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_safe_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) IsFloorStopOnSlopeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFloorConstantSpeedEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFloorConstantSpeedEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFloorBlockOnWallEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSlideOnCeilingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlatformFloorLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlatformWallLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMaxSlides() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_max_slides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_max_slides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFloorMaxAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFloorSnapLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWallMinSlideAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpDirection() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_up_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpDirection(up_direction gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_up_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMotionMode(mode classdb.CharacterBody2DMotionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_motion_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMotionMode() classdb.CharacterBody2DMotionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CharacterBody2DMotionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_motion_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlatformOnLeave(on_leave_apply_velocity classdb.CharacterBody2DPlatformOnLeave)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, on_leave_apply_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_set_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlatformOnLeave() classdb.CharacterBody2DPlatformOnLeave {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CharacterBody2DPlatformOnLeave](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_on_floor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_on_floor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_on_ceiling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_on_ceiling_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_on_wall, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_is_on_wall_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetFloorNormal() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_floor_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetWallNormal() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_wall_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the last motion applied to the [CharacterBody2D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
//go:nosplit
func (self class) GetLastMotion() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_last_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetPositionDelta() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_position_delta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
//go:nosplit
func (self class) GetRealVelocity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_real_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) GetFloorAngle(up_direction gd.Vector2) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_floor_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
//go:nosplit
func (self class) GetPlatformVelocity() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_platform_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_slide_collision_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [KinematicCollision2D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
[b]Example usage:[/b]
[codeblocks]
[gdscript]
for i in get_slide_collision_count():
    var collision = get_slide_collision(i)
    print("Collided with: ", collision.get_collider().name)
[/gdscript]
[csharp]
for (int i = 0; i < GetSlideCollisionCount(); i++)
{
    KinematicCollision2D collision = GetSlideCollision(i);
    GD.Print("Collided with: ", (collision.GetCollider() as Node).Name);
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetSlideCollision(ctx gd.Lifetime, slide_idx gd.Int) gdclass.KinematicCollision2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, slide_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_slide_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.KinematicCollision2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [KinematicCollision2D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetLastSlideCollision(ctx gd.Lifetime) gdclass.KinematicCollision2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CharacterBody2D.Bind_get_last_slide_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.KinematicCollision2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsCharacterBody2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCharacterBody2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody2D() PhysicsBody2D.GD { return *((*PhysicsBody2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody2D() PhysicsBody2D.Go { return *((*PhysicsBody2D.Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsPhysicsBody2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPhysicsBody2D(), name)
	}
}
func init() {classdb.Register("CharacterBody2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type MotionMode = classdb.CharacterBody2DMotionMode

const (
/*Apply when notions of walls, ceiling and floor are relevant. In this mode the body motion will react to slopes (acceleration/slowdown). This mode is suitable for sided games like platformers.*/
	MotionModeGrounded MotionMode = 0
/*Apply when there is no notion of floor or ceiling. All collisions will be reported as [code]on_wall[/code]. In this mode, when you slide, the speed will always be constant. This mode is suitable for top-down games.*/
	MotionModeFloating MotionMode = 1
)
type PlatformOnLeave = classdb.CharacterBody2DPlatformOnLeave

const (
/*Add the last platform velocity to the [member velocity] when you leave a moving platform.*/
	PlatformOnLeaveAddVelocity PlatformOnLeave = 0
/*Add the last platform velocity to the [member velocity] when you leave a moving platform, but any downward motion is ignored. It's useful to keep full jump height even when the platform is moving down.*/
	PlatformOnLeaveAddUpwardVelocity PlatformOnLeave = 1
/*Do nothing when leaving a platform.*/
	PlatformOnLeaveDoNothing PlatformOnLeave = 2
)
