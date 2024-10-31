package AudioEffectEQ21

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffectEQ"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

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
type Instance [1]classdb.AudioEffectEQ21

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioEffectEQ21

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ21"))
	return Instance{classdb.AudioEffectEQ21(object)}
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
	classdb.Register("AudioEffectEQ21", func(ptr gd.Object) any { return classdb.AudioEffectEQ21(ptr) })
}
