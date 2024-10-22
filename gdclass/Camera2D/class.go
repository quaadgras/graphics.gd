package Camera2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Camera node for 2D scenes. It forces the screen (current layer) to scroll following this node. This makes it easier (and faster) to program scrollable scenes than manually changing the position of [CanvasItem]-based nodes.
Cameras register themselves in the nearest [Viewport] node (when ascending the tree). Only one camera can be active per viewport. If no viewport is available ascending the tree, the camera will register in the global viewport.
This node is intended to be a simple helper to get things going quickly, but more functionality may be desired to change how the camera works. To make your own custom camera node, inherit it from [Node2D] and change the transform of the canvas by setting [member Viewport.canvas_transform] in [Viewport] (you can obtain the current [Viewport] by using [method Node.get_viewport]).
Note that the [Camera2D] node's [code]position[/code] doesn't represent the actual position of the screen, which may differ due to applied smoothing or limits. You can use [method get_screen_center_position] to get the real position.

*/
type Go [1]classdb.Camera2D

/*
Forces this [Camera2D] to become the current active one. [member enabled] must be [code]true[/code].
*/
func (self Go) MakeCurrent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MakeCurrent()
}

/*
Returns [code]true[/code] if this [Camera2D] is the active camera (see [method Viewport.get_camera_2d]).
*/
func (self Go) IsCurrent() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCurrent())
}

/*
Returns this camera's target position, in global coordinates.
[b]Note:[/b] The returned value is not the same as [member Node2D.global_position], as it is affected by the drag properties. It is also not the same as the current position if [member position_smoothing_enabled] is [code]true[/code] (see [method get_screen_center_position]).
*/
func (self Go) GetTargetPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetTargetPosition())
}

/*
Returns the center of the screen from this camera's point of view, in global coordinates.
[b]Note:[/b] The exact targeted position of the camera may be different. See [method get_target_position].
*/
func (self Go) GetScreenCenterPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetScreenCenterPosition())
}

/*
Forces the camera to update scroll immediately.
*/
func (self Go) ForceUpdateScroll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ForceUpdateScroll()
}

/*
Sets the camera's position immediately to its current smoothing destination.
This method has no effect if [member position_smoothing_enabled] is [code]false[/code].
*/
func (self Go) ResetSmoothing() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ResetSmoothing()
}

/*
Aligns the camera to the tracked node.
*/
func (self Go) Align() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Align()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Camera2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Camera2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Offset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetOffset())
}

func (self Go) SetOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffset(value)
}

func (self Go) AnchorMode() classdb.Camera2DAnchorMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Camera2DAnchorMode(class(self).GetAnchorMode())
}

func (self Go) SetAnchorMode(value classdb.Camera2DAnchorMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAnchorMode(value)
}

func (self Go) IgnoreRotation() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsIgnoringRotation())
}

func (self Go) SetIgnoreRotation(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIgnoreRotation(value)
}

func (self Go) Enabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEnabled())
}

func (self Go) SetEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEnabled(value)
}

func (self Go) Zoom() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetZoom())
}

func (self Go) SetZoom(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetZoom(value)
}

func (self Go) CustomViewport() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Node(class(self).GetCustomViewport(gc))
}

func (self Go) SetCustomViewport(value gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCustomViewport(value)
}

func (self Go) ProcessCallback() classdb.Camera2DCamera2DProcessCallback {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.Camera2DCamera2DProcessCallback(class(self).GetProcessCallback())
}

func (self Go) SetProcessCallback(value classdb.Camera2DCamera2DProcessCallback) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProcessCallback(value)
}

func (self Go) LimitLeft() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLimit(0)))
}

func (self Go) SetLimitLeft(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLimit(0, gd.Int(value))
}

func (self Go) LimitTop() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLimit(1)))
}

func (self Go) SetLimitTop(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLimit(1, gd.Int(value))
}

func (self Go) LimitRight() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLimit(2)))
}

func (self Go) SetLimitRight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLimit(2, gd.Int(value))
}

func (self Go) LimitBottom() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLimit(3)))
}

func (self Go) SetLimitBottom(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLimit(3, gd.Int(value))
}

func (self Go) LimitSmoothed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsLimitSmoothingEnabled())
}

func (self Go) SetLimitSmoothed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLimitSmoothingEnabled(value)
}

func (self Go) PositionSmoothingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPositionSmoothingEnabled())
}

func (self Go) SetPositionSmoothingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPositionSmoothingEnabled(value)
}

func (self Go) PositionSmoothingSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPositionSmoothingSpeed()))
}

func (self Go) SetPositionSmoothingSpeed(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPositionSmoothingSpeed(gd.Float(value))
}

func (self Go) RotationSmoothingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRotationSmoothingEnabled())
}

func (self Go) SetRotationSmoothingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotationSmoothingEnabled(value)
}

func (self Go) RotationSmoothingSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRotationSmoothingSpeed()))
}

func (self Go) SetRotationSmoothingSpeed(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRotationSmoothingSpeed(gd.Float(value))
}

func (self Go) DragHorizontalEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDragHorizontalEnabled())
}

func (self Go) SetDragHorizontalEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragHorizontalEnabled(value)
}

func (self Go) DragVerticalEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDragVerticalEnabled())
}

func (self Go) SetDragVerticalEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragVerticalEnabled(value)
}

func (self Go) DragHorizontalOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragHorizontalOffset()))
}

func (self Go) SetDragHorizontalOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragHorizontalOffset(gd.Float(value))
}

func (self Go) DragVerticalOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragVerticalOffset()))
}

func (self Go) SetDragVerticalOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragVerticalOffset(gd.Float(value))
}

func (self Go) DragLeftMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragMargin(0)))
}

func (self Go) SetDragLeftMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragMargin(0, gd.Float(value))
}

