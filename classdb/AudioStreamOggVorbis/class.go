// Package AudioStreamOggVorbis provides methods for working with AudioStreamOggVorbis object instances.
package AudioStreamOggVorbis

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The AudioStreamOggVorbis class is a specialized [AudioStream] for handling Ogg Vorbis file formats. It offers functionality for loading and playing back Ogg Vorbis files, as well as managing looping and other playback properties. This class is part of the audio stream system, which also supports WAV files through the [AudioStreamWAV] class.
*/
type Instance [1]gdclass.AudioStreamOggVorbis
type Any interface {
	gd.IsClass
	AsAudioStreamOggVorbis() Instance
}

/*
Creates a new AudioStreamOggVorbis instance from the given buffer. The buffer must contain Ogg Vorbis data.
*/
func LoadFromBuffer(buffer []byte) [1]gdclass.AudioStreamOggVorbis {
	self := Instance{}
	return [1]gdclass.AudioStreamOggVorbis(class(self).LoadFromBuffer(gd.NewPackedByteSlice(buffer)))
}

/*
Creates a new AudioStreamOggVorbis instance from the given file path. The file must be in Ogg Vorbis format.
*/
func LoadFromFile(path string) [1]gdclass.AudioStreamOggVorbis {
	self := Instance{}
	return [1]gdclass.AudioStreamOggVorbis(class(self).LoadFromFile(gd.NewString(path)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamOggVorbis

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamOggVorbis"))
	return Instance{*(*gdclass.AudioStreamOggVorbis)(unsafe.Pointer(&object))}
}

func (self Instance) PacketSequence() [1]gdclass.OggPacketSequence {
	return [1]gdclass.OggPacketSequence(class(self).GetPacketSequence())
}

func (self Instance) SetPacketSequence(value [1]gdclass.OggPacketSequence) {
	class(self).SetPacketSequence(value)
}

func (self Instance) Bpm() Float.X {
	return Float.X(Float.X(class(self).GetBpm()))
}

func (self Instance) SetBpm(value Float.X) {
	class(self).SetBpm(gd.Float(value))
}

func (self Instance) BeatCount() int {
	return int(int(class(self).GetBeatCount()))
}

func (self Instance) SetBeatCount(value int) {
	class(self).SetBeatCount(gd.Int(value))
}

func (self Instance) BarBeats() int {
	return int(int(class(self).GetBarBeats()))
}

func (self Instance) SetBarBeats(value int) {
	class(self).SetBarBeats(gd.Int(value))
}

func (self Instance) Loop() bool {
	return bool(class(self).HasLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

func (self Instance) LoopOffset() Float.X {
	return Float.X(Float.X(class(self).GetLoopOffset()))
}

func (self Instance) SetLoopOffset(value Float.X) {
	class(self).SetLoopOffset(gd.Float(value))
}

/*
Creates a new AudioStreamOggVorbis instance from the given buffer. The buffer must contain Ogg Vorbis data.
*/
//go:nosplit
func (self class) LoadFromBuffer(buffer gd.PackedByteArray) [1]gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_load_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStreamOggVorbis{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamOggVorbis](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new AudioStreamOggVorbis instance from the given file path. The file must be in Ogg Vorbis format.
*/
//go:nosplit
func (self class) LoadFromFile(path gd.String) [1]gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.AudioStreamOggVorbis{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamOggVorbis](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPacketSequence(packet_sequence [1]gdclass.OggPacketSequence) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(packet_sequence[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_set_packet_sequence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPacketSequence() [1]gdclass.OggPacketSequence {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_get_packet_sequence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.OggPacketSequence{gd.PointerWithOwnershipTransferredToGo[gdclass.OggPacketSequence](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoopOffset(seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_set_loop_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_get_loop_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBpm(bpm gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bpm)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_set_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBpm() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBeatCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_set_beat_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBeatCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_get_beat_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBarBeats(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_set_bar_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBarBeats() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamOggVorbis.Bind_get_bar_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamOggVorbis() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamOggVorbis() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.Advanced {
	return *((*AudioStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStream() AudioStream.Instance {
	return *((*AudioStream.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {
	gdclass.Register("AudioStreamOggVorbis", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamOggVorbis{*(*gdclass.AudioStreamOggVorbis)(unsafe.Pointer(&ptr))}
	})
}
