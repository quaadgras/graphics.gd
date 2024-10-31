package ResourceImporterOBJ

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ResourceImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Unlike [ResourceImporterScene], [ResourceImporterOBJ] will import a single [Mesh] resource by default instead of importing a [PackedScene]. This makes it easier to use the [Mesh] resource in nodes that expect direct [Mesh] resources, such as [GridMap], [GPUParticles3D] or [CPUParticles3D]. Note that it is still possible to save mesh resources from 3D scenes using the [b]Advanced Import Settings[/b] dialog, regardless of the source format.
See also [ResourceImporterScene], which is used for more advanced 3D formats such as glTF.
*/
type Instance [1]classdb.ResourceImporterOBJ

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceImporterOBJ

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterOBJ"))
	return Instance{classdb.ResourceImporterOBJ(object)}
}

func (self class) AsResourceImporterOBJ() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceImporterOBJ() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("ResourceImporterOBJ", func(ptr gd.Object) any { return classdb.ResourceImporterOBJ(ptr) })
}
