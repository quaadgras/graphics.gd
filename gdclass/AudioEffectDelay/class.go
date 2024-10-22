package AudioEffectDelay

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
Plays input signal back after a period of time. The delayed signal may be played back multiple times to create the sound of a repeating, decaying echo. Delay effects range from a subtle echo effect to a pronounced blending of previous sounds with new sounds.

*/
type Go [1]classdb.AudioEffectDelay
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectDelay
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectDelay"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Dry() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDry()))
}

func (self Go) SetDry(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDry(gd.Float(value))
}

func (self Go) Tap1Active() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsTap1Active())
}

func (self Go) SetTap1Active(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap1Active(value)
}

func (self Go) Tap1DelayMs() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTap1DelayMs()))
}

func (self Go) SetTap1DelayMs(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap1DelayMs(gd.Float(value))
}

func (self Go) Tap1LevelDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTap1LevelDb()))
}

func (self Go) SetTap1LevelDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap1LevelDb(gd.Float(value))
}

func (self Go) Tap1Pan() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTap1Pan()))
}

func (self Go) SetTap1Pan(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap1Pan(gd.Float(value))
}

func (self Go) Tap2Active() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsTap2Active())
}

func (self Go) SetTap2Active(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap2Active(value)
}

func (self Go) Tap2DelayMs() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTap2DelayMs()))
}

func (self Go) SetTap2DelayMs(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap2DelayMs(gd.Float(value))
}

func (self Go) Tap2LevelDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTap2LevelDb()))
}

func (self Go) SetTap2LevelDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap2LevelDb(gd.Float(value))
}

func (self Go) Tap2Pan() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTap2Pan()))
}

func (self Go) SetTap2Pan(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTap2Pan(gd.Float(value))
}

func (self Go) FeedbackActive() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFeedbackActive())
}

func (self Go) SetFeedbackActive(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFeedbackActive(value)
}

func (self Go) FeedbackDelayMs() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFeedbackDelayMs()))
}

func (self Go) SetFeedbackDelayMs(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFeedbackDelayMs(gd.Float(value))
}

func (self Go) FeedbackLevelDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFeedbackLevelDb()))
}

func (self Go) SetFeedbackLevelDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFeedbackLevelDb(gd.Float(value))
}

func (self Go) FeedbackLowpass() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFeedbackLowpass()))
}

func (self Go) SetFeedbackLowpass(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFeedbackLowpass(gd.Float(value))
}

//go:nosplit
func (self class) SetDry(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_dry, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDry() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_dry, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap1Active(amount bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap1_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTap1Active() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_is_tap1_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap1DelayMs(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap1_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTap1DelayMs() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_tap1_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap1LevelDb(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap1_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTap1LevelDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_tap1_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap1Pan(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap1_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTap1Pan() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_tap1_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap2Active(amount bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap2_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsTap2Active() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_is_tap2_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap2DelayMs(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap2_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTap2DelayMs() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_tap2_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap2LevelDb(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap2_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTap2LevelDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_tap2_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTap2Pan(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_tap2_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTap2Pan() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_tap2_pan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFeedbackActive(amount bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_feedback_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFeedbackActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_is_feedback_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFeedbackDelayMs(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_feedback_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFeedbackDelayMs() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_feedback_delay_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFeedbackLevelDb(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_feedback_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFeedbackLevelDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_feedback_level_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFeedbackLowpass(amount gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_set_feedback_lowpass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFeedbackLowpass() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectDelay.Bind_get_feedback_lowpass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectDelay() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectDelay() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectDelay", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
