// Package AnimationNodeStateMachine provides methods for working with AnimationNodeStateMachine object instances.
package AnimationNodeStateMachine

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
type Instance [1]gdclass.AnimationNodeStateMachine

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeStateMachine() Instance
}

/*
Adds a new animation node to the graph. The [param position] is used for display in the editor.
*/
func (self Instance) AddNode(name string, node [1]gdclass.AnimationNode) {
	class(self).AddNode(gd.NewStringName(name), node, gd.Vector2(gd.Vector2{0, 0}))
}

/*
Replaces the given animation node with a new animation node.
*/
func (self Instance) ReplaceNode(name string, node [1]gdclass.AnimationNode) {
	class(self).ReplaceNode(gd.NewStringName(name), node)
}

/*
Returns the animation node with the given name.
*/
func (self Instance) GetNode(name string) [1]gdclass.AnimationNode {
	return [1]gdclass.AnimationNode(class(self).GetNode(gd.NewStringName(name)))
}

/*
Deletes the given animation node from the graph.
*/
func (self Instance) RemoveNode(name string) {
	class(self).RemoveNode(gd.NewStringName(name))
}

/*
Renames the given animation node.
*/
func (self Instance) RenameNode(name string, new_name string) {
	class(self).RenameNode(gd.NewStringName(name), gd.NewStringName(new_name))
}

/*
Returns [code]true[/code] if the graph contains the given animation node.
*/
func (self Instance) HasNode(name string) bool {
	return bool(class(self).HasNode(gd.NewStringName(name)))
}

/*
Returns the given animation node's name.
*/
func (self Instance) GetNodeName(node [1]gdclass.AnimationNode) string {
	return string(class(self).GetNodeName(node).String())
}

/*
Sets the animation node's coordinates. Used for display in the editor.
*/
func (self Instance) SetNodePosition(name string, position Vector2.XY) {
	class(self).SetNodePosition(gd.NewStringName(name), gd.Vector2(position))
}

/*
Returns the given animation node's coordinates. Used for display in the editor.
*/
func (self Instance) GetNodePosition(name string) Vector2.XY {
	return Vector2.XY(class(self).GetNodePosition(gd.NewStringName(name)))
}

/*
Returns [code]true[/code] if there is a transition between the given animation nodes.
*/
func (self Instance) HasTransition(from string, to string) bool {
	return bool(class(self).HasTransition(gd.NewStringName(from), gd.NewStringName(to)))
}

/*
Adds a transition between the given animation nodes.
*/
func (self Instance) AddTransition(from string, to string, transition [1]gdclass.AnimationNodeStateMachineTransition) {
	class(self).AddTransition(gd.NewStringName(from), gd.NewStringName(to), transition)
}

/*
Returns the given transition.
*/
func (self Instance) GetTransition(idx int) [1]gdclass.AnimationNodeStateMachineTransition {
	return [1]gdclass.AnimationNodeStateMachineTransition(class(self).GetTransition(gd.Int(idx)))
}

/*
Returns the given transition's start node.
*/
func (self Instance) GetTransitionFrom(idx int) string {
	return string(class(self).GetTransitionFrom(gd.Int(idx)).String())
}

/*
Returns the given transition's end node.
*/
func (self Instance) GetTransitionTo(idx int) string {
	return string(class(self).GetTransitionTo(gd.Int(idx)).String())
}

/*
Returns the number of connections in the graph.
*/
func (self Instance) GetTransitionCount() int {
	return int(int(class(self).GetTransitionCount()))
}

/*
Deletes the given transition by index.
*/
func (self Instance) RemoveTransitionByIndex(idx int) {
	class(self).RemoveTransitionByIndex(gd.Int(idx))
}

/*
Deletes the transition between the two specified animation nodes.
*/
func (self Instance) RemoveTransition(from string, to string) {
	class(self).RemoveTransition(gd.NewStringName(from), gd.NewStringName(to))
}

/*
Sets the draw offset of the graph. Used for display in the editor.
*/
func (self Instance) SetGraphOffset(offset Vector2.XY) {
	class(self).SetGraphOffset(gd.Vector2(offset))
}

