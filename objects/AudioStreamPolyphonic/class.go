package AudioStreamPolyphonic

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioStream"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
AudioStream that lets the user play custom streams at any time from code, simultaneously using a single player.
Playback control is done via the [AudioStreamPlaybackPolyphonic] instance set inside the player, which can be obtained via [method AudioStreamPlayer.get_stream_playback], [method AudioStreamPlayer2D.get_stream_playback] or [method AudioStreamPlayer3D.get_stream_playback] methods. Obtaining the playback instance is only valid after the [code]stream[/code] property is set as an [AudioStreamPolyphonic] in those players.
*/
type Instance [1]classdb.AudioStreamPolyphonic

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamPolyphonic

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPolyphonic"))
	return Instance{classdb.AudioStreamPolyphonic(object)}
}

func (self Instance) Polyphony() int {
	return int(int(class(self).GetPolyphony()))
}

func (self Instance) SetPolyphony(value int) {
	class(self).SetPolyphony(gd.Int(value))
}

//go:nosplit
func (self class) SetPolyphony(voices gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, voices)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPolyphonic.Bind_set_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolyphony() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPolyphonic.Bind_get_polyphony, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPolyphonic() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPolyphonic() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("AudioStreamPolyphonic", func(ptr gd.Object) any { return classdb.AudioStreamPolyphonic(ptr) })
}
