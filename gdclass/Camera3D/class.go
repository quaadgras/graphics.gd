package Camera3D

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
[Camera3D] is a special node that displays what is visible from its current location. Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the camera will register in the global viewport. In other words, a camera just provides 3D display capabilities to a [Viewport], and, without one, a scene registered in that [Viewport] (or higher viewports) can't be displayed.

*/
type Go [1]classdb.Camera3D

/*
Returns a normal vector in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (self Go) ProjectRayNormal(screen_point gd.Vector2) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ProjectRayNormal(screen_point))
}

/*
Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
func (self Go) ProjectLocalRayNormal(screen_point gd.Vector2) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ProjectLocalRayNormal(screen_point))
}

/*
Returns a 3D position in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (self Go) ProjectRayOrigin(screen_point gd.Vector2) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ProjectRayOrigin(screen_point))
}

/*
Returns the 2D coordinate in the [Viewport] rectangle that maps to the given 3D point in world space.
[b]Note:[/b] When using this to position GUI elements over a 3D viewport, use [method is_position_behind] to prevent them from appearing if the 3D point is behind the camera:
[codeblock]
# This code block is part of a script that inherits from Node3D.
# `control` is a reference to a node inheriting from Control.
control.visible = not get_viewport().get_camera_3d().is_position_behind(global_transform.origin)
control.position = get_viewport().get_camera_3d().unproject_position(global_transform.origin)
[/codeblock]
*/
func (self Go) UnprojectPosition(world_point gd.Vector3) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).UnprojectPosition(world_point))
}

/*
Returns [code]true[/code] if the given position is behind the camera (the blue part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
[b]Note:[/b] A position which returns [code]false[/code] may still be outside the camera's field of view.
*/
func (self Go) IsPositionBehind(world_point gd.Vector3) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPositionBehind(world_point))
}

/*
Returns the 3D point in world space that maps to the given 2D coordinate in the [Viewport] rectangle on a plane that is the given [param z_depth] distance into the scene away from the camera.
*/
func (self Go) ProjectPosition(screen_point gd.Vector2, z_depth float64) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).ProjectPosition(screen_point, gd.Float(z_depth)))
}

/*
Sets the camera projection to perspective mode (see [constant PROJECTION_PERSPECTIVE]), by specifying a [param fov] (field of view) angle in degrees, and the [param z_near] and [param z_far] clip planes in world space units.
*/
func (self Go) SetPerspective(fov float64, z_near float64, z_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPerspective(gd.Float(fov), gd.Float(z_near), gd.Float(z_far))
}

/*
Sets the camera projection to orthogonal mode (see [constant PROJECTION_ORTHOGONAL]), by specifying a [param size], and the [param z_near] and [param z_far] clip planes in world space units. (As a hint, 2D games often use this projection, with values specified in pixels.)
*/
func (self Go) SetOrthogonal(size float64, z_near float64, z_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOrthogonal(gd.Float(size), gd.Float(z_near), gd.Float(z_far))
}

/*
Sets the camera projection to frustum mode (see [constant PROJECTION_FRUSTUM]), by specifying a [param size], an [param offset], and the [param z_near] and [param z_far] clip planes in world space units. See also [member frustum_offset].
*/
func (self Go) SetFrustum(size float64, offset gd.Vector2, z_near float64, z_far float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrustum(gd.Float(size), offset, gd.Float(z_near), gd.Float(z_far))
}

/*
Makes this camera the current camera for the [Viewport] (see class description). If the camera node is outside the scene tree, it will attempt to become current once it's added.
*/
func (self Go) MakeCurrent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MakeCurrent()
}

/*
If this is the current camera, remove it from being current. If [param enable_next] is [code]true[/code], request to make the next camera current, if any.
*/
func (self Go) ClearCurrent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearCurrent(true)
}

