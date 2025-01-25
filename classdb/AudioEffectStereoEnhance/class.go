// Package AudioEffectStereoEnhance provides methods for working with AudioEffectStereoEnhance object instances.
package AudioEffectStereoEnhance

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
An audio effect that can be used to adjust the intensity of stereo panning.
*/
type Instance [1]gdclass.AudioEffectStereoEnhance

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectStereoEnhance() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectStereoEnhance

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectStereoEnhance"))
	casted := Instance{*(*gdclass.AudioEffectStereoEnhance)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) PanPullout() Float.X {
	return Float.X(Float.X(class(self).GetPanPullout()))
}

func (self Instance) SetPanPullout(value Float.X) {
	class(self).SetPanPullout(gd.Float(value))
}

func (self Instance) TimePulloutMs() Float.X {
	return Float.X(Float.X(class(self).GetTimePullout()))
}

func (self Instance) SetTimePulloutMs(value Float.X) {
	class(self).SetTimePullout(gd.Float(value))
}

func (self Instance) Surround() Float.X {
	return Float.X(Float.X(class(self).GetSurround()))
}

func (self Instance) SetSurround(value Float.X) {
	class(self).SetSurround(gd.Float(value))
}

//go:nosplit
func (self class) SetPanPullout(amount gd.Float) { //gd:AudioEffectStereoEnhance.set_pan_pullout
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectStereoEnhance.Bind_set_pan_pullout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPanPullout() gd.Float { //gd:AudioEffectStereoEnhance.get_pan_pullout
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectStereoEnhance.Bind_get_pan_pullout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimePullout(amount gd.Float) { //gd:AudioEffectStereoEnhance.set_time_pullout
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectStereoEnhance.Bind_set_time_pullout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimePullout() gd.Float { //gd:AudioEffectStereoEnhance.get_time_pullout
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectStereoEnhance.Bind_get_time_pullout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSurround(amount gd.Float) { //gd:AudioEffectStereoEnhance.set_surround
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectStereoEnhance.Bind_set_surround, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSurround() gd.Float { //gd:AudioEffectStereoEnhance.get_surround
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectStereoEnhance.Bind_get_surround, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectStereoEnhance() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectStereoEnhance() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	gdclass.Register("AudioEffectStereoEnhance", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectStereoEnhance{*(*gdclass.AudioEffectStereoEnhance)(unsafe.Pointer(&ptr))}
	})
}
