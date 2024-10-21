package AudioStreamPlayer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [AudioStreamPlayer] node plays an audio stream non-positionally. It is ideal for user interfaces, menus, or background music.
To use this node, [member stream] needs to be set to a valid [AudioStream] resource. Playing more than one sound at the same time is also supported, see [member max_polyphony].
If you need to play audio at a specific position, use [AudioStreamPlayer2D] or [AudioStreamPlayer3D] instead.

*/
type Simple [1]classdb.AudioStreamPlayer
func (self Simple) SetStream(stream [1]classdb.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStream(stream)
}
func (self Simple) GetStream() [1]classdb.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStream(Expert(self).GetStream(gc))
}
func (self Simple) SetVolumeDb(volume_db float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVolumeDb(gd.Float(volume_db))
}
func (self Simple) GetVolumeDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVolumeDb()))
}
func (self Simple) SetPitchScale(pitch_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPitchScale(gd.Float(pitch_scale))
}
func (self Simple) GetPitchScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPitchScale()))
}
func (self Simple) Play(from_position float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Play(gd.Float(from_position))
}
func (self Simple) SeekTo(to_position float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SeekTo(gd.Float(to_position))
}
func (self Simple) Stop() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Stop()
}
func (self Simple) IsPlaying() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPlaying())
}
func (self Simple) GetPlaybackPosition() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPlaybackPosition()))
}
func (self Simple) SetBus(bus string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBus(gc.StringName(bus))
}
func (self Simple) GetBus() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetBus(gc).String())
}
func (self Simple) SetAutoplay(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoplay(enable)
}
func (self Simple) IsAutoplayEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAutoplayEnabled())
}
func (self Simple) SetMixTarget(mix_target classdb.AudioStreamPlayerMixTarget) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMixTarget(mix_target)
}
func (self Simple) GetMixTarget() classdb.AudioStreamPlayerMixTarget {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AudioStreamPlayerMixTarget(Expert(self).GetMixTarget())
}
func (self Simple) SetStreamPaused(pause bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamPaused(pause)
}
func (self Simple) GetStreamPaused() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetStreamPaused())
}
func (self Simple) SetMaxPolyphony(max_polyphony int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxPolyphony(gd.Int(max_polyphony))
}
func (self Simple) GetMaxPolyphony() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxPolyphony()))
}
func (self Simple) HasStreamPlayback() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasStreamPlayback())
}
func (self Simple) GetStreamPlayback() [1]classdb.AudioStreamPlayback {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStreamPlayback(Expert(self).GetStreamPlayback(gc))
}
func (self Simple) SetPlaybackType(playback_type classdb.AudioServerPlaybackType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlaybackType(playback_type)
}
func (self Simple) GetPlaybackType() classdb.AudioServerPlaybackType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AudioServerPlaybackType(Expert(self).GetPlaybackType())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamPlayer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetStream(stream object.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStream(ctx gd.Lifetime) object.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStream
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
func (self class) GetStreamPlayback(ctx gd.Lifetime) object.AudioStreamPlayback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayer.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStreamPlayback
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

//go:nosplit
func (self class) AsAudioStreamPlayer() Expert { return self[0].AsAudioStreamPlayer() }


//go:nosplit
func (self Simple) AsAudioStreamPlayer() Simple { return self[0].AsAudioStreamPlayer() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStreamPlayer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type MixTarget = classdb.AudioStreamPlayerMixTarget

const (
/*The audio will be played only on the first channel. This is the default.*/
	MixTargetStereo MixTarget = 0
/*The audio will be played on all surround channels.*/
	MixTargetSurround MixTarget = 1
/*The audio will be played on the second channel, which is usually the center.*/
	MixTargetCenter MixTarget = 2
)
