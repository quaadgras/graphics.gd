/*
{
	"negotiated": true, # When set to true (default off), means the channel is negotiated out of band. "id" must be set too. "data_channel_received" will not be called.
	"id": 1, # When "negotiated" is true this value must also be set to the same value on both peer.

	# Only one of maxRetransmits and maxPacketLifeTime can be specified, not both. They make the channel unreliable (but also better at real time).
	"maxRetransmits": 1, # Specify the maximum number of attempt the peer will make to retransmits packets if they are not acknowledged.
	"maxPacketLifeTime": 100, # Specify the maximum amount of time before giving up retransmitions of unacknowledged packets (in milliseconds).
	"ordered": true, # When in unreliable mode (i.e. either "maxRetransmits" or "maxPacketLifetime" is set), "ordered" (true by default) specify if packet ordering is to be enforced.

	"protocol": "my-custom-protocol", # A custom sub-protocol string for this channel.
}
*/

package main

import "graphics.gd/classdb/WebRTCPeerConnection"

func WebRTCPeerConnection_CreateDataChannel() {
	var example = WebRTCPeerConnection.Options{
		Negotiated: true, // When set to true (default off), means the channel is negotiated out of band. "id" must be set too. "data_channel_received" will not be called.
		ID:         1,    // When "negotiated" is true this value must also be set to the same value on both peer.
		// Only one of maxRetransmits and maxPacketLifeTime can be specified, not both. They make the channel unreliable (but also better at real time).
		MaxRetransmits:    1,                    // Specify the maximum number of attempt the peer will make to retransmits packets if they are not acknowledged.
		MaxPacketLifeTime: 100,                  // Specify the maximum amount of time before giving up retransmitions of unacknowledged packets (in milliseconds).
		Ordered:           true,                 // When in unreliable mode (i.e. either "maxRetransmits" or "maxPacketLifetime" is set), "ordered" (true by default) specify if packet ordering is to be enforced.
		Protocol:          "my-custom-protocol", // A custom sub-protocol string for this channel.
	}
	_ = example
}
