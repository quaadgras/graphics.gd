// Package Camera2D provides methods for working with Camera2D object instances.
package Camera2D

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
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

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
Camera node for 2D scenes. It forces the screen (current layer) to scroll following this node. This makes it easier (and faster) to program scrollable scenes than manually changing the position of [CanvasItem]-based nodes.
Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the camera will register in the global viewport.
This node is intended to be a simple helper to get things going quickly, but more functionality may be desired to change how the camera works. To make your own custom camera node, inherit it from [Node2D] and change the transform of the canvas by setting [member Viewport.canvas_transform] in [Viewport] (you can obtain the current [Viewport] by using [method Node.get_viewport]).
Note that the [Camera2D] node's [code]position[/code] doesn't represent the actual position of the screen, which may differ due to applied smoothing or limits. You can use [method get_screen_center_position] to get the real position.
*/
type Instance [1]gdclass.Camera2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCamera2D() Instance
}

/*
Forces this [Camera2D] to become the current active one. [member enabled] must be [code]true[/code].
*/
func (self Instance) MakeCurrent() { //gd:Camera2D.make_current
	class(self).MakeCurrent()
}

/*
Returns [code]true[/code] if this [Camera2D] is the active camera (see [method Viewport.get_camera_2d]).
*/
func (self Instance) IsCurrent() bool { //gd:Camera2D.is_current
	return bool(class(self).IsCurrent())
}

/*
Returns this camera's target position, in global coordinates.
[b]Note:[/b] The returned value is not the same as [member Node2D.global_position], as it is affected by the drag properties. It is also not the same as the current position if [member position_smoothing_enabled] is [code]true[/code] (see [method get_screen_center_position]).
*/
func (self Instance) GetTargetPosition() Vector2.XY { //gd:Camera2D.get_target_position
	return Vector2.XY(class(self).GetTargetPosition())
}

/*
Returns the center of the screen from this camera's point of view, in global coordinates.
[b]Note:[/b] The exact targeted position of the camera may be different. See [method get_target_position].
*/
func (self Instance) GetScreenCenterPosition() Vector2.XY { //gd:Camera2D.get_screen_center_position
	return Vector2.XY(class(self).GetScreenCenterPosition())
}

/*
Forces the camera to update scroll immediately.
*/
func (self Instance) ForceUpdateScroll() { //gd:Camera2D.force_update_scroll
	class(self).ForceUpdateScroll()
}

/*
Sets the camera's position immediately to its current smoothing destination.
This method has no effect if [member position_smoothing_enabled] is [code]false[/code].
*/
func (self Instance) ResetSmoothing() { //gd:Camera2D.reset_smoothing
	class(self).ResetSmoothing()
}

