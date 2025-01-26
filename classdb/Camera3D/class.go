// Package Camera3D provides methods for working with Camera3D object instances.
package Camera3D

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
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Projection"
import "graphics.gd/variant/Plane"
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
[Camera3D] is a special node that displays what is visible from its current location. Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the camera will register in the global viewport. In other words, a camera just provides 3D display capabilities to a [Viewport], and, without one, a scene registered in that [Viewport] (or higher viewports) can't be displayed.
*/
type Instance [1]gdclass.Camera3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCamera3D() Instance
}

/*
Returns a normal vector in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (self Instance) ProjectRayNormal(screen_point Vector2.XY) Vector3.XYZ { //gd:Camera3D.project_ray_normal
	return Vector3.XYZ(class(self).ProjectRayNormal(gd.Vector2(screen_point)))
}

/*
Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
func (self Instance) ProjectLocalRayNormal(screen_point Vector2.XY) Vector3.XYZ { //gd:Camera3D.project_local_ray_normal
	return Vector3.XYZ(class(self).ProjectLocalRayNormal(gd.Vector2(screen_point)))
}

/*
Returns a 3D position in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
func (self Instance) ProjectRayOrigin(screen_point Vector2.XY) Vector3.XYZ { //gd:Camera3D.project_ray_origin
	return Vector3.XYZ(class(self).ProjectRayOrigin(gd.Vector2(screen_point)))
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
func (self Instance) UnprojectPosition(world_point Vector3.XYZ) Vector2.XY { //gd:Camera3D.unproject_position
	return Vector2.XY(class(self).UnprojectPosition(gd.Vector3(world_point)))
}

/*
Returns [code]true[/code] if the given position is behind the camera (the blue part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
[b]Note:[/b] A position which returns [code]false[/code] may still be outside the camera's field of view.
*/
func (self Instance) IsPositionBehind(world_point Vector3.XYZ) bool { //gd:Camera3D.is_position_behind
	return bool(class(self).IsPositionBehind(gd.Vector3(world_point)))
}

/*
Returns the 3D point in world space that maps to the given 2D coordinate in the [Viewport] rectangle on a plane that is the given [param z_depth] distance into the scene away from the camera.
*/
func (self Instance) ProjectPosition(screen_point Vector2.XY, z_depth Float.X) Vector3.XYZ { //gd:Camera3D.project_position
	return Vector3.XYZ(class(self).ProjectPosition(gd.Vector2(screen_point), gd.Float(z_depth)))
}

/*
Sets the camera projection to perspective mode (see [constant PROJECTION_PERSPECTIVE]), by specifying a [param fov] (field of view) angle in degrees, and the [param z_near] and [param z_far] clip planes in world space units.
*/
func (self Instance) SetPerspective(fov Float.X, z_near Float.X, z_far Float.X) { //gd:Camera3D.set_perspective
	class(self).SetPerspective(gd.Float(fov), gd.Float(z_near), gd.Float(z_far))
}

/*
Sets the camera projection to orthogonal mode (see [constant PROJECTION_ORTHOGONAL]), by specifying a [param size], and the [param z_near] and [param z_far] clip planes in world space units. (As a hint, 2D games often use this projection, with values specified in pixels.)
*/
func (self Instance) SetOrthogonal(size Float.X, z_near Float.X, z_far Float.X) { //gd:Camera3D.set_orthogonal
	class(self).SetOrthogonal(gd.Float(size), gd.Float(z_near), gd.Float(z_far))
}

/*
Sets the camera projection to frustum mode (see [constant PROJECTION_FRUSTUM]), by specifying a [param size], an [param offset], and the [param z_near] and [param z_far] clip planes in world space units. See also [member frustum_offset].
*/
func (self Instance) SetFrustum(size Float.X, offset Vector2.XY, z_near Float.X, z_far Float.X) { //gd:Camera3D.set_frustum
	class(self).SetFrustum(gd.Float(size), gd.Vector2(offset), gd.Float(z_near), gd.Float(z_far))
}

/*
Makes this camera the current camera for the [Viewport] (see class description). If the camera node is outside the scene tree, it will attempt to become current once it's added.
*/
func (self Instance) MakeCurrent() { //gd:Camera3D.make_current
	class(self).MakeCurrent()
}

/*
If this is the current camera, remove it from being current. If [param enable_next] is [code]true[/code], request to make the next camera current, if any.
*/
func (self Instance) ClearCurrent() { //gd:Camera3D.clear_current
	class(self).ClearCurrent(true)
}

