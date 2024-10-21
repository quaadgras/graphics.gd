package AnimationNodeStateMachinePlayback

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Simple [1]classdb.AnimationNodeStateMachinePlayback
func (self Simple) Travel(to_node string, reset_on_teleport bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Travel(gc.StringName(to_node), reset_on_teleport)
}
func (self Simple) Start(node string, reset bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Start(gc.StringName(node), reset)
}
func (self Simple) Next() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Next()
}
func (self Simple) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Stop()
}
func (self Simple) IsPlaying() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPlaying())
}
func (self Simple) GetCurrentNode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCurrentNode(gc).String())
}
func (self Simple) GetCurrentPlayPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCurrentPlayPosition()))
}
func (self Simple) GetCurrentLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCurrentLength()))
}
func (self Simple) GetFadingFromNode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetFadingFromNode(gc).String())
}
func (self Simple) GetTravelPath() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](Expert(self).GetTravelPath(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeStateMachinePlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Transitions from the current state to another one, following the shortest path.
If the path does not connect from the current state, the animation will play after the state teleports.
If [param reset_on_teleport] is [code]true[/code], the animation is played from the beginning when the travel cause a teleportation.
*/
//go:nosplit
func (self class) Travel(to_node gd.StringName, reset_on_teleport bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(to_node))
	callframe.Arg(frame, reset_on_teleport)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_travel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts playing the given animation.
If [param reset] is [code]true[/code], the animation is played from the beginning.
*/
//go:nosplit
func (self class) Start(node gd.StringName, reset bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node))
	callframe.Arg(frame, reset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If there is a next path by travel or auto advance, immediately transitions from the current state to the next state.
*/
//go:nosplit
func (self class) Next()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the currently playing animation.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if an animation is playing.
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the currently playing animation state.
[b]Note:[/b] When using a cross-fade, the current state changes to the next state immediately after the cross-fade begins.
*/
//go:nosplit
func (self class) GetCurrentNode(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_get_current_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the playback position within the current animation state.
*/
//go:nosplit
func (self class) GetCurrentPlayPosition() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_get_current_play_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_get_current_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the starting state of currently fading animation.
*/
//go:nosplit
func (self class) GetFadingFromNode(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_get_fading_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current travel path as computed internally by the A* algorithm.
*/
//go:nosplit
func (self class) GetTravelPath(ctx gd.Lifetime) gd.ArrayOf[gd.StringName] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachinePlayback.Bind_get_travel_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.StringName](ret)
}

//go:nosplit
func (self class) AsAnimationNodeStateMachinePlayback() Expert { return self[0].AsAnimationNodeStateMachinePlayback() }


//go:nosplit
func (self Simple) AsAnimationNodeStateMachinePlayback() Simple { return self[0].AsAnimationNodeStateMachinePlayback() }


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
func init() {classdb.Register("AnimationNodeStateMachinePlayback", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
