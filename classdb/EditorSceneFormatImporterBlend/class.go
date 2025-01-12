// Package EditorSceneFormatImporterBlend provides methods for working with EditorSceneFormatImporterBlend object instances.
package EditorSceneFormatImporterBlend

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/EditorSceneFormatImporter"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Imports Blender scenes in the [code].blend[/code] file format through the glTF 2.0 3D import pipeline. This importer requires Blender to be installed by the user, so that it can be used to export the scene as glTF 2.0.
The location of the Blender binary is set via the [code]filesystem/import/blender/blender_path[/code] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/blender/enabled] is enabled, otherwise [code].blend[/code] files present in the project folder are not imported.
Blend import requires Blender 3.0.
Internally, the EditorSceneFormatImporterBlend uses the Blender glTF "Use Original" mode to reference external textures.
*/
type Instance [1]gdclass.EditorSceneFormatImporterBlend

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSceneFormatImporterBlend() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSceneFormatImporterBlend

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterBlend"))
	casted := Instance{*(*gdclass.EditorSceneFormatImporterBlend)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsEditorSceneFormatImporterBlend() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporterBlend() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Advanced {
	return *((*EditorSceneFormatImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Instance {
	return *((*EditorSceneFormatImporter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(EditorSceneFormatImporter.Advanced(self.AsEditorSceneFormatImporter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(EditorSceneFormatImporter.Instance(self.AsEditorSceneFormatImporter()), name)
	}
}
func init() {
	gdclass.Register("EditorSceneFormatImporterBlend", func(ptr gd.Object) any {
		return [1]gdclass.EditorSceneFormatImporterBlend{*(*gdclass.EditorSceneFormatImporterBlend)(unsafe.Pointer(&ptr))}
	})
}
