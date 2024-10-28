package Node2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A 2D game object, with a transform (position, rotation, and scale). All 2D nodes, including physics objects and sprites, inherit from Node2D. Use Node2D as a parent node to move, scale and rotate children in a 2D project. Also gives control of the node's render order.

*/
type Go [1]classdb.Node2D

/*
Applies a rotation to the node, in radians, starting from its current rotation.
*/
func (self Go) Rotate(radians float64) {
	class(self).Rotate(gd.Float(radians))
}

/*
Applies a local translation on the node's X axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
func (self Go) MoveLocalX(delta float64) {
	class(self).MoveLocalX(gd.Float(delta), false)
}

/*
Applies a local translation on the node's Y axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
func (self Go) MoveLocalY(delta float64) {
	class(self).MoveLocalY(gd.Float(delta), false)
}

/*
Translates the node by the given [param offset] in local coordinates.
*/
func (self Go) Translate(offset gd.Vector2) {
	class(self).Translate(offset)
}

/*
Adds the [param offset] vector to the node's global position.
*/
func (self Go) GlobalTranslate(offset gd.Vector2) {
	class(self).GlobalTranslate(offset)
}

/*
Multiplies the current scale by the [param ratio] vector.
*/
func (self Go) ApplyScale(ratio gd.Vector2) {
	class(self).ApplyScale(ratio)
}

/*
Rotates the node so that its local +X axis points towards the [param point], which is expected to use global coordinates.
[param point] should not be the same as the node's position, otherwise the node always looks to the right.
*/
func (self Go) LookAt(point gd.Vector2) {
	class(self).LookAt(point)
}

/*
Returns the angle between the node and the [param point] in radians.
[url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/node2d_get_angle_to.png]Illustration of the returned angle.[/url]
*/
func (self Go) GetAngleTo(point gd.Vector2) float64 {
	return float64(float64(class(self).GetAngleTo(point)))
}

/*
Transforms the provided global position into a position in local coordinate space. The output will be local relative to the [Node2D] it is called on. e.g. It is appropriate for determining the positions of child nodes, but it is not appropriate for determining its own position relative to its parent.
*/
func (self Go) ToLocal(global_point gd.Vector2) gd.Vector2 {
	return gd.Vector2(class(self).ToLocal(global_point))
}

/*
Transforms the provided local position into a position in global coordinate space. The input is expected to be local relative to the [Node2D] it is called on. e.g. Applying this method to the positions of child nodes will correctly transform their positions into the global coordinate space, but applying it to a node's own position will give an incorrect result, as it will incorporate the node's own transformation into its global position.
*/
func (self Go) ToGlobal(local_point gd.Vector2) gd.Vector2 {
	return gd.Vector2(class(self).ToGlobal(local_point))
}

/*
Returns the [Transform2D] relative to this node's parent.
*/
func (self Go) GetRelativeTransformToParent(parent gdclass.Node) gd.Transform2D {
	return gd.Transform2D(class(self).GetRelativeTransformToParent(parent))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Node2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Node2D"))
	return Go{classdb.Node2D(object)}
}

func (self Go) Position() gd.Vector2 {
		return gd.Vector2(class(self).GetPosition())
}

func (self Go) SetPosition(value gd.Vector2) {
	class(self).SetPosition(value)
}

func (self Go) Rotation() float64 {
		return float64(float64(class(self).GetRotation()))
}

func (self Go) SetRotation(value float64) {
	class(self).SetRotation(gd.Float(value))
}

func (self Go) RotationDegrees() float64 {
		return float64(float64(class(self).GetRotationDegrees()))
}

func (self Go) SetRotationDegrees(value float64) {
	class(self).SetRotationDegrees(gd.Float(value))
}

func (self Go) Scale() gd.Vector2 {
		return gd.Vector2(class(self).GetScale())
}

func (self Go) SetScale(value gd.Vector2) {
	class(self).SetScale(value)
}

func (self Go) Skew() float64 {
		return float64(float64(class(self).GetSkew()))
}

func (self Go) SetSkew(value float64) {
	class(self).SetSkew(gd.Float(value))
}

func (self Go) GlobalPosition() gd.Vector2 {
		return gd.Vector2(class(self).GetGlobalPosition())
}

