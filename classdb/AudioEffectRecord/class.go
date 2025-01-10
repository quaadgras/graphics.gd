// Package AudioEffectRecord provides methods for working with AudioEffectRecord object instances.
package AudioEffectRecord

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Allows the user to record the sound from an audio bus into an [AudioStreamWAV]. When used on the "Master" audio bus, this includes all audio output by Godot.
Unlike [AudioEffectCapture], this effect encodes the recording with the given format (8-bit, 16-bit, or compressed) instead of giving access to the raw audio samples.
Can be used (with an [AudioStreamMicrophone]) to record from a microphone.
[b]Note:[/b] [member ProjectSettings.audio/driver/enable_input] must be [code]true[/code] for audio input to work. See also that setting's description for caveats related to permissions and operating system privacy settings.
*/
type Instance [1]gdclass.AudioEffectRecord

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectRecord() Instance
}

/*
If [code]true[/code], the sound will be recorded. Note that restarting the recording will remove the previously recorded sample.
*/
func (self Instance) SetRecordingActive(record bool) {
	class(self).SetRecordingActive(record)
}

/*
Returns whether the recording is active or not.
*/
func (self Instance) IsRecordingActive() bool {
	return bool(class(self).IsRecordingActive())
}

/*
Returns the recorded sample.
*/
func (self Instance) GetRecording() [1]gdclass.AudioStreamWAV {
	return [1]gdclass.AudioStreamWAV(class(self).GetRecording())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectRecord

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectRecord"))
	casted := Instance{*(*gdclass.AudioEffectRecord)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Format() gdclass.AudioStreamWAVFormat {
	return gdclass.AudioStreamWAVFormat(class(self).GetFormat())
}

func (self Instance) SetFormat(value gdclass.AudioStreamWAVFormat) {
	class(self).SetFormat(value)
}

/*
If [code]true[/code], the sound will be recorded. Note that restarting the recording will remove the previously recorded sample.
*/
//go:nosplit
func (self class) SetRecordingActive(record bool) {
	var frame = callframe.New()
	callframe.Arg(frame, record)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectRecord.Bind_set_recording_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the recording is active or not.
*/
//go:nosplit
func (self class) IsRecordingActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectRecord.Bind_is_recording_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFormat(format gdclass.AudioStreamWAVFormat) {
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectRecord.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() gdclass.AudioStreamWAVFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamWAVFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectRecord.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the recorded sample.
*/
//go:nosplit
func (self class) GetRecording() [1]gdclass.AudioStreamWAV {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectRecord.Bind_get_recording, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStreamWAV{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamWAV](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsAudioEffectRecord() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectRecord() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectRecord", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectRecord{*(*gdclass.AudioEffectRecord)(unsafe.Pointer(&ptr))}
	})
}
