package AudioEffectEQ6

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
Band 1: 32 Hz
Band 2: 100 Hz
Band 3: 320 Hz
Band 4: 1000 Hz
Band 5: 3200 Hz
Band 6: 10000 Hz
See also [AudioEffectEQ], [AudioEffectEQ10], [AudioEffectEQ21].

*/
type Simple [1]classdb.AudioEffectEQ6
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioEffectEQ6
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsAudioEffectEQ6() Expert { return self[0].AsAudioEffectEQ6() }


//go:nosplit
func (self Simple) AsAudioEffectEQ6() Simple { return self[0].AsAudioEffectEQ6() }


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
func init() {classdb.Register("AudioEffectEQ6", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
