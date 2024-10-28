package AudioEffectEQ6

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

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
type Go [1]classdb.AudioEffectEQ6
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectEQ6
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ6"))
	return Go{classdb.AudioEffectEQ6(object)}
}

func (self class) AsAudioEffectEQ6() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectEQ6() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("AudioEffectEQ6", func(ptr gd.Object) any { return classdb.AudioEffectEQ6(ptr) })}
