package ResourceImporterOBJ

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ResourceImporter"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Unlike [ResourceImporterScene], [ResourceImporterOBJ] will import a single [Mesh] resource by default instead of importing a [PackedScene]. This makes it easier to use the [Mesh] resource in nodes that expect direct [Mesh] resources, such as [GridMap], [GPUParticles3D] or [CPUParticles3D]. Note that it is still possible to save mesh resources from 3D scenes using the [b]Advanced Import Settings[/b] dialog, regardless of the source format.
See also [ResourceImporterScene], which is used for more advanced 3D formats such as glTF.

*/
type Simple [1]classdb.ResourceImporterOBJ
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ResourceImporterOBJ
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsResourceImporterOBJ() Expert { return self[0].AsResourceImporterOBJ() }


//go:nosplit
func (self Simple) AsResourceImporterOBJ() Simple { return self[0].AsResourceImporterOBJ() }


//go:nosplit
func (self class) AsResourceImporter() ResourceImporter.Expert { return self[0].AsResourceImporter() }


//go:nosplit
func (self Simple) AsResourceImporter() ResourceImporter.Simple { return self[0].AsResourceImporter() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ResourceImporterOBJ", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
