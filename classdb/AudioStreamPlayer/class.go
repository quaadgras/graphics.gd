// Package AudioStreamPlayer provides methods for working with AudioStreamPlayer object instances.
package AudioStreamPlayer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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

/*
The [AudioStreamPlayer] node plays an audio stream non-positionally. It is ideal for user interfaces, menus, or background music.
To use this node, [member stream] needs to be set to a valid [AudioStream] resource. Playing more than one sound at the same time is also supported, see [member max_polyphony].
If you need to play audio at a specific position, use [AudioStreamPlayer2D] or [AudioStreamPlayer3D] instead.
*/
type Instance [1]gdclass.AudioStreamPlayer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlayer() Instance
}

/*
Plays a sound from the beginning, or the given [param from_position] in seconds.
*/
func (self Instance) Play() { //gd:AudioStreamPlayer.play
	class(self).Play(gd.Float(0.0))
}

/*
Restarts all sounds to be played from the given [param to_position], in seconds. Does nothing if no sounds are playing.
*/
func (self Instance) SeekTo(to_position Float.X) { //gd:AudioStreamPlayer.seek
	class(self).SeekTo(gd.Float(to_position))
}

/*
Stops all sounds from this node.
*/
func (self Instance) Stop() { //gd:AudioStreamPlayer.stop
	class(self).Stop()
}

/*
Returns the position in the [AudioStream] of the latest sound, in seconds. Returns [code]0.0[/code] if no sounds are playing.
[b]Note:[/b] The position is not always accurate, as the [AudioServer] does not mix audio every processed frame. To get more accurate results, add [method AudioServer.get_time_since_last_mix] to the returned position.
*/
func (self Instance) GetPlaybackPosition() Float.X { //gd:AudioStreamPlayer.get_playback_position
	return Float.X(Float.X(class(self).GetPlaybackPosition()))
}

/*
Returns [code]true[/code] if any sound is active, even if [member stream_paused] is set to [code]true[/code]. See also [member playing] and [method get_stream_playback].
*/
func (self Instance) HasStreamPlayback() bool { //gd:AudioStreamPlayer.has_stream_playback
	return bool(class(self).HasStreamPlayback())
}

