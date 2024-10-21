package AudioStreamSynchronized

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioStream"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This is a stream that can be fitted with sub-streams, which will be played in-sync. The streams being at exactly the same time when play is pressed, and will end when the last of them ends. If one of the sub-streams loops, then playback will continue.

*/
type Simple [1]classdb.AudioStreamSynchronized
func (self Simple) SetStreamCount(stream_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamCount(gd.Int(stream_count))
}
func (self Simple) GetStreamCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStreamCount()))
}
func (self Simple) SetSyncStream(stream_index int, audio_stream [1]classdb.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSyncStream(gd.Int(stream_index), audio_stream)
}
func (self Simple) GetSyncStream(stream_index int) [1]classdb.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStream(Expert(self).GetSyncStream(gc, gd.Int(stream_index)))
}
func (self Simple) SetSyncStreamVolume(stream_index int, volume_db float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSyncStreamVolume(gd.Int(stream_index), gd.Float(volume_db))
}
func (self Simple) GetSyncStreamVolume(stream_index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSyncStreamVolume(gd.Int(stream_index))))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamSynchronized
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) SetSyncStream(stream_index gd.Int, audio_stream object.AudioStream)  {
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
func (self class) GetSyncStream(ctx gd.Lifetime, stream_index gd.Int) object.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamSynchronized.Bind_get_sync_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStream
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

//go:nosplit
func (self class) AsAudioStreamSynchronized() Expert { return self[0].AsAudioStreamSynchronized() }


//go:nosplit
func (self Simple) AsAudioStreamSynchronized() Simple { return self[0].AsAudioStreamSynchronized() }


//go:nosplit
func (self class) AsAudioStream() AudioStream.Expert { return self[0].AsAudioStream() }


//go:nosplit
func (self Simple) AsAudioStream() AudioStream.Simple { return self[0].AsAudioStream() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStreamSynchronized", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
