// Package AudioStreamPlaybackInteractive provides methods for working with AudioStreamPlaybackInteractive object instances.
package AudioStreamPlaybackInteractive

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AudioStreamPlayback"

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
Playback component of [AudioStreamInteractive]. Contains functions to change the currently played clip.
*/
type Instance [1]gdclass.AudioStreamPlaybackInteractive

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlaybackInteractive() Instance
}

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
type class [1]gdclass.AudioStreamPlaybackInteractive

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackInteractive"))
	casted := Instance{*(*gdclass.AudioStreamPlaybackInteractive)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStreamPlayback.Advanced(self.AsAudioStreamPlayback()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStreamPlayback.Instance(self.AsAudioStreamPlayback()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlaybackInteractive", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlaybackInteractive{*(*gdclass.AudioStreamPlaybackInteractive)(unsafe.Pointer(&ptr))}
	})
}
