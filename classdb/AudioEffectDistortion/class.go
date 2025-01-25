// Package AudioEffectDistortion provides methods for working with AudioEffectDistortion object instances.
package AudioEffectDistortion

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

/*
Different types are available: clip, tan, lo-fi (bit crushing), overdrive, or waveshape.
By distorting the waveform the frequency content changes, which will often make the sound "crunchy" or "abrasive". For games, it can simulate sound coming from some saturated device or speaker very efficiently.
*/
type Instance [1]gdclass.AudioEffectDistortion

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectDistortion() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectDistortion

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectDistortion"))
	casted := Instance{*(*gdclass.AudioEffectDistortion)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Mode() gdclass.AudioEffectDistortionMode {
	return gdclass.AudioEffectDistortionMode(class(self).GetMode())
}

func (self Instance) SetMode(value gdclass.AudioEffectDistortionMode) {
	class(self).SetMode(value)
}

func (self Instance) PreGain() Float.X {
	return Float.X(Float.X(class(self).GetPreGain()))
}

func (self Instance) SetPreGain(value Float.X) {
	class(self).SetPreGain(gd.Float(value))
}

func (self Instance) KeepHfHz() Float.X {
	return Float.X(Float.X(class(self).GetKeepHfHz()))
}

func (self Instance) SetKeepHfHz(value Float.X) {
	class(self).SetKeepHfHz(gd.Float(value))
}

func (self Instance) Drive() Float.X {
	return Float.X(Float.X(class(self).GetDrive()))
}

func (self Instance) SetDrive(value Float.X) {
	class(self).SetDrive(gd.Float(value))
}

func (self Instance) PostGain() Float.X {
	return Float.X(Float.X(class(self).GetPostGain()))
}

func (self Instance) SetPostGain(value Float.X) {
	class(self).SetPostGain(gd.Float(value))
}

//go:nosplit
func (self class) SetMode(mode gdclass.AudioEffectDistortionMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() gdclass.AudioEffectDistortionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioEffectDistortionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreGain(pre_gain gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pre_gain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_pre_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreGain() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_pre_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepHfHz(keep_hf_hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, keep_hf_hz)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_keep_hf_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetKeepHfHz() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_keep_hf_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrive(drive gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, drive)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_drive, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrive() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_drive, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPostGain(post_gain gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, post_gain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_post_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPostGain() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_post_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectDistortion() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectDistortion() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectDistortion", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectDistortion{*(*gdclass.AudioEffectDistortion)(unsafe.Pointer(&ptr))}
	})
}

type Mode = gdclass.AudioEffectDistortionMode

const (
	/*Digital distortion effect which cuts off peaks at the top and bottom of the waveform.*/
	ModeClip Mode = 0
	ModeAtan Mode = 1
	/*Low-resolution digital distortion effect (bit depth reduction). You can use it to emulate the sound of early digital audio devices.*/
	ModeLofi Mode = 2
	/*Emulates the warm distortion produced by a field effect transistor, which is commonly used in solid-state musical instrument amplifiers. The [member drive] property has no effect in this mode.*/
	ModeOverdrive Mode = 3
	/*Waveshaper distortions are used mainly by electronic musicians to achieve an extra-abrasive sound.*/
	ModeWaveshape Mode = 4
)
