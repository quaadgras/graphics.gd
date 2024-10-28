package ResourceSaver

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A singleton for saving resource types to the filesystem.
It uses the many [ResourceFormatSaver] classes registered in the engine (either built-in or from a plugin) to save resource data to text-based (e.g. [code].tres[/code] or [code].tscn[/code]) or binary files (e.g. [code].res[/code] or [code].scn[/code]).

*/
var self gdclass.ResourceSaver
var once sync.Once
func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceSaver)
	self = *(*gdclass.ResourceSaver)(unsafe.Pointer(&obj))
}

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
func Save(resource gdclass.Resource) gd.Error {
	once.Do(singleton)
	return gd.Error(class(self).Save(resource, gd.NewString(""), 0))
}

/*
Returns the list of extensions available for saving a resource of a given type.
*/
func GetRecognizedExtensions(atype gdclass.Resource) []string {
	once.Do(singleton)
	return []string(class(self).GetRecognizedExtensions(atype).Strings())
}

/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
func AddResourceFormatSaver(format_saver gdclass.ResourceFormatSaver) {
	once.Do(singleton)
	class(self).AddResourceFormatSaver(format_saver, false)
}

/*
Unregisters the given [ResourceFormatSaver].
*/
func RemoveResourceFormatSaver(format_saver gdclass.ResourceFormatSaver) {
	once.Do(singleton)
	class(self).RemoveResourceFormatSaver(format_saver)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.ResourceSaver
func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
//go:nosplit
func (self class) Save(resource gdclass.Resource, path gd.String, flags classdb.ResourceSaverSaverFlags) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(resource[0])[0])
	callframe.Arg(frame, discreet.Get(path))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of extensions available for saving a resource of a given type.
*/
//go:nosplit
func (self class) GetRecognizedExtensions(atype gdclass.Resource) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(atype[0])[0])
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_get_recognized_extensions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
//go:nosplit
func (self class) AddResourceFormatSaver(format_saver gdclass.ResourceFormatSaver, at_front bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(format_saver[0])[0])
	callframe.Arg(frame, at_front)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_add_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the given [ResourceFormatSaver].
*/
//go:nosplit
func (self class) RemoveResourceFormatSaver(format_saver gdclass.ResourceFormatSaver)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(format_saver[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_remove_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ResourceSaver", func(ptr gd.Object) any { return classdb.ResourceSaver(ptr) })}
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