/*
Returns the transform of the camera plus the vertical ([member v_offset]) and horizontal ([member h_offset]) offsets; and any other adjustments made to the position and orientation of the camera by subclassed cameras such as [XRCamera3D].
*/
func (self Go) GetCameraTransform() gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(class(self).GetCameraTransform())
}

/*
Returns the projection matrix that this camera uses to render to its associated viewport. The camera must be part of the scene tree to function.
*/
func (self Go) GetCameraProjection() gd.Projection {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Projection(class(self).GetCameraProjection())
}

/*
Returns the camera's frustum planes in world space units as an array of [Plane]s in the following order: near, far, left, top, right, bottom. Not to be confused with [member frustum_offset].
*/
func (self Go) GetFrustum() gd.ArrayOf[gd.Plane] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Plane](class(self).GetFrustum(gc))
}

/*
Returns [code]true[/code] if the given position is inside the camera's frustum (the green part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
*/
func (self Go) IsPositionInFrustum(world_point gd.Vector3) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPositionInFrustum(world_point))
}

/*
Returns the camera's RID from the [RenderingServer].
*/
func (self Go) GetCameraRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetCameraRid())
}

/*
Returns the RID of a pyramid shape encompassing the camera's view frustum, ignoring the camera's near plane. The tip of the pyramid represents the position of the camera.
*/
func (self Go) GetPyramidShapeRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetPyramidShapeRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member cull_mask], given a [param layer_number] between 1 and 20.
*/
func (self Go) SetCullMaskValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCullMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member cull_mask] is enabled, given a [param layer_number] between 1 and 20.
*/
func (self Go) GetCullMaskValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetCullMaskValue(gd.Int(layer_number)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Camera3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Camera3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) KeepAspect() classdb.Camera3DKeepAspect {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Camera3DKeepAspect(class(self).GetKeepAspectMode())
}

func (self Go) SetKeepAspect(value classdb.Camera3DKeepAspect) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeepAspectMode(value)
}

func (self Go) CullMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCullMask()))
}

func (self Go) SetCullMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCullMask(gd.Int(value))
}

func (self Go) Environment() gdclass.Environment {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Environment(class(self).GetEnvironment(gc))
}

func (self Go) SetEnvironment(value gdclass.Environment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnvironment(value)
}

func (self Go) Attributes() gdclass.CameraAttributes {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.CameraAttributes(class(self).GetAttributes(gc))
}

func (self Go) SetAttributes(value gdclass.CameraAttributes) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAttributes(value)
}

func (self Go) Compositor() gdclass.Compositor {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Compositor(class(self).GetCompositor(gc))
}

func (self Go) SetCompositor(value gdclass.Compositor) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCompositor(value)
}

func (self Go) HOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetHOffset()))
}

func (self Go) SetHOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHOffset(gd.Float(value))
}

func (self Go) VOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVOffset()))
}

func (self Go) SetVOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVOffset(gd.Float(value))
}

func (self Go) DopplerTracking() classdb.Camera3DDopplerTracking {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Camera3DDopplerTracking(class(self).GetDopplerTracking())
}

func (self Go) SetDopplerTracking(value classdb.Camera3DDopplerTracking) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDopplerTracking(value)
}

func (self Go) Projection() classdb.Camera3DProjectionType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Camera3DProjectionType(class(self).GetProjection())
}

func (self Go) SetProjection(value classdb.Camera3DProjectionType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProjection(value)
}

func (self Go) Current() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCurrent())
}

func (self Go) SetCurrent(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrent(value)
}

func (self Go) Fov() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFov()))
}

func (self Go) SetFov(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFov(gd.Float(value))
}

func (self Go) Size() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSize()))
}

func (self Go) SetSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSize(gd.Float(value))
}

func (self Go) FrustumOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetFrustumOffset())
}

func (self Go) SetFrustumOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFrustumOffset(value)
}

func (self Go) Near() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetNear()))
}

func (self Go) SetNear(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNear(gd.Float(value))
}

func (self Go) Far() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFar()))
}

func (self Go) SetFar(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFar(gd.Float(value))
}