/*
Returns the transform of the camera plus the vertical ([member v_offset]) and horizontal ([member h_offset]) offsets; and any other adjustments made to the position and orientation of the camera by subclassed cameras such as [XRCamera3D].
*/
func (self Instance) GetCameraTransform() Transform3D.BasisOrigin { //gd:Camera3D.get_camera_transform
	return Transform3D.BasisOrigin(class(self).GetCameraTransform())
}

/*
Returns the projection matrix that this camera uses to render to its associated viewport. The camera must be part of the scene tree to function.
*/
func (self Instance) GetCameraProjection() Projection.XYZW { //gd:Camera3D.get_camera_projection
	return Projection.XYZW(class(self).GetCameraProjection())
}

/*
Returns the camera's frustum planes in world space units as an array of [Plane]s in the following order: near, far, left, top, right, bottom. Not to be confused with [member frustum_offset].
*/
func (self Instance) GetFrustum() []Plane.NormalD { //gd:Camera3D.get_frustum
	return []Plane.NormalD(gd.ArrayAs[[]Plane.NormalD](gd.InternalArray(class(self).GetFrustum())))
}

/*
Returns [code]true[/code] if the given position is inside the camera's frustum (the green part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
*/
func (self Instance) IsPositionInFrustum(world_point Vector3.XYZ) bool { //gd:Camera3D.is_position_in_frustum
	return bool(class(self).IsPositionInFrustum(gd.Vector3(world_point)))
}

/*
Returns the camera's RID from the [RenderingServer].
*/
func (self Instance) GetCameraRid() RID.Camera { //gd:Camera3D.get_camera_rid
	return RID.Camera(class(self).GetCameraRid())
}