/*
Aligns the camera to the tracked node.
*/
func (self Instance) Align() { //gd:Camera2D.align
	class(self).Align()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Camera2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Camera2D"))
	casted := Instance{*(*gdclass.Camera2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(gd.Vector2(value))
}

func (self Instance) AnchorMode() gdclass.Camera2DAnchorMode {
	return gdclass.Camera2DAnchorMode(class(self).GetAnchorMode())
}

func (self Instance) SetAnchorMode(value gdclass.Camera2DAnchorMode) {
	class(self).SetAnchorMode(value)
}

func (self Instance) IgnoreRotation() bool {
	return bool(class(self).IsIgnoringRotation())
}

func (self Instance) SetIgnoreRotation(value bool) {
	class(self).SetIgnoreRotation(value)
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) Zoom() Vector2.XY {
	return Vector2.XY(class(self).GetZoom())
}

func (self Instance) SetZoom(value Vector2.XY) {
	class(self).SetZoom(gd.Vector2(value))
}

func (self Instance) CustomViewport() [1]gdclass.Node {
	return [1]gdclass.Node(class(self).GetCustomViewport())
}

func (self Instance) SetCustomViewport(value [1]gdclass.Node) {
	class(self).SetCustomViewport(value)
}

func (self Instance) ProcessCallback() gdclass.Camera2DCamera2DProcessCallback {
	return gdclass.Camera2DCamera2DProcessCallback(class(self).GetProcessCallback())
}

func (self Instance) SetProcessCallback(value gdclass.Camera2DCamera2DProcessCallback) {
	class(self).SetProcessCallback(value)
}

func (self Instance) LimitLeft() int {
	return int(int(class(self).GetLimit(0)))
}

func (self Instance) SetLimitLeft(value int) {
	class(self).SetLimit(0, gd.Int(value))
}

func (self Instance) LimitTop() int {
	return int(int(class(self).GetLimit(1)))
}

func (self Instance) SetLimitTop(value int) {
	class(self).SetLimit(1, gd.Int(value))
}

func (self Instance) LimitRight() int {
	return int(int(class(self).GetLimit(2)))
}

func (self Instance) SetLimitRight(value int) {
	class(self).SetLimit(2, gd.Int(value))
}

func (self Instance) LimitBottom() int {
	return int(int(class(self).GetLimit(3)))
}

func (self Instance) SetLimitBottom(value int) {
	class(self).SetLimit(3, gd.Int(value))
}

func (self Instance) LimitSmoothed() bool {
	return bool(class(self).IsLimitSmoothingEnabled())
}

func (self Instance) SetLimitSmoothed(value bool) {
	class(self).SetLimitSmoothingEnabled(value)
}

func (self Instance) PositionSmoothingEnabled() bool {
	return bool(class(self).IsPositionSmoothingEnabled())
}

func (self Instance) SetPositionSmoothingEnabled(value bool) {
	class(self).SetPositionSmoothingEnabled(value)
}

func (self Instance) PositionSmoothingSpeed() Float.X {
	return Float.X(Float.X(class(self).GetPositionSmoothingSpeed()))
}

func (self Instance) SetPositionSmoothingSpeed(value Float.X) {
	class(self).SetPositionSmoothingSpeed(gd.Float(value))
}

func (self Instance) RotationSmoothingEnabled() bool {
	return bool(class(self).IsRotationSmoothingEnabled())
}

func (self Instance) SetRotationSmoothingEnabled(value bool) {
	class(self).SetRotationSmoothingEnabled(value)
}

func (self Instance) RotationSmoothingSpeed() Float.X {
	return Float.X(Float.X(class(self).GetRotationSmoothingSpeed()))
}

func (self Instance) SetRotationSmoothingSpeed(value Float.X) {
	class(self).SetRotationSmoothingSpeed(gd.Float(value))
}

func (self Instance) DragHorizontalEnabled() bool {
	return bool(class(self).IsDragHorizontalEnabled())
}

func (self Instance) SetDragHorizontalEnabled(value bool) {
	class(self).SetDragHorizontalEnabled(value)
}

func (self Instance) DragVerticalEnabled() bool {
	return bool(class(self).IsDragVerticalEnabled())
}

func (self Instance) SetDragVerticalEnabled(value bool) {
	class(self).SetDragVerticalEnabled(value)
}

func (self Instance) DragHorizontalOffset() Float.X {
	return Float.X(Float.X(class(self).GetDragHorizontalOffset()))
}

func (self Instance) SetDragHorizontalOffset(value Float.X) {
	class(self).SetDragHorizontalOffset(gd.Float(value))
}

func (self Instance) DragVerticalOffset() Float.X {
	return Float.X(Float.X(class(self).GetDragVerticalOffset()))
}

func (self Instance) SetDragVerticalOffset(value Float.X) {
	class(self).SetDragVerticalOffset(gd.Float(value))
}

func (self Instance) DragLeftMargin() Float.X {
	return Float.X(Float.X(class(self).GetDragMargin(0)))
}

func (self Instance) SetDragLeftMargin(value Float.X) {
	class(self).SetDragMargin(0, gd.Float(value))
}

func (self Instance) DragTopMargin() Float.X {
	return Float.X(Float.X(class(self).GetDragMargin(1)))
}

func (self Instance) SetDragTopMargin(value Float.X) {
	class(self).SetDragMargin(1, gd.Float(value))
}

func (self Instance) DragRightMargin() Float.X {
	return Float.X(Float.X(class(self).GetDragMargin(2)))
}

func (self Instance) SetDragRightMargin(value Float.X) {
	class(self).SetDragMargin(2, gd.Float(value))
}

func (self Instance) DragBottomMargin() Float.X {
	return Float.X(Float.X(class(self).GetDragMargin(3)))
}

func (self Instance) SetDragBottomMargin(value Float.X) {
	class(self).SetDragMargin(3, gd.Float(value))
}

func (self Instance) EditorDrawScreen() bool {
	return bool(class(self).IsScreenDrawingEnabled())
}

func (self Instance) SetEditorDrawScreen(value bool) {
	class(self).SetScreenDrawingEnabled(value)
}

func (self Instance) EditorDrawLimits() bool {
	return bool(class(self).IsLimitDrawingEnabled())
}

func (self Instance) SetEditorDrawLimits(value bool) {
	class(self).SetLimitDrawingEnabled(value)
}

func (self Instance) EditorDrawDragMargin() bool {
	return bool(class(self).IsMarginDrawingEnabled())
}

func (self Instance) SetEditorDrawDragMargin(value bool) {
	class(self).SetMarginDrawingEnabled(value)
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) { //gd:Camera2D.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 { //gd:Camera2D.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAnchorMode(anchor_mode gdclass.Camera2DAnchorMode) { //gd:Camera2D.set_anchor_mode
	var frame = callframe.New()
	callframe.Arg(frame, anchor_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_anchor_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnchorMode() gdclass.Camera2DAnchorMode { //gd:Camera2D.get_anchor_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Camera2DAnchorMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_anchor_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIgnoreRotation(ignore bool) { //gd:Camera2D.set_ignore_rotation
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_ignore_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsIgnoringRotation() bool { //gd:Camera2D.is_ignoring_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_ignoring_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProcessCallback(mode gdclass.Camera2DCamera2DProcessCallback) { //gd:Camera2D.set_process_callback
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProcessCallback() gdclass.Camera2DCamera2DProcessCallback { //gd:Camera2D.get_process_callback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Camera2DCamera2DProcessCallback](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) { //gd:Camera2D.set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool { //gd:Camera2D.is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces this [Camera2D] to become the current active one. [member enabled] must be [code]true[/code].
*/
//go:nosplit
func (self class) MakeCurrent() { //gd:Camera2D.make_current
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if this [Camera2D] is the active camera (see [method Viewport.get_camera_2d]).
*/
//go:nosplit
func (self class) IsCurrent() bool { //gd:Camera2D.is_current
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the camera limit for the specified [enum Side]. See also [member limit_bottom], [member limit_top], [member limit_left], and [member limit_right].
*/
//go:nosplit
func (self class) SetLimit(margin Side, limit gd.Int) { //gd:Camera2D.set_limit
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, limit)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_limit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the camera limit for the specified [enum Side]. See also [member limit_bottom], [member limit_top], [member limit_left], and [member limit_right].
*/
//go:nosplit
func (self class) GetLimit(margin Side) gd.Int { //gd:Camera2D.get_limit
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_limit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLimitSmoothingEnabled(limit_smoothing_enabled bool) { //gd:Camera2D.set_limit_smoothing_enabled
	var frame = callframe.New()
	callframe.Arg(frame, limit_smoothing_enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_limit_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLimitSmoothingEnabled() bool { //gd:Camera2D.is_limit_smoothing_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_limit_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragVerticalEnabled(enabled bool) { //gd:Camera2D.set_drag_vertical_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_drag_vertical_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDragVerticalEnabled() bool { //gd:Camera2D.is_drag_vertical_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_drag_vertical_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragHorizontalEnabled(enabled bool) { //gd:Camera2D.set_drag_horizontal_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_drag_horizontal_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDragHorizontalEnabled() bool { //gd:Camera2D.is_drag_horizontal_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_drag_horizontal_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragVerticalOffset(offset gd.Float) { //gd:Camera2D.set_drag_vertical_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_drag_vertical_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDragVerticalOffset() gd.Float { //gd:Camera2D.get_drag_vertical_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_drag_vertical_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragHorizontalOffset(offset gd.Float) { //gd:Camera2D.set_drag_horizontal_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_drag_horizontal_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDragHorizontalOffset() gd.Float { //gd:Camera2D.get_drag_horizontal_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_drag_horizontal_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the specified [enum Side]'s margin. See also [member drag_bottom_margin], [member drag_top_margin], [member drag_left_margin], and [member drag_right_margin].
*/
//go:nosplit
func (self class) SetDragMargin(margin Side, drag_margin gd.Float) { //gd:Camera2D.set_drag_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, drag_margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_drag_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the specified [enum Side]'s margin. See also [member drag_bottom_margin], [member drag_top_margin], [member drag_left_margin], and [member drag_right_margin].
*/
//go:nosplit
func (self class) GetDragMargin(margin Side) gd.Float { //gd:Camera2D.get_drag_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_drag_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns this camera's target position, in global coordinates.
[b]Note:[/b] The returned value is not the same as [member Node2D.global_position], as it is affected by the drag properties. It is also not the same as the current position if [member position_smoothing_enabled] is [code]true[/code] (see [method get_screen_center_position]).
*/
//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 { //gd:Camera2D.get_target_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the center of the screen from this camera's point of view, in global coordinates.
[b]Note:[/b] The exact targeted position of the camera may be different. See [method get_target_position].
*/
//go:nosplit
func (self class) GetScreenCenterPosition() gd.Vector2 { //gd:Camera2D.get_screen_center_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_screen_center_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoom(zoom gd.Vector2) { //gd:Camera2D.set_zoom
	var frame = callframe.New()
	callframe.Arg(frame, zoom)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_zoom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoom() gd.Vector2 { //gd:Camera2D.get_zoom
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_zoom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomViewport(viewport [1]gdclass.Node) { //gd:Camera2D.set_custom_viewport
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(viewport[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomViewport() [1]gdclass.Node { //gd:Camera2D.get_custom_viewport
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerMustAssertInstanceID[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPositionSmoothingSpeed(position_smoothing_speed gd.Float) { //gd:Camera2D.set_position_smoothing_speed
	var frame = callframe.New()
	callframe.Arg(frame, position_smoothing_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_position_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPositionSmoothingSpeed() gd.Float { //gd:Camera2D.get_position_smoothing_speed
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_position_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPositionSmoothingEnabled(position_smoothing_speed bool) { //gd:Camera2D.set_position_smoothing_enabled
	var frame = callframe.New()
	callframe.Arg(frame, position_smoothing_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_position_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPositionSmoothingEnabled() bool { //gd:Camera2D.is_position_smoothing_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_position_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationSmoothingEnabled(enabled bool) { //gd:Camera2D.set_rotation_smoothing_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_rotation_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRotationSmoothingEnabled() bool { //gd:Camera2D.is_rotation_smoothing_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_rotation_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRotationSmoothingSpeed(speed gd.Float) { //gd:Camera2D.set_rotation_smoothing_speed
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_rotation_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRotationSmoothingSpeed() gd.Float { //gd:Camera2D.get_rotation_smoothing_speed
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_get_rotation_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces the camera to update scroll immediately.
*/
//go:nosplit
func (self class) ForceUpdateScroll() { //gd:Camera2D.force_update_scroll
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_force_update_scroll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the camera's position immediately to its current smoothing destination.
This method has no effect if [member position_smoothing_enabled] is [code]false[/code].
*/
//go:nosplit
func (self class) ResetSmoothing() { //gd:Camera2D.reset_smoothing
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_reset_smoothing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Aligns the camera to the tracked node.
*/
//go:nosplit
func (self class) Align() { //gd:Camera2D.align
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_align, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetScreenDrawingEnabled(screen_drawing_enabled bool) { //gd:Camera2D.set_screen_drawing_enabled
	var frame = callframe.New()
	callframe.Arg(frame, screen_drawing_enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_screen_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsScreenDrawingEnabled() bool { //gd:Camera2D.is_screen_drawing_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_screen_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLimitDrawingEnabled(limit_drawing_enabled bool) { //gd:Camera2D.set_limit_drawing_enabled
	var frame = callframe.New()
	callframe.Arg(frame, limit_drawing_enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_limit_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLimitDrawingEnabled() bool { //gd:Camera2D.is_limit_drawing_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_limit_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMarginDrawingEnabled(margin_drawing_enabled bool) { //gd:Camera2D.set_margin_drawing_enabled
	var frame = callframe.New()
	callframe.Arg(frame, margin_drawing_enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_set_margin_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMarginDrawingEnabled() bool { //gd:Camera2D.is_margin_drawing_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Camera2D.Bind_is_margin_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCamera2D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCamera2D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("Camera2D", func(ptr gd.Object) any { return [1]gdclass.Camera2D{*(*gdclass.Camera2D)(unsafe.Pointer(&ptr))} })
}

type AnchorMode = gdclass.Camera2DAnchorMode //gd:Camera2D.AnchorMode

const (
	/*The camera's position is fixed so that the top-left corner is always at the origin.*/
	AnchorModeFixedTopLeft AnchorMode = 0
	/*The camera's position takes into account vertical/horizontal offsets and the screen size.*/
	AnchorModeDragCenter AnchorMode = 1
)

type Camera2DProcessCallback = gdclass.Camera2DCamera2DProcessCallback //gd:Camera2D.Camera2DProcessCallback

const (
	/*The camera updates during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	Camera2dProcessPhysics Camera2DProcessCallback = 0
	/*The camera updates during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	Camera2dProcessIdle Camera2DProcessCallback = 1
)

type Side int

const (
	/*Left side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideLeft Side = 0
	/*Top side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideTop Side = 1
	/*Right side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideRight Side = 2
	/*Bottom side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideBottom Side = 3
)
