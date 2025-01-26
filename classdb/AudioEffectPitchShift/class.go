// Package AudioEffectPitchShift provides methods for working with AudioEffectPitchShift object instances.
package AudioEffectPitchShift

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
import "graphics.gd/variant/Dictionary"
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
var _ Dictionary.Any

/*
Allows modulation of pitch independently of tempo. All frequencies can be increased/decreased with minimal effect on transients.
*/
type Instance [1]gdclass.AudioEffectPitchShift

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectPitchShift() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectPitchShift

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectPitchShift"))
	casted := Instance{*(*gdclass.AudioEffectPitchShift)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) PitchScale() Float.X {
	return Float.X(Float.X(class(self).GetPitchScale()))
}

func (self Instance) SetPitchScale(value Float.X) {
	class(self).SetPitchScale(gd.Float(value))
}

func (self Instance) Oversampling() int {
	return int(int(class(self).GetOversampling()))
}

func (self Instance) SetOversampling(value int) {
	class(self).SetOversampling(gd.Int(value))
}

func (self Instance) FftSize() gdclass.AudioEffectPitchShiftFFTSize {
	return gdclass.AudioEffectPitchShiftFFTSize(class(self).GetFftSize())
}

func (self Instance) SetFftSize(value gdclass.AudioEffectPitchShiftFFTSize) {
	class(self).SetFftSize(value)
}

//go:nosplit
func (self class) SetPitchScale(rate gd.Float) { //gd:AudioEffectPitchShift.set_pitch_scale
	var frame = callframe.New()
	callframe.Arg(frame, rate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPitchShift.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitchScale() gd.Float { //gd:AudioEffectPitchShift.get_pitch_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPitchShift.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOversampling(amount gd.Int) { //gd:AudioEffectPitchShift.set_oversampling
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPitchShift.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOversampling() gd.Int { //gd:AudioEffectPitchShift.get_oversampling
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPitchShift.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFftSize(size gdclass.AudioEffectPitchShiftFFTSize) { //gd:AudioEffectPitchShift.set_fft_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPitchShift.Bind_set_fft_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFftSize() gdclass.AudioEffectPitchShiftFFTSize { //gd:AudioEffectPitchShift.get_fft_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioEffectPitchShiftFFTSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPitchShift.Bind_get_fft_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectPitchShift() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectPitchShift() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectPitchShift", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectPitchShift{*(*gdclass.AudioEffectPitchShift)(unsafe.Pointer(&ptr))}
	})
}

type FFTSize = gdclass.AudioEffectPitchShiftFFTSize //gd:AudioEffectPitchShift.FFTSize

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
