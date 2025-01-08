// Package Line2D provides methods for working with Line2D object instances.
package Line2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This node draws a 2D polyline, i.e. a shape consisting of several points connected by segments. [Line2D] is not a mathematical polyline, i.e. the segments are not infinitely thin. It is intended for rendering and it can be colored and optionally textured.
[b]Warning:[/b] Certain configurations may be impossible to draw nicely, such as very sharp angles. In these situations, the node uses fallback drawing logic to look decent.
[b]Note:[/b] [Line2D] is drawn using a 2D mesh.
*/
type Instance [1]gdclass.Line2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLine2D() Instance
}

/*
Overwrites the position of the point at the given [param index] with the supplied [param position].
*/
func (self Instance) SetPointPosition(index int, position Vector2.XY) {
	class(self).SetPointPosition(gd.Int(index), gd.Vector2(position))
}

/*
Returns the position of the point at index [param index].
*/
func (self Instance) GetPointPosition(index int) Vector2.XY {
	return Vector2.XY(class(self).GetPointPosition(gd.Int(index)))
}

/*
Returns the number of points in the polyline.
*/
func (self Instance) GetPointCount() int {
	return int(int(class(self).GetPointCount()))
}

/*
Adds a point with the specified [param position] relative to the polyline's own position. If no [param index] is provided, the new point will be added to the end of the points array.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. The indices of the points after the new point get increased by 1. The provided [param index] must not exceed the number of existing points in the polyline. See [method get_point_count].
*/
func (self Instance) AddPoint(position Vector2.XY) {
	class(self).AddPoint(gd.Vector2(position), gd.Int(-1))
}

/*
Removes the point at index [param index] from the polyline.
*/
func (self Instance) RemovePoint(index int) {
	class(self).RemovePoint(gd.Int(index))
}

/*
Removes all points from the polyline, making it empty.
*/
func (self Instance) ClearPoints() {
	class(self).ClearPoints()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Line2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Line2D"))
	casted := Instance{*(*gdclass.Line2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Points() []Vector2.XY {
	return []Vector2.XY(class(self).GetPoints().AsSlice())
}

func (self Instance) SetPoints(value []Vector2.XY) {
	class(self).SetPoints(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&value))))
}

func (self Instance) Closed() bool {
	return bool(class(self).IsClosed())
}

func (self Instance) SetClosed(value bool) {
	class(self).SetClosed(value)
}

func (self Instance) Width() Float.X {
	return Float.X(Float.X(class(self).GetWidth()))
}

func (self Instance) SetWidth(value Float.X) {
	class(self).SetWidth(gd.Float(value))
}

func (self Instance) WidthCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetCurve())
}

func (self Instance) SetWidthCurve(value [1]gdclass.Curve) {
	class(self).SetCurve(value)
}

func (self Instance) DefaultColor() Color.RGBA {
	return Color.RGBA(class(self).GetDefaultColor())
}

func (self Instance) SetDefaultColor(value Color.RGBA) {
	class(self).SetDefaultColor(gd.Color(value))
}

func (self Instance) Gradient() [1]gdclass.Gradient {
	return [1]gdclass.Gradient(class(self).GetGradient())
}

func (self Instance) SetGradient(value [1]gdclass.Gradient) {
	class(self).SetGradient(value)
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) TextureMode() gdclass.Line2DLineTextureMode {
	return gdclass.Line2DLineTextureMode(class(self).GetTextureMode())
}

func (self Instance) SetTextureMode(value gdclass.Line2DLineTextureMode) {
	class(self).SetTextureMode(value)
}

func (self Instance) JointMode() gdclass.Line2DLineJointMode {
	return gdclass.Line2DLineJointMode(class(self).GetJointMode())
}

func (self Instance) SetJointMode(value gdclass.Line2DLineJointMode) {
	class(self).SetJointMode(value)
}

func (self Instance) BeginCapMode() gdclass.Line2DLineCapMode {
	return gdclass.Line2DLineCapMode(class(self).GetBeginCapMode())
}

func (self Instance) SetBeginCapMode(value gdclass.Line2DLineCapMode) {
	class(self).SetBeginCapMode(value)
}

func (self Instance) EndCapMode() gdclass.Line2DLineCapMode {
	return gdclass.Line2DLineCapMode(class(self).GetEndCapMode())
}

func (self Instance) SetEndCapMode(value gdclass.Line2DLineCapMode) {
	class(self).SetEndCapMode(value)
}

func (self Instance) SharpLimit() Float.X {
	return Float.X(Float.X(class(self).GetSharpLimit()))
}

