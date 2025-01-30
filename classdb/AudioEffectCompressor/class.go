// Package AudioEffectCompressor provides methods for working with AudioEffectCompressor object instances.
package AudioEffectCompressor

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
Dynamic range compressor reduces the level of the sound when the amplitude goes over a certain threshold in Decibels. One of the main uses of a compressor is to increase the dynamic range by clipping as little as possible (when sound goes over 0dB).
Compressor has many uses in the mix:
- In the Master bus to compress the whole output (although an [AudioEffectLimiter] is probably better).
- In voice channels to ensure they sound as balanced as possible.
- Sidechained. This can reduce the sound level sidechained with another audio bus for threshold detection. This technique is common in video game mixing to the level of music and SFX while voices are being heard.
- Accentuates transients by using a wider attack, making effects sound more punchy.
*/
type Instance [1]gdclass.AudioEffectCompressor

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectCompressor() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectCompressor

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectCompressor"))
	casted := Instance{*(*gdclass.AudioEffectCompressor)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Threshold() Float.X {
	return Float.X(Float.X(class(self).GetThreshold()))
}

func (self Instance) SetThreshold(value Float.X) {
	class(self).SetThreshold(float64(value))
}

func (self Instance) Ratio() Float.X {
	return Float.X(Float.X(class(self).GetRatio()))
}

func (self Instance) SetRatio(value Float.X) {
	class(self).SetRatio(float64(value))
}

func (self Instance) Gain() Float.X {
	return Float.X(Float.X(class(self).GetGain()))
}

func (self Instance) SetGain(value Float.X) {
	class(self).SetGain(float64(value))
}

func (self Instance) AttackUs() Float.X {
	return Float.X(Float.X(class(self).GetAttackUs()))
}

func (self Instance) SetAttackUs(value Float.X) {
	class(self).SetAttackUs(float64(value))
}

func (self Instance) ReleaseMs() Float.X {
	return Float.X(Float.X(class(self).GetReleaseMs()))
}

func (self Instance) SetReleaseMs(value Float.X) {
	class(self).SetReleaseMs(float64(value))
}

func (self Instance) Mix() Float.X {
	return Float.X(Float.X(class(self).GetMix()))
}

func (self Instance) SetMix(value Float.X) {
	class(self).SetMix(float64(value))
}

func (self Instance) Sidechain() string {
	return string(class(self).GetSidechain().String())
}

func (self Instance) SetSidechain(value string) {
	class(self).SetSidechain(String.Name(String.New(value)))
}

//go:nosplit
func (self class) SetThreshold(threshold float64) { //gd:AudioEffectCompressor.set_threshold
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetThreshold() float64 { //gd:AudioEffectCompressor.get_threshold
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_threshold, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRatio(ratio float64) { //gd:AudioEffectCompressor.set_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRatio() float64 { //gd:AudioEffectCompressor.get_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGain(gain float64) { //gd:AudioEffectCompressor.set_gain
	var frame = callframe.New()
	callframe.Arg(frame, gain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGain() float64 { //gd:AudioEffectCompressor.get_gain
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttackUs(attack_us float64) { //gd:AudioEffectCompressor.set_attack_us
	var frame = callframe.New()
	callframe.Arg(frame, attack_us)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_attack_us, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttackUs() float64 { //gd:AudioEffectCompressor.get_attack_us
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_attack_us, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReleaseMs(release_ms float64) { //gd:AudioEffectCompressor.set_release_ms
	var frame = callframe.New()
	callframe.Arg(frame, release_ms)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_release_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReleaseMs() float64 { //gd:AudioEffectCompressor.get_release_ms
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_release_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMix(mix float64) { //gd:AudioEffectCompressor.set_mix
	var frame = callframe.New()
	callframe.Arg(frame, mix)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMix() float64 { //gd:AudioEffectCompressor.get_mix
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSidechain(sidechain String.Name) { //gd:AudioEffectCompressor.set_sidechain
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(sidechain)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_sidechain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSidechain() String.Name { //gd:AudioEffectCompressor.get_sidechain
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_sidechain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) AsAudioEffectCompressor() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectCompressor() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectCompressor", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectCompressor{*(*gdclass.AudioEffectCompressor)(unsafe.Pointer(&ptr))}
	})
}
