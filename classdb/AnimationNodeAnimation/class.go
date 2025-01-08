// Package AnimationNodeAnimation provides methods for working with AnimationNodeAnimation object instances.
package AnimationNodeAnimation

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AnimationRootNode"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A resource to add to an [AnimationNodeBlendTree]. Only has one output port using the [member animation] property. Used as an input for [AnimationNode]s that blend animations together.
*/
type Instance [1]gdclass.AnimationNodeAnimation

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeAnimation() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeAnimation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeAnimation"))
	casted := Instance{*(*gdclass.AnimationNodeAnimation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Animation() string {
	return string(class(self).GetAnimation().String())
}

func (self Instance) SetAnimation(value string) {
	class(self).SetAnimation(gd.NewStringName(value))
}

func (self Instance) PlayMode() gdclass.AnimationNodeAnimationPlayMode {
	return gdclass.AnimationNodeAnimationPlayMode(class(self).GetPlayMode())
}

func (self Instance) SetPlayMode(value gdclass.AnimationNodeAnimationPlayMode) {
	class(self).SetPlayMode(value)
}

func (self Instance) UseCustomTimeline() bool {
	return bool(class(self).IsUsingCustomTimeline())
}

func (self Instance) SetUseCustomTimeline(value bool) {
	class(self).SetUseCustomTimeline(value)
}

func (self Instance) TimelineLength() Float.X {
	return Float.X(Float.X(class(self).GetTimelineLength()))
}

func (self Instance) SetTimelineLength(value Float.X) {
	class(self).SetTimelineLength(gd.Float(value))
}

func (self Instance) StretchTimeScale() bool {
	return bool(class(self).IsStretchingTimeScale())
}

func (self Instance) SetStretchTimeScale(value bool) {
	class(self).SetStretchTimeScale(value)
}

func (self Instance) StartOffset() Float.X {
	return Float.X(Float.X(class(self).GetStartOffset()))
}

func (self Instance) SetStartOffset(value Float.X) {
	class(self).SetStartOffset(gd.Float(value))
}

func (self Instance) LoopMode() gdclass.AnimationLoopMode {
	return gdclass.AnimationLoopMode(class(self).GetLoopMode())
}

func (self Instance) SetLoopMode(value gdclass.AnimationLoopMode) {
	class(self).SetLoopMode(value)
}

//go:nosplit
func (self class) SetAnimation(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAnimation() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlayMode(mode gdclass.AnimationNodeAnimationPlayMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_play_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlayMode() gdclass.AnimationNodeAnimationPlayMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationNodeAnimationPlayMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_play_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomTimeline(use_custom_timeline bool) {
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
func (self class) SetTimelineLength(timeline_length gd.Float) {
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
func (self class) SetStretchTimeScale(stretch_time_scale bool) {
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
func (self class) SetStartOffset(start_offset gd.Float) {
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
func (self class) SetLoopMode(loop_mode gdclass.AnimationLoopMode) {
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopMode() gdclass.AnimationLoopMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AnimationLoopMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeAnimation.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeAnimation() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeAnimation() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationRootNode() AnimationRootNode.Advanced {
	return *((*AnimationRootNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationRootNode() AnimationRootNode.Instance {
	return *((*AnimationRootNode.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AnimationRootNode.Advanced(self.AsAnimationRootNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationRootNode.Instance(self.AsAnimationRootNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeAnimation", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeAnimation{*(*gdclass.AnimationNodeAnimation)(unsafe.Pointer(&ptr))}
	})
}

type PlayMode = gdclass.AnimationNodeAnimationPlayMode

const (
	/*Plays animation in forward direction.*/
	PlayModeForward PlayMode = 0
	/*Plays animation in backward direction.*/
	PlayModeBackward PlayMode = 1
)
