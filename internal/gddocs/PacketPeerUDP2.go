/*
var peer

func _ready():
	peer = PacketPeerUDP.new()
	peer.bind(4433)


func _process(_delta):
	if peer.get_available_packet_count() > 0:
		var array_bytes = peer.get_packet()
		var packet_string = array_bytes.get_string_from_ascii()
		print("Received message: ", packet_string)
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PacketPeerUDP"
)

type MyPacketPeerUDP struct {
	Node.Extension[MyPacketPeerUDP]

	peer PacketPeerUDP.Instance
}

func (m *MyPacketPeerUDP) Ready() {
	m.peer = PacketPeerUDP.New()
	m.peer.Bind(4433)
}

func (m *MyPacketPeerUDP) Process(_ float64) {
	if m.peer.AsPacketPeer().GetAvailablePacketCount() > 0 {
		var array_bytes = m.peer.AsPacketPeer().GetPacket()
		var packet_string = string(array_bytes)
		print("Received message: ", packet_string)
	}
}
