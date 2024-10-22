package GraphElement

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
[GraphElement] allows to create custom elements for a [GraphEdit] graph. By default such elements can be selected, resized, and repositioned, but they cannot be connected. For a graph element that allows for connections see [GraphNode].

*/
type Go [1]classdb.GraphElement
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GraphElement
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GraphElement"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PositionOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetPositionOffset())
}

func (self Go) SetPositionOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPositionOffset(value)
}

func (self Go) Resizable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsResizable())
}

func (self Go) SetResizable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetResizable(value)
}

func (self Go) Draggable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDraggable())
}

func (self Go) SetDraggable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDraggable(value)
}

func (self Go) Selectable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelectable())
}

func (self Go) SetSelectable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectable(value)
}

func (self Go) Selected() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelected())
}

func (self Go) SetSelected(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelected(value)
}

//go:nosplit
func (self class) SetResizable(resizable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, resizable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_set_resizable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsResizable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_is_resizable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDraggable(draggable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draggable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_set_draggable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDraggable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_is_draggable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectable(selectable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelectable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelected(selected bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, selected)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelected() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPositionOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_set_position_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPositionOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GraphElement.Bind_get_position_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnNodeSelected(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_selected"), gc.Callable(cb), 0)
}


func (self Go) OnNodeDeselected(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("node_deselected"), gc.Callable(cb), 0)
}


func (self Go) OnRaiseRequest(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("raise_request"), gc.Callable(cb), 0)
}


func (self Go) OnDeleteRequest(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("delete_request"), gc.Callable(cb), 0)
}


func (self Go) OnResizeRequest(cb func(new_size gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("resize_request"), gc.Callable(cb), 0)
}


func (self Go) OnResizeEnd(cb func(new_size gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("resize_end"), gc.Callable(cb), 0)
}


func (self Go) OnDragged(cb func(from gd.Vector2, to gd.Vector2)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("dragged"), gc.Callable(cb), 0)
}


func (self Go) OnPositionOffsetChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("position_offset_changed"), gc.Callable(cb), 0)
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
func init() {classdb.Register("GraphElement", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
