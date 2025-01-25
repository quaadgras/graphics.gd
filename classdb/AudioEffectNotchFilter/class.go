// Package AudioEffectNotchFilter provides methods for working with AudioEffectNotchFilter object instances.
package AudioEffectNotchFilter

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
import "graphics.gd/classdb/AudioEffectFilter"
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

/*
Attenuates frequencies in a narrow band around the [member AudioEffectFilter.cutoff_hz] and cuts frequencies outside of this range.
*/
type Instance [1]gdclass.AudioEffectNotchFilter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectNotchFilter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectNotchFilter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectNotchFilter"))
	casted := Instance{*(*gdclass.AudioEffectNotchFilter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsAudioEffectNotchFilter() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectNotchFilter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioEffectFilter() AudioEffectFilter.Advanced {
	return *((*AudioEffectFilter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectFilter() AudioEffectFilter.Instance {
	return *((*AudioEffectFilter.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioEffectFilter.Advanced(self.AsAudioEffectFilter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffectFilter.Instance(self.AsAudioEffectFilter()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectNotchFilter", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectNotchFilter{*(*gdclass.AudioEffectNotchFilter)(unsafe.Pointer(&ptr))}
	})
}
