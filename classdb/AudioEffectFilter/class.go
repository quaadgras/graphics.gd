// Package AudioEffectFilter provides methods for working with AudioEffectFilter object instances.
package AudioEffectFilter

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
Allows frequencies other than the [member cutoff_hz] to pass.
*/
type Instance [1]gdclass.AudioEffectFilter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectFilter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectFilter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectFilter"))
	casted := Instance{*(*gdclass.AudioEffectFilter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) CutoffHz() Float.X {
	return Float.X(Float.X(class(self).GetCutoff()))
}

func (self Instance) SetCutoffHz(value Float.X) {
	class(self).SetCutoff(gd.Float(value))
}

func (self Instance) Resonance() Float.X {
	return Float.X(Float.X(class(self).GetResonance()))
}

func (self Instance) SetResonance(value Float.X) {
	class(self).SetResonance(gd.Float(value))
}

func (self Instance) Gain() Float.X {
	return Float.X(Float.X(class(self).GetGain()))
}

func (self Instance) SetGain(value Float.X) {
	class(self).SetGain(gd.Float(value))
}

func (self Instance) Db() gdclass.AudioEffectFilterFilterDB {
	return gdclass.AudioEffectFilterFilterDB(class(self).GetDb())
}

func (self Instance) SetDb(value gdclass.AudioEffectFilterFilterDB) {
	class(self).SetDb(value)
}

//go:nosplit
func (self class) SetCutoff(freq gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, freq)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_set_cutoff, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCutoff() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_get_cutoff, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResonance(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_set_resonance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetResonance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_get_resonance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGain(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_set_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGain() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_get_gain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDb(amount gdclass.AudioEffectFilterFilterDB) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_set_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDb() gdclass.AudioEffectFilterFilterDB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioEffectFilterFilterDB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectFilter.Bind_get_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectFilter() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectFilter() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectFilter", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectFilter{*(*gdclass.AudioEffectFilter)(unsafe.Pointer(&ptr))}
	})
}

type FilterDB = gdclass.AudioEffectFilterFilterDB

const (
	Filter6db  FilterDB = 0
	Filter12db FilterDB = 1
	Filter18db FilterDB = 2
	Filter24db FilterDB = 3
)
