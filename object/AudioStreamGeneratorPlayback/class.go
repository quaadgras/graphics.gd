package AudioStreamGeneratorPlayback

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioStreamPlaybackResampled"
import "grow.graphics/gd/object/AudioStreamPlayback"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class is meant to be used with [AudioStreamGenerator] to play back the generated audio in real-time.

*/
type Simple [1]classdb.AudioStreamGeneratorPlayback
func (self Simple) PushFrame(frame_ gd.Vector2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).PushFrame(frame_))
}
func (self Simple) CanPushBuffer(amount int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanPushBuffer(gd.Int(amount)))
}
func (self Simple) PushBuffer(frames gd.PackedVector2Array) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).PushBuffer(frames))
}
func (self Simple) GetFramesAvailable() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFramesAvailable()))
}
func (self Simple) GetSkips() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSkips()))
}
func (self Simple) ClearBuffer() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearBuffer()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamGeneratorPlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Pushes a single audio data frame to the buffer. This is usually less efficient than [method push_buffer] in C# and compiled languages via GDExtension, but [method push_frame] may be [i]more[/i] efficient in GDScript.
*/
//go:nosplit
func (self class) PushFrame(frame_ gd.Vector2) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGeneratorPlayback.Bind_push_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a buffer of the size [param amount] can be pushed to the audio sample data buffer without overflowing it, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) CanPushBuffer(amount gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGeneratorPlayback.Bind_can_push_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Pushes several audio data frames to the buffer. This is usually more efficient than [method push_frame] in C# and compiled languages via GDExtension, but [method push_buffer] may be [i]less[/i] efficient in GDScript.
*/
//go:nosplit
func (self class) PushBuffer(frames gd.PackedVector2Array) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(frames))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGeneratorPlayback.Bind_push_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of frames that can be pushed to the audio sample data buffer without overflowing it. If the result is [code]0[/code], the buffer is full.
*/
//go:nosplit
func (self class) GetFramesAvailable() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGeneratorPlayback.Bind_get_frames_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of times the playback skipped due to a buffer underrun in the audio sample data. This value is reset at the start of the playback.
*/
//go:nosplit
func (self class) GetSkips() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGeneratorPlayback.Bind_get_skips, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears the audio sample data buffer.
*/
//go:nosplit
func (self class) ClearBuffer()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamGeneratorPlayback.Bind_clear_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsAudioStreamGeneratorPlayback() Expert { return self[0].AsAudioStreamGeneratorPlayback() }


//go:nosplit
func (self Simple) AsAudioStreamGeneratorPlayback() Simple { return self[0].AsAudioStreamGeneratorPlayback() }


//go:nosplit
func (self class) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Expert { return self[0].AsAudioStreamPlaybackResampled() }


//go:nosplit
func (self Simple) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Simple { return self[0].AsAudioStreamPlaybackResampled() }


//go:nosplit
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Expert { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self Simple) AsAudioStreamPlayback() AudioStreamPlayback.Simple { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStreamGeneratorPlayback", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
