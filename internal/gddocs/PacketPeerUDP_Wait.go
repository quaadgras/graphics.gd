/*
[gdscript]
socket = PacketPeerUDP.new()
# Server
socket.set_dest_address("127.0.0.1", 789)
socket.put_packet("Time to stop".to_ascii_buffer())

# Client
while socket.wait() == OK:
	var data = socket.get_packet().get_string_from_ascii()
	if data == "Time to stop":
		return
[/gdscript]
[csharp]
var socket = new PacketPeerUdp();
// Server
socket.SetDestAddress("127.0.0.1", 789);
socket.PutPacket("Time to stop".ToAsciiBuffer());

// Client
while (socket.Wait() == OK)
{
	string data = socket.GetPacket().GetStringFromASCII();
	if (data == "Time to stop")
	{
		return;
	}
}
[/csharp]
*/

package main

import "graphics.gd/classdb/PacketPeerUDP"

func ExamplePacketPeerUDPWait(socket PacketPeerUDP.Instance) {
	// Server
	socket.SetDestAddress("127.0.0.1", 789)
	socket.AsPacketPeer().PutPacket([]byte("Time to stop"))

	// Client
	for socket.Wait() == nil {
		var data = string(socket.AsPacketPeer().GetPacket())
		if data == "Time to stop" {
			return
		}
	}
}
