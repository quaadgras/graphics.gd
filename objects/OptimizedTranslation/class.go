package OptimizedTranslation

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Translation"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An optimized translation, used by default for CSV Translations. Uses real-time compressed translations, which results in very small dictionaries.
*/
type Instance [1]classdb.OptimizedTranslation

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
func (self Instance) Generate(from objects.Translation) {
	class(self).Generate(from)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OptimizedTranslation

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OptimizedTranslation"))
	return Instance{classdb.OptimizedTranslation(object)}
}

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
//go:nosplit
func (self class) Generate(from objects.Translation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptimizedTranslation.Bind_generate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOptimizedTranslation() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOptimizedTranslation() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTranslation() Translation.Advanced {
	return *((*Translation.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTranslation() Translation.Instance {
	return *((*Translation.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsTranslation(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTranslation(), name)
	}
}
func init() {
	classdb.Register("OptimizedTranslation", func(ptr gd.Object) any { return classdb.OptimizedTranslation(ptr) })
}
