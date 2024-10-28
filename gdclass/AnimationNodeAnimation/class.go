package AnimationNodeAnimation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationRootNode"
import "grow.graphics/gd/gdclass/AnimationNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A resource to add to an [AnimationNodeBlendTree]. Only has one output port using the [member animation] property. Used as an input for [AnimationNode]s that blend animations together.

*/
type Go [1]classdb.AnimationNodeAnimation
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNodeAnimation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeAnimation"))
	return Go{classdb.AnimationNodeAnimation(object)}
}

func (self Go) Animation() string {
		return string(class(self).GetAnimation().String())
}

func (self Go) SetAnimation(value string) {
	class(self).SetAnimation(gd.NewStringName(value))
}

func (self Go) PlayMode() classdb.AnimationNodeAnimationPlayMode {
		return classdb.AnimationNodeAnimationPlayMode(class(self).GetPlayMode())
}

func (self Go) SetPlayMode(value classdb.AnimationNodeAnimationPlayMode) {
	class(self).SetPlayMode(value)
}

func (self Go) UseCustomTimeline() bool {
		return bool(class(self).IsUsingCustomTimeline())
}

func (self Go) SetUseCustomTimeline(value bool) {
	class(self).SetUseCustomTimeline(value)
}

func (self Go) TimelineLength() float64 {
		return float64(float64(class(self).GetTimelineLength()))
}

func (self Go) SetTimelineLength(value float64) {
	class(self).SetTimelineLength(gd.Float(value))
}

func (self Go) StretchTimeScale() bool {
		return bool(class(self).IsStretchingTimeScale())
}

func (self Go) SetStretchTimeScale(value bool) {
	class(self).SetStretchTimeScale(value)
}

func (self Go) StartOffset() float64 {
		return float64(float64(class(self).GetStartOffset()))
}

func (self Go) SetStartOffset(value float64) {
	class(self).SetStartOffset(gd.Float(value))
}

func (self Go) LoopMode() classdb.AnimationLoopMode {
		return classdb.AnimationLoopMode(class(self).GetLoopMode())
}

func (self Go) SetLoopMode(value classdb.AnimationLoopMode) {
	class(self).SetLoopMode(value)
}

//go:nosplit
func (self class) SetAnimation(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnimation() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlayMode(mode classdb.AnimationNodeAnimationPlayMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_play_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlayMode() classdb.AnimationNodeAnimationPlayMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeAnimationPlayMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_play_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomTimeline(use_custom_timeline bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, use_custom_timeline)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_use_custom_timeline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomTimeline() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_is_using_custom_timeline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTimelineLength(timeline_length gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, timeline_length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_timeline_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimelineLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_timeline_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchTimeScale(stretch_time_scale bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, stretch_time_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_stretch_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsStretchingTimeScale() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_is_stretching_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStartOffset(start_offset gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, start_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_start_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStartOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_start_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopMode(loop_mode classdb.AnimationLoopMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopMode() classdb.AnimationLoopMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationLoopMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeAnimation() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNodeAnimation() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationRootNode() AnimationRootNode.GD { return *((*AnimationRootNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationRootNode() AnimationRootNode.Go { return *((*AnimationRootNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNode() AnimationNode.GD { return *((*AnimationNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNode() AnimationNode.Go { return *((*AnimationNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationRootNode(), name)
	}
}
func init() {classdb.Register("AnimationNodeAnimation", func(ptr gd.Object) any { return classdb.AnimationNodeAnimation(ptr) })}
type PlayMode = classdb.AnimationNodeAnimationPlayMode

const (
/*Plays animation in forward direction.*/
	PlayModeForward PlayMode = 0
/*Plays animation in backward direction.*/
	PlayModeBackward PlayMode = 1
)
