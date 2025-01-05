// Package StreamPeerGZIP provides methods for working with StreamPeerGZIP object instances.
package StreamPeerGZIP

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/StreamPeer"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class allows to compress or decompress data using GZIP/deflate in a streaming fashion. This is particularly useful when compressing or decompressing files that have to be sent through the network without needing to allocate them all in memory.
After starting the stream via [method start_compression] (or [method start_decompression]), calling [method StreamPeer.put_partial_data] on this stream will compress (or decompress) the data, writing it to the internal buffer. Calling [method StreamPeer.get_available_bytes] will return the pending bytes in the internal buffer, and [method StreamPeer.get_partial_data] will retrieve the compressed (or decompressed) bytes from it. When the stream is over, you must call [method finish] to ensure the internal buffer is properly flushed (make sure to call [method StreamPeer.get_available_bytes] on last time to check if more data needs to be read after that).
*/
type Instance [1]gdclass.StreamPeerGZIP
type Any interface {
	gd.IsClass
	AsStreamPeerGZIP() Instance
}

/*
Start the stream in compression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
func (self Instance) StartCompression() error {
	return error(class(self).StartCompression(false, gd.Int(65535)))
}

/*
Start the stream in decompression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
func (self Instance) StartDecompression() error {
	return error(class(self).StartDecompression(false, gd.Int(65535)))
}

/*
Finalizes the stream, compressing or decompressing any buffered chunk left.
*/
func (self Instance) Finish() error {
	return error(class(self).Finish())
}

/*
Clears this stream, resetting the internal state.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StreamPeerGZIP

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerGZIP"))
	return Instance{*(*gdclass.StreamPeerGZIP)(unsafe.Pointer(&object))}
}

/*
Start the stream in compression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
//go:nosplit
func (self class) StartCompression(use_deflate bool, buffer_size gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, use_deflate)
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_start_compression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Start the stream in decompression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
//go:nosplit
func (self class) StartDecompression(use_deflate bool, buffer_size gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, use_deflate)
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_start_decompression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Finalizes the stream, compressing or decompressing any buffered chunk left.
*/
//go:nosplit
func (self class) Finish() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears this stream, resetting the internal state.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerGZIP.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsStreamPeerGZIP() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerGZIP() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StreamPeer.Advanced(self.AsStreamPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StreamPeer.Instance(self.AsStreamPeer()), name)
	}
}
func init() {
	gdclass.Register("StreamPeerGZIP", func(ptr gd.Object) any {
		return [1]gdclass.StreamPeerGZIP{*(*gdclass.StreamPeerGZIP)(unsafe.Pointer(&ptr))}
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
