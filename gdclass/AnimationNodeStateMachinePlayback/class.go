package AnimationNodeStateMachinePlayback

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
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
type Go [1]classdb.AnimationNodeStateMachinePlayback

/*
Transitions from the current state to another one, following the shortest path.
If the path does not connect from the current state, the animation will play after the state teleports.
If [param reset_on_teleport] is [code]true[/code], the animation is played from the beginning when the travel cause a teleportation.
*/
func (self Go) Travel(to_node string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Travel(gc.StringName(to_node), true)
}

/*
Starts playing the given animation.
If [param reset] is [code]true[/code], the animation is played from the beginning.
*/
func (self Go) Start(node string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Start(gc.StringName(node), true)
}

/*
If there is a next path by travel or auto advance, immediately transitions from the current state to the next state.
*/
func (self Go) Next() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Next()
}

/*
Stops the currently playing animation.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}

/*
Returns [code]true[/code] if an animation is playing.
*/
func (self Go) IsPlaying() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPlaying())
}

/*
Returns the currently playing animation state.
[b]Note:[/b] When using a cross-fade, the current state changes to the next state immediately after the cross-fade begins.
*/
func (self Go) GetCurrentNode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCurrentNode(gc).String())
}

/*
Returns the playback position within the current animation state.
*/
func (self Go) GetCurrentPlayPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCurrentPlayPosition()))
}

/*
Returns the current state length.
[b]Note:[/b] It is possible that any [AnimationRootNode] can be nodes as well as animations. This means that there can be multiple animations within a single state. Which animation length has priority depends on the nodes connected inside it. Also, if a transition does not reset, the remaining length at that point will be returned.
*/
func (self Go) GetCurrentLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetCurrentLength()))
}

/*
Returns the starting state of currently fading animation.
*/
func (self Go) GetFadingFromNode() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFadingFromNode(gc).String())
}

/*
Returns the current travel path as computed internally by the A* algorithm.
*/
func (self Go) GetTravelPath() gd.ArrayOf[gd.StringName] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.StringName](class(self).GetTravelPath(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeStateMachinePlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationNodeStateMachinePlayback"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (self class) AsAnimationNodeStateMachinePlayback() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeStateMachinePlayback() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("AnimationNodeStateMachinePlayback", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
