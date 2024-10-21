package AudioEffectEQ21

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioEffectEQ"
import "grow.graphics/gd/object/AudioEffect"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
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
type Simple [1]classdb.AudioEffectEQ21
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioEffectEQ21
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsAudioEffectEQ21() Expert { return self[0].AsAudioEffectEQ21() }


//go:nosplit
func (self Simple) AsAudioEffectEQ21() Simple { return self[0].AsAudioEffectEQ21() }


//go:nosplit
func (self class) AsAudioEffectEQ() AudioEffectEQ.Expert { return self[0].AsAudioEffectEQ() }


//go:nosplit
func (self Simple) AsAudioEffectEQ() AudioEffectEQ.Simple { return self[0].AsAudioEffectEQ() }


//go:nosplit
func (self class) AsAudioEffect() AudioEffect.Expert { return self[0].AsAudioEffect() }


//go:nosplit
func (self Simple) AsAudioEffect() AudioEffect.Simple { return self[0].AsAudioEffect() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioEffectEQ21", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
