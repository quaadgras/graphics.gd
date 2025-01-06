// Package ResourceImporter provides methods for working with ResourceImporter object instances.
package ResourceImporter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This is the base class for Godot's resource importers. To implement your own resource importers using editor plugins, see [EditorImportPlugin].
*/
type Instance [1]gdclass.ResourceImporter
type Any interface {
	gd.IsClass
	AsResourceImporter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourceImporter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporter"))
	casted := Instance{*(*gdclass.ResourceImporter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsResourceImporter() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceImporter() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("ResourceImporter", func(ptr gd.Object) any {
		return [1]gdclass.ResourceImporter{*(*gdclass.ResourceImporter)(unsafe.Pointer(&ptr))}
	})
}

type ImportOrder = gdclass.ResourceImporterImportOrder

const (
	/*The default import order.*/
	ImportOrderDefault ImportOrder = 0
	/*The import order for scenes, which ensures scenes are imported [i]after[/i] all other core resources such as textures. Custom importers should generally have an import order lower than [code]100[/code] to avoid issues when importing scenes that rely on custom resources.*/
	ImportOrderScene ImportOrder = 100
)
