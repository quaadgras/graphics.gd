package AudioEffectRecord

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioEffect"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Allows the user to record the sound from an audio bus into an [AudioStreamWAV]. When used on the "Master" audio bus, this includes all audio output by Godot.
Unlike [AudioEffectCapture], this effect encodes the recording with the given format (8-bit, 16-bit, or compressed) instead of giving access to the raw audio samples.
Can be used (with an [AudioStreamMicrophone]) to record from a microphone.
[b]Note:[/b] [member ProjectSettings.audio/driver/enable_input] must be [code]true[/code] for audio input to work. See also that setting's description for caveats related to permissions and operating system privacy settings.

*/
type Simple [1]classdb.AudioEffectRecord
func (self Simple) SetRecordingActive(record bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRecordingActive(record)
}
func (self Simple) IsRecordingActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRecordingActive())
}
func (self Simple) SetFormat(format classdb.AudioStreamWAVFormat) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFormat(format)
}
func (self Simple) GetFormat() classdb.AudioStreamWAVFormat {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AudioStreamWAVFormat(Expert(self).GetFormat())
}
func (self Simple) GetRecording() [1]classdb.AudioStreamWAV {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AudioStreamWAV(Expert(self).GetRecording(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioEffectRecord
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
If [code]true[/code], the sound will be recorded. Note that restarting the recording will remove the previously recorded sample.
*/
//go:nosplit
func (self class) SetRecordingActive(record bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, record)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectRecord.Bind_set_recording_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the recording is active or not.
*/
//go:nosplit
func (self class) IsRecordingActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectRecord.Bind_is_recording_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFormat(format classdb.AudioStreamWAVFormat)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectRecord.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFormat() classdb.AudioStreamWAVFormat {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioStreamWAVFormat](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectRecord.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the recorded sample.
*/
//go:nosplit
func (self class) GetRecording(ctx gd.Lifetime) object.AudioStreamWAV {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectRecord.Bind_get_recording, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AudioStreamWAV
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAudioEffectRecord() Expert { return self[0].AsAudioEffectRecord() }


//go:nosplit
func (self Simple) AsAudioEffectRecord() Simple { return self[0].AsAudioEffectRecord() }


//go:nosplit
func (self class) AsAudioEffect() AudioEffect.Expert { return self[0].AsAudioEffect() }


//go:nosplit
func (self Simple) AsAudioEffect() AudioEffect.Simple { return self[0].AsAudioEffect() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioEffectRecord", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
