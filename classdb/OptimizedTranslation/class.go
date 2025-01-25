// Package OptimizedTranslation provides methods for working with OptimizedTranslation object instances.
package OptimizedTranslation

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
import "graphics.gd/classdb/Translation"
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
An optimized translation, used by default for CSV Translations. Uses real-time compressed translations, which results in very small dictionaries.
*/
type Instance [1]gdclass.OptimizedTranslation

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOptimizedTranslation() Instance
}

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
func (self Instance) Generate(from [1]gdclass.Translation) {
	class(self).Generate(from)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OptimizedTranslation

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OptimizedTranslation"))
	casted := Instance{*(*gdclass.OptimizedTranslation)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Generates and sets an optimized translation from the given [Translation] resource.
*/
//go:nosplit
func (self class) Generate(from [1]gdclass.Translation) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptimizedTranslation.Bind_generate, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Translation.Advanced(self.AsTranslation()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Translation.Instance(self.AsTranslation()), name)
	}
}
func init() {
	gdclass.Register("OptimizedTranslation", func(ptr gd.Object) any {
		return [1]gdclass.OptimizedTranslation{*(*gdclass.OptimizedTranslation)(unsafe.Pointer(&ptr))}
	})
}
