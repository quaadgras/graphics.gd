package AudioStreamRandomizer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
Picks a random AudioStream from the pool, depending on the playback mode, and applies random pitch shifting and volume shifting during playback.

*/
type Simple [1]classdb.AudioStreamRandomizer
func (self Simple) AddStream(index int, stream [1]classdb.AudioStream, weight float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddStream(gd.Int(index), stream, gd.Float(weight))
}
func (self Simple) MoveStream(index_from int, index_to int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveStream(gd.Int(index_from), gd.Int(index_to))
}
func (self Simple) RemoveStream(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveStream(gd.Int(index))
}
func (self Simple) SetStream(index int, stream [1]classdb.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStream(gd.Int(index), stream)
}
func (self Simple) GetStream(index int) [1]classdb.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStream(Expert(self).GetStream(gc, gd.Int(index)))
}
func (self Simple) SetStreamProbabilityWeight(index int, weight float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamProbabilityWeight(gd.Int(index), gd.Float(weight))
}
func (self Simple) GetStreamProbabilityWeight(index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetStreamProbabilityWeight(gd.Int(index))))
}
func (self Simple) SetStreamsCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamsCount(gd.Int(count))
}
func (self Simple) GetStreamsCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStreamsCount()))
}
func (self Simple) SetRandomPitch(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRandomPitch(gd.Float(scale))
}
func (self Simple) GetRandomPitch() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRandomPitch()))
}
func (self Simple) SetRandomVolumeOffsetDb(db_offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRandomVolumeOffsetDb(gd.Float(db_offset))
}
func (self Simple) GetRandomVolumeOffsetDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRandomVolumeOffsetDb()))
}
func (self Simple) SetPlaybackMode(mode classdb.AudioStreamRandomizerPlaybackMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlaybackMode(mode)
}
func (self Simple) GetPlaybackMode() classdb.AudioStreamRandomizerPlaybackMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AudioStreamRandomizerPlaybackMode(Expert(self).GetPlaybackMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamRandomizer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Insert a stream at the specified index. If the index is less than zero, the insertion occurs at the end of the underlying pool.
*/
//go:nosplit
func (self class) AddStream(index gd.Int, stream object.AudioStream, weight gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_add_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Move a stream from one index to another.
*/
//go:nosplit
func (self class) MoveStream(index_from gd.Int, index_to gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index_from)
	callframe.Arg(frame, index_to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_move_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove the stream at the specified index.
*/
//go:nosplit
func (self class) RemoveStream(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_remove_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the AudioStream at the specified index.
*/
//go:nosplit
func (self class) SetStream(index gd.Int, stream object.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stream at the specified index.
*/
//go:nosplit
func (self class) GetStream(ctx gd.Lifetime, index gd.Int) object.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Set the probability weight of the stream at the specified index. The higher this value, the more likely that the randomizer will choose this stream during random playback modes.
*/
//go:nosplit
func (self class) SetStreamProbabilityWeight(index gd.Int, weight gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_set_stream_probability_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the probability weight associated with the stream at the given index.
*/
//go:nosplit
func (self class) GetStreamProbabilityWeight(index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_get_stream_probability_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStreamsCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_set_streams_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamsCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_get_streams_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRandomPitch(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_set_random_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRandomPitch() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_get_random_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRandomVolumeOffsetDb(db_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, db_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_set_random_volume_offset_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRandomVolumeOffsetDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_get_random_volume_offset_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaybackMode(mode classdb.AudioStreamRandomizerPlaybackMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_set_playback_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaybackMode() classdb.AudioStreamRandomizerPlaybackMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamRandomizerPlaybackMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamRandomizer.Bind_get_playback_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAudioStreamRandomizer() Expert { return self[0].AsAudioStreamRandomizer() }


//go:nosplit
func (self Simple) AsAudioStreamRandomizer() Simple { return self[0].AsAudioStreamRandomizer() }


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
func init() {classdb.Register("AudioStreamRandomizer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type PlaybackMode = classdb.AudioStreamRandomizerPlaybackMode

const (
/*Pick a stream at random according to the probability weights chosen for each stream, but avoid playing the same stream twice in a row whenever possible. If only 1 sound is present in the pool, the same sound will always play, effectively allowing repeats to occur.*/
	PlaybackRandomNoRepeats PlaybackMode = 0
/*Pick a stream at random according to the probability weights chosen for each stream. If only 1 sound is present in the pool, the same sound will always play.*/
	PlaybackRandom PlaybackMode = 1
/*Play streams in the order they appear in the stream pool. If only 1 sound is present in the pool, the same sound will always play.*/
	PlaybackSequential PlaybackMode = 2
)
