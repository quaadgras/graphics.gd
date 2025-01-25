// Package AudioEffectReverb provides methods for working with AudioEffectReverb object instances.
package AudioEffectReverb

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
Simulates the sound of acoustic environments such as rooms, concert halls, caverns, or an open spaces.
*/
type Instance [1]gdclass.AudioEffectReverb

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectReverb() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectReverb

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectReverb"))
	casted := Instance{*(*gdclass.AudioEffectReverb)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) PredelayMsec() Float.X {
	return Float.X(Float.X(class(self).GetPredelayMsec()))
}

func (self Instance) SetPredelayMsec(value Float.X) {
	class(self).SetPredelayMsec(gd.Float(value))
}

func (self Instance) PredelayFeedback() Float.X {
	return Float.X(Float.X(class(self).GetPredelayFeedback()))
}

func (self Instance) SetPredelayFeedback(value Float.X) {
	class(self).SetPredelayFeedback(gd.Float(value))
}

func (self Instance) RoomSize() Float.X {
	return Float.X(Float.X(class(self).GetRoomSize()))
}

func (self Instance) SetRoomSize(value Float.X) {
	class(self).SetRoomSize(gd.Float(value))
}

func (self Instance) Damping() Float.X {
	return Float.X(Float.X(class(self).GetDamping()))
}

func (self Instance) SetDamping(value Float.X) {
	class(self).SetDamping(gd.Float(value))
}

func (self Instance) Spread() Float.X {
	return Float.X(Float.X(class(self).GetSpread()))
}

func (self Instance) SetSpread(value Float.X) {
	class(self).SetSpread(gd.Float(value))
}

func (self Instance) Hipass() Float.X {
	return Float.X(Float.X(class(self).GetHpf()))
}

func (self Instance) SetHipass(value Float.X) {
	class(self).SetHpf(gd.Float(value))
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
func (self class) SetPredelayMsec(msec gd.Float) { //gd:AudioEffectReverb.set_predelay_msec
	var frame = callframe.New()
	callframe.Arg(frame, msec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_predelay_msec, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPredelayMsec() gd.Float { //gd:AudioEffectReverb.get_predelay_msec
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_predelay_msec, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPredelayFeedback(feedback gd.Float) { //gd:AudioEffectReverb.set_predelay_feedback
	var frame = callframe.New()
	callframe.Arg(frame, feedback)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_predelay_feedback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPredelayFeedback() gd.Float { //gd:AudioEffectReverb.get_predelay_feedback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_predelay_feedback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoomSize(size gd.Float) { //gd:AudioEffectReverb.set_room_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_room_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoomSize() gd.Float { //gd:AudioEffectReverb.get_room_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_room_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDamping(amount gd.Float) { //gd:AudioEffectReverb.set_damping
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDamping() gd.Float { //gd:AudioEffectReverb.get_damping
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpread(amount gd.Float) { //gd:AudioEffectReverb.set_spread
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpread() gd.Float { //gd:AudioEffectReverb.get_spread
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDry(amount gd.Float) { //gd:AudioEffectReverb.set_dry
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_dry, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDry() gd.Float { //gd:AudioEffectReverb.get_dry
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_dry, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWet(amount gd.Float) { //gd:AudioEffectReverb.set_wet
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_wet, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWet() gd.Float { //gd:AudioEffectReverb.get_wet
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_wet, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHpf(amount gd.Float) { //gd:AudioEffectReverb.set_hpf
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_set_hpf, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHpf() gd.Float { //gd:AudioEffectReverb.get_hpf
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectReverb.Bind_get_hpf, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectReverb() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectReverb() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectReverb", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectReverb{*(*gdclass.AudioEffectReverb)(unsafe.Pointer(&ptr))}
	})
}
