package ResourceSaver

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A singleton for saving resource types to the filesystem.
It uses the many [ResourceFormatSaver] classes registered in the engine (either built-in or from a plugin) to save resource data to text-based (e.g. [code].tres[/code] or [code].tscn[/code]) or binary files (e.g. [code].res[/code] or [code].scn[/code]).

*/
type Simple [1]classdb.ResourceSaver
func (self Simple) Save(resource [1]classdb.Resource, path string, flags classdb.ResourceSaverSaverFlags) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Save(resource, gc.String(path), flags))
}
func (self Simple) GetRecognizedExtensions(atype [1]classdb.Resource) gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetRecognizedExtensions(gc, atype))
}
func (self Simple) AddResourceFormatSaver(format_saver [1]classdb.ResourceFormatSaver, at_front bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddResourceFormatSaver(format_saver, at_front)
}
func (self Simple) RemoveResourceFormatSaver(format_saver [1]classdb.ResourceFormatSaver) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveResourceFormatSaver(format_saver)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ResourceSaver
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
//go:nosplit
func (self class) Save(resource object.Resource, path gd.String, flags classdb.ResourceSaverSaverFlags) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(resource[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceSaver.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of extensions available for saving a resource of a given type.
*/
//go:nosplit
func (self class) GetRecognizedExtensions(ctx gd.Lifetime, atype object.Resource) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(atype[0].AsPointer())[0])
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceSaver.Bind_get_recognized_extensions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
//go:nosplit
func (self class) AddResourceFormatSaver(format_saver object.ResourceFormatSaver, at_front bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(format_saver[0].AsPointer())[0])
	callframe.Arg(frame, at_front)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceSaver.Bind_add_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the given [ResourceFormatSaver].
*/
//go:nosplit
func (self class) RemoveResourceFormatSaver(format_saver object.ResourceFormatSaver)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(format_saver[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceSaver.Bind_remove_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsResourceSaver() Expert { return self[0].AsResourceSaver() }


//go:nosplit
func (self Simple) AsResourceSaver() Simple { return self[0].AsResourceSaver() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ResourceSaver", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type SaverFlags = classdb.ResourceSaverSaverFlags

const (
/*No resource saving option.*/
	FlagNone SaverFlags = 0
/*Save the resource with a path relative to the scene which uses it.*/
	FlagRelativePaths SaverFlags = 1
/*Bundles external resources.*/
	FlagBundleResources SaverFlags = 2
/*Changes the [member Resource.resource_path] of the saved resource to match its new location.*/
	FlagChangePath SaverFlags = 4
/*Do not save editor-specific metadata (identified by their [code]__editor[/code] prefix).*/
	FlagOmitEditorProperties SaverFlags = 8
/*Save as big endian (see [member FileAccess.big_endian]).*/
	FlagSaveBigEndian SaverFlags = 16
/*Compress the resource on save using [constant FileAccess.COMPRESSION_ZSTD]. Only available for binary resource types.*/
	FlagCompress SaverFlags = 32
/*Take over the paths of the saved subresources (see [method Resource.take_over_path]).*/
	FlagReplaceSubresourcePaths SaverFlags = 64
)
