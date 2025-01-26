// Package AudioEffectEQ21 provides methods for working with AudioEffectEQ21 object instances.
package AudioEffectEQ21

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/AudioEffectEQ"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
Frequency bands:
Band 1: 22 Hz
Band 2: 32 Hz
Band 3: 44 Hz
Band 4: 63 Hz
Band 5: 90 Hz
Band 6: 125 Hz
Band 7: 175 Hz
Band 8: 250 Hz
Band 9: 350 Hz
Band 10: 500 Hz
Band 11: 700 Hz
Band 12: 1000 Hz
Band 13: 1400 Hz
Band 14: 2000 Hz
Band 15: 2800 Hz
Band 16: 4000 Hz
Band 17: 5600 Hz
Band 18: 8000 Hz
Band 19: 11000 Hz
Band 20: 16000 Hz
Band 21: 22000 Hz
See also [AudioEffectEQ], [AudioEffectEQ6], [AudioEffectEQ10].
*/
type Instance [1]gdclass.AudioEffectEQ21

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectEQ21() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectEQ21

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ21"))
	casted := Instance{*(*gdclass.AudioEffectEQ21)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsAudioEffectEQ21() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectEQ21() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffectEQ.Advanced(self.AsAudioEffectEQ()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffectEQ.Instance(self.AsAudioEffectEQ()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectEQ21", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectEQ21{*(*gdclass.AudioEffectEQ21)(unsafe.Pointer(&ptr))}
	})
}
