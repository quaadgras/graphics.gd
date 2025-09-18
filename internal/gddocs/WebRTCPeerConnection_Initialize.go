/*
{
	"iceServers": [
		{
			"urls": [ "stun:stun.example.com:3478" ], # One or more STUN servers.
		},
		{
			"urls": [ "turn:turn.example.com:3478" ], # One or more TURN servers.
			"username": "a_username", # Optional username for the TURN server.
			"credential": "a_password", # Optional password for the TURN server.
		}
	]
}
*/

package main

import "graphics.gd/classdb/WebRTCPeerConnection"

func WebRTCPeerConnection_Initialize() {
	var example = WebRTCPeerConnection.Configuration{
		IceServers: []WebRTCPeerConnection.IceServer{
			{
				URLs: []string{"stun:stun.example.com:3478"}, // One or more STUN servers.
			},
			{
				URLs:       []string{"turn:turn.example.com:3478"}, // One or more TURN servers.
				Username:   "a_username",                           // Optional username for the TURN server.
				Credential: "a_password",                           // Optional password for the TURN server.
			},
		},
	}
	_ = example
}
