package Line2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This node draws a 2D polyline, i.e. a shape consisting of several points connected by segments. [Line2D] is not a mathematical polyline, i.e. the segments are not infinitely thin. It is intended for rendering and it can be colored and optionally textured.
[b]Warning:[/b] Certain configurations may be impossible to draw nicely, such as very sharp angles. In these situations, the node uses fallback drawing logic to look decent.
[b]Note:[/b] [Line2D] is drawn using a 2D mesh.
*/
type Instance [1]classdb.Line2D

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
type class [1]classdb.Line2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Line2D"))
	return Instance{classdb.Line2D(object)}
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

func (self Instance) WidthCurve() objects.Curve {
	return objects.Curve(class(self).GetCurve())
}

func (self Instance) SetWidthCurve(value objects.Curve) {
	class(self).SetCurve(value)
}

func (self Instance) DefaultColor() Color.RGBA {
	return Color.RGBA(class(self).GetDefaultColor())
}

func (self Instance) SetDefaultColor(value Color.RGBA) {
	class(self).SetDefaultColor(gd.Color(value))
}

func (self Instance) Gradient() objects.Gradient {
	return objects.Gradient(class(self).GetGradient())
}

func (self Instance) SetGradient(value objects.Gradient) {
	class(self).SetGradient(value)
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) TextureMode() classdb.Line2DLineTextureMode {
	return classdb.Line2DLineTextureMode(class(self).GetTextureMode())
}

func (self Instance) SetTextureMode(value classdb.Line2DLineTextureMode) {
	class(self).SetTextureMode(value)
}

func (self Instance) JointMode() classdb.Line2DLineJointMode {
	return classdb.Line2DLineJointMode(class(self).GetJointMode())
}

func (self Instance) SetJointMode(value classdb.Line2DLineJointMode) {
	class(self).SetJointMode(value)
}

func (self Instance) BeginCapMode() classdb.Line2DLineCapMode {
	return classdb.Line2DLineCapMode(class(self).GetBeginCapMode())
}

func (self Instance) SetBeginCapMode(value classdb.Line2DLineCapMode) {
	class(self).SetBeginCapMode(value)
}

func (self Instance) EndCapMode() classdb.Line2DLineCapMode {
	return classdb.Line2DLineCapMode(class(self).GetEndCapMode())
}

func (self Instance) SetEndCapMode(value classdb.Line2DLineCapMode) {
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
func (self class) SetCurve(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
func (self class) SetGradient(color objects.Gradient) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(color[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGradient() objects.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMode(mode classdb.Line2DLineTextureMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMode() classdb.Line2DLineTextureMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Line2DLineTextureMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJointMode(mode classdb.Line2DLineJointMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_joint_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetJointMode() classdb.Line2DLineJointMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Line2DLineJointMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_joint_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBeginCapMode(mode classdb.Line2DLineCapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_begin_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBeginCapMode() classdb.Line2DLineCapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Line2DLineCapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_get_begin_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEndCapMode(mode classdb.Line2DLineCapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Line2D.Bind_set_end_cap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEndCapMode() classdb.Line2DLineCapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Line2DLineCapMode](frame)
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
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() { classdb.Register("Line2D", func(ptr gd.Object) any { return classdb.Line2D(ptr) }) }

type LineJointMode = classdb.Line2DLineJointMode

const (
	/*Makes the polyline's joints pointy, connecting the sides of the two segments by extending them until they intersect. If the rotation of a joint is too big (based on [member sharp_limit]), the joint falls back to [constant LINE_JOINT_BEVEL] to prevent very long miters.*/
	LineJointSharp LineJointMode = 0
	/*Makes the polyline's joints bevelled/chamfered, connecting the sides of the two segments with a simple line.*/
	LineJointBevel LineJointMode = 1
	/*Makes the polyline's joints rounded, connecting the sides of the two segments with an arc. The detail of this arc depends on [member round_precision].*/
	LineJointRound LineJointMode = 2
)

type LineCapMode = classdb.Line2DLineCapMode

const (
	/*Draws no line cap.*/
	LineCapNone LineCapMode = 0
	/*Draws the line cap as a box, slightly extending the first/last segment.*/
	LineCapBox LineCapMode = 1
	/*Draws the line cap as a semicircle attached to the first/last segment.*/
	LineCapRound LineCapMode = 2
)

type LineTextureMode = classdb.Line2DLineTextureMode

const (
	/*Takes the left pixels of the texture and renders them over the whole polyline.*/
	LineTextureNone LineTextureMode = 0
	/*Tiles the texture over the polyline. [member CanvasItem.texture_repeat] of the [Line2D] node must be [constant CanvasItem.TEXTURE_REPEAT_ENABLED] or [constant CanvasItem.TEXTURE_REPEAT_MIRROR] for it to work properly.*/
	LineTextureTile LineTextureMode = 1
	/*Stretches the texture across the polyline. [member CanvasItem.texture_repeat] of the [Line2D] node must be [constant CanvasItem.TEXTURE_REPEAT_DISABLED] for best results.*/
	LineTextureStretch LineTextureMode = 2
)
