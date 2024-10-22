package AudioStreamPlayback

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Can play, loop, pause a scroll through audio. See [AudioStream] and [AudioStreamOggVorbis] for usage.
	// AudioStreamPlayback methods that can be overridden by a [Class] that extends it.
	type AudioStreamPlayback interface {
		//Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
		Start(from_pos float64) 
		//Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
		Stop() 
		//Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
		IsPlaying() bool
		//Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
		GetLoopCount() int
		//Overridable method. Should return the current progress along the audio stream, in seconds.
		GetPlaybackPosition() float64
		//Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
		Seek(position float64) 
		//Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
		//[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
		Mix(buffer *classdb.AudioFrame, rate_scale float64, frames int) int
		//Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
		TagUsedStreams() 
		//Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
		SetParameter(name string, value gd.Variant) 
		//Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
		GetParameter(name string) gd.Variant
	}

*/
type Go [1]classdb.AudioStreamPlayback

/*
Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
*/
func (Go) _start(impl func(ptr unsafe.Pointer, from_pos float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var from_pos = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(from_pos))
		gc.End()
	}
}

/*
Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
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
Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
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
Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
*/
func (Go) _get_loop_count(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable method. Should return the current progress along the audio stream, in seconds.
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
Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
*/
func (Go) _seek(impl func(ptr unsafe.Pointer, position float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var position = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(position))
		gc.End()
	}
}

/*
Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (Go) _mix(impl func(ptr unsafe.Pointer, buffer *classdb.AudioFrame, rate_scale float64, frames int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,0)
		var rate_scale = gd.UnsafeGet[gd.Float](p_args,1)
		var frames = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, buffer, float64(rate_scale), int(frames))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
*/
func (Go) _tag_used_streams(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (Go) _set_parameter(impl func(ptr unsafe.Pointer, name string, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var value = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name.String(), value)
		gc.End()
	}
}

/*
Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (Go) _get_parameter(impl func(ptr unsafe.Pointer, name string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.StringName](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}

/*
Associates [AudioSamplePlayback] to this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
func (self Go) SetSamplePlayback(playback_sample gdclass.AudioSamplePlayback) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSamplePlayback(playback_sample)
}

/*
Returns the [AudioSamplePlayback] associated with this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
func (self Go) GetSamplePlayback() gdclass.AudioSamplePlayback {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioSamplePlayback(class(self).GetSamplePlayback(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamPlayback"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
*/
func (class) _start(impl func(ptr unsafe.Pointer, from_pos gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var from_pos = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, from_pos)
		ctx.End()
	}
}

/*
Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
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
Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
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
Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
*/
func (class) _get_loop_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
Overridable method. Should return the current progress along the audio stream, in seconds.
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
Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
*/
func (class) _seek(impl func(ptr unsafe.Pointer, position gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var position = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, position)
		ctx.End()
	}
}

/*
Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (class) _mix(impl func(ptr unsafe.Pointer, buffer *classdb.AudioFrame, rate_scale gd.Float, frames gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,0)
		var rate_scale = gd.UnsafeGet[gd.Float](p_args,1)
		var frames = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, buffer, rate_scale, frames)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
*/
func (class) _tag_used_streams(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (class) _set_parameter(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var value = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, name, value)
		ctx.End()
	}
}

/*
Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (class) _get_parameter(impl func(ptr unsafe.Pointer, name gd.StringName) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Associates [AudioSamplePlayback] to this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
//go:nosplit
func (self class) SetSamplePlayback(playback_sample gdclass.AudioSamplePlayback)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(playback_sample[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayback.Bind_set_sample_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [AudioSamplePlayback] associated with this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
//go:nosplit
func (self class) GetSamplePlayback(ctx gd.Lifetime) gdclass.AudioSamplePlayback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlayback.Bind_get_sample_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioSamplePlayback
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPlayback() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlayback() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_start": return reflect.ValueOf(self._start);
	case "_stop": return reflect.ValueOf(self._stop);
	case "_is_playing": return reflect.ValueOf(self._is_playing);
	case "_get_loop_count": return reflect.ValueOf(self._get_loop_count);
	case "_get_playback_position": return reflect.ValueOf(self._get_playback_position);
	case "_seek": return reflect.ValueOf(self._seek);
	case "_mix": return reflect.ValueOf(self._mix);
	case "_tag_used_streams": return reflect.ValueOf(self._tag_used_streams);
	case "_set_parameter": return reflect.ValueOf(self._set_parameter);
	case "_get_parameter": return reflect.ValueOf(self._get_parameter);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_start": return reflect.ValueOf(self._start);
	case "_stop": return reflect.ValueOf(self._stop);
	case "_is_playing": return reflect.ValueOf(self._is_playing);
	case "_get_loop_count": return reflect.ValueOf(self._get_loop_count);
	case "_get_playback_position": return reflect.ValueOf(self._get_playback_position);
	case "_seek": return reflect.ValueOf(self._seek);
	case "_mix": return reflect.ValueOf(self._mix);
	case "_tag_used_streams": return reflect.ValueOf(self._tag_used_streams);
	case "_set_parameter": return reflect.ValueOf(self._set_parameter);
	case "_get_parameter": return reflect.ValueOf(self._get_parameter);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("AudioStreamPlayback", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
