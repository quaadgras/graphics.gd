package AnimationNodeTransition

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationNodeSync"
import "grow.graphics/gd/object/AnimationNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Simple state machine for cases which don't require a more advanced [AnimationNodeStateMachine]. Animations can be connected to the inputs and transition times can be specified.
After setting the request and changing the animation playback, the transition node automatically clears the request on the next process frame by setting its [code]transition_request[/code] value to empty.
[b]Note:[/b] When using a cross-fade, [code]current_state[/code] and [code]current_index[/code] change to the next state immediately after the cross-fade begins.
[codeblocks]
[gdscript]
# Play child animation connected to "state_2" port.
animation_tree.set("parameters/Transition/transition_request", "state_2")
# Alternative syntax (same result as above).
animation_tree["parameters/Transition/transition_request"] = "state_2"

# Get current state name (read-only).
animation_tree.get("parameters/Transition/current_state")
# Alternative syntax (same result as above).
animation_tree["parameters/Transition/current_state"]

# Get current state index (read-only).
animation_tree.get("parameters/Transition/current_index")
# Alternative syntax (same result as above).
animation_tree["parameters/Transition/current_index"]
[/gdscript]
[csharp]
// Play child animation connected to "state_2" port.
animationTree.Set("parameters/Transition/transition_request", "state_2");

// Get current state name (read-only).
animationTree.Get("parameters/Transition/current_state");

// Get current state index (read-only).
animationTree.Get("parameters/Transition/current_index");
[/csharp]
[/codeblocks]

*/
type Simple [1]classdb.AnimationNodeTransition
func (self Simple) SetInputCount(input_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputCount(gd.Int(input_count))
}
func (self Simple) SetInputAsAutoAdvance(input int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputAsAutoAdvance(gd.Int(input), enable)
}
func (self Simple) IsInputSetAsAutoAdvance(input int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInputSetAsAutoAdvance(gd.Int(input)))
}
func (self Simple) SetInputBreakLoopAtEnd(input int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputBreakLoopAtEnd(gd.Int(input), enable)
}
func (self Simple) IsInputLoopBrokenAtEnd(input int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInputLoopBrokenAtEnd(gd.Int(input)))
}
func (self Simple) SetInputReset(input int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInputReset(gd.Int(input), enable)
}
func (self Simple) IsInputReset(input int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInputReset(gd.Int(input)))
}
func (self Simple) SetXfadeTime(time float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetXfadeTime(gd.Float(time))
}
func (self Simple) GetXfadeTime() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetXfadeTime()))
}
func (self Simple) SetXfadeCurve(curve [1]classdb.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetXfadeCurve(curve)
}
func (self Simple) GetXfadeCurve() [1]classdb.Curve {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Curve(Expert(self).GetXfadeCurve(gc))
}
func (self Simple) SetAllowTransitionToSelf(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowTransitionToSelf(enable)
}
func (self Simple) IsAllowTransitionToSelf() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAllowTransitionToSelf())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeTransition
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetInputCount(input_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_input_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables auto-advance for the given [param input] index. If enabled, state changes to the next input after playing the animation once. If enabled for the last input state, it loops to the first.
*/
//go:nosplit
func (self class) SetInputAsAutoAdvance(input gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_input_as_auto_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if auto-advance is enabled for the given [param input] index.
*/
//go:nosplit
func (self class) IsInputSetAsAutoAdvance(input gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_is_input_set_as_auto_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], breaks the loop at the end of the loop cycle for transition, even if the animation is looping.
*/
//go:nosplit
func (self class) SetInputBreakLoopAtEnd(input gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_input_break_loop_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the animation breaks the loop at the end of the loop cycle for transition.
*/
//go:nosplit
func (self class) IsInputLoopBrokenAtEnd(input gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_is_input_loop_broken_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], the destination animation is restarted when the animation transitions.
*/
//go:nosplit
func (self class) SetInputReset(input gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_input_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the animation restarts when the animation transitions from the other animation.
*/
//go:nosplit
func (self class) IsInputReset(input gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_is_input_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetXfadeTime(time gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetXfadeTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_get_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetXfadeCurve(curve object.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetXfadeCurve(ctx gd.Lifetime) object.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_get_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowTransitionToSelf(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_set_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllowTransitionToSelf() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeTransition.Bind_is_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationNodeTransition() Expert { return self[0].AsAnimationNodeTransition() }


//go:nosplit
func (self Simple) AsAnimationNodeTransition() Simple { return self[0].AsAnimationNodeTransition() }


//go:nosplit
func (self class) AsAnimationNodeSync() AnimationNodeSync.Expert { return self[0].AsAnimationNodeSync() }


//go:nosplit
func (self Simple) AsAnimationNodeSync() AnimationNodeSync.Simple { return self[0].AsAnimationNodeSync() }


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
func init() {classdb.Register("AnimationNodeTransition", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