func (self Go) DragTopMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragMargin(1)))
}

func (self Go) SetDragTopMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragMargin(1, gd.Float(value))
}

func (self Go) DragRightMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragMargin(2)))
}

func (self Go) SetDragRightMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragMargin(2, gd.Float(value))
}

func (self Go) DragBottomMargin() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDragMargin(3)))
}

func (self Go) SetDragBottomMargin(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragMargin(3, gd.Float(value))
}

func (self Go) EditorDrawScreen() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsScreenDrawingEnabled())
}

func (self Go) SetEditorDrawScreen(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScreenDrawingEnabled(value)
}

func (self Go) EditorDrawLimits() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsLimitDrawingEnabled())
}

func (self Go) SetEditorDrawLimits(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLimitDrawingEnabled(value)
}

func (self Go) EditorDrawDragMargin() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMarginDrawingEnabled())
}

func (self Go) SetEditorDrawDragMargin(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMarginDrawingEnabled(value)
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnchorMode(anchor_mode classdb.Camera2DAnchorMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, anchor_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_anchor_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnchorMode() classdb.Camera2DAnchorMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera2DAnchorMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_anchor_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIgnoreRotation(ignore bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_ignore_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIgnoringRotation() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_ignoring_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProcessCallback(mode classdb.Camera2DCamera2DProcessCallback)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProcessCallback() classdb.Camera2DCamera2DProcessCallback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Camera2DCamera2DProcessCallback](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces this [Camera2D] to become the current active one. [member enabled] must be [code]true[/code].
*/
//go:nosplit
func (self class) MakeCurrent()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_make_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if this [Camera2D] is the active camera (see [method Viewport.get_camera_2d]).
*/
//go:nosplit
func (self class) IsCurrent() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the camera limit for the specified [enum Side]. See also [member limit_bottom], [member limit_top], [member limit_left], and [member limit_right].
*/
//go:nosplit
func (self class) SetLimit(margin gd.Side, limit gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the camera limit for the specified [enum Side]. See also [member limit_bottom], [member limit_top], [member limit_left], and [member limit_right].
*/
//go:nosplit
func (self class) GetLimit(margin gd.Side) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLimitSmoothingEnabled(limit_smoothing_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, limit_smoothing_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_limit_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLimitSmoothingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_limit_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragVerticalEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_drag_vertical_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDragVerticalEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_drag_vertical_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragHorizontalEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_drag_horizontal_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDragHorizontalEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_drag_horizontal_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragVerticalOffset(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_drag_vertical_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragVerticalOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_drag_vertical_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragHorizontalOffset(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_drag_horizontal_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDragHorizontalOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_drag_horizontal_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the specified [enum Side]'s margin. See also [member drag_bottom_margin], [member drag_top_margin], [member drag_left_margin], and [member drag_right_margin].
*/
//go:nosplit
func (self class) SetDragMargin(margin gd.Side, drag_margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, drag_margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_drag_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the specified [enum Side]'s margin. See also [member drag_bottom_margin], [member drag_top_margin], [member drag_left_margin], and [member drag_right_margin].
*/
//go:nosplit
func (self class) GetDragMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_drag_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns this camera's target position, in global coordinates.
[b]Note:[/b] The returned value is not the same as [member Node2D.global_position], as it is affected by the drag properties. It is also not the same as the current position if [member position_smoothing_enabled] is [code]true[/code] (see [method get_screen_center_position]).
*/
//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the center of the screen from this camera's point of view, in global coordinates.
[b]Note:[/b] The exact targeted position of the camera may be different. See [method get_target_position].
*/
//go:nosplit
func (self class) GetScreenCenterPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_screen_center_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetZoom(zoom gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, zoom)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetZoom() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_zoom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomViewport(viewport gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(viewport[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomViewport(ctx gd.Lifetime) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_custom_viewport, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPositionSmoothingSpeed(position_smoothing_speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position_smoothing_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_position_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPositionSmoothingSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_position_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPositionSmoothingEnabled(position_smoothing_speed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position_smoothing_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_position_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPositionSmoothingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_position_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotationSmoothingEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_rotation_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRotationSmoothingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_rotation_smoothing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRotationSmoothingSpeed(speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_rotation_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRotationSmoothingSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_get_rotation_smoothing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces the camera to update scroll immediately.
*/
//go:nosplit
func (self class) ForceUpdateScroll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_force_update_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the camera's position immediately to its current smoothing destination.
This method has no effect if [member position_smoothing_enabled] is [code]false[/code].
*/
//go:nosplit
func (self class) ResetSmoothing()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_reset_smoothing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Aligns the camera to the tracked node.
*/
//go:nosplit
func (self class) Align()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetScreenDrawingEnabled(screen_drawing_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, screen_drawing_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_screen_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsScreenDrawingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_screen_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLimitDrawingEnabled(limit_drawing_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, limit_drawing_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_limit_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLimitDrawingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_limit_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMarginDrawingEnabled(margin_drawing_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin_drawing_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_set_margin_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMarginDrawingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Camera2D.Bind_is_margin_drawing_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCamera2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCamera2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("Camera2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AnchorMode = classdb.Camera2DAnchorMode

const (
/*The camera's position is fixed so that the top-left corner is always at the origin.*/
	AnchorModeFixedTopLeft AnchorMode = 0
/*The camera's position takes into account vertical/horizontal offsets and the screen size.*/
	AnchorModeDragCenter AnchorMode = 1
)
type Camera2DProcessCallback = classdb.Camera2DCamera2DProcessCallback

const (
/*The camera updates during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	Camera2dProcessPhysics Camera2DProcessCallback = 0
/*The camera updates during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	Camera2dProcessIdle Camera2DProcessCallback = 1
)
