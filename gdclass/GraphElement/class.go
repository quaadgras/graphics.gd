package GraphElement

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[GraphElement] allows to create custom elements for a [GraphEdit] graph. By default such elements can be selected, resized, and repositioned, but they cannot be connected. For a graph element that allows for connections see [GraphNode].
*/
type Instance [1]classdb.GraphElement

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GraphElement

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphElement"))
	return Instance{classdb.GraphElement(object)}
}

func (self Instance) PositionOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetPositionOffset())
}

func (self Instance) SetPositionOffset(value gd.Vector2) {
	class(self).SetPositionOffset(value)
}

func (self Instance) Resizable() bool {
	return bool(class(self).IsResizable())
}

func (self Instance) SetResizable(value bool) {
	class(self).SetResizable(value)
}

func (self Instance) Draggable() bool {
	return bool(class(self).IsDraggable())
}

func (self Instance) SetDraggable(value bool) {
	class(self).SetDraggable(value)
}

func (self Instance) Selectable() bool {
	return bool(class(self).IsSelectable())
}

func (self Instance) SetSelectable(value bool) {
	class(self).SetSelectable(value)
}

func (self Instance) Selected() bool {
	return bool(class(self).IsSelected())
}

func (self Instance) SetSelected(value bool) {
	class(self).SetSelected(value)
}

//go:nosplit
func (self class) SetResizable(resizable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, resizable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_resizable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsResizable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_resizable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDraggable(draggable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, draggable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_draggable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDraggable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_draggable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectable(selectable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelected(selected bool) {
	var frame = callframe.New()
	callframe.Arg(frame, selected)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelected() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPositionOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_position_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPositionOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_get_position_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnNodeSelected(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("node_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeDeselected(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("node_deselected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRaiseRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("raise_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDeleteRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("delete_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResizeRequest(cb func(new_size gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("resize_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResizeEnd(cb func(new_size gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("resize_end"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDragged(cb func(from gd.Vector2, to gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("dragged"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPositionOffsetChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("position_offset_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGraphElement() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGraphElement() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {
	classdb.Register("GraphElement", func(ptr gd.Object) any { return classdb.GraphElement(ptr) })
}
