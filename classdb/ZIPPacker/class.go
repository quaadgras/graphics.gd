// Package ZIPPacker provides methods for working with ZIPPacker object instances.
package ZIPPacker

import "unsafe"
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

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This class implements a writer that allows storing the multiple blobs in a zip archive.
[codeblock]
func write_zip_file():

	var writer := ZIPPacker.new()
	var err := writer.open("user://archive.zip")
	if err != OK:
	    return err
	writer.start_file("hello.txt")
	writer.write_file("Hello World".to_utf8_buffer())
	writer.close_file()

	writer.close()
	return OK

[/codeblock]
*/
type Instance [1]gdclass.ZIPPacker

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsZIPPacker() Instance
}

/*
Opens a zip file for writing at the given path using the specified write mode.
This must be called before everything else.
*/
func (self Instance) Open(path string) error { //gd:ZIPPacker.open
	return error(gd.ToError(class(self).Open(gd.NewString(path), 0)))
}

/*
Starts writing to a file within the archive. Only one file can be written at the same time.
Must be called after [method open].
*/
func (self Instance) StartFile(path string) error { //gd:ZIPPacker.start_file
	return error(gd.ToError(class(self).StartFile(gd.NewString(path))))
}

/*
Write the given [param data] to the file.
Needs to be called after [method start_file].
*/
func (self Instance) WriteFile(data []byte) error { //gd:ZIPPacker.write_file
	return error(gd.ToError(class(self).WriteFile(gd.NewPackedByteSlice(data))))
}

/*
Stops writing to a file within the archive.
It will fail if there is no open file.
*/
func (self Instance) CloseFile() error { //gd:ZIPPacker.close_file
	return error(gd.ToError(class(self).CloseFile()))
}

/*
Closes the underlying resources used by this instance.
*/
func (self Instance) Close() error { //gd:ZIPPacker.close
	return error(gd.ToError(class(self).Close()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ZIPPacker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ZIPPacker"))
	casted := Instance{*(*gdclass.ZIPPacker)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Opens a zip file for writing at the given path using the specified write mode.
This must be called before everything else.
*/
//go:nosplit
func (self class) Open(path gd.String, append gdclass.ZIPPackerZipAppend) gd.Error { //gd:ZIPPacker.open
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, append)
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts writing to a file within the archive. Only one file can be written at the same time.
Must be called after [method open].
*/
//go:nosplit
func (self class) StartFile(path gd.String) gd.Error { //gd:ZIPPacker.start_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_start_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Write the given [param data] to the file.
Needs to be called after [method start_file].
*/
//go:nosplit
func (self class) WriteFile(data gd.PackedByteArray) gd.Error { //gd:ZIPPacker.write_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_write_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Stops writing to a file within the archive.
It will fail if there is no open file.
*/
//go:nosplit
func (self class) CloseFile() gd.Error { //gd:ZIPPacker.close_file
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_close_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Closes the underlying resources used by this instance.
*/
//go:nosplit
func (self class) Close() gd.Error { //gd:ZIPPacker.close
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPPacker.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsZIPPacker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsZIPPacker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("ZIPPacker", func(ptr gd.Object) any { return [1]gdclass.ZIPPacker{*(*gdclass.ZIPPacker)(unsafe.Pointer(&ptr))} })
}

type ZipAppend = gdclass.ZIPPackerZipAppend //gd:ZIPPacker.ZipAppend

const (
	/*Create a new zip archive at the given path.*/
	AppendCreate ZipAppend = 0
	/*Append a new zip archive to the end of the already existing file at the given path.*/
	AppendCreateafter ZipAppend = 1
	/*Add new files to the existing zip archive at the given path.*/
	AppendAddinzip ZipAppend = 2
)

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
