// Package AnimationNodeBlendTree provides methods for working with AnimationNodeBlendTree object instances.
package AnimationNodeBlendTree

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AnimationRootNode"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This animation node may contain a sub-tree of any other type animation nodes, such as [AnimationNodeTransition], [AnimationNodeBlend2], [AnimationNodeBlend3], [AnimationNodeOneShot], etc. This is one of the most commonly used animation node roots.
An [AnimationNodeOutput] node named [code]output[/code] is created by default.
*/
type Instance [1]gdclass.AnimationNodeBlendTree

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeBlendTree() Instance
}

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
func (self Instance) AddNode(name string, node [1]gdclass.AnimationNode) {
	class(self).AddNode(gd.NewStringName(name), node, gd.Vector2(gd.Vector2{0, 0}))
}

/*
Returns the sub animation node with the specified [param name].
*/
func (self Instance) GetNode(name string) [1]gdclass.AnimationNode {
	return [1]gdclass.AnimationNode(class(self).GetNode(gd.NewStringName(name)))
}

/*
Removes a sub animation node.
*/
func (self Instance) RemoveNode(name string) {
	class(self).RemoveNode(gd.NewStringName(name))
}

/*
Changes the name of a sub animation node.
*/
func (self Instance) RenameNode(name string, new_name string) {
	class(self).RenameNode(gd.NewStringName(name), gd.NewStringName(new_name))
}

/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
func (self Instance) HasNode(name string) bool {
	return bool(class(self).HasNode(gd.NewStringName(name)))
}

/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
func (self Instance) ConnectNode(input_node string, input_index int, output_node string) {
	class(self).ConnectNode(gd.NewStringName(input_node), gd.Int(input_index), gd.NewStringName(output_node))
}

/*
Disconnects the animation node connected to the specified input.
*/
func (self Instance) DisconnectNode(input_node string, input_index int) {
	class(self).DisconnectNode(gd.NewStringName(input_node), gd.Int(input_index))
}

/*
Modifies the position of a sub animation node.
*/
func (self Instance) SetNodePosition(name string, position Vector2.XY) {
	class(self).SetNodePosition(gd.NewStringName(name), gd.Vector2(position))
}

/*
Returns the position of the sub animation node with the specified [param name].
*/
func (self Instance) GetNodePosition(name string) Vector2.XY {
	return Vector2.XY(class(self).GetNodePosition(gd.NewStringName(name)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeBlendTree

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlendTree"))
	casted := Instance{*(*gdclass.AnimationNodeBlendTree)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) GraphOffset() Vector2.XY {
	return Vector2.XY(class(self).GetGraphOffset())
}

func (self Instance) SetGraphOffset(value Vector2.XY) {
	class(self).SetGraphOffset(gd.Vector2(value))
}

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
//go:nosplit
func (self class) AddNode(name gd.StringName, node [1]gdclass.AnimationNode, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNode(name gd.StringName) [1]gdclass.AnimationNode {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AnimationNode{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationNode](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes a sub animation node.
*/
//go:nosplit
func (self class) RemoveNode(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the name of a sub animation node.
*/
//go:nosplit
func (self class) RenameNode(name gd.StringName, new_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(new_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_rename_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
//go:nosplit
func (self class) HasNode(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
//go:nosplit
func (self class) ConnectNode(input_node gd.StringName, input_index gd.Int, output_node gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(input_node))
	callframe.Arg(frame, input_index)
	callframe.Arg(frame, pointers.Get(output_node))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disconnects the animation node connected to the specified input.
*/
//go:nosplit
func (self class) DisconnectNode(input_node gd.StringName, input_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(input_node))
	callframe.Arg(frame, input_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Modifies the position of a sub animation node.
*/
//go:nosplit
func (self class) SetNodePosition(name gd.StringName, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNodePosition(name gd.StringName) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnNodeChanged(cb func(node_name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationNodeBlendTree() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeBlendTree() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationRootNode() AnimationRootNode.Advanced {
	return *((*AnimationRootNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationRootNode() AnimationRootNode.Instance {
	return *((*AnimationRootNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AnimationRootNode.Advanced(self.AsAnimationRootNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationRootNode.Instance(self.AsAnimationRootNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeBlendTree", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeBlendTree{*(*gdclass.AnimationNodeBlendTree)(unsafe.Pointer(&ptr))}
	})
}
