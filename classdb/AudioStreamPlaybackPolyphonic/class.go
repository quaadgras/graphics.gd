// Package AudioStreamPlaybackPolyphonic provides methods for working with AudioStreamPlaybackPolyphonic object instances.
package AudioStreamPlaybackPolyphonic

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioStreamPlayback"
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
Playback instance for [AudioStreamPolyphonic]. After setting the [code]stream[/code] property of [AudioStreamPlayer], [AudioStreamPlayer2D], or [AudioStreamPlayer3D], the playback instance can be obtained by calling [method AudioStreamPlayer.get_stream_playback], [method AudioStreamPlayer2D.get_stream_playback] or [method AudioStreamPlayer3D.get_stream_playback] methods.
*/
type Instance [1]gdclass.AudioStreamPlaybackPolyphonic

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlaybackPolyphonic() Instance
}

/*
Play an [AudioStream] at a given offset, volume, pitch scale, playback type, and bus. Playback starts immediately.
The return value is a unique integer ID that is associated to this playback stream and which can be used to control it.
This ID becomes invalid when the stream ends (if it does not loop), when the [AudioStreamPlaybackPolyphonic] is stopped, or when [method stop_stream] is called.
This function returns [constant INVALID_ID] if the amount of streams currently playing equals [member AudioStreamPolyphonic.polyphony]. If you need a higher amount of maximum polyphony, raise this value.
*/
func (self Instance) PlayStream(stream [1]gdclass.AudioStream) int { //gd:AudioStreamPlaybackPolyphonic.play_stream
	return int(int(class(self).PlayStream(stream, float64(0), float64(0), float64(1.0), 0, String.Name(String.New("Master")))))
}

/*
Change the stream volume (in db). The [param stream] argument is an integer ID returned by [method play_stream].
*/
func (self Instance) SetStreamVolume(stream int, volume_db Float.X) { //gd:AudioStreamPlaybackPolyphonic.set_stream_volume
	class(self).SetStreamVolume(int64(stream), float64(volume_db))
}

/*
Change the stream pitch scale. The [param stream] argument is an integer ID returned by [method play_stream].
*/
func (self Instance) SetStreamPitchScale(stream int, pitch_scale Float.X) { //gd:AudioStreamPlaybackPolyphonic.set_stream_pitch_scale
	class(self).SetStreamPitchScale(int64(stream), float64(pitch_scale))
}

/*
Returns [code]true[/code] if the stream associated with the given integer ID is still playing. Check [method play_stream] for information on when this ID becomes invalid.
*/
func (self Instance) IsStreamPlaying(stream int) bool { //gd:AudioStreamPlaybackPolyphonic.is_stream_playing
	return bool(class(self).IsStreamPlaying(int64(stream)))
}

/*
Stop a stream. The [param stream] argument is an integer ID returned by [method play_stream], which becomes invalid after calling this function.
*/
func (self Instance) StopStream(stream int) { //gd:AudioStreamPlaybackPolyphonic.stop_stream
	class(self).StopStream(int64(stream))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlaybackPolyphonic

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackPolyphonic"))
	casted := Instance{*(*gdclass.AudioStreamPlaybackPolyphonic)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Play an [AudioStream] at a given offset, volume, pitch scale, playback type, and bus. Playback starts immediately.
The return value is a unique integer ID that is associated to this playback stream and which can be used to control it.
This ID becomes invalid when the stream ends (if it does not loop), when the [AudioStreamPlaybackPolyphonic] is stopped, or when [method stop_stream] is called.
This function returns [constant INVALID_ID] if the amount of streams currently playing equals [member AudioStreamPolyphonic.polyphony]. If you need a higher amount of maximum polyphony, raise this value.
*/
//go:nosplit
func (self class) PlayStream(stream [1]gdclass.AudioStream, from_offset float64, volume_db float64, pitch_scale float64, playback_type gdclass.AudioServerPlaybackType, bus String.Name) int64 { //gd:AudioStreamPlaybackPolyphonic.play_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, from_offset)
	callframe.Arg(frame, volume_db)
	callframe.Arg(frame, pitch_scale)
	callframe.Arg(frame, playback_type)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(bus)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackPolyphonic.Bind_play_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Change the stream volume (in db). The [param stream] argument is an integer ID returned by [method play_stream].
*/
//go:nosplit
func (self class) SetStreamVolume(stream int64, volume_db float64) { //gd:AudioStreamPlaybackPolyphonic.set_stream_volume
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	callframe.Arg(frame, volume_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackPolyphonic.Bind_set_stream_volume, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Change the stream pitch scale. The [param stream] argument is an integer ID returned by [method play_stream].
*/
//go:nosplit
func (self class) SetStreamPitchScale(stream int64, pitch_scale float64) { //gd:AudioStreamPlaybackPolyphonic.set_stream_pitch_scale
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	callframe.Arg(frame, pitch_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackPolyphonic.Bind_set_stream_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the stream associated with the given integer ID is still playing. Check [method play_stream] for information on when this ID becomes invalid.
*/
//go:nosplit
func (self class) IsStreamPlaying(stream int64) bool { //gd:AudioStreamPlaybackPolyphonic.is_stream_playing
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackPolyphonic.Bind_is_stream_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Stop a stream. The [param stream] argument is an integer ID returned by [method play_stream], which becomes invalid after calling this function.
*/
//go:nosplit
func (self class) StopStream(stream int64) { //gd:AudioStreamPlaybackPolyphonic.stop_stream
	var frame = callframe.New()
	callframe.Arg(frame, stream)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackPolyphonic.Bind_stop_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsAudioStreamPlaybackPolyphonic() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlaybackPolyphonic() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Advanced {
	return *((*AudioStreamPlayback.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlayback() AudioStreamPlayback.Instance {
	return *((*AudioStreamPlayback.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioStreamPlayback.Advanced(self.AsAudioStreamPlayback()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStreamPlayback.Instance(self.AsAudioStreamPlayback()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlaybackPolyphonic", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlaybackPolyphonic{*(*gdclass.AudioStreamPlaybackPolyphonic)(unsafe.Pointer(&ptr))}
	})
}
