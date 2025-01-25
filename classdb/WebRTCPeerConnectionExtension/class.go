// Package WebRTCPeerConnectionExtension provides methods for working with WebRTCPeerConnectionExtension object instances.
package WebRTCPeerConnectionExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/WebRTCPeerConnection"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.WebRTCPeerConnectionExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCPeerConnectionExtension() Instance
}
type Interface interface {
	GetConnectionState() gdclass.WebRTCPeerConnectionConnectionState
	GetGatheringState() gdclass.WebRTCPeerConnectionGatheringState
	GetSignalingState() gdclass.WebRTCPeerConnectionSignalingState
	Initialize(p_config map[any]any) error
	CreateDataChannel(p_label string, p_config map[any]any) [1]gdclass.WebRTCDataChannel
	CreateOffer() error
	SetRemoteDescription(p_type string, p_sdp string) error
	SetLocalDescription(p_type string, p_sdp string) error
	AddIceCandidate(p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) error
	Poll() error
	Close()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetConnectionState() (_ gdclass.WebRTCPeerConnectionConnectionState) {
	return
}
func (self implementation) GetGatheringState() (_ gdclass.WebRTCPeerConnectionGatheringState) { return }
func (self implementation) GetSignalingState() (_ gdclass.WebRTCPeerConnectionSignalingState) { return }
func (self implementation) Initialize(p_config map[any]any) (_ error)                         { return }
func (self implementation) CreateDataChannel(p_label string, p_config map[any]any) (_ [1]gdclass.WebRTCDataChannel) {
	return
}
func (self implementation) CreateOffer() (_ error)                                     { return }
func (self implementation) SetRemoteDescription(p_type string, p_sdp string) (_ error) { return }
func (self implementation) SetLocalDescription(p_type string, p_sdp string) (_ error)  { return }
func (self implementation) AddIceCandidate(p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) (_ error) {
	return
}
func (self implementation) Poll() (_ error) { return }
func (self implementation) Close()          { return }
func (Instance) _get_connection_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionConnectionState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_gathering_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionGatheringState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_signaling_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionSignalingState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _initialize(impl func(ptr unsafe.Pointer, p_config map[any]any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(p_config)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gd.DictionaryAs[any, any](p_config))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _create_data_channel(impl func(ptr unsafe.Pointer, p_label string, p_config map[any]any) [1]gdclass.WebRTCDataChannel) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_label = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(p_label)
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(p_config)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label.String(), gd.DictionaryAs[any, any](p_config))
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _create_offer(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _set_remote_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(p_type)
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(p_sdp)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _set_local_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(p_type)
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		defer pointers.End(p_sdp)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_sdp_mid_name = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(p_sdp_mid_name)
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var p_sdp_name = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(p_sdp_name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name.String(), int(p_sdp_mline_index), p_sdp_name.String())
		gd.UnsafeSet(p_back, ret)
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

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCPeerConnectionExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCPeerConnectionExtension"))
	casted := Instance{*(*gdclass.WebRTCPeerConnectionExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _get_connection_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionConnectionState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_gathering_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionGatheringState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_signaling_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionSignalingState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _initialize(impl func(ptr unsafe.Pointer, p_config gd.Dictionary) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _create_data_channel(impl func(ptr unsafe.Pointer, p_label gd.String, p_config gd.Dictionary) [1]gdclass.WebRTCDataChannel) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_label = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var p_config = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label, p_config)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _create_offer(impl func(ptr unsafe.Pointer) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_remote_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _set_local_description(impl func(ptr unsafe.Pointer, p_type gd.String, p_sdp gd.String) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var p_sdp = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name gd.String, p_sdp_mline_index gd.Int, p_sdp_name gd.String) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_sdp_mid_name = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		var p_sdp_mline_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var p_sdp_name = pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name, p_sdp_mline_index, p_sdp_name)
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

func (self class) AsWebRTCPeerConnectionExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCPeerConnectionExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWebRTCPeerConnection() WebRTCPeerConnection.Advanced {
	return *((*WebRTCPeerConnection.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCPeerConnection() WebRTCPeerConnection.Instance {
	return *((*WebRTCPeerConnection.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state":
		return reflect.ValueOf(self._get_connection_state)
	case "_get_gathering_state":
		return reflect.ValueOf(self._get_gathering_state)
	case "_get_signaling_state":
		return reflect.ValueOf(self._get_signaling_state)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_create_data_channel":
		return reflect.ValueOf(self._create_data_channel)
	case "_create_offer":
		return reflect.ValueOf(self._create_offer)
	case "_set_remote_description":
		return reflect.ValueOf(self._set_remote_description)
	case "_set_local_description":
		return reflect.ValueOf(self._set_local_description)
	case "_add_ice_candidate":
		return reflect.ValueOf(self._add_ice_candidate)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	default:
		return gd.VirtualByName(WebRTCPeerConnection.Advanced(self.AsWebRTCPeerConnection()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state":
		return reflect.ValueOf(self._get_connection_state)
	case "_get_gathering_state":
		return reflect.ValueOf(self._get_gathering_state)
	case "_get_signaling_state":
		return reflect.ValueOf(self._get_signaling_state)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_create_data_channel":
		return reflect.ValueOf(self._create_data_channel)
	case "_create_offer":
		return reflect.ValueOf(self._create_offer)
	case "_set_remote_description":
		return reflect.ValueOf(self._set_remote_description)
	case "_set_local_description":
		return reflect.ValueOf(self._set_local_description)
	case "_add_ice_candidate":
		return reflect.ValueOf(self._add_ice_candidate)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	default:
		return gd.VirtualByName(WebRTCPeerConnection.Instance(self.AsWebRTCPeerConnection()), name)
	}
}
func init() {
	gdclass.Register("WebRTCPeerConnectionExtension", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCPeerConnectionExtension{*(*gdclass.WebRTCPeerConnectionExtension)(unsafe.Pointer(&ptr))}
	})
}

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
