package AudioStreamPlaybackPolyphonic

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioStreamPlayback"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Playback instance for [AudioStreamPolyphonic]. After setting the [code]stream[/code] property of [AudioStreamPlayer], [AudioStreamPlayer2D], or [AudioStreamPlayer3D], the playback instance can be obtained by calling [method AudioStreamPlayer.get_stream_playback], [method AudioStreamPlayer2D.get_stream_playback] or [method AudioStreamPlayer3D.get_stream_playback] methods.

*/
type Simple [1]classdb.AudioStreamPlaybackPolyphonic
func (self Simple) PlayStream(stream [1]classdb.AudioStream, from_offset float64, volume_db float64, pitch_scale float64, playback_type classdb.AudioServerPlaybackType, bus string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).PlayStream(stream, gd.Float(from_offset), gd.Float(volume_db), gd.Float(pitch_scale), playback_type, gc.StringName(bus))))
}
func (self Simple) SetStreamVolume(stream int, volume_db float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamVolume(gd.Int(stream), gd.Float(volume_db))
}
func (self Simple) SetStreamPitchScale(stream int, pitch_scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamPitchScale(gd.Int(stream), gd.Float(pitch_scale))
}
func (self Simple) IsStreamPlaying(stream int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsStreamPlaying(gd.Int(stream)))
}
func (self Simple) StopStream(stream int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).StopStream(gd.Int(stream))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamPlaybackPolyphonic
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Play an [AudioStream] at a given offset, volume, pitch scale, playback type, and bus. Playback starts immediately.
The return value is a unique integer ID that is associated to this playback stream and which can be used to control it.
This ID becomes invalid when the stream ends (if it does not loop), when the [AudioStreamPlaybackPolyphonic] is stopped, or when [method stop_stream] is called.
This function returns [constant INVALID_ID] if the amount of streams currently playing equals [member AudioStreamPolyphonic.polyphony]. If you need a higher amount of maximum polyphony, raise this value.
*/
//go:nosplit
func (self class) PlayStream(stream object.AudioStream, from_offset gd.Float, volume_db gd.Float, pitch_scale gd.Float, playback_type classdb.AudioServerPlaybackType, bus gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	callframe.Arg(frame, from_offset)
	callframe.Arg(frame, volume_db)
	callframe.Arg(frame, pitch_scale)
	callframe.Arg(frame, playback_type)
	callframe.Arg(frame, mmm.Get(bus))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackPolyphonic.Bind_play_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Change the stream volume (in db). The [param stream] argument is an integer ID returned by [method play_stream].
*/
//go:nosplit
func (self class) SetStreamVolume(stream gd.Int, volume_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackPolyphonic.Bind_set_stream_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Change the stream pitch scale. The [param stream] argument is an integer ID returned by [method play_stream].
*/
//go:nosplit
func (self class) SetStreamPitchScale(stream gd.Int, pitch_scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	callframe.Arg(frame, pitch_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackPolyphonic.Bind_set_stream_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return true whether the stream associated with an integer ID is still playing. Check [method play_stream] for information on when this ID becomes invalid.
*/
//go:nosplit
func (self class) IsStreamPlaying(stream gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackPolyphonic.Bind_is_stream_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Stop a stream. The [param stream] argument is an integer ID returned by [method play_stream], which becomes invalid after calling this function.
*/
//go:nosplit
func (self class) StopStream(stream gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackPolyphonic.Bind_stop_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsAudioStreamPlaybackPolyphonic() Expert { return self[0].AsAudioStreamPlaybackPolyphonic() }


//go:nosplit
func (self Simple) AsAudioStreamPlaybackPolyphonic() Simple { return self[0].AsAudioStreamPlaybackPolyphonic() }


//go:nosplit
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Expert { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self Simple) AsAudioStreamPlayback() AudioStreamPlayback.Simple { return self[0].AsAudioStreamPlayback() }


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
func init() {classdb.Register("AudioStreamPlaybackPolyphonic", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
