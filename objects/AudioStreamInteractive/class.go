package AudioStreamInteractive

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioStream"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This is an audio stream that can playback music interactively, combining clips and a transition table. Clips must be added first, and the transition rules via the [method add_transition]. Additionally, this stream export a property parameter to control the playback via [AudioStreamPlayer], [AudioStreamPlayer2D], or [AudioStreamPlayer3D].
The way this is used is by filling a number of clips, then configuring the transition table. From there, clips are selected for playback and the music will smoothly go from the current to the new one while using the corresponding transition rule defined in the transition table.
*/
type Instance [1]classdb.AudioStreamInteractive

/*
Set the name of the current clip (for easier identification).
*/
func (self Instance) SetClipName(clip_index int, name string) {
	class(self).SetClipName(gd.Int(clip_index), gd.NewStringName(name))
}

/*
Return the name of a clip.
*/
func (self Instance) GetClipName(clip_index int) string {
	return string(class(self).GetClipName(gd.Int(clip_index)).String())
}

/*
Set the [AudioStream] associated with the current clip.
*/
func (self Instance) SetClipStream(clip_index int, stream objects.AudioStream) {
	class(self).SetClipStream(gd.Int(clip_index), stream)
}

/*
Return the [AudioStream] associated with a clip.
*/
func (self Instance) GetClipStream(clip_index int) objects.AudioStream {
	return objects.AudioStream(class(self).GetClipStream(gd.Int(clip_index)))
}

/*
Set whether a clip will auto-advance by changing the auto-advance mode.
*/
func (self Instance) SetClipAutoAdvance(clip_index int, mode classdb.AudioStreamInteractiveAutoAdvanceMode) {
	class(self).SetClipAutoAdvance(gd.Int(clip_index), mode)
}

/*
Return whether a clip has auto-advance enabled. See [method set_clip_auto_advance].
*/
func (self Instance) GetClipAutoAdvance(clip_index int) classdb.AudioStreamInteractiveAutoAdvanceMode {
	return classdb.AudioStreamInteractiveAutoAdvanceMode(class(self).GetClipAutoAdvance(gd.Int(clip_index)))
}

/*
Set the index of the next clip towards which this clip will auto advance to when finished. If the clip being played loops, then auto-advance will be ignored.
*/
func (self Instance) SetClipAutoAdvanceNextClip(clip_index int, auto_advance_next_clip int) {
	class(self).SetClipAutoAdvanceNextClip(gd.Int(clip_index), gd.Int(auto_advance_next_clip))
}

/*
Return the clip towards which the clip referenced by [param clip_index] will auto-advance to.
*/
func (self Instance) GetClipAutoAdvanceNextClip(clip_index int) int {
	return int(int(class(self).GetClipAutoAdvanceNextClip(gd.Int(clip_index))))
}

/*
Add a transition between two clips. Provide the indices of the source and destination clips, or use the [constant CLIP_ANY] constant to indicate that transition happens to/from any clip to this one.
* [param from_time] indicates the moment in the current clip the transition will begin after triggered.
* [param to_time] indicates the time in the next clip that the playback will start from.
* [param fade_mode] indicates how the fade will happen between clips. If unsure, just use [constant FADE_AUTOMATIC] which uses the most common type of fade for each situation.
* [param fade_beats] indicates how many beats the fade will take. Using decimals is allowed.
* [param use_filler_clip] indicates that there will be a filler clip used between the source and destination clips.
* [param filler_clip] the index of the filler clip.
* If [param hold_previous] is used, then this clip will be remembered. This can be used together with [constant AUTO_ADVANCE_RETURN_TO_HOLD] to return to this clip after another is done playing.
*/
func (self Instance) AddTransition(from_clip int, to_clip int, from_time classdb.AudioStreamInteractiveTransitionFromTime, to_time classdb.AudioStreamInteractiveTransitionToTime, fade_mode classdb.AudioStreamInteractiveFadeMode, fade_beats float64) {
	class(self).AddTransition(gd.Int(from_clip), gd.Int(to_clip), from_time, to_time, fade_mode, gd.Float(fade_beats), false, gd.Int(-1), false)
}

/*
Return true if a given transition exists (was added via [method add_transition]).
*/
func (self Instance) HasTransition(from_clip int, to_clip int) bool {
	return bool(class(self).HasTransition(gd.Int(from_clip), gd.Int(to_clip)))
}

/*
Erase a transition by providing [param from_clip] and [param to_clip] clip indices. [constant CLIP_ANY] can be used for either argument or both.
*/
func (self Instance) EraseTransition(from_clip int, to_clip int) {
	class(self).EraseTransition(gd.Int(from_clip), gd.Int(to_clip))
}

/*
Return the list of transitions (from, to interleaved).
*/
func (self Instance) GetTransitionList() []int32 {
	return []int32(class(self).GetTransitionList().AsSlice())
}