/*
Returns the draw offset of the graph. Used for display in the editor.
*/
func (self Instance) GetGraphOffset() Vector2.XY {
	return Vector2.XY(class(self).GetGraphOffset())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeStateMachine

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeStateMachine"))
	casted := Instance{*(*gdclass.AnimationNodeStateMachine)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) StateMachineType() gdclass.AnimationNodeStateMachineStateMachineType {
	return gdclass.AnimationNodeStateMachineStateMachineType(class(self).GetStateMachineType())
}

func (self Instance) SetStateMachineType(value gdclass.AnimationNodeStateMachineStateMachineType) {
	class(self).SetStateMachineType(value)
}

func (self Instance) AllowTransitionToSelf() bool {
	return bool(class(self).IsAllowTransitionToSelf())
}

func (self Instance) SetAllowTransitionToSelf(value bool) {
	class(self).SetAllowTransitionToSelf(value)
}

func (self Instance) ResetEnds() bool {
	return bool(class(self).AreEndsReset())
}

func (self Instance) SetResetEnds(value bool) {
	class(self).SetResetEnds(value)
}

/*
Adds a new animation node to the graph. The [param position] is used for display in the editor.
*/
//go:nosplit
func (self class) AddNode(name gd.StringName, node [1]gdclass.AnimationNode, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Replaces the given animation node with a new animation node.
*/
//go:nosplit
func (self class) ReplaceNode(name gd.StringName, node [1]gdclass.AnimationNode) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_replace_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the animation node with the given name.
*/
//go:nosplit
func (self class) GetNode(name gd.StringName) [1]gdclass.AnimationNode {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AnimationNode{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationNode](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Deletes the given animation node from the graph.
*/
//go:nosplit
func (self class) RemoveNode(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Renames the given animation node.
*/
//go:nosplit
func (self class) RenameNode(name gd.StringName, new_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(new_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_rename_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the graph contains the given animation node.
*/
//go:nosplit
func (self class) HasNode(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the given animation node's name.
*/
//go:nosplit
func (self class) GetNodeName(node [1]gdclass.AnimationNode) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the animation node's coordinates. Used for display in the editor.
*/
//go:nosplit
func (self class) SetNodePosition(name gd.StringName, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given animation node's coordinates. Used for display in the editor.
*/
//go:nosplit
func (self class) GetNodePosition(name gd.StringName) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if there is a transition between the given animation nodes.
*/
//go:nosplit
func (self class) HasTransition(from gd.StringName, to gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_has_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a transition between the given animation nodes.
*/
//go:nosplit
func (self class) AddTransition(from gd.StringName, to gd.StringName, transition [1]gdclass.AnimationNodeStateMachineTransition) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	callframe.Arg(frame, pointers.Get(transition[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_add_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given transition.
*/
//go:nosplit
func (self class) GetTransition(idx gd.Int) [1]gdclass.AnimationNodeStateMachineTransition {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AnimationNodeStateMachineTransition{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationNodeStateMachineTransition](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the given transition's start node.
*/
//go:nosplit
func (self class) GetTransitionFrom(idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_transition_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the given transition's end node.
*/
//go:nosplit
func (self class) GetTransitionTo(idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_transition_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of connections in the graph.
*/
//go:nosplit
func (self class) GetTransitionCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_transition_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Deletes the given transition by index.
*/
//go:nosplit
func (self class) RemoveTransitionByIndex(idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_remove_transition_by_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deletes the transition between the two specified animation nodes.
*/
//go:nosplit
func (self class) RemoveTransition(from gd.StringName, to gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_remove_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the draw offset of the graph. Used for display in the editor.
*/
//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the draw offset of the graph. Used for display in the editor.
*/
//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStateMachineType(state_machine_type gdclass.AnimationNodeStateMachineStateMachineType) {
	var frame = callframe.New()
	callframe.Arg(frame, state_machine_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_set_state_machine_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStateMachineType() gdclass.AnimationNodeStateMachineStateMachineType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationNodeStateMachineStateMachineType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_get_state_machine_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowTransitionToSelf(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_set_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowTransitionToSelf() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_is_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResetEnds(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_set_reset_ends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AreEndsReset() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachine.Bind_are_ends_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeStateMachine() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeStateMachine() Instance {
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
	gdclass.Register("AnimationNodeStateMachine", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeStateMachine{*(*gdclass.AnimationNodeStateMachine)(unsafe.Pointer(&ptr))}
	})
}

type StateMachineType = gdclass.AnimationNodeStateMachineStateMachineType

const (
	/*Seeking to the beginning is treated as playing from the start state. Transition to the end state is treated as exiting the state machine.*/
	StateMachineTypeRoot StateMachineType = 0
	/*Seeking to the beginning is treated as seeking to the beginning of the animation in the current state. Transition to the end state, or the absence of transitions in each state, is treated as exiting the state machine.*/
	StateMachineTypeNested StateMachineType = 1
	/*This is a grouped state machine that can be controlled from a parent state machine. It does not work independently. There must be a state machine with [member state_machine_type] of [constant STATE_MACHINE_TYPE_ROOT] or [constant STATE_MACHINE_TYPE_NESTED] in the parent or ancestor.*/
	StateMachineTypeGrouped StateMachineType = 2
)
