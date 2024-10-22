package AudioEffectSpectrumAnalyzer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This audio effect does not affect sound output, but can be used for real-time audio visualizations.
This resource configures an [AudioEffectSpectrumAnalyzerInstance], which performs the actual analysis at runtime. An instance can be acquired with [method AudioServer.get_bus_effect_instance].
See also [AudioStreamGenerator] for procedurally generating sounds.

*/
type Go [1]classdb.AudioEffectSpectrumAnalyzer
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectSpectrumAnalyzer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectSpectrumAnalyzer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BufferLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetBufferLength()))
}

func (self Go) SetBufferLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBufferLength(gd.Float(value))
}

func (self Go) TapBackPos() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTapBackPos()))
}

func (self Go) SetTapBackPos(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTapBackPos(gd.Float(value))
}

func (self Go) FftSize() classdb.AudioEffectSpectrumAnalyzerFFTSize {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioEffectSpectrumAnalyzerFFTSize(class(self).GetFftSize())
}

func (self Go) SetFftSize(value classdb.AudioEffectSpectrumAnalyzerFFTSize) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFftSize(value)
}

//go:nosplit
func (self class) SetBufferLength(seconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzer.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBufferLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzer.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTapBackPos(seconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzer.Bind_set_tap_back_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTapBackPos() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzer.Bind_get_tap_back_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFftSize(size classdb.AudioEffectSpectrumAnalyzerFFTSize)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzer.Bind_set_fft_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFftSize() classdb.AudioEffectSpectrumAnalyzerFFTSize {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioEffectSpectrumAnalyzerFFTSize](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzer.Bind_get_fft_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectSpectrumAnalyzer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectSpectrumAnalyzer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.GD { return *((*AudioEffect.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() AudioEffect.Go { return *((*AudioEffect.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {classdb.Register("AudioEffectSpectrumAnalyzer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type FFTSize = classdb.AudioEffectSpectrumAnalyzerFFTSize

const (
/*Use a buffer of 256 samples for the Fast Fourier transform. Lowest latency, but least stable over time.*/
	FftSize256 FFTSize = 0
/*Use a buffer of 512 samples for the Fast Fourier transform. Low latency, but less stable over time.*/
	FftSize512 FFTSize = 1
/*Use a buffer of 1024 samples for the Fast Fourier transform. This is a compromise between latency and stability over time.*/
	FftSize1024 FFTSize = 2
/*Use a buffer of 2048 samples for the Fast Fourier transform. High latency, but stable over time.*/
	FftSize2048 FFTSize = 3
/*Use a buffer of 4096 samples for the Fast Fourier transform. Highest latency, but most stable over time.*/
	FftSize4096 FFTSize = 4
/*Represents the size of the [enum FFTSize] enum.*/
	FftSizeMax FFTSize = 5
)
