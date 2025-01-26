// Package AudioEffectEQ provides methods for working with AudioEffectEQ object instances.
package AudioEffectEQ

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
import "graphics.gd/variant/Dictionary"
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
var _ Dictionary.Any

/*
AudioEffectEQ gives you control over frequencies. Use it to compensate for existing deficiencies in audio. AudioEffectEQs are useful on the Master bus to completely master a mix and give it more character. They are also useful when a game is run on a mobile device, to adjust the mix to that kind of speakers (it can be added but disabled when headphones are plugged).
*/
type Instance [1]gdclass.AudioEffectEQ

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectEQ() Instance
}

/*
Sets band's gain at the specified index, in dB.
*/
func (self Instance) SetBandGainDb(band_idx int, volume_db Float.X) { //gd:AudioEffectEQ.set_band_gain_db
	class(self).SetBandGainDb(gd.Int(band_idx), gd.Float(volume_db))
}

/*
Returns the band's gain at the specified index, in dB.
*/
func (self Instance) GetBandGainDb(band_idx int) Float.X { //gd:AudioEffectEQ.get_band_gain_db
	return Float.X(Float.X(class(self).GetBandGainDb(gd.Int(band_idx))))
}

/*
Returns the number of bands of the equalizer.
*/
func (self Instance) GetBandCount() int { //gd:AudioEffectEQ.get_band_count
	return int(int(class(self).GetBandCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectEQ

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ"))
	casted := Instance{*(*gdclass.AudioEffectEQ)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sets band's gain at the specified index, in dB.
*/
//go:nosplit
func (self class) SetBandGainDb(band_idx gd.Int, volume_db gd.Float) { //gd:AudioEffectEQ.set_band_gain_db
	var frame = callframe.New()
	callframe.Arg(frame, band_idx)
	callframe.Arg(frame, volume_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectEQ.Bind_set_band_gain_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the band's gain at the specified index, in dB.
*/
//go:nosplit
func (self class) GetBandGainDb(band_idx gd.Int) gd.Float { //gd:AudioEffectEQ.get_band_gain_db
	var frame = callframe.New()
	callframe.Arg(frame, band_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectEQ.Bind_get_band_gain_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of bands of the equalizer.
*/
//go:nosplit
func (self class) GetBandCount() gd.Int { //gd:AudioEffectEQ.get_band_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectEQ.Bind_get_band_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectEQ() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectEQ() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectEQ", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectEQ{*(*gdclass.AudioEffectEQ)(unsafe.Pointer(&ptr))}
	})
}
