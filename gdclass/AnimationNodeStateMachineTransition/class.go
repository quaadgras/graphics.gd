package AnimationNodeStateMachineTransition

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
The path generated when using [method AnimationNodeStateMachinePlayback.travel] is limited to the nodes connected by [AnimationNodeStateMachineTransition].
You can set the timing and conditions of the transition in detail.

*/
type Go [1]classdb.AnimationNodeStateMachineTransition
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeStateMachineTransition
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationNodeStateMachineTransition"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) XfadeTime() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetXfadeTime()))
}

func (self Go) SetXfadeTime(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetXfadeTime(gd.Float(value))
}

func (self Go) XfadeCurve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetXfadeCurve(gc))
}

func (self Go) SetXfadeCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetXfadeCurve(value)
}

func (self Go) BreakLoopAtEnd() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsLoopBrokenAtEnd())
}

func (self Go) SetBreakLoopAtEnd(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBreakLoopAtEnd(value)
}

func (self Go) Reset() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsReset())
}

func (self Go) SetReset(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReset(value)
}

func (self Go) Priority() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPriority()))
}

func (self Go) SetPriority(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPriority(gd.Int(value))
}

func (self Go) SwitchMode() classdb.AnimationNodeStateMachineTransitionSwitchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AnimationNodeStateMachineTransitionSwitchMode(class(self).GetSwitchMode())
}

func (self Go) SetSwitchMode(value classdb.AnimationNodeStateMachineTransitionSwitchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwitchMode(value)
}

func (self Go) AdvanceMode() classdb.AnimationNodeStateMachineTransitionAdvanceMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AnimationNodeStateMachineTransitionAdvanceMode(class(self).GetAdvanceMode())
}

func (self Go) SetAdvanceMode(value classdb.AnimationNodeStateMachineTransitionAdvanceMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdvanceMode(value)
}

func (self Go) AdvanceCondition() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAdvanceCondition(gc).String())
}

func (self Go) SetAdvanceCondition(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdvanceCondition(gc.StringName(value))
}

func (self Go) AdvanceExpression() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAdvanceExpression(gc).String())
}

func (self Go) SetAdvanceExpression(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAdvanceExpression(gc.String(value))
}

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
func (self class) SetXfadeCurve(curve gdclass.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_set_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetXfadeCurve(ctx gd.Lifetime) gdclass.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeStateMachineTransition.Bind_get_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Curve
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
func (self Go) OnAdvanceConditionChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("advance_condition_changed"), gc.Callable(cb), 0)
}


func (self class) AsAnimationNodeStateMachineTransition() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeStateMachineTransition() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AnimationNodeStateMachineTransition", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
