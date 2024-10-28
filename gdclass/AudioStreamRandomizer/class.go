package AudioStreamRandomizer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
Picks a random AudioStream from the pool, depending on the playback mode, and applies random pitch shifting and volume shifting during playback.

*/
type Go [1]classdb.AudioStreamRandomizer

/*
Insert a stream at the specified index. If the index is less than zero, the insertion occurs at the end of the underlying pool.
*/
func (self Go) AddStream(index int, stream gdclass.AudioStream) {
	class(self).AddStream(gd.Int(index), stream, gd.Float(1.0))
}

/*
Move a stream from one index to another.
*/
func (self Go) MoveStream(index_from int, index_to int) {
	class(self).MoveStream(gd.Int(index_from), gd.Int(index_to))
}

/*
Remove the stream at the specified index.
*/
func (self Go) RemoveStream(index int) {
	class(self).RemoveStream(gd.Int(index))
}

/*
Set the AudioStream at the specified index.
*/
func (self Go) SetStream(index int, stream gdclass.AudioStream) {
	class(self).SetStream(gd.Int(index), stream)
}

/*
Returns the stream at the specified index.
*/
func (self Go) GetStream(index int) gdclass.AudioStream {
	return gdclass.AudioStream(class(self).GetStream(gd.Int(index)))
}

/*
Set the probability weight of the stream at the specified index. The higher this value, the more likely that the randomizer will choose this stream during random playback modes.
*/
func (self Go) SetStreamProbabilityWeight(index int, weight float64) {
	class(self).SetStreamProbabilityWeight(gd.Int(index), gd.Float(weight))
}

/*
Returns the probability weight associated with the stream at the given index.
*/
func (self Go) GetStreamProbabilityWeight(index int) float64 {
	return float64(float64(class(self).GetStreamProbabilityWeight(gd.Int(index))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamRandomizer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamRandomizer"))
	return Go{classdb.AudioStreamRandomizer(object)}
}

func (self Go) PlaybackMode() classdb.AudioStreamRandomizerPlaybackMode {
		return classdb.AudioStreamRandomizerPlaybackMode(class(self).GetPlaybackMode())
}

func (self Go) SetPlaybackMode(value classdb.AudioStreamRandomizerPlaybackMode) {
	class(self).SetPlaybackMode(value)
}

func (self Go) RandomPitch() float64 {
		return float64(float64(class(self).GetRandomPitch()))
}

func (self Go) SetRandomPitch(value float64) {
	class(self).SetRandomPitch(gd.Float(value))
}

func (self Go) RandomVolumeOffsetDb() float64 {
		return float64(float64(class(self).GetRandomVolumeOffsetDb()))
}

func (self Go) SetRandomVolumeOffsetDb(value float64) {
	class(self).SetRandomVolumeOffsetDb(gd.Float(value))
}

func (self Go) StreamsCount() int {
		return int(int(class(self).GetStreamsCount()))
}

func (self Go) SetStreamsCount(value int) {
	class(self).SetStreamsCount(gd.Int(value))
}

/*
Insert a stream at the specified index. If the index is less than zero, the insertion occurs at the end of the underlying pool.
*/
//go:nosplit
func (self class) AddStream(index gd.Int, stream gdclass.AudioStream, weight gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, discreet.Get(stream[0])[0])
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_add_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Move a stream from one index to another.
*/
//go:nosplit
func (self class) MoveStream(index_from gd.Int, index_to gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index_from)
	callframe.Arg(frame, index_to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_move_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove the stream at the specified index.
*/
//go:nosplit
func (self class) RemoveStream(index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_remove_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Set the AudioStream at the specified index.
*/
//go:nosplit
func (self class) SetStream(index gd.Int, stream gdclass.AudioStream)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, discreet.Get(stream[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stream at the specified index.
*/
//go:nosplit
func (self class) GetStream(index gd.Int) gdclass.AudioStream {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.AudioStream{classdb.AudioStream(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Set the probability weight of the stream at the specified index. The higher this value, the more likely that the randomizer will choose this stream during random playback modes.
*/
//go:nosplit
func (self class) SetStreamProbabilityWeight(index gd.Int, weight gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_stream_probability_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the probability weight associated with the stream at the given index.
*/
//go:nosplit
func (self class) GetStreamProbabilityWeight(index gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_stream_probability_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStreamsCount(count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_streams_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamsCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_streams_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRandomPitch(scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_random_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRandomPitch() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_random_pitch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRandomVolumeOffsetDb(db_offset gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, db_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_random_volume_offset_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRandomVolumeOffsetDb() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_random_volume_offset_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaybackMode(mode classdb.AudioStreamRandomizerPlaybackMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_playback_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaybackMode() classdb.AudioStreamRandomizerPlaybackMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamRandomizerPlaybackMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_playback_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamRandomizer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamRandomizer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioStreamRandomizer", func(ptr gd.Object) any { return classdb.AudioStreamRandomizer(ptr) })}
type PlaybackMode = classdb.AudioStreamRandomizerPlaybackMode

const (
/*Pick a stream at random according to the probability weights chosen for each stream, but avoid playing the same stream twice in a row whenever possible. If only 1 sound is present in the pool, the same sound will always play, effectively allowing repeats to occur.*/
	PlaybackRandomNoRepeats PlaybackMode = 0
/*Pick a stream at random according to the probability weights chosen for each stream. If only 1 sound is present in the pool, the same sound will always play.*/
	PlaybackRandom PlaybackMode = 1
/*Play streams in the order they appear in the stream pool. If only 1 sound is present in the pool, the same sound will always play.*/
	PlaybackSequential PlaybackMode = 2
)
