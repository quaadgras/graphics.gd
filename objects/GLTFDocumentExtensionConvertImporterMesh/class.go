package GLTFDocumentExtensionConvertImporterMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/GLTFDocumentExtension"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

type Instance [1]classdb.GLTFDocumentExtensionConvertImporterMesh

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFDocumentExtensionConvertImporterMesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFDocumentExtensionConvertImporterMesh"))
	return Instance{classdb.GLTFDocumentExtensionConvertImporterMesh(object)}
}

func (self class) AsGLTFDocumentExtensionConvertImporterMesh() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFDocumentExtensionConvertImporterMesh() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGLTFDocumentExtension() GLTFDocumentExtension.Advanced {
	return *((*GLTFDocumentExtension.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFDocumentExtension() GLTFDocumentExtension.Instance {
	return *((*GLTFDocumentExtension.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsGLTFDocumentExtension(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGLTFDocumentExtension(), name)
	}
}
func init() {
	classdb.Register("GLTFDocumentExtensionConvertImporterMesh", func(ptr gd.Object) any { return classdb.GLTFDocumentExtensionConvertImporterMesh(ptr) })
}
