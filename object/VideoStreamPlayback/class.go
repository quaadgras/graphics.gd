package VideoStreamPlayback

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
This class is intended to be overridden by video decoder extensions with custom implementations of [VideoStream].
	// VideoStreamPlayback methods that can be overridden by a [Class] that extends it.
	type VideoStreamPlayback interface {
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
		GetLength() gd.Float
		//Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
		GetPlaybackPosition() gd.Float
		//Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
		Seek(time gd.Float) 
		//Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
		SetAudioTrack(idx gd.Int) 
		//Allocates a [Texture2D] in which decoded video frames will be drawn.
		GetTexture() object.Texture2D
		//Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
		Update(delta gd.Float) 
		//Returns the number of audio channels.
		GetChannels() gd.Int
		//Returns the audio sample rate used for mixing.
		GetMixRate() gd.Int
	}

*/
type Simple [1]classdb.VideoStreamPlayback
func (self Simple) MixAudio(num_frames int, buffer gd.PackedFloat32Array, offset int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MixAudio(gd.Int(num_frames), buffer, gd.Int(offset))))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VideoStreamPlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Stops playback. May be called multiple times before [method _play], or in response to [method VideoStreamPlayer.stop]. [method _is_playing] should return false once stopped.
*/
func (class) _stop(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called in response to [member VideoStreamPlayer.autoplay] or [method VideoStreamPlayer.play]. Note that manual playback may also invoke [method _stop] multiple times before this method is called. [method _is_playing] should return true once playing.
*/
func (class) _play(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Returns the playback state, as determined by calls to [method _play] and [method _stop].
*/
func (class) _is_playing(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Set the paused status of video playback. [method _is_paused] must return [param paused]. Called in response to the [member VideoStreamPlayer.paused] setter.
*/
func (class) _set_paused(impl func(ptr unsafe.Pointer, paused bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var paused = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, paused)
		ctx.End()
	}
}

/*
Returns the paused status, as set by [method _set_paused].
*/
func (class) _is_paused(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns the video duration in seconds, if known, or 0 if unknown.
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
Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
*/
func (class) _get_playback_position(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
*/
func (class) _seek(impl func(ptr unsafe.Pointer, time gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var time = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, time)
		ctx.End()
	}
}

/*
Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
*/
func (class) _set_audio_track(impl func(ptr unsafe.Pointer, idx gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, idx)
		ctx.End()
	}
}

/*
Allocates a [Texture2D] in which decoded video frames will be drawn.
*/
func (class) _get_texture(impl func(ptr unsafe.Pointer) object.Texture2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
*/
func (class) _update(impl func(ptr unsafe.Pointer, delta gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, delta)
		ctx.End()
	}
}

/*
Returns the number of audio channels.
*/
func (class) _get_channels(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns the audio sample rate used for mixing.
*/
func (class) _get_mix_rate(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Render [param num_frames] audio frames (of [method _get_channels] floats each) from [param buffer], starting from index [param offset] in the array. Returns the number of audio frames rendered, or -1 on error.
*/
//go:nosplit
func (self class) MixAudio(num_frames gd.Int, buffer gd.PackedFloat32Array, offset gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, num_frames)
	callframe.Arg(frame, mmm.Get(buffer))
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VideoStreamPlayback.Bind_mix_audio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVideoStreamPlayback() Expert { return self[0].AsVideoStreamPlayback() }


//go:nosplit
func (self Simple) AsVideoStreamPlayback() Simple { return self[0].AsVideoStreamPlayback() }


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
	case "_stop": return reflect.ValueOf(self._stop);
	case "_play": return reflect.ValueOf(self._play);
	case "_is_playing": return reflect.ValueOf(self._is_playing);
	case "_set_paused": return reflect.ValueOf(self._set_paused);
	case "_is_paused": return reflect.ValueOf(self._is_paused);
	case "_get_length": return reflect.ValueOf(self._get_length);
	case "_get_playback_position": return reflect.ValueOf(self._get_playback_position);
	case "_seek": return reflect.ValueOf(self._seek);
	case "_set_audio_track": return reflect.ValueOf(self._set_audio_track);
	case "_get_texture": return reflect.ValueOf(self._get_texture);
	case "_update": return reflect.ValueOf(self._update);
	case "_get_channels": return reflect.ValueOf(self._get_channels);
	case "_get_mix_rate": return reflect.ValueOf(self._get_mix_rate);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VideoStreamPlayback", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
