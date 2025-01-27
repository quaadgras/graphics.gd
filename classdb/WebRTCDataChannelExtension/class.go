// Package WebRTCDataChannelExtension provides methods for working with WebRTCDataChannelExtension object instances.
package WebRTCDataChannelExtension

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/WebRTCDataChannel"
import "graphics.gd/classdb/PacketPeer"

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
var _ Path.ToNode

type Instance [1]gdclass.WebRTCDataChannelExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCDataChannelExtension() Instance
}
type Interface interface {
	GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) error
	PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) error
	GetAvailablePacketCount() int
	GetMaxPacketSize() int
	Poll() error
	Close()
	SetWriteMode(p_write_mode gdclass.WebRTCDataChannelWriteMode)
	GetWriteMode() gdclass.WebRTCDataChannelWriteMode
	WasStringPacket() bool
	GetReadyState() gdclass.WebRTCDataChannelChannelState
	GetLabel() string
	IsOrdered() bool
	GetId() int
	GetMaxPacketLifeTime() int
	GetMaxRetransmits() int
	GetProtocol() string
	IsNegotiated() bool
	GetBufferedAmount() int
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) (_ error) { return }
func (self implementation) PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) (_ error)    { return }
func (self implementation) GetAvailablePacketCount() (_ int)                                  { return }
func (self implementation) GetMaxPacketSize() (_ int)                                         { return }
func (self implementation) Poll() (_ error)                                                   { return }
func (self implementation) Close()                                                            { return }
func (self implementation) SetWriteMode(p_write_mode gdclass.WebRTCDataChannelWriteMode)      { return }
func (self implementation) GetWriteMode() (_ gdclass.WebRTCDataChannelWriteMode)              { return }
func (self implementation) WasStringPacket() (_ bool)                                         { return }
func (self implementation) GetReadyState() (_ gdclass.WebRTCDataChannelChannelState)          { return }
func (self implementation) GetLabel() (_ string)                                              { return }
func (self implementation) IsOrdered() (_ bool)                                               { return }
func (self implementation) GetId() (_ int)                                                    { return }
func (self implementation) GetMaxPacketLifeTime() (_ int)                                     { return }
func (self implementation) GetMaxRetransmits() (_ int)                                        { return }
func (self implementation) GetProtocol() (_ string)                                           { return }
func (self implementation) IsNegotiated() (_ bool)                                            { return }
func (self implementation) GetBufferedAmount() (_ int)                                        { return }
func (Instance) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_buffer_size = gd.UnsafeGet[*int32](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_buffer_size)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_buffer_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer, int(p_buffer_size))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_available_packet_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_max_packet_size(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _poll(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _set_write_mode(impl func(ptr unsafe.Pointer, p_write_mode gdclass.WebRTCDataChannelWriteMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_write_mode = gd.UnsafeGet[gdclass.WebRTCDataChannelWriteMode](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_write_mode)
	}
}
func (Instance) _get_write_mode(impl func(ptr unsafe.Pointer) gdclass.WebRTCDataChannelWriteMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _was_string_packet(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_ready_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCDataChannelChannelState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_label(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_ordered(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_max_packet_life_time(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_max_retransmits(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_protocol(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_negotiated(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_buffered_amount(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCDataChannelExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCDataChannelExtension"))
	casted := Instance{*(*gdclass.WebRTCDataChannelExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_buffer_size = gd.UnsafeGet[*int32](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_buffer_size)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size gd.Int) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_buffer_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer, p_buffer_size)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_available_packet_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_max_packet_size(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _poll(impl func(ptr unsafe.Pointer) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _set_write_mode(impl func(ptr unsafe.Pointer, p_write_mode gdclass.WebRTCDataChannelWriteMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_write_mode = gd.UnsafeGet[gdclass.WebRTCDataChannelWriteMode](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_write_mode)
	}
}

func (class) _get_write_mode(impl func(ptr unsafe.Pointer) gdclass.WebRTCDataChannelWriteMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _was_string_packet(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_ready_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCDataChannelChannelState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_label(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_ordered(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_max_packet_life_time(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_max_retransmits(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_protocol(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_negotiated(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_buffered_amount(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsWebRTCDataChannelExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCDataChannelExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWebRTCDataChannel() WebRTCDataChannel.Advanced {
	return *((*WebRTCDataChannel.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCDataChannel() WebRTCDataChannel.Instance {
	return *((*WebRTCDataChannel.Instance)(unsafe.Pointer(&self)))
}
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
	case "_get_packet":
		return reflect.ValueOf(self._get_packet)
	case "_put_packet":
		return reflect.ValueOf(self._put_packet)
	case "_get_available_packet_count":
		return reflect.ValueOf(self._get_available_packet_count)
	case "_get_max_packet_size":
		return reflect.ValueOf(self._get_max_packet_size)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	case "_set_write_mode":
		return reflect.ValueOf(self._set_write_mode)
	case "_get_write_mode":
		return reflect.ValueOf(self._get_write_mode)
	case "_was_string_packet":
		return reflect.ValueOf(self._was_string_packet)
	case "_get_ready_state":
		return reflect.ValueOf(self._get_ready_state)
	case "_get_label":
		return reflect.ValueOf(self._get_label)
	case "_is_ordered":
		return reflect.ValueOf(self._is_ordered)
	case "_get_id":
		return reflect.ValueOf(self._get_id)
	case "_get_max_packet_life_time":
		return reflect.ValueOf(self._get_max_packet_life_time)
	case "_get_max_retransmits":
		return reflect.ValueOf(self._get_max_retransmits)
	case "_get_protocol":
		return reflect.ValueOf(self._get_protocol)
	case "_is_negotiated":
		return reflect.ValueOf(self._is_negotiated)
	case "_get_buffered_amount":
		return reflect.ValueOf(self._get_buffered_amount)
	default:
		return gd.VirtualByName(WebRTCDataChannel.Advanced(self.AsWebRTCDataChannel()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_packet":
		return reflect.ValueOf(self._get_packet)
	case "_put_packet":
		return reflect.ValueOf(self._put_packet)
	case "_get_available_packet_count":
		return reflect.ValueOf(self._get_available_packet_count)
	case "_get_max_packet_size":
		return reflect.ValueOf(self._get_max_packet_size)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	case "_set_write_mode":
		return reflect.ValueOf(self._set_write_mode)
	case "_get_write_mode":
		return reflect.ValueOf(self._get_write_mode)
	case "_was_string_packet":
		return reflect.ValueOf(self._was_string_packet)
	case "_get_ready_state":
		return reflect.ValueOf(self._get_ready_state)
	case "_get_label":
		return reflect.ValueOf(self._get_label)
	case "_is_ordered":
		return reflect.ValueOf(self._is_ordered)
	case "_get_id":
		return reflect.ValueOf(self._get_id)
	case "_get_max_packet_life_time":
		return reflect.ValueOf(self._get_max_packet_life_time)
	case "_get_max_retransmits":
		return reflect.ValueOf(self._get_max_retransmits)
	case "_get_protocol":
		return reflect.ValueOf(self._get_protocol)
	case "_is_negotiated":
		return reflect.ValueOf(self._is_negotiated)
	case "_get_buffered_amount":
		return reflect.ValueOf(self._get_buffered_amount)
	default:
		return gd.VirtualByName(WebRTCDataChannel.Instance(self.AsWebRTCDataChannel()), name)
	}
}
func init() {
	gdclass.Register("WebRTCDataChannelExtension", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCDataChannelExtension{*(*gdclass.WebRTCDataChannelExtension)(unsafe.Pointer(&ptr))}
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
