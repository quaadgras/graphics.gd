package AnimationPlayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationMixer"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An animation player is used for general-purpose playback of animations. It contains a dictionary of [AnimationLibrary] resources and custom blend times between animation transitions.
Some methods and properties use a single key to reference an animation directly. These keys are formatted as the key for the library, followed by a forward slash, then the key for the animation within the library, for example [code]"movement/run"[/code]. If the library's key is an empty string (known as the default library), the forward slash is omitted, being the same key used by the library.
[AnimationPlayer] is better-suited than [Tween] for more complex animations, for example ones with non-trivial timings. It can also be used over [Tween] if the animation track editor is more convenient than doing it in code.
Updating the target properties of animations occurs at the process frame.

*/
type Go [1]classdb.AnimationPlayer

/*
Triggers the [param animation_to] animation when the [param animation_from] animation completes.
*/
func (self Go) AnimationSetNext(animation_from string, animation_to string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AnimationSetNext(gc.StringName(animation_from), gc.StringName(animation_to))
}

/*
Returns the key of the animation which is queued to play after the [param animation_from] animation.
*/
func (self Go) AnimationGetNext(animation_from string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).AnimationGetNext(gc, gc.StringName(animation_from)).String())
}

/*
Specifies a blend time (in seconds) between two animations, referenced by their keys.
*/
func (self Go) SetBlendTime(animation_from string, animation_to string, sec float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlendTime(gc.StringName(animation_from), gc.StringName(animation_to), gd.Float(sec))
}

/*
Returns the blend time (in seconds) between two animations, referenced by their keys.
*/
func (self Go) GetBlendTime(animation_from string, animation_to string) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetBlendTime(gc.StringName(animation_from), gc.StringName(animation_to))))
}

/*
Plays the animation with key [param name]. Custom blend times and speed can be set.
The [param from_end] option only affects when switching to a new animation track, or if the same track but at the start or end. It does not affect resuming playback that was paused in the middle of an animation. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
The [AnimationPlayer] keeps track of its current or last played animation with [member assigned_animation]. If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
[b]Note:[/b] The animation will be updated the next time the [AnimationPlayer] is processed. If other variables are updated at the same time this is called, they may be updated too early. To perform the update immediately, call [code]advance(0)[/code].
*/
func (self Go) Play() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Play(gc.StringName(""), gd.Float(-1), gd.Float(1.0), false)
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
func (self Go) PlayBackwards() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PlayBackwards(gc.StringName(""), gd.Float(-1))
}

/*
See also [method AnimationMixer.capture].
You can use this method to use more detailed options for capture than those performed by [member playback_auto_capture]. When [member playback_auto_capture] is [code]false[/code], this method is almost the same as the following:
[codeblock]
capture(name, duration, trans_type, ease_type)
play(name, custom_blend, custom_speed, from_end)
[/codeblock]
If [param name] is blank, it specifies [member assigned_animation].
If [param duration] is a negative value, the duration is set to the interval between the current position and the first key, when [param from_end] is [code]true[/code], uses the interval between the current position and the last key instead.
[b]Note:[/b] The [param duration] takes [member speed_scale] into account, but [param custom_speed] does not, because the capture cache is interpolated with the blend result and the result may contain multiple animations.
*/
func (self Go) PlayWithCapture() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PlayWithCapture(gc.StringName(""), gd.Float(-1.0), gd.Float(-1), gd.Float(1.0), false, 0, 0)
}

/*
Pauses the currently playing animation. The [member current_animation_position] will be kept and calling [method play] or [method play_backwards] without arguments or with the same animation name as [member assigned_animation] will resume the animation.
See also [method stop].
*/
func (self Go) Pause() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Pause()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
If [param keep_state] is [code]true[/code], the animation state is not updated visually.
[b]Note:[/b] The method / audio / animation playback tracks will not be processed by this method.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop(false)
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
func (self Go) IsPlaying() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsPlaying())
}

/*
Queues an animation for playback once the current animation and all previously queued animations are done.
[b]Note:[/b] If a looped animation is currently playing, the queued animation will never play unless the looped animation is stopped somehow.
*/
func (self Go) Queue(name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Queue(gc.StringName(name))
}

