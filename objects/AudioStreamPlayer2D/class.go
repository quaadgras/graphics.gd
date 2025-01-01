package AudioStreamPlayer2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Plays audio that is attenuated with distance to the listener.
By default, audio is heard from the screen center. This can be changed by adding an [AudioListener2D] node to the scene and enabling it by calling [method AudioListener2D.make_current] on it.
See also [AudioStreamPlayer] to play a sound non-positionally.
[b]Note:[/b] Hiding an [AudioStreamPlayer2D] node does not disable its audio output. To temporarily disable an [AudioStreamPlayer2D]'s audio output, set [member volume_db] to a very low value like [code]-100[/code] (which isn't audible to human hearing).
*/
type Instance [1]classdb.AudioStreamPlayer2D
type Any interface {
	gd.IsClass
	AsAudioStreamPlayer2D() Instance
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
func (self Instance) Play() {
	class(self).Play(gd.Float(0.0))
}

/*
Sets the position from which audio will be played, in seconds.
*/
func (self Instance) SeekTo(to_position Float.X) {
	class(self).SeekTo(gd.Float(to_position))
}

/*
Stops the audio.
*/
func (self Instance) Stop() {
	class(self).Stop()
}

/*
Returns the position in the [AudioStream].
*/
func (self Instance) GetPlaybackPosition() Float.X {
	return Float.X(Float.X(class(self).GetPlaybackPosition()))
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
func (self Instance) HasStreamPlayback() bool {
	return bool(class(self).HasStreamPlayback())
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer2D].
*/
func (self Instance) GetStreamPlayback() objects.AudioStreamPlayback {
	return objects.AudioStreamPlayback(class(self).GetStreamPlayback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamPlayer2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlayer2D"))
	return Instance{classdb.AudioStreamPlayer2D(object)}
}

func (self Instance) Stream() objects.AudioStream {
	return objects.AudioStream(class(self).GetStream())
}

func (self Instance) SetStream(value objects.AudioStream) {
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

func (self Instance) MaxDistance() Float.X {
	return Float.X(Float.X(class(self).GetMaxDistance()))
}

func (self Instance) SetMaxDistance(value Float.X) {
	class(self).SetMaxDistance(gd.Float(value))
}

func (self Instance) Attenuation() Float.X {
	return Float.X(Float.X(class(self).GetAttenuation()))
}

func (self Instance) SetAttenuation(value Float.X) {
	class(self).SetAttenuation(gd.Float(value))
}

func (self Instance) MaxPolyphony() int {
	return int(int(class(self).GetMaxPolyphony()))
}

func (self Instance) SetMaxPolyphony(value int) {
	class(self).SetMaxPolyphony(gd.Int(value))
}

func (self Instance) PanningStrength() Float.X {
	return Float.X(Float.X(class(self).GetPanningStrength()))
}

func (self Instance) SetPanningStrength(value Float.X) {
	class(self).SetPanningStrength(gd.Float(value))
}

func (self Instance) Bus() string {
	return string(class(self).GetBus().String())
}

func (self Instance) SetBus(value string) {
	class(self).SetBus(gd.NewStringName(value))
}

func (self Instance) AreaMask() int {
	return int(int(class(self).GetAreaMask()))
}

func (self Instance) SetAreaMask(value int) {
	class(self).SetAreaMask(gd.Int(value))
}

func (self Instance) PlaybackType() classdb.AudioServerPlaybackType {
	return classdb.AudioServerPlaybackType(class(self).GetPlaybackType())
}

func (self Instance) SetPlaybackType(value classdb.AudioServerPlaybackType) {
	class(self).SetPlaybackType(value)
}

//go:nosplit
func (self class) SetStream(stream objects.AudioStream) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStream() objects.AudioStream {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AudioStream{classdb.AudioStream(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumeDb(volume_db gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumeDb() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPitchScale(pitch_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pitch_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitchScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
//go:nosplit
func (self class) Play(from_position gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, from_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the position from which audio will be played, in seconds.
*/
//go:nosplit
func (self class) SeekTo(to_position gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Stops the audio.
*/
//go:nosplit
func (self class) Stop() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsPlaying() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position in the [AudioStream].
*/
//go:nosplit
func (self class) GetPlaybackPosition() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_playback_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBus(bus gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bus))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBus() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoplayEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_is_autoplay_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxDistance(pixels gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttenuation(curve gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_attenuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttenuation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_attenuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAreaMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_area_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAreaMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_area_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStreamPaused(pause bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_stream_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPaused() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_stream_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxPolyphony(max_polyphony gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxPolyphony() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPanningStrength(panning_strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, panning_strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_panning_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPanningStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_panning_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
//go:nosplit
func (self class) HasStreamPlayback() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_has_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer2D].
*/
//go:nosplit
func (self class) GetStreamPlayback() objects.AudioStreamPlayback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AudioStreamPlayback{classdb.AudioStreamPlayback(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaybackType(playback_type classdb.AudioServerPlaybackType) {
	var frame = callframe.New()
	callframe.Arg(frame, playback_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_playback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaybackType() classdb.AudioServerPlaybackType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioServerPlaybackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_playback_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsAudioStreamPlayer2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlayer2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced          { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance       { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("AudioStreamPlayer2D", func(ptr gd.Object) any { return [1]classdb.AudioStreamPlayer2D{classdb.AudioStreamPlayer2D(ptr)} })
}
