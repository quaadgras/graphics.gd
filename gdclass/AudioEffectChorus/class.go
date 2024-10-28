package AudioEffectChorus

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
Adds a chorus audio effect. The effect applies a filter with voices to duplicate the audio source and manipulate it through the filter.

*/
type Go [1]classdb.AudioEffectChorus
func (self Go) SetVoiceDelayMs(voice_idx int, delay_ms float64) {
	class(self).SetVoiceDelayMs(gd.Int(voice_idx), gd.Float(delay_ms))
}
func (self Go) GetVoiceDelayMs(voice_idx int) float64 {
	return float64(float64(class(self).GetVoiceDelayMs(gd.Int(voice_idx))))
}
func (self Go) SetVoiceRateHz(voice_idx int, rate_hz float64) {
	class(self).SetVoiceRateHz(gd.Int(voice_idx), gd.Float(rate_hz))
}
func (self Go) GetVoiceRateHz(voice_idx int) float64 {
	return float64(float64(class(self).GetVoiceRateHz(gd.Int(voice_idx))))
}
func (self Go) SetVoiceDepthMs(voice_idx int, depth_ms float64) {
	class(self).SetVoiceDepthMs(gd.Int(voice_idx), gd.Float(depth_ms))
}
func (self Go) GetVoiceDepthMs(voice_idx int) float64 {
	return float64(float64(class(self).GetVoiceDepthMs(gd.Int(voice_idx))))
}
func (self Go) SetVoiceLevelDb(voice_idx int, level_db float64) {
	class(self).SetVoiceLevelDb(gd.Int(voice_idx), gd.Float(level_db))
}
func (self Go) GetVoiceLevelDb(voice_idx int) float64 {
	return float64(float64(class(self).GetVoiceLevelDb(gd.Int(voice_idx))))
}
func (self Go) SetVoiceCutoffHz(voice_idx int, cutoff_hz float64) {
	class(self).SetVoiceCutoffHz(gd.Int(voice_idx), gd.Float(cutoff_hz))
}
func (self Go) GetVoiceCutoffHz(voice_idx int) float64 {
	return float64(float64(class(self).GetVoiceCutoffHz(gd.Int(voice_idx))))
}
func (self Go) SetVoicePan(voice_idx int, pan float64) {
	class(self).SetVoicePan(gd.Int(voice_idx), gd.Float(pan))
}
func (self Go) GetVoicePan(voice_idx int) float64 {
	return float64(float64(class(self).GetVoicePan(gd.Int(voice_idx))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectChorus
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectChorus"))
	return Go{classdb.AudioEffectChorus(object)}
}

func (self Go) VoiceCount() int {
		return int(int(class(self).GetVoiceCount()))
}

func (self Go) SetVoiceCount(value int) {
	class(self).SetVoiceCount(gd.Int(value))
}

func (self Go) Dry() float64 {
		return float64(float64(class(self).GetDry()))
}

func (self Go) SetDry(value float64) {
	class(self).SetDry(gd.Float(value))
}

func (self Go) Wet() float64 {
		return float64(float64(class(self).GetWet()))
}

func (self Go) SetWet(value float64) {
	class(self).SetWet(gd.Float(value))
}

//go:nosplit
func (self class) SetVoiceCount(voices gd.Int)  {
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
func (self class) SetVoiceDelayMs(voice_idx gd.Int, delay_ms gd.Float)  {
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
func (self class) SetVoiceRateHz(voice_idx gd.Int, rate_hz gd.Float)  {
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
func (self class) SetVoiceDepthMs(voice_idx gd.Int, depth_ms gd.Float)  {
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
func (self class) SetVoiceLevelDb(voice_idx gd.Int, level_db gd.Float)  {
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
func (self class) SetVoiceCutoffHz(voice_idx gd.Int, cutoff_hz gd.Float)  {
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
func (self class) SetVoicePan(voice_idx gd.Int, pan gd.Float)  {
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
func (self class) SetWet(amount gd.Float)  {
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
func (self class) SetDry(amount gd.Float)  {
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
func (self class) AsAudioEffectChorus() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectChorus() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectChorus", func(ptr gd.Object) any { return classdb.AudioEffectChorus(ptr) })}
