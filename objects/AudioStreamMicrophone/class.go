package AudioStreamMicrophone

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AudioStream"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
When used directly in an [AudioStreamPlayer] node, [AudioStreamMicrophone] plays back microphone input in real-time. This can be used in conjunction with [AudioEffectCapture] to process the data or save it.
[b]Note:[/b] [member ProjectSettings.audio/driver/enable_input] must be [code]true[/code] for audio input to work. See also that setting's description for caveats related to permissions and operating system privacy settings.
*/
type Instance [1]classdb.AudioStreamMicrophone
type Any interface {
	gd.IsClass
	AsAudioStreamMicrophone() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamMicrophone

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamMicrophone"))
	return Instance{*(*classdb.AudioStreamMicrophone)(unsafe.Pointer(&object))}
}

func (self class) AsAudioStreamMicrophone() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamMicrophone() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("AudioStreamMicrophone", func(ptr gd.Object) any {
		return [1]classdb.AudioStreamMicrophone{*(*classdb.AudioStreamMicrophone)(unsafe.Pointer(&ptr))}
	})
}
