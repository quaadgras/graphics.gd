package AudioStreamPlaylist

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

type Simple [1]classdb.AudioStreamPlaylist
func (self Simple) SetStreamCount(stream_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStreamCount(gd.Int(stream_count))
}
func (self Simple) GetStreamCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetStreamCount()))
}
func (self Simple) GetBpm() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetBpm()))
}
func (self Simple) SetListStream(stream_index int, audio_stream [1]classdb.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetListStream(gd.Int(stream_index), audio_stream)
}
func (self Simple) GetListStream(stream_index int) [1]classdb.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStream(Expert(self).GetListStream(gc, gd.Int(stream_index)))
}
func (self Simple) SetShuffle(shuffle bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShuffle(shuffle)
}
func (self Simple) GetShuffle() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetShuffle())
}
func (self Simple) SetFadeTime(dec float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFadeTime(gd.Float(dec))
}
func (self Simple) GetFadeTime() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetFadeTime()))
}
func (self Simple) SetLoop(loop bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLoop(loop)
}
func (self Simple) HasLoop() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasLoop())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamPlaylist
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
//go:nosplit
func (self class) GetBpm() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the stream at playback position index.
*/
//go:nosplit
func (self class) SetListStream(stream_index gd.Int, audio_stream object.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, mmm.Get(audio_stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_list_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stream at playback position index.
*/
//go:nosplit
func (self class) GetListStream(ctx gd.Lifetime, stream_index gd.Int) object.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_list_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShuffle(shuffle bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shuffle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_shuffle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShuffle() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_shuffle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFadeTime(dec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, dec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_fade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFadeTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_fade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoop(loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAudioStreamPlaylist() Expert { return self[0].AsAudioStreamPlaylist() }


//go:nosplit
func (self Simple) AsAudioStreamPlaylist() Simple { return self[0].AsAudioStreamPlaylist() }


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
func init() {classdb.Register("AudioStreamPlaylist", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
