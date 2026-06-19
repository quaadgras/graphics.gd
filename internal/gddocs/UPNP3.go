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

	if err != UPNP.UPNP_RESULT_SUCCESS:
		push_error(str(err))
		upnp_completed.emit(err)
		return

	if upnp.get_gateway() and upnp.get_gateway().is_valid_gateway():
		upnp.add_port_mapping(server_port, server_port, ProjectSettings.get_setting("application/config/name"), "UDP")
		upnp.add_port_mapping(server_port, server_port, ProjectSettings.get_setting("application/config/name"), "TCP")
		upnp_completed.emit(UPNP.UPNP_RESULT_SUCCESS)

func _ready():
	thread = Thread.new()
	thread.start(_upnp_setup.bind(SERVER_PORT))

func _exit_tree():
	# Wait for thread finish here to handle game exit while the thread is running.
	thread.wait_to_finish()
*/

package main

import (
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/UPNP"
	"graphics.gd/classdb/UPNPDevice"
)

// Core of _upnp_setup; the surrounding Thread + upnp_completed signal are omitted.
func exampleUPNPSetup(serverPort int) {
	// UPNP queries take some time.
	var upnp = UPNP.New()
	var err = upnp.Discover()
	if UPNP.UPNPResult(err) != UPNP.UpnpResultSuccess {
		// push_error(str(err)); upnp_completed.emit(err)
		return
	}
	if gateway := upnp.GetGateway(); gateway != UPNPDevice.Nil && gateway.IsValidGateway() {
		name, _ := ProjectSettings.GetSetting("application/config/name", "").(string)
		upnp.MoreArgs().AddPortMapping(serverPort, serverPort, name, "UDP", 0)
		upnp.MoreArgs().AddPortMapping(serverPort, serverPort, name, "TCP", 0)
		// upnp_completed.emit(UPNP.UpnpResultSuccess)
	}
}
