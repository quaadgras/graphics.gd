// Package MultiplayerPeerExtension provides methods for working with MultiplayerPeerExtension object instances.
package MultiplayerPeerExtension

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
import "graphics.gd/classdb/MultiplayerPeer"
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

/*
This class is designed to be inherited from a GDExtension plugin to implement custom networking layers for the multiplayer API (such as WebRTC). All the methods below [b]must[/b] be implemented to have a working custom multiplayer implementation. See also [MultiplayerAPI].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=MultiplayerPeerExtension)
*/
type Instance [1]gdclass.MultiplayerPeerExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiplayerPeerExtension() Instance
}
type Interface interface {
	//Called when a packet needs to be received by the [MultiplayerAPI], with [param r_buffer_size] being the size of the binary [param r_buffer] in bytes.
	GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) error
	//Called when a packet needs to be sent by the [MultiplayerAPI], with [param p_buffer_size] being the size of the binary [param p_buffer] in bytes.
	PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) error
	//Called when the available packet count is internally requested by the [MultiplayerAPI].
	GetAvailablePacketCount() int
	//Called when the maximum allowed packet size (in bytes) is requested by the [MultiplayerAPI].
	GetMaxPacketSize() int
	//Called when a packet needs to be received by the [MultiplayerAPI], if [method _get_packet] isn't implemented. Use this when extending this class via GDScript.
	GetPacketScript() []byte
	//Called when a packet needs to be sent by the [MultiplayerAPI], if [method _put_packet] isn't implemented. Use this when extending this class via GDScript.
	PutPacketScript(p_buffer []byte) error
	//Called to get the channel over which the next available packet was received. See [method MultiplayerPeer.get_packet_channel].
	GetPacketChannel() int
	//Called to get the transfer mode the remote peer used to send the next available packet. See [method MultiplayerPeer.get_packet_mode].
	GetPacketMode() gdclass.MultiplayerPeerTransferMode
	//Called when the channel to use is set for this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
	SetTransferChannel(p_channel int)
	//Called when the transfer channel to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
	GetTransferChannel() int
	//Called when the transfer mode is set on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
	SetTransferMode(p_mode gdclass.MultiplayerPeerTransferMode)
	//Called when the transfer mode to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
	GetTransferMode() gdclass.MultiplayerPeerTransferMode
	//Called when the target peer to use is set for this [MultiplayerPeer] (see [method MultiplayerPeer.set_target_peer]).
	SetTargetPeer(p_peer int)
	//Called when the ID of the [MultiplayerPeer] who sent the most recent packet is requested (see [method MultiplayerPeer.get_packet_peer]).
	GetPacketPeer() int
	//Called when the "is server" status is requested on the [MultiplayerAPI]. See [method MultiplayerAPI.is_server].
	IsServer() bool
	//Called when the [MultiplayerAPI] is polled. See [method MultiplayerAPI.poll].
	Poll()
	//Called when the multiplayer peer should be immediately closed (see [method MultiplayerPeer.close]).
	Close()
	//Called when the connected [param p_peer] should be forcibly disconnected (see [method MultiplayerPeer.disconnect_peer]).
	DisconnectPeer(p_peer int, p_force bool)
	//Called when the unique ID of this [MultiplayerPeer] is requested (see [method MultiplayerPeer.get_unique_id]). The value must be between [code]1[/code] and [code]2147483647[/code].
	GetUniqueId() int
	//Called when the "refuse new connections" status is set on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
	SetRefuseNewConnections(p_enable bool)
	//Called when the "refuse new connections" status is requested on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
	IsRefusingNewConnections() bool
	//Called to check if the server can act as a relay in the current configuration. See [method MultiplayerPeer.is_server_relay_supported].
	IsServerRelaySupported() bool
	//Called when the connection status is requested on the [MultiplayerPeer] (see [method MultiplayerPeer.get_connection_status]).
	GetConnectionStatus() gdclass.MultiplayerPeerConnectionStatus
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) (_ error) { return }
func (self implementation) PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) (_ error)    { return }
func (self implementation) GetAvailablePacketCount() (_ int)                                  { return }
func (self implementation) GetMaxPacketSize() (_ int)                                         { return }
func (self implementation) GetPacketScript() (_ []byte)                                       { return }
func (self implementation) PutPacketScript(p_buffer []byte) (_ error)                         { return }
func (self implementation) GetPacketChannel() (_ int)                                         { return }
func (self implementation) GetPacketMode() (_ gdclass.MultiplayerPeerTransferMode)            { return }
func (self implementation) SetTransferChannel(p_channel int)                                  { return }
func (self implementation) GetTransferChannel() (_ int)                                       { return }
func (self implementation) SetTransferMode(p_mode gdclass.MultiplayerPeerTransferMode)        { return }
func (self implementation) GetTransferMode() (_ gdclass.MultiplayerPeerTransferMode)          { return }
func (self implementation) SetTargetPeer(p_peer int)                                          { return }
func (self implementation) GetPacketPeer() (_ int)                                            { return }
func (self implementation) IsServer() (_ bool)                                                { return }
func (self implementation) Poll()                                                             { return }
func (self implementation) Close()                                                            { return }
func (self implementation) DisconnectPeer(p_peer int, p_force bool)                           { return }
func (self implementation) GetUniqueId() (_ int)                                              { return }
func (self implementation) SetRefuseNewConnections(p_enable bool)                             { return }
func (self implementation) IsRefusingNewConnections() (_ bool)                                { return }
func (self implementation) IsServerRelaySupported() (_ bool)                                  { return }
func (self implementation) GetConnectionStatus() (_ gdclass.MultiplayerPeerConnectionStatus)  { return }