/*
Returns a list of the animation keys that are currently queued to play.
*/
func (self Go) GetQueue() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetQueue(gc).Strings(gc))
}

/*
Clears all queued, unplayed animations.
*/
func (self Go) ClearQueue() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearQueue()
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
func (self Go) GetPlayingSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetPlayingSpeed()))
}

/*
Seeks the animation to the [param seconds] point in time (in seconds). If [param update] is [code]true[/code], the animation updates too, otherwise it updates at process time. Events between the current frame and [param seconds] are skipped.
If [param update_only] is [code]true[/code], the method / audio / animation playback tracks will not be processed.
[b]Note:[/b] Seeking to the end of the animation doesn't emit [signal AnimationMixer.animation_finished]. If you want to skip animation and emit the signal, use [method AnimationMixer.advance].
*/
func (self Go) SeekTo(seconds float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SeekTo(gd.Float(seconds), false, false)
}

/*
Sets the process notification in which to update animations.
*/
func (self Go) SetProcessCallback(mode classdb.AnimationPlayerAnimationProcessCallback) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProcessCallback(mode)
}

/*
Returns the process notification in which to update animations.
*/
func (self Go) GetProcessCallback() classdb.AnimationPlayerAnimationProcessCallback {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationPlayerAnimationProcessCallback(class(self).GetProcessCallback())
}

/*
Sets the call mode used for "Call Method" tracks.
*/
func (self Go) SetMethodCallMode(mode classdb.AnimationPlayerAnimationMethodCallMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMethodCallMode(mode)
}

/*
Returns the call mode used for "Call Method" tracks.
*/
func (self Go) GetMethodCallMode() classdb.AnimationPlayerAnimationMethodCallMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationPlayerAnimationMethodCallMode(class(self).GetMethodCallMode())
}

/*
Sets the node which node path references will travel from.
*/
func (self Go) SetRoot(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRoot(gc.String(path).NodePath(gc))
}

/*
Returns the node which node path references will travel from.
*/
func (self Go) GetRoot() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetRoot(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationPlayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AnimationPlayer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) CurrentAnimation() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetCurrentAnimation(gc).String())
}

func (self Go) SetCurrentAnimation(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurrentAnimation(gc.String(value))
}

func (self Go) AssignedAnimation() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAssignedAnimation(gc).String())
}

func (self Go) SetAssignedAnimation(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAssignedAnimation(gc.String(value))
}

func (self Go) Autoplay() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetAutoplay(gc).String())
}

func (self Go) SetAutoplay(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoplay(gc.String(value))
}

func (self Go) CurrentAnimationLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCurrentAnimationLength()))
}

func (self Go) CurrentAnimationPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCurrentAnimationPosition()))
}

func (self Go) PlaybackAutoCapture() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoCapture())
}

func (self Go) SetPlaybackAutoCapture(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoCapture(value)
}

func (self Go) PlaybackAutoCaptureDuration() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAutoCaptureDuration()))
}

func (self Go) SetPlaybackAutoCaptureDuration(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoCaptureDuration(gd.Float(value))
}

func (self Go) PlaybackAutoCaptureTransitionType() classdb.TweenTransitionType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TweenTransitionType(class(self).GetAutoCaptureTransitionType())
}

func (self Go) SetPlaybackAutoCaptureTransitionType(value classdb.TweenTransitionType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoCaptureTransitionType(value)
}

func (self Go) PlaybackAutoCaptureEaseType() classdb.TweenEaseType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TweenEaseType(class(self).GetAutoCaptureEaseType())
}

func (self Go) SetPlaybackAutoCaptureEaseType(value classdb.TweenEaseType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoCaptureEaseType(value)
}

func (self Go) PlaybackDefaultBlendTime() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDefaultBlendTime()))
}

func (self Go) SetPlaybackDefaultBlendTime(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDefaultBlendTime(gd.Float(value))
}

func (self Go) SpeedScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSpeedScale()))
}

