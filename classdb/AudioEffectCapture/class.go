// Package AudioEffectCapture provides methods for working with AudioEffectCapture object instances.
package AudioEffectCapture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
AudioEffectCapture is an AudioEffect which copies all audio frames from the attached audio effect bus into its internal ring buffer.
Application code should consume these audio frames from this ring buffer using [method get_buffer] and process it as needed, for example to capture data from an [AudioStreamMicrophone], implement application-defined effects, or to transmit audio over the network. When capturing audio data from a microphone, the format of the samples will be stereo 32-bit floating-point PCM.
Unlike [AudioEffectRecord], this effect only returns the raw audio samples instead of encoding them into an [AudioStream].
*/
type Instance [1]gdclass.AudioEffectCapture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectCapture() Instance
}

/*
Returns [code]true[/code] if at least [param frames] audio frames are available to read in the internal ring buffer.
*/
func (self Instance) CanGetBuffer(frames int) bool {
	return bool(class(self).CanGetBuffer(gd.Int(frames)))
}

/*
Gets the next [param frames] audio samples from the internal ring buffer.
Returns a [PackedVector2Array] containing exactly [param frames] audio samples if available, or an empty [PackedVector2Array] if insufficient data was available.
The samples are signed floating-point PCM between [code]-1[/code] and [code]1[/code]. You will have to scale them if you want to use them as 8 or 16-bit integer samples. ([code]v = 0x7fff * samples[0].x[/code])
*/
func (self Instance) GetBuffer(frames int) []Vector2.XY {
	return []Vector2.XY(class(self).GetBuffer(gd.Int(frames)).AsSlice())
}

/*
Clears the internal ring buffer.
[b]Note:[/b] Calling this during a capture can cause the loss of samples which causes popping in the playback.
*/
func (self Instance) ClearBuffer() {
	class(self).ClearBuffer()
}

/*
Returns the number of frames available to read using [method get_buffer].
*/
func (self Instance) GetFramesAvailable() int {
	return int(int(class(self).GetFramesAvailable()))
}

/*
Returns the number of audio frames discarded from the audio bus due to full buffer.
*/
func (self Instance) GetDiscardedFrames() int {
	return int(int(class(self).GetDiscardedFrames()))
}

/*
Returns the total size of the internal ring buffer in frames.
*/
func (self Instance) GetBufferLengthFrames() int {
	return int(int(class(self).GetBufferLengthFrames()))
}

/*
Returns the number of audio frames inserted from the audio bus.
*/
func (self Instance) GetPushedFrames() int {
	return int(int(class(self).GetPushedFrames()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectCapture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectCapture"))
	casted := Instance{*(*gdclass.AudioEffectCapture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BufferLength() Float.X {
	return Float.X(Float.X(class(self).GetBufferLength()))
}

func (self Instance) SetBufferLength(value Float.X) {
	class(self).SetBufferLength(gd.Float(value))
}

/*
Returns [code]true[/code] if at least [param frames] audio frames are available to read in the internal ring buffer.
*/
//go:nosplit
func (self class) CanGetBuffer(frames gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_can_get_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the next [param frames] audio samples from the internal ring buffer.
Returns a [PackedVector2Array] containing exactly [param frames] audio samples if available, or an empty [PackedVector2Array] if insufficient data was available.
The samples are signed floating-point PCM between [code]-1[/code] and [code]1[/code]. You will have to scale them if you want to use them as 8 or 16-bit integer samples. ([code]v = 0x7fff * samples[0].x[/code])
*/
//go:nosplit
func (self class) GetBuffer(frames gd.Int) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the internal ring buffer.
[b]Note:[/b] Calling this during a capture can cause the loss of samples which causes popping in the playback.
*/
//go:nosplit
func (self class) ClearBuffer() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_clear_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetBufferLength(buffer_length_seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_length_seconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of frames available to read using [method get_buffer].
*/
//go:nosplit
func (self class) GetFramesAvailable() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_frames_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of audio frames discarded from the audio bus due to full buffer.
*/
//go:nosplit
func (self class) GetDiscardedFrames() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_discarded_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total size of the internal ring buffer in frames.
*/
//go:nosplit
func (self class) GetBufferLengthFrames() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_buffer_length_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of audio frames inserted from the audio bus.
*/
//go:nosplit
func (self class) GetPushedFrames() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_pushed_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectCapture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectCapture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.Advanced {
	return *((*AudioEffect.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffect() AudioEffect.Instance {
	return *((*AudioEffect.Instance)(unsafe.Pointer(&self)))
}
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
	default:
		return gd.VirtualByName(AudioEffect.Advanced(self.AsAudioEffect()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffect.Instance(self.AsAudioEffect()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectCapture", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectCapture{*(*gdclass.AudioEffectCapture)(unsafe.Pointer(&ptr))}
	})
}