/*
Called when a packet needs to be received by the [MultiplayerAPI], with [param r_buffer_size] being the size of the binary [param r_buffer] in bytes.
*/
func (Instance) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_buffer_size = gd.UnsafeGet[*int32](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_buffer_size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when a packet needs to be sent by the [MultiplayerAPI], with [param p_buffer_size] being the size of the binary [param p_buffer] in bytes.
*/
func (Instance) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_buffer_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer, int(p_buffer_size))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the available packet count is internally requested by the [MultiplayerAPI].
*/
func (Instance) _get_available_packet_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the maximum allowed packet size (in bytes) is requested by the [MultiplayerAPI].
*/
func (Instance) _get_max_packet_size(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when a packet needs to be received by the [MultiplayerAPI], if [method _get_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (Instance) _get_packet_script(impl func(ptr unsafe.Pointer) []byte) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedByteSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when a packet needs to be sent by the [MultiplayerAPI], if [method _put_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (Instance) _put_packet_script(impl func(ptr unsafe.Pointer, p_buffer []byte) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = pointers.New[gd.PackedByteArray](gd.UnsafeGet[gd.PackedPointers](p_args, 0))
		defer pointers.End(p_buffer)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer.Bytes())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to get the channel over which the next available packet was received. See [method MultiplayerPeer.get_packet_channel].
*/
func (Instance) _get_packet_channel(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called to get the transfer mode the remote peer used to send the next available packet. See [method MultiplayerPeer.get_packet_mode].
*/
func (Instance) _get_packet_mode(impl func(ptr unsafe.Pointer) gdclass.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the channel to use is set for this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (Instance) _set_transfer_channel(impl func(ptr unsafe.Pointer, p_channel int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_channel = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(p_channel))
	}
}

/*
Called when the transfer channel to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (Instance) _get_transfer_channel(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the transfer mode is set on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (Instance) _set_transfer_mode(impl func(ptr unsafe.Pointer, p_mode gdclass.MultiplayerPeerTransferMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_mode = gd.UnsafeGet[gdclass.MultiplayerPeerTransferMode](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_mode)
	}
}

/*
Called when the transfer mode to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (Instance) _get_transfer_mode(impl func(ptr unsafe.Pointer) gdclass.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the target peer to use is set for this [MultiplayerPeer] (see [method MultiplayerPeer.set_target_peer]).
*/
func (Instance) _set_target_peer(impl func(ptr unsafe.Pointer, p_peer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_peer = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(p_peer))
	}
}

/*
Called when the ID of the [MultiplayerPeer] who sent the most recent packet is requested (see [method MultiplayerPeer.get_packet_peer]).
*/
func (Instance) _get_packet_peer(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the "is server" status is requested on the [MultiplayerAPI]. See [method MultiplayerAPI.is_server].
*/
func (Instance) _is_server(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [MultiplayerAPI] is polled. See [method MultiplayerAPI.poll].
*/
func (Instance) _poll(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the multiplayer peer should be immediately closed (see [method MultiplayerPeer.close]).
*/
func (Instance) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the connected [param p_peer] should be forcibly disconnected (see [method MultiplayerPeer.disconnect_peer]).
*/
func (Instance) _disconnect_peer(impl func(ptr unsafe.Pointer, p_peer int, p_force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_peer = gd.UnsafeGet[gd.Int](p_args, 0)

		var p_force = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(p_peer), p_force)
	}
}

/*
Called when the unique ID of this [MultiplayerPeer] is requested (see [method MultiplayerPeer.get_unique_id]). The value must be between [code]1[/code] and [code]2147483647[/code].
*/
func (Instance) _get_unique_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the "refuse new connections" status is set on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (Instance) _set_refuse_new_connections(impl func(ptr unsafe.Pointer, p_enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_enable = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_enable)
	}
}

/*
Called when the "refuse new connections" status is requested on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (Instance) _is_refusing_new_connections(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to check if the server can act as a relay in the current configuration. See [method MultiplayerPeer.is_server_relay_supported].
*/
func (Instance) _is_server_relay_supported(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the connection status is requested on the [MultiplayerPeer] (see [method MultiplayerPeer.get_connection_status]).
*/
func (Instance) _get_connection_status(impl func(ptr unsafe.Pointer) gdclass.MultiplayerPeerConnectionStatus) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerPeerExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerPeerExtension"))
	casted := Instance{*(*gdclass.MultiplayerPeerExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Called when a packet needs to be received by the [MultiplayerAPI], with [param r_buffer_size] being the size of the binary [param r_buffer] in bytes.
*/
func (class) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_buffer_size = gd.UnsafeGet[*int32](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_buffer_size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when a packet needs to be sent by the [MultiplayerAPI], with [param p_buffer_size] being the size of the binary [param p_buffer] in bytes.
*/
func (class) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size gd.Int) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_buffer_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer, p_buffer_size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the available packet count is internally requested by the [MultiplayerAPI].
*/
func (class) _get_available_packet_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the maximum allowed packet size (in bytes) is requested by the [MultiplayerAPI].
*/
func (class) _get_max_packet_size(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when a packet needs to be received by the [MultiplayerAPI], if [method _get_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (class) _get_packet_script(impl func(ptr unsafe.Pointer) gd.PackedByteArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
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
Called when a packet needs to be sent by the [MultiplayerAPI], if [method _put_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (class) _put_packet_script(impl func(ptr unsafe.Pointer, p_buffer gd.PackedByteArray) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = pointers.New[gd.PackedByteArray](gd.UnsafeGet[gd.PackedPointers](p_args, 0))
		defer pointers.End(p_buffer)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to get the channel over which the next available packet was received. See [method MultiplayerPeer.get_packet_channel].
*/
func (class) _get_packet_channel(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to get the transfer mode the remote peer used to send the next available packet. See [method MultiplayerPeer.get_packet_mode].
*/
func (class) _get_packet_mode(impl func(ptr unsafe.Pointer) gdclass.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the channel to use is set for this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (class) _set_transfer_channel(impl func(ptr unsafe.Pointer, p_channel gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_channel = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_channel)
	}
}

/*
Called when the transfer channel to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (class) _get_transfer_channel(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the transfer mode is set on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (class) _set_transfer_mode(impl func(ptr unsafe.Pointer, p_mode gdclass.MultiplayerPeerTransferMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_mode = gd.UnsafeGet[gdclass.MultiplayerPeerTransferMode](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_mode)
	}
}

/*
Called when the transfer mode to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (class) _get_transfer_mode(impl func(ptr unsafe.Pointer) gdclass.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the target peer to use is set for this [MultiplayerPeer] (see [method MultiplayerPeer.set_target_peer]).
*/
func (class) _set_target_peer(impl func(ptr unsafe.Pointer, p_peer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_peer = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_peer)
	}
}

/*
Called when the ID of the [MultiplayerPeer] who sent the most recent packet is requested (see [method MultiplayerPeer.get_packet_peer]).
*/
func (class) _get_packet_peer(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the "is server" status is requested on the [MultiplayerAPI]. See [method MultiplayerAPI.is_server].
*/
func (class) _is_server(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [MultiplayerAPI] is polled. See [method MultiplayerAPI.poll].
*/
func (class) _poll(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the multiplayer peer should be immediately closed (see [method MultiplayerPeer.close]).
*/
func (class) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the connected [param p_peer] should be forcibly disconnected (see [method MultiplayerPeer.disconnect_peer]).
*/
func (class) _disconnect_peer(impl func(ptr unsafe.Pointer, p_peer gd.Int, p_force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_peer = gd.UnsafeGet[gd.Int](p_args, 0)

		var p_force = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_peer, p_force)
	}
}

/*
Called when the unique ID of this [MultiplayerPeer] is requested (see [method MultiplayerPeer.get_unique_id]). The value must be between [code]1[/code] and [code]2147483647[/code].
*/
func (class) _get_unique_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the "refuse new connections" status is set on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (class) _set_refuse_new_connections(impl func(ptr unsafe.Pointer, p_enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_enable = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_enable)
	}
}

/*
Called when the "refuse new connections" status is requested on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (class) _is_refusing_new_connections(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to check if the server can act as a relay in the current configuration. See [method MultiplayerPeer.is_server_relay_supported].
*/
func (class) _is_server_relay_supported(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the connection status is requested on the [MultiplayerPeer] (see [method MultiplayerPeer.get_connection_status]).
*/
func (class) _get_connection_status(impl func(ptr unsafe.Pointer) gdclass.MultiplayerPeerConnectionStatus) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsMultiplayerPeerExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerPeerExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMultiplayerPeer() MultiplayerPeer.Advanced {
	return *((*MultiplayerPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerPeer() MultiplayerPeer.Instance {
	return *((*MultiplayerPeer.Instance)(unsafe.Pointer(&self)))
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
	case "_get_packet_script":
		return reflect.ValueOf(self._get_packet_script)
	case "_put_packet_script":
		return reflect.ValueOf(self._put_packet_script)
	case "_get_packet_channel":
		return reflect.ValueOf(self._get_packet_channel)
	case "_get_packet_mode":
		return reflect.ValueOf(self._get_packet_mode)
	case "_set_transfer_channel":
		return reflect.ValueOf(self._set_transfer_channel)
	case "_get_transfer_channel":
		return reflect.ValueOf(self._get_transfer_channel)
	case "_set_transfer_mode":
		return reflect.ValueOf(self._set_transfer_mode)
	case "_get_transfer_mode":
		return reflect.ValueOf(self._get_transfer_mode)
	case "_set_target_peer":
		return reflect.ValueOf(self._set_target_peer)
	case "_get_packet_peer":
		return reflect.ValueOf(self._get_packet_peer)
	case "_is_server":
		return reflect.ValueOf(self._is_server)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	case "_disconnect_peer":
		return reflect.ValueOf(self._disconnect_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_set_refuse_new_connections":
		return reflect.ValueOf(self._set_refuse_new_connections)
	case "_is_refusing_new_connections":
		return reflect.ValueOf(self._is_refusing_new_connections)
	case "_is_server_relay_supported":
		return reflect.ValueOf(self._is_server_relay_supported)
	case "_get_connection_status":
		return reflect.ValueOf(self._get_connection_status)
	default:
		return gd.VirtualByName(MultiplayerPeer.Advanced(self.AsMultiplayerPeer()), name)
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
	case "_get_packet_script":
		return reflect.ValueOf(self._get_packet_script)
	case "_put_packet_script":
		return reflect.ValueOf(self._put_packet_script)
	case "_get_packet_channel":
		return reflect.ValueOf(self._get_packet_channel)
	case "_get_packet_mode":
		return reflect.ValueOf(self._get_packet_mode)
	case "_set_transfer_channel":
		return reflect.ValueOf(self._set_transfer_channel)
	case "_get_transfer_channel":
		return reflect.ValueOf(self._get_transfer_channel)
	case "_set_transfer_mode":
		return reflect.ValueOf(self._set_transfer_mode)
	case "_get_transfer_mode":
		return reflect.ValueOf(self._get_transfer_mode)
	case "_set_target_peer":
		return reflect.ValueOf(self._set_target_peer)
	case "_get_packet_peer":
		return reflect.ValueOf(self._get_packet_peer)
	case "_is_server":
		return reflect.ValueOf(self._is_server)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	case "_disconnect_peer":
		return reflect.ValueOf(self._disconnect_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_set_refuse_new_connections":
		return reflect.ValueOf(self._set_refuse_new_connections)
	case "_is_refusing_new_connections":
		return reflect.ValueOf(self._is_refusing_new_connections)
	case "_is_server_relay_supported":
		return reflect.ValueOf(self._is_server_relay_supported)
	case "_get_connection_status":
		return reflect.ValueOf(self._get_connection_status)
	default:
		return gd.VirtualByName(MultiplayerPeer.Instance(self.AsMultiplayerPeer()), name)
	}
}
func init() {
	gdclass.Register("MultiplayerPeerExtension", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerPeerExtension{*(*gdclass.MultiplayerPeerExtension)(unsafe.Pointer(&ptr))}
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
