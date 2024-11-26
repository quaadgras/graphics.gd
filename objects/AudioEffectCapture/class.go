package AudioEffectCapture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioEffect"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
AudioEffectCapture is an AudioEffect which copies all audio frames from the attached audio effect bus into its internal ring buffer.
Application code should consume these audio frames from this ring buffer using [method get_buffer] and process it as needed, for example to capture data from an [AudioStreamMicrophone], implement application-defined effects, or to transmit audio over the network. When capturing audio data from a microphone, the format of the samples will be stereo 32-bit floating-point PCM.
Unlike [AudioEffectRecord], this effect only returns the raw audio samples instead of encoding them into an [AudioStream].
*/
type Instance [1]classdb.AudioEffectCapture

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
type class [1]classdb.AudioEffectCapture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectCapture"))
	return Instance{classdb.AudioEffectCapture(object)}
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_can_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_clear_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetBufferLength(buffer_length_seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_length_seconds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_frames_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_discarded_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_buffer_length_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCapture.Bind_get_pushed_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {
	classdb.Register("AudioEffectCapture", func(ptr gd.Object) any { return classdb.AudioEffectCapture(ptr) })
}
