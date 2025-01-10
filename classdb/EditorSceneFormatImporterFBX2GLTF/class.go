// Package EditorSceneFormatImporterFBX2GLTF provides methods for working with EditorSceneFormatImporterFBX2GLTF object instances.
package EditorSceneFormatImporterFBX2GLTF

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

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Imports Autodesk FBX 3D scenes by way of converting them to glTF 2.0 using the FBX2glTF command line tool.
The location of the FBX2glTF binary is set via the [member EditorSettings.filesystem/import/fbx/fbx2gltf_path] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/fbx2gltf/enabled] is set to [code]true[/code].
*/
type Instance [1]gdclass.EditorSceneFormatImporterFBX2GLTF

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSceneFormatImporterFBX2GLTF() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSceneFormatImporterFBX2GLTF

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterFBX2GLTF"))
	casted := Instance{*(*gdclass.EditorSceneFormatImporterFBX2GLTF)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsEditorSceneFormatImporterFBX2GLTF() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporterFBX2GLTF() Instance {
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
	gdclass.Register("EditorSceneFormatImporterFBX2GLTF", func(ptr gd.Object) any {
		return [1]gdclass.EditorSceneFormatImporterFBX2GLTF{*(*gdclass.EditorSceneFormatImporterFBX2GLTF)(unsafe.Pointer(&ptr))}
	})
}
