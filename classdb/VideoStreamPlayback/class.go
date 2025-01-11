// Package VideoStreamPlayback provides methods for working with VideoStreamPlayback object instances.
package VideoStreamPlayback

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

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
This class is intended to be overridden by video decoder extensions with custom implementations of [VideoStream].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=VideoStreamPlayback)
*/
type Instance [1]gdclass.VideoStreamPlayback

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVideoStreamPlayback() Instance
}
type Interface interface {
	//Stops playback. May be called multiple times before [method _play], or in response to [method VideoStreamPlayer.stop]. [method _is_playing] should return false once stopped.
	Stop()
	//Called in response to [member VideoStreamPlayer.autoplay] or [method VideoStreamPlayer.play]. Note that manual playback may also invoke [method _stop] multiple times before this method is called. [method _is_playing] should return true once playing.
	Play()
	//Returns the playback state, as determined by calls to [method _play] and [method _stop].
	IsPlaying() bool
	//Set the paused status of video playback. [method _is_paused] must return [param paused]. Called in response to the [member VideoStreamPlayer.paused] setter.
	SetPaused(paused bool)
	//Returns the paused status, as set by [method _set_paused].
	IsPaused() bool
	//Returns the video duration in seconds, if known, or 0 if unknown.
	GetLength() Float.X
	//Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
	GetPlaybackPosition() Float.X
	//Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
	Seek(time Float.X)
	//Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
	SetAudioTrack(idx int)
	//Allocates a [Texture2D] in which decoded video frames will be drawn.
	GetTexture() [1]gdclass.Texture2D
	//Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
	Update(delta Float.X)
	//Returns the number of audio channels.
	GetChannels() int
	//Returns the audio sample rate used for mixing.
	GetMixRate() int
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) Stop()                                { return }
func (self Implementation) Play()                                { return }
func (self Implementation) IsPlaying() (_ bool)                  { return }
func (self Implementation) SetPaused(paused bool)                { return }
func (self Implementation) IsPaused() (_ bool)                   { return }
func (self Implementation) GetLength() (_ Float.X)               { return }
func (self Implementation) GetPlaybackPosition() (_ Float.X)     { return }
func (self Implementation) Seek(time Float.X)                    { return }
func (self Implementation) SetAudioTrack(idx int)                { return }
func (self Implementation) GetTexture() (_ [1]gdclass.Texture2D) { return }
func (self Implementation) Update(delta Float.X)                 { return }
func (self Implementation) GetChannels() (_ int)                 { return }
func (self Implementation) GetMixRate() (_ int)                  { return }

