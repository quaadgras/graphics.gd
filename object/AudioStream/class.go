package AudioStream

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for audio streams. Audio streams are used for sound effects and music playback, and support WAV (via [AudioStreamWAV]) and Ogg (via [AudioStreamOggVorbis]) file formats.
	// AudioStream methods that can be overridden by a [Class] that extends it.
	type AudioStream interface {
		//Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
		InstantiatePlayback() object.AudioStreamPlayback
		//Override this method to customize the name assigned to this audio stream. Unused by the engine.
		GetStreamName() gd.String
		//Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
		GetLength() gd.Float
		//Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
		IsMonophonic() bool
		//Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
		//Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
		GetBpm() gd.Float
		//Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
		//Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
		GetBeatCount() gd.Int
		//Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
		GetParameterList() gd.ArrayOf[gd.Dictionary]
	}

*/
type Simple [1]classdb.AudioStream
func (Simple) _instantiate_playback(impl func(ptr unsafe.Pointer) [1]classdb.AudioStreamPlayback, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _get_stream_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _get_length(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _is_monophonic(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_bpm(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _get_beat_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (self Simple) GetLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLength()))
}
func (self Simple) IsMonophonic() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMonophonic())
}
func (self Simple) InstantiatePlayback() [1]classdb.AudioStreamPlayback {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStreamPlayback(Expert(self).InstantiatePlayback(gc))
}
func (self Simple) CanBeSampled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanBeSampled())
}
func (self Simple) GenerateSample() [1]classdb.AudioSample {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioSample(Expert(self).GenerateSample(gc))
}
func (self Simple) IsMetaStream() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMetaStream())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStream
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
*/
func (class) _instantiate_playback(impl func(ptr unsafe.Pointer) object.AudioStreamPlayback, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (self class) InstantiatePlayback(ctx gd.Lifetime) object.AudioStreamPlayback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_instantiate_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStreamPlayback
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
func (self class) GenerateSample(ctx gd.Lifetime) object.AudioSample {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStream.Bind_generate_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioSample
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

//go:nosplit
func (self class) AsAudioStream() Expert { return self[0].AsAudioStream() }


//go:nosplit
func (self Simple) AsAudioStream() Simple { return self[0].AsAudioStream() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback": return reflect.ValueOf(self._instantiate_playback);
	case "_get_stream_name": return reflect.ValueOf(self._get_stream_name);
	case "_get_length": return reflect.ValueOf(self._get_length);
	case "_is_monophonic": return reflect.ValueOf(self._is_monophonic);
	case "_get_bpm": return reflect.ValueOf(self._get_bpm);
	case "_get_beat_count": return reflect.ValueOf(self._get_beat_count);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback": return reflect.ValueOf(self._instantiate_playback);
	case "_get_stream_name": return reflect.ValueOf(self._get_stream_name);
	case "_get_length": return reflect.ValueOf(self._get_length);
	case "_is_monophonic": return reflect.ValueOf(self._is_monophonic);
	case "_get_bpm": return reflect.ValueOf(self._get_bpm);
	case "_get_beat_count": return reflect.ValueOf(self._get_beat_count);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStream", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
