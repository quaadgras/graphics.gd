// Package WebRTCDataChannel provides methods for working with WebRTCDataChannel object instances.
package WebRTCDataChannel

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/PacketPeer"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

type Instance [1]gdclass.WebRTCDataChannel

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCDataChannel() Instance
}

/*
Reserved, but not used for now.
*/
func (self Instance) Poll() error {
	return error(class(self).Poll())
}

/*
Closes this data channel, notifying the other peer.
*/
func (self Instance) Close() {
	class(self).Close()
}

/*
Returns [code]true[/code] if the last received packet was transferred as text. See [member write_mode].
*/
func (self Instance) WasStringPacket() bool {
	return bool(class(self).WasStringPacket())
}

/*
Returns the current state of this channel, see [enum ChannelState].
*/
func (self Instance) GetReadyState() gdclass.WebRTCDataChannelChannelState {
	return gdclass.WebRTCDataChannelChannelState(class(self).GetReadyState())
}

/*
Returns the label assigned to this channel during creation.
*/
func (self Instance) GetLabel() string {
	return string(class(self).GetLabel().String())
}

/*
Returns [code]true[/code] if this channel was created with ordering enabled (default).
*/
func (self Instance) IsOrdered() bool {
	return bool(class(self).IsOrdered())
}

/*
Returns the ID assigned to this channel during creation (or auto-assigned during negotiation).
If the channel is not negotiated out-of-band the ID will only be available after the connection is established (will return [code]65535[/code] until then).
*/
func (self Instance) GetId() int {
	return int(int(class(self).GetId()))
}

/*
Returns the [code]maxPacketLifeTime[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
func (self Instance) GetMaxPacketLifeTime() int {
	return int(int(class(self).GetMaxPacketLifeTime()))
}

/*
Returns the [code]maxRetransmits[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
func (self Instance) GetMaxRetransmits() int {
	return int(int(class(self).GetMaxRetransmits()))
}

/*
Returns the sub-protocol assigned to this channel during creation. An empty string if not specified.
*/
func (self Instance) GetProtocol() string {
	return string(class(self).GetProtocol().String())
}

/*
Returns [code]true[/code] if this channel was created with out-of-band configuration.
*/
func (self Instance) IsNegotiated() bool {
	return bool(class(self).IsNegotiated())
}

/*
Returns the number of bytes currently queued to be sent over this channel.
*/
func (self Instance) GetBufferedAmount() int {
	return int(int(class(self).GetBufferedAmount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCDataChannel

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCDataChannel"))
	casted := Instance{*(*gdclass.WebRTCDataChannel)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) WriteMode() gdclass.WebRTCDataChannelWriteMode {
	return gdclass.WebRTCDataChannelWriteMode(class(self).GetWriteMode())
}

func (self Instance) SetWriteMode(value gdclass.WebRTCDataChannelWriteMode) {
	class(self).SetWriteMode(value)
}

/*
Reserved, but not used for now.
*/
//go:nosplit
func (self class) Poll() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Closes this data channel, notifying the other peer.
*/
//go:nosplit
func (self class) Close() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the last received packet was transferred as text. See [member write_mode].
*/
//go:nosplit
func (self class) WasStringPacket() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_was_string_packet, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWriteMode(write_mode gdclass.WebRTCDataChannelWriteMode) {
	var frame = callframe.New()
	callframe.Arg(frame, write_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_set_write_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWriteMode() gdclass.WebRTCDataChannelWriteMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WebRTCDataChannelWriteMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_write_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current state of this channel, see [enum ChannelState].
*/
//go:nosplit
func (self class) GetReadyState() gdclass.WebRTCDataChannelChannelState {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.WebRTCDataChannelChannelState](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_ready_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the label assigned to this channel during creation.
*/
//go:nosplit
func (self class) GetLabel() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this channel was created with ordering enabled (default).
*/
//go:nosplit
func (self class) IsOrdered() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_is_ordered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ID assigned to this channel during creation (or auto-assigned during negotiation).
If the channel is not negotiated out-of-band the ID will only be available after the connection is established (will return [code]65535[/code] until then).
*/
//go:nosplit
func (self class) GetId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [code]maxPacketLifeTime[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
//go:nosplit
func (self class) GetMaxPacketLifeTime() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_max_packet_life_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [code]maxRetransmits[/code] value assigned to this channel during creation.
Will be [code]65535[/code] if not specified.
*/
//go:nosplit
func (self class) GetMaxRetransmits() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_max_retransmits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the sub-protocol assigned to this channel during creation. An empty string if not specified.
*/
//go:nosplit
func (self class) GetProtocol() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_protocol, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this channel was created with out-of-band configuration.
*/
//go:nosplit
func (self class) IsNegotiated() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_is_negotiated, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of bytes currently queued to be sent over this channel.
*/
//go:nosplit
func (self class) GetBufferedAmount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCDataChannel.Bind_get_buffered_amount, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWebRTCDataChannel() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebRTCDataChannel() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(PacketPeer.Advanced(self.AsPacketPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PacketPeer.Instance(self.AsPacketPeer()), name)
	}
}
func init() {
	gdclass.Register("WebRTCDataChannel", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCDataChannel{*(*gdclass.WebRTCDataChannel)(unsafe.Pointer(&ptr))}
	})
}

type WriteMode = gdclass.WebRTCDataChannelWriteMode

const (
	/*Tells the channel to send data over this channel as text. An external peer (non-Godot) would receive this as a string.*/
	WriteModeText WriteMode = 0
	/*Tells the channel to send data over this channel as binary. An external peer (non-Godot) would receive this as array buffer or blob.*/
	WriteModeBinary WriteMode = 1
)

type ChannelState = gdclass.WebRTCDataChannelChannelState

const (
	/*The channel was created, but it's still trying to connect.*/
	StateConnecting ChannelState = 0
	/*The channel is currently open, and data can flow over it.*/
	StateOpen ChannelState = 1
	/*The channel is being closed, no new messages will be accepted, but those already in queue will be flushed.*/
	StateClosing ChannelState = 2
	/*The channel was closed, or connection failed.*/
	StateClosed ChannelState = 3
)

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