/*
Returns a normal vector in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
//go:nosplit
func (self class) ProjectRayNormal(screen_point gd.Vector2) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_project_ray_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
//go:nosplit
func (self class) ProjectLocalRayNormal(screen_point gd.Vector2) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_project_local_ray_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a 3D position in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
//go:nosplit
func (self class) ProjectRayOrigin(screen_point gd.Vector2) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_project_ray_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 2D coordinate in the [Viewport] rectangle that maps to the given 3D point in world space.
[b]Note:[/b] When using this to position GUI elements over a 3D viewport, use [method is_position_behind] to prevent them from appearing if the 3D point is behind the camera:
[codeblock]
# This code block is part of a script that inherits from Node3D.
# `control` is a reference to a node inheriting from Control.
control.visible = not get_viewport().get_camera_3d().is_position_behind(global_transform.origin)
control.position = get_viewport().get_camera_3d().unproject_position(global_transform.origin)
[/codeblock]
*/
//go:nosplit
func (self class) UnprojectPosition(world_point gd.Vector3) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_unproject_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given position is behind the camera (the blue part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
[b]Note:[/b] A position which returns [code]false[/code] may still be outside the camera's field of view.
*/
//go:nosplit
func (self class) IsPositionBehind(world_point gd.Vector3) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_is_position_behind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 3D point in world space that maps to the given 2D coordinate in the [Viewport] rectangle on a plane that is the given [param z_depth] distance into the scene away from the camera.
*/
//go:nosplit
func (self class) ProjectPosition(screen_point gd.Vector2, z_depth gd.Float) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	callframe.Arg(frame, z_depth)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_project_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the camera projection to perspective mode (see [constant PROJECTION_PERSPECTIVE]), by specifying a [param fov] (field of view) angle in degrees, and the [param z_near] and [param z_far] clip planes in world space units.
*/
//go:nosplit
func (self class) SetPerspective(fov gd.Float, z_near gd.Float, z_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the camera projection to orthogonal mode (see [constant PROJECTION_ORTHOGONAL]), by specifying a [param size], and the [param z_near] and [param z_far] clip planes in world space units. (As a hint, 2D games often use this projection, with values specified in pixels.)
*/
//go:nosplit
func (self class) SetOrthogonal(size gd.Float, z_near gd.Float, z_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_orthogonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the camera projection to frustum mode (see [constant PROJECTION_FRUSTUM]), by specifying a [param size], an [param offset], and the [param z_near] and [param z_far] clip planes in world space units. See also [member frustum_offset].
*/
//go:nosplit
func (self class) SetFrustum(size gd.Float, offset gd.Vector2, z_near gd.Float, z_far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Makes this camera the current camera for the [Viewport] (see class description). If the camera node is outside the scene tree, it will attempt to become current once it's added.
*/
//go:nosplit
func (self class) MakeCurrent()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If this is the current camera, remove it from being current. If [param enable_next] is [code]true[/code], request to make the next camera current, if any.
*/
//go:nosplit
func (self class) ClearCurrent(enable_next bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable_next)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_clear_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurrent(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCurrent() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transform of the camera plus the vertical ([member v_offset]) and horizontal ([member h_offset]) offsets; and any other adjustments made to the position and orientation of the camera by subclassed cameras such as [XRCamera3D].
*/
//go:nosplit
func (self class) GetCameraTransform() gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_camera_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the projection matrix that this camera uses to render to its associated viewport. The camera must be part of the scene tree to function.
*/
//go:nosplit
func (self class) GetCameraProjection() gd.Projection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Projection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_camera_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFov() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFrustumOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_frustum_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFar() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetNear() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFov(fov gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFrustumOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_frustum_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFar(far gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, far)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetNear(near gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, near)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProjection() classdb.Camera3DProjectionType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera3DProjectionType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProjection(mode classdb.Camera3DProjectionType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHOffset(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVOffset(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCullMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCullMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironment(env gdclass.Environment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(env[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironment(ctx gd.Lifetime) gdclass.Environment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Environment
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAttributes(env gdclass.CameraAttributes)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(env[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttributes(ctx gd.Lifetime) gdclass.CameraAttributes {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CameraAttributes
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCompositor(compositor gdclass.Compositor)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(compositor[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCompositor(ctx gd.Lifetime) gdclass.Compositor {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Compositor
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepAspectMode(mode classdb.Camera3DKeepAspect)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_keep_aspect_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeepAspectMode() classdb.Camera3DKeepAspect {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera3DKeepAspect](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_keep_aspect_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDopplerTracking(mode classdb.Camera3DDopplerTracking)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDopplerTracking() classdb.Camera3DDopplerTracking {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera3DDopplerTracking](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the camera's frustum planes in world space units as an array of [Plane]s in the following order: near, far, left, top, right, bottom. Not to be confused with [member frustum_offset].
*/
//go:nosplit
func (self class) GetFrustum(ctx gd.Lifetime) gd.ArrayOf[gd.Plane] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Plane](ret)
}
/*
Returns [code]true[/code] if the given position is inside the camera's frustum (the green part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
*/
//go:nosplit
func (self class) IsPositionInFrustum(world_point gd.Vector3) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_is_position_in_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the camera's RID from the [RenderingServer].
*/
//go:nosplit
func (self class) GetCameraRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_camera_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the RID of a pyramid shape encompassing the camera's view frustum, ignoring the camera's near plane. The tip of the pyramid represents the position of the camera.
*/
//go:nosplit
func (self class) GetPyramidShapeRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_pyramid_shape_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member cull_mask], given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) SetCullMaskValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_set_cull_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member cull_mask] is enabled, given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) GetCullMaskValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera3D.Bind_get_cull_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCamera3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCamera3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Camera3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ProjectionType = classdb.Camera3DProjectionType

const (
/*Perspective projection. Objects on the screen becomes smaller when they are far away.*/
	ProjectionPerspective ProjectionType = 0
/*Orthogonal projection, also known as orthographic projection. Objects remain the same size on the screen no matter how far away they are.*/
	ProjectionOrthogonal ProjectionType = 1
/*Frustum projection. This mode allows adjusting [member frustum_offset] to create "tilted frustum" effects.*/
	ProjectionFrustum ProjectionType = 2
)
type KeepAspect = classdb.Camera3DKeepAspect

const (
/*Preserves the horizontal aspect ratio; also known as Vert- scaling. This is usually the best option for projects running in portrait mode, as taller aspect ratios will benefit from a wider vertical FOV.*/
	KeepWidth KeepAspect = 0
/*Preserves the vertical aspect ratio; also known as Hor+ scaling. This is usually the best option for projects running in landscape mode, as wider aspect ratios will automatically benefit from a wider horizontal FOV.*/
	KeepHeight KeepAspect = 1
)
type DopplerTracking = classdb.Camera3DDopplerTracking

const (
/*Disables [url=https://en.wikipedia.org/wiki/Doppler_effect]Doppler effect[/url] simulation (default).*/
	DopplerTrackingDisabled DopplerTracking = 0
/*Simulate [url=https://en.wikipedia.org/wiki/Doppler_effect]Doppler effect[/url] by tracking positions of objects that are changed in [code]_process[/code]. Changes in the relative velocity of this camera compared to those objects affect how audio is perceived (changing the audio's [member AudioStreamPlayer3D.pitch_scale]).*/
	DopplerTrackingIdleStep DopplerTracking = 1
/*Simulate [url=https://en.wikipedia.org/wiki/Doppler_effect]Doppler effect[/url] by tracking positions of objects that are changed in [code]_physics_process[/code]. Changes in the relative velocity of this camera compared to those objects affect how audio is perceived (changing the audio's [member AudioStreamPlayer3D.pitch_scale]).*/
	DopplerTrackingPhysicsStep DopplerTracking = 2
)
