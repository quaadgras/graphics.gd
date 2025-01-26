// Package AudioStreamSynchronized provides methods for working with AudioStreamSynchronized object instances.
package AudioStreamSynchronized

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
This is a stream that can be fitted with sub-streams, which will be played in-sync. The streams being at exactly the same time when play is pressed, and will end when the last of them ends. If one of the sub-streams loops, then playback will continue.
*/
type Instance [1]gdclass.AudioStreamSynchronized

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamSynchronized() Instance
}

/*
Set one of the synchronized streams, by index.
*/
func (self Instance) SetSyncStream(stream_index int, audio_stream [1]gdclass.AudioStream) { //gd:AudioStreamSynchronized.set_sync_stream
	class(self).SetSyncStream(gd.Int(stream_index), audio_stream)
}

/*
Get one of the synchronized streams, by index.
*/
func (self Instance) GetSyncStream(stream_index int) [1]gdclass.AudioStream { //gd:AudioStreamSynchronized.get_sync_stream
	return [1]gdclass.AudioStream(class(self).GetSyncStream(gd.Int(stream_index)))
}

/*
Set the volume of one of the synchronized streams, by index.
*/
func (self Instance) SetSyncStreamVolume(stream_index int, volume_db Float.X) { //gd:AudioStreamSynchronized.set_sync_stream_volume
	class(self).SetSyncStreamVolume(gd.Int(stream_index), gd.Float(volume_db))
}

/*
Get the volume of one of the synchronized streams, by index.
*/
func (self Instance) GetSyncStreamVolume(stream_index int) Float.X { //gd:AudioStreamSynchronized.get_sync_stream_volume
	return Float.X(Float.X(class(self).GetSyncStreamVolume(gd.Int(stream_index))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamSynchronized

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamSynchronized"))
	casted := Instance{*(*gdclass.AudioStreamSynchronized)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) StreamCount() int {
	return int(int(class(self).GetStreamCount()))
}

func (self Instance) SetStreamCount(value int) {
	class(self).SetStreamCount(gd.Int(value))
}

//go:nosplit
func (self class) SetStreamCount(stream_count gd.Int) { //gd:AudioStreamSynchronized.set_stream_count
	var frame = callframe.New()
	callframe.Arg(frame, stream_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamCount() gd.Int { //gd:AudioStreamSynchronized.get_stream_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) SetSyncStream(stream_index gd.Int, audio_stream [1]gdclass.AudioStream) { //gd:AudioStreamSynchronized.set_sync_stream
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, pointers.Get(audio_stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_set_sync_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Get one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) GetSyncStream(stream_index gd.Int) [1]gdclass.AudioStream { //gd:AudioStreamSynchronized.get_sync_stream
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_get_sync_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Set the volume of one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) SetSyncStreamVolume(stream_index gd.Int, volume_db gd.Float) { //gd:AudioStreamSynchronized.set_sync_stream_volume
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, volume_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_set_sync_stream_volume, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Get the volume of one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) GetSyncStreamVolume(stream_index gd.Int) gd.Float { //gd:AudioStreamSynchronized.get_sync_stream_volume
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_get_sync_stream_volume, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamSynchronized() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamSynchronized() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStream() AudioStream.Advanced {
	return *((*AudioStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStream() AudioStream.Instance {
	return *((*AudioStream.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioStream.Advanced(self.AsAudioStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Instance(self.AsAudioStream()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamSynchronized", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamSynchronized{*(*gdclass.AudioStreamSynchronized)(unsafe.Pointer(&ptr))}
	})
}
