package AudioEffectEQ6

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AudioEffectEQ"
import "graphics.gd/objects/AudioEffect"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Frequency bands:
Band 1: 32 Hz
Band 2: 100 Hz
Band 3: 320 Hz
Band 4: 1000 Hz
Band 5: 3200 Hz
Band 6: 10000 Hz
See also [AudioEffectEQ], [AudioEffectEQ10], [AudioEffectEQ21].
*/
type Instance [1]classdb.AudioEffectEQ6
type Any interface {
	gd.IsClass
	AsAudioEffectEQ6() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioEffectEQ6

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ6"))
	return Instance{*(*classdb.AudioEffectEQ6)(unsafe.Pointer(&object))}
}

func (self class) AsAudioEffectEQ6() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectEQ6() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffectEQ() AudioEffectEQ.Advanced {
	return *((*AudioEffectEQ.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectEQ() AudioEffectEQ.Instance {
	return *((*AudioEffectEQ.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioEffect() AudioEffect.Advanced {
	return *((*AudioEffect.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffect() AudioEffect.Instance {
	return *((*AudioEffect.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsAudioEffectEQ(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffectEQ(), name)
	}
}
func init() {
	classdb.Register("AudioEffectEQ6", func(ptr gd.Object) any {
		return [1]classdb.AudioEffectEQ6{*(*classdb.AudioEffectEQ6)(unsafe.Pointer(&ptr))}
	})
}
