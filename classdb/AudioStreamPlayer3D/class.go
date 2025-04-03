// Package AudioStreamPlayer3D provides methods for working with AudioStreamPlayer3D object instances.
package AudioStreamPlayer3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
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
Plays audio with positional sound effects, based on the relative position of the audio listener. Positional effects include distance attenuation, directionality, and the Doppler effect. For greater realism, a low-pass filter is applied to distant sounds. This can be disabled by setting [member attenuation_filter_cutoff_hz] to [code]20500[/code].
By default, audio is heard from the camera position. This can be changed by adding an [AudioListener3D] node to the scene and enabling it by calling [method AudioListener3D.make_current] on it.
See also [AudioStreamPlayer] to play a sound non-positionally.
[b]Note:[/b] Hiding an [AudioStreamPlayer3D] node does not disable its audio output. To temporarily disable an [AudioStreamPlayer3D]'s audio output, set [member volume_db] to a very low value like [code]-100[/code] (which isn't audible to human hearing).
*/
type Instance [1]gdclass.AudioStreamPlayer3D
type Expanded [1]gdclass.AudioStreamPlayer3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlayer3D() Instance
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
func (self Instance) Play() { //gd:AudioStreamPlayer3D.play
	Advanced(self).Play(float64(0.0))
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
func (self Expanded) Play(from_position Float.X) { //gd:AudioStreamPlayer3D.play
	Advanced(self).Play(float64(from_position))
}

/*
Sets the position from which audio will be played, in seconds.
*/
func (self Instance) SeekTo(to_position Float.X) { //gd:AudioStreamPlayer3D.seek
	Advanced(self).SeekTo(float64(to_position))
}

/*
Stops the audio.
*/
func (self Instance) Stop() { //gd:AudioStreamPlayer3D.stop
	Advanced(self).Stop()
}

/*
Returns the position in the [AudioStream].
*/
func (self Instance) GetPlaybackPosition() Float.X { //gd:AudioStreamPlayer3D.get_playback_position
	return Float.X(Float.X(Advanced(self).GetPlaybackPosition()))
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
func (self Instance) HasStreamPlayback() bool { //gd:AudioStreamPlayer3D.has_stream_playback
	return bool(Advanced(self).HasStreamPlayback())
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer3D].
*/
func (self Instance) GetStreamPlayback() [1]gdclass.AudioStreamPlayback { //gd:AudioStreamPlayer3D.get_stream_playback
	return [1]gdclass.AudioStreamPlayback(Advanced(self).GetStreamPlayback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlayer3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlayer3D"))
	casted := Instance{*(*gdclass.AudioStreamPlayer3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Stream() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetStream())
}

func (self Instance) SetStream(value [1]gdclass.AudioStream) {
	class(self).SetStream(value)
}

func (self Instance) AttenuationModel() gdclass.AudioStreamPlayer3DAttenuationModel {
	return gdclass.AudioStreamPlayer3DAttenuationModel(class(self).GetAttenuationModel())
}

func (self Instance) SetAttenuationModel(value gdclass.AudioStreamPlayer3DAttenuationModel) {
	class(self).SetAttenuationModel(value)
}

func (self Instance) VolumeDb() Float.X {
	return Float.X(Float.X(class(self).GetVolumeDb()))
}

func (self Instance) SetVolumeDb(value Float.X) {
	class(self).SetVolumeDb(float64(value))
}

func (self Instance) VolumeLinear() Float.X {
	return Float.X(Float.X(class(self).GetVolumeLinear()))
}

func (self Instance) SetVolumeLinear(value Float.X) {
	class(self).SetVolumeLinear(float64(value))
}

func (self Instance) UnitSize() Float.X {
	return Float.X(Float.X(class(self).GetUnitSize()))
}

func (self Instance) SetUnitSize(value Float.X) {
	class(self).SetUnitSize(float64(value))
}

func (self Instance) MaxDb() Float.X {
	return Float.X(Float.X(class(self).GetMaxDb()))
}

func (self Instance) SetMaxDb(value Float.X) {
	class(self).SetMaxDb(float64(value))
}

func (self Instance) PitchScale() Float.X {
	return Float.X(Float.X(class(self).GetPitchScale()))
}

func (self Instance) SetPitchScale(value Float.X) {
	class(self).SetPitchScale(float64(value))
}

func (self Instance) Playing() bool {
	return bool(class(self).IsPlaying())
}

func (self Instance) SetPlaying(value bool) {
	class(self).SetPlaying(value)
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
	class(self).SetMaxDistance(float64(value))
}

func (self Instance) MaxPolyphony() int {
	return int(int(class(self).GetMaxPolyphony()))
}

func (self Instance) SetMaxPolyphony(value int) {
	class(self).SetMaxPolyphony(int64(value))
}

func (self Instance) PanningStrength() Float.X {
	return Float.X(Float.X(class(self).GetPanningStrength()))
}

func (self Instance) SetPanningStrength(value Float.X) {
	class(self).SetPanningStrength(float64(value))
}

func (self Instance) Bus() string {
	return string(class(self).GetBus().String())
}

func (self Instance) SetBus(value string) {
	class(self).SetBus(String.Name(String.New(value)))
}

func (self Instance) AreaMask() int {
	return int(int(class(self).GetAreaMask()))
}

func (self Instance) SetAreaMask(value int) {
	class(self).SetAreaMask(int64(value))
}

func (self Instance) PlaybackType() gdclass.AudioServerPlaybackType {
	return gdclass.AudioServerPlaybackType(class(self).GetPlaybackType())
}

func (self Instance) SetPlaybackType(value gdclass.AudioServerPlaybackType) {
	class(self).SetPlaybackType(value)
}

func (self Instance) EmissionAngleEnabled() bool {
	return bool(class(self).IsEmissionAngleEnabled())
}

func (self Instance) SetEmissionAngleEnabled(value bool) {
	class(self).SetEmissionAngleEnabled(value)
}

func (self Instance) EmissionAngleDegrees() Float.X {
	return Float.X(Float.X(class(self).GetEmissionAngle()))
}

func (self Instance) SetEmissionAngleDegrees(value Float.X) {
	class(self).SetEmissionAngle(float64(value))
}

func (self Instance) EmissionAngleFilterAttenuationDb() Float.X {
	return Float.X(Float.X(class(self).GetEmissionAngleFilterAttenuationDb()))
}

func (self Instance) SetEmissionAngleFilterAttenuationDb(value Float.X) {
	class(self).SetEmissionAngleFilterAttenuationDb(float64(value))
}

func (self Instance) AttenuationFilterCutoffHz() Float.X {
	return Float.X(Float.X(class(self).GetAttenuationFilterCutoffHz()))
}

func (self Instance) SetAttenuationFilterCutoffHz(value Float.X) {
	class(self).SetAttenuationFilterCutoffHz(float64(value))
}

func (self Instance) AttenuationFilterDb() Float.X {
	return Float.X(Float.X(class(self).GetAttenuationFilterDb()))
}

func (self Instance) SetAttenuationFilterDb(value Float.X) {
	class(self).SetAttenuationFilterDb(float64(value))
}

func (self Instance) DopplerTracking() gdclass.AudioStreamPlayer3DDopplerTracking {
	return gdclass.AudioStreamPlayer3DDopplerTracking(class(self).GetDopplerTracking())
}

func (self Instance) SetDopplerTracking(value gdclass.AudioStreamPlayer3DDopplerTracking) {
	class(self).SetDopplerTracking(value)
}

//go:nosplit
func (self class) SetStream(stream [1]gdclass.AudioStream) { //gd:AudioStreamPlayer3D.set_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStream() [1]gdclass.AudioStream { //gd:AudioStreamPlayer3D.get_stream
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumeDb(volume_db float64) { //gd:AudioStreamPlayer3D.set_volume_db
	var frame = callframe.New()
	callframe.Arg(frame, volume_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumeDb() float64 { //gd:AudioStreamPlayer3D.get_volume_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumeLinear(volume_linear float64) { //gd:AudioStreamPlayer3D.set_volume_linear
	var frame = callframe.New()
	callframe.Arg(frame, volume_linear)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_volume_linear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumeLinear() float64 { //gd:AudioStreamPlayer3D.get_volume_linear
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_volume_linear, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUnitSize(unit_size float64) { //gd:AudioStreamPlayer3D.set_unit_size
	var frame = callframe.New()
	callframe.Arg(frame, unit_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_unit_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUnitSize() float64 { //gd:AudioStreamPlayer3D.get_unit_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_unit_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxDb(max_db float64) { //gd:AudioStreamPlayer3D.set_max_db
	var frame = callframe.New()
	callframe.Arg(frame, max_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_max_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDb() float64 { //gd:AudioStreamPlayer3D.get_max_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_max_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPitchScale(pitch_scale float64) { //gd:AudioStreamPlayer3D.set_pitch_scale
	var frame = callframe.New()
	callframe.Arg(frame, pitch_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitchScale() float64 { //gd:AudioStreamPlayer3D.get_pitch_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
//go:nosplit
func (self class) Play(from_position float64) { //gd:AudioStreamPlayer3D.play
	var frame = callframe.New()
	callframe.Arg(frame, from_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the position from which audio will be played, in seconds.
*/
//go:nosplit
func (self class) SeekTo(to_position float64) { //gd:AudioStreamPlayer3D.seek
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the audio.
*/
//go:nosplit
func (self class) Stop() { //gd:AudioStreamPlayer3D.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPlaying() bool { //gd:AudioStreamPlayer3D.is_playing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position in the [AudioStream].
*/
//go:nosplit
func (self class) GetPlaybackPosition() float64 { //gd:AudioStreamPlayer3D.get_playback_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_playback_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBus(bus String.Name) { //gd:AudioStreamPlayer3D.set_bus
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(bus)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBus() String.Name { //gd:AudioStreamPlayer3D.get_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(enable bool) { //gd:AudioStreamPlayer3D.set_autoplay
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoplayEnabled() bool { //gd:AudioStreamPlayer3D.is_autoplay_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_is_autoplay_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaying(enable bool) { //gd:AudioStreamPlayer3D.set_playing
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMaxDistance(meters float64) { //gd:AudioStreamPlayer3D.set_max_distance
	var frame = callframe.New()
	callframe.Arg(frame, meters)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDistance() float64 { //gd:AudioStreamPlayer3D.get_max_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAreaMask(mask int64) { //gd:AudioStreamPlayer3D.set_area_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_area_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAreaMask() int64 { //gd:AudioStreamPlayer3D.get_area_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_area_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionAngle(degrees float64) { //gd:AudioStreamPlayer3D.set_emission_angle
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_emission_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionAngle() float64 { //gd:AudioStreamPlayer3D.get_emission_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_emission_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionAngleEnabled(enabled bool) { //gd:AudioStreamPlayer3D.set_emission_angle_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_emission_angle_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmissionAngleEnabled() bool { //gd:AudioStreamPlayer3D.is_emission_angle_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_is_emission_angle_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionAngleFilterAttenuationDb(db float64) { //gd:AudioStreamPlayer3D.set_emission_angle_filter_attenuation_db
	var frame = callframe.New()
	callframe.Arg(frame, db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_emission_angle_filter_attenuation_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionAngleFilterAttenuationDb() float64 { //gd:AudioStreamPlayer3D.get_emission_angle_filter_attenuation_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_emission_angle_filter_attenuation_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttenuationFilterCutoffHz(degrees float64) { //gd:AudioStreamPlayer3D.set_attenuation_filter_cutoff_hz
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_attenuation_filter_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttenuationFilterCutoffHz() float64 { //gd:AudioStreamPlayer3D.get_attenuation_filter_cutoff_hz
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_attenuation_filter_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttenuationFilterDb(db float64) { //gd:AudioStreamPlayer3D.set_attenuation_filter_db
	var frame = callframe.New()
	callframe.Arg(frame, db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_attenuation_filter_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttenuationFilterDb() float64 { //gd:AudioStreamPlayer3D.get_attenuation_filter_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_attenuation_filter_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttenuationModel(model gdclass.AudioStreamPlayer3DAttenuationModel) { //gd:AudioStreamPlayer3D.set_attenuation_model
	var frame = callframe.New()
	callframe.Arg(frame, model)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_attenuation_model, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttenuationModel() gdclass.AudioStreamPlayer3DAttenuationModel { //gd:AudioStreamPlayer3D.get_attenuation_model
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamPlayer3DAttenuationModel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_attenuation_model, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDopplerTracking(mode gdclass.AudioStreamPlayer3DDopplerTracking) { //gd:AudioStreamPlayer3D.set_doppler_tracking
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDopplerTracking() gdclass.AudioStreamPlayer3DDopplerTracking { //gd:AudioStreamPlayer3D.get_doppler_tracking
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamPlayer3DDopplerTracking](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_doppler_tracking, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStreamPaused(pause bool) { //gd:AudioStreamPlayer3D.set_stream_paused
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_stream_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPaused() bool { //gd:AudioStreamPlayer3D.get_stream_paused
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_stream_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxPolyphony(max_polyphony int64) { //gd:AudioStreamPlayer3D.set_max_polyphony
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxPolyphony() int64 { //gd:AudioStreamPlayer3D.get_max_polyphony
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPanningStrength(panning_strength float64) { //gd:AudioStreamPlayer3D.set_panning_strength
	var frame = callframe.New()
	callframe.Arg(frame, panning_strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_panning_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPanningStrength() float64 { //gd:AudioStreamPlayer3D.get_panning_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_panning_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
//go:nosplit
func (self class) HasStreamPlayback() bool { //gd:AudioStreamPlayer3D.has_stream_playback
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_has_stream_playback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer3D].
*/
//go:nosplit
func (self class) GetStreamPlayback() [1]gdclass.AudioStreamPlayback { //gd:AudioStreamPlayer3D.get_stream_playback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStreamPlayback{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamPlayback](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaybackType(playback_type gdclass.AudioServerPlaybackType) { //gd:AudioStreamPlayer3D.set_playback_type
	var frame = callframe.New()
	callframe.Arg(frame, playback_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_set_playback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaybackType() gdclass.AudioServerPlaybackType { //gd:AudioStreamPlayer3D.get_playback_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioServerPlaybackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer3D.Bind_get_playback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsAudioStreamPlayer3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlayer3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced          { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance       { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced              { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance           { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlayer3D", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlayer3D{*(*gdclass.AudioStreamPlayer3D)(unsafe.Pointer(&ptr))}
	})
}

type AttenuationModel = gdclass.AudioStreamPlayer3DAttenuationModel //gd:AudioStreamPlayer3D.AttenuationModel

const (
	/*Attenuation of loudness according to linear distance.*/
	AttenuationInverseDistance AttenuationModel = 0
	/*Attenuation of loudness according to squared distance.*/
	AttenuationInverseSquareDistance AttenuationModel = 1
	/*Attenuation of loudness according to logarithmic distance.*/
	AttenuationLogarithmic AttenuationModel = 2
	/*No attenuation of loudness according to distance. The sound will still be heard positionally, unlike an [AudioStreamPlayer]. [constant ATTENUATION_DISABLED] can be combined with a [member max_distance] value greater than [code]0.0[/code] to achieve linear attenuation clamped to a sphere of a defined size.*/
	AttenuationDisabled AttenuationModel = 3
)

type DopplerTracking = gdclass.AudioStreamPlayer3DDopplerTracking //gd:AudioStreamPlayer3D.DopplerTracking

const (
	/*Disables doppler tracking.*/
	DopplerTrackingDisabled DopplerTracking = 0
	/*Executes doppler tracking during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	DopplerTrackingIdleStep DopplerTracking = 1
	/*Executes doppler tracking during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	DopplerTrackingPhysicsStep DopplerTracking = 2
)