/*
Return the source time position for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFromTime(from_clip int, to_clip int) classdb.AudioStreamInteractiveTransitionFromTime {
	return classdb.AudioStreamInteractiveTransitionFromTime(class(self).GetTransitionFromTime(gd.Int(from_clip), gd.Int(to_clip)))
}

/*
Return the destination time position for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionToTime(from_clip int, to_clip int) classdb.AudioStreamInteractiveTransitionToTime {
	return classdb.AudioStreamInteractiveTransitionToTime(class(self).GetTransitionToTime(gd.Int(from_clip), gd.Int(to_clip)))
}

/*
Return the mode for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFadeMode(from_clip int, to_clip int) classdb.AudioStreamInteractiveFadeMode {
	return classdb.AudioStreamInteractiveFadeMode(class(self).GetTransitionFadeMode(gd.Int(from_clip), gd.Int(to_clip)))
}

/*
Return the time (in beats) for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFadeBeats(from_clip int, to_clip int) float64 {
	return float64(float64(class(self).GetTransitionFadeBeats(gd.Int(from_clip), gd.Int(to_clip))))
}

/*
Return whether a transition uses the [i]filler clip[/i] functionality (see [method add_transition]).
*/
func (self Instance) IsTransitionUsingFillerClip(from_clip int, to_clip int) bool {
	return bool(class(self).IsTransitionUsingFillerClip(gd.Int(from_clip), gd.Int(to_clip)))
}

