package AudioEffectEQ

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
AudioEffectEQ gives you control over frequencies. Use it to compensate for existing deficiencies in audio. AudioEffectEQs are useful on the Master bus to completely master a mix and give it more character. They are also useful when a game is run on a mobile device, to adjust the mix to that kind of speakers (it can be added but disabled when headphones are plugged).

*/
type Go [1]classdb.AudioEffectEQ

/*
Sets band's gain at the specified index, in dB.
*/
func (self Go) SetBandGainDb(band_idx int, volume_db float64) {
	class(self).SetBandGainDb(gd.Int(band_idx), gd.Float(volume_db))
}

/*
Returns the band's gain at the specified index, in dB.
*/
func (self Go) GetBandGainDb(band_idx int) float64 {
	return float64(float64(class(self).GetBandGainDb(gd.Int(band_idx))))
}

/*
Returns the number of bands of the equalizer.
*/
func (self Go) GetBandCount() int {
	return int(int(class(self).GetBandCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectEQ
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ"))
	return Go{classdb.AudioEffectEQ(object)}
}

/*
Sets band's gain at the specified index, in dB.
*/
//go:nosplit
func (self class) SetBandGainDb(band_idx gd.Int, volume_db gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, band_idx)
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectEQ.Bind_set_band_gain_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the band's gain at the specified index, in dB.
*/
//go:nosplit
func (self class) GetBandGainDb(band_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, band_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectEQ.Bind_get_band_gain_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of bands of the equalizer.
*/
//go:nosplit
func (self class) GetBandCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectEQ.Bind_get_band_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectEQ() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectEQ() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectEQ", func(ptr gd.Object) any { return classdb.AudioEffectEQ(ptr) })}
