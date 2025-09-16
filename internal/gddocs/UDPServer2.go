/*
[gdscript]
# client_node.gd
class_name ClientNode
extends Node

var udp = PacketPeerUDP.new()
var connected = false

func _ready():
	udp.connect_to_host("127.0.0.1", 4242)

func _process(delta):
	if !connected:
		# Try to contact server
		udp.put_packet("The answer is... 42!".to_utf8_buffer())
	if udp.get_available_packet_count() > 0:
		print("Connected: %s" % udp.get_packet().get_string_from_utf8())
		connected = true
[/gdscript]
[csharp]
// ClientNode.cs
using Godot;

public partial class ClientNode : Node
{
	private PacketPeerUdp _udp = new PacketPeerUdp();
	private bool _connected = false;

	public override void _Ready()
	{
		_udp.ConnectToHost("127.0.0.1", 4242);
	}

	public override void _Process(double delta)
	{
		if (!_connected)
		{
			// Try to contact server
			_udp.PutPacket("The Answer Is..42!".ToUtf8Buffer());
		}
		if (_udp.GetAvailablePacketCount() > 0)
		{
			GD.Print($"Connected: {_udp.GetPacket().GetStringFromUtf8()}");
			_connected = true;
		}
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PacketPeerUDP"
)

type ClientNode struct {
	Node.Extension[ClientNode]

	udp       PacketPeerUDP.Instance
	connected bool
}

func (c *ClientNode) Ready() {
	c.udp = PacketPeerUDP.New()
	c.udp.ConnectToHost("127.0.0.1", 4242)
}

func (c *ClientNode) Process(delta float64) {
	if !c.connected {
		// Try to contact server
		c.udp.AsPacketPeer().PutPacket([]byte("The answer is... 42!"))
	}
	if c.udp.AsPacketPeer().GetAvailablePacketCount() > 0 {
		fmt.Println("Connected: ", string(c.udp.AsPacketPeer().GetPacket()))
		c.connected = true
	}
}
