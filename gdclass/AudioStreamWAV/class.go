package AudioStreamWAV

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
AudioStreamWAV stores sound samples loaded from WAV files. To play the stored sound, use an [AudioStreamPlayer] (for non-positional audio) or [AudioStreamPlayer2D]/[AudioStreamPlayer3D] (for positional audio). The sound can be looped.
This class can also be used to store dynamically-generated PCM audio data. See also [AudioStreamGenerator] for procedural audio generation.

*/
type Go [1]classdb.AudioStreamWAV

/*
Saves the AudioStreamWAV as a WAV file to [param path]. Samples with IMA ADPCM or QOA formats can't be saved.
[b]Note:[/b] A [code].wav[/code] extension is automatically appended to [param path] if it is missing.
*/
func (self Go) SaveToWav(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).SaveToWav(gc.String(path)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamWAV
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamWAV"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Data() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetData(gc).Bytes())
}

func (self Go) SetData(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetData(gc.PackedByteSlice(value))
}

func (self Go) Format() classdb.AudioStreamWAVFormat {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioStreamWAVFormat(class(self).GetFormat())
}

func (self Go) SetFormat(value classdb.AudioStreamWAVFormat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFormat(value)
}

func (self Go) LoopMode() classdb.AudioStreamWAVLoopMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioStreamWAVLoopMode(class(self).GetLoopMode())
}

func (self Go) SetLoopMode(value classdb.AudioStreamWAVLoopMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoopMode(value)
}

func (self Go) LoopBegin() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLoopBegin()))
}

func (self Go) SetLoopBegin(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoopBegin(gd.Int(value))
}

func (self Go) LoopEnd() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetLoopEnd()))
}

func (self Go) SetLoopEnd(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoopEnd(gd.Int(value))
}

func (self Go) MixRate() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMixRate()))
}

func (self Go) SetMixRate(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMixRate(gd.Int(value))
}

func (self Go) Stereo() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsStereo())
}

func (self Go) SetStereo(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStereo(value)
}

//go:nosplit
func (self class) SetData(data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetData(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFormat(format classdb.AudioStreamWAVFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFormat() classdb.AudioStreamWAVFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamWAVFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopMode(loop_mode classdb.AudioStreamWAVLoopMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopMode() classdb.AudioStreamWAVLoopMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamWAVLoopMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopBegin(loop_begin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop_begin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_loop_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopBegin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_get_loop_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoopEnd(loop_end gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop_end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_loop_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLoopEnd() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_get_loop_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMixRate(mix_rate gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mix_rate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_mix_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMixRate() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_get_mix_rate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStereo(stereo bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stereo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_set_stereo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsStereo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_is_stereo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Saves the AudioStreamWAV as a WAV file to [param path]. Samples with IMA ADPCM or QOA formats can't be saved.
[b]Note:[/b] A [code].wav[/code] extension is automatically appended to [param path] if it is missing.
*/
//go:nosplit
func (self class) SaveToWav(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamWAV.Bind_save_to_wav, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamWAV() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamWAV() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioStreamWAV", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Format = classdb.AudioStreamWAVFormat

const (
/*8-bit audio codec.*/
	Format8Bits Format = 0
/*16-bit audio codec.*/
	Format16Bits Format = 1
/*Audio is compressed using IMA ADPCM.*/
	FormatImaAdpcm Format = 2
/*Audio is compressed as QOA ([url=https://qoaformat.org/]Quite OK Audio[/url]).*/
	FormatQoa Format = 3
)
type LoopMode = classdb.AudioStreamWAVLoopMode

const (
/*Audio does not loop.*/
	LoopDisabled LoopMode = 0
/*Audio loops the data between [member loop_begin] and [member loop_end], playing forward only.*/
	LoopForward LoopMode = 1
/*Audio loops the data between [member loop_begin] and [member loop_end], playing back and forth.*/
	LoopPingpong LoopMode = 2
/*Audio loops the data between [member loop_begin] and [member loop_end], playing backward only.*/
	LoopBackward LoopMode = 3
)
