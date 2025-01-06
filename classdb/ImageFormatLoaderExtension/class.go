// Package ImageFormatLoaderExtension provides methods for working with ImageFormatLoaderExtension object instances.
package ImageFormatLoaderExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/ImageFormatLoader"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The engine supports multiple image formats out of the box (PNG, SVG, JPEG, WebP to name a few), but you can choose to implement support for additional image formats by extending this class.
Be sure to respect the documented return types and values. You should create an instance of it, and call [method add_format_loader] to register that loader during the initialization phase.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=ImageFormatLoaderExtension)
*/
type Instance [1]gdclass.ImageFormatLoaderExtension
type Any interface {
	gd.IsClass
	AsImageFormatLoaderExtension() Instance
}
type Interface interface {
	//Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
	GetRecognizedExtensions() []string
	//Loads the content of [param fileaccess] into the provided [param image].
	LoadImage(image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale Float.X) error
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) GetRecognizedExtensions() (_ []string) { return }
func (self Implementation) LoadImage(image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale Float.X) (_ error) {
	return
}

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Loads the content of [param fileaccess] into the provided [param image].
*/
func (Instance) _load_image(impl func(ptr unsafe.Pointer, image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale Float.X) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(image[0])
		var fileaccess = [1]gdclass.FileAccess{pointers.New[gdclass.FileAccess]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(fileaccess[0])
		var flags = gd.UnsafeGet[gdclass.ImageFormatLoaderLoaderFlags](p_args, 2)
		var scale = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, Float.X(scale))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
func (self Instance) AddFormatLoader() {
	class(self).AddFormatLoader()
}

/*
Remove this format loader from the engine.
*/
func (self Instance) RemoveFormatLoader() {
	class(self).RemoveFormatLoader()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImageFormatLoaderExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageFormatLoaderExtension"))
	casted := Instance{*(*gdclass.ImageFormatLoaderExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the list of file extensions for this image format. Files with the given extensions will be treated as image file and loaded using this class.
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Loads the content of [param fileaccess] into the provided [param image].
*/
func (class) _load_image(impl func(ptr unsafe.Pointer, image [1]gdclass.Image, fileaccess [1]gdclass.FileAccess, flags gdclass.ImageFormatLoaderLoaderFlags, scale gd.Float) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(image[0])
		var fileaccess = [1]gdclass.FileAccess{pointers.New[gdclass.FileAccess]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(fileaccess[0])
		var flags = gd.UnsafeGet[gdclass.ImageFormatLoaderLoaderFlags](p_args, 2)
		var scale = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, image, fileaccess, flags, scale)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Add this format loader to the engine, allowing it to recognize the file extensions returned by [method _get_recognized_extensions].
*/
//go:nosplit
func (self class) AddFormatLoader() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageFormatLoaderExtension.Bind_add_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Remove this format loader from the engine.
*/
//go:nosplit
func (self class) RemoveFormatLoader() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageFormatLoaderExtension.Bind_remove_format_loader, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageFormatLoaderExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageFormatLoaderExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsImageFormatLoader() ImageFormatLoader.Advanced {
	return *((*ImageFormatLoader.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageFormatLoader() ImageFormatLoader.Instance {
	return *((*ImageFormatLoader.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_load_image":
		return reflect.ValueOf(self._load_image)
	default:
		return gd.VirtualByName(ImageFormatLoader.Advanced(self.AsImageFormatLoader()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_load_image":
		return reflect.ValueOf(self._load_image)
	default:
		return gd.VirtualByName(ImageFormatLoader.Instance(self.AsImageFormatLoader()), name)
	}
}
func init() {
	gdclass.Register("ImageFormatLoaderExtension", func(ptr gd.Object) any {
		return [1]gdclass.ImageFormatLoaderExtension{*(*gdclass.ImageFormatLoaderExtension)(unsafe.Pointer(&ptr))}
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
