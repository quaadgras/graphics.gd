package GLTFDocumentExtensionConvertImporterMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GLTFDocumentExtension"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Go [1]classdb.GLTFDocumentExtensionConvertImporterMesh
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFDocumentExtensionConvertImporterMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFDocumentExtensionConvertImporterMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsGLTFDocumentExtensionConvertImporterMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFDocumentExtensionConvertImporterMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGLTFDocumentExtension() GLTFDocumentExtension.GD { return *((*GLTFDocumentExtension.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFDocumentExtension() GLTFDocumentExtension.Go { return *((*GLTFDocumentExtension.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGLTFDocumentExtension(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGLTFDocumentExtension(), name)
	}
}
func init() {classdb.Register("GLTFDocumentExtensionConvertImporterMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
