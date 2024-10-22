package AudioEffectDistortion

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
Different types are available: clip, tan, lo-fi (bit crushing), overdrive, or waveshape.
By distorting the waveform the frequency content changes, which will often make the sound "crunchy" or "abrasive". For games, it can simulate sound coming from some saturated device or speaker very efficiently.

*/
type Go [1]classdb.AudioEffectDistortion
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectDistortion
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectDistortion"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Mode() classdb.AudioEffectDistortionMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioEffectDistortionMode(class(self).GetMode())
}

func (self Go) SetMode(value classdb.AudioEffectDistortionMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMode(value)
}

func (self Go) PreGain() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPreGain()))
}

func (self Go) SetPreGain(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPreGain(gd.Float(value))
}

func (self Go) KeepHfHz() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetKeepHfHz()))
}

func (self Go) SetKeepHfHz(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeepHfHz(gd.Float(value))
}

func (self Go) Drive() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDrive()))
}

func (self Go) SetDrive(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrive(gd.Float(value))
}

func (self Go) PostGain() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPostGain()))
}

func (self Go) SetPostGain(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPostGain(gd.Float(value))
}

//go:nosplit
func (self class) SetMode(mode classdb.AudioEffectDistortionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMode() classdb.AudioEffectDistortionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioEffectDistortionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreGain(pre_gain gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pre_gain)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_set_pre_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreGain() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_get_pre_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepHfHz(keep_hf_hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, keep_hf_hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_set_keep_hf_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeepHfHz() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_get_keep_hf_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrive(drive gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, drive)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_set_drive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDrive() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_get_drive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPostGain(post_gain gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, post_gain)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_set_post_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPostGain() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDistortion.Bind_get_post_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectDistortion() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectDistortion() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectDistortion", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Mode = classdb.AudioEffectDistortionMode

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
