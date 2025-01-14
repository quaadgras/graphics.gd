// Package AESContext provides methods for working with AESContext object instances.
package AESContext

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class holds the context information required for encryption and decryption operations with AES (Advanced Encryption Standard). Both AES-ECB and AES-CBC modes are supported.
[codeblocks]
[gdscript]
extends Node

var aes = AESContext.new()

func _ready():

	var key = "My secret key!!!" # Key must be either 16 or 32 bytes.
	var data = "My secret text!!" # Data size must be multiple of 16 bytes, apply padding if needed.
	# Encrypt ECB
	aes.start(AESContext.MODE_ECB_ENCRYPT, key.to_utf8_buffer())
	var encrypted = aes.update(data.to_utf8_buffer())
	aes.finish()
	# Decrypt ECB
	aes.start(AESContext.MODE_ECB_DECRYPT, key.to_utf8_buffer())
	var decrypted = aes.update(encrypted)
	aes.finish()
	# Check ECB
	assert(decrypted == data.to_utf8_buffer())

	var iv = "My secret iv!!!!" # IV must be of exactly 16 bytes.
	# Encrypt CBC
	aes.start(AESContext.MODE_CBC_ENCRYPT, key.to_utf8_buffer(), iv.to_utf8_buffer())
	encrypted = aes.update(data.to_utf8_buffer())
	aes.finish()
	# Decrypt CBC
	aes.start(AESContext.MODE_CBC_DECRYPT, key.to_utf8_buffer(), iv.to_utf8_buffer())
	decrypted = aes.update(encrypted)
	aes.finish()
	# Check CBC
	assert(decrypted == data.to_utf8_buffer())

[/gdscript]
[csharp]
using Godot;
using System.Diagnostics;

public partial class MyNode : Node

	{
	    private AesContext _aes = new AesContext();

	    public override void _Ready()
	    {
	        string key = "My secret key!!!"; // Key must be either 16 or 32 bytes.
	        string data = "My secret text!!"; // Data size must be multiple of 16 bytes, apply padding if needed.
	        // Encrypt ECB
	        _aes.Start(AesContext.Mode.EcbEncrypt, key.ToUtf8Buffer());
	        byte[] encrypted = _aes.Update(data.ToUtf8Buffer());
	        _aes.Finish();
	        // Decrypt ECB
	        _aes.Start(AesContext.Mode.EcbDecrypt, key.ToUtf8Buffer());
	        byte[] decrypted = _aes.Update(encrypted);
	        _aes.Finish();
	        // Check ECB
	        Debug.Assert(decrypted == data.ToUtf8Buffer());

	        string iv = "My secret iv!!!!"; // IV must be of exactly 16 bytes.
	        // Encrypt CBC
	        _aes.Start(AesContext.Mode.EcbEncrypt, key.ToUtf8Buffer(), iv.ToUtf8Buffer());
	        encrypted = _aes.Update(data.ToUtf8Buffer());
	        _aes.Finish();
	        // Decrypt CBC
	        _aes.Start(AesContext.Mode.EcbDecrypt, key.ToUtf8Buffer(), iv.ToUtf8Buffer());
	        decrypted = _aes.Update(encrypted);
	        _aes.Finish();
	        // Check CBC
	        Debug.Assert(decrypted == data.ToUtf8Buffer());
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.AESContext

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAESContext() Instance
}

/*
Start the AES context in the given [param mode]. A [param key] of either 16 or 32 bytes must always be provided, while an [param iv] (initialization vector) of exactly 16 bytes, is only needed when [param mode] is either [constant MODE_CBC_ENCRYPT] or [constant MODE_CBC_DECRYPT].
*/
func (self Instance) Start(mode gdclass.AESContextMode, key []byte) error {
	return error(class(self).Start(mode, gd.NewPackedByteSlice(key), gd.NewPackedByteSlice([1][]byte{}[0])))
}

/*
Run the desired operation for this AES context. Will return a [PackedByteArray] containing the result of encrypting (or decrypting) the given [param src]. See [method start] for mode of operation.
[b]Note:[/b] The size of [param src] must be a multiple of 16. Apply some padding if needed.
*/
func (self Instance) Update(src []byte) []byte {
	return []byte(class(self).Update(gd.NewPackedByteSlice(src)).Bytes())
}

/*
Get the current IV state for this context (IV gets updated when calling [method update]). You normally don't need this function.
[b]Note:[/b] This function only makes sense when the context is started with [constant MODE_CBC_ENCRYPT] or [constant MODE_CBC_DECRYPT].
*/
func (self Instance) GetIvState() []byte {
	return []byte(class(self).GetIvState().Bytes())
}

/*
Close this AES context so it can be started again. See [method start].
*/
func (self Instance) Finish() {
	class(self).Finish()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AESContext

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AESContext"))
	casted := Instance{*(*gdclass.AESContext)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Start the AES context in the given [param mode]. A [param key] of either 16 or 32 bytes must always be provided, while an [param iv] (initialization vector) of exactly 16 bytes, is only needed when [param mode] is either [constant MODE_CBC_ENCRYPT] or [constant MODE_CBC_DECRYPT].
*/
//go:nosplit
func (self class) Start(mode gdclass.AESContextMode, key gd.PackedByteArray, iv gd.PackedByteArray) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, pointers.Get(iv))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AESContext.Bind_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Run the desired operation for this AES context. Will return a [PackedByteArray] containing the result of encrypting (or decrypting) the given [param src]. See [method start] for mode of operation.
[b]Note:[/b] The size of [param src] must be a multiple of 16. Apply some padding if needed.
*/
//go:nosplit
func (self class) Update(src gd.PackedByteArray) gd.PackedByteArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(src))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AESContext.Bind_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Get the current IV state for this context (IV gets updated when calling [method update]). You normally don't need this function.
[b]Note:[/b] This function only makes sense when the context is started with [constant MODE_CBC_ENCRYPT] or [constant MODE_CBC_DECRYPT].
*/
//go:nosplit
func (self class) GetIvState() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AESContext.Bind_get_iv_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Close this AES context so it can be started again. See [method start].
*/
//go:nosplit
func (self class) Finish() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AESContext.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsAESContext() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAESContext() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AESContext", func(ptr gd.Object) any { return [1]gdclass.AESContext{*(*gdclass.AESContext)(unsafe.Pointer(&ptr))} })
}

type Mode = gdclass.AESContextMode

const (
	/*AES electronic codebook encryption mode.*/
	ModeEcbEncrypt Mode = 0
	/*AES electronic codebook decryption mode.*/
	ModeEcbDecrypt Mode = 1
	/*AES cipher blocker chaining encryption mode.*/
	ModeCbcEncrypt Mode = 2
	/*AES cipher blocker chaining decryption mode.*/
	ModeCbcDecrypt Mode = 3
	/*Maximum value for the mode enum.*/
	ModeMax Mode = 4
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