/*
Return the filler clip for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFillerClip(from_clip int, to_clip int) int {
	return int(int(class(self).GetTransitionFillerClip(gd.Int(from_clip), gd.Int(to_clip))))
}

/*
Return whether a transition uses the [i]hold previous[/i] functionality (see [method add_transition]).
*/
func (self Instance) IsTransitionHoldingPrevious(from_clip int, to_clip int) bool {
	return bool(class(self).IsTransitionHoldingPrevious(gd.Int(from_clip), gd.Int(to_clip)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamInteractive

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamInteractive"))
	return Instance{classdb.AudioStreamInteractive(object)}
}

func (self Instance) InitialClip() int {
	return int(int(class(self).GetInitialClip()))
}

func (self Instance) SetInitialClip(value int) {
	class(self).SetInitialClip(gd.Int(value))
}

func (self Instance) ClipCount() int {
	return int(int(class(self).GetClipCount()))
}

func (self Instance) SetClipCount(value int) {
	class(self).SetClipCount(gd.Int(value))
}

//go:nosplit
func (self class) SetClipCount(clip_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInitialClip(clip_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_initial_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInitialClip() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_initial_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the name of the current clip (for easier identification).
*/
//go:nosplit
func (self class) SetClipName(clip_index gd.Int, name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return the name of a clip.
*/
//go:nosplit
func (self class) GetClipName(clip_index gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Set the [AudioStream] associated with the current clip.
*/
//go:nosplit
func (self class) SetClipStream(clip_index gd.Int, stream objects.AudioStream) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return the [AudioStream] associated with a clip.
*/
//go:nosplit
func (self class) GetClipStream(clip_index gd.Int) objects.AudioStream {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AudioStream{classdb.AudioStream(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Set whether a clip will auto-advance by changing the auto-advance mode.
*/
//go:nosplit
func (self class) SetClipAutoAdvance(clip_index gd.Int, mode classdb.AudioStreamInteractiveAutoAdvanceMode) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_auto_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return whether a clip has auto-advance enabled. See [method set_clip_auto_advance].
*/
//go:nosplit
func (self class) GetClipAutoAdvance(clip_index gd.Int) classdb.AudioStreamInteractiveAutoAdvanceMode {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[classdb.AudioStreamInteractiveAutoAdvanceMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_auto_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the index of the next clip towards which this clip will auto advance to when finished. If the clip being played loops, then auto-advance will be ignored.
*/
//go:nosplit
func (self class) SetClipAutoAdvanceNextClip(clip_index gd.Int, auto_advance_next_clip gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, auto_advance_next_clip)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_auto_advance_next_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return the clip towards which the clip referenced by [param clip_index] will auto-advance to.
*/
//go:nosplit
func (self class) GetClipAutoAdvanceNextClip(clip_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_auto_advance_next_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Add a transition between two clips. Provide the indices of the source and destination clips, or use the [constant CLIP_ANY] constant to indicate that transition happens to/from any clip to this one.
* [param from_time] indicates the moment in the current clip the transition will begin after triggered.
* [param to_time] indicates the time in the next clip that the playback will start from.
* [param fade_mode] indicates how the fade will happen between clips. If unsure, just use [constant FADE_AUTOMATIC] which uses the most common type of fade for each situation.
* [param fade_beats] indicates how many beats the fade will take. Using decimals is allowed.
* [param use_filler_clip] indicates that there will be a filler clip used between the source and destination clips.
* [param filler_clip] the index of the filler clip.
* If [param hold_previous] is used, then this clip will be remembered. This can be used together with [constant AUTO_ADVANCE_RETURN_TO_HOLD] to return to this clip after another is done playing.
*/
//go:nosplit
func (self class) AddTransition(from_clip gd.Int, to_clip gd.Int, from_time classdb.AudioStreamInteractiveTransitionFromTime, to_time classdb.AudioStreamInteractiveTransitionToTime, fade_mode classdb.AudioStreamInteractiveFadeMode, fade_beats gd.Float, use_filler_clip bool, filler_clip gd.Int, hold_previous bool) {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	callframe.Arg(frame, from_time)
	callframe.Arg(frame, to_time)
	callframe.Arg(frame, fade_mode)
	callframe.Arg(frame, fade_beats)
	callframe.Arg(frame, use_filler_clip)
	callframe.Arg(frame, filler_clip)
	callframe.Arg(frame, hold_previous)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_add_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return true if a given transition exists (was added via [method add_transition]).
*/
//go:nosplit
func (self class) HasTransition(from_clip gd.Int, to_clip gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_has_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Erase a transition by providing [param from_clip] and [param to_clip] clip indices. [constant CLIP_ANY] can be used for either argument or both.
*/
//go:nosplit
func (self class) EraseTransition(from_clip gd.Int, to_clip gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_erase_transition, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Return the list of transitions (from, to interleaved).
*/
//go:nosplit
func (self class) GetTransitionList() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Return the source time position for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFromTime(from_clip gd.Int, to_clip gd.Int) classdb.AudioStreamInteractiveTransitionFromTime {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[classdb.AudioStreamInteractiveTransitionFromTime](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_from_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the destination time position for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionToTime(from_clip gd.Int, to_clip gd.Int) classdb.AudioStreamInteractiveTransitionToTime {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[classdb.AudioStreamInteractiveTransitionToTime](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_to_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the mode for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFadeMode(from_clip gd.Int, to_clip gd.Int) classdb.AudioStreamInteractiveFadeMode {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[classdb.AudioStreamInteractiveFadeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the time (in beats) for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFadeBeats(from_clip gd.Int, to_clip gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_fade_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return whether a transition uses the [i]filler clip[/i] functionality (see [method add_transition]).
*/
//go:nosplit
func (self class) IsTransitionUsingFillerClip(from_clip gd.Int, to_clip gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_is_transition_using_filler_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the filler clip for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFillerClip(from_clip gd.Int, to_clip gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_filler_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return whether a transition uses the [i]hold previous[/i] functionality (see [method add_transition]).
*/
//go:nosplit
func (self class) IsTransitionHoldingPrevious(from_clip gd.Int, to_clip gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_is_transition_holding_previous, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamInteractive() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamInteractive() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStream() AudioStream.Advanced {
	return *((*AudioStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStream() AudioStream.Instance {
	return *((*AudioStream.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {
	classdb.Register("AudioStreamInteractive", func(ptr gd.Object) any { return classdb.AudioStreamInteractive(ptr) })
}

type TransitionFromTime = classdb.AudioStreamInteractiveTransitionFromTime

const (
	/*Start transition as soon as possible, don't wait for any specific time position.*/
	TransitionFromTimeImmediate TransitionFromTime = 0
	/*Transition when the clip playback position reaches the next beat.*/
	TransitionFromTimeNextBeat TransitionFromTime = 1
	/*Transition when the clip playback position reaches the next bar.*/
	TransitionFromTimeNextBar TransitionFromTime = 2
	/*Transition when the current clip finished playing.*/
	TransitionFromTimeEnd TransitionFromTime = 3
)

type TransitionToTime = classdb.AudioStreamInteractiveTransitionToTime

const (
	/*Transition to the same position in the destination clip. This is useful when both clips have exactly the same length and the music should fade between them.*/
	TransitionToTimeSamePosition TransitionToTime = 0
	/*Transition to the start of the destination clip.*/
	TransitionToTimeStart TransitionToTime = 1
)

type FadeMode = classdb.AudioStreamInteractiveFadeMode

const (
	/*Do not use fade for the transition. This is useful when transitioning from a clip-end to clip-beginning, and each clip has their begin/end.*/
	FadeDisabled FadeMode = 0
	/*Use a fade-in in the next clip, let the current clip finish.*/
	FadeIn FadeMode = 1
	/*Use a fade-out in the current clip, the next clip will start by itself.*/
	FadeOut FadeMode = 2
	/*Use a cross-fade between clips.*/
	FadeCross FadeMode = 3
	/*Use automatic fade logic depending on the transition from/to. It is recommended to use this by default.*/
	FadeAutomatic FadeMode = 4
)

type AutoAdvanceMode = classdb.AudioStreamInteractiveAutoAdvanceMode

const (
	/*Disable auto-advance (default).*/
	AutoAdvanceDisabled AutoAdvanceMode = 0
	/*Enable auto-advance, a clip must be specified.*/
	AutoAdvanceEnabled AutoAdvanceMode = 1
	/*Enable auto-advance, but instead of specifying a clip, the playback will return to hold (see [method add_transition]).*/
	AutoAdvanceReturnToHold AutoAdvanceMode = 2
)
