package StreamPeerExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/StreamPeer"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.StreamPeerExtension
type Any interface {
	gd.IsClass
	AsStreamPeerExtension() Instance
}

func (Instance) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int, r_received *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, int(r_bytes), r_received)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int, r_received *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, int(r_bytes), r_received)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int, r_sent *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, int(p_bytes), r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int, r_sent *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, int(p_bytes), r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_available_bytes(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StreamPeerExtension

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerExtension"))
	return Instance{*(*gdclass.StreamPeerExtension)(unsafe.Pointer(&object))}
}

func (class) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes gd.Int, r_received *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes gd.Int, r_received *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_received = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes gd.Int, r_sent *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes gd.Int, r_sent *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_bytes = gd.UnsafeGet[gd.Int](p_args, 1)
		var r_sent = gd.UnsafeGet[*int32](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_available_bytes(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsStreamPeerExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	case "_get_partial_data":
		return reflect.ValueOf(self._get_partial_data)
	case "_put_data":
		return reflect.ValueOf(self._put_data)
	case "_put_partial_data":
		return reflect.ValueOf(self._put_partial_data)
	case "_get_available_bytes":
		return reflect.ValueOf(self._get_available_bytes)
	default:
		return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	case "_get_partial_data":
		return reflect.ValueOf(self._get_partial_data)
	case "_put_data":
		return reflect.ValueOf(self._put_data)
	case "_put_partial_data":
		return reflect.ValueOf(self._put_partial_data)
	case "_get_available_bytes":
		return reflect.ValueOf(self._get_available_bytes)
	default:
		return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}
func init() {
	gdclass.Register("StreamPeerExtension", func(ptr gd.Object) any {
		return [1]gdclass.StreamPeerExtension{*(*gdclass.StreamPeerExtension)(unsafe.Pointer(&ptr))}
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
