// Package EditorExportPlatformPC provides methods for working with EditorExportPlatformPC object instances.
package EditorExportPlatformPC

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/EditorExportPlatform"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The base class for the desktop platform exporters. These include Windows and Linux/BSD, but not macOS. See the classes inheriting this one for more details.
*/
type Instance [1]gdclass.EditorExportPlatformPC
type Any interface {
	gd.IsClass
	AsEditorExportPlatformPC() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorExportPlatformPC

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlatformPC"))
	return Instance{*(*gdclass.EditorExportPlatformPC)(unsafe.Pointer(&object))}
}

func (self class) AsEditorExportPlatformPC() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorExportPlatformPC() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorExportPlatform() EditorExportPlatform.Advanced {
	return *((*EditorExportPlatform.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatform() EditorExportPlatform.Instance {
	return *((*EditorExportPlatform.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(EditorExportPlatform.Advanced(self.AsEditorExportPlatform()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(EditorExportPlatform.Instance(self.AsEditorExportPlatform()), name)
	}
}
func init() {
	gdclass.Register("EditorExportPlatformPC", func(ptr gd.Object) any {
		return [1]gdclass.EditorExportPlatformPC{*(*gdclass.EditorExportPlatformPC)(unsafe.Pointer(&ptr))}
	})
}
