package GraphElement

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[GraphElement] allows to create custom elements for a [GraphEdit] graph. By default such elements can be selected, resized, and repositioned, but they cannot be connected. For a graph element that allows for connections see [GraphNode].

*/
type Go [1]classdb.GraphElement
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GraphElement
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphElement"))
	return Go{classdb.GraphElement(object)}
}

func (self Go) PositionOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetPositionOffset())
}

func (self Go) SetPositionOffset(value gd.Vector2) {
	class(self).SetPositionOffset(value)
}

func (self Go) Resizable() bool {
		return bool(class(self).IsResizable())
}

func (self Go) SetResizable(value bool) {
	class(self).SetResizable(value)
}

func (self Go) Draggable() bool {
		return bool(class(self).IsDraggable())
}

func (self Go) SetDraggable(value bool) {
	class(self).SetDraggable(value)
}

func (self Go) Selectable() bool {
		return bool(class(self).IsSelectable())
}

func (self Go) SetSelectable(value bool) {
	class(self).SetSelectable(value)
}

func (self Go) Selected() bool {
		return bool(class(self).IsSelected())
}

func (self Go) SetSelected(value bool) {
	class(self).SetSelected(value)
}

//go:nosplit
func (self class) SetResizable(resizable bool)  {
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
func (self class) SetDraggable(draggable bool)  {
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
func (self class) SetSelectable(selectable bool)  {
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
func (self class) SetSelected(selected bool)  {
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
func (self class) SetPositionOffset(offset gd.Vector2)  {
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
func (self Go) OnNodeSelected(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("node_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnNodeDeselected(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("node_deselected"), gd.NewCallable(cb), 0)
}


func (self Go) OnRaiseRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("raise_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnDeleteRequest(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("delete_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnResizeRequest(cb func(new_size gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("resize_request"), gd.NewCallable(cb), 0)
}


func (self Go) OnResizeEnd(cb func(new_size gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("resize_end"), gd.NewCallable(cb), 0)
}


func (self Go) OnDragged(cb func(from gd.Vector2, to gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("dragged"), gd.NewCallable(cb), 0)
}


func (self Go) OnPositionOffsetChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("position_offset_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsGraphElement() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGraphElement() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsContainer(), name)
	}
}
func init() {classdb.Register("GraphElement", func(ptr gd.Object) any { return classdb.GraphElement(ptr) })}
