package AudioStreamPlaybackInteractive

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioStreamPlayback"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Playback component of [AudioStreamInteractive]. Contains functions to change the currently played clip.
*/
type Instance [1]classdb.AudioStreamPlaybackInteractive

/*
Switch to a clip (by name).
*/
func (self Instance) SwitchToClipByName(clip_name string) {
	class(self).SwitchToClipByName(gd.NewStringName(clip_name))
}

/*
Switch to a clip (by index).
*/
func (self Instance) SwitchToClip(clip_index int) {
	class(self).SwitchToClip(gd.Int(clip_index))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamPlaybackInteractive

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackInteractive"))
	return Instance{classdb.AudioStreamPlaybackInteractive(object)}
}

/*
Switch to a clip (by name).
*/
//go:nosplit
func (self class) SwitchToClipByName(clip_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(clip_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackInteractive.Bind_switch_to_clip_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Switch to a clip (by index).
*/
//go:nosplit
func (self class) SwitchToClip(clip_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackInteractive.Bind_switch_to_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsAudioStreamPlaybackInteractive() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlaybackInteractive() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Advanced {
	return *((*AudioStreamPlayback.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlayback() AudioStreamPlayback.Instance {
	return *((*AudioStreamPlayback.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStreamPlayback(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStreamPlayback(), name)
	}
}
func init() {
	classdb.Register("AudioStreamPlaybackInteractive", func(ptr gd.Object) any { return classdb.AudioStreamPlaybackInteractive(ptr) })
}
