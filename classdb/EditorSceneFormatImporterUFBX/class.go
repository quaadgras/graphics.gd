// Package EditorSceneFormatImporterUFBX provides methods for working with EditorSceneFormatImporterUFBX object instances.
package EditorSceneFormatImporterUFBX

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
EditorSceneFormatImporterUFBX is designed to load FBX files and supports both binary and ASCII FBX files from version 3000 onward. This class supports various 3D object types like meshes, skins, blend shapes, materials, and rigging information. The class aims for feature parity with the official FBX SDK and supports FBX 7.4 specifications.
*/
type Instance [1]gdclass.EditorSceneFormatImporterUFBX
type Any interface {
	gd.IsClass
	AsEditorSceneFormatImporterUFBX() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSceneFormatImporterUFBX

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterUFBX"))
	return Instance{*(*gdclass.EditorSceneFormatImporterUFBX)(unsafe.Pointer(&object))}
}

func (self class) AsEditorSceneFormatImporterUFBX() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporterUFBX() Instance {
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
	gdclass.Register("EditorSceneFormatImporterUFBX", func(ptr gd.Object) any {
		return [1]gdclass.EditorSceneFormatImporterUFBX{*(*gdclass.EditorSceneFormatImporterUFBX)(unsafe.Pointer(&ptr))}
	})
}
