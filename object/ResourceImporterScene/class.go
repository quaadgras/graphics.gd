package ResourceImporterScene

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
See also [ResourceImporterOBJ], which is used for OBJ models that can be imported as an independent [Mesh] or a scene.
Additional options (such as extracting individual meshes or materials to files) are available in the [b]Advanced Import Settings[/b] dialog. This dialog can be accessed by double-clicking a 3D scene in the FileSystem dock or by selecting a 3D scene in the FileSystem dock, going to the Import dock and choosing [b]Advanced[/b].
[b]Note:[/b] [ResourceImporterScene] is [i]not[/i] used for [PackedScene]s, such as [code].tscn[/code] and [code].scn[/code] files.

*/
type Simple [1]classdb.ResourceImporterScene
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ResourceImporterScene
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsResourceImporterScene() Expert { return self[0].AsResourceImporterScene() }


//go:nosplit
func (self Simple) AsResourceImporterScene() Simple { return self[0].AsResourceImporterScene() }


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
func init() {classdb.Register("ResourceImporterScene", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
