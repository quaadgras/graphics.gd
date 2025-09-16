/*
[gdscript]
extends Node

var socket = WebSocketPeer.new()

func _ready():
	socket.connect_to_url("wss://example.com")

func _process(delta):
	socket.poll()
	var state = socket.get_ready_state()
	if state == WebSocketPeer.STATE_OPEN:
		while socket.get_available_packet_count():
			print("Packet: ", socket.get_packet())
	elif state == WebSocketPeer.STATE_CLOSING:
		# Keep polling to achieve proper close.
		pass
	elif state == WebSocketPeer.STATE_CLOSED:
		var code = socket.get_close_code()
		var reason = socket.get_close_reason()
		print("WebSocket closed with code: %d, reason %s. Clean: %s" % [code, reason, code != -1])
		set_process(false) # Stop processing.
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/WebSocketPeer"
)

type WebSocketNode struct {
	Node.Extension[WebSocketNode]

	socket WebSocketPeer.Instance
}

func (w *WebSocketNode) Ready() {
	w.socket = WebSocketPeer.New()
	w.socket.ConnectToUrl("wss://example.com")
}

func (w *WebSocketNode) Process(_ float64) {
	w.socket.Poll()
	var state = w.socket.GetReadyState()
	switch state {
	case WebSocketPeer.StateOpen:
		for w.socket.AsPacketPeer().GetAvailablePacketCount() > 0 {
			print("Packet: ", w.socket.AsPacketPeer().GetPacket())
		}
	case WebSocketPeer.StateClosing:
		// Keep polling to achieve proper close.
	case WebSocketPeer.StateClosed:
		var code = w.socket.GetCloseCode()
		var reason = w.socket.GetCloseReason()
		print("WebSocket closed with code: %d, reason %s. Clean: %s", code, reason, code != -1)
		w.AsNode().SetProcess(false) // Stop processing.
	}
}
