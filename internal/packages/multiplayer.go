package packages

// Multiplayer (high level)
var Multiplayer = []string{
	"MultiplayerAPI",
	"MultiplayerAPIExtension",
	"MultiplayerPeer",
	"MultiplayerPeerExtension",
	"MultiplayerSpawner",
	"MultiplayerSynchronizer",
	"SceneMultiplayer",
	"SceneReplicationConfig",
}

// MultiplayerNetworking, low level RPCs and replication, HTTP/TCP/UDP/DNS, WebSockets, ENet, encryption.
var MultiplayerNetworking = []string{
	"PacketPeer",
	"PacketPeerExtension",
	"PacketPeerStream",
	"StreamPeer",
	"StreamPeerBuffer",
	"StreamPeerExtension",
}

// MultiplayerNetworkingDTLS
var MultiplayerNetworkingDTLS = []string{
	"DTLSServer",
	"PacketPeerDTLS",
}

// MultiplayerNetworkingENet
var MultiplayerNetworkingENet = []string{
	"ENetConnection",
	"ENetMultiplayerPeer",
	"ENetPacketPeer",
}

// MultiplayerNetworkingHTTP
var MultiplayerNetworkingHTTP = []string{
	"HTTPClient",
	"HTTPRequest",
}

// MultiplayerNetworkingIP
var MultiplayerNetworkingIP = []string{
	"IP",
	"IPUnix",
}

// MultiplayerNetworkingSSL
var MultiplayerNetworkingSSL = []string{
	"StreamPeerSSL",
}

// MultiplayerNetworkingTCP
var MultiplayerNetworkingTCP = []string{
	"StreamPeerTCP",
	"TCPServer",
}

// MultiplayerNetworkingUDP
var MultiplayerNetworkingUDP = []string{
	"PacketPeerUDP",
	"UDPServer",
}

// MultiplayerNetworkingUPNP
var MultiplayerNetworkingUPNP = []string{
	"UPNP",
	"UPNPDevice",
}

// MultiplayerNetworkingWebRTC
var MultiplayerNetworkingWebRTC = []string{
	"WebRTCDataChannel",
	"WebRTCDataChannelExtension",
	"WebRTCMultiplayerPeer",
	"WebRTCPeerConnection",
	"WebRTCPeerConnectionExtension",
}

// MultiplayerNetworkingWebsocket
var MultiplayerNetworkingWebsocket = []string{
	"WebSocketClient",
	"WebSocketMultiplayerPeer",
	"WebSocketPeer",
	"WebSocketServer",
}