func (self Go) SetSpeedScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpeedScale(gd.Float(value))
}

func (self Go) MovieQuitOnFinish() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMovieQuitOnFinishEnabled())
}

func (self Go) SetMovieQuitOnFinish(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMovieQuitOnFinishEnabled(value)
}

/*
Triggers the [param animation_to] animation when the [param animation_from] animation completes.
*/
//go:nosplit
func (self class) AnimationSetNext(animation_from gd.StringName, animation_to gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation_from))
	callframe.Arg(frame, mmm.Get(animation_to))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_animation_set_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the key of the animation which is queued to play after the [param animation_from] animation.
*/
//go:nosplit
func (self class) AnimationGetNext(ctx gd.Lifetime, animation_from gd.StringName) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation_from))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_animation_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Specifies a blend time (in seconds) between two animations, referenced by their keys.
*/
//go:nosplit
func (self class) SetBlendTime(animation_from gd.StringName, animation_to gd.StringName, sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation_from))
	callframe.Arg(frame, mmm.Get(animation_to))
	callframe.Arg(frame, sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the blend time (in seconds) between two animations, referenced by their keys.
*/
//go:nosplit
func (self class) GetBlendTime(animation_from gd.StringName, animation_to gd.StringName) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation_from))
	callframe.Arg(frame, mmm.Get(animation_to))
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultBlendTime(sec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_default_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultBlendTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_default_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoCapture(auto_capture bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_auto_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoCapture() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_is_auto_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoCaptureDuration(auto_capture_duration gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture_duration)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_auto_capture_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoCaptureDuration() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_auto_capture_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoCaptureTransitionType(auto_capture_transition_type classdb.TweenTransitionType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture_transition_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_auto_capture_transition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoCaptureTransitionType() classdb.TweenTransitionType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TweenTransitionType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_auto_capture_transition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoCaptureEaseType(auto_capture_ease_type classdb.TweenEaseType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture_ease_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_auto_capture_ease_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoCaptureEaseType() classdb.TweenEaseType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TweenEaseType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_auto_capture_ease_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Plays the animation with key [param name]. Custom blend times and speed can be set.
The [param from_end] option only affects when switching to a new animation track, or if the same track but at the start or end. It does not affect resuming playback that was paused in the middle of an animation. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
The [AnimationPlayer] keeps track of its current or last played animation with [member assigned_animation]. If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
[b]Note:[/b] The animation will be updated the next time the [AnimationPlayer] is processed. If other variables are updated at the same time this is called, they may be updated too early. To perform the update immediately, call [code]advance(0)[/code].
*/
//go:nosplit
func (self class) Play(name gd.StringName, custom_blend gd.Float, custom_speed gd.Float, from_end bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, custom_blend)
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
//go:nosplit
func (self class) PlayBackwards(name gd.StringName, custom_blend gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, custom_blend)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_play_backwards, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
See also [method AnimationMixer.capture].
You can use this method to use more detailed options for capture than those performed by [member playback_auto_capture]. When [member playback_auto_capture] is [code]false[/code], this method is almost the same as the following:
[codeblock]
capture(name, duration, trans_type, ease_type)
play(name, custom_blend, custom_speed, from_end)
[/codeblock]
If [param name] is blank, it specifies [member assigned_animation].
If [param duration] is a negative value, the duration is set to the interval between the current position and the first key, when [param from_end] is [code]true[/code], uses the interval between the current position and the last key instead.
[b]Note:[/b] The [param duration] takes [member speed_scale] into account, but [param custom_speed] does not, because the capture cache is interpolated with the blend result and the result may contain multiple animations.
*/
//go:nosplit
func (self class) PlayWithCapture(name gd.StringName, duration gd.Float, custom_blend gd.Float, custom_speed gd.Float, from_end bool, trans_type classdb.TweenTransitionType, ease_type classdb.TweenEaseType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, duration)
	callframe.Arg(frame, custom_blend)
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	callframe.Arg(frame, trans_type)
	callframe.Arg(frame, ease_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_play_with_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pauses the currently playing animation. The [member current_animation_position] will be kept and calling [method play] or [method play_backwards] without arguments or with the same animation name as [member assigned_animation] will resume the animation.
See also [method stop].
*/
//go:nosplit
func (self class) Pause()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
If [param keep_state] is [code]true[/code], the animation state is not updated visually.
[b]Note:[/b] The method / audio / animation playback tracks will not be processed by this method.
*/
//go:nosplit
func (self class) Stop(keep_state bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keep_state)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurrentAnimation(animation gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_current_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurrentAnimation(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_current_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAssignedAnimation(animation gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_assigned_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAssignedAnimation(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_assigned_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Queues an animation for playback once the current animation and all previously queued animations are done.
[b]Note:[/b] If a looped animation is currently playing, the queued animation will never play unless the looped animation is stopped somehow.
*/
//go:nosplit
func (self class) Queue(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_queue, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a list of the animation keys that are currently queued to play.
*/
//go:nosplit
func (self class) GetQueue(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_queue, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears all queued, unplayed animations.
*/
//go:nosplit
func (self class) ClearQueue()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_clear_queue, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSpeedScale(speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
//go:nosplit
func (self class) GetPlayingSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_playing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoplay(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoplay(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMovieQuitOnFinishEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_movie_quit_on_finish_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMovieQuitOnFinishEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_is_movie_quit_on_finish_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentAnimationPosition() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_current_animation_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCurrentAnimationLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_current_animation_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Seeks the animation to the [param seconds] point in time (in seconds). If [param update] is [code]true[/code], the animation updates too, otherwise it updates at process time. Events between the current frame and [param seconds] are skipped.
If [param update_only] is [code]true[/code], the method / audio / animation playback tracks will not be processed.
[b]Note:[/b] Seeking to the end of the animation doesn't emit [signal AnimationMixer.animation_finished]. If you want to skip animation and emit the signal, use [method AnimationMixer.advance].
*/
//go:nosplit
func (self class) SeekTo(seconds gd.Float, update bool, update_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	callframe.Arg(frame, update)
	callframe.Arg(frame, update_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the process notification in which to update animations.
*/
//go:nosplit
func (self class) SetProcessCallback(mode classdb.AnimationPlayerAnimationProcessCallback)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the process notification in which to update animations.
*/
//go:nosplit
func (self class) GetProcessCallback() classdb.AnimationPlayerAnimationProcessCallback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationPlayerAnimationProcessCallback](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the call mode used for "Call Method" tracks.
*/
//go:nosplit
func (self class) SetMethodCallMode(mode classdb.AnimationPlayerAnimationMethodCallMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_method_call_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the call mode used for "Call Method" tracks.
*/
//go:nosplit
func (self class) GetMethodCallMode() classdb.AnimationPlayerAnimationMethodCallMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationPlayerAnimationMethodCallMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_method_call_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the node which node path references will travel from.
*/
//go:nosplit
func (self class) SetRoot(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_set_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the node which node path references will travel from.
*/
//go:nosplit
func (self class) GetRoot(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationPlayer.Bind_get_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnCurrentAnimationChanged(cb func(name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("current_animation_changed"), gc.Callable(cb), 0)
}


func (self Go) OnAnimationChanged(cb func(old_name string, new_name string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("animation_changed"), gc.Callable(cb), 0)
}


func (self class) AsAnimationPlayer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationPlayer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAnimationMixer() AnimationMixer.GD { return *((*AnimationMixer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationMixer() AnimationMixer.Go { return *((*AnimationMixer.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}
func init() {classdb.Register("AnimationPlayer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AnimationProcessCallback = classdb.AnimationPlayerAnimationProcessCallback

const (
	AnimationProcessPhysics AnimationProcessCallback = 0
	AnimationProcessIdle AnimationProcessCallback = 1
	AnimationProcessManual AnimationProcessCallback = 2
)
type AnimationMethodCallMode = classdb.AnimationPlayerAnimationMethodCallMode

const (
	AnimationMethodCallDeferred AnimationMethodCallMode = 0
	AnimationMethodCallImmediate AnimationMethodCallMode = 1
)
