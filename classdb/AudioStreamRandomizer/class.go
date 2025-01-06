// Package AudioStreamRandomizer provides methods for working with AudioStreamRandomizer object instances.
package AudioStreamRandomizer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Picks a random AudioStream from the pool, depending on the playback mode, and applies random pitch shifting and volume shifting during playback.
*/
type Instance [1]gdclass.AudioStreamRandomizer
type Any interface {
	gd.IsClass
	AsAudioStreamRandomizer() Instance
}

/*
Insert a stream at the specified index. If the index is less than zero, the insertion occurs at the end of the underlying pool.
*/
func (self Instance) AddStream(index int, stream [1]gdclass.AudioStream) {
	class(self).AddStream(gd.Int(index), stream, gd.Float(1.0))
}

/*
Move a stream from one index to another.
*/
func (self Instance) MoveStream(index_from int, index_to int) {
	class(self).MoveStream(gd.Int(index_from), gd.Int(index_to))
}

/*
Remove the stream at the specified index.
*/
func (self Instance) RemoveStream(index int) {
	class(self).RemoveStream(gd.Int(index))
}

/*
Set the AudioStream at the specified index.
*/
func (self Instance) SetStream(index int, stream [1]gdclass.AudioStream) {
	class(self).SetStream(gd.Int(index), stream)
}

/*
Returns the stream at the specified index.
*/
func (self Instance) GetStream(index int) [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetStream(gd.Int(index)))
}

/*
Set the probability weight of the stream at the specified index. The higher this value, the more likely that the randomizer will choose this stream during random playback modes.
*/
func (self Instance) SetStreamProbabilityWeight(index int, weight Float.X) {
	class(self).SetStreamProbabilityWeight(gd.Int(index), gd.Float(weight))
}

/*
Returns the probability weight associated with the stream at the given index.
*/
func (self Instance) GetStreamProbabilityWeight(index int) Float.X {
	return Float.X(Float.X(class(self).GetStreamProbabilityWeight(gd.Int(index))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamRandomizer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamRandomizer"))
	casted := Instance{*(*gdclass.AudioStreamRandomizer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) PlaybackMode() gdclass.AudioStreamRandomizerPlaybackMode {
	return gdclass.AudioStreamRandomizerPlaybackMode(class(self).GetPlaybackMode())
}

func (self Instance) SetPlaybackMode(value gdclass.AudioStreamRandomizerPlaybackMode) {
	class(self).SetPlaybackMode(value)
}

func (self Instance) RandomPitch() Float.X {
	return Float.X(Float.X(class(self).GetRandomPitch()))
}

func (self Instance) SetRandomPitch(value Float.X) {
	class(self).SetRandomPitch(gd.Float(value))
}

func (self Instance) RandomVolumeOffsetDb() Float.X {
	return Float.X(Float.X(class(self).GetRandomVolumeOffsetDb()))
}

func (self Instance) SetRandomVolumeOffsetDb(value Float.X) {
	class(self).SetRandomVolumeOffsetDb(gd.Float(value))
}

func (self Instance) StreamsCount() int {
	return int(int(class(self).GetStreamsCount()))
}

func (self Instance) SetStreamsCount(value int) {
	class(self).SetStreamsCount(gd.Int(value))
}

/*
Insert a stream at the specified index. If the index is less than zero, the insertion occurs at the end of the underlying pool.
*/
//go:nosplit
func (self class) AddStream(index gd.Int, stream [1]gdclass.AudioStream, weight gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_add_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Move a stream from one index to another.
*/
//go:nosplit
func (self class) MoveStream(index_from gd.Int, index_to gd.Int) {
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
func (self class) RemoveStream(index gd.Int) {
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
func (self class) SetStream(index gd.Int, stream [1]gdclass.AudioStream) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the stream at the specified index.
*/
//go:nosplit
func (self class) GetStream(index gd.Int) [1]gdclass.AudioStream {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Set the probability weight of the stream at the specified index. The higher this value, the more likely that the randomizer will choose this stream during random playback modes.
*/
//go:nosplit
func (self class) SetStreamProbabilityWeight(index gd.Int, weight gd.Float) {
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
func (self class) SetStreamsCount(count gd.Int) {
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
func (self class) SetRandomPitch(scale gd.Float) {
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
func (self class) SetRandomVolumeOffsetDb(db_offset gd.Float) {
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
func (self class) SetPlaybackMode(mode gdclass.AudioStreamRandomizerPlaybackMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_set_playback_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaybackMode() gdclass.AudioStreamRandomizerPlaybackMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamRandomizerPlaybackMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamRandomizer.Bind_get_playback_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamRandomizer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamRandomizer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioStreamRandomizer", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamRandomizer{*(*gdclass.AudioStreamRandomizer)(unsafe.Pointer(&ptr))}
	})
}

type PlaybackMode = gdclass.AudioStreamRandomizerPlaybackMode

const (
	/*Pick a stream at random according to the probability weights chosen for each stream, but avoid playing the same stream twice in a row whenever possible. If only 1 sound is present in the pool, the same sound will always play, effectively allowing repeats to occur.*/
	PlaybackRandomNoRepeats PlaybackMode = 0
	/*Pick a stream at random according to the probability weights chosen for each stream. If only 1 sound is present in the pool, the same sound will always play.*/
	PlaybackRandom PlaybackMode = 1
	/*Play streams in the order they appear in the stream pool. If only 1 sound is present in the pool, the same sound will always play.*/
	PlaybackSequential PlaybackMode = 2
)
