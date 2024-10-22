package AudioStreamPolyphonic

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStream"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
AudioStream that lets the user play custom streams at any time from code, simultaneously using a single player.
Playback control is done via the [AudioStreamPlaybackPolyphonic] instance set inside the player, which can be obtained via [method AudioStreamPlayer.get_stream_playback], [method AudioStreamPlayer2D.get_stream_playback] or [method AudioStreamPlayer3D.get_stream_playback] methods. Obtaining the playback instance is only valid after the [code]stream[/code] property is set as an [AudioStreamPolyphonic] in those players.

*/
type Go [1]classdb.AudioStreamPolyphonic
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPolyphonic
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamPolyphonic"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Polyphony() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetPolyphony()))
}

func (self Go) SetPolyphony(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPolyphony(gd.Int(value))
}

//go:nosplit
func (self class) SetPolyphony(voices gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, voices)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPolyphonic.Bind_set_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolyphony() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPolyphonic.Bind_get_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPolyphonic() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPolyphonic() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.GD { return *((*AudioStream.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStream() AudioStream.Go { return *((*AudioStream.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {classdb.Register("AudioStreamPolyphonic", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
