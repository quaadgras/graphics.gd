package ResourceImporterShaderFile

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/ResourceImporter"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This imports native GLSL shaders as [RDShaderFile] resources, for use with low-level [RenderingDevice] operations. This importer does [i]not[/i] handle [code].gdshader[/code] files.
*/
type Instance [1]classdb.ResourceImporterShaderFile

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceImporterShaderFile

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterShaderFile"))
	return Instance{classdb.ResourceImporterShaderFile(object)}
}

func (self class) AsResourceImporterShaderFile() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporterShaderFile() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResourceImporter() ResourceImporter.Advanced {
	return *((*ResourceImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResourceImporter() ResourceImporter.Instance {
	return *((*ResourceImporter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}
func init() {
	classdb.Register("ResourceImporterShaderFile", func(ptr gd.Object) any { return classdb.ResourceImporterShaderFile(ptr) })
}
