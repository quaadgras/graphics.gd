// Package Node2D provides methods for working with Node2D object instances.
package Node2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A 2D game object, with a transform (position, rotation, and scale). All 2D nodes, including physics objects and sprites, inherit from Node2D. Use Node2D as a parent node to move, scale and rotate children in a 2D project. Also gives control of the node's render order.
[b]Note:[/b] Since both [Node2D] and [Control] inherit from [CanvasItem], they share several concepts from the class such as the [member CanvasItem.z_index] and [member CanvasItem.visible] properties.
*/
type Instance [1]gdclass.Node2D
type Expanded [1]gdclass.Node2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNode2D() Instance
}

/*
Applies a rotation to the node, in radians, starting from its current rotation.
*/
func (self Instance) Rotate(radians Float.X) { //gd:Node2D.rotate
	Advanced(self).Rotate(float64(radians))
}

/*
Applies a local translation on the node's X axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
func (self Instance) MoveLocalX(delta Float.X) { //gd:Node2D.move_local_x
	Advanced(self).MoveLocalX(float64(delta), false)
}

/*
Applies a local translation on the node's X axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
func (self Expanded) MoveLocalX(delta Float.X, scaled bool) { //gd:Node2D.move_local_x
	Advanced(self).MoveLocalX(float64(delta), scaled)
}

/*
Applies a local translation on the node's Y axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
func (self Instance) MoveLocalY(delta Float.X) { //gd:Node2D.move_local_y
	Advanced(self).MoveLocalY(float64(delta), false)
}

/*
Applies a local translation on the node's Y axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
func (self Expanded) MoveLocalY(delta Float.X, scaled bool) { //gd:Node2D.move_local_y
	Advanced(self).MoveLocalY(float64(delta), scaled)
}

/*
Translates the node by the given [param offset] in local coordinates.
*/
func (self Instance) Translate(offset Vector2.XY) { //gd:Node2D.translate
	Advanced(self).Translate(Vector2.XY(offset))
}

/*
Adds the [param offset] vector to the node's global position.
*/
func (self Instance) GlobalTranslate(offset Vector2.XY) { //gd:Node2D.global_translate
	Advanced(self).GlobalTranslate(Vector2.XY(offset))
}

/*
Multiplies the current scale by the [param ratio] vector.
*/
func (self Instance) ApplyScale(ratio Vector2.XY) { //gd:Node2D.apply_scale
	Advanced(self).ApplyScale(Vector2.XY(ratio))
}

/*
Rotates the node so that its local +X axis points towards the [param point], which is expected to use global coordinates.
[param point] should not be the same as the node's position, otherwise the node always looks to the right.
*/
func (self Instance) LookAt(point Vector2.XY) { //gd:Node2D.look_at
	Advanced(self).LookAt(Vector2.XY(point))
}

/*
Returns the angle between the node and the [param point] in radians.
[url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/node2d_get_angle_to.png]Illustration of the returned angle.[/url]
*/
func (self Instance) GetAngleTo(point Vector2.XY) Float.X { //gd:Node2D.get_angle_to
	return Float.X(Float.X(Advanced(self).GetAngleTo(Vector2.XY(point))))
}

/*
Transforms the provided global position into a position in local coordinate space. The output will be local relative to the [Node2D] it is called on. e.g. It is appropriate for determining the positions of child nodes, but it is not appropriate for determining its own position relative to its parent.
*/
func (self Instance) ToLocal(global_point Vector2.XY) Vector2.XY { //gd:Node2D.to_local
	return Vector2.XY(Advanced(self).ToLocal(Vector2.XY(global_point)))
}

/*
Transforms the provided local position into a position in global coordinate space. The input is expected to be local relative to the [Node2D] it is called on. e.g. Applying this method to the positions of child nodes will correctly transform their positions into the global coordinate space, but applying it to a node's own position will give an incorrect result, as it will incorporate the node's own transformation into its global position.
*/
func (self Instance) ToGlobal(local_point Vector2.XY) Vector2.XY { //gd:Node2D.to_global
	return Vector2.XY(Advanced(self).ToGlobal(Vector2.XY(local_point)))
}

