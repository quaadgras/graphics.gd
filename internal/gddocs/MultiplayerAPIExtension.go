/*
[gdscript]
extends MultiplayerAPIExtension
class_name LogMultiplayer

# We want to extend the default SceneMultiplayer.
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

# These can be optional, but in our case we want to extend SceneMultiplayer, so forward everything.
func _set_multiplayer_peer(p_peer: MultiplayerPeer):
    base_multiplayer.multiplayer_peer = p_peer

func _get_multiplayer_peer() -> MultiplayerPeer:
    return base_multiplayer.multiplayer_peer

func _get_unique_id() -> int:
    return base_multiplayer.get_unique_id()

func _get_peer_ids() -> PackedInt32Array:
    return base_multiplayer.get_peers()
[/gdscript]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/MultiplayerAPI"
	"graphics.gd/classdb/MultiplayerAPIExtension"
	"graphics.gd/classdb/MultiplayerPeer"
	"graphics.gd/classdb/MultiplayerSpawner"
	"graphics.gd/classdb/MultiplayerSynchronizer"
	"graphics.gd/variant"
	"graphics.gd/variant/Object"
)

type LogMultiplayer struct {
	MultiplayerAPIExtension.Extension[LogMultiplayer]

	base MultiplayerAPI.Instance
}

func (m *LogMultiplayer) Init() {
	m.base.OnConnectedToServer(func() { MultiplayerAPI.Advanced(m.AsMultiplayerAPI()).ConnectedToServer().Emit() })
	m.base.OnConnectionFailed(func() { MultiplayerAPI.Advanced(m.AsMultiplayerAPI()).ConnectionFailed().Emit() })
	m.base.OnPeerConnected(func(id int) { MultiplayerAPI.Advanced(m.AsMultiplayerAPI()).PeerConnected().Emit(variant.New(id)) })
	m.base.OnPeerDisconnected(func(id int) { MultiplayerAPI.Advanced(m.AsMultiplayerAPI()).PeerDisconnected().Emit(variant.New(id)) })
}

func (m *LogMultiplayer) Poll() error { return m.base.Poll() }

func (m *LogMultiplayer) Rpc(peer int, object Object.Instance, method string, args []any) error {
	fmt.Println("Got RPC for", peer, ":", object, "::", method, "(", args, ")")
	return MultiplayerAPI.Expanded(m.base).Rpc(peer, object, method, args)
}

func (m *LogMultiplayer) ObjectConfigurationAdd(object Object.Instance, config any) error {
	if _, ok := config.(MultiplayerSynchronizer.Instance); ok {
		fmt.Println("Adding synchronization configuration for", object, ". Synchronizer:", config)
	} else if _, ok := config.(MultiplayerSpawner.Instance); ok {
		fmt.Println("Adding node", object, "to the spawn list. Spawner:", config)
	}
	return m.base.ObjectConfigurationAdd(object, config)
}

func (m *LogMultiplayer) ObjectConfigurationRemove(object Object.Instance, config any) error {
	if _, ok := config.(MultiplayerSynchronizer.Instance); ok {
		fmt.Println("Removing synchronization configuration for", object, ". Synchronizer:", config)
	} else if _, ok := config.(MultiplayerSpawner.Instance); ok {
		fmt.Println("Removing node", object, "from the spawn list. Spawner:", config)
	}
	return m.base.ObjectConfigurationRemove(object, config)
}

func (m *LogMultiplayer) SetMultiplayerPeer(p_peer MultiplayerPeer.Instance) {
	m.base.SetMultiplayerPeer(p_peer)
}

func (m *LogMultiplayer) GetMultiplayerPeer() MultiplayerPeer.Instance {
	return m.base.MultiplayerPeer()
}

func (m *LogMultiplayer) GetUniqueID() int {
	return m.base.GetUniqueId()
}

func (m *LogMultiplayer) GetPeerIDs() []int32 {
	return m.base.GetPeers()
}
