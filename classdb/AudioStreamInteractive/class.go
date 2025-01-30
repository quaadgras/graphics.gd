// Package AudioStreamInteractive provides methods for working with AudioStreamInteractive object instances.
package AudioStreamInteractive

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This is an audio stream that can playback music interactively, combining clips and a transition table. Clips must be added first, and the transition rules via the [method add_transition]. Additionally, this stream export a property parameter to control the playback via [AudioStreamPlayer], [AudioStreamPlayer2D], or [AudioStreamPlayer3D].
The way this is used is by filling a number of clips, then configuring the transition table. From there, clips are selected for playback and the music will smoothly go from the current to the new one while using the corresponding transition rule defined in the transition table.
*/
type Instance [1]gdclass.AudioStreamInteractive

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamInteractive() Instance
}

/*
Set the name of the current clip (for easier identification).
*/
func (self Instance) SetClipName(clip_index int, name string) { //gd:AudioStreamInteractive.set_clip_name
	class(self).SetClipName(int64(clip_index), String.Name(String.New(name)))
}

/*
Return the name of a clip.
*/
func (self Instance) GetClipName(clip_index int) string { //gd:AudioStreamInteractive.get_clip_name
	return string(class(self).GetClipName(int64(clip_index)).String())
}

/*
Set the [AudioStream] associated with the current clip.
*/
func (self Instance) SetClipStream(clip_index int, stream [1]gdclass.AudioStream) { //gd:AudioStreamInteractive.set_clip_stream
	class(self).SetClipStream(int64(clip_index), stream)
}

/*
Return the [AudioStream] associated with a clip.
*/
func (self Instance) GetClipStream(clip_index int) [1]gdclass.AudioStream { //gd:AudioStreamInteractive.get_clip_stream
	return [1]gdclass.AudioStream(class(self).GetClipStream(int64(clip_index)))
}

/*
Set whether a clip will auto-advance by changing the auto-advance mode.
*/
func (self Instance) SetClipAutoAdvance(clip_index int, mode gdclass.AudioStreamInteractiveAutoAdvanceMode) { //gd:AudioStreamInteractive.set_clip_auto_advance
	class(self).SetClipAutoAdvance(int64(clip_index), mode)
}

/*
Return whether a clip has auto-advance enabled. See [method set_clip_auto_advance].
*/
func (self Instance) GetClipAutoAdvance(clip_index int) gdclass.AudioStreamInteractiveAutoAdvanceMode { //gd:AudioStreamInteractive.get_clip_auto_advance
	return gdclass.AudioStreamInteractiveAutoAdvanceMode(class(self).GetClipAutoAdvance(int64(clip_index)))
}

/*
Set the index of the next clip towards which this clip will auto advance to when finished. If the clip being played loops, then auto-advance will be ignored.
*/
func (self Instance) SetClipAutoAdvanceNextClip(clip_index int, auto_advance_next_clip int) { //gd:AudioStreamInteractive.set_clip_auto_advance_next_clip
	class(self).SetClipAutoAdvanceNextClip(int64(clip_index), int64(auto_advance_next_clip))
}

