/*
[gdscript]
# client_node.gd
extends Node

var dtls = PacketPeerDTLS.new()
var udp = PacketPeerUDP.new()
var connected = false

func _ready():
	udp.connect_to_host("127.0.0.1", 4242)
	dtls.connect_to_peer(udp, false) # Use true in production for certificate validation!

func _process(delta):
	dtls.poll()
	if dtls.get_status() == PacketPeerDTLS.STATUS_CONNECTED:
		if !connected:
			# Try to contact server
			dtls.put_packet("The answer is... 42!".to_utf8_buffer())
		while dtls.get_available_packet_count() > 0:
			print("Connected: %s" % dtls.get_packet().get_string_from_utf8())
			connected = true
[/gdscript]
[csharp]
// ClientNode.cs
using Godot;
using System.Text;

public partial class ClientNode : Node
{
	private PacketPeerDtls _dtls = new PacketPeerDtls();
	private PacketPeerUdp _udp = new PacketPeerUdp();
	private bool _connected = false;

	public override void _Ready()
	{
		_udp.ConnectToHost("127.0.0.1", 4242);
		_dtls.ConnectToPeer(_udp, validateCerts: false); // Use true in production for certificate validation!
	}

	public override void _Process(double delta)
	{
		_dtls.Poll();
		if (_dtls.GetStatus() == PacketPeerDtls.Status.Connected)
		{
			if (!_connected)
			{
				// Try to contact server
				_dtls.PutPacket("The Answer Is..42!".ToUtf8Buffer());
			}
			while (_dtls.GetAvailablePacketCount() > 0)
			{
				GD.Print($"Connected: {_dtls.GetPacket().GetStringFromUtf8()}");
				_connected = true;
			}
		}
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PacketPeerDTLS"
	"graphics.gd/classdb/PacketPeerUDP"
	"graphics.gd/variant/Float"
)

type ClientDTLS struct {
	Node.Extension[ClientDTLS]

	dtls PacketPeerDTLS.Instance
	udp  PacketPeerUDP.Instance

	connected bool
}

func (srv *ClientDTLS) Ready() {
	srv.udp.ConnectToHost("127.0.0.1", 4242)
	srv.dtls.ConnectToPeer(srv.udp, "") // Use hostname in production for certificate validation!
}

func (srv *ClientDTLS) Process(delta Float.X) {
	srv.dtls.Poll()
	if srv.dtls.GetStatus() == PacketPeerDTLS.StatusConnected {
		if !srv.connected {
			// Try to contact server
			srv.dtls.AsPacketPeer().PutPacket([]byte("The Answer Is... 42!"))
		}
		for srv.dtls.AsPacketPeer().GetAvailablePacketCount() > 0 {
			fmt.Println("Connected: " + string(srv.dtls.AsPacketPeer().GetPacket()))
			srv.connected = true
		}
	}
}