/*
Returns the RID of a pyramid shape encompassing the camera's view frustum, ignoring the camera's near plane. The tip of the pyramid represents the position of the camera.
*/
func (self Instance) GetPyramidShapeRid() RID.Shape3D { //gd:Camera3D.get_pyramid_shape_rid
	return RID.Shape3D(class(self).GetPyramidShapeRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member cull_mask], given a [param layer_number] between 1 and 20.
*/
func (self Instance) SetCullMaskValue(layer_number int, value bool) { //gd:Camera3D.set_cull_mask_value
	class(self).SetCullMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member cull_mask] is enabled, given a [param layer_number] between 1 and 20.
*/
func (self Instance) GetCullMaskValue(layer_number int) bool { //gd:Camera3D.get_cull_mask_value
	return bool(class(self).GetCullMaskValue(gd.Int(layer_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Camera3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Camera3D"))
	casted := Instance{*(*gdclass.Camera3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) KeepAspect() gdclass.Camera3DKeepAspect {
	return gdclass.Camera3DKeepAspect(class(self).GetKeepAspectMode())
}

func (self Instance) SetKeepAspect(value gdclass.Camera3DKeepAspect) {
	class(self).SetKeepAspectMode(value)
}

func (self Instance) CullMask() int {
	return int(int(class(self).GetCullMask()))
}

func (self Instance) SetCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

func (self Instance) Environment() [1]gdclass.Environment {
	return [1]gdclass.Environment(class(self).GetEnvironment())
}

func (self Instance) SetEnvironment(value [1]gdclass.Environment) {
	class(self).SetEnvironment(value)
}

func (self Instance) Attributes() [1]gdclass.CameraAttributes {
	return [1]gdclass.CameraAttributes(class(self).GetAttributes())
}

func (self Instance) SetAttributes(value [1]gdclass.CameraAttributes) {
	class(self).SetAttributes(value)
}

func (self Instance) Compositor() [1]gdclass.Compositor {
	return [1]gdclass.Compositor(class(self).GetCompositor())
}

func (self Instance) SetCompositor(value [1]gdclass.Compositor) {
	class(self).SetCompositor(value)
}

func (self Instance) HOffset() Float.X {
	return Float.X(Float.X(class(self).GetHOffset()))
}

func (self Instance) SetHOffset(value Float.X) {
	class(self).SetHOffset(gd.Float(value))
}

func (self Instance) VOffset() Float.X {
	return Float.X(Float.X(class(self).GetVOffset()))
}

func (self Instance) SetVOffset(value Float.X) {
	class(self).SetVOffset(gd.Float(value))
}

func (self Instance) DopplerTracking() gdclass.Camera3DDopplerTracking {
	return gdclass.Camera3DDopplerTracking(class(self).GetDopplerTracking())
}

func (self Instance) SetDopplerTracking(value gdclass.Camera3DDopplerTracking) {
	class(self).SetDopplerTracking(value)
}

func (self Instance) Projection() gdclass.Camera3DProjectionType {
	return gdclass.Camera3DProjectionType(class(self).GetProjection())
}

func (self Instance) SetProjection(value gdclass.Camera3DProjectionType) {
	class(self).SetProjection(value)
}

func (self Instance) Current() bool {
	return bool(class(self).IsCurrent())
}

func (self Instance) SetCurrent(value bool) {
	class(self).SetCurrent(value)
}

func (self Instance) Fov() Float.X {
	return Float.X(Float.X(class(self).GetFov()))
}

func (self Instance) SetFov(value Float.X) {
	class(self).SetFov(gd.Float(value))
}

func (self Instance) Size() Float.X {
	return Float.X(Float.X(class(self).GetSize()))
}

func (self Instance) SetSize(value Float.X) {
	class(self).SetSize(gd.Float(value))
}

func (self Instance) FrustumOffset() Vector2.XY {
	return Vector2.XY(class(self).GetFrustumOffset())
}

func (self Instance) SetFrustumOffset(value Vector2.XY) {
	class(self).SetFrustumOffset(gd.Vector2(value))
}

func (self Instance) Near() Float.X {
	return Float.X(Float.X(class(self).GetNear()))
}

func (self Instance) SetNear(value Float.X) {
	class(self).SetNear(gd.Float(value))
}

func (self Instance) Far() Float.X {
	return Float.X(Float.X(class(self).GetFar()))
}

func (self Instance) SetFar(value Float.X) {
	class(self).SetFar(gd.Float(value))
}

/*
Returns a normal vector in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
//go:nosplit
func (self class) ProjectRayNormal(screen_point gd.Vector2) gd.Vector3 { //gd:Camera3D.project_ray_normal
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_ray_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a normal vector from the screen point location directed along the camera. Orthogonal cameras are normalized. Perspective cameras account for perspective, screen width/height, etc.
*/
//go:nosplit
func (self class) ProjectLocalRayNormal(screen_point gd.Vector2) gd.Vector3 { //gd:Camera3D.project_local_ray_normal
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_local_ray_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a 3D position in world space, that is the result of projecting a point on the [Viewport] rectangle by the inverse camera projection. This is useful for casting rays in the form of (origin, normal) for object intersection or picking.
*/
//go:nosplit
func (self class) ProjectRayOrigin(screen_point gd.Vector2) gd.Vector3 { //gd:Camera3D.project_ray_origin
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_ray_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) UnprojectPosition(world_point gd.Vector3) gd.Vector2 { //gd:Camera3D.unproject_position
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_unproject_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given position is behind the camera (the blue part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
[b]Note:[/b] A position which returns [code]false[/code] may still be outside the camera's field of view.
*/
//go:nosplit
func (self class) IsPositionBehind(world_point gd.Vector3) bool { //gd:Camera3D.is_position_behind
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_is_position_behind, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D point in world space that maps to the given 2D coordinate in the [Viewport] rectangle on a plane that is the given [param z_depth] distance into the scene away from the camera.
*/
//go:nosplit
func (self class) ProjectPosition(screen_point gd.Vector2, z_depth gd.Float) gd.Vector3 { //gd:Camera3D.project_position
	var frame = callframe.New()
	callframe.Arg(frame, screen_point)
	callframe.Arg(frame, z_depth)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_project_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the camera projection to perspective mode (see [constant PROJECTION_PERSPECTIVE]), by specifying a [param fov] (field of view) angle in degrees, and the [param z_near] and [param z_far] clip planes in world space units.
*/
//go:nosplit
func (self class) SetPerspective(fov gd.Float, z_near gd.Float, z_far gd.Float) { //gd:Camera3D.set_perspective
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_perspective, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the camera projection to orthogonal mode (see [constant PROJECTION_ORTHOGONAL]), by specifying a [param size], and the [param z_near] and [param z_far] clip planes in world space units. (As a hint, 2D games often use this projection, with values specified in pixels.)
*/
//go:nosplit
func (self class) SetOrthogonal(size gd.Float, z_near gd.Float, z_far gd.Float) { //gd:Camera3D.set_orthogonal
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_orthogonal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the camera projection to frustum mode (see [constant PROJECTION_FRUSTUM]), by specifying a [param size], an [param offset], and the [param z_near] and [param z_far] clip planes in world space units. See also [member frustum_offset].
*/
//go:nosplit
func (self class) SetFrustum(size gd.Float, offset gd.Vector2, z_near gd.Float, z_far gd.Float) { //gd:Camera3D.set_frustum
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, z_near)
	callframe.Arg(frame, z_far)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_frustum, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Makes this camera the current camera for the [Viewport] (see class description). If the camera node is outside the scene tree, it will attempt to become current once it's added.
*/
//go:nosplit
func (self class) MakeCurrent() { //gd:Camera3D.make_current
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If this is the current camera, remove it from being current. If [param enable_next] is [code]true[/code], request to make the next camera current, if any.
*/
//go:nosplit
func (self class) ClearCurrent(enable_next bool) { //gd:Camera3D.clear_current
	var frame = callframe.New()
	callframe.Arg(frame, enable_next)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_clear_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCurrent(enabled bool) { //gd:Camera3D.set_current
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCurrent() bool { //gd:Camera3D.is_current
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform of the camera plus the vertical ([member v_offset]) and horizontal ([member h_offset]) offsets; and any other adjustments made to the position and orientation of the camera by subclassed cameras such as [XRCamera3D].
*/
//go:nosplit
func (self class) GetCameraTransform() gd.Transform3D { //gd:Camera3D.get_camera_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_camera_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the projection matrix that this camera uses to render to its associated viewport. The camera must be part of the scene tree to function.
*/
//go:nosplit
func (self class) GetCameraProjection() gd.Projection { //gd:Camera3D.get_camera_projection
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Projection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_camera_projection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFov() gd.Float { //gd:Camera3D.get_fov
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFrustumOffset() gd.Vector2 { //gd:Camera3D.get_frustum_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_frustum_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSize() gd.Float { //gd:Camera3D.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFar() gd.Float { //gd:Camera3D.get_far
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_far, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetNear() gd.Float { //gd:Camera3D.get_near
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_near, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFov(fov gd.Float) { //gd:Camera3D.set_fov
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_fov, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFrustumOffset(offset gd.Vector2) { //gd:Camera3D.set_frustum_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_frustum_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSize(size gd.Float) { //gd:Camera3D.set_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFar(far gd.Float) { //gd:Camera3D.set_far
	var frame = callframe.New()
	callframe.Arg(frame, far)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_far, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetNear(near gd.Float) { //gd:Camera3D.set_near
	var frame = callframe.New()
	callframe.Arg(frame, near)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_near, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProjection() gdclass.Camera3DProjectionType { //gd:Camera3D.get_projection
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Camera3DProjectionType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_projection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProjection(mode gdclass.Camera3DProjectionType) { //gd:Camera3D.set_projection
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_projection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetHOffset(offset gd.Float) { //gd:Camera3D.set_h_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_h_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHOffset() gd.Float { //gd:Camera3D.get_h_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_h_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVOffset(offset gd.Float) { //gd:Camera3D.set_v_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_v_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVOffset() gd.Float { //gd:Camera3D.get_v_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_v_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMask(mask gd.Int) { //gd:Camera3D.set_cull_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMask() gd.Int { //gd:Camera3D.get_cull_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironment(env [1]gdclass.Environment) { //gd:Camera3D.set_environment
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() [1]gdclass.Environment { //gd:Camera3D.get_environment
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Environment{gd.PointerWithOwnershipTransferredToGo[gdclass.Environment](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttributes(env [1]gdclass.CameraAttributes) { //gd:Camera3D.set_attributes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(env[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_attributes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttributes() [1]gdclass.CameraAttributes { //gd:Camera3D.get_attributes
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_attributes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CameraAttributes{gd.PointerWithOwnershipTransferredToGo[gdclass.CameraAttributes](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCompositor(compositor [1]gdclass.Compositor) { //gd:Camera3D.set_compositor
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(compositor[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_compositor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompositor() [1]gdclass.Compositor { //gd:Camera3D.get_compositor
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_compositor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Compositor{gd.PointerWithOwnershipTransferredToGo[gdclass.Compositor](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepAspectMode(mode gdclass.Camera3DKeepAspect) { //gd:Camera3D.set_keep_aspect_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_keep_aspect_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetKeepAspectMode() gdclass.Camera3DKeepAspect { //gd:Camera3D.get_keep_aspect_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Camera3DKeepAspect](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_keep_aspect_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDopplerTracking(mode gdclass.Camera3DDopplerTracking) { //gd:Camera3D.set_doppler_tracking
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDopplerTracking() gdclass.Camera3DDopplerTracking { //gd:Camera3D.get_doppler_tracking
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Camera3DDopplerTracking](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the camera's frustum planes in world space units as an array of [Plane]s in the following order: near, far, left, top, right, bottom. Not to be confused with [member frustum_offset].
*/
//go:nosplit
func (self class) GetFrustum() Array.Contains[gd.Plane] { //gd:Camera3D.get_frustum
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_frustum, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Plane]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given position is inside the camera's frustum (the green part of the linked diagram). [url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/camera3d_position_frustum.png]See this diagram[/url] for an overview of position query methods.
*/
//go:nosplit
func (self class) IsPositionInFrustum(world_point gd.Vector3) bool { //gd:Camera3D.is_position_in_frustum
	var frame = callframe.New()
	callframe.Arg(frame, world_point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_is_position_in_frustum, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the camera's RID from the [RenderingServer].
*/
//go:nosplit
func (self class) GetCameraRid() gd.RID { //gd:Camera3D.get_camera_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_camera_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the RID of a pyramid shape encompassing the camera's view frustum, ignoring the camera's near plane. The tip of the pyramid represents the position of the camera.
*/
//go:nosplit
func (self class) GetPyramidShapeRid() gd.RID { //gd:Camera3D.get_pyramid_shape_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_pyramid_shape_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member cull_mask], given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) SetCullMaskValue(layer_number gd.Int, value bool) { //gd:Camera3D.set_cull_mask_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_set_cull_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member cull_mask] is enabled, given a [param layer_number] between 1 and 20.
*/
//go:nosplit
func (self class) GetCullMaskValue(layer_number gd.Int) bool { //gd:Camera3D.get_cull_mask_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera3D.Bind_get_cull_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("Camera3D", func(ptr gd.Object) any { return [1]gdclass.Camera3D{*(*gdclass.Camera3D)(unsafe.Pointer(&ptr))} })
}

type ProjectionType = gdclass.Camera3DProjectionType //gd:Camera3D.ProjectionType

const (
	/*Perspective projection. Objects on the screen becomes smaller when they are far away.*/
	ProjectionPerspective ProjectionType = 0
	/*Orthogonal projection, also known as orthographic projection. Objects remain the same size on the screen no matter how far away they are.*/
	ProjectionOrthogonal ProjectionType = 1
	/*Frustum projection. This mode allows adjusting [member frustum_offset] to create "tilted frustum" effects.*/
	ProjectionFrustum ProjectionType = 2
)

type KeepAspect = gdclass.Camera3DKeepAspect //gd:Camera3D.KeepAspect

const (
	/*Preserves the horizontal aspect ratio; also known as Vert- scaling. This is usually the best option for projects running in portrait mode, as taller aspect ratios will benefit from a wider vertical FOV.*/
	KeepWidth KeepAspect = 0
	/*Preserves the vertical aspect ratio; also known as Hor+ scaling. This is usually the best option for projects running in landscape mode, as wider aspect ratios will automatically benefit from a wider horizontal FOV.*/
	KeepHeight KeepAspect = 1
)

type DopplerTracking = gdclass.Camera3DDopplerTracking //gd:Camera3D.DopplerTracking

const (
	/*Disables [url=https://en.wikipedia.org/wiki/Doppler_effect]Doppler effect[/url] simulation (default).*/
	DopplerTrackingDisabled DopplerTracking = 0
	/*Simulate [url=https://en.wikipedia.org/wiki/Doppler_effect]Doppler effect[/url] by tracking positions of objects that are changed in [code]_process[/code]. Changes in the relative velocity of this camera compared to those objects affect how audio is perceived (changing the audio's [member AudioStreamPlayer3D.pitch_scale]).*/
	DopplerTrackingIdleStep DopplerTracking = 1
	/*Simulate [url=https://en.wikipedia.org/wiki/Doppler_effect]Doppler effect[/url] by tracking positions of objects that are changed in [code]_physics_process[/code]. Changes in the relative velocity of this camera compared to those objects affect how audio is perceived (changing the audio's [member AudioStreamPlayer3D.pitch_scale]).*/
	DopplerTrackingPhysicsStep DopplerTracking = 2
)
