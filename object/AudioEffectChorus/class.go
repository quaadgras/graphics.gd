package AudioEffectChorus

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioEffect"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Adds a chorus audio effect. The effect applies a filter with voices to duplicate the audio source and manipulate it through the filter.

*/
type Simple [1]classdb.AudioEffectChorus
func (self Simple) SetVoiceCount(voices int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoiceCount(gd.Int(voices))
}
func (self Simple) GetVoiceCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVoiceCount()))
}
func (self Simple) SetVoiceDelayMs(voice_idx int, delay_ms float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoiceDelayMs(gd.Int(voice_idx), gd.Float(delay_ms))
}
func (self Simple) GetVoiceDelayMs(voice_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVoiceDelayMs(gd.Int(voice_idx))))
}
func (self Simple) SetVoiceRateHz(voice_idx int, rate_hz float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoiceRateHz(gd.Int(voice_idx), gd.Float(rate_hz))
}
func (self Simple) GetVoiceRateHz(voice_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVoiceRateHz(gd.Int(voice_idx))))
}
func (self Simple) SetVoiceDepthMs(voice_idx int, depth_ms float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoiceDepthMs(gd.Int(voice_idx), gd.Float(depth_ms))
}
func (self Simple) GetVoiceDepthMs(voice_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVoiceDepthMs(gd.Int(voice_idx))))
}
func (self Simple) SetVoiceLevelDb(voice_idx int, level_db float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoiceLevelDb(gd.Int(voice_idx), gd.Float(level_db))
}
func (self Simple) GetVoiceLevelDb(voice_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVoiceLevelDb(gd.Int(voice_idx))))
}
func (self Simple) SetVoiceCutoffHz(voice_idx int, cutoff_hz float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoiceCutoffHz(gd.Int(voice_idx), gd.Float(cutoff_hz))
}
func (self Simple) GetVoiceCutoffHz(voice_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVoiceCutoffHz(gd.Int(voice_idx))))
}
func (self Simple) SetVoicePan(voice_idx int, pan float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVoicePan(gd.Int(voice_idx), gd.Float(pan))
}
func (self Simple) GetVoicePan(voice_idx int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVoicePan(gd.Int(voice_idx))))
}
func (self Simple) SetWet(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWet(gd.Float(amount))
}
func (self Simple) GetWet() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetWet()))
}
func (self Simple) SetDry(amount float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDry(gd.Float(amount))
}
func (self Simple) GetDry() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDry()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioEffectChorus
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetVoiceCount(voices gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voices)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoiceCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVoiceDelayMs(voice_idx gd.Int, delay_ms gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, delay_ms)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoiceDelayMs(voice_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVoiceRateHz(voice_idx gd.Int, rate_hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, rate_hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoiceRateHz(voice_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVoiceDepthMs(voice_idx gd.Int, depth_ms gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, depth_ms)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_depth_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoiceDepthMs(voice_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_depth_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVoiceLevelDb(voice_idx gd.Int, level_db gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, level_db)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoiceLevelDb(voice_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVoiceCutoffHz(voice_idx gd.Int, cutoff_hz gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, cutoff_hz)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoiceCutoffHz(voice_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVoicePan(voice_idx gd.Int, pan gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, pan)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_voice_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVoicePan(voice_idx gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_voice_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWet(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_wet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWet() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_wet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDry(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_set_dry, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDry() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectChorus.Bind_get_dry, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAudioEffectChorus() Expert { return self[0].AsAudioEffectChorus() }


//go:nosplit
func (self Simple) AsAudioEffectChorus() Simple { return self[0].AsAudioEffectChorus() }


//go:nosplit
func (self class) AsAudioEffect() AudioEffect.Expert { return self[0].AsAudioEffect() }


//go:nosplit
func (self Simple) AsAudioEffect() AudioEffect.Simple { return self[0].AsAudioEffect() }


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
func init() {classdb.Register("AudioEffectChorus", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
