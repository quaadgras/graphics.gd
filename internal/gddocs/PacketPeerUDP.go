/*
var peer = PacketPeerUDP.new()

# Optionally, you can select the local port used to send the packet.
peer.bind(4444)

peer.set_dest_address("1.1.1.1", 4433)
peer.put_packet("hello".to_utf8_buffer())
*/

package main

import "graphics.gd/classdb/PacketPeerUDP"

func ExamplePacketPeerUDP() {
	var peer = PacketPeerUDP.New()
	peer.Bind(4444)
	peer.SetDestAddress("1.1.1.1", 4433)
	peer.AsPacketPeer().PutPacket([]byte("hello"))
}
