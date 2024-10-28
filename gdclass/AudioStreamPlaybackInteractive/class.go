package AudioStreamPlaybackInteractive

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStreamPlayback"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Playback component of [AudioStreamInteractive]. Contains functions to change the currently played clip.

*/
type Go [1]classdb.AudioStreamPlaybackInteractive

/*
Switch to a clip (by name).
*/
func (self Go) SwitchToClipByName(clip_name string) {
	class(self).SwitchToClipByName(gd.NewStringName(clip_name))
}

/*
Switch to a clip (by index).
*/
func (self Go) SwitchToClip(clip_index int) {
	class(self).SwitchToClip(gd.Int(clip_index))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlaybackInteractive
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackInteractive"))
	return Go{classdb.AudioStreamPlaybackInteractive(object)}
}

/*
Switch to a clip (by name).
*/
//go:nosplit
func (self class) SwitchToClipByName(clip_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(clip_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackInteractive.Bind_switch_to_clip_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Switch to a clip (by index).
*/
//go:nosplit
func (self class) SwitchToClip(clip_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackInteractive.Bind_switch_to_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsAudioStreamPlaybackInteractive() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlaybackInteractive() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.GD { return *((*AudioStreamPlayback.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlayback() AudioStreamPlayback.Go { return *((*AudioStreamPlayback.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStreamPlayback(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStreamPlayback(), name)
	}
}
func init() {classdb.Register("AudioStreamPlaybackInteractive", func(ptr gd.Object) any { return classdb.AudioStreamPlaybackInteractive(ptr) })}