/*
Stops playback. May be called multiple times before [method _play], or in response to [method VideoStreamPlayer.stop]. [method _is_playing] should return false once stopped.
*/
func (Instance) _stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called in response to [member VideoStreamPlayer.autoplay] or [method VideoStreamPlayer.play]. Note that manual playback may also invoke [method _stop] multiple times before this method is called. [method _is_playing] should return true once playing.
*/
func (Instance) _play(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns the playback state, as determined by calls to [method _play] and [method _stop].
*/
func (Instance) _is_playing(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Set the paused status of video playback. [method _is_paused] must return [param paused]. Called in response to the [member VideoStreamPlayer.paused] setter.
*/
func (Instance) _set_paused(impl func(ptr unsafe.Pointer, paused bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var paused = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, paused)
	}
}

/*
Returns the paused status, as set by [method _set_paused].
*/
func (Instance) _is_paused(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the video duration in seconds, if known, or 0 if unknown.
*/
func (Instance) _get_length(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
*/
func (Instance) _get_playback_position(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
*/
func (Instance) _seek(impl func(ptr unsafe.Pointer, time Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var time = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(time))
	}
}

/*
Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
*/
func (Instance) _set_audio_track(impl func(ptr unsafe.Pointer, idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(idx))
	}
}

/*
Allocates a [Texture2D] in which decoded video frames will be drawn.
*/
func (Instance) _get_texture(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
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
Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
*/
func (Instance) _update(impl func(ptr unsafe.Pointer, delta Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(delta))
	}
}

/*
Returns the number of audio channels.
*/
func (Instance) _get_channels(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns the audio sample rate used for mixing.
*/
func (Instance) _get_mix_rate(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Render [param num_frames] audio frames (of [method _get_channels] floats each) from [param buffer], starting from index [param offset] in the array. Returns the number of audio frames rendered, or -1 on error.
*/
func (self Instance) MixAudio(num_frames int) int {
	return int(int(class(self).MixAudio(gd.Int(num_frames), gd.NewPackedFloat32Slice([1][]float32{}[0]), gd.Int(0))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VideoStreamPlayback

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VideoStreamPlayback"))
	casted := Instance{*(*gdclass.VideoStreamPlayback)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Stops playback. May be called multiple times before [method _play], or in response to [method VideoStreamPlayer.stop]. [method _is_playing] should return false once stopped.
*/
func (class) _stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called in response to [member VideoStreamPlayer.autoplay] or [method VideoStreamPlayer.play]. Note that manual playback may also invoke [method _stop] multiple times before this method is called. [method _is_playing] should return true once playing.
*/
func (class) _play(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Returns the playback state, as determined by calls to [method _play] and [method _stop].
*/
func (class) _is_playing(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Set the paused status of video playback. [method _is_paused] must return [param paused]. Called in response to the [member VideoStreamPlayer.paused] setter.
*/
func (class) _set_paused(impl func(ptr unsafe.Pointer, paused bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var paused = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, paused)
	}
}

/*
Returns the paused status, as set by [method _set_paused].
*/
func (class) _is_paused(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the video duration in seconds, if known, or 0 if unknown.
*/
func (class) _get_length(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
*/
func (class) _get_playback_position(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
*/
func (class) _seek(impl func(ptr unsafe.Pointer, time gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var time = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, time)
	}
}

/*
Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
*/
func (class) _set_audio_track(impl func(ptr unsafe.Pointer, idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var idx = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, idx)
	}
}

/*
Allocates a [Texture2D] in which decoded video frames will be drawn.
*/
func (class) _get_texture(impl func(ptr unsafe.Pointer) [1]gdclass.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
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
Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
*/
func (class) _update(impl func(ptr unsafe.Pointer, delta gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var delta = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, delta)
	}
}

/*
Returns the number of audio channels.
*/
func (class) _get_channels(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the audio sample rate used for mixing.
*/
func (class) _get_mix_rate(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Render [param num_frames] audio frames (of [method _get_channels] floats each) from [param buffer], starting from index [param offset] in the array. Returns the number of audio frames rendered, or -1 on error.
*/
//go:nosplit
func (self class) MixAudio(num_frames gd.Int, buffer gd.PackedFloat32Array, offset gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, num_frames)
	callframe.Arg(frame, pointers.Get(buffer))
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayback.Bind_mix_audio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVideoStreamPlayback() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVideoStreamPlayback() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_stop":
		return reflect.ValueOf(self._stop)
	case "_play":
		return reflect.ValueOf(self._play)
	case "_is_playing":
		return reflect.ValueOf(self._is_playing)
	case "_set_paused":
		return reflect.ValueOf(self._set_paused)
	case "_is_paused":
		return reflect.ValueOf(self._is_paused)
	case "_get_length":
		return reflect.ValueOf(self._get_length)
	case "_get_playback_position":
		return reflect.ValueOf(self._get_playback_position)
	case "_seek":
		return reflect.ValueOf(self._seek)
	case "_set_audio_track":
		return reflect.ValueOf(self._set_audio_track)
	case "_get_texture":
		return reflect.ValueOf(self._get_texture)
	case "_update":
		return reflect.ValueOf(self._update)
	case "_get_channels":
		return reflect.ValueOf(self._get_channels)
	case "_get_mix_rate":
		return reflect.ValueOf(self._get_mix_rate)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_stop":
		return reflect.ValueOf(self._stop)
	case "_play":
		return reflect.ValueOf(self._play)
	case "_is_playing":
		return reflect.ValueOf(self._is_playing)
	case "_set_paused":
		return reflect.ValueOf(self._set_paused)
	case "_is_paused":
		return reflect.ValueOf(self._is_paused)
	case "_get_length":
		return reflect.ValueOf(self._get_length)
	case "_get_playback_position":
		return reflect.ValueOf(self._get_playback_position)
	case "_seek":
		return reflect.ValueOf(self._seek)
	case "_set_audio_track":
		return reflect.ValueOf(self._set_audio_track)
	case "_get_texture":
		return reflect.ValueOf(self._get_texture)
	case "_update":
		return reflect.ValueOf(self._update)
	case "_get_channels":
		return reflect.ValueOf(self._get_channels)
	case "_get_mix_rate":
		return reflect.ValueOf(self._get_mix_rate)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("VideoStreamPlayback", func(ptr gd.Object) any {
		return [1]gdclass.VideoStreamPlayback{*(*gdclass.VideoStreamPlayback)(unsafe.Pointer(&ptr))}
	})
}
