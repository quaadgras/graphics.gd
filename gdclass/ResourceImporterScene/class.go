package ResourceImporterScene

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ResourceImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
See also [ResourceImporterOBJ], which is used for OBJ models that can be imported as an independent [Mesh] or a scene.
Additional options (such as extracting individual meshes or materials to files) are available in the [b]Advanced Import Settings[/b] dialog. This dialog can be accessed by double-clicking a 3D scene in the FileSystem dock or by selecting a 3D scene in the FileSystem dock, going to the Import dock and choosing [b]Advanced[/b].
[b]Note:[/b] [ResourceImporterScene] is [i]not[/i] used for [PackedScene]s, such as [code].tscn[/code] and [code].scn[/code] files.

*/
type Go [1]classdb.ResourceImporterScene
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ResourceImporterScene
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ResourceImporterScene"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsResourceImporterScene() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporterScene() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResourceImporter() ResourceImporter.GD { return *((*ResourceImporter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporter() ResourceImporter.Go { return *((*ResourceImporter.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}
func init() {classdb.Register("ResourceImporterScene", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
