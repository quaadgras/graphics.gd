package EditorSceneFormatImporterUFBX

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
EditorSceneFormatImporterUFBX is designed to load FBX files and supports both binary and ASCII FBX files from version 3000 onward. This class supports various 3D object types like meshes, skins, blend shapes, materials, and rigging information. The class aims for feature parity with the official FBX SDK and supports FBX 7.4 specifications.

*/
type Go [1]classdb.EditorSceneFormatImporterUFBX
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSceneFormatImporterUFBX
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterUFBX"))
	return Go{classdb.EditorSceneFormatImporterUFBX(object)}
}

func (self class) AsEditorSceneFormatImporterUFBX() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSceneFormatImporterUFBX() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("EditorSceneFormatImporterUFBX", func(ptr gd.Object) any { return classdb.EditorSceneFormatImporterUFBX(ptr) })}
