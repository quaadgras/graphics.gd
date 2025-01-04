package AudioEffectChorus

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Adds a chorus audio effect. The effect applies a filter with voices to duplicate the audio source and manipulate it through the filter.
*/
type Instance [1]gdclass.AudioEffectChorus
type Any interface {
	gd.IsClass
	AsAudioEffectChorus() Instance
}

func (self Instance) SetVoiceDelayMs(voice_idx int, delay_ms Float.X) {
	class(self).SetVoiceDelayMs(gd.Int(voice_idx), gd.Float(delay_ms))
}
func (self Instance) GetVoiceDelayMs(voice_idx int) Float.X {
	return Float.X(Float.X(class(self).GetVoiceDelayMs(gd.Int(voice_idx))))
}
func (self Instance) SetVoiceRateHz(voice_idx int, rate_hz Float.X) {
	class(self).SetVoiceRateHz(gd.Int(voice_idx), gd.Float(rate_hz))
}
func (self Instance) GetVoiceRateHz(voice_idx int) Float.X {
	return Float.X(Float.X(class(self).GetVoiceRateHz(gd.Int(voice_idx))))
}
func (self Instance) SetVoiceDepthMs(voice_idx int, depth_ms Float.X) {
	class(self).SetVoiceDepthMs(gd.Int(voice_idx), gd.Float(depth_ms))
}
func (self Instance) GetVoiceDepthMs(voice_idx int) Float.X {
	return Float.X(Float.X(class(self).GetVoiceDepthMs(gd.Int(voice_idx))))
}
func (self Instance) SetVoiceLevelDb(voice_idx int, level_db Float.X) {
	class(self).SetVoiceLevelDb(gd.Int(voice_idx), gd.Float(level_db))
}
func (self Instance) GetVoiceLevelDb(voice_idx int) Float.X {
	return Float.X(Float.X(class(self).GetVoiceLevelDb(gd.Int(voice_idx))))
}
func (self Instance) SetVoiceCutoffHz(voice_idx int, cutoff_hz Float.X) {
	class(self).SetVoiceCutoffHz(gd.Int(voice_idx), gd.Float(cutoff_hz))
}
func (self Instance) GetVoiceCutoffHz(voice_idx int) Float.X {
	return Float.X(Float.X(class(self).GetVoiceCutoffHz(gd.Int(voice_idx))))
}
func (self Instance) SetVoicePan(voice_idx int, pan Float.X) {
	class(self).SetVoicePan(gd.Int(voice_idx), gd.Float(pan))
}
func (self Instance) GetVoicePan(voice_idx int) Float.X {
	return Float.X(Float.X(class(self).GetVoicePan(gd.Int(voice_idx))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectChorus

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectChorus"))
	return Instance{*(*gdclass.AudioEffectChorus)(unsafe.Pointer(&object))}
}

func (self Instance) VoiceCount() int {
	return int(int(class(self).GetVoiceCount()))
}

func (self Instance) SetVoiceCount(value int) {
	class(self).SetVoiceCount(gd.Int(value))
}

func (self Instance) Dry() Float.X {
	return Float.X(Float.X(class(self).GetDry()))
}

func (self Instance) SetDry(value Float.X) {
	class(self).SetDry(gd.Float(value))
}

func (self Instance) Wet() Float.X {
	return Float.X(Float.X(class(self).GetWet()))
}

func (self Instance) SetWet(value Float.X) {
	class(self).SetWet(gd.Float(value))
}

//go:nosplit
func (self class) SetVoiceCount(voices gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, voices)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceDelayMs(voice_idx gd.Int, delay_ms gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, delay_ms)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceDelayMs(voice_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceRateHz(voice_idx gd.Int, rate_hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, rate_hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceRateHz(voice_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceDepthMs(voice_idx gd.Int, depth_ms gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, depth_ms)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_depth_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceDepthMs(voice_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_depth_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceLevelDb(voice_idx gd.Int, level_db gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, level_db)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceLevelDb(voice_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceCutoffHz(voice_idx gd.Int, cutoff_hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, cutoff_hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceCutoffHz(voice_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoicePan(voice_idx gd.Int, pan gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, pan)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoicePan(voice_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWet(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_wet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWet() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_wet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDry(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_dry, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDry() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_dry, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectChorus() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectChorus() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {
	gdclass.Register("AudioEffectChorus", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectChorus{*(*gdclass.AudioEffectChorus)(unsafe.Pointer(&ptr))}
	})
}
