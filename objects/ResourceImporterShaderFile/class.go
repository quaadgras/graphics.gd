package ResourceImporterShaderFile

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/ResourceImporter"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This imports native GLSL shaders as [RDShaderFile] resources, for use with low-level [RenderingDevice] operations. This importer does [i]not[/i] handle [code].gdshader[/code] files.
*/
type Instance [1]classdb.ResourceImporterShaderFile
type Any interface {
	gd.IsClass
	AsResourceImporterShaderFile() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceImporterShaderFile

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterShaderFile"))
	return Instance{*(*classdb.ResourceImporterShaderFile)(unsafe.Pointer(&object))}
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
	classdb.Register("ResourceImporterShaderFile", func(ptr gd.Object) any {
		return [1]classdb.ResourceImporterShaderFile{*(*classdb.ResourceImporterShaderFile)(unsafe.Pointer(&ptr))}
	})
}
