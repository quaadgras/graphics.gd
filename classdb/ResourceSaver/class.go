// Package ResourceSaver provides methods for working with ResourceSaver object instances.
package ResourceSaver

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
A singleton for saving resource types to the filesystem.
It uses the many [ResourceFormatSaver] classes registered in the engine (either built-in or from a plugin) to save resource data to text-based (e.g. [code].tres[/code] or [code].tscn[/code]) or binary files (e.g. [code].res[/code] or [code].scn[/code]).
*/
var self [1]gdclass.ResourceSaver
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceSaver)
	self = *(*[1]gdclass.ResourceSaver)(unsafe.Pointer(&obj))
}

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
func Save(resource [1]gdclass.Resource) error {
	once.Do(singleton)
	return error(gd.ToError(class(self).Save(resource, gd.NewString(""), 0)))
}

/*
Returns the list of extensions available for saving a resource of a given type.
*/
func GetRecognizedExtensions(atype [1]gdclass.Resource) []string {
	once.Do(singleton)
	return []string(class(self).GetRecognizedExtensions(atype).Strings())
}

/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
func AddResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver) {
	once.Do(singleton)
	class(self).AddResourceFormatSaver(format_saver, false)
}

/*
Unregisters the given [ResourceFormatSaver].
*/
func RemoveResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver) {
	once.Do(singleton)
	class(self).RemoveResourceFormatSaver(format_saver)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ResourceSaver

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
//go:nosplit
func (self class) Save(resource [1]gdclass.Resource, path gd.String, flags gdclass.ResourceSaverSaverFlags) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_save, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of extensions available for saving a resource of a given type.
*/
//go:nosplit
func (self class) GetRecognizedExtensions(atype [1]gdclass.Resource) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype[0])[0])
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_get_recognized_extensions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
//go:nosplit
func (self class) AddResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver, at_front bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_saver[0])[0])
	callframe.Arg(frame, at_front)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_add_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the given [ResourceFormatSaver].
*/
//go:nosplit
func (self class) RemoveResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_saver[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_remove_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ResourceSaver", func(ptr gd.Object) any {
		return [1]gdclass.ResourceSaver{*(*gdclass.ResourceSaver)(unsafe.Pointer(&ptr))}
	})
}

type SaverFlags = gdclass.ResourceSaverSaverFlags

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

type Error = gd.Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