func (self Go) SetGlobalPosition(value gd.Vector2) {
	class(self).SetGlobalPosition(value)
}

func (self Go) GlobalRotation() float64 {
		return float64(float64(class(self).GetGlobalRotation()))
}

func (self Go) SetGlobalRotation(value float64) {
	class(self).SetGlobalRotation(gd.Float(value))
}

func (self Go) GlobalRotationDegrees() float64 {
		return float64(float64(class(self).GetGlobalRotationDegrees()))
}

func (self Go) SetGlobalRotationDegrees(value float64) {
	class(self).SetGlobalRotationDegrees(gd.Float(value))
}

func (self Go) GlobalScale() gd.Vector2 {
		return gd.Vector2(class(self).GetGlobalScale())
}

func (self Go) SetGlobalScale(value gd.Vector2) {
	class(self).SetGlobalScale(value)
}

func (self Go) GlobalSkew() float64 {
		return float64(float64(class(self).GetGlobalSkew()))
}

func (self Go) SetGlobalSkew(value float64) {
	class(self).SetGlobalSkew(gd.Float(value))
}

//go:nosplit
func (self class) SetPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetRotation(radians gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetRotationDegrees(degrees gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSkew(radians gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetScale(scale gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetRotation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetRotationDegrees() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSkew() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetScale() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Applies a rotation to the node, in radians, starting from its current rotation.
*/
//go:nosplit
func (self class) Rotate(radians gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_rotate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a local translation on the node's X axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
//go:nosplit
func (self class) MoveLocalX(delta gd.Float, scaled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, scaled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_move_local_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a local translation on the node's Y axis based on the [method Node._process]'s [param delta]. If [param scaled] is [code]false[/code], normalizes the movement.
*/
//go:nosplit
func (self class) MoveLocalY(delta gd.Float, scaled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, scaled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_move_local_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Translates the node by the given [param offset] in local coordinates.
*/
//go:nosplit
func (self class) Translate(offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the [param offset] vector to the node's global position.
*/
//go:nosplit
func (self class) GlobalTranslate(offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_global_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Multiplies the current scale by the [param ratio] vector.
*/
//go:nosplit
func (self class) ApplyScale(ratio gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_apply_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetGlobalPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalRotation(radians gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetGlobalRotationDegrees(degrees gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalRotation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetGlobalRotationDegrees() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalSkew(radians gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalSkew() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlobalScale(scale gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlobalScale() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_global_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransform(xform gd.Transform2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetGlobalTransform(xform gd.Transform2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_set_global_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Rotates the node so that its local +X axis points towards the [param point], which is expected to use global coordinates.
[param point] should not be the same as the node's position, otherwise the node always looks to the right.
*/
//go:nosplit
func (self class) LookAt(point gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_look_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the angle between the node and the [param point] in radians.
[url=https://raw.githubusercontent.com/godotengine/godot-docs/master/img/node2d_get_angle_to.png]Illustration of the returned angle.[/url]
*/
//go:nosplit
func (self class) GetAngleTo(point gd.Vector2) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_angle_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Transforms the provided global position into a position in local coordinate space. The output will be local relative to the [Node2D] it is called on. e.g. It is appropriate for determining the positions of child nodes, but it is not appropriate for determining its own position relative to its parent.
*/
//go:nosplit
func (self class) ToLocal(global_point gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, global_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Transforms the provided local position into a position in global coordinate space. The input is expected to be local relative to the [Node2D] it is called on. e.g. Applying this method to the positions of child nodes will correctly transform their positions into the global coordinate space, but applying it to a node's own position will give an incorrect result, as it will incorporate the node's own transformation into its global position.
*/
//go:nosplit
func (self class) ToGlobal(local_point gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_to_global, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Transform2D] relative to this node's parent.
*/
//go:nosplit
func (self class) GetRelativeTransformToParent(parent gdclass.Node) gd.Transform2D {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(parent[0])[0])
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Node2D.Bind_get_relative_transform_to_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNode2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCanvasItem(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCanvasItem(), name)
	}
}
func init() {classdb.Register("Node2D", func(ptr gd.Object) any { return classdb.Node2D(ptr) })}
