// Package MultiplayerAPIExtension provides methods for working with MultiplayerAPIExtension object instances.
package MultiplayerAPIExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/MultiplayerAPI"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class can be used to augment or replace the default [MultiplayerAPI] implementation via script or extensions.
The following example augment the default implementation ([SceneMultiplayer]) by logging every RPC being made, and every object being configured for replication.
[codeblocks]
[gdscript]
extends MultiplayerAPIExtension
class_name LogMultiplayer

# We want to augment the default SceneMultiplayer.
var base_multiplayer = SceneMultiplayer.new()

func _init():

	# Just passthrough base signals (copied to var to avoid cyclic reference)
	var cts = connected_to_server
	var cf = connection_failed
	var pc = peer_connected
	var pd = peer_disconnected
	base_multiplayer.connected_to_server.connect(func(): cts.emit())
	base_multiplayer.connection_failed.connect(func(): cf.emit())
	base_multiplayer.peer_connected.connect(func(id): pc.emit(id))
	base_multiplayer.peer_disconnected.connect(func(id): pd.emit(id))

func _poll():

	return base_multiplayer.poll()

# Log RPC being made and forward it to the default multiplayer.
func _rpc(peer: int, object: Object, method: StringName, args: Array) -> Error:

	print("Got RPC for %d: %s::%s(%s)" % [peer, object, method, args])
	return base_multiplayer.rpc(peer, object, method, args)

# Log configuration add. E.g. root path (nullptr, NodePath), replication (Node, Spawner|Synchronizer), custom.
func _object_configuration_add(object, config: Variant) -> Error:

	if config is MultiplayerSynchronizer:
	    print("Adding synchronization configuration for %s. Synchronizer: %s" % [object, config])
	elif config is MultiplayerSpawner:
	    print("Adding node %s to the spawn list. Spawner: %s" % [object, config])
	return base_multiplayer.object_configuration_add(object, config)

# Log configuration remove. E.g. root path (nullptr, NodePath), replication (Node, Spawner|Synchronizer), custom.
func _object_configuration_remove(object, config: Variant) -> Error:

	if config is MultiplayerSynchronizer:
	    print("Removing synchronization configuration for %s. Synchronizer: %s" % [object, config])
	elif config is MultiplayerSpawner:
	    print("Removing node %s from the spawn list. Spawner: %s" % [object, config])
	return base_multiplayer.object_configuration_remove(object, config)

# These can be optional, but in our case we want to augment SceneMultiplayer, so forward everything.
func _set_multiplayer_peer(p_peer: MultiplayerPeer):

	base_multiplayer.multiplayer_peer = p_peer

func _get_multiplayer_peer() -> MultiplayerPeer:

	return base_multiplayer.multiplayer_peer

func _get_unique_id() -> int:

	return base_multiplayer.get_unique_id()

func _get_peer_ids() -> PackedInt32Array:

	return base_multiplayer.get_peers()

[/gdscript]
[/codeblocks]
Then in your main scene or in an autoload call [method SceneTree.set_multiplayer] to start using your custom [MultiplayerAPI]:
[codeblocks]
[gdscript]
# autoload.gd
func _enter_tree():

	# Sets our custom multiplayer as the main one in SceneTree.

get_tree().set_multiplayer(LogMultiplayer.new())
[/gdscript]
[/codeblocks]
Native extensions can alternatively use the [method MultiplayerAPI.set_default_interface] method during initialization to configure themselves as the default implementation.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=MultiplayerAPIExtension)
*/
type Instance [1]gdclass.MultiplayerAPIExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiplayerAPIExtension() Instance
}
type Interface interface {
	//Callback for [method MultiplayerAPI.poll].
	Poll() error
	//Called when the [member MultiplayerAPI.multiplayer_peer] is set.
	SetMultiplayerPeer(multiplayer_peer [1]gdclass.MultiplayerPeer)
	//Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
	GetMultiplayerPeer() [1]gdclass.MultiplayerPeer
	//Callback for [method MultiplayerAPI.get_unique_id].
	GetUniqueId() int
	//Callback for [method MultiplayerAPI.get_peers].
	GetPeerIds() []int32
	//Callback for [method MultiplayerAPI.rpc].
	Rpc(peer int, obj Object.Instance, method string, args Array.Any) error
	//Callback for [method MultiplayerAPI.get_remote_sender_id].
	GetRemoteSenderId() int
	//Callback for [method MultiplayerAPI.object_configuration_add].
	ObjectConfigurationAdd(obj Object.Instance, configuration any) error
	//Callback for [method MultiplayerAPI.object_configuration_remove].
	ObjectConfigurationRemove(obj Object.Instance, configuration any) error
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) Poll() (_ error)                                                { return }
func (self Implementation) SetMultiplayerPeer(multiplayer_peer [1]gdclass.MultiplayerPeer) { return }
func (self Implementation) GetMultiplayerPeer() (_ [1]gdclass.MultiplayerPeer)             { return }
func (self Implementation) GetUniqueId() (_ int)                                           { return }
func (self Implementation) GetPeerIds() (_ []int32)                                        { return }
func (self Implementation) Rpc(peer int, obj Object.Instance, method string, args Array.Any) (_ error) {
	return
}
func (self Implementation) GetRemoteSenderId() (_ int) { return }
func (self Implementation) ObjectConfigurationAdd(obj Object.Instance, configuration any) (_ error) {
	return
}
func (self Implementation) ObjectConfigurationRemove(obj Object.Instance, configuration any) (_ error) {
	return
}

