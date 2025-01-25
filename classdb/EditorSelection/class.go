// Package EditorSelection provides methods for working with EditorSelection object instances.
package EditorSelection

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This object manages the SceneTree selection in the editor.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_selection].
*/
type Instance [1]gdclass.EditorSelection

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSelection() Instance
}

/*
Clear the selection.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Adds a node to the selection.
[b]Note:[/b] The newly selected node will not be automatically edited in the inspector. If you want to edit a node, use [method EditorInterface.edit_node].
*/
func (self Instance) AddNode(node [1]gdclass.Node) {
	class(self).AddNode(node)
}

/*
Removes a node from the selection.
*/
func (self Instance) RemoveNode(node [1]gdclass.Node) {
	class(self).RemoveNode(node)
}

/*
Returns the list of selected nodes.
*/
func (self Instance) GetSelectedNodes() [][1]gdclass.Node {
	return [][1]gdclass.Node(gd.ArrayAs[[][1]gdclass.Node](class(self).GetSelectedNodes()))
}

/*
Returns the list of selected nodes, optimized for transform operations (i.e. moving them, rotating, etc.). This list can be used to avoid situations where a node is selected and is also a child/grandchild.
*/
func (self Instance) GetTransformableSelectedNodes() [][1]gdclass.Node {
	return [][1]gdclass.Node(gd.ArrayAs[[][1]gdclass.Node](class(self).GetTransformableSelectedNodes()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSelection

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSelection"))
	casted := Instance{*(*gdclass.EditorSelection)(unsafe.Pointer(&object))}
	return casted
}

/*
Clear the selection.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSelection.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a node to the selection.
[b]Note:[/b] The newly selected node will not be automatically edited in the inspector. If you want to edit a node, use [method EditorInterface.edit_node].
*/
//go:nosplit
func (self class) AddNode(node [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(node[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSelection.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a node from the selection.
*/
//go:nosplit
func (self class) RemoveNode(node [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSelection.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the list of selected nodes.
*/
//go:nosplit
func (self class) GetSelectedNodes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSelection.Bind_get_selected_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of selected nodes, optimized for transform operations (i.e. moving them, rotating, etc.). This list can be used to avoid situations where a node is selected and is also a child/grandchild.
*/
//go:nosplit
func (self class) GetTransformableSelectedNodes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorSelection.Bind_get_transformable_selected_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnSelectionChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("selection_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorSelection() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorSelection() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("EditorSelection", func(ptr gd.Object) any {
		return [1]gdclass.EditorSelection{*(*gdclass.EditorSelection)(unsafe.Pointer(&ptr))}
	})
}
