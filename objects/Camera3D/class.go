package Camera3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[Camera3D] is a special node that displays what is visible from its current location. Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the camera will register in the global viewport. In other words, a camera just provides 3D display capabilities to a [Viewport], and, without one, a scene registered in that [Viewport] (or higher viewports) can't be displayed.
*/
type Instance [1]classdb.Camera3D

/*
Returns a normal vector in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (self Instance) ProjectRayNormal(screen_point gd.Vector2) gd.Vector3 {
	return gd.Vector3(class(self).ProjectRayNormal(screen_point))
}

/*
Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
func (self Instance) ProjectLocalRayNormal(screen_point gd.Vector2) gd.Vector3 {
	return gd.Vector3(class(self).ProjectLocalRayNormal(screen_point))
}

/*
Returns a 3D position in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (self Instance) ProjectRayOrigin(screen_point gd.Vector2) gd.Vector3 {
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
func (self Instance) UnprojectPosition(world_point gd.Vector3) gd.Vector2 {
	return gd.Vector2(class(self).UnprojectPosition(world_point))
}

/*
Returns [code]true[/code] if the given position is behind the camera (the blue part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
[b]Note:[/b] A position which returns [code]false[/code] may still be outside the camera's field of view.
*/
func (self Instance) IsPositionBehind(world_point gd.Vector3) bool {
	return bool(class(self).IsPositionBehind(world_point))
}

/*
Returns the 3D point in world space that maps to the given 2D coordinate in the [Viewport] rectangle on a plane that is the given [param z_depth] distance into the scene away from the camera.
*/
func (self Instance) ProjectPosition(screen_point gd.Vector2, z_depth float64) gd.Vector3 {
	return gd.Vector3(class(self).ProjectPosition(screen_point, gd.Float(z_depth)))
}

/*
Sets the camera projection to perspective mode (see [constant PROJECTION_PERSPECTIVE]), by specifying a [param fov] (field of view) angle in degrees, and the [param z_near] and [param z_far] clip planes in world space units.
*/
func (self Instance) SetPerspective(fov float64, z_near float64, z_far float64) {
	class(self).SetPerspective(gd.Float(fov), gd.Float(z_near), gd.Float(z_far))
}

/*
Sets the camera projection to orthogonal mode (see [constant PROJECTION_ORTHOGONAL]), by specifying a [param size], and the [param z_near] and [param z_far] clip planes in world space units. (As a hint, 2D games often use this projection, with values specified in pixels.)
*/
func (self Instance) SetOrthogonal(size float64, z_near float64, z_far float64) {
	class(self).SetOrthogonal(gd.Float(size), gd.Float(z_near), gd.Float(z_far))
}

/*
Sets the camera projection to frustum mode (see [constant PROJECTION_FRUSTUM]), by specifying a [param size], an [param offset], and the [param z_near] and [param z_far] clip planes in world space units. See also [member frustum_offset].
*/
func (self Instance) SetFrustum(size float64, offset gd.Vector2, z_near float64, z_far float64) {
	class(self).SetFrustum(gd.Float(size), offset, gd.Float(z_near), gd.Float(z_far))
}

/*
Makes this camera the current camera for the [Viewport] (see class description). If the camera node is outside the scene tree, it will attempt to become current once it's added.
*/
func (self Instance) MakeCurrent() {
	class(self).MakeCurrent()
}

/*
If this is the current camera, remove it from being current. If [param enable_next] is [code]true[/code], request to make the next camera current, if any.
*/
func (self Instance) ClearCurrent() {
	class(self).ClearCurrent(true)
}

/*
Returns the transform of the camera plus the vertical ([member v_offset]) and horizontal ([member h_offset]) offsets; and any other adjustments made to the position and orientation of the camera by subclassed cameras such as [XRCamera3D].
*/
func (self Instance) GetCameraTransform() gd.Transform3D {
	return gd.Transform3D(class(self).GetCameraTransform())
}

/*
Returns the projection matrix that this camera uses to render to its associated viewport. The camera must be part of the scene tree to function.
*/
func (self Instance) GetCameraProjection() gd.Projection {
	return gd.Projection(class(self).GetCameraProjection())
}

/*
Returns the camera's frustum planes in world space units as an array of [Plane]s in the following order: near, far, left, top, right, bottom. Not to be confused with [member frustum_offset].
*/
func (self Instance) GetFrustum() gd.Array {
	return gd.Array(class(self).GetFrustum())
}

