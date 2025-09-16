/*
# Emitted when UPnP port mapping setup is completed (regardless of success or failure).
signal upnp_completed(error)

# Replace this with your own server port number between 1024 and 65535.
const SERVER_PORT = 3928
var thread = null

func _upnp_setup(server_port):
	# UPNP queries take some time.
	var upnp = UPNP.new()
	var err = upnp.discover()

	if err != OK:
		push_error(str(err))
		upnp_completed.emit(err)
		return

	if upnp.get_gateway() and upnp.get_gateway().is_valid_gateway():
		upnp.add_port_mapping(server_port, server_port, ProjectSettings.get_setting("application/config/name"), "UDP")
		upnp.add_port_mapping(server_port, server_port, ProjectSettings.get_setting("application/config/name"), "TCP")
		upnp_completed.emit(OK)

func _ready():
	thread = Thread.new()
	thread.start(_upnp_setup.bind(SERVER_PORT))

func _exit_tree():
	# Wait for thread finish here to handle game exit while the thread is running.
	thread.wait_to_finish()
*/

package main

import (
	"errors"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Thread"
	"graphics.gd/classdb/UPNP"
	"graphics.gd/classdb/UPNPDevice"
	"graphics.gd/variant/Error"
	"graphics.gd/variant/Signal"
)

const ServerPort = 3928 // Replace this with your own server port number between 1024 and 65535.

type MyPortForwarding struct {
	Node.Extension[MyPortForwarding]

	Completed Signal.Solo[Error.Code] // Emitted when UPnP port mapping setup is completed (regardless of success or failure).

	thread Thread.Instance
}

func (m *MyPortForwarding) setup(port int) {
	var upnp = UPNP.New()
	if err := upnp.Discover(); err != 0 {
		Engine.Raise(errors.New("error UPNP"))
		m.Completed.Emit(1)
		return
	}
	if gateway := upnp.GetGateway(); gateway != UPNPDevice.Nil && gateway.IsValidGateway() {
		UPNPDevice.Expanded(gateway).AddPortMapping(port, port, "MyGame", "UDP", 0)
		UPNPDevice.Expanded(gateway).AddPortMapping(port, port, "MyGame", "TCP", 0)
		m.Completed.Emit(0)
	}
}

func (m *MyPortForwarding) Ready() {
	m.thread = Thread.New()
	m.thread.Start(func() { m.setup(ServerPort) })
}

func (m *MyPortForwarding) ExitTree() {
	m.thread.WaitToFinish() // Wait for thread finish here to handle game exit while the thread is running.
}
