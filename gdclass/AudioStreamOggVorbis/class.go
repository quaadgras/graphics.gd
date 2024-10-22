package AudioStreamOggVorbis

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
The AudioStreamOggVorbis class is a specialized [AudioStream] for handling Ogg Vorbis file formats. It offers functionality for loading and playing back Ogg Vorbis files, as well as managing looping and other playback properties. This class is part of the audio stream system, which also supports WAV files through the [AudioStreamWAV] class.

*/
type Go [1]classdb.AudioStreamOggVorbis

/*
Creates a new AudioStreamOggVorbis instance from the given buffer. The buffer must contain Ogg Vorbis data.
*/
func (self Go) LoadFromBuffer(buffer []byte) gdclass.AudioStreamOggVorbis {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStreamOggVorbis(class(self).LoadFromBuffer(gc, gc.PackedByteSlice(buffer)))
}

/*
Creates a new AudioStreamOggVorbis instance from the given file path. The file must be in Ogg Vorbis format.
*/
func (self Go) LoadFromFile(path string) gdclass.AudioStreamOggVorbis {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStreamOggVorbis(class(self).LoadFromFile(gc, gc.String(path)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamOggVorbis
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamOggVorbis"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) PacketSequence() gdclass.OggPacketSequence {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.OggPacketSequence(class(self).GetPacketSequence(gc))
}

func (self Go) SetPacketSequence(value gdclass.OggPacketSequence) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPacketSequence(value)
}

func (self Go) Bpm() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBpm()))
}

func (self Go) SetBpm(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBpm(gd.Float(value))
}

func (self Go) BeatCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBeatCount()))
}

func (self Go) SetBeatCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBeatCount(gd.Int(value))
}

func (self Go) BarBeats() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBarBeats()))
}

func (self Go) SetBarBeats(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBarBeats(gd.Int(value))
}

func (self Go) Loop() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasLoop())
}

func (self Go) SetLoop(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoop(value)
}

func (self Go) LoopOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetLoopOffset()))
}

func (self Go) SetLoopOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoopOffset(gd.Float(value))
}

/*
Creates a new AudioStreamOggVorbis instance from the given buffer. The buffer must contain Ogg Vorbis data.
*/
//go:nosplit
func (self class) LoadFromBuffer(ctx gd.Lifetime, buffer gd.PackedByteArray) gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(buffer))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.AudioStreamOggVorbis.Bind_load_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStreamOggVorbis
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Creates a new AudioStreamOggVorbis instance from the given file path. The file must be in Ogg Vorbis format.
*/
//go:nosplit
func (self class) LoadFromFile(ctx gd.Lifetime, path gd.String) gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.AudioStreamOggVorbis.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStreamOggVorbis
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPacketSequence(packet_sequence gdclass.OggPacketSequence)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packet_sequence[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_set_packet_sequence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPacketSequence(ctx gd.Lifetime) gdclass.OggPacketSequence {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamOggVorbis.Bind_get_packet_sequence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OggPacketSequence
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
func (self class) AsAudioStreamOggVorbis() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamOggVorbis() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioStreamOggVorbis", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
