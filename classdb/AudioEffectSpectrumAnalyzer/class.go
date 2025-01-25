// Package AudioEffectSpectrumAnalyzer provides methods for working with AudioEffectSpectrumAnalyzer object instances.
package AudioEffectSpectrumAnalyzer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This audio effect does not affect sound output, but can be used for real-time audio visualizations.
This resource configures an [AudioEffectSpectrumAnalyzerInstance], which performs the actual analysis at runtime. An instance can be acquired with [method AudioServer.get_bus_effect_instance].
See also [AudioStreamGenerator] for procedurally generating sounds.
*/
type Instance [1]gdclass.AudioEffectSpectrumAnalyzer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectSpectrumAnalyzer() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectSpectrumAnalyzer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectSpectrumAnalyzer"))
	casted := Instance{*(*gdclass.AudioEffectSpectrumAnalyzer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BufferLength() Float.X {
	return Float.X(Float.X(class(self).GetBufferLength()))
}

func (self Instance) SetBufferLength(value Float.X) {
	class(self).SetBufferLength(gd.Float(value))
}

func (self Instance) TapBackPos() Float.X {
	return Float.X(Float.X(class(self).GetTapBackPos()))
}

func (self Instance) SetTapBackPos(value Float.X) {
	class(self).SetTapBackPos(gd.Float(value))
}

func (self Instance) FftSize() gdclass.AudioEffectSpectrumAnalyzerFFTSize {
	return gdclass.AudioEffectSpectrumAnalyzerFFTSize(class(self).GetFftSize())
}

func (self Instance) SetFftSize(value gdclass.AudioEffectSpectrumAnalyzerFFTSize) {
	class(self).SetFftSize(value)
}

//go:nosplit
func (self class) SetBufferLength(seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzer.Bind_set_buffer_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzer.Bind_get_buffer_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTapBackPos(seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzer.Bind_set_tap_back_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTapBackPos() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzer.Bind_get_tap_back_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFftSize(size gdclass.AudioEffectSpectrumAnalyzerFFTSize) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzer.Bind_set_fft_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFftSize() gdclass.AudioEffectSpectrumAnalyzerFFTSize {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioEffectSpectrumAnalyzerFFTSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzer.Bind_get_fft_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectSpectrumAnalyzer() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectSpectrumAnalyzer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioEffect() AudioEffect.Advanced {
	return *((*AudioEffect.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffect() AudioEffect.Instance {
	return *((*AudioEffect.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioEffect.Advanced(self.AsAudioEffect()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffect.Instance(self.AsAudioEffect()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectSpectrumAnalyzer", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectSpectrumAnalyzer{*(*gdclass.AudioEffectSpectrumAnalyzer)(unsafe.Pointer(&ptr))}
	})
}

type FFTSize = gdclass.AudioEffectSpectrumAnalyzerFFTSize

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
