package AnimationNodeAnimation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationRootNode"
import "grow.graphics/gd/object/AnimationNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource to add to an [AnimationNodeBlendTree]. Only has one output port using the [member animation] property. Used as an input for [AnimationNode]s that blend animations together.

*/
type Simple [1]classdb.AnimationNodeAnimation
func (self Simple) SetAnimation(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnimation(gc.StringName(name))
}
func (self Simple) GetAnimation() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetAnimation(gc).String())
}
func (self Simple) SetPlayMode(mode classdb.AnimationNodeAnimationPlayMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlayMode(mode)
}
func (self Simple) GetPlayMode() classdb.AnimationNodeAnimationPlayMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationNodeAnimationPlayMode(Expert(self).GetPlayMode())
}
func (self Simple) SetUseCustomTimeline(use_custom_timeline bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseCustomTimeline(use_custom_timeline)
}
func (self Simple) IsUsingCustomTimeline() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingCustomTimeline())
}
func (self Simple) SetTimelineLength(timeline_length float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTimelineLength(gd.Float(timeline_length))
}
func (self Simple) GetTimelineLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTimelineLength()))
}
func (self Simple) SetStretchTimeScale(stretch_time_scale bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStretchTimeScale(stretch_time_scale)
}
func (self Simple) IsStretchingTimeScale() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsStretchingTimeScale())
}
func (self Simple) SetStartOffset(start_offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStartOffset(gd.Float(start_offset))
}
func (self Simple) GetStartOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetStartOffset()))
}
func (self Simple) SetLoopMode(loop_mode classdb.AnimationLoopMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLoopMode(loop_mode)
}
func (self Simple) GetLoopMode() classdb.AnimationLoopMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationLoopMode(Expert(self).GetLoopMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeAnimation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetAnimation(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnimation(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlayMode(mode classdb.AnimationNodeAnimationPlayMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_play_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlayMode() classdb.AnimationNodeAnimationPlayMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeAnimationPlayMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_get_play_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomTimeline(use_custom_timeline bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_custom_timeline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_use_custom_timeline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomTimeline() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_is_using_custom_timeline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTimelineLength(timeline_length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, timeline_length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_timeline_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimelineLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_get_timeline_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchTimeScale(stretch_time_scale bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch_time_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_stretch_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsStretchingTimeScale() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_is_stretching_time_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStartOffset(start_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, start_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_start_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStartOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_get_start_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopMode(loop_mode classdb.AnimationLoopMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopMode() classdb.AnimationLoopMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationLoopMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationNodeAnimation.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationNodeAnimation() Expert { return self[0].AsAnimationNodeAnimation() }


//go:nosplit
func (self Simple) AsAnimationNodeAnimation() Simple { return self[0].AsAnimationNodeAnimation() }


//go:nosplit
func (self class) AsAnimationRootNode() AnimationRootNode.Expert { return self[0].AsAnimationRootNode() }


//go:nosplit
func (self Simple) AsAnimationRootNode() AnimationRootNode.Simple { return self[0].AsAnimationRootNode() }


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
func init() {classdb.Register("AnimationNodeAnimation", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type PlayMode = classdb.AnimationNodeAnimationPlayMode

const (
/*Plays animation in forward direction.*/
	PlayModeForward PlayMode = 0
/*Plays animation in backward direction.*/
	PlayModeBackward PlayMode = 1
)
