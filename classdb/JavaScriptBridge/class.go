// Package JavaScriptBridge provides methods for working with JavaScriptBridge object instances.
package JavaScriptBridge

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
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable

/*
The JavaScriptBridge singleton is implemented only in the Web export. It's used to access the browser's JavaScript context. This allows interaction with embedding pages or calling third-party JavaScript APIs.
[b]Note:[/b] This singleton can be disabled at build-time to improve security. By default, the JavaScriptBridge singleton is enabled. Official export templates also have the JavaScriptBridge singleton enabled. See [url=$DOCS_URL/contributing/development/compiling/compiling_for_web.html]Compiling for the Web[/url] in the documentation for more information.
*/
var self [1]gdclass.JavaScriptBridge
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.JavaScriptBridge)
	self = *(*[1]gdclass.JavaScriptBridge)(unsafe.Pointer(&obj))
}

/*
Execute the string [param code] as JavaScript code within the browser window. This is a call to the actual global JavaScript function [code skip-lint]eval()[/code].
If [param use_global_execution_context] is [code]true[/code], the code will be evaluated in the global execution context. Otherwise, it is evaluated in the execution context of a function within the engine's runtime environment.
*/
func Eval(code string) any { //gd:JavaScriptBridge.eval
	once.Do(singleton)
	return any(class(self).Eval(String.New(code), false).Interface())
}

/*
Returns an interface to a JavaScript object that can be used by scripts. The [param interface] must be a valid property of the JavaScript [code]window[/code]. The callback must accept a single [Array] argument, which will contain the JavaScript [code]arguments[/code]. See [JavaScriptObject] for usage.
*/
func GetInterface(intf string) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.get_interface
	once.Do(singleton)
	return [1]gdclass.JavaScriptObject(class(self).GetInterface(String.New(intf)))
}

/*
Creates a reference to a [Callable] that can be used as a callback by JavaScript. The reference must be kept until the callback happens, or it won't be called at all. See [JavaScriptObject] for usage.
*/
func CreateCallback(callable Callable.Function) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.create_callback
	once.Do(singleton)
	return [1]gdclass.JavaScriptObject(class(self).CreateCallback(Callable.New(callable)))
}

/*
Prompts the user to download a file containing the specified [param buffer]. The file will have the given [param name] and [param mime] type.
[b]Note:[/b] The browser may override the [url=https://en.wikipedia.org/wiki/Media_type]MIME type[/url] provided based on the file [param name]'s extension.
[b]Note:[/b] Browsers might block the download if [method download_buffer] is not being called from a user interaction (e.g. button click).
[b]Note:[/b] Browsers might ask the user for permission or block the download if multiple download requests are made in a quick succession.
*/
func DownloadBuffer(buffer []byte, name string) { //gd:JavaScriptBridge.download_buffer
	once.Do(singleton)
	class(self).DownloadBuffer(gd.NewPackedByteSlice(buffer), String.New(name), String.New("application/octet-stream"))
}

/*
Returns [code]true[/code] if a new version of the progressive web app is waiting to be activated.
[b]Note:[/b] Only relevant when exported as a Progressive Web App.
*/
func PwaNeedsUpdate() bool { //gd:JavaScriptBridge.pwa_needs_update
	once.Do(singleton)
	return bool(class(self).PwaNeedsUpdate())
}

/*
Performs the live update of the progressive web app. Forcing the new version to be installed and the page to be reloaded.
[b]Note:[/b] Your application will be [b]reloaded in all browser tabs[/b].
[b]Note:[/b] Only relevant when exported as a Progressive Web App and [method pwa_needs_update] returns [code]true[/code].
*/
func PwaUpdate() error { //gd:JavaScriptBridge.pwa_update
	once.Do(singleton)
	return error(gd.ToError(class(self).PwaUpdate()))
}

/*
Force synchronization of the persistent file system (when enabled).
[b]Note:[/b] This is only useful for modules or extensions that can't use [FileAccess] to write files.
*/
func ForceFsSync() { //gd:JavaScriptBridge.force_fs_sync
	once.Do(singleton)
	class(self).ForceFsSync()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.JavaScriptBridge

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Execute the string [param code] as JavaScript code within the browser window. This is a call to the actual global JavaScript function [code skip-lint]eval()[/code].
If [param use_global_execution_context] is [code]true[/code], the code will be evaluated in the global execution context. Otherwise, it is evaluated in the execution context of a function within the engine's runtime environment.
*/
//go:nosplit
func (self class) Eval(code String.Readable, use_global_execution_context bool) gd.Variant { //gd:JavaScriptBridge.eval
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(code)))
	callframe.Arg(frame, use_global_execution_context)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_eval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an interface to a JavaScript object that can be used by scripts. The [param interface] must be a valid property of the JavaScript [code]window[/code]. The callback must accept a single [Array] argument, which will contain the JavaScript [code]arguments[/code]. See [JavaScriptObject] for usage.
*/
//go:nosplit
func (self class) GetInterface(intf String.Readable) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.get_interface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(intf)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaScriptObject{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaScriptObject](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a reference to a [Callable] that can be used as a callback by JavaScript. The reference must be kept until the callback happens, or it won't be called at all. See [JavaScriptObject] for usage.
*/
//go:nosplit
func (self class) CreateCallback(callable Callable.Function) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.create_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_create_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaScriptObject{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaScriptObject](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Prompts the user to download a file containing the specified [param buffer]. The file will have the given [param name] and [param mime] type.
[b]Note:[/b] The browser may override the [url=https://en.wikipedia.org/wiki/Media_type]MIME type[/url] provided based on the file [param name]'s extension.
[b]Note:[/b] Browsers might block the download if [method download_buffer] is not being called from a user interaction (e.g. button click).
[b]Note:[/b] Browsers might ask the user for permission or block the download if multiple download requests are made in a quick succession.
*/
//go:nosplit
func (self class) DownloadBuffer(buffer gd.PackedByteArray, name String.Readable, mime String.Readable) { //gd:JavaScriptBridge.download_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(buffer))
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(mime)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_download_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a new version of the progressive web app is waiting to be activated.
[b]Note:[/b] Only relevant when exported as a Progressive Web App.
*/
//go:nosplit
func (self class) PwaNeedsUpdate() bool { //gd:JavaScriptBridge.pwa_needs_update
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_pwa_needs_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Performs the live update of the progressive web app. Forcing the new version to be installed and the page to be reloaded.
[b]Note:[/b] Your application will be [b]reloaded in all browser tabs[/b].
[b]Note:[/b] Only relevant when exported as a Progressive Web App and [method pwa_needs_update] returns [code]true[/code].
*/
//go:nosplit
func (self class) PwaUpdate() gd.Error { //gd:JavaScriptBridge.pwa_update
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_pwa_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Force synchronization of the persistent file system (when enabled).
[b]Note:[/b] This is only useful for modules or extensions that can't use [FileAccess] to write files.
*/
//go:nosplit
func (self class) ForceFsSync() { //gd:JavaScriptBridge.force_fs_sync
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_force_fs_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func OnPwaUpdateAvailable(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pwa_update_available"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("JavaScriptBridge", func(ptr gd.Object) any {
		return [1]gdclass.JavaScriptBridge{*(*gdclass.JavaScriptBridge)(unsafe.Pointer(&ptr))}
	})
}

type Error = gd.Error //gd:Error

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
