package AnimationNodeStateMachine

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
Contains multiple [AnimationRootNode]s representing animation states, connected in a graph. State transitions can be configured to happen automatically or via code, using a shortest-path algorithm. Retrieve the [AnimationNodeStateMachinePlayback] object from the [AnimationTree] node to control it programmatically.
[b]Example:[/b]
[codeblocks]
[gdscript]
var state_machine = $AnimationTree.get("parameters/playback")
state_machine.travel("some_state")
[/gdscript]
[csharp]
var stateMachine = GetNode<AnimationTree>("AnimationTree").Get("parameters/playback") as AnimationNodeStateMachinePlayback;
stateMachine.Travel("some_state");
[/csharp]
[/codeblocks]

*/
type Simple [1]classdb.AnimationNodeStateMachine
func (self Simple) AddNode(name string, node [1]classdb.AnimationNode, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddNode(gc.StringName(name), node, position)
}
func (self Simple) ReplaceNode(name string, node [1]classdb.AnimationNode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReplaceNode(gc.StringName(name), node)
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
func (self Simple) GetNodeName(node [1]classdb.AnimationNode) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetNodeName(gc, node).String())
}
func (self Simple) SetNodePosition(name string, position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNodePosition(gc.StringName(name), position)
}
func (self Simple) GetNodePosition(name string) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetNodePosition(gc.StringName(name)))
}
func (self Simple) HasTransition(from string, to string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasTransition(gc.StringName(from), gc.StringName(to)))
}
func (self Simple) AddTransition(from string, to string, transition [1]classdb.AnimationNodeStateMachineTransition) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddTransition(gc.StringName(from), gc.StringName(to), transition)
}
func (self Simple) GetTransition(idx int) [1]classdb.AnimationNodeStateMachineTransition {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AnimationNodeStateMachineTransition(Expert(self).GetTransition(gc, gd.Int(idx)))
}
func (self Simple) GetTransitionFrom(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTransitionFrom(gc, gd.Int(idx)).String())
}
func (self Simple) GetTransitionTo(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTransitionTo(gc, gd.Int(idx)).String())
}
func (self Simple) GetTransitionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTransitionCount()))
}
func (self Simple) RemoveTransitionByIndex(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTransitionByIndex(gd.Int(idx))
}
func (self Simple) RemoveTransition(from string, to string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTransition(gc.StringName(from), gc.StringName(to))
}
func (self Simple) SetGraphOffset(offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGraphOffset(offset)
}
func (self Simple) GetGraphOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGraphOffset())
}
func (self Simple) SetStateMachineType(state_machine_type classdb.AnimationNodeStateMachineStateMachineType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStateMachineType(state_machine_type)
}
func (self Simple) GetStateMachineType() classdb.AnimationNodeStateMachineStateMachineType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationNodeStateMachineStateMachineType(Expert(self).GetStateMachineType())
}
func (self Simple) SetAllowTransitionToSelf(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowTransitionToSelf(enable)
}
func (self Simple) IsAllowTransitionToSelf() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAllowTransitionToSelf())
}
func (self Simple) SetResetEnds(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetResetEnds(enable)
}
func (self Simple) AreEndsReset() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AreEndsReset())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeStateMachine
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds a new animation node to the graph. The [param position] is used for display in the editor.
*/
//go:nosplit
func (self class) AddNode(name gd.StringName, node object.AnimationNode, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the given animation node with a new animation node.
*/
//go:nosplit
func (self class) ReplaceNode(name gd.StringName, node object.AnimationNode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_replace_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the animation node with the given name.
*/
//go:nosplit
func (self class) GetNode(ctx gd.Lifetime, name gd.StringName) object.AnimationNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Deletes the given animation node from the graph.
*/
//go:nosplit
func (self class) RemoveNode(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renames the given animation node.
*/
//go:nosplit
func (self class) RenameNode(name gd.StringName, new_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(new_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_rename_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the graph contains the given animation node.
*/
//go:nosplit
func (self class) HasNode(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the given animation node's name.
*/
//go:nosplit
func (self class) GetNodeName(ctx gd.Lifetime, node object.AnimationNode) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the animation node's coordinates. Used for display in the editor.
*/
//go:nosplit
func (self class) SetNodePosition(name gd.StringName, position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given animation node's coordinates. Used for display in the editor.
*/
//go:nosplit
func (self class) GetNodePosition(name gd.StringName) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a transition between the given animation nodes.
*/
//go:nosplit
func (self class) HasTransition(from gd.StringName, to gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_has_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a transition between the given animation nodes.
*/
//go:nosplit
func (self class) AddTransition(from gd.StringName, to gd.StringName, transition object.AnimationNodeStateMachineTransition)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	callframe.Arg(frame, mmm.Get(transition[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_add_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given transition.
*/
//go:nosplit
func (self class) GetTransition(ctx gd.Lifetime, idx gd.Int) object.AnimationNodeStateMachineTransition {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationNodeStateMachineTransition
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the given transition's start node.
*/
//go:nosplit
func (self class) GetTransitionFrom(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_transition_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the given transition's end node.
*/
//go:nosplit
func (self class) GetTransitionTo(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_transition_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of connections in the graph.
*/
//go:nosplit
func (self class) GetTransitionCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_transition_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Deletes the given transition by index.
*/
//go:nosplit
func (self class) RemoveTransitionByIndex(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_remove_transition_by_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes the transition between the two specified animation nodes.
*/
//go:nosplit
func (self class) RemoveTransition(from gd.StringName, to gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from))
	callframe.Arg(frame, mmm.Get(to))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_remove_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the draw offset of the graph. Used for display in the editor.
*/
//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the draw offset of the graph. Used for display in the editor.
*/
//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStateMachineType(state_machine_type classdb.AnimationNodeStateMachineStateMachineType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, state_machine_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_set_state_machine_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStateMachineType() classdb.AnimationNodeStateMachineStateMachineType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeStateMachineStateMachineType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_get_state_machine_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowTransitionToSelf(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_set_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllowTransitionToSelf() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_is_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetResetEnds(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_set_reset_ends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreEndsReset() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachine.Bind_are_ends_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationNodeStateMachine() Expert { return self[0].AsAnimationNodeStateMachine() }


//go:nosplit
func (self Simple) AsAnimationNodeStateMachine() Simple { return self[0].AsAnimationNodeStateMachine() }


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
func init() {classdb.Register("AnimationNodeStateMachine", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type StateMachineType = classdb.AnimationNodeStateMachineStateMachineType

const (
/*Seeking to the beginning is treated as playing from the start state. Transition to the end state is treated as exiting the state machine.*/
	StateMachineTypeRoot StateMachineType = 0
/*Seeking to the beginning is treated as seeking to the beginning of the animation in the current state. Transition to the end state, or the absence of transitions in each state, is treated as exiting the state machine.*/
	StateMachineTypeNested StateMachineType = 1
/*This is a grouped state machine that can be controlled from a parent state machine. It does not work independently. There must be a state machine with [member state_machine_type] of [constant STATE_MACHINE_TYPE_ROOT] or [constant STATE_MACHINE_TYPE_NESTED] in the parent or ancestor.*/
	StateMachineTypeGrouped StateMachineType = 2
)
