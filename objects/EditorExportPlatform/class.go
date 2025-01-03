package EditorExportPlatform

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Base resource that provides the functionality of exporting a release build of a project to a platform, from the editor. Stores platform-specific metadata such as the name and supported features of the platform, and performs the exporting of projects, PCK files, and ZIP files. Uses an export template for the platform provided at the time of project exporting.
Used in scripting by [EditorExportPlugin] to configure platform-specific customization of scenes and resources. See [method EditorExportPlugin._begin_customize_scenes] and [method EditorExportPlugin._begin_customize_resources] for more details.
*/
type Instance [1]classdb.EditorExportPlatform
type Any interface {
	gd.IsClass
	AsEditorExportPlatform() Instance
}

/*
Returns the name of the export operating system handled by this [EditorExportPlatform] class, as a friendly string. Possible return values are [code]Windows[/code], [code]Linux[/code], [code]macOS[/code], [code]Android[/code], [code]iOS[/code], and [code]Web[/code].
*/
func (self Instance) GetOsName() string {
	return string(class(self).GetOsName().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorExportPlatform

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlatform"))
	return Instance{*(*classdb.EditorExportPlatform)(unsafe.Pointer(&object))}
}

/*
Returns the name of the export operating system handled by this [EditorExportPlatform] class, as a friendly string. Possible return values are [code]Windows[/code], [code]Linux[/code], [code]macOS[/code], [code]Android[/code], [code]iOS[/code], and [code]Web[/code].
*/
//go:nosplit
func (self class) GetOsName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorExportPlatform.Bind_get_os_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorExportPlatform() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorExportPlatform() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted         { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorExportPlatform", func(ptr gd.Object) any {
		return [1]classdb.EditorExportPlatform{*(*classdb.EditorExportPlatform)(unsafe.Pointer(&ptr))}
	})
}
