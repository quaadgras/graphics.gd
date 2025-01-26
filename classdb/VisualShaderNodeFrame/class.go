// Package VisualShaderNodeFrame provides methods for working with VisualShaderNodeFrame object instances.
package VisualShaderNodeFrame

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
import "graphics.gd/classdb/VisualShaderNodeResizableBase"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"

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
A rectangular frame that can be used to group visual shader nodes together to improve organization.
Nodes attached to the frame will move with it when it is dragged and it can automatically resize to enclose all attached nodes.
Its title, description and color can be customized.
*/
type Instance [1]gdclass.VisualShaderNodeFrame

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeFrame() Instance
}

/*
Adds a node to the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.attach_node_to_frame] method instead.
*/
func (self Instance) AddAttachedNode(node int) { //gd:VisualShaderNodeFrame.add_attached_node
	class(self).AddAttachedNode(gd.Int(node))
}

/*
Removes a node from the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.detach_node_from_frame] method instead.
*/
func (self Instance) RemoveAttachedNode(node int) { //gd:VisualShaderNodeFrame.remove_attached_node
	class(self).RemoveAttachedNode(gd.Int(node))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeFrame

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeFrame"))
	casted := Instance{*(*gdclass.VisualShaderNodeFrame)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Title() string {
	return string(class(self).GetTitle().String())
}

func (self Instance) SetTitle(value string) {
	class(self).SetTitle(gd.NewString(value))
}

func (self Instance) TintColorEnabled() bool {
	return bool(class(self).IsTintColorEnabled())
}

func (self Instance) SetTintColorEnabled(value bool) {
	class(self).SetTintColorEnabled(value)
}

func (self Instance) TintColor() Color.RGBA {
	return Color.RGBA(class(self).GetTintColor())
}

func (self Instance) SetTintColor(value Color.RGBA) {
	class(self).SetTintColor(gd.Color(value))
}

func (self Instance) Autoshrink() bool {
	return bool(class(self).IsAutoshrinkEnabled())
}

func (self Instance) SetAutoshrink(value bool) {
	class(self).SetAutoshrinkEnabled(value)
}

func (self Instance) AttachedNodes() []int32 {
	return []int32(class(self).GetAttachedNodes().AsSlice())
}

func (self Instance) SetAttachedNodes(value []int32) {
	class(self).SetAttachedNodes(gd.NewPackedInt32Slice(value))
}

//go:nosplit
func (self class) SetTitle(title gd.String) { //gd:VisualShaderNodeFrame.set_title
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_set_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTitle() gd.String { //gd:VisualShaderNodeFrame.get_title
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_get_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTintColorEnabled(enable bool) { //gd:VisualShaderNodeFrame.set_tint_color_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_set_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTintColorEnabled() bool { //gd:VisualShaderNodeFrame.is_tint_color_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_is_tint_color_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTintColor(color gd.Color) { //gd:VisualShaderNodeFrame.set_tint_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_set_tint_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTintColor() gd.Color { //gd:VisualShaderNodeFrame.get_tint_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_get_tint_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoshrinkEnabled(enable bool) { //gd:VisualShaderNodeFrame.set_autoshrink_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_set_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoshrinkEnabled() bool { //gd:VisualShaderNodeFrame.is_autoshrink_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_is_autoshrink_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a node to the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.attach_node_to_frame] method instead.
*/
//go:nosplit
func (self class) AddAttachedNode(node gd.Int) { //gd:VisualShaderNodeFrame.add_attached_node
	var frame = callframe.New()
	callframe.Arg(frame, node)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_add_attached_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a node from the list of nodes attached to the frame. Should not be called directly, use the [method VisualShader.detach_node_from_frame] method instead.
*/
//go:nosplit
func (self class) RemoveAttachedNode(node gd.Int) { //gd:VisualShaderNodeFrame.remove_attached_node
	var frame = callframe.New()
	callframe.Arg(frame, node)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_remove_attached_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAttachedNodes(attached_nodes gd.PackedInt32Array) { //gd:VisualShaderNodeFrame.set_attached_nodes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(attached_nodes))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_set_attached_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttachedNodes() gd.PackedInt32Array { //gd:VisualShaderNodeFrame.get_attached_nodes
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFrame.Bind_get_attached_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeFrame() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeFrame() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Advanced {
	return *((*VisualShaderNodeResizableBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Instance {
	return *((*VisualShaderNodeResizableBase.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeResizableBase.Advanced(self.AsVisualShaderNodeResizableBase()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeResizableBase.Instance(self.AsVisualShaderNodeResizableBase()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeFrame", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeFrame{*(*gdclass.VisualShaderNodeFrame)(unsafe.Pointer(&ptr))}
	})
}