func (self Instance) SetSharpLimit(value Float.X) {
	class(self).SetSharpLimit(gd.Float(value))
}

func (self Instance) RoundPrecision() int {
	return int(int(class(self).GetRoundPrecision()))
}

func (self Instance) SetRoundPrecision(value int) {
	class(self).SetRoundPrecision(gd.Int(value))
}

func (self Instance) Antialiased() bool {
	return bool(class(self).GetAntialiased())
}

func (self Instance) SetAntialiased(value bool) {
	class(self).SetAntialiased(value)
}

//go:nosplit
func (self class) SetPoints(points gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPoints() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Overwrites the position of the point at the given [param index] with the supplied [param position].
*/
//go:nosplit
func (self class) SetPointPosition(index gd.Int, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the position of the point at index [param index].
*/
//go:nosplit
func (self class) GetPointPosition(index gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of points in the polyline.
*/
//go:nosplit
func (self class) GetPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a point with the specified [param position] relative to the polyline's own position. If no [param index] is provided, the new point will be added to the end of the points array.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. The indices of the points after the new point get increased by 1. The provided [param index] must not exceed the number of existing points in the polyline. See [method get_point_count].
*/
//go:nosplit
func (self class) AddPoint(position gd.Vector2, index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the point at index [param index] from the polyline.
*/
//go:nosplit
func (self class) RemovePoint(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all points from the polyline, making it empty.
*/
//go:nosplit
func (self class) ClearPoints() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetClosed(closed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, closed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_closed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsClosed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_is_closed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_default_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_default_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGradient(color [1]gdclass.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(color[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGradient() [1]gdclass.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMode(mode gdclass.Line2DLineTextureMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMode() gdclass.Line2DLineTextureMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Line2DLineTextureMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointMode(mode gdclass.Line2DLineJointMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_joint_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointMode() gdclass.Line2DLineJointMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Line2DLineJointMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_joint_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBeginCapMode(mode gdclass.Line2DLineCapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_begin_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBeginCapMode() gdclass.Line2DLineCapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Line2DLineCapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_begin_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEndCapMode(mode gdclass.Line2DLineCapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_end_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEndCapMode() gdclass.Line2DLineCapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Line2DLineCapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_end_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSharpLimit(limit gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_sharp_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSharpLimit() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_sharp_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoundPrecision(precision gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, precision)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_round_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoundPrecision() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_round_precision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAntialiased(antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiased() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLine2D() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLine2D() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Line2D", func(ptr gd.Object) any { return [1]gdclass.Line2D{*(*gdclass.Line2D)(unsafe.Pointer(&ptr))} })
}

type LineJointMode = gdclass.Line2DLineJointMode

const (
	/*Makes the polyline's joints pointy, connecting the sides of the two segments by extending them until they intersect. If the rotation of a joint is too big (based on [member sharp_limit]), the joint falls back to [constant LINE_JOINT_BEVEL] to prevent very long miters.*/
	LineJointSharp LineJointMode = 0
	/*Makes the polyline's joints bevelled/chamfered, connecting the sides of the two segments with a simple line.*/
	LineJointBevel LineJointMode = 1
	/*Makes the polyline's joints rounded, connecting the sides of the two segments with an arc. The detail of this arc depends on [member round_precision].*/
	LineJointRound LineJointMode = 2
)

type LineCapMode = gdclass.Line2DLineCapMode

const (
	/*Draws no line cap.*/
	LineCapNone LineCapMode = 0
	/*Draws the line cap as a box, slightly extending the first/last segment.*/
	LineCapBox LineCapMode = 1
	/*Draws the line cap as a semicircle attached to the first/last segment.*/
	LineCapRound LineCapMode = 2
)

type LineTextureMode = gdclass.Line2DLineTextureMode

const (
	/*Takes the left pixels of the texture and renders them over the whole polyline.*/
	LineTextureNone LineTextureMode = 0
	/*Tiles the texture over the polyline. [member CanvasItem.texture_repeat] of the [Line2D] node must be [constant CanvasItem.TEXTURE_REPEAT_ENABLED] or [constant CanvasItem.TEXTURE_REPEAT_MIRROR] for it to work properly.*/
	LineTextureTile LineTextureMode = 1
	/*Stretches the texture across the polyline. [member CanvasItem.texture_repeat] of the [Line2D] node must be [constant CanvasItem.TEXTURE_REPEAT_DISABLED] for best results.*/
	LineTextureStretch LineTextureMode = 2
)
