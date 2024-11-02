package EditorExportPlatformLinuxBSD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/EditorExportPlatformPC"
import "grow.graphics/gd/objects/EditorExportPlatform"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

type Instance [1]classdb.EditorExportPlatformLinuxBSD

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorExportPlatformLinuxBSD

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlatformLinuxBSD"))
	return Instance{classdb.EditorExportPlatformLinuxBSD(object)}
}

func (self class) AsEditorExportPlatformLinuxBSD() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatformLinuxBSD() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorExportPlatformPC() EditorExportPlatformPC.Advanced {
	return *((*EditorExportPlatformPC.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatformPC() EditorExportPlatformPC.Instance {
	return *((*EditorExportPlatformPC.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorExportPlatform() EditorExportPlatform.Advanced {
	return *((*EditorExportPlatform.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatform() EditorExportPlatform.Instance {
	return *((*EditorExportPlatform.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorExportPlatformPC(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorExportPlatformPC(), name)
	}
}
func init() {
	classdb.Register("EditorExportPlatformLinuxBSD", func(ptr gd.Object) any { return classdb.EditorExportPlatformLinuxBSD(ptr) })
}
