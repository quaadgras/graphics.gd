// Package AnimationNodeStateMachinePlayback provides methods for working with AnimationNodeStateMachinePlayback object instances.
package AnimationNodeStateMachinePlayback

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Allows control of [AnimationTree] state machines created with [AnimationNodeStateMachine]. Retrieve with [code]$AnimationTree.get("parameters/playback")[/code].
[b]Example:[/b]
[codeblocks]
[gdscript]
var state_machine = $AnimationTree.get("parameters/playback")
state_machine.travel("some_state")
[/gdscript]
[csharp]
var stateMachine = GetNode<AnimationTree>("AnimationTree").Get("parameters/playback").As<AnimationNodeStateMachinePlayback>();
stateMachine.Travel("some_state");
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.AnimationNodeStateMachinePlayback
type Any interface {
	gd.IsClass
	AsAnimationNodeStateMachinePlayback() Instance
}

/*
Transitions from the current state to another one, following the shortest path.
If the path does not connect from the current state, the animation will play after the state teleports.
If [param reset_on_teleport] is [code]true[/code], the animation is played from the beginning when the travel cause a teleportation.
*/
func (self Instance) Travel(to_node string) {
	class(self).Travel(gd.NewStringName(to_node), true)
}

/*
Starts playing the given animation.
If [param reset] is [code]true[/code], the animation is played from the beginning.
*/
func (self Instance) Start(node string) {
	class(self).Start(gd.NewStringName(node), true)
}

/*
If there is a next path by travel or auto advance, immediately transitions from the current state to the next state.
*/
func (self Instance) Next() {
	class(self).Next()
}

/*
Stops the currently playing animation.
*/
func (self Instance) Stop() {
	class(self).Stop()
}

/*
Returns [code]true[/code] if an animation is playing.
*/
func (self Instance) IsPlaying() bool {
	return bool(class(self).IsPlaying())
}

/*
Returns the currently playing animation state.
[b]Note:[/b] When using a cross-fade, the current state changes to the next state immediately after the cross-fade begins.
*/
func (self Instance) GetCurrentNode() string {
	return string(class(self).GetCurrentNode().String())
}

/*
Returns the playback position within the current animation state.
*/
func (self Instance) GetCurrentPlayPosition() Float.X {
	return Float.X(Float.X(class(self).GetCurrentPlayPosition()))
}

/*
Returns the current state length.
[b]Note:[/b] It is possible that any [AnimationRootNode] can be nodes as well as animations. This means that there can be multiple animations within a single state. Which animation length has priority depends on the nodes connected inside it. Also, if a transition does not reset, the remaining length at that point will be returned.
*/
func (self Instance) GetCurrentLength() Float.X {
	return Float.X(Float.X(class(self).GetCurrentLength()))
}

/*
Returns the starting state of currently fading animation.
*/
func (self Instance) GetFadingFromNode() string {
	return string(class(self).GetFadingFromNode().String())
}

/*
Returns the current travel path as computed internally by the A* algorithm.
*/
func (self Instance) GetTravelPath() gd.Array {
	return gd.Array(class(self).GetTravelPath())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeStateMachinePlayback

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeStateMachinePlayback"))
	casted := Instance{*(*gdclass.AnimationNodeStateMachinePlayback)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Transitions from the current state to another one, following the shortest path.
If the path does not connect from the current state, the animation will play after the state teleports.
If [param reset_on_teleport] is [code]true[/code], the animation is played from the beginning when the travel cause a teleportation.
*/
//go:nosplit
func (self class) Travel(to_node gd.StringName, reset_on_teleport bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(to_node))
	callframe.Arg(frame, reset_on_teleport)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Starts playing the given animation.
If [param reset] is [code]true[/code], the animation is played from the beginning.
*/
//go:nosplit
func (self class) Start(node gd.StringName, reset bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	callframe.Arg(frame, reset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If there is a next path by travel or auto advance, immediately transitions from the current state to the next state.
*/
//go:nosplit
func (self class) Next() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Stops the currently playing animation.
*/
//go:nosplit
func (self class) Stop() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if an animation is playing.
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the currently playing animation state.
[b]Note:[/b] When using a cross-fade, the current state changes to the next state immediately after the cross-fade begins.
*/
//go:nosplit
func (self class) GetCurrentNode() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_get_current_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the playback position within the current animation state.
*/
//go:nosplit
func (self class) GetCurrentPlayPosition() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_get_current_play_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current state length.
[b]Note:[/b] It is possible that any [AnimationRootNode] can be nodes as well as animations. This means that there can be multiple animations within a single state. Which animation length has priority depends on the nodes connected inside it. Also, if a transition does not reset, the remaining length at that point will be returned.
*/
//go:nosplit
func (self class) GetCurrentLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_get_current_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the starting state of currently fading animation.
*/
//go:nosplit
func (self class) GetFadingFromNode() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_get_fading_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the current travel path as computed internally by the A* algorithm.
*/
//go:nosplit
func (self class) GetTravelPath() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachinePlayback.Bind_get_travel_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeStateMachinePlayback() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeStateMachinePlayback() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeStateMachinePlayback", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeStateMachinePlayback{*(*gdclass.AnimationNodeStateMachinePlayback)(unsafe.Pointer(&ptr))}
	})
}
