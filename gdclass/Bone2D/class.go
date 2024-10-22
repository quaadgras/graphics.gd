package Bone2D

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
A hierarchy of [Bone2D]s can be bound to a [Skeleton2D] to control and animate other [Node2D] nodes.
You can use [Bone2D] and [Skeleton2D] nodes to animate 2D meshes created with the [Polygon2D] UV editor.
Each bone has a [member rest] transform that you can reset to with [method apply_rest]. These rest poses are relative to the bone's parent.
If in the editor, you can set the rest pose of an entire skeleton using a menu option, from the code, you need to iterate over the bones to set their individual rest poses.

*/
type Go [1]classdb.Bone2D

/*
Resets the bone to the rest pose. This is equivalent to setting [member Node2D.transform] to [member rest].
*/
func (self Go) ApplyRest() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ApplyRest()
}

/*
Returns the node's [member rest] [Transform2D] if it doesn't have a parent, or its rest pose relative to its parent.
*/
func (self Go) GetSkeletonRest() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(class(self).GetSkeletonRest())
}

/*
Returns the node's index as part of the entire skeleton. See [Skeleton2D].
*/
func (self Go) GetIndexInSkeleton() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetIndexInSkeleton()))
}

/*
When set to [code]true[/code], the [Bone2D] node will attempt to automatically calculate the bone angle and length using the first child [Bone2D] node, if one exists. If none exist, the [Bone2D] cannot automatically calculate these values and will print a warning.
*/
func (self Go) SetAutocalculateLengthAndAngle(auto_calculate bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutocalculateLengthAndAngle(auto_calculate)
}

/*
Returns whether this [Bone2D] is going to autocalculate its length and bone angle using its first [Bone2D] child node, if one exists. If there are no [Bone2D] children, then it cannot autocalculate these values and will print a warning.
*/
func (self Go) GetAutocalculateLengthAndAngle() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetAutocalculateLengthAndAngle())
}

/*
Sets the length of the bone in the [Bone2D].
*/
func (self Go) SetLength(length float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLength(gd.Float(length))
}

/*
Returns the length of the bone in the [Bone2D] node.
*/
func (self Go) GetLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLength()))
}

/*
Sets the bone angle for the [Bone2D]. This is typically set to the rotation from the [Bone2D] to a child [Bone2D] node.
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
func (self Go) SetBoneAngle(angle float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBoneAngle(gd.Float(angle))
}

/*
Returns the angle of the bone in the [Bone2D].
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
func (self Go) GetBoneAngle() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetBoneAngle()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Bone2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Bone2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Rest() gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Transform2D(class(self).GetRest())
}

func (self Go) SetRest(value gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRest(value)
}

//go:nosplit
func (self class) SetRest(rest gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rest)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_set_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRest() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_get_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resets the bone to the rest pose. This is equivalent to setting [member Node2D.transform] to [member rest].
*/
//go:nosplit
func (self class) ApplyRest()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_apply_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the node's [member rest] [Transform2D] if it doesn't have a parent, or its rest pose relative to its parent.
*/
//go:nosplit
func (self class) GetSkeletonRest() gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_get_skeleton_rest, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the node's index as part of the entire skeleton. See [Skeleton2D].
*/
//go:nosplit
func (self class) GetIndexInSkeleton() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_get_index_in_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
When set to [code]true[/code], the [Bone2D] node will attempt to automatically calculate the bone angle and length using the first child [Bone2D] node, if one exists. If none exist, the [Bone2D] cannot automatically calculate these values and will print a warning.
*/
//go:nosplit
func (self class) SetAutocalculateLengthAndAngle(auto_calculate bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, auto_calculate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_set_autocalculate_length_and_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether this [Bone2D] is going to autocalculate its length and bone angle using its first [Bone2D] child node, if one exists. If there are no [Bone2D] children, then it cannot autocalculate these values and will print a warning.
*/
//go:nosplit
func (self class) GetAutocalculateLengthAndAngle() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_get_autocalculate_length_and_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the length of the bone in the [Bone2D].
*/
//go:nosplit
func (self class) SetLength(length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the length of the bone in the [Bone2D] node.
*/
//go:nosplit
func (self class) GetLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the bone angle for the [Bone2D]. This is typically set to the rotation from the [Bone2D] to a child [Bone2D] node.
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
//go:nosplit
func (self class) SetBoneAngle(angle gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_set_bone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the angle of the bone in the [Bone2D].
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
//go:nosplit
func (self class) GetBoneAngle() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Bone2D.Bind_get_bone_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBone2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBone2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Bone2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
