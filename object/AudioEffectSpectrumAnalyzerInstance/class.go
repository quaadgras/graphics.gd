package AudioEffectSpectrumAnalyzerInstance

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioEffectInstance"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The runtime part of an [AudioEffectSpectrumAnalyzer], which can be used to query the magnitude of a frequency range on its host bus.
An instance of this class can be acquired with [method AudioServer.get_bus_effect_instance].

*/
type Simple [1]classdb.AudioEffectSpectrumAnalyzerInstance
func (self Simple) GetMagnitudeForFrequencyRange(from_hz float64, to_hz float64, mode classdb.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMagnitudeForFrequencyRange(gd.Float(from_hz), gd.Float(to_hz), mode))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioEffectSpectrumAnalyzerInstance
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the magnitude of the frequencies from [param from_hz] to [param to_hz] in linear energy as a Vector2. The [code]x[/code] component of the return value represents the left stereo channel, and [code]y[/code] represents the right channel.
[param mode] determines how the frequency range will be processed. See [enum MagnitudeMode].
*/
//go:nosplit
func (self class) GetMagnitudeForFrequencyRange(from_hz gd.Float, to_hz gd.Float, mode classdb.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_hz)
	callframe.Arg(frame, to_hz)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectSpectrumAnalyzerInstance.Bind_get_magnitude_for_frequency_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAudioEffectSpectrumAnalyzerInstance() Expert { return self[0].AsAudioEffectSpectrumAnalyzerInstance() }


//go:nosplit
func (self Simple) AsAudioEffectSpectrumAnalyzerInstance() Simple { return self[0].AsAudioEffectSpectrumAnalyzerInstance() }


//go:nosplit
func (self class) AsAudioEffectInstance() AudioEffectInstance.Expert { return self[0].AsAudioEffectInstance() }


//go:nosplit
func (self Simple) AsAudioEffectInstance() AudioEffectInstance.Simple { return self[0].AsAudioEffectInstance() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioEffectSpectrumAnalyzerInstance", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type MagnitudeMode = classdb.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode

const (
/*Use the average value across the frequency range as magnitude.*/
	MagnitudeAverage MagnitudeMode = 0
/*Use the maximum value of the frequency range as magnitude.*/
	MagnitudeMax MagnitudeMode = 1
)
