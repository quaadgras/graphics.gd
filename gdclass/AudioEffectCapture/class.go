package AudioEffectCapture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
AudioEffectCapture is an AudioEffect which copies all audio frames from the attached audio effect bus into its internal ring buffer.
Application code should consume these audio frames from this ring buffer using [method get_buffer] and process it as needed, for example to capture data from an [AudioStreamMicrophone], implement application-defined effects, or to transmit audio over the network. When capturing audio data from a microphone, the format of the samples will be stereo 32-bit floating-point PCM.
Unlike [AudioEffectRecord], this effect only returns the raw audio samples instead of encoding them into an [AudioStream].

*/
type Go [1]classdb.AudioEffectCapture

/*
Returns [code]true[/code] if at least [param frames] audio frames are available to read in the internal ring buffer.
*/
func (self Go) CanGetBuffer(frames int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).CanGetBuffer(gd.Int(frames)))
}

/*
Gets the next [param frames] audio samples from the internal ring buffer.
Returns a [PackedVector2Array] containing exactly [param frames] audio samples if available, or an empty [PackedVector2Array] if insufficient data was available.
The samples are signed floating-point PCM between [code]-1[/code] and [code]1[/code]. You will have to scale them if you want to use them as 8 or 16-bit integer samples. ([code]v = 0x7fff * samples[0].x[/code])
*/
func (self Go) GetBuffer(frames int) []gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector2(class(self).GetBuffer(gc, gd.Int(frames)).AsSlice())
}

/*
Clears the internal ring buffer.
[b]Note:[/b] Calling this during a capture can cause the loss of samples which causes popping in the playback.
*/
func (self Go) ClearBuffer() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearBuffer()
}

/*
Returns the number of frames available to read using [method get_buffer].
*/
func (self Go) GetFramesAvailable() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFramesAvailable()))
}

/*
Returns the number of audio frames discarded from the audio bus due to full buffer.
*/
func (self Go) GetDiscardedFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetDiscardedFrames()))
}

/*
Returns the total size of the internal ring buffer in frames.
*/
func (self Go) GetBufferLengthFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBufferLengthFrames()))
}

/*
Returns the number of audio frames inserted from the audio bus.
*/
func (self Go) GetPushedFrames() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPushedFrames()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectCapture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectCapture"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BufferLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBufferLength()))
}

func (self Go) SetBufferLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBufferLength(gd.Float(value))
}

/*
Returns [code]true[/code] if at least [param frames] audio frames are available to read in the internal ring buffer.
*/
//go:nosplit
func (self class) CanGetBuffer(frames gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_can_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GetBuffer(ctx gd.Lifetime, frames gd.Int) gd.PackedVector2Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector2Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears the internal ring buffer.
[b]Note:[/b] Calling this during a capture can cause the loss of samples which causes popping in the playback.
*/
//go:nosplit
func (self class) ClearBuffer()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_clear_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBufferLength(buffer_length_seconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer_length_seconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBufferLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of frames available to read using [method get_buffer].
*/
//go:nosplit
func (self class) GetFramesAvailable() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_get_frames_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of audio frames discarded from the audio bus due to full buffer.
*/
//go:nosplit
func (self class) GetDiscardedFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_get_discarded_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total size of the internal ring buffer in frames.
*/
//go:nosplit
func (self class) GetBufferLengthFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_get_buffer_length_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of audio frames inserted from the audio bus.
*/
//go:nosplit
func (self class) GetPushedFrames() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectCapture.Bind_get_pushed_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectCapture() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectCapture() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.GD { return *((*AudioEffect.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() AudioEffect.Go { return *((*AudioEffect.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {classdb.Register("AudioEffectCapture", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
