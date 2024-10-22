package OptimizedTranslation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Translation"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An optimized translation, used by default for CSV Translations. Uses real-time compressed translations, which results in very small dictionaries.

*/
type Go [1]classdb.OptimizedTranslation

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
func (self Go) Generate(from gdclass.Translation) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Generate(from)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OptimizedTranslation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OptimizedTranslation"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
//go:nosplit
func (self class) Generate(from gdclass.Translation)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptimizedTranslation.Bind_generate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOptimizedTranslation() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOptimizedTranslation() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTranslation() Translation.GD { return *((*Translation.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTranslation() Translation.Go { return *((*Translation.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTranslation(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTranslation(), name)
	}
}
func init() {classdb.Register("OptimizedTranslation", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
