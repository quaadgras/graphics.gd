package EditorSceneFormatImporterBlend

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/EditorSceneFormatImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Imports Blender scenes in the [code].blend[/code] file format through the glTF 2.0 3D import pipeline. This importer requires Blender to be installed by the user, so that it can be used to export the scene as glTF 2.0.
The location of the Blender binary is set via the [code]filesystem/import/blender/blender_path[/code] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/blender/enabled] is enabled, otherwise [code].blend[/code] files present in the project folder are not imported.
Blend import requires Blender 3.0.
Internally, the EditorSceneFormatImporterBlend uses the Blender glTF "Use Original" mode to reference external textures.

*/
type Go [1]classdb.EditorSceneFormatImporterBlend
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSceneFormatImporterBlend
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterBlend"))
	return Go{classdb.EditorSceneFormatImporterBlend(object)}
}

func (self class) AsEditorSceneFormatImporterBlend() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporterBlend() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("EditorSceneFormatImporterBlend", func(ptr gd.Object) any { return classdb.EditorSceneFormatImporterBlend(ptr) })}
