package AnimationNodeBlendTree

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationRootNode"
import "grow.graphics/gd/gdclass/AnimationNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This animation node may contain a sub-tree of any other type animation nodes, such as [AnimationNodeTransition], [AnimationNodeBlend2], [AnimationNodeBlend3], [AnimationNodeOneShot], etc. This is one of the most commonly used animation node roots.
An [AnimationNodeOutput] node named [code]output[/code] is created by default.

*/
type Go [1]classdb.AnimationNodeBlendTree

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
func (self Go) AddNode(name string, node gdclass.AnimationNode) {
	class(self).AddNode(gd.NewStringName(name), node, gd.Vector2{0, 0})
}

/*
Returns the sub animation node with the specified [param name].
*/
func (self Go) GetNode(name string) gdclass.AnimationNode {
	return gdclass.AnimationNode(class(self).GetNode(gd.NewStringName(name)))
}

/*
Removes a sub animation node.
*/
func (self Go) RemoveNode(name string) {
	class(self).RemoveNode(gd.NewStringName(name))
}

/*
Changes the name of a sub animation node.
*/
func (self Go) RenameNode(name string, new_name string) {
	class(self).RenameNode(gd.NewStringName(name), gd.NewStringName(new_name))
}

/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
func (self Go) HasNode(name string) bool {
	return bool(class(self).HasNode(gd.NewStringName(name)))
}

/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
func (self Go) ConnectNode(input_node string, input_index int, output_node string) {
	class(self).ConnectNode(gd.NewStringName(input_node), gd.Int(input_index), gd.NewStringName(output_node))
}

/*
Disconnects the animation node connected to the specified input.
*/
func (self Go) DisconnectNode(input_node string, input_index int) {
	class(self).DisconnectNode(gd.NewStringName(input_node), gd.Int(input_index))
}

/*
Modifies the position of a sub animation node.
*/
func (self Go) SetNodePosition(name string, position gd.Vector2) {
	class(self).SetNodePosition(gd.NewStringName(name), position)
}

/*
Returns the position of the sub animation node with the specified [param name].
*/
func (self Go) GetNodePosition(name string) gd.Vector2 {
	return gd.Vector2(class(self).GetNodePosition(gd.NewStringName(name)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeBlendTree
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlendTree"))
	return Go{classdb.AnimationNodeBlendTree(object)}
}

func (self Go) GraphOffset() gd.Vector2 {
		return gd.Vector2(class(self).GetGraphOffset())
}

func (self Go) SetGraphOffset(value gd.Vector2) {
	class(self).SetGraphOffset(value)
}

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
//go:nosplit
func (self class) AddNode(name gd.StringName, node gdclass.AnimationNode, position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(node[0])[0])
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNode(name gd.StringName) gdclass.AnimationNode {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.AnimationNode{classdb.AnimationNode(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Removes a sub animation node.
*/
//go:nosplit
func (self class) RemoveNode(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the name of a sub animation node.
*/
//go:nosplit
func (self class) RenameNode(name gd.StringName, new_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(new_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_rename_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
//go:nosplit
func (self class) HasNode(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
//go:nosplit
func (self class) ConnectNode(input_node gd.StringName, input_index gd.Int, output_node gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(input_node))
	callframe.Arg(frame, input_index)
	callframe.Arg(frame, discreet.Get(output_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disconnects the animation node connected to the specified input.
*/
//go:nosplit
func (self class) DisconnectNode(input_node gd.StringName, input_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(input_node))
	callframe.Arg(frame, input_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Modifies the position of a sub animation node.
*/
//go:nosplit
func (self class) SetNodePosition(name gd.StringName, position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNodePosition(name gd.StringName) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnNodeChanged(cb func(node_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("node_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsAnimationNodeBlendTree() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeBlendTree() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationRootNode() AnimationRootNode.GD { return *((*AnimationRootNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationRootNode() AnimationRootNode.Go { return *((*AnimationRootNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNode() AnimationNode.GD { return *((*AnimationNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNode() AnimationNode.Go { return *((*AnimationNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}
func init() {classdb.Register("AnimationNodeBlendTree", func(ptr gd.Object) any { return classdb.AnimationNodeBlendTree(ptr) })}
