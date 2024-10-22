package EditorSceneFormatImporterFBX2GLTF

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/EditorSceneFormatImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Imports Autodesk FBX 3D scenes by way of converting them to glTF 2.0 using the FBX2glTF command line tool.
The location of the FBX2glTF binary is set via the [member EditorSettings.filesystem/import/fbx/fbx2gltf_path] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/fbx2gltf/enabled] is set to [code]true[/code].

*/
type Go [1]classdb.EditorSceneFormatImporterFBX2GLTF
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSceneFormatImporterFBX2GLTF
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorSceneFormatImporterFBX2GLTF"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsEditorSceneFormatImporterFBX2GLTF() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporterFBX2GLTF() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.GD { return *((*EditorSceneFormatImporter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Go { return *((*EditorSceneFormatImporter.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorSceneFormatImporter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorSceneFormatImporter(), name)
	}
}
func init() {classdb.Register("EditorSceneFormatImporterFBX2GLTF", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
