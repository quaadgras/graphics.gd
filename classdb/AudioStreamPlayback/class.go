// Package AudioStreamPlayback provides methods for working with AudioStreamPlayback object instances.
package AudioStreamPlayback

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
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
Can play, loop, pause a scroll through audio. See [AudioStream] and [AudioStreamOggVorbis] for usage.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=AudioStreamPlayback)
*/
type Instance [1]gdclass.AudioStreamPlayback

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlayback() Instance
}
type Interface interface {
	//Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
	Start(from_pos Float.X)
	//Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
	Stop()
	//Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
	IsPlaying() bool
	//Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
	GetLoopCount() int
	//Overridable method. Should return the current progress along the audio stream, in seconds.
	GetPlaybackPosition() Float.X
	//Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
	Seek(position Float.X)
	//Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
	//[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
	Mix(buffer *AudioFrame, rate_scale Float.X, frames int) int
	//Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
	TagUsedStreams()
	//Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
	SetParameter(name string, value any)
	//Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
	GetParameter(name string) any
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) Start(from_pos Float.X)                                         { return }
func (self Implementation) Stop()                                                          { return }
func (self Implementation) IsPlaying() (_ bool)                                            { return }
func (self Implementation) GetLoopCount() (_ int)                                          { return }
func (self Implementation) GetPlaybackPosition() (_ Float.X)                               { return }
func (self Implementation) Seek(position Float.X)                                          { return }
func (self Implementation) Mix(buffer *AudioFrame, rate_scale Float.X, frames int) (_ int) { return }
func (self Implementation) TagUsedStreams()                                                { return }
func (self Implementation) SetParameter(name string, value any)                            { return }
func (self Implementation) GetParameter(name string) (_ any)                               { return }

/*
Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
*/
func (Instance) _start(impl func(ptr unsafe.Pointer, from_pos Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_pos = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(from_pos))
	}
}

/*
Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
*/
func (Instance) _stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
*/
func (Instance) _is_playing(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
*/
func (Instance) _get_loop_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable method. Should return the current progress along the audio stream, in seconds.
*/
func (Instance) _get_playback_position(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
*/
func (Instance) _seek(impl func(ptr unsafe.Pointer, position Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var position = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(position))
	}
}

/*
Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (Instance) _mix(impl func(ptr unsafe.Pointer, buffer *AudioFrame, rate_scale Float.X, frames int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var buffer = gd.UnsafeGet[*AudioFrame](p_args, 0)
		var rate_scale = gd.UnsafeGet[gd.Float](p_args, 1)
		var frames = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, buffer, Float.X(rate_scale), int(frames))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
*/
func (Instance) _tag_used_streams(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (Instance) _set_parameter(impl func(ptr unsafe.Pointer, name string, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(name)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 1))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name.String(), value.Interface())
	}
}

/*
Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (Instance) _get_parameter(impl func(ptr unsafe.Pointer, name string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Associates [AudioSamplePlayback] to this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
func (self Instance) SetSamplePlayback(playback_sample [1]gdclass.AudioSamplePlayback) {
	class(self).SetSamplePlayback(playback_sample)
}

/*
Returns the [AudioSamplePlayback] associated with this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
func (self Instance) GetSamplePlayback() [1]gdclass.AudioSamplePlayback {
	return [1]gdclass.AudioSamplePlayback(class(self).GetSamplePlayback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlayback

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlayback"))
	casted := Instance{*(*gdclass.AudioStreamPlayback)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
*/
func (class) _start(impl func(ptr unsafe.Pointer, from_pos gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_pos = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, from_pos)
	}
}

/*
Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
*/
func (class) _stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
*/
func (class) _is_playing(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
*/
func (class) _get_loop_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return the current progress along the audio stream, in seconds.
*/
func (class) _get_playback_position(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
*/
func (class) _seek(impl func(ptr unsafe.Pointer, position gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var position = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, position)
	}
}

/*
Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (class) _mix(impl func(ptr unsafe.Pointer, buffer *AudioFrame, rate_scale gd.Float, frames gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var buffer = gd.UnsafeGet[*AudioFrame](p_args, 0)
		var rate_scale = gd.UnsafeGet[gd.Float](p_args, 1)
		var frames = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, buffer, rate_scale, frames)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
*/
func (class) _tag_used_streams(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (class) _set_parameter(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var value = pointers.New[gd.Variant](gd.UnsafeGet[variantPointers](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name, value)
	}
}

/*
Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (class) _get_parameter(impl func(ptr unsafe.Pointer, name gd.StringName) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Associates [AudioSamplePlayback] to this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
//go:nosplit
func (self class) SetSamplePlayback(playback_sample [1]gdclass.AudioSamplePlayback) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(playback_sample[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayback.Bind_set_sample_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [AudioSamplePlayback] associated with this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
//go:nosplit
func (self class) GetSamplePlayback() [1]gdclass.AudioSamplePlayback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayback.Bind_get_sample_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioSamplePlayback{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioSamplePlayback](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPlayback() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlayback() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_start":
		return reflect.ValueOf(self._start)
	case "_stop":
		return reflect.ValueOf(self._stop)
	case "_is_playing":
		return reflect.ValueOf(self._is_playing)
	case "_get_loop_count":
		return reflect.ValueOf(self._get_loop_count)
	case "_get_playback_position":
		return reflect.ValueOf(self._get_playback_position)
	case "_seek":
		return reflect.ValueOf(self._seek)
	case "_mix":
		return reflect.ValueOf(self._mix)
	case "_tag_used_streams":
		return reflect.ValueOf(self._tag_used_streams)
	case "_set_parameter":
		return reflect.ValueOf(self._set_parameter)
	case "_get_parameter":
		return reflect.ValueOf(self._get_parameter)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_start":
		return reflect.ValueOf(self._start)
	case "_stop":
		return reflect.ValueOf(self._stop)
	case "_is_playing":
		return reflect.ValueOf(self._is_playing)
	case "_get_loop_count":
		return reflect.ValueOf(self._get_loop_count)
	case "_get_playback_position":
		return reflect.ValueOf(self._get_playback_position)
	case "_seek":
		return reflect.ValueOf(self._seek)
	case "_mix":
		return reflect.ValueOf(self._mix)
	case "_tag_used_streams":
		return reflect.ValueOf(self._tag_used_streams)
	case "_set_parameter":
		return reflect.ValueOf(self._set_parameter)
	case "_get_parameter":
		return reflect.ValueOf(self._get_parameter)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlayback", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlayback{*(*gdclass.AudioStreamPlayback)(unsafe.Pointer(&ptr))}
	})
}