/*
Returns the [Transform2D] relative to this node's parent.
*/
func (self Instance) GetRelativeTransformToParent(parent [1]gdclass.Node) Transform2D.OriginXY { //gd:Node2D.get_relative_transform_to_parent
	return Transform2D.OriginXY(Advanced(self).GetRelativeTransformToParent(parent))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Node2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Node2D"))
	casted := Instance{*(*gdclass.Node2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Position() Vector2.XY {
	return Vector2.XY(class(self).GetPosition())
}

func (self Instance) SetPosition(value Vector2.XY) {
	class(self).SetPosition(Vector2.XY(value))
}

func (self Instance) Rotation() Float.X {
	return Float.X(Float.X(class(self).GetRotation()))
}

func (self Instance) SetRotation(value Float.X) {
	class(self).SetRotation(float64(value))
}

func (self Instance) RotationDegrees() Float.X {
	return Float.X(Float.X(class(self).GetRotationDegrees()))
}

func (self Instance) SetRotationDegrees(value Float.X) {
	class(self).SetRotationDegrees(float64(value))
}

func (self Instance) Scale() Vector2.XY {
	return Vector2.XY(class(self).GetScale())
}

func (self Instance) SetScale(value Vector2.XY) {
	class(self).SetScale(Vector2.XY(value))
}

func (self Instance) Skew() Float.X {
	return Float.X(Float.X(class(self).GetSkew()))
}

func (self Instance) SetSkew(value Float.X) {
	class(self).SetSkew(float64(value))
}

func (self Instance) SetTransform(value Transform2D.OriginXY) {
	class(self).SetTransform(Transform2D.OriginXY(value))
}

func (self Instance) GlobalPosition() Vector2.XY {
	return Vector2.XY(class(self).GetGlobalPosition())
}

func (self Instance) SetGlobalPosition(value Vector2.XY) {
	class(self).SetGlobalPosition(Vector2.XY(value))
}

func (self Instance) GlobalRotation() Float.X {
	return Float.X(Float.X(class(self).GetGlobalRotation()))
}

func (self Instance) SetGlobalRotation(value Float.X) {
	class(self).SetGlobalRotation(float64(value))
}

func (self Instance) GlobalRotationDegrees() Float.X {
	return Float.X(Float.X(class(self).GetGlobalRotationDegrees()))
}

func (self Instance) SetGlobalRotationDegrees(value Float.X) {
	class(self).SetGlobalRotationDegrees(float64(value))
}

func (self Instance) GlobalScale() Vector2.XY {
	return Vector2.XY(class(self).GetGlobalScale())
}

func (self Instance) SetGlobalScale(value Vector2.XY) {
	class(self).SetGlobalScale(Vector2.XY(value))
}

func (self Instance) GlobalSkew() Float.X {
	return Float.X(Float.X(class(self).GetGlobalSkew()))
}

func (self Instance) SetGlobalSkew(value Float.X) {
	class(self).SetGlobalSkew(float64(value))
}

func (self Instance) SetGlobalTransform(value Transform2D.OriginXY) {
	class(self).SetGlobalTransform(Transform2D.OriginXY(value))
}

//go:nosplit
func (self class) SetPosition(position Vector2.XY) { //gd:Node2D.set_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetRotation(radians float64) { //gd:Node2D.set_rotation
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetRotationDegrees(degrees float64) { //gd:Node2D.set_rotation_degrees
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSkew(radians float64) { //gd:Node2D.set_skew
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_skew, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetScale(scale Vector2.XY) { //gd:Node2D.set_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() Vector2.XY { //gd:Node2D.get_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRotation() float64 { //gd:Node2D.get_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetRotationDegrees() float64 { //gd:Node2D.get_rotation_degrees
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSkew() float64 { //gd:Node2D.get_skew
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_skew, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetScale() Vector2.XY { //gd:Node2D.get_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Applies a rotation to the node, in radians, starting from its current rotation.
*/
//go:nosplit
func (self class) Rotate(radians float64) { //gd:Node2D.rotate
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_rotate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a local translation on the node's X axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
//go:nosplit
func (self class) MoveLocalX(delta float64, scaled bool) { //gd:Node2D.move_local_x
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, scaled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_move_local_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Applies a local translation on the node's Y axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
//go:nosplit
func (self class) MoveLocalY(delta float64, scaled bool) { //gd:Node2D.move_local_y
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, scaled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_move_local_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Translates the node by the given [param offset] in local coordinates.
*/
//go:nosplit
func (self class) Translate(offset Vector2.XY) { //gd:Node2D.translate
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds the [param offset] vector to the node's global position.
*/
//go:nosplit
func (self class) GlobalTranslate(offset Vector2.XY) { //gd:Node2D.global_translate
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_global_translate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Multiplies the current scale by the [param ratio] vector.
*/
//go:nosplit
func (self class) ApplyScale(ratio Vector2.XY) { //gd:Node2D.apply_scale
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_apply_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetGlobalPosition(position Vector2.XY) { //gd:Node2D.set_global_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalPosition() Vector2.XY { //gd:Node2D.get_global_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalRotation(radians float64) { //gd:Node2D.set_global_rotation
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetGlobalRotationDegrees(degrees float64) { //gd:Node2D.set_global_rotation_degrees
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalRotation() float64 { //gd:Node2D.get_global_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetGlobalRotationDegrees() float64 { //gd:Node2D.get_global_rotation_degrees
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalSkew(radians float64) { //gd:Node2D.set_global_skew
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_skew, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalSkew() float64 { //gd:Node2D.get_global_skew
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_skew, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalScale(scale Vector2.XY) { //gd:Node2D.set_global_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalScale() Vector2.XY { //gd:Node2D.get_global_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(xform Transform2D.OriginXY) { //gd:Node2D.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetGlobalTransform(xform Transform2D.OriginXY) { //gd:Node2D.set_global_transform
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Rotates the node so that its local +X axis points towards the [param point], which is expected to use global coordinates.
[param point] should not be the same as the node's position, otherwise the node always looks to the right.
*/
//go:nosplit
func (self class) LookAt(point Vector2.XY) { //gd:Node2D.look_at
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_look_at, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the angle between the node and the [param point] in radians.
[url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/node2d_get_angle_to.png]Illustration of the returned angle.[/url]
*/
//go:nosplit
func (self class) GetAngleTo(point Vector2.XY) float64 { //gd:Node2D.get_angle_to
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_angle_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Transforms the provided global position into a position in local coordinate space. The output will be local relative to the [Node2D] it is called on. e.g. It is appropriate for determining the positions of child nodes, but it is not appropriate for determining its own position relative to its parent.
*/
//go:nosplit
func (self class) ToLocal(global_point Vector2.XY) Vector2.XY { //gd:Node2D.to_local
	var frame = callframe.New()
	callframe.Arg(frame, global_point)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_to_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Transforms the provided local position into a position in global coordinate space. The input is expected to be local relative to the [Node2D] it is called on. e.g. Applying this method to the positions of child nodes will correctly transform their positions into the global coordinate space, but applying it to a node's own position will give an incorrect result, as it will incorporate the node's own transformation into its global position.
*/
//go:nosplit
func (self class) ToGlobal(local_point Vector2.XY) Vector2.XY { //gd:Node2D.to_global
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_to_global, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Transform2D] relative to this node's parent.
*/
//go:nosplit
func (self class) GetRelativeTransformToParent(parent [1]gdclass.Node) Transform2D.OriginXY { //gd:Node2D.get_relative_transform_to_parent
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parent[0])[0])
	var r_ret = callframe.Ret[Transform2D.OriginXY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_relative_transform_to_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNode2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(CanvasItem.Advanced(self.AsCanvasItem()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(CanvasItem.Instance(self.AsCanvasItem()), name)
	}
}
func init() {
	gdclass.Register("Node2D", func(ptr gd.Object) any { return [1]gdclass.Node2D{*(*gdclass.Node2D)(unsafe.Pointer(&ptr))} })
}
