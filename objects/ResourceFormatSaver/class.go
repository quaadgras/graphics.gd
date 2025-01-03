package ResourceFormatSaver

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
The engine can save resources when you do it from the editor, or when you use the [ResourceSaver] singleton. This is accomplished thanks to multiple [ResourceFormatSaver]s, each handling its own format and called automatically by the engine.
By default, Godot saves resources as [code].tres[/code] (text-based), [code].res[/code] (binary) or another built-in format, but you can choose to create your own format by extending this class. Be sure to respect the documented return types and values. You should give it a global class name with [code]class_name[/code] for it to be registered. Like built-in ResourceFormatSavers, it will be called automatically when saving resources of its recognized type(s). You may also implement a [ResourceFormatLoader].

	// ResourceFormatSaver methods that can be overridden by a [Class] that extends it.
	type ResourceFormatSaver interface {
		//Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
		//Returns [constant OK] on success, or an [enum Error] constant in case of failure.
		Save(resource objects.Resource, path string, flags int) error
		//Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
		SetUid(path string, uid int) error
		//Returns whether the given resource object can be saved by this saver.
		Recognize(resource objects.Resource) bool
		//Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
		GetRecognizedExtensions(resource objects.Resource) []string
		//Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
		//If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
		RecognizePath(resource objects.Resource, path string) bool
	}
*/
type Instance [1]classdb.ResourceFormatSaver
type Any interface {
	gd.IsClass
	AsResourceFormatSaver() Instance
}

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _save(impl func(ptr unsafe.Pointer, resource objects.Resource, path string, flags int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(path)
		var flags = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String(), int(flags))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _set_uid(impl func(ptr unsafe.Pointer, path string, uid int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var uid = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(uid))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (Instance) _recognize(impl func(ptr unsafe.Pointer, resource objects.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource objects.Resource) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
*/
func (Instance) _recognize_path(impl func(ptr unsafe.Pointer, resource objects.Resource, path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(path)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ResourceFormatSaver

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceFormatSaver"))
	return Instance{*(*classdb.ResourceFormatSaver)(unsafe.Pointer(&object))}
}

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _save(impl func(ptr unsafe.Pointer, resource objects.Resource, path gd.String, flags gd.Int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var flags = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path, flags)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _set_uid(impl func(ptr unsafe.Pointer, path gd.String, uid gd.Int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var uid = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, uid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (class) _recognize(impl func(ptr unsafe.Pointer, resource objects.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource objects.Resource) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
*/
func (class) _recognize_path(impl func(ptr unsafe.Pointer, resource objects.Resource, path gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsResourceFormatSaver() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceFormatSaver() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted        { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted     { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_save":
		return reflect.ValueOf(self._save)
	case "_set_uid":
		return reflect.ValueOf(self._set_uid)
	case "_recognize":
		return reflect.ValueOf(self._recognize)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_save":
		return reflect.ValueOf(self._save)
	case "_set_uid":
		return reflect.ValueOf(self._set_uid)
	case "_recognize":
		return reflect.ValueOf(self._recognize)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("ResourceFormatSaver", func(ptr gd.Object) any {
		return [1]classdb.ResourceFormatSaver{*(*classdb.ResourceFormatSaver)(unsafe.Pointer(&ptr))}
	})
}

type Error int

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