/*
Callback for [method MultiplayerAPI.poll].
*/
func (Instance) _poll(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (Instance) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer [1]gdclass.MultiplayerPeer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var multiplayer_peer = [1]gdclass.MultiplayerPeer{pointers.New[gdclass.MultiplayerPeer]([3]uint64{uint64(gd.UnsafeGet[uintptr](p_args, 0))})}
		defer pointers.End(multiplayer_peer[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, multiplayer_peer)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (Instance) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) [1]gdclass.MultiplayerPeer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (Instance) _get_unique_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (Instance) _get_peer_ids(impl func(ptr unsafe.Pointer) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.rpc].
*/
func (Instance) _rpc(impl func(ptr unsafe.Pointer, peer int, obj Object.Instance, method string, args Array.Any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var peer = gd.UnsafeGet[gd.Int](p_args, 0)
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(obj[0])
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		defer pointers.End(method)
		var args = pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		defer pointers.End(args)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(peer), obj, method.String(), args)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (Instance) _get_remote_sender_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (Instance) _object_configuration_add(impl func(ptr unsafe.Pointer, obj Object.Instance, configuration any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(configuration)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration.Interface())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (Instance) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj Object.Instance, configuration any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(configuration)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration.Interface())
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerAPIExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerAPIExtension"))
	casted := Instance{*(*gdclass.MultiplayerAPIExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Callback for [method MultiplayerAPI.poll].
*/
func (class) _poll(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (class) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer [1]gdclass.MultiplayerPeer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var multiplayer_peer = [1]gdclass.MultiplayerPeer{pointers.New[gdclass.MultiplayerPeer]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(multiplayer_peer[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, multiplayer_peer)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (class) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) [1]gdclass.MultiplayerPeer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (class) _get_unique_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (class) _get_peer_ids(impl func(ptr unsafe.Pointer) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
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
Callback for [method MultiplayerAPI.rpc].
*/
func (class) _rpc(impl func(ptr unsafe.Pointer, peer gd.Int, obj [1]gd.Object, method gd.StringName, args gd.Array) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var peer = gd.UnsafeGet[gd.Int](p_args, 0)
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(obj[0])
		var method = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))
		var args = pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, peer, obj, method, args)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (class) _get_remote_sender_id(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (class) _object_configuration_add(impl func(ptr unsafe.Pointer, obj [1]gd.Object, configuration gd.Variant) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (class) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj [1]gd.Object, configuration gd.Variant) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsMultiplayerAPIExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerAPIExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMultiplayerAPI() MultiplayerAPI.Advanced {
	return *((*MultiplayerAPI.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerAPI() MultiplayerAPI.Instance {
	return *((*MultiplayerAPI.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_set_multiplayer_peer":
		return reflect.ValueOf(self._set_multiplayer_peer)
	case "_get_multiplayer_peer":
		return reflect.ValueOf(self._get_multiplayer_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_get_peer_ids":
		return reflect.ValueOf(self._get_peer_ids)
	case "_rpc":
		return reflect.ValueOf(self._rpc)
	case "_get_remote_sender_id":
		return reflect.ValueOf(self._get_remote_sender_id)
	case "_object_configuration_add":
		return reflect.ValueOf(self._object_configuration_add)
	case "_object_configuration_remove":
		return reflect.ValueOf(self._object_configuration_remove)
	default:
		return gd.VirtualByName(MultiplayerAPI.Advanced(self.AsMultiplayerAPI()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_set_multiplayer_peer":
		return reflect.ValueOf(self._set_multiplayer_peer)
	case "_get_multiplayer_peer":
		return reflect.ValueOf(self._get_multiplayer_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_get_peer_ids":
		return reflect.ValueOf(self._get_peer_ids)
	case "_rpc":
		return reflect.ValueOf(self._rpc)
	case "_get_remote_sender_id":
		return reflect.ValueOf(self._get_remote_sender_id)
	case "_object_configuration_add":
		return reflect.ValueOf(self._object_configuration_add)
	case "_object_configuration_remove":
		return reflect.ValueOf(self._object_configuration_remove)
	default:
		return gd.VirtualByName(MultiplayerAPI.Instance(self.AsMultiplayerAPI()), name)
	}
}
func init() {
	gdclass.Register("MultiplayerAPIExtension", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerAPIExtension{*(*gdclass.MultiplayerAPIExtension)(unsafe.Pointer(&ptr))}
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
