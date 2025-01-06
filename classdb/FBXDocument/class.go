// Package FBXDocument provides methods for working with FBXDocument object instances.
package FBXDocument

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/GLTFDocument"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The FBXDocument handles FBX documents. It provides methods to append data from buffers or files, generate scenes, and register/unregister document extensions.
When exporting FBX from Blender, use the "FBX Units Scale" option. The "FBX Units Scale" option sets the correct scale factor and avoids manual adjustments when re-importing into Blender, such as through glTF export.
*/
type Instance [1]gdclass.FBXDocument
type Any interface {
	gd.IsClass
	AsFBXDocument() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FBXDocument

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FBXDocument"))
	casted := Instance{*(*gdclass.FBXDocument)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsFBXDocument() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFBXDocument() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGLTFDocument() GLTFDocument.Advanced {
	return *((*GLTFDocument.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFDocument() GLTFDocument.Instance {
	return *((*GLTFDocument.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(GLTFDocument.Advanced(self.AsGLTFDocument()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GLTFDocument.Instance(self.AsGLTFDocument()), name)
	}
}
func init() {
	gdclass.Register("FBXDocument", func(ptr gd.Object) any { return [1]gdclass.FBXDocument{*(*gdclass.FBXDocument)(unsafe.Pointer(&ptr))} })
}
