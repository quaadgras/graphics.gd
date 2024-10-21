package EditorSceneFormatImporterFBX2GLTF

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
Imports Autodesk FBX 3D scenes by way of converting them to glTF 2.0 using the FBX2glTF command line tool.
The location of the FBX2glTF binary is set via the [member EditorSettings.filesystem/import/fbx/fbx2gltf_path] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/fbx2gltf/enabled] is set to [code]true[/code].

*/
type Simple [1]classdb.EditorSceneFormatImporterFBX2GLTF
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorSceneFormatImporterFBX2GLTF
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsEditorSceneFormatImporterFBX2GLTF() Expert { return self[0].AsEditorSceneFormatImporterFBX2GLTF() }


//go:nosplit
func (self Simple) AsEditorSceneFormatImporterFBX2GLTF() Simple { return self[0].AsEditorSceneFormatImporterFBX2GLTF() }


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
func init() {classdb.Register("EditorSceneFormatImporterFBX2GLTF", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
