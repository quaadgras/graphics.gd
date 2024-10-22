package EditorExportPlatform

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base resource that provides the functionality of exporting a release build of a project to a platform, from the editor. Stores platform-specific metadata such as the name and supported features of the platform, and performs the exporting of projects, PCK files, and ZIP files. Uses an export template for the platform provided at the time of project exporting.
Used in scripting by [EditorExportPlugin] to configure platform-specific customization of scenes and resources. See [method EditorExportPlugin._begin_customize_scenes] and [method EditorExportPlugin._begin_customize_resources] for more details.

*/
type Go [1]classdb.EditorExportPlatform

/*
Returns the name of the export operating system handled by this [EditorExportPlatform] class, as a friendly string. Possible return values are [code]Windows[/code], [code]Linux[/code], [code]macOS[/code], [code]Android[/code], [code]iOS[/code], and [code]Web[/code].
*/
func (self Go) GetOsName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetOsName(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorExportPlatform
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorExportPlatform"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the name of the export operating system handled by this [EditorExportPlatform] class, as a friendly string. Possible return values are [code]Windows[/code], [code]Linux[/code], [code]macOS[/code], [code]Android[/code], [code]iOS[/code], and [code]Web[/code].
*/
//go:nosplit
func (self class) GetOsName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlatform.Bind_get_os_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorExportPlatform() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlatform() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorExportPlatform", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
