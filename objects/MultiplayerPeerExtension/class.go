package MultiplayerPeerExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/MultiplayerPeer"
import "grow.graphics/gd/objects/PacketPeer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class is designed to be inherited from a GDExtension plugin to implement custom networking layers for the multiplayer API (such as WebRTC). All the methods below [b]must[/b] be implemented to have a working custom multiplayer implementation. See also [MultiplayerAPI].

	// MultiplayerPeerExtension methods that can be overridden by a [Class] that extends it.
	type MultiplayerPeerExtension interface {
		//Called when a packet needs to be received by the [MultiplayerAPI], with [param r_buffer_size] being the size of the binary [param r_buffer] in bytes.
		GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) gd.Error
		//Called when a packet needs to be sent by the [MultiplayerAPI], with [param p_buffer_size] being the size of the binary [param p_buffer] in bytes.
		PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) gd.Error
		//Called when the available packet count is internally requested by the [MultiplayerAPI].
		GetAvailablePacketCount() int
		//Called when the maximum allowed packet size (in bytes) is requested by the [MultiplayerAPI].
		GetMaxPacketSize() int
		//Called when a packet needs to be received by the [MultiplayerAPI], if [method _get_packet] isn't implemented. Use this when extending this class via GDScript.
		GetPacketScript() []byte
		//Called when a packet needs to be sent by the [MultiplayerAPI], if [method _put_packet] isn't implemented. Use this when extending this class via GDScript.
		PutPacketScript(p_buffer []byte) gd.Error
		//Called to get the channel over which the next available packet was received. See [method MultiplayerPeer.get_packet_channel].
		GetPacketChannel() int
		//Called to get the transfer mode the remote peer used to send the next available packet. See [method MultiplayerPeer.get_packet_mode].
		GetPacketMode() classdb.MultiplayerPeerTransferMode
		//Called when the channel to use is set for this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
		SetTransferChannel(p_channel int)
		//Called when the transfer channel to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
		GetTransferChannel() int
		//Called when the transfer mode is set on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
		SetTransferMode(p_mode classdb.MultiplayerPeerTransferMode)
		//Called when the transfer mode to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
		GetTransferMode() classdb.MultiplayerPeerTransferMode
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
		GetConnectionStatus() classdb.MultiplayerPeerConnectionStatus
	}
*/
type Instance [1]classdb.MultiplayerPeerExtension

