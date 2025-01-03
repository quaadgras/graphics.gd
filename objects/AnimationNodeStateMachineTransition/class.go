package AnimationNodeStateMachineTransition

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The path generated when using [method AnimationNodeStateMachinePlayback.travel] is limited to the nodes connected by [AnimationNodeStateMachineTransition].
You can set the timing and conditions of the transition in detail.
*/
type Instance [1]classdb.AnimationNodeStateMachineTransition
type Any interface {
	gd.IsClass
	AsAnimationNodeStateMachineTransition() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeStateMachineTransition

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeStateMachineTransition"))
	return Instance{*(*classdb.AnimationNodeStateMachineTransition)(unsafe.Pointer(&object))}
}

func (self Instance) XfadeTime() Float.X {
	return Float.X(Float.X(class(self).GetXfadeTime()))
}

func (self Instance) SetXfadeTime(value Float.X) {
	class(self).SetXfadeTime(gd.Float(value))
}

func (self Instance) XfadeCurve() objects.Curve {
	return objects.Curve(class(self).GetXfadeCurve())
}

func (self Instance) SetXfadeCurve(value objects.Curve) {
	class(self).SetXfadeCurve(value)
}

func (self Instance) BreakLoopAtEnd() bool {
	return bool(class(self).IsLoopBrokenAtEnd())
}

func (self Instance) SetBreakLoopAtEnd(value bool) {
	class(self).SetBreakLoopAtEnd(value)
}

func (self Instance) Reset() bool {
	return bool(class(self).IsReset())
}

func (self Instance) SetReset(value bool) {
	class(self).SetReset(value)
}

func (self Instance) Priority() int {
	return int(int(class(self).GetPriority()))
}

func (self Instance) SetPriority(value int) {
	class(self).SetPriority(gd.Int(value))
}

func (self Instance) SwitchMode() classdb.AnimationNodeStateMachineTransitionSwitchMode {
	return classdb.AnimationNodeStateMachineTransitionSwitchMode(class(self).GetSwitchMode())
}

func (self Instance) SetSwitchMode(value classdb.AnimationNodeStateMachineTransitionSwitchMode) {
	class(self).SetSwitchMode(value)
}

func (self Instance) AdvanceMode() classdb.AnimationNodeStateMachineTransitionAdvanceMode {
	return classdb.AnimationNodeStateMachineTransitionAdvanceMode(class(self).GetAdvanceMode())
}

func (self Instance) SetAdvanceMode(value classdb.AnimationNodeStateMachineTransitionAdvanceMode) {
	class(self).SetAdvanceMode(value)
}

func (self Instance) AdvanceCondition() string {
	return string(class(self).GetAdvanceCondition().String())
}

func (self Instance) SetAdvanceCondition(value string) {
	class(self).SetAdvanceCondition(gd.NewStringName(value))
}

func (self Instance) AdvanceExpression() string {
	return string(class(self).GetAdvanceExpression().String())
}

func (self Instance) SetAdvanceExpression(value string) {
	class(self).SetAdvanceExpression(gd.NewString(value))
}

//go:nosplit
func (self class) SetSwitchMode(mode classdb.AnimationNodeStateMachineTransitionSwitchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_switch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSwitchMode() classdb.AnimationNodeStateMachineTransitionSwitchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeStateMachineTransitionSwitchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_switch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdvanceMode(mode classdb.AnimationNodeStateMachineTransitionAdvanceMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_advance_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdvanceMode() classdb.AnimationNodeStateMachineTransitionAdvanceMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeStateMachineTransitionAdvanceMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_advance_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdvanceCondition(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_advance_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdvanceCondition() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_advance_condition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXfadeTime(secs gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, secs)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetXfadeTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXfadeCurve(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetXfadeCurve() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{gd.PointerWithOwnershipTransferredToGo[classdb.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBreakLoopAtEnd(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_break_loop_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsLoopBrokenAtEnd() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_is_loop_broken_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReset(reset bool) {
	var frame = callframe.New()
	callframe.Arg(frame, reset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsReset() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_is_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAdvanceExpression(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_set_advance_expression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAdvanceExpression() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeStateMachineTransition.Bind_get_advance_expression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnAdvanceConditionChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("advance_condition_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationNodeStateMachineTransition() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeStateMachineTransition() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("AnimationNodeStateMachineTransition", func(ptr gd.Object) any {
		return [1]classdb.AnimationNodeStateMachineTransition{*(*classdb.AnimationNodeStateMachineTransition)(unsafe.Pointer(&ptr))}
	})
}

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
