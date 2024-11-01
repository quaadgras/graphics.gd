package AudioEffectSpectrumAnalyzerInstance

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioEffectInstance"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
The runtime part of an [AudioEffectSpectrumAnalyzer], which can be used to query the magnitude of a frequency range on its host bus.
An instance of this class can be acquired with [method AudioServer.get_bus_effect_instance].
*/
type Instance [1]classdb.AudioEffectSpectrumAnalyzerInstance

/*
Returns the magnitude of the frequencies from [param from_hz] to [param to_hz] in linear energy as a Vector2. The [code]x[/code] component of the return value represents the left stereo channel, and [code]y[/code] represents the right channel.
[param mode] determines how the frequency range will be processed. See [enum MagnitudeMode].
*/
func (self Instance) GetMagnitudeForFrequencyRange(from_hz float64, to_hz float64) gd.Vector2 {
	return gd.Vector2(class(self).GetMagnitudeForFrequencyRange(gd.Float(from_hz), gd.Float(to_hz), 1))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioEffectSpectrumAnalyzerInstance

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectSpectrumAnalyzerInstance"))
	return Instance{classdb.AudioEffectSpectrumAnalyzerInstance(object)}
}

/*
Returns the magnitude of the frequencies from [param from_hz] to [param to_hz] in linear energy as a Vector2. The [code]x[/code] component of the return value represents the left stereo channel, and [code]y[/code] represents the right channel.
[param mode] determines how the frequency range will be processed. See [enum MagnitudeMode].
*/
//go:nosplit
func (self class) GetMagnitudeForFrequencyRange(from_hz gd.Float, to_hz gd.Float, mode classdb.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, from_hz)
	callframe.Arg(frame, to_hz)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzerInstance.Bind_get_magnitude_for_frequency_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectSpectrumAnalyzerInstance() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectSpectrumAnalyzerInstance() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioEffectInstance() AudioEffectInstance.Advanced {
	return *((*AudioEffectInstance.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectInstance() AudioEffectInstance.Instance {
	return *((*AudioEffectInstance.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffectInstance(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffectInstance(), name)
	}
}
func init() {
	classdb.Register("AudioEffectSpectrumAnalyzerInstance", func(ptr gd.Object) any { return classdb.AudioEffectSpectrumAnalyzerInstance(ptr) })
}

type MagnitudeMode = classdb.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode

const (
	/*Use the average value across the frequency range as magnitude.*/
	MagnitudeAverage MagnitudeMode = 0
	/*Use the maximum value of the frequency range as magnitude.*/
	MagnitudeMax MagnitudeMode = 1
)