/*
Called when a packet needs to be received by the [MultiplayerAPI], with [param r_buffer_size] being the size of the binary [param r_buffer] in bytes.
*/
func (Instance) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (Instance) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size int) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the maximum allowed packet size (in bytes) is requested by the [MultiplayerAPI].
*/
func (Instance) _get_max_packet_size(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when a packet needs to be received by the [MultiplayerAPI], if [method _get_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (Instance) _get_packet_script(impl func(ptr unsafe.Pointer) []byte) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (Instance) _put_packet_script(impl func(ptr unsafe.Pointer, p_buffer []byte) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_buffer = pointers.New[gd.PackedByteArray](gd.UnsafeGet[[2]uintptr](p_args, 0))
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
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called to get the transfer mode the remote peer used to send the next available packet. See [method MultiplayerPeer.get_packet_mode].
*/
func (Instance) _get_packet_mode(impl func(ptr unsafe.Pointer) classdb.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the channel to use is set for this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (Instance) _set_transfer_channel(impl func(ptr unsafe.Pointer, p_channel int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_channel = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(p_channel))
	}
}

/*
Called when the transfer channel to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (Instance) _get_transfer_channel(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the transfer mode is set on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (Instance) _set_transfer_mode(impl func(ptr unsafe.Pointer, p_mode classdb.MultiplayerPeerTransferMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_mode = gd.UnsafeGet[classdb.MultiplayerPeerTransferMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_mode)
	}
}

/*
Called when the transfer mode to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (Instance) _get_transfer_mode(impl func(ptr unsafe.Pointer) classdb.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the target peer to use is set for this [MultiplayerPeer] (see [method MultiplayerPeer.set_target_peer]).
*/
func (Instance) _set_target_peer(impl func(ptr unsafe.Pointer, p_peer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_peer = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(p_peer))
	}
}

/*
Called when the ID of the [MultiplayerPeer] who sent the most recent packet is requested (see [method MultiplayerPeer.get_packet_peer]).
*/
func (Instance) _get_packet_peer(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the "is server" status is requested on the [MultiplayerAPI]. See [method MultiplayerAPI.is_server].
*/
func (Instance) _is_server(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [MultiplayerAPI] is polled. See [method MultiplayerAPI.poll].
*/
func (Instance) _poll(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the multiplayer peer should be immediately closed (see [method MultiplayerPeer.close]).
*/
func (Instance) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the connected [param p_peer] should be forcibly disconnected (see [method MultiplayerPeer.disconnect_peer]).
*/
func (Instance) _disconnect_peer(impl func(ptr unsafe.Pointer, p_peer int, p_force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called when the "refuse new connections" status is set on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (Instance) _set_refuse_new_connections(impl func(ptr unsafe.Pointer, p_enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_enable = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_enable)
	}
}

/*
Called when the "refuse new connections" status is requested on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (Instance) _is_refusing_new_connections(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to check if the server can act as a relay in the current configuration. See [method MultiplayerPeer.is_server_relay_supported].
*/
func (Instance) _is_server_relay_supported(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the connection status is requested on the [MultiplayerPeer] (see [method MultiplayerPeer.get_connection_status]).
*/
func (Instance) _get_connection_status(impl func(ptr unsafe.Pointer) classdb.MultiplayerPeerConnectionStatus) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MultiplayerPeerExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerPeerExtension"))
	return Instance{classdb.MultiplayerPeerExtension(object)}
}

/*
Called when a packet needs to be received by the [MultiplayerAPI], with [param r_buffer_size] being the size of the binary [param r_buffer] in bytes.
*/
func (class) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (class) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size gd.Int) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the maximum allowed packet size (in bytes) is requested by the [MultiplayerAPI].
*/
func (class) _get_max_packet_size(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when a packet needs to be received by the [MultiplayerAPI], if [method _get_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (class) _get_packet_script(impl func(ptr unsafe.Pointer) gd.PackedByteArray) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called when a packet needs to be sent by the [MultiplayerAPI], if [method _put_packet] isn't implemented. Use this when extending this class via GDScript.
*/
func (class) _put_packet_script(impl func(ptr unsafe.Pointer, p_buffer gd.PackedByteArray) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_buffer = pointers.New[gd.PackedByteArray](gd.UnsafeGet[[2]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to get the channel over which the next available packet was received. See [method MultiplayerPeer.get_packet_channel].
*/
func (class) _get_packet_channel(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to get the transfer mode the remote peer used to send the next available packet. See [method MultiplayerPeer.get_packet_mode].
*/
func (class) _get_packet_mode(impl func(ptr unsafe.Pointer) classdb.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the channel to use is set for this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (class) _set_transfer_channel(impl func(ptr unsafe.Pointer, p_channel gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_channel = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_channel)
	}
}

/*
Called when the transfer channel to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_channel]).
*/
func (class) _get_transfer_channel(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the transfer mode is set on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (class) _set_transfer_mode(impl func(ptr unsafe.Pointer, p_mode classdb.MultiplayerPeerTransferMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_mode = gd.UnsafeGet[classdb.MultiplayerPeerTransferMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_mode)
	}
}

/*
Called when the transfer mode to use is read on this [MultiplayerPeer] (see [member MultiplayerPeer.transfer_mode]).
*/
func (class) _get_transfer_mode(impl func(ptr unsafe.Pointer) classdb.MultiplayerPeerTransferMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the target peer to use is set for this [MultiplayerPeer] (see [method MultiplayerPeer.set_target_peer]).
*/
func (class) _set_target_peer(impl func(ptr unsafe.Pointer, p_peer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_peer = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_peer)
	}
}

/*
Called when the ID of the [MultiplayerPeer] who sent the most recent packet is requested (see [method MultiplayerPeer.get_packet_peer]).
*/
func (class) _get_packet_peer(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the "is server" status is requested on the [MultiplayerAPI]. See [method MultiplayerAPI.is_server].
*/
func (class) _is_server(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [MultiplayerAPI] is polled. See [method MultiplayerAPI.poll].
*/
func (class) _poll(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the multiplayer peer should be immediately closed (see [method MultiplayerPeer.close]).
*/
func (class) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the connected [param p_peer] should be forcibly disconnected (see [method MultiplayerPeer.disconnect_peer]).
*/
func (class) _disconnect_peer(impl func(ptr unsafe.Pointer, p_peer gd.Int, p_force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the "refuse new connections" status is set on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (class) _set_refuse_new_connections(impl func(ptr unsafe.Pointer, p_enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var p_enable = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_enable)
	}
}

/*
Called when the "refuse new connections" status is requested on this [MultiplayerPeer] (see [member MultiplayerPeer.refuse_new_connections]).
*/
func (class) _is_refusing_new_connections(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called to check if the server can act as a relay in the current configuration. See [method MultiplayerPeer.is_server_relay_supported].
*/
func (class) _is_server_relay_supported(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the connection status is requested on the [MultiplayerPeer] (see [method MultiplayerPeer.get_connection_status]).
*/
func (class) _get_connection_status(impl func(ptr unsafe.Pointer) classdb.MultiplayerPeerConnectionStatus) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
		return gd.VirtualByName(self.AsMultiplayerPeer(), name)
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
		return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}
func init() {
	classdb.Register("MultiplayerPeerExtension", func(ptr gd.Object) any { return classdb.MultiplayerPeerExtension(ptr) })
}
