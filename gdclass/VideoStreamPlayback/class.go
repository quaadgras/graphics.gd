package VideoStreamPlayback

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
		GetLength() float64
		//Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
		GetPlaybackPosition() float64
		//Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
		Seek(time float64) 
		//Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
		SetAudioTrack(idx int) 
		//Allocates a [Texture2D] in which decoded video frames will be drawn.
		GetTexture() gdclass.Texture2D
		//Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
		Update(delta float64) 
		//Returns the number of audio channels.
		GetChannels() int
		//Returns the audio sample rate used for mixing.
		GetMixRate() int
	}

*/
type Go [1]classdb.VideoStreamPlayback

/*
Stops playback. May be called multiple times before [method _play], or in response to [method VideoStreamPlayer.stop]. [method _is_playing] should return false once stopped.
*/
func (Go) _stop(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called in response to [member VideoStreamPlayer.autoplay] or [method VideoStreamPlayer.play]. Note that manual playback may also invoke [method _stop] multiple times before this method is called. [method _is_playing] should return true once playing.
*/
func (Go) _play(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Returns the playback state, as determined by calls to [method _play] and [method _stop].
*/
func (Go) _is_playing(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Set the paused status of video playback. [method _is_paused] must return [param paused]. Called in response to the [member VideoStreamPlayer.paused] setter.
*/
func (Go) _set_paused(impl func(ptr unsafe.Pointer, paused bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var paused = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, paused)
		gc.End()
	}
}

/*
Returns the paused status, as set by [method _set_paused].
*/
func (Go) _is_paused(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns the video duration in seconds, if known, or 0 if unknown.
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
Return the current playback timestamp. Called in response to the [member VideoStreamPlayer.stream_position] getter.
*/
func (Go) _get_playback_position(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Seeks to [param time] seconds. Called in response to the [member VideoStreamPlayer.stream_position] setter.
*/
func (Go) _seek(impl func(ptr unsafe.Pointer, time float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var time = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(time))
		gc.End()
	}
}

/*
Select the audio track [param idx]. Called when playback starts, and in response to the [member VideoStreamPlayer.audio_track] setter.
*/
func (Go) _set_audio_track(impl func(ptr unsafe.Pointer, idx int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(idx))
		gc.End()
	}
}

/*
Allocates a [Texture2D] in which decoded video frames will be drawn.
*/
func (Go) _get_texture(impl func(ptr unsafe.Pointer) gdclass.Texture2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Ticks video playback for [param delta] seconds. Called every frame as long as [method _is_paused] and [method _is_playing] return true.
*/
func (Go) _update(impl func(ptr unsafe.Pointer, delta float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var delta = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(delta))
		gc.End()
	}
}

/*
Returns the number of audio channels.
*/
func (Go) _get_channels(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Returns the audio sample rate used for mixing.
*/
func (Go) _get_mix_rate(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Render [param num_frames] audio frames (of [method _get_channels] floats each) from [param buffer], starting from index [param offset] in the array. Returns the number of audio frames rendered, or -1 on error.
*/
func (self Go) MixAudio(num_frames int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).MixAudio(gd.Int(num_frames), gc.PackedFloat32Slice(([1][]float32{}[0])), gd.Int(0))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VideoStreamPlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("VideoStreamPlayback"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (class) _get_texture(impl func(ptr unsafe.Pointer) gdclass.Texture2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (self class) AsVideoStreamPlayback() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVideoStreamPlayback() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
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
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("VideoStreamPlayback", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
