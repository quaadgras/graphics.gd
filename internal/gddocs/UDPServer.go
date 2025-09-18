/*
[gdscript]
# server_node.gd
class_name ServerNode
extends Node

var server = UDPServer.new()
var peers = []

func _ready():
	server.listen(4242)

func _process(delta):
	server.poll() # Important!
	if server.is_connection_available():
		var peer = server.take_connection()
		var packet = peer.get_packet()
		print("Accepted peer: %s:%s" % [peer.get_packet_ip(), peer.get_packet_port()])
		print("Received data: %s" % [packet.get_string_from_utf8()])
		# Reply so it knows we received the message.
		peer.put_packet(packet)
		# Keep a reference so we can keep contacting the remote peer.
		peers.append(peer)

	for i in range(0, peers.size()):
		pass # Do something with the connected peers.
[/gdscript]
[csharp]
// ServerNode.cs
using Godot;
using System.Collections.Generic;

public partial class ServerNode : Node
{
	private UdpServer _server = new UdpServer();
	private List<PacketPeerUdp> _peers  = new List<PacketPeerUdp>();

	public override void _Ready()
	{
		_server.Listen(4242);
	}

	public override void _Process(double delta)
	{
		_server.Poll(); // Important!
		if (_server.IsConnectionAvailable())
		{
			PacketPeerUdp peer = _server.TakeConnection();
			byte[] packet = peer.GetPacket();
			GD.Print($"Accepted Peer: {peer.GetPacketIP()}:{peer.GetPacketPort()}");
			GD.Print($"Received Data: {packet.GetStringFromUtf8()}");
			// Reply so it knows we received the message.
			peer.PutPacket(packet);
			// Keep a reference so we can keep contacting the remote peer.
			_peers.Add(peer);
		}
		foreach (var peer in _peers)
		{
			// Do something with the peers.
		}
	}
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PacketPeerUDP"
	"graphics.gd/classdb/UDPServer"
)

type ServerNode struct {
	Node.Extension[ServerNode]

	udp   UDPServer.Instance
	peers []PacketPeerUDP.Instance
}

func (s *ServerNode) Ready() {
	s.udp = UDPServer.New()
	s.udp.Listen(4242)
}

func (s *ServerNode) Process(_ float64) {
	s.udp.Poll() // Important!
	if s.udp.IsConnectionAvailable() {
		peer := s.udp.TakeConnection()
		packet := peer.AsPacketPeer().GetPacket()
		print("Accepted peer: ", peer.GetPacketIp(), ":", peer.GetPacketPort())
		print("Received data: ", string(packet))
		// Reply so it knows we received the message.
		peer.AsPacketPeer().PutPacket(packet)
		// Keep a reference so we can keep contacting the remote peer.
		s.peers = append(s.peers, peer)
	}
	for _, peer := range s.peers {
		_ = peer // Do something with the connected peers.
	}
}
