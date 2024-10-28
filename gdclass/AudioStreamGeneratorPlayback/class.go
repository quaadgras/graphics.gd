package AudioStreamGeneratorPlayback

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStreamPlaybackResampled"
import "grow.graphics/gd/gdclass/AudioStreamPlayback"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class is meant to be used with [AudioStreamGenerator] to play back the generated audio in real-time.

*/
type Go [1]classdb.AudioStreamGeneratorPlayback

/*
Pushes a single audio data frame to the buffer. This is usually less efficient than [method push_buffer] in C# and compiled languages via GDExtension, but [method push_frame] may be [i]more[/i] efficient in GDScript.
*/
func (self Go) PushFrame(frame_ gd.Vector2) bool {
	return bool(class(self).PushFrame(frame_))
}

/*
Returns [code]true[/code] if a buffer of the size [param amount] can be pushed to the audio sample data buffer without overflowing it, [code]false[/code] otherwise.
*/
func (self Go) CanPushBuffer(amount int) bool {
	return bool(class(self).CanPushBuffer(gd.Int(amount)))
}

/*
Pushes several audio data frames to the buffer. This is usually more efficient than [method push_frame] in C# and compiled languages via GDExtension, but [method push_buffer] may be [i]less[/i] efficient in GDScript.
*/
func (self Go) PushBuffer(frames []gd.Vector2) bool {
	return bool(class(self).PushBuffer(gd.NewPackedVector2Slice(frames)))
}

/*
Returns the number of frames that can be pushed to the audio sample data buffer without overflowing it. If the result is [code]0[/code], the buffer is full.
*/
func (self Go) GetFramesAvailable() int {
	return int(int(class(self).GetFramesAvailable()))
}

/*
Returns the number of times the playback skipped due to a buffer underrun in the audio sample data. This value is reset at the start of the playback.
*/
func (self Go) GetSkips() int {
	return int(int(class(self).GetSkips()))
}

/*
Clears the audio sample data buffer.
*/
func (self Go) ClearBuffer() {
	class(self).ClearBuffer()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamGeneratorPlayback
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamGeneratorPlayback"))
	return Go{classdb.AudioStreamGeneratorPlayback(object)}
}

/*
Pushes a single audio data frame to the buffer. This is usually less efficient than [method push_buffer] in C# and compiled languages via GDExtension, but [method push_frame] may be [i]more[/i] efficient in GDScript.
*/
//go:nosplit
func (self class) PushFrame(frame_ gd.Vector2) bool {
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_push_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a buffer of the size [param amount] can be pushed to the audio sample data buffer without overflowing it, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) CanPushBuffer(amount gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_can_push_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Pushes several audio data frames to the buffer. This is usually more efficient than [method push_frame] in C# and compiled languages via GDExtension, but [method push_buffer] may be [i]less[/i] efficient in GDScript.
*/
//go:nosplit
func (self class) PushBuffer(frames gd.PackedVector2Array) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(frames))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_push_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of frames that can be pushed to the audio sample data buffer without overflowing it. If the result is [code]0[/code], the buffer is full.
*/
//go:nosplit
func (self class) GetFramesAvailable() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_get_frames_available, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of times the playback skipped due to a buffer underrun in the audio sample data. This value is reset at the start of the playback.
*/
//go:nosplit
func (self class) GetSkips() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_get_skips, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears the audio sample data buffer.
*/
//go:nosplit
func (self class) ClearBuffer()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_clear_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsAudioStreamGeneratorPlayback() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamGeneratorPlayback() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.GD { return *((*AudioStreamPlaybackResampled.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Go { return *((*AudioStreamPlaybackResampled.Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.GD { return *((*AudioStreamPlayback.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlayback() AudioStreamPlayback.Go { return *((*AudioStreamPlayback.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStreamPlaybackResampled(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStreamPlaybackResampled(), name)
	}
}
func init() {classdb.Register("AudioStreamGeneratorPlayback", func(ptr gd.Object) any { return classdb.AudioStreamGeneratorPlayback(ptr) })}
