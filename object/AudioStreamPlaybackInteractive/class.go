package AudioStreamPlaybackInteractive

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioStreamPlayback"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Playback component of [AudioStreamInteractive]. Contains functions to change the currently played clip.

*/
type Simple [1]classdb.AudioStreamPlaybackInteractive
func (self Simple) SwitchToClipByName(clip_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SwitchToClipByName(gc.StringName(clip_name))
}
func (self Simple) SwitchToClip(clip_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SwitchToClip(gd.Int(clip_index))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamPlaybackInteractive
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Switch to a clip (by name).
*/
//go:nosplit
func (self class) SwitchToClipByName(clip_name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(clip_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackInteractive.Bind_switch_to_clip_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Switch to a clip (by index).
*/
//go:nosplit
func (self class) SwitchToClip(clip_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, clip_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackInteractive.Bind_switch_to_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsAudioStreamPlaybackInteractive() Expert { return self[0].AsAudioStreamPlaybackInteractive() }


//go:nosplit
func (self Simple) AsAudioStreamPlaybackInteractive() Simple { return self[0].AsAudioStreamPlaybackInteractive() }


//go:nosplit
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Expert { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self Simple) AsAudioStreamPlayback() AudioStreamPlayback.Simple { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStreamPlaybackInteractive", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
