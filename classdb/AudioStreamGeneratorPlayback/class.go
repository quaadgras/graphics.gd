// Package AudioStreamGeneratorPlayback provides methods for working with AudioStreamGeneratorPlayback object instances.
package AudioStreamGeneratorPlayback

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
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/AudioStreamPlaybackResampled"
import "graphics.gd/classdb/AudioStreamPlayback"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This class is meant to be used with [AudioStreamGenerator] to play back the generated audio in real-time.
*/
type Instance [1]gdclass.AudioStreamGeneratorPlayback

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamGeneratorPlayback() Instance
}

/*
Pushes a single audio data frame to the buffer. This is usually less efficient than [method push_buffer] in C# and compiled languages via GDExtension, but [method push_frame] may be [i]more[/i] efficient in GDScript.
*/
func (self Instance) PushFrame(frame_ Vector2.XY) bool { //gd:AudioStreamGeneratorPlayback.push_frame
	return bool(class(self).PushFrame(gd.Vector2(frame_)))
}

/*
Returns [code]true[/code] if a buffer of the size [param amount] can be pushed to the audio sample data buffer without overflowing it, [code]false[/code] otherwise.
*/
func (self Instance) CanPushBuffer(amount int) bool { //gd:AudioStreamGeneratorPlayback.can_push_buffer
	return bool(class(self).CanPushBuffer(gd.Int(amount)))
}

/*
Pushes several audio data frames to the buffer. This is usually more efficient than [method push_frame] in C# and compiled languages via GDExtension, but [method push_buffer] may be [i]less[/i] efficient in GDScript.
*/
func (self Instance) PushBuffer(frames []Vector2.XY) bool { //gd:AudioStreamGeneratorPlayback.push_buffer
	return bool(class(self).PushBuffer(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&frames)))))
}

/*
Returns the number of frames that can be pushed to the audio sample data buffer without overflowing it. If the result is [code]0[/code], the buffer is full.
*/
func (self Instance) GetFramesAvailable() int { //gd:AudioStreamGeneratorPlayback.get_frames_available
	return int(int(class(self).GetFramesAvailable()))
}

/*
Returns the number of times the playback skipped due to a buffer underrun in the audio sample data. This value is reset at the start of the playback.
*/
func (self Instance) GetSkips() int { //gd:AudioStreamGeneratorPlayback.get_skips
	return int(int(class(self).GetSkips()))
}

/*
Clears the audio sample data buffer.
*/
func (self Instance) ClearBuffer() { //gd:AudioStreamGeneratorPlayback.clear_buffer
	class(self).ClearBuffer()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamGeneratorPlayback

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamGeneratorPlayback"))
	casted := Instance{*(*gdclass.AudioStreamGeneratorPlayback)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Pushes a single audio data frame to the buffer. This is usually less efficient than [method push_buffer] in C# and compiled languages via GDExtension, but [method push_frame] may be [i]more[/i] efficient in GDScript.
*/
//go:nosplit
func (self class) PushFrame(frame_ gd.Vector2) bool { //gd:AudioStreamGeneratorPlayback.push_frame
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_push_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a buffer of the size [param amount] can be pushed to the audio sample data buffer without overflowing it, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) CanPushBuffer(amount gd.Int) bool { //gd:AudioStreamGeneratorPlayback.can_push_buffer
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_can_push_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Pushes several audio data frames to the buffer. This is usually more efficient than [method push_frame] in C# and compiled languages via GDExtension, but [method push_buffer] may be [i]less[/i] efficient in GDScript.
*/
//go:nosplit
func (self class) PushBuffer(frames gd.PackedVector2Array) bool { //gd:AudioStreamGeneratorPlayback.push_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(frames))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_push_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of frames that can be pushed to the audio sample data buffer without overflowing it. If the result is [code]0[/code], the buffer is full.
*/
//go:nosplit
func (self class) GetFramesAvailable() gd.Int { //gd:AudioStreamGeneratorPlayback.get_frames_available
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_get_frames_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of times the playback skipped due to a buffer underrun in the audio sample data. This value is reset at the start of the playback.
*/
//go:nosplit
func (self class) GetSkips() gd.Int { //gd:AudioStreamGeneratorPlayback.get_skips
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_get_skips, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears the audio sample data buffer.
*/
//go:nosplit
func (self class) ClearBuffer() { //gd:AudioStreamGeneratorPlayback.clear_buffer
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamGeneratorPlayback.Bind_clear_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsAudioStreamGeneratorPlayback() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamGeneratorPlayback() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Advanced {
	return *((*AudioStreamPlaybackResampled.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Instance {
	return *((*AudioStreamPlaybackResampled.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Advanced {
	return *((*AudioStreamPlayback.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlayback() AudioStreamPlayback.Instance {
	return *((*AudioStreamPlayback.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioStreamPlaybackResampled.Advanced(self.AsAudioStreamPlaybackResampled()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStreamPlaybackResampled.Instance(self.AsAudioStreamPlaybackResampled()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamGeneratorPlayback", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamGeneratorPlayback{*(*gdclass.AudioStreamGeneratorPlayback)(unsafe.Pointer(&ptr))}
	})
}