/*
Returns the latest [AudioStreamPlayback] of this node, usually the most recently created by [method play]. If no sounds are playing, this method fails and returns an empty playback.
*/
func (self Instance) GetStreamPlayback() [1]gdclass.AudioStreamPlayback { //gd:AudioStreamPlayer.get_stream_playback
	return [1]gdclass.AudioStreamPlayback(class(self).GetStreamPlayback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlayer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlayer"))
	casted := Instance{*(*gdclass.AudioStreamPlayer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Stream() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetStream())
}

func (self Instance) SetStream(value [1]gdclass.AudioStream) {
	class(self).SetStream(value)
}

func (self Instance) VolumeDb() Float.X {
	return Float.X(Float.X(class(self).GetVolumeDb()))
}

func (self Instance) SetVolumeDb(value Float.X) {
	class(self).SetVolumeDb(gd.Float(value))
}

func (self Instance) PitchScale() Float.X {
	return Float.X(Float.X(class(self).GetPitchScale()))
}

func (self Instance) SetPitchScale(value Float.X) {
	class(self).SetPitchScale(gd.Float(value))
}

func (self Instance) Playing() bool {
	return bool(class(self).IsPlaying())
}

func (self Instance) Autoplay() bool {
	return bool(class(self).IsAutoplayEnabled())
}

func (self Instance) SetAutoplay(value bool) {
	class(self).SetAutoplay(value)
}

func (self Instance) StreamPaused() bool {
	return bool(class(self).GetStreamPaused())
}

func (self Instance) SetStreamPaused(value bool) {
	class(self).SetStreamPaused(value)
}

func (self Instance) MixTarget() gdclass.AudioStreamPlayerMixTarget {
	return gdclass.AudioStreamPlayerMixTarget(class(self).GetMixTarget())
}

func (self Instance) SetMixTarget(value gdclass.AudioStreamPlayerMixTarget) {
	class(self).SetMixTarget(value)
}

func (self Instance) MaxPolyphony() int {
	return int(int(class(self).GetMaxPolyphony()))
}

func (self Instance) SetMaxPolyphony(value int) {
	class(self).SetMaxPolyphony(gd.Int(value))
}

func (self Instance) Bus() string {
	return string(class(self).GetBus().String())
}

func (self Instance) SetBus(value string) {
	class(self).SetBus(String.Name(String.New(value)))
}

func (self Instance) PlaybackType() gdclass.AudioServerPlaybackType {
	return gdclass.AudioServerPlaybackType(class(self).GetPlaybackType())
}

func (self Instance) SetPlaybackType(value gdclass.AudioServerPlaybackType) {
	class(self).SetPlaybackType(value)
}

//go:nosplit
func (self class) SetStream(stream [1]gdclass.AudioStream) { //gd:AudioStreamPlayer.set_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStream() [1]gdclass.AudioStream { //gd:AudioStreamPlayer.get_stream
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumeDb(volume_db gd.Float) { //gd:AudioStreamPlayer.set_volume_db
	var frame = callframe.New()
	callframe.Arg(frame, volume_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumeDb() gd.Float { //gd:AudioStreamPlayer.get_volume_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPitchScale(pitch_scale gd.Float) { //gd:AudioStreamPlayer.set_pitch_scale
	var frame = callframe.New()
	callframe.Arg(frame, pitch_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitchScale() gd.Float { //gd:AudioStreamPlayer.get_pitch_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Plays a sound from the beginning, or the given [param from_position] in seconds.
*/
//go:nosplit
func (self class) Play(from_position gd.Float) { //gd:AudioStreamPlayer.play
	var frame = callframe.New()
	callframe.Arg(frame, from_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Restarts all sounds to be played from the given [param to_position], in seconds. Does nothing if no sounds are playing.
*/
//go:nosplit
func (self class) SeekTo(to_position gd.Float) { //gd:AudioStreamPlayer.seek
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops all sounds from this node.
*/
//go:nosplit
func (self class) Stop() { //gd:AudioStreamPlayer.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPlaying() bool { //gd:AudioStreamPlayer.is_playing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position in the [AudioStream] of the latest sound, in seconds. Returns [code]0.0[/code] if no sounds are playing.
[b]Note:[/b] The position is not always accurate, as the [AudioServer] does not mix audio every processed frame. To get more accurate results, add [method AudioServer.get_time_since_last_mix] to the returned position.
*/
//go:nosplit
func (self class) GetPlaybackPosition() gd.Float { //gd:AudioStreamPlayer.get_playback_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_playback_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBus(bus String.Name) { //gd:AudioStreamPlayer.set_bus
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(bus)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBus() String.Name { //gd:AudioStreamPlayer.get_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(enable bool) { //gd:AudioStreamPlayer.set_autoplay
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoplayEnabled() bool { //gd:AudioStreamPlayer.is_autoplay_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_is_autoplay_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMixTarget(mix_target gdclass.AudioStreamPlayerMixTarget) { //gd:AudioStreamPlayer.set_mix_target
	var frame = callframe.New()
	callframe.Arg(frame, mix_target)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_mix_target, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMixTarget() gdclass.AudioStreamPlayerMixTarget { //gd:AudioStreamPlayer.get_mix_target
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamPlayerMixTarget](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_mix_target, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStreamPaused(pause bool) { //gd:AudioStreamPlayer.set_stream_paused
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_stream_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPaused() bool { //gd:AudioStreamPlayer.get_stream_paused
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_stream_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxPolyphony(max_polyphony gd.Int) { //gd:AudioStreamPlayer.set_max_polyphony
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxPolyphony() gd.Int { //gd:AudioStreamPlayer.get_max_polyphony
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if any sound is active, even if [member stream_paused] is set to [code]true[/code]. See also [member playing] and [method get_stream_playback].
*/
//go:nosplit
func (self class) HasStreamPlayback() bool { //gd:AudioStreamPlayer.has_stream_playback
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_has_stream_playback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the latest [AudioStreamPlayback] of this node, usually the most recently created by [method play]. If no sounds are playing, this method fails and returns an empty playback.
*/
//go:nosplit
func (self class) GetStreamPlayback() [1]gdclass.AudioStreamPlayback { //gd:AudioStreamPlayer.get_stream_playback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStreamPlayback{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamPlayback](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaybackType(playback_type gdclass.AudioServerPlaybackType) { //gd:AudioStreamPlayer.set_playback_type
	var frame = callframe.New()
	callframe.Arg(frame, playback_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_set_playback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaybackType() gdclass.AudioServerPlaybackType { //gd:AudioStreamPlayer.get_playback_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioServerPlaybackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer.Bind_get_playback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsAudioStreamPlayer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlayer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced            { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance         { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlayer", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlayer{*(*gdclass.AudioStreamPlayer)(unsafe.Pointer(&ptr))}
	})
}

type MixTarget = gdclass.AudioStreamPlayerMixTarget //gd:AudioStreamPlayer.MixTarget

const (
	/*The audio will be played only on the first channel. This is the default.*/
	MixTargetStereo MixTarget = 0
	/*The audio will be played on all surround channels.*/
	MixTargetSurround MixTarget = 1
	/*The audio will be played on the second channel, which is usually the center.*/
	MixTargetCenter MixTarget = 2
)
