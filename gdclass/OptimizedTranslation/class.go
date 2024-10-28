package OptimizedTranslation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
An optimized translation, used by default for CSV Translations. Uses real-time compressed translations, which results in very small dictionaries.

*/
type Go [1]classdb.OptimizedTranslation

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
func (self Go) Generate(from gdclass.Translation) {
	class(self).Generate(from)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OptimizedTranslation
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OptimizedTranslation"))
	return Go{classdb.OptimizedTranslation(object)}
}

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
//go:nosplit
func (self class) Generate(from gdclass.Translation)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(from[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptimizedTranslation.Bind_generate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("OptimizedTranslation", func(ptr gd.Object) any { return classdb.OptimizedTranslation(ptr) })}
