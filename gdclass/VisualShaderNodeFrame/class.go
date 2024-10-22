package VisualShaderNodeFrame

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeResizableBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A rectangular frame that can be used to group visual shader nodes together to improve organization.
Nodes attached to the frame will move with it when it is dragged and it can automatically resize to enclose all attached nodes.
Its title, description and color can be customized.

*/
type Go [1]classdb.VisualShaderNodeFrame

/*
Adds a node to the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.attach_node_to_frame] method instead.
*/
func (self Go) AddAttachedNode(node int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddAttachedNode(gd.Int(node))
}

/*
Removes a node from the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.detach_node_from_frame] method instead.
*/
func (self Go) RemoveAttachedNode(node int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveAttachedNode(gd.Int(node))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeFrame
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VisualShaderNodeFrame"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Title() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetTitle(gc).String())
}

func (self Go) SetTitle(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTitle(gc.String(value))
}

func (self Go) TintColorEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsTintColorEnabled())
}

func (self Go) SetTintColorEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTintColorEnabled(value)
}

func (self Go) TintColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetTintColor())
}

func (self Go) SetTintColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTintColor(value)
}

func (self Go) Autoshrink() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoshrinkEnabled())
}

func (self Go) SetAutoshrink(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoshrinkEnabled(value)
}

func (self Go) AttachedNodes() []int32 {
	gc := gd.GarbageCollector(); _ = gc
		return []int32(class(self).GetAttachedNodes(gc).AsSlice())
}

func (self Go) SetAttachedNodes(value []int32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAttachedNodes(gc.PackedInt32Slice(value))
}

//go:nosplit
func (self class) SetTitle(title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTitle(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintColorEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_set_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTintColorEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_is_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTintColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_set_tint_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTintColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_get_tint_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoshrinkEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_set_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoshrinkEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_is_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a node to the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.attach_node_to_frame] method instead.
*/
//go:nosplit
func (self class) AddAttachedNode(node gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, node)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_add_attached_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a node from the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.detach_node_from_frame] method instead.
*/
//go:nosplit
func (self class) RemoveAttachedNode(node gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, node)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_remove_attached_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAttachedNodes(attached_nodes gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(attached_nodes))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_set_attached_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAttachedNodes(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeFrame.Bind_get_attached_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeFrame() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeFrame() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.GD { return *((*VisualShaderNodeResizableBase.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Go { return *((*VisualShaderNodeResizableBase.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeResizableBase(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeResizableBase(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeFrame", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
