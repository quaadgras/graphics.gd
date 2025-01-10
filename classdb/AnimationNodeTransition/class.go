// Package AnimationNodeTransition provides methods for working with AnimationNodeTransition object instances.
package AnimationNodeTransition

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AnimationNodeSync"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

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
type Instance [1]gdclass.AnimationNodeTransition

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeTransition() Instance
}

/*
Enables or disables auto-advance for the given [param input] index. If enabled, state changes to the next input after playing the animation once. If enabled for the last input state, it loops to the first.
*/
func (self Instance) SetInputAsAutoAdvance(input int, enable bool) {
	class(self).SetInputAsAutoAdvance(gd.Int(input), enable)
}

/*
Returns [code]true[/code] if auto-advance is enabled for the given [param input] index.
*/
func (self Instance) IsInputSetAsAutoAdvance(input int) bool {
	return bool(class(self).IsInputSetAsAutoAdvance(gd.Int(input)))
}

/*
If [code]true[/code], breaks the loop at the end of the loop cycle for transition, even if the animation is looping.
*/
func (self Instance) SetInputBreakLoopAtEnd(input int, enable bool) {
	class(self).SetInputBreakLoopAtEnd(gd.Int(input), enable)
}

/*
Returns whether the animation breaks the loop at the end of the loop cycle for transition.
*/
func (self Instance) IsInputLoopBrokenAtEnd(input int) bool {
	return bool(class(self).IsInputLoopBrokenAtEnd(gd.Int(input)))
}

/*
If [code]true[/code], the destination animation is restarted when the animation transitions.
*/
func (self Instance) SetInputReset(input int, enable bool) {
	class(self).SetInputReset(gd.Int(input), enable)
}

/*
Returns whether the animation restarts when the animation transitions from the other animation.
*/
func (self Instance) IsInputReset(input int) bool {
	return bool(class(self).IsInputReset(gd.Int(input)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeTransition

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeTransition"))
	casted := Instance{*(*gdclass.AnimationNodeTransition)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) XfadeTime() Float.X {
	return Float.X(Float.X(class(self).GetXfadeTime()))
}

func (self Instance) SetXfadeTime(value Float.X) {
	class(self).SetXfadeTime(gd.Float(value))
}

func (self Instance) XfadeCurve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetXfadeCurve())
}

func (self Instance) SetXfadeCurve(value [1]gdclass.Curve) {
	class(self).SetXfadeCurve(value)
}

func (self Instance) AllowTransitionToSelf() bool {
	return bool(class(self).IsAllowTransitionToSelf())
}

func (self Instance) SetAllowTransitionToSelf(value bool) {
	class(self).SetAllowTransitionToSelf(value)
}

func (self Instance) SetInputCount(value int) {
	class(self).SetInputCount(gd.Int(value))
}

//go:nosplit
func (self class) SetInputCount(input_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, input_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_input_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Enables or disables auto-advance for the given [param input] index. If enabled, state changes to the next input after playing the animation once. If enabled for the last input state, it loops to the first.
*/
//go:nosplit
func (self class) SetInputAsAutoAdvance(input gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_input_as_auto_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if auto-advance is enabled for the given [param input] index.
*/
//go:nosplit
func (self class) IsInputSetAsAutoAdvance(input gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_is_input_set_as_auto_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], breaks the loop at the end of the loop cycle for transition, even if the animation is looping.
*/
//go:nosplit
func (self class) SetInputBreakLoopAtEnd(input gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_input_break_loop_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the animation breaks the loop at the end of the loop cycle for transition.
*/
//go:nosplit
func (self class) IsInputLoopBrokenAtEnd(input gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_is_input_loop_broken_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the destination animation is restarted when the animation transitions.
*/
//go:nosplit
func (self class) SetInputReset(input gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_input_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the animation restarts when the animation transitions from the other animation.
*/
//go:nosplit
func (self class) IsInputReset(input gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_is_input_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXfadeTime(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetXfadeTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_get_xfade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetXfadeCurve(curve [1]gdclass.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetXfadeCurve() [1]gdclass.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_get_xfade_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowTransitionToSelf(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_set_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowTransitionToSelf() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeTransition.Bind_is_allow_transition_to_self, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeTransition() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeTransition() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNodeSync() AnimationNodeSync.Advanced {
	return *((*AnimationNodeSync.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeSync() AnimationNodeSync.Instance {
	return *((*AnimationNodeSync.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AnimationNodeSync.Advanced(self.AsAnimationNodeSync()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationNodeSync.Instance(self.AsAnimationNodeSync()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeTransition", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeTransition{*(*gdclass.AnimationNodeTransition)(unsafe.Pointer(&ptr))}
	})
}
