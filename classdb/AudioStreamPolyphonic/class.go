// Package AudioStreamPolyphonic provides methods for working with AudioStreamPolyphonic object instances.
package AudioStreamPolyphonic

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
AudioStream that lets the user play custom streams at any time from code, simultaneously using a single player.
Playback control is done via the [AudioStreamPlaybackPolyphonic] instance set inside the player, which can be obtained via [method AudioStreamPlayer.get_stream_playback], [method AudioStreamPlayer2D.get_stream_playback] or [method AudioStreamPlayer3D.get_stream_playback] methods. Obtaining the playback instance is only valid after the [code]stream[/code] property is set as an [AudioStreamPolyphonic] in those players.
*/
type Instance [1]gdclass.AudioStreamPolyphonic
type Any interface {
	gd.IsClass
	AsAudioStreamPolyphonic() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPolyphonic

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPolyphonic"))
	return Instance{*(*gdclass.AudioStreamPolyphonic)(unsafe.Pointer(&object))}
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Advanced(self.AsAudioStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Instance(self.AsAudioStream()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPolyphonic", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPolyphonic{*(*gdclass.AudioStreamPolyphonic)(unsafe.Pointer(&ptr))}
	})
}
