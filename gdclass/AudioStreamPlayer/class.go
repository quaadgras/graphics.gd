package AudioStreamPlayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [AudioStreamPlayer] node plays an audio stream non-positionally. It is ideal for user interfaces, menus, or background music.
To use this node, [member stream] needs to be set to a valid [AudioStream] resource. Playing more than one sound at the same time is also supported, see [member max_polyphony].
If you need to play audio at a specific position, use [AudioStreamPlayer2D] or [AudioStreamPlayer3D] instead.

*/
type Go [1]classdb.AudioStreamPlayer

/*
Plays a sound from the beginning, or the given [param from_position] in seconds.
*/
func (self Go) Play() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Play(gd.Float(0.0))
}

/*
Restarts all sounds to be played from the given [param to_position], in seconds. Does nothing if no sounds are playing.
*/
func (self Go) SeekTo(to_position float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SeekTo(gd.Float(to_position))
}

/*
Stops all sounds from this node.
*/
func (self Go) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Stop()
}

/*
Returns the position in the [AudioStream] of the latest sound, in seconds. Returns [code]0.0[/code] if no sounds are playing.
[b]Note:[/b] The position is not always accurate, as the [AudioServer] does not mix audio every processed frame. To get more accurate results, add [method AudioServer.get_time_since_last_mix] to the returned position.
*/
func (self Go) GetPlaybackPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetPlaybackPosition()))
}

/*
Returns [code]true[/code] if any sound is active, even if [member stream_paused] is set to [code]true[/code]. See also [member playing] and [method get_stream_playback].
*/
func (self Go) HasStreamPlayback() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasStreamPlayback())
}

/*
Returns the latest [AudioStreamPlayback] of this node, usually the most recently created by [method play]. If no sounds are playing, this method fails and returns an empty playback.
*/
func (self Go) GetStreamPlayback() gdclass.AudioStreamPlayback {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStreamPlayback(class(self).GetStreamPlayback(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamPlayer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Stream() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetStream(gc))
}

func (self Go) SetStream(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStream(value)
}

func (self Go) VolumeDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumeDb()))
}

func (self Go) SetVolumeDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumeDb(gd.Float(value))
}

func (self Go) PitchScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPitchScale()))
}

func (self Go) SetPitchScale(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPitchScale(gd.Float(value))
}

func (self Go) Playing() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPlaying())
}

func (self Go) Autoplay() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoplayEnabled())
}

func (self Go) SetAutoplay(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoplay(value)
}

func (self Go) StreamPaused() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetStreamPaused())
}

func (self Go) SetStreamPaused(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStreamPaused(value)
}

func (self Go) MixTarget() classdb.AudioStreamPlayerMixTarget {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioStreamPlayerMixTarget(class(self).GetMixTarget())
}

func (self Go) SetMixTarget(value classdb.AudioStreamPlayerMixTarget) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMixTarget(value)
}

func (self Go) MaxPolyphony() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxPolyphony()))
}

func (self Go) SetMaxPolyphony(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxPolyphony(gd.Int(value))
}

func (self Go) Bus() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetBus(gc).String())
}

func (self Go) SetBus(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBus(gc.StringName(value))
}

func (self Go) PlaybackType() classdb.AudioServerPlaybackType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioServerPlaybackType(class(self).GetPlaybackType())
}

func (self Go) SetPlaybackType(value classdb.AudioServerPlaybackType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlaybackType(value)
}

//go:nosplit
func (self class) SetStream(stream gdclass.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStream(ctx gd.Lifetime) gdclass.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVolumeDb(volume_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumeDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPitchScale(pitch_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pitch_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPitchScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Plays a sound from the beginning, or the given [param from_position] in seconds.
*/
//go:nosplit
func (self class) Play(from_position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Restarts all sounds to be played from the given [param to_position], in seconds. Does nothing if no sounds are playing.
*/
//go:nosplit
func (self class) SeekTo(to_position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops all sounds from this node.
*/
//go:nosplit
func (self class) Stop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPlaying() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position in the [AudioStream] of the latest sound, in seconds. Returns [code]0.0[/code] if no sounds are playing.
[b]Note:[/b] The position is not always accurate, as the [AudioServer] does not mix audio every processed frame. To get more accurate results, add [method AudioServer.get_time_since_last_mix] to the returned position.
*/
//go:nosplit
func (self class) GetPlaybackPosition() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_playback_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBus(bus gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bus))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBus(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoplay(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoplayEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_is_autoplay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMixTarget(mix_target classdb.AudioStreamPlayerMixTarget)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mix_target)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_mix_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMixTarget() classdb.AudioStreamPlayerMixTarget {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamPlayerMixTarget](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_mix_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStreamPaused(pause bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_stream_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamPaused() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_stream_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxPolyphony(max_polyphony gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxPolyphony() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if any sound is active, even if [member stream_paused] is set to [code]true[/code]. See also [member playing] and [method get_stream_playback].
*/
//go:nosplit
func (self class) HasStreamPlayback() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_has_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the latest [AudioStreamPlayback] of this node, usually the most recently created by [method play]. If no sounds are playing, this method fails and returns an empty playback.
*/
//go:nosplit
func (self class) GetStreamPlayback(ctx gd.Lifetime) gdclass.AudioStreamPlayback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStreamPlayback
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaybackType(playback_type classdb.AudioServerPlaybackType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, playback_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_playback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaybackType() classdb.AudioServerPlaybackType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioServerPlaybackType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_playback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("finished"), gc.Callable(cb), 0)
}


func (self class) AsAudioStreamPlayer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlayer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("AudioStreamPlayer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type MixTarget = classdb.AudioStreamPlayerMixTarget

const (
/*The audio will be played only on the first channel. This is the default.*/
	MixTargetStereo MixTarget = 0
/*The audio will be played on all surround channels.*/
	MixTargetSurround MixTarget = 1
/*The audio will be played on the second channel, which is usually the center.*/
	MixTargetCenter MixTarget = 2
)
