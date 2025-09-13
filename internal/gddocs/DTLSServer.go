/*
[gdscript]
# server_node.gd
extends Node

var dtls = DTLSServer.new()
var server = UDPServer.new()
var peers = []

func _ready():
    server.listen(4242)
    var key = load("key.key") # Your private key.
    var cert = load("cert.crt") # Your X509 certificate.
    dtls.setup(TlsOptions.server(key, cert))

func _process(delta):
    while server.is_connection_available():
        var peer = server.take_connection()
        var dtls_peer = dtls.take_connection(peer)
        if dtls_peer.get_status() != PacketPeerDTLS.STATUS_HANDSHAKING:
            continue # It is normal that 50% of the connections fails due to cookie exchange.
        print("Peer connected!")
        peers.append(dtls_peer)

    for p in peers:
        p.poll() # Must poll to update the state.
        if p.get_status() == PacketPeerDTLS.STATUS_CONNECTED:
            while p.get_available_packet_count() > 0:
                print("Received message from client: %s" % p.get_packet().get_string_from_utf8())
                p.put_packet("Hello DTLS client".to_utf8_buffer())
[/gdscript]
[csharp]
// ServerNode.cs
using Godot;

public partial class ServerNode : Node
{
    private DtlsServer _dtls = new DtlsServer();
    private UdpServer _server = new UdpServer();
    private Godot.Collections.Array<PacketPeerDtls> _peers = [];

    public override void _Ready()
    {
        _server.Listen(4242);
        var key = GD.Load<CryptoKey>("key.key"); // Your private key.
        var cert = GD.Load<X509Certificate>("cert.crt"); // Your X509 certificate.
        _dtls.Setup(TlsOptions.Server(key, cert));
    }

    public override void _Process(double delta)
    {
        while (_server.IsConnectionAvailable())
        {
            PacketPeerUdp peer = _server.TakeConnection();
            PacketPeerDtls dtlsPeer = _dtls.TakeConnection(peer);
            if (dtlsPeer.GetStatus() != PacketPeerDtls.Status.Handshaking)
            {
                continue; // It is normal that 50% of the connections fails due to cookie exchange.
            }
            GD.Print("Peer connected!");
            _peers.Add(dtlsPeer);
        }

        foreach (var p in _peers)
        {
            p.Poll(); // Must poll to update the state.
            if (p.GetStatus() == PacketPeerDtls.Status.Connected)
            {
                while (p.GetAvailablePacketCount() > 0)
                {
                    GD.Print($"Received Message From Client: {p.GetPacket().GetStringFromUtf8()}");
                    p.PutPacket("Hello DTLS Client".ToUtf8Buffer());
                }
            }
        }
    }
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/CryptoKey"
	"graphics.gd/classdb/DTLSServer"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PacketPeerDTLS"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/TLSOptions"
	"graphics.gd/classdb/UDPServer"
	"graphics.gd/classdb/X509Certificate"
	"graphics.gd/variant/Float"
)

type ServerDTLS struct {
	Node.Extension[ServerDTLS]

	dtls  DTLSServer.Instance
	udp   UDPServer.Instance
	peers []PacketPeerDTLS.Instance
}

func (srv *ServerDTLS) Ready() {
	srv.udp.Listen(4242)
	var key = Resource.Load[CryptoKey.Instance]("key.key")
	var cert = Resource.Load[X509Certificate.Instance]("cert.crt")
	srv.dtls.Setup(TLSOptions.Server(key, cert))
}

func (srv *ServerDTLS) Process(delta Float.X) {
	for srv.udp.IsConnectionAvailable() {
		var peer = srv.udp.TakeConnection()
		var dtlsPeer = srv.dtls.TakeConnection(peer)
		if dtlsPeer.GetStatus() != PacketPeerDTLS.StatusHandshaking {
			continue // It is normal that 50%!o(MISSING)f the connections fails due to cookie exchange.
		}
		fmt.Println("Peer connected!")
		srv.peers = append(srv.peers, dtlsPeer)
	}
	for _, p := range srv.peers {
		p.Poll() // Must poll to update the state.
		if p.GetStatus() == PacketPeerDTLS.StatusConnected {
			for p.AsPacketPeer().GetAvailablePacketCount() > 0 {
				fmt.Println("Received message from client: ", string(p.AsPacketPeer().GetPacket()))
				p.AsPacketPeer().PutPacket([]byte("Hello DTLS client"))
			}
		}
	}
}
