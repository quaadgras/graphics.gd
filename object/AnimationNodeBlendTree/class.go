package AnimationNodeBlendTree

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationRootNode"
import "grow.graphics/gd/object/AnimationNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This animation node may contain a sub-tree of any other type animation nodes, such as [AnimationNodeTransition], [AnimationNodeBlend2], [AnimationNodeBlend3], [AnimationNodeOneShot], etc. This is one of the most commonly used animation node roots.
An [AnimationNodeOutput] node named [code]output[/code] is created by default.

*/
type Simple [1]classdb.AnimationNodeBlendTree
func (self Simple) AddNode(name string, node [1]classdb.AnimationNode, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddNode(gc.StringName(name), node, position)
}
func (self Simple) GetNode(name string) [1]classdb.AnimationNode {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AnimationNode(Expert(self).GetNode(gc, gc.StringName(name)))
}
func (self Simple) RemoveNode(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveNode(gc.StringName(name))
}
func (self Simple) RenameNode(name string, new_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RenameNode(gc.StringName(name), gc.StringName(new_name))
}
func (self Simple) HasNode(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasNode(gc.StringName(name)))
}
func (self Simple) ConnectNode(input_node string, input_index int, output_node string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ConnectNode(gc.StringName(input_node), gd.Int(input_index), gc.StringName(output_node))
}
func (self Simple) DisconnectNode(input_node string, input_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DisconnectNode(gc.StringName(input_node), gd.Int(input_index))
}
func (self Simple) SetNodePosition(name string, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNodePosition(gc.StringName(name), position)
}
func (self Simple) GetNodePosition(name string) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetNodePosition(gc.StringName(name)))
}
func (self Simple) SetGraphOffset(offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGraphOffset(offset)
}
func (self Simple) GetGraphOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGraphOffset())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeBlendTree
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
//go:nosplit
func (self class) AddNode(name gd.StringName, node object.AnimationNode, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNode(ctx gd.Lifetime, name gd.StringName) object.AnimationNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Removes a sub animation node.
*/
//go:nosplit
func (self class) RemoveNode(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes the name of a sub animation node.
*/
//go:nosplit
func (self class) RenameNode(name gd.StringName, new_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(new_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_rename_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
//go:nosplit
func (self class) HasNode(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
//go:nosplit
func (self class) ConnectNode(input_node gd.StringName, input_index gd.Int, output_node gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(input_node))
	callframe.Arg(frame, input_index)
	callframe.Arg(frame, mmm.Get(output_node))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disconnects the animation node connected to the specified input.
*/
//go:nosplit
func (self class) DisconnectNode(input_node gd.StringName, input_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(input_node))
	callframe.Arg(frame, input_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Modifies the position of a sub animation node.
*/
//go:nosplit
func (self class) SetNodePosition(name gd.StringName, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNodePosition(name gd.StringName) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeBlendTree.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationNodeBlendTree() Expert { return self[0].AsAnimationNodeBlendTree() }


//go:nosplit
func (self Simple) AsAnimationNodeBlendTree() Simple { return self[0].AsAnimationNodeBlendTree() }


//go:nosplit
func (self class) AsAnimationRootNode() AnimationRootNode.Expert { return self[0].AsAnimationRootNode() }


//go:nosplit
func (self Simple) AsAnimationRootNode() AnimationRootNode.Simple { return self[0].AsAnimationRootNode() }


//go:nosplit
func (self class) AsAnimationNode() AnimationNode.Expert { return self[0].AsAnimationNode() }


//go:nosplit
func (self Simple) AsAnimationNode() AnimationNode.Simple { return self[0].AsAnimationNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AnimationNodeBlendTree", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
