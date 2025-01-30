// Package AudioEffectDelay provides methods for working with AudioEffectDelay object instances.
package AudioEffectDelay

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Plays input signal back after a period of time. The delayed signal may be played back multiple times to create the sound of a repeating, decaying echo. Delay effects range from a subtle echo effect to a pronounced blending of previous sounds with new sounds.
*/
type Instance [1]gdclass.AudioEffectDelay

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectDelay() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectDelay

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectDelay"))
	casted := Instance{*(*gdclass.AudioEffectDelay)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Dry() Float.X {
	return Float.X(Float.X(class(self).GetDry()))
}

func (self Instance) SetDry(value Float.X) {
	class(self).SetDry(float64(value))
}

func (self Instance) Tap1Active() bool {
	return bool(class(self).IsTap1Active())
}

func (self Instance) SetTap1Active(value bool) {
	class(self).SetTap1Active(value)
}

func (self Instance) Tap1DelayMs() Float.X {
	return Float.X(Float.X(class(self).GetTap1DelayMs()))
}

func (self Instance) SetTap1DelayMs(value Float.X) {
	class(self).SetTap1DelayMs(float64(value))
}

func (self Instance) Tap1LevelDb() Float.X {
	return Float.X(Float.X(class(self).GetTap1LevelDb()))
}

func (self Instance) SetTap1LevelDb(value Float.X) {
	class(self).SetTap1LevelDb(float64(value))
}

func (self Instance) Tap1Pan() Float.X {
	return Float.X(Float.X(class(self).GetTap1Pan()))
}

func (self Instance) SetTap1Pan(value Float.X) {
	class(self).SetTap1Pan(float64(value))
}

func (self Instance) Tap2Active() bool {
	return bool(class(self).IsTap2Active())
}

func (self Instance) SetTap2Active(value bool) {
	class(self).SetTap2Active(value)
}

func (self Instance) Tap2DelayMs() Float.X {
	return Float.X(Float.X(class(self).GetTap2DelayMs()))
}

func (self Instance) SetTap2DelayMs(value Float.X) {
	class(self).SetTap2DelayMs(float64(value))
}

func (self Instance) Tap2LevelDb() Float.X {
	return Float.X(Float.X(class(self).GetTap2LevelDb()))
}

func (self Instance) SetTap2LevelDb(value Float.X) {
	class(self).SetTap2LevelDb(float64(value))
}

func (self Instance) Tap2Pan() Float.X {
	return Float.X(Float.X(class(self).GetTap2Pan()))
}

func (self Instance) SetTap2Pan(value Float.X) {
	class(self).SetTap2Pan(float64(value))
}

func (self Instance) FeedbackActive() bool {
	return bool(class(self).IsFeedbackActive())
}

func (self Instance) SetFeedbackActive(value bool) {
	class(self).SetFeedbackActive(value)
}

func (self Instance) FeedbackDelayMs() Float.X {
	return Float.X(Float.X(class(self).GetFeedbackDelayMs()))
}

func (self Instance) SetFeedbackDelayMs(value Float.X) {
	class(self).SetFeedbackDelayMs(float64(value))
}

func (self Instance) FeedbackLevelDb() Float.X {
	return Float.X(Float.X(class(self).GetFeedbackLevelDb()))
}

func (self Instance) SetFeedbackLevelDb(value Float.X) {
	class(self).SetFeedbackLevelDb(float64(value))
}

func (self Instance) FeedbackLowpass() Float.X {
	return Float.X(Float.X(class(self).GetFeedbackLowpass()))
}

func (self Instance) SetFeedbackLowpass(value Float.X) {
	class(self).SetFeedbackLowpass(float64(value))
}

//go:nosplit
func (self class) SetDry(amount float64) { //gd:AudioEffectDelay.set_dry
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_dry, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDry() float64 { //gd:AudioEffectDelay.get_dry
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_dry, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap1Active(amount bool) { //gd:AudioEffectDelay.set_tap1_active
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap1_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTap1Active() bool { //gd:AudioEffectDelay.is_tap1_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_is_tap1_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap1DelayMs(amount float64) { //gd:AudioEffectDelay.set_tap1_delay_ms
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap1_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTap1DelayMs() float64 { //gd:AudioEffectDelay.get_tap1_delay_ms
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_tap1_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap1LevelDb(amount float64) { //gd:AudioEffectDelay.set_tap1_level_db
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap1_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTap1LevelDb() float64 { //gd:AudioEffectDelay.get_tap1_level_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_tap1_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap1Pan(amount float64) { //gd:AudioEffectDelay.set_tap1_pan
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap1_pan, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTap1Pan() float64 { //gd:AudioEffectDelay.get_tap1_pan
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_tap1_pan, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap2Active(amount bool) { //gd:AudioEffectDelay.set_tap2_active
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap2_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsTap2Active() bool { //gd:AudioEffectDelay.is_tap2_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_is_tap2_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap2DelayMs(amount float64) { //gd:AudioEffectDelay.set_tap2_delay_ms
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap2_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTap2DelayMs() float64 { //gd:AudioEffectDelay.get_tap2_delay_ms
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_tap2_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap2LevelDb(amount float64) { //gd:AudioEffectDelay.set_tap2_level_db
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap2_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTap2LevelDb() float64 { //gd:AudioEffectDelay.get_tap2_level_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_tap2_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTap2Pan(amount float64) { //gd:AudioEffectDelay.set_tap2_pan
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_tap2_pan, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTap2Pan() float64 { //gd:AudioEffectDelay.get_tap2_pan
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_tap2_pan, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFeedbackActive(amount bool) { //gd:AudioEffectDelay.set_feedback_active
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_feedback_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFeedbackActive() bool { //gd:AudioEffectDelay.is_feedback_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_is_feedback_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFeedbackDelayMs(amount float64) { //gd:AudioEffectDelay.set_feedback_delay_ms
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_feedback_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFeedbackDelayMs() float64 { //gd:AudioEffectDelay.get_feedback_delay_ms
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_feedback_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFeedbackLevelDb(amount float64) { //gd:AudioEffectDelay.set_feedback_level_db
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_feedback_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFeedbackLevelDb() float64 { //gd:AudioEffectDelay.get_feedback_level_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_feedback_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFeedbackLowpass(amount float64) { //gd:AudioEffectDelay.set_feedback_lowpass
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_set_feedback_lowpass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFeedbackLowpass() float64 { //gd:AudioEffectDelay.get_feedback_lowpass
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDelay.Bind_get_feedback_lowpass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectDelay() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectDelay() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectDelay", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectDelay{*(*gdclass.AudioEffectDelay)(unsafe.Pointer(&ptr))}
	})
}