/*
Returns [code]true[/code] if the given position is inside the camera's frustum (the green part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
*/
func (self Instance) IsPositionInFrustum(world_point gd.Vector3) bool {
	return bool(class(self).IsPositionInFrustum(world_point))
}

/*
Returns the camera's RID from the [RenderingServer].
*/
func (self Instance) GetCameraRid() gd.RID {
	return gd.RID(class(self).GetCameraRid())
}

/*
Returns the RID of a pyramid shape encompassing the camera's view frustum, ignoring the camera's near plane. The tip of the pyramid represents the position of the camera.
*/
func (self Instance) GetPyramidShapeRid() gd.RID {
	return gd.RID(class(self).GetPyramidShapeRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member cull_mask], given a [param layer_number] between 1 and 20.
*/
func (self Instance) SetCullMaskValue(layer_number int, value bool) {
	class(self).SetCullMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member cull_mask] is enabled, given a [param layer_number] between 1 and 20.
*/
func (self Instance) GetCullMaskValue(layer_number int) bool {
	return bool(class(self).GetCullMaskValue(gd.Int(layer_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Camera3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Camera3D"))
	return Instance{classdb.Camera3D(object)}
}

func (self Instance) KeepAspect() classdb.Camera3DKeepAspect {
	return classdb.Camera3DKeepAspect(class(self).GetKeepAspectMode())
}

func (self Instance) SetKeepAspect(value classdb.Camera3DKeepAspect) {
	class(self).SetKeepAspectMode(value)
}

func (self Instance) CullMask() int {
	return int(int(class(self).GetCullMask()))
}

func (self Instance) SetCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

func (self Instance) Environment() objects.Environment {
	return objects.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value objects.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) Attributes() objects.CameraAttributes {
	return objects.CameraAttributes(class(self).GetAttributes())
}

func (self Instance) SetAttributes(value objects.CameraAttributes) {
	class(self).SetAttributes(value)
}

func (self Instance) Compositor() objects.Compositor {
	return objects.Compositor(class(self).GetCompositor())
}

func (self Instance) SetCompositor(value objects.Compositor) {
	class(self).SetCompositor(value)
}

func (self Instance) HOffset() float64 {
	return float64(float64(class(self).GetHOffset()))
}

func (self Instance) SetHOffset(value float64) {
	class(self).SetHOffset(gd.Float(value))
}

func (self Instance) VOffset() float64 {
	return float64(float64(class(self).GetVOffset()))
}

func (self Instance) SetVOffset(value float64) {
	class(self).SetVOffset(gd.Float(value))
}

func (self Instance) DopplerTracking() classdb.Camera3DDopplerTracking {
	return classdb.Camera3DDopplerTracking(class(self).GetDopplerTracking())
}

func (self Instance) SetDopplerTracking(value classdb.Camera3DDopplerTracking) {
	class(self).SetDopplerTracking(value)
}

func (self Instance) Projection() classdb.Camera3DProjectionType {
	return classdb.Camera3DProjectionType(class(self).GetProjection())
}

func (self Instance) SetProjection(value classdb.Camera3DProjectionType) {
	class(self).SetProjection(value)
}

func (self Instance) Current() bool {
	return bool(class(self).IsCurrent())
}

func (self Instance) SetCurrent(value bool) {
	class(self).SetCurrent(value)
}

func (self Instance) Fov() float64 {
	return float64(float64(class(self).GetFov()))
}

func (self Instance) SetFov(value float64) {
	class(self).SetFov(gd.Float(value))
}

func (self Instance) Size() float64 {
	return float64(float64(class(self).GetSize()))
}

func (self Instance) SetSize(value float64) {
	class(self).SetSize(gd.Float(value))
}

func (self Instance) FrustumOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetFrustumOffset())
}

func (self Instance) SetFrustumOffset(value gd.Vector2) {
	class(self).SetFrustumOffset(value)
}

func (self Instance) Near() float64 {
	return float64(float64(class(self).GetNear()))
}

func (self Instance) SetNear(value float64) {
	class(self).SetNear(gd.Float(value))
}

func (self Instance) Far() float64 {
	return float64(float64(class(self).GetFar()))
}

func (self Instance) SetFar(value float64) {
	class(self).SetFar(gd.Float(value))
}

/*
Returns a normal vector in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
//go:nosplit
func (self class) ProjectRayNormal(screen_point gd.Vector2) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_ray_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
//go:nosplit
func (self class) ProjectLocalRayNormal(screen_point gd.Vector2) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_local_ray_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a 3D position in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
//go:nosplit
func (self class) ProjectRayOrigin(screen_point gd.Vector2) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_ray_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_unproject_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_is_position_behind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D point in world space that maps to the given 2D coordinate in the [Viewport] rectangle on a plane that is the given [param z_depth] distance into the scene away from the camera.
*/
//go:nosplit
func (self class) ProjectPosition(screen_point gd.Vector2, z_depth gd.Float) gd.Vector3 {
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	callframe.Arg(frame, z_depth)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the camera projection to perspective mode (see [constant PROJECTION_PERSPECTIVE]), by specifying a [param fov] (field of view) angle in degrees, and the [param z_near] and [param z_far] clip planes in world space units.
*/
//go:nosplit
func (self class) SetPerspective(fov gd.Float, z_near gd.Float, z_far gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the camera projection to orthogonal mode (see [constant PROJECTION_ORTHOGONAL]), by specifying a [param size], and the [param z_near] and [param z_far] clip planes in world space units. (As a hint, 2D games often use this projection, with values specified in pixels.)
*/
//go:nosplit
func (self class) SetOrthogonal(size gd.Float, z_near gd.Float, z_far gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_orthogonal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the camera projection to frustum mode (see [constant PROJECTION_FRUSTUM]), by specifying a [param size], an [param offset], and the [param z_near] and [param z_far] clip planes in world space units. See also [member frustum_offset].
*/
//go:nosplit
func (self class) SetFrustum(size gd.Float, offset gd.Vector2, z_near gd.Float, z_far gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Makes this camera the current camera for the [Viewport] (see class description). If the camera node is outside the scene tree, it will attempt to become current once it's added.
*/
//go:nosplit
func (self class) MakeCurrent() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If this is the current camera, remove it from being current. If [param enable_next] is [code]true[/code], request to make the next camera current, if any.
*/
//go:nosplit
func (self class) ClearCurrent(enable_next bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable_next)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_clear_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurrent(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCurrent() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform of the camera plus the vertical ([member v_offset]) and horizontal ([member h_offset]) offsets; and any other adjustments made to the position and orientation of the camera by subclassed cameras such as [XRCamera3D].
*/
//go:nosplit
func (self class) GetCameraTransform() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_camera_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the projection matrix that this camera uses to render to its associated viewport. The camera must be part of the scene tree to function.
*/
//go:nosplit
func (self class) GetCameraProjection() gd.Projection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Projection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_camera_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFov() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFrustumOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_frustum_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFar() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetNear() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFov(fov gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFrustumOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_frustum_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSize(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFar(far gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, far)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetNear(near gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, near)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProjection() classdb.Camera3DProjectionType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera3DProjectionType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProjection(mode classdb.Camera3DProjectionType) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_projection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetHOffset(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVOffset(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironment(env objects.Environment) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() objects.Environment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Environment{classdb.Environment(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttributes(env objects.CameraAttributes) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttributes() objects.CameraAttributes {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.CameraAttributes{classdb.CameraAttributes(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCompositor(compositor objects.Compositor) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(compositor[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompositor() objects.Compositor {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Compositor{classdb.Compositor(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepAspectMode(mode classdb.Camera3DKeepAspect) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_keep_aspect_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetKeepAspectMode() classdb.Camera3DKeepAspect {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera3DKeepAspect](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_keep_aspect_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDopplerTracking(mode classdb.Camera3DDopplerTracking) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDopplerTracking() classdb.Camera3DDopplerTracking {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera3DDopplerTracking](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the camera's frustum planes in world space units as an array of [Plane]s in the following order: near, far, left, top, right, bottom. Not to be confused with [member frustum_offset].
*/
//go:nosplit
func (self class) GetFrustum() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given position is inside the camera's frustum (the green part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
*/
//go:nosplit
func (self class) IsPositionInFrustum(world_point gd.Vector3) bool {
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_is_position_in_frustum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the camera's RID from the [RenderingServer].
*/
//go:nosplit
func (self class) GetCameraRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_camera_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the RID of a pyramid shape encompassing the camera's view frustum, ignoring the camera's near plane. The tip of the pyramid represents the position of the camera.
*/
//go:nosplit
func (self class) GetPyramidShapeRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_pyramid_shape_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member cull_mask], given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) SetCullMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_cull_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member cull_mask] is enabled, given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) GetCullMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_cull_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCamera3D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCamera3D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

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
func init() { classdb.Register("Camera3D", func(ptr gd.Object) any { return classdb.Camera3D(ptr) }) }

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
