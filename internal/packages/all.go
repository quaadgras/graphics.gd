// Package packages provides a metapackage that aims to organise the Godot classes into sensible
// divisions based on https://godotengine.org/teams to help reduce the intimidating amount of
// global classes in the Godot API into digestible and explorable packages.
package packages

// List of packages. GD is excluded and needs to be manually added
// to the root of the repo.
var List = map[string][]string{
	"audio":     Audio,
	"audio/ogg": AudioOgg,

	"core":        Core,
	"core/binary": CoreBinary,
	"core/editor": CoreEditor,
	"core/engine": CoreEngine,
	"core/script": CoreScript,

	"encryption":      Encryption,
	"encryption/aes":  EncryptionAES,
	"encryption/hmac": EncryptionHMAC,
	"encryption/x509": EncryptionX509,

	"navigation": Navigation,

	"multiplayer":                      Multiplayer,
	"multiplayer/networking":           MultiplayerNetworking,
	"multiplayer/networking/dtls":      MultiplayerNetworkingDTLS,
	"multiplayer/networking/enet":      MultiplayerNetworkingENet,
	"multiplayer/networking/http":      MultiplayerNetworkingHTTP,
	"multiplayer/networking/ip":        MultiplayerNetworkingIP,
	"multiplayer/networking/ssl":       MultiplayerNetworkingSSL,
	"multiplayer/networking/tcp":       MultiplayerNetworkingTCP,
	"multiplayer/networking/udp":       MultiplayerNetworkingUDP,
	"multiplayer/networking/upnp":      MultiplayerNetworkingUPNP,
	"multiplayer/networking/webrtc":    MultiplayerNetworkingWebRTC,
	"multiplayer/networking/websocket": MultiplayerNetworkingWebsocket,

	"physics": Physics,
	"spatial": Spatial,

	"system":       System,
	"system/input": SystemInput,

	"textual":     Textual,
	"textual/xml": TextualXML,

	"visual":                     Visual,
	"visual/animation":           VisualAnimation,
	"visual/display":             VisualDisplay,
	"visual/model":               VisualModel,
	"visual/model/csg":           VisualModelCSG,
	"visual/model/gltf":          VisualModelGLTF,
	"visual/reality":             VisualReality,
	"visual/rendering":           VisualRendering,
	"visual/rendering/gpu":       VisualRenderingGPU,
	"visual/rendering/particles": VisualRenderingParticles,
	"visual/shader":              VisualShader,
}
