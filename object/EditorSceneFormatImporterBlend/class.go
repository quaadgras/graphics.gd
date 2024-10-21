package EditorSceneFormatImporterBlend

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/EditorSceneFormatImporter"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Imports Blender scenes in the [code].blend[/code] file format through the glTF 2.0 3D import pipeline. This importer requires Blender to be installed by the user, so that it can be used to export the scene as glTF 2.0.
The location of the Blender binary is set via the [code]filesystem/import/blender/blender_path[/code] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/blender/enabled] is enabled, otherwise [code].blend[/code] files present in the project folder are not imported.
Blend import requires Blender 3.0.
Internally, the EditorSceneFormatImporterBlend uses the Blender glTF "Use Original" mode to reference external textures.

*/
type Simple [1]classdb.EditorSceneFormatImporterBlend
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorSceneFormatImporterBlend
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsEditorSceneFormatImporterBlend() Expert { return self[0].AsEditorSceneFormatImporterBlend() }


//go:nosplit
func (self Simple) AsEditorSceneFormatImporterBlend() Simple { return self[0].AsEditorSceneFormatImporterBlend() }


//go:nosplit
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Expert { return self[0].AsEditorSceneFormatImporter() }


//go:nosplit
func (self Simple) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Simple { return self[0].AsEditorSceneFormatImporter() }


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
func init() {classdb.Register("EditorSceneFormatImporterBlend", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
