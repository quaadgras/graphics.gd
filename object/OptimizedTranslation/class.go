package OptimizedTranslation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Translation"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An optimized translation, used by default for CSV Translations. Uses real-time compressed translations, which results in very small dictionaries.

*/
type Simple [1]classdb.OptimizedTranslation
func (self Simple) Generate(from [1]classdb.Translation) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Generate(from)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OptimizedTranslation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
//go:nosplit
func (self class) Generate(from object.Translation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptimizedTranslation.Bind_generate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsOptimizedTranslation() Expert { return self[0].AsOptimizedTranslation() }


//go:nosplit
func (self Simple) AsOptimizedTranslation() Simple { return self[0].AsOptimizedTranslation() }


//go:nosplit
func (self class) AsTranslation() Translation.Expert { return self[0].AsTranslation() }


//go:nosplit
func (self Simple) AsTranslation() Translation.Simple { return self[0].AsTranslation() }


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
func init() {classdb.Register("OptimizedTranslation", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
