package AudioStreamOggVorbis

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
The AudioStreamOggVorbis class is a specialized [AudioStream] for handling Ogg Vorbis file formats. It offers functionality for loading and playing back Ogg Vorbis files, as well as managing looping and other playback properties. This class is part of the audio stream system, which also supports WAV files through the [AudioStreamWAV] class.

*/
type Simple [1]classdb.AudioStreamOggVorbis
func (self Simple) LoadFromBuffer(buffer []byte) [1]classdb.AudioStreamOggVorbis {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStreamOggVorbis(Expert(self).LoadFromBuffer(gc, gc.PackedByteSlice(buffer)))
}
func (self Simple) LoadFromFile(path string) [1]classdb.AudioStreamOggVorbis {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStreamOggVorbis(Expert(self).LoadFromFile(gc, gc.String(path)))
}
func (self Simple) SetPacketSequence(packet_sequence [1]classdb.OggPacketSequence) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPacketSequence(packet_sequence)
}
func (self Simple) GetPacketSequence() [1]classdb.OggPacketSequence {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.OggPacketSequence(Expert(self).GetPacketSequence(gc))
}
func (self Simple) SetLoop(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLoop(enable)
}
func (self Simple) HasLoop() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasLoop())
}
func (self Simple) SetLoopOffset(seconds float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLoopOffset(gd.Float(seconds))
}
func (self Simple) GetLoopOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLoopOffset()))
}
func (self Simple) SetBpm(bpm float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBpm(gd.Float(bpm))
}
func (self Simple) GetBpm() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBpm()))
}
func (self Simple) SetBeatCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBeatCount(gd.Int(count))
}
func (self Simple) GetBeatCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBeatCount()))
}
func (self Simple) SetBarBeats(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBarBeats(gd.Int(count))
}
func (self Simple) GetBarBeats() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBarBeats()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamOggVorbis
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Creates a new AudioStreamOggVorbis instance from the given buffer. The buffer must contain Ogg Vorbis data.
*/
//go:nosplit
func (self class) LoadFromBuffer(ctx gd.Lifetime, buffer gd.PackedByteArray) object.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.AudioStreamOggVorbis.Bind_load_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStreamOggVorbis
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new AudioStreamOggVorbis instance from the given file path. The file must be in Ogg Vorbis format.
*/
//go:nosplit
func (self class) LoadFromFile(ctx gd.Lifetime, path gd.String) object.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.AudioStreamOggVorbis.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStreamOggVorbis
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPacketSequence(packet_sequence object.OggPacketSequence)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packet_sequence[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_packet_sequence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPacketSequence(ctx gd.Lifetime) object.OggPacketSequence {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_get_packet_sequence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.OggPacketSequence
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoop(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopOffset(seconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_loop_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_get_loop_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBpm(bpm gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bpm)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBpm() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBeatCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_beat_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBeatCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_get_beat_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBarBeats(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_bar_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBarBeats() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_get_bar_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAudioStreamOggVorbis() Expert { return self[0].AsAudioStreamOggVorbis() }


//go:nosplit
func (self Simple) AsAudioStreamOggVorbis() Simple { return self[0].AsAudioStreamOggVorbis() }


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
func init() {classdb.Register("AudioStreamOggVorbis", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