/*
Return the clip towards which the clip referenced by [param clip_index] will auto-advance to.
*/
func (self Instance) GetClipAutoAdvanceNextClip(clip_index int) int { //gd:AudioStreamInteractive.get_clip_auto_advance_next_clip
	return int(int(class(self).GetClipAutoAdvanceNextClip(int64(clip_index))))
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
func (self Instance) AddTransition(from_clip int, to_clip int, from_time gdclass.AudioStreamInteractiveTransitionFromTime, to_time gdclass.AudioStreamInteractiveTransitionToTime, fade_mode gdclass.AudioStreamInteractiveFadeMode, fade_beats Float.X) { //gd:AudioStreamInteractive.add_transition
	class(self).AddTransition(int64(from_clip), int64(to_clip), from_time, to_time, fade_mode, float64(fade_beats), false, int64(-1), false)
}

/*
Return true if a given transition exists (was added via [method add_transition]).
*/
func (self Instance) HasTransition(from_clip int, to_clip int) bool { //gd:AudioStreamInteractive.has_transition
	return bool(class(self).HasTransition(int64(from_clip), int64(to_clip)))
}

/*
Erase a transition by providing [param from_clip] and [param to_clip] clip indices. [constant CLIP_ANY] can be used for either argument or both.
*/
func (self Instance) EraseTransition(from_clip int, to_clip int) { //gd:AudioStreamInteractive.erase_transition
	class(self).EraseTransition(int64(from_clip), int64(to_clip))
}

/*
Return the list of transitions (from, to interleaved).
*/
func (self Instance) GetTransitionList() []int32 { //gd:AudioStreamInteractive.get_transition_list
	return []int32(slices.Collect(class(self).GetTransitionList().Values()))
}

/*
Return the source time position for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFromTime(from_clip int, to_clip int) gdclass.AudioStreamInteractiveTransitionFromTime { //gd:AudioStreamInteractive.get_transition_from_time
	return gdclass.AudioStreamInteractiveTransitionFromTime(class(self).GetTransitionFromTime(int64(from_clip), int64(to_clip)))
}

/*
Return the destination time position for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionToTime(from_clip int, to_clip int) gdclass.AudioStreamInteractiveTransitionToTime { //gd:AudioStreamInteractive.get_transition_to_time
	return gdclass.AudioStreamInteractiveTransitionToTime(class(self).GetTransitionToTime(int64(from_clip), int64(to_clip)))
}

/*
Return the mode for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFadeMode(from_clip int, to_clip int) gdclass.AudioStreamInteractiveFadeMode { //gd:AudioStreamInteractive.get_transition_fade_mode
	return gdclass.AudioStreamInteractiveFadeMode(class(self).GetTransitionFadeMode(int64(from_clip), int64(to_clip)))
}

/*
Return the time (in beats) for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFadeBeats(from_clip int, to_clip int) Float.X { //gd:AudioStreamInteractive.get_transition_fade_beats
	return Float.X(Float.X(class(self).GetTransitionFadeBeats(int64(from_clip), int64(to_clip))))
}

/*
Return whether a transition uses the [i]filler clip[/i] functionality (see [method add_transition]).
*/
func (self Instance) IsTransitionUsingFillerClip(from_clip int, to_clip int) bool { //gd:AudioStreamInteractive.is_transition_using_filler_clip
	return bool(class(self).IsTransitionUsingFillerClip(int64(from_clip), int64(to_clip)))
}

/*
Return the filler clip for a transition (see [method add_transition]).
*/
func (self Instance) GetTransitionFillerClip(from_clip int, to_clip int) int { //gd:AudioStreamInteractive.get_transition_filler_clip
	return int(int(class(self).GetTransitionFillerClip(int64(from_clip), int64(to_clip))))
}

/*
Return whether a transition uses the [i]hold previous[/i] functionality (see [method add_transition]).
*/
func (self Instance) IsTransitionHoldingPrevious(from_clip int, to_clip int) bool { //gd:AudioStreamInteractive.is_transition_holding_previous
	return bool(class(self).IsTransitionHoldingPrevious(int64(from_clip), int64(to_clip)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamInteractive

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamInteractive"))
	casted := Instance{*(*gdclass.AudioStreamInteractive)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) InitialClip() int {
	return int(int(class(self).GetInitialClip()))
}

func (self Instance) SetInitialClip(value int) {
	class(self).SetInitialClip(int64(value))
}

func (self Instance) ClipCount() int {
	return int(int(class(self).GetClipCount()))
}

func (self Instance) SetClipCount(value int) {
	class(self).SetClipCount(int64(value))
}

//go:nosplit
func (self class) SetClipCount(clip_count int64) { //gd:AudioStreamInteractive.set_clip_count
	var frame = callframe.New()
	callframe.Arg(frame, clip_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipCount() int64 { //gd:AudioStreamInteractive.get_clip_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInitialClip(clip_index int64) { //gd:AudioStreamInteractive.set_initial_clip
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_initial_clip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInitialClip() int64 { //gd:AudioStreamInteractive.get_initial_clip
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_initial_clip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the name of the current clip (for easier identification).
*/
//go:nosplit
func (self class) SetClipName(clip_index int64, name String.Name) { //gd:AudioStreamInteractive.set_clip_name
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return the name of a clip.
*/
//go:nosplit
func (self class) GetClipName(clip_index int64) String.Name { //gd:AudioStreamInteractive.get_clip_name
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Set the [AudioStream] associated with the current clip.
*/
//go:nosplit
func (self class) SetClipStream(clip_index int64, stream [1]gdclass.AudioStream) { //gd:AudioStreamInteractive.set_clip_stream
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return the [AudioStream] associated with a clip.
*/
//go:nosplit
func (self class) GetClipStream(clip_index int64) [1]gdclass.AudioStream { //gd:AudioStreamInteractive.get_clip_stream
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Set whether a clip will auto-advance by changing the auto-advance mode.
*/
//go:nosplit
func (self class) SetClipAutoAdvance(clip_index int64, mode gdclass.AudioStreamInteractiveAutoAdvanceMode) { //gd:AudioStreamInteractive.set_clip_auto_advance
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_auto_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return whether a clip has auto-advance enabled. See [method set_clip_auto_advance].
*/
//go:nosplit
func (self class) GetClipAutoAdvance(clip_index int64) gdclass.AudioStreamInteractiveAutoAdvanceMode { //gd:AudioStreamInteractive.get_clip_auto_advance
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[gdclass.AudioStreamInteractiveAutoAdvanceMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_auto_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the index of the next clip towards which this clip will auto advance to when finished. If the clip being played loops, then auto-advance will be ignored.
*/
//go:nosplit
func (self class) SetClipAutoAdvanceNextClip(clip_index int64, auto_advance_next_clip int64) { //gd:AudioStreamInteractive.set_clip_auto_advance_next_clip
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	callframe.Arg(frame, auto_advance_next_clip)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_set_clip_auto_advance_next_clip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return the clip towards which the clip referenced by [param clip_index] will auto-advance to.
*/
//go:nosplit
func (self class) GetClipAutoAdvanceNextClip(clip_index int64) int64 { //gd:AudioStreamInteractive.get_clip_auto_advance_next_clip
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_clip_auto_advance_next_clip, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AddTransition(from_clip int64, to_clip int64, from_time gdclass.AudioStreamInteractiveTransitionFromTime, to_time gdclass.AudioStreamInteractiveTransitionToTime, fade_mode gdclass.AudioStreamInteractiveFadeMode, fade_beats float64, use_filler_clip bool, filler_clip int64, hold_previous bool) { //gd:AudioStreamInteractive.add_transition
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_add_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return true if a given transition exists (was added via [method add_transition]).
*/
//go:nosplit
func (self class) HasTransition(from_clip int64, to_clip int64) bool { //gd:AudioStreamInteractive.has_transition
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_has_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Erase a transition by providing [param from_clip] and [param to_clip] clip indices. [constant CLIP_ANY] can be used for either argument or both.
*/
//go:nosplit
func (self class) EraseTransition(from_clip int64, to_clip int64) { //gd:AudioStreamInteractive.erase_transition
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_erase_transition, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return the list of transitions (from, to interleaved).
*/
//go:nosplit
func (self class) GetTransitionList() Packed.Array[int32] { //gd:AudioStreamInteractive.get_transition_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Return the source time position for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFromTime(from_clip int64, to_clip int64) gdclass.AudioStreamInteractiveTransitionFromTime { //gd:AudioStreamInteractive.get_transition_from_time
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[gdclass.AudioStreamInteractiveTransitionFromTime](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_from_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the destination time position for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionToTime(from_clip int64, to_clip int64) gdclass.AudioStreamInteractiveTransitionToTime { //gd:AudioStreamInteractive.get_transition_to_time
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[gdclass.AudioStreamInteractiveTransitionToTime](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_to_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the mode for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFadeMode(from_clip int64, to_clip int64) gdclass.AudioStreamInteractiveFadeMode { //gd:AudioStreamInteractive.get_transition_fade_mode
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[gdclass.AudioStreamInteractiveFadeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_fade_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the time (in beats) for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFadeBeats(from_clip int64, to_clip int64) float64 { //gd:AudioStreamInteractive.get_transition_fade_beats
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_fade_beats, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return whether a transition uses the [i]filler clip[/i] functionality (see [method add_transition]).
*/
//go:nosplit
func (self class) IsTransitionUsingFillerClip(from_clip int64, to_clip int64) bool { //gd:AudioStreamInteractive.is_transition_using_filler_clip
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_is_transition_using_filler_clip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return the filler clip for a transition (see [method add_transition]).
*/
//go:nosplit
func (self class) GetTransitionFillerClip(from_clip int64, to_clip int64) int64 { //gd:AudioStreamInteractive.get_transition_filler_clip
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_get_transition_filler_clip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Return whether a transition uses the [i]hold previous[/i] functionality (see [method add_transition]).
*/
//go:nosplit
func (self class) IsTransitionHoldingPrevious(from_clip int64, to_clip int64) bool { //gd:AudioStreamInteractive.is_transition_holding_previous
	var frame = callframe.New()
	callframe.Arg(frame, from_clip)
	callframe.Arg(frame, to_clip)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamInteractive.Bind_is_transition_holding_previous, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Advanced(self.AsAudioStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Instance(self.AsAudioStream()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamInteractive", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamInteractive{*(*gdclass.AudioStreamInteractive)(unsafe.Pointer(&ptr))}
	})
}

type TransitionFromTime = gdclass.AudioStreamInteractiveTransitionFromTime //gd:AudioStreamInteractive.TransitionFromTime

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

type TransitionToTime = gdclass.AudioStreamInteractiveTransitionToTime //gd:AudioStreamInteractive.TransitionToTime

const (
	/*Transition to the same position in the destination clip. This is useful when both clips have exactly the same length and the music should fade between them.*/
	TransitionToTimeSamePosition TransitionToTime = 0
	/*Transition to the start of the destination clip.*/
	TransitionToTimeStart TransitionToTime = 1
)

type FadeMode = gdclass.AudioStreamInteractiveFadeMode //gd:AudioStreamInteractive.FadeMode

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

type AutoAdvanceMode = gdclass.AudioStreamInteractiveAutoAdvanceMode //gd:AudioStreamInteractive.AutoAdvanceMode

const (
	/*Disable auto-advance (default).*/
	AutoAdvanceDisabled AutoAdvanceMode = 0
	/*Enable auto-advance, a clip must be specified.*/
	AutoAdvanceEnabled AutoAdvanceMode = 1
	/*Enable auto-advance, but instead of specifying a clip, the playback will return to hold (see [method add_transition]).*/
	AutoAdvanceReturnToHold AutoAdvanceMode = 2
)
