package AudioStreamSynchronized

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStream"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This is a stream that can be fitted with sub-streams, which will be played in-sync. The streams being at exactly the same time when play is pressed, and will end when the last of them ends. If one of the sub-streams loops, then playback will continue.

*/
type Go [1]classdb.AudioStreamSynchronized

/*
Set one of the synchronized streams, by index.
*/
func (self Go) SetSyncStream(stream_index int, audio_stream gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSyncStream(gd.Int(stream_index), audio_stream)
}

/*
Get one of the synchronized streams, by index.
*/
func (self Go) GetSyncStream(stream_index int) gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStream(class(self).GetSyncStream(gc, gd.Int(stream_index)))
}

/*
Set the volume of one of the synchronized streams, by index.
*/
func (self Go) SetSyncStreamVolume(stream_index int, volume_db float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSyncStreamVolume(gd.Int(stream_index), gd.Float(volume_db))
}

/*
Get the volume of one of the synchronized streams, by index.
*/
func (self Go) GetSyncStreamVolume(stream_index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetSyncStreamVolume(gd.Int(stream_index))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamSynchronized
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamSynchronized"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) StreamCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetStreamCount()))
}

func (self Go) SetStreamCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStreamCount(gd.Int(value))
}

//go:nosplit
func (self class) SetStreamCount(stream_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) SetSyncStream(stream_index gd.Int, audio_stream gdclass.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, mmm.Get(audio_stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_set_sync_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Get one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) GetSyncStream(ctx gd.Lifetime, stream_index gd.Int) gdclass.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_get_sync_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Set the volume of one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) SetSyncStreamVolume(stream_index gd.Int, volume_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_set_sync_stream_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Get the volume of one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) GetSyncStreamVolume(stream_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_get_sync_stream_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamSynchronized() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamSynchronized() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.GD { return *((*AudioStream.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStream() AudioStream.Go { return *((*AudioStream.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {classdb.Register("AudioStreamSynchronized", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
