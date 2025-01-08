// Package ResourceImporterImageFont provides methods for working with ResourceImporterImageFont object instances.
package ResourceImporterImageFont

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/ResourceImporter"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This image-based workflow can be easier to use than [ResourceImporterBMFont], but it requires all glyphs to have the same width and height, glyph advances and drawing offsets can be customized. This makes [ResourceImporterImageFont] most suited to fixed-width fonts.
See also [ResourceImporterDynamicFont].
*/
type Instance [1]gdclass.ResourceImporterImageFont

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsResourceImporterImageFont() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourceImporterImageFont

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterImageFont"))
	casted := Instance{*(*gdclass.ResourceImporterImageFont)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsResourceImporterImageFont() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporterImageFont() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResourceImporter() ResourceImporter.Advanced {
	return *((*ResourceImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporter() ResourceImporter.Instance {
	return *((*ResourceImporter.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(ResourceImporter.Advanced(self.AsResourceImporter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ResourceImporter.Instance(self.AsResourceImporter()), name)
	}
}
func init() {
	gdclass.Register("ResourceImporterImageFont", func(ptr gd.Object) any {
		return [1]gdclass.ResourceImporterImageFont{*(*gdclass.ResourceImporterImageFont)(unsafe.Pointer(&ptr))}
	})
}
