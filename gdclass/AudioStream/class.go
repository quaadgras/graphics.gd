package AudioStream

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for audio streams. Audio streams are used for sound effects and music playback, and support WAV (via [AudioStreamWAV]) and Ogg (via [AudioStreamOggVorbis]) file formats.
	// AudioStream methods that can be overridden by a [Class] that extends it.
	type AudioStream interface {
		//Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
		InstantiatePlayback() gdclass.AudioStreamPlayback
		//Override this method to customize the name assigned to this audio stream. Unused by the engine.
		GetStreamName() string
		//Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
		GetLength() float64
		//Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
		IsMonophonic() bool
		//Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
		//Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
		GetBpm() float64
		//Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
		//Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
		GetBeatCount() int
		//Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
		GetParameterList() gd.ArrayOf[gd.Dictionary]
	}

*/
type Go [1]classdb.AudioStream

/*
Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
*/
func (Go) _instantiate_playback(impl func(ptr unsafe.Pointer) gdclass.AudioStreamPlayback, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}

/*
Override this method to customize the name assigned to this audio stream. Unused by the engine.
*/
func (Go) _get_stream_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
*/
func (Go) _get_length(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}

/*
Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
*/
func (Go) _is_monophonic(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (Go) _get_bpm(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}

/*
Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (Go) _get_beat_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
*/
func (Go) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}

/*
Returns the length of the audio stream in seconds.
*/
func (self Go) GetLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLength()))
}

/*
Returns [code]true[/code] if this audio stream only supports one channel ([i]monophony[/i]), or [code]false[/code] if the audio stream supports two or more channels ([i]polyphony[/i]).
*/
func (self Go) IsMonophonic() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMonophonic())
}

/*
Returns a newly created [AudioStreamPlayback] intended to play this audio stream. Useful for when you want to extend [method _instantiate_playback] but call [method instantiate_playback] from an internally held AudioStream subresource. An example of this can be found in the source code for [code]AudioStreamRandomPitch::instantiate_playback[/code].
*/
func (self Go) InstantiatePlayback() gdclass.AudioStreamPlayback {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStreamPlayback(class(self).InstantiatePlayback(gc))
}

/*
Returns if the current [AudioStream] can be used as a sample. Only static streams can be sampled.
*/
func (self Go) CanBeSampled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).CanBeSampled())
}

/*
Generates an [AudioSample] based on the current stream.
*/
func (self Go) GenerateSample() gdclass.AudioSample {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioSample(class(self).GenerateSample(gc))
}

/*
Returns [code]true[/code] if the stream is a collection of other streams, [code]false[/code] otherwise.
*/
func (self Go) IsMetaStream() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMetaStream())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStream
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStream"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
*/
func (class) _instantiate_playback(impl func(ptr unsafe.Pointer) gdclass.AudioStreamPlayback, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Override this method to customize the name assigned to this audio stream. Unused by the engine.
*/
func (class) _get_stream_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
*/
func (class) _get_length(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
*/
func (class) _is_monophonic(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (class) _get_bpm(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (class) _get_beat_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
*/
func (class) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Returns the length of the audio stream in seconds.
*/
//go:nosplit
func (self class) GetLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this audio stream only supports one channel ([i]monophony[/i]), or [code]false[/code] if the audio stream supports two or more channels ([i]polyphony[/i]).
*/
//go:nosplit
func (self class) IsMonophonic() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_is_monophonic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a newly created [AudioStreamPlayback] intended to play this audio stream. Useful for when you want to extend [method _instantiate_playback] but call [method instantiate_playback] from an internally held AudioStream subresource. An example of this can be found in the source code for [code]AudioStreamRandomPitch::instantiate_playback[/code].
*/
//go:nosplit
func (self class) InstantiatePlayback(ctx gd.Lifetime) gdclass.AudioStreamPlayback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_instantiate_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStreamPlayback
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns if the current [AudioStream] can be used as a sample. Only static streams can be sampled.
*/
//go:nosplit
func (self class) CanBeSampled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_can_be_sampled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Generates an [AudioSample] based on the current stream.
*/
//go:nosplit
func (self class) GenerateSample(ctx gd.Lifetime) gdclass.AudioSample {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_generate_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioSample
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the stream is a collection of other streams, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsMetaStream() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_is_meta_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnParameterListChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("parameter_list_changed"), gc.Callable(cb), 0)
}


func (self class) AsAudioStream() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStream() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback": return reflect.ValueOf(self._instantiate_playback);
	case "_get_stream_name": return reflect.ValueOf(self._get_stream_name);
	case "_get_length": return reflect.ValueOf(self._get_length);
	case "_is_monophonic": return reflect.ValueOf(self._is_monophonic);
	case "_get_bpm": return reflect.ValueOf(self._get_bpm);
	case "_get_beat_count": return reflect.ValueOf(self._get_beat_count);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback": return reflect.ValueOf(self._instantiate_playback);
	case "_get_stream_name": return reflect.ValueOf(self._get_stream_name);
	case "_get_length": return reflect.ValueOf(self._get_length);
	case "_is_monophonic": return reflect.ValueOf(self._is_monophonic);
	case "_get_bpm": return reflect.ValueOf(self._get_bpm);
	case "_get_beat_count": return reflect.ValueOf(self._get_beat_count);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("AudioStream", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
