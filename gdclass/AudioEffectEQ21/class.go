package AudioEffectEQ21

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffectEQ"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.AudioEffectEQ21
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectEQ21
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectEQ21"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsAudioEffectEQ21() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectEQ21() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffectEQ() AudioEffectEQ.GD { return *((*AudioEffectEQ.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectEQ() AudioEffectEQ.Go { return *((*AudioEffectEQ.Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.GD { return *((*AudioEffect.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() AudioEffect.Go { return *((*AudioEffect.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffectEQ(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffectEQ(), name)
	}
}
func init() {classdb.Register("AudioEffectEQ21", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
