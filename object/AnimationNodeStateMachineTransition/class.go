package AnimationNodeStateMachineTransition

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
The path generated when using [method AnimationNodeStateMachinePlayback.travel] is limited to the nodes connected by [AnimationNodeStateMachineTransition].
You can set the timing and conditions of the transition in detail.

*/
type Simple [1]classdb.AnimationNodeStateMachineTransition
func (self Simple) SetSwitchMode(mode classdb.AnimationNodeStateMachineTransitionSwitchMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSwitchMode(mode)
}
func (self Simple) GetSwitchMode() classdb.AnimationNodeStateMachineTransitionSwitchMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationNodeStateMachineTransitionSwitchMode(Expert(self).GetSwitchMode())
}
func (self Simple) SetAdvanceMode(mode classdb.AnimationNodeStateMachineTransitionAdvanceMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdvanceMode(mode)
}
func (self Simple) GetAdvanceMode() classdb.AnimationNodeStateMachineTransitionAdvanceMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationNodeStateMachineTransitionAdvanceMode(Expert(self).GetAdvanceMode())
}
func (self Simple) SetAdvanceCondition(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdvanceCondition(gc.StringName(name))
}
func (self Simple) GetAdvanceCondition() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetAdvanceCondition(gc).String())
}
func (self Simple) SetXfadeTime(secs float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetXfadeTime(gd.Float(secs))
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
func (self Simple) SetBreakLoopAtEnd(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBreakLoopAtEnd(enable)
}
func (self Simple) IsLoopBrokenAtEnd() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLoopBrokenAtEnd())
}
func (self Simple) SetReset(reset bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReset(reset)
}
func (self Simple) IsReset() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsReset())
}
func (self Simple) SetPriority(priority int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPriority(gd.Int(priority))
}
func (self Simple) GetPriority() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPriority()))
}
func (self Simple) SetAdvanceExpression(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdvanceExpression(gc.String(text))
}
func (self Simple) GetAdvanceExpression() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetAdvanceExpression(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeStateMachineTransition
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSwitchMode(mode classdb.AnimationNodeStateMachineTransitionSwitchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_switch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSwitchMode() classdb.AnimationNodeStateMachineTransitionSwitchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeStateMachineTransitionSwitchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_switch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdvanceMode(mode classdb.AnimationNodeStateMachineTransitionAdvanceMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_advance_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdvanceMode() classdb.AnimationNodeStateMachineTransitionAdvanceMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeStateMachineTransitionAdvanceMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_advance_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdvanceCondition(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_advance_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdvanceCondition(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_advance_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetXfadeTime(secs gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetXfadeTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetXfadeCurve(ctx gd.Lifetime) object.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBreakLoopAtEnd(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_break_loop_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLoopBrokenAtEnd() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_is_loop_broken_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReset(reset bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, reset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsReset() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_is_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPriority(priority gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPriority() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdvanceExpression(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_advance_expression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdvanceExpression(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_advance_expression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationNodeStateMachineTransition() Expert { return self[0].AsAnimationNodeStateMachineTransition() }


//go:nosplit
func (self Simple) AsAnimationNodeStateMachineTransition() Simple { return self[0].AsAnimationNodeStateMachineTransition() }


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
func init() {classdb.Register("AnimationNodeStateMachineTransition", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type SwitchMode = classdb.AnimationNodeStateMachineTransitionSwitchMode

const (
/*Switch to the next state immediately. The current state will end and blend into the beginning of the new one.*/
	SwitchModeImmediate SwitchMode = 0
/*Switch to the next state immediately, but will seek the new state to the playback position of the old state.*/
	SwitchModeSync SwitchMode = 1
/*Wait for the current state playback to end, then switch to the beginning of the next state animation.*/
	SwitchModeAtEnd SwitchMode = 2
)
type AdvanceMode = classdb.AnimationNodeStateMachineTransitionAdvanceMode

const (
/*Don't use this transition.*/
	AdvanceModeDisabled AdvanceMode = 0
/*Only use this transition during [method AnimationNodeStateMachinePlayback.travel].*/
	AdvanceModeEnabled AdvanceMode = 1
/*Automatically use this transition if the [member advance_condition] and [member advance_expression] checks are true (if assigned).*/
	AdvanceModeAuto AdvanceMode = 2
)
