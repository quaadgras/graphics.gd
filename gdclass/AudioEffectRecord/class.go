package AudioEffectRecord

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
Allows the user to record the sound from an audio bus into an [AudioStreamWAV]. When used on the "Master" audio bus, this includes all audio output by Godot.
Unlike [AudioEffectCapture], this effect encodes the recording with the given format (8-bit, 16-bit, or compressed) instead of giving access to the raw audio samples.
Can be used (with an [AudioStreamMicrophone]) to record from a microphone.
[b]Note:[/b] [member ProjectSettings.audio/driver/enable_input] must be [code]true[/code] for audio input to work. See also that setting's description for caveats related to permissions and operating system privacy settings.

*/
type Go [1]classdb.AudioEffectRecord

/*
If [code]true[/code], the sound will be recorded. Note that restarting the recording will remove the previously recorded sample.
*/
func (self Go) SetRecordingActive(record bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRecordingActive(record)
}

/*
Returns whether the recording is active or not.
*/
func (self Go) IsRecordingActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsRecordingActive())
}

/*
Returns the recorded sample.
*/
func (self Go) GetRecording() gdclass.AudioStreamWAV {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.AudioStreamWAV(class(self).GetRecording(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectRecord
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectRecord"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Format() classdb.AudioStreamWAVFormat {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.AudioStreamWAVFormat(class(self).GetFormat())
}

func (self Go) SetFormat(value classdb.AudioStreamWAVFormat) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFormat(value)
}

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
func (self class) GetRecording(ctx gd.Lifetime) gdclass.AudioStreamWAV {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectRecord.Bind_get_recording, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStreamWAV
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsAudioEffectRecord() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectRecord() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectRecord", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
