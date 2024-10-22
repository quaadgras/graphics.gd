package EditorSelection

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object manages the SceneTree selection in the editor.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_selection].

*/
type Go [1]classdb.EditorSelection

/*
Clear the selection.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Adds a node to the selection.
[b]Note:[/b] The newly selected node will not be automatically edited in the inspector. If you want to edit a node, use [method EditorInterface.edit_node].
*/
func (self Go) AddNode(node gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddNode(node)
}

/*
Removes a node from the selection.
*/
func (self Go) RemoveNode(node gdclass.Node) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveNode(node)
}

/*
Returns the list of selected nodes.
*/
func (self Go) GetSelectedNodes() gd.ArrayOf[gdclass.Node] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Node](class(self).GetSelectedNodes(gc))
}

/*
Returns the list of selected nodes, optimized for transform operations (i.e. moving them, rotating, etc.). This list can be used to avoid situations where a node is selected and is also a child/grandchild.
*/
func (self Go) GetTransformableSelectedNodes() gd.ArrayOf[gdclass.Node] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Node](class(self).GetTransformableSelectedNodes(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSelection
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorSelection"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Clear the selection.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSelection.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a node to the selection.
[b]Note:[/b] The newly selected node will not be automatically edited in the inspector. If you want to edit a node, use [method EditorInterface.edit_node].
*/
//go:nosplit
func (self class) AddNode(node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSelection.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a node from the selection.
*/
//go:nosplit
func (self class) RemoveNode(node gdclass.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSelection.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of selected nodes.
*/
//go:nosplit
func (self class) GetSelectedNodes(ctx gd.Lifetime) gd.ArrayOf[gdclass.Node] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSelection.Bind_get_selected_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Node](ret)
}
/*
Returns the list of selected nodes, optimized for transform operations (i.e. moving them, rotating, etc.). This list can be used to avoid situations where a node is selected and is also a child/grandchild.
*/
//go:nosplit
func (self class) GetTransformableSelectedNodes(ctx gd.Lifetime) gd.ArrayOf[gdclass.Node] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSelection.Bind_get_transformable_selected_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Node](ret)
}
func (self Go) OnSelectionChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("selection_changed"), gc.Callable(cb), 0)
}


func (self class) AsEditorSelection() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSelection() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("EditorSelection", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
