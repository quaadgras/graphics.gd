// Package AudioStream provides methods for working with AudioStream object instances.
package AudioStream

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Base class for audio streams. Audio streams are used for sound effects and music playback, and support WAV (via [AudioStreamWAV]) and Ogg (via [AudioStreamOggVorbis]) file formats.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=AudioStream)
*/
type Instance [1]gdclass.AudioStream
type Any interface {
	gd.IsClass
	AsAudioStream() Instance
}
type Interface interface {
	//Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
	InstantiatePlayback() [1]gdclass.AudioStreamPlayback
	//Override this method to customize the name assigned to this audio stream. Unused by the engine.
	GetStreamName() string
	//Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
	GetLength() Float.X
	//Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
	IsMonophonic() bool
	//Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
	//Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
	GetBpm() Float.X
	//Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
	//Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
	GetBeatCount() int
	//Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
	GetParameterList() gd.Array
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) InstantiatePlayback() (_ [1]gdclass.AudioStreamPlayback) { return }
func (self Implementation) GetStreamName() (_ string)                               { return }
func (self Implementation) GetLength() (_ Float.X)                                  { return }
func (self Implementation) IsMonophonic() (_ bool)                                  { return }
func (self Implementation) GetBpm() (_ Float.X)                                     { return }
func (self Implementation) GetBeatCount() (_ int)                                   { return }
func (self Implementation) GetParameterList() (_ gd.Array)                          { return }

/*
Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
*/
func (Instance) _instantiate_playback(impl func(ptr unsafe.Pointer) [1]gdclass.AudioStreamPlayback) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to customize the name assigned to this audio stream. Unused by the engine.
*/
func (Instance) _get_stream_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
*/
func (Instance) _get_length(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
*/
func (Instance) _is_monophonic(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (Instance) _get_bpm(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (Instance) _get_beat_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
*/
func (Instance) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the length of the audio stream in seconds.
*/
func (self Instance) GetLength() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

/*
Returns [code]true[/code] if this audio stream only supports one channel ([i]monophony[/i]), or [code]false[/code] if the audio stream supports two or more channels ([i]polyphony[/i]).
*/
func (self Instance) IsMonophonic() bool {
	return bool(class(self).IsMonophonic())
}

/*
Returns a newly created [AudioStreamPlayback] intended to play this audio stream. Useful for when you want to extend [method _instantiate_playback] but call [method instantiate_playback] from an internally held AudioStream subresource. An example of this can be found in the source code for [code]AudioStreamRandomPitch::instantiate_playback[/code].
*/
func (self Instance) InstantiatePlayback() [1]gdclass.AudioStreamPlayback {
	return [1]gdclass.AudioStreamPlayback(class(self).InstantiatePlayback())
}

/*
Returns if the current [AudioStream] can be used as a sample. Only static streams can be sampled.
*/
func (self Instance) CanBeSampled() bool {
	return bool(class(self).CanBeSampled())
}

/*
Generates an [AudioSample] based on the current stream.
*/
func (self Instance) GenerateSample() [1]gdclass.AudioSample {
	return [1]gdclass.AudioSample(class(self).GenerateSample())
}

/*
Returns [code]true[/code] if the stream is a collection of other streams, [code]false[/code] otherwise.
*/
func (self Instance) IsMetaStream() bool {
	return bool(class(self).IsMetaStream())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStream

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStream"))
	casted := Instance{*(*gdclass.AudioStream)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to customize the returned value of [method instantiate_playback]. Should returned a new [AudioStreamPlayback] created when the stream is played (such as by an [AudioStreamPlayer])..
*/
func (class) _instantiate_playback(impl func(ptr unsafe.Pointer) [1]gdclass.AudioStreamPlayback) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to customize the name assigned to this audio stream. Unused by the engine.
*/
func (class) _get_stream_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to customize the returned value of [method get_length]. Should return the length of this audio stream, in seconds.
*/
func (class) _get_length(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to customize the returned value of [method is_monophonic]. Should return [code]true[/code] if this audio stream only supports one channel.
*/
func (class) _is_monophonic(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return the tempo of this audio stream, in beats per minute (BPM). Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (class) _get_bpm(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return the total number of beats of this audio stream. Used by the engine to determine the position of every beat.
Ideally, the returned value should be based off the stream's sample rate ([member AudioStreamWAV.mix_rate], for example).
*/
func (class) _get_beat_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the controllable parameters of this stream. This array contains dictionaries with a property info description format (see [method Object.get_property_list]). Additionally, the default value for this parameter must be added tho each dictionary in "default_value" field.
*/
func (class) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the length of the audio stream in seconds.
*/
//go:nosplit
func (self class) GetLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStream.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this audio stream only supports one channel ([i]monophony[/i]), or [code]false[/code] if the audio stream supports two or more channels ([i]polyphony[/i]).
*/
//go:nosplit
func (self class) IsMonophonic() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStream.Bind_is_monophonic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a newly created [AudioStreamPlayback] intended to play this audio stream. Useful for when you want to extend [method _instantiate_playback] but call [method instantiate_playback] from an internally held AudioStream subresource. An example of this can be found in the source code for [code]AudioStreamRandomPitch::instantiate_playback[/code].
*/
//go:nosplit
func (self class) InstantiatePlayback() [1]gdclass.AudioStreamPlayback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStream.Bind_instantiate_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStreamPlayback{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamPlayback](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns if the current [AudioStream] can be used as a sample. Only static streams can be sampled.
*/
//go:nosplit
func (self class) CanBeSampled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStream.Bind_can_be_sampled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Generates an [AudioSample] based on the current stream.
*/
//go:nosplit
func (self class) GenerateSample() [1]gdclass.AudioSample {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStream.Bind_generate_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioSample{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioSample](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the stream is a collection of other streams, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) IsMetaStream() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStream.Bind_is_meta_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnParameterListChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("parameter_list_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAudioStream() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStream() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_instantiate_playback":
		return reflect.ValueOf(self._instantiate_playback)
	case "_get_stream_name":
		return reflect.ValueOf(self._get_stream_name)
	case "_get_length":
		return reflect.ValueOf(self._get_length)
	case "_is_monophonic":
		return reflect.ValueOf(self._is_monophonic)
	case "_get_bpm":
		return reflect.ValueOf(self._get_bpm)
	case "_get_beat_count":
		return reflect.ValueOf(self._get_beat_count)
	case "_get_parameter_list":
		return reflect.ValueOf(self._get_parameter_list)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback":
		return reflect.ValueOf(self._instantiate_playback)
	case "_get_stream_name":
		return reflect.ValueOf(self._get_stream_name)
	case "_get_length":
		return reflect.ValueOf(self._get_length)
	case "_is_monophonic":
		return reflect.ValueOf(self._is_monophonic)
	case "_get_bpm":
		return reflect.ValueOf(self._get_bpm)
	case "_get_beat_count":
		return reflect.ValueOf(self._get_beat_count)
	case "_get_parameter_list":
		return reflect.ValueOf(self._get_parameter_list)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("AudioStream", func(ptr gd.Object) any { return [1]gdclass.AudioStream{*(*gdclass.AudioStream)(unsafe.Pointer(&ptr))} })
}
