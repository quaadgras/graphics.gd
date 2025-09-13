/*
upnp.delete_port_mapping(port)
*/

package main

import "graphics.gd/classdb/UPNP"

func ExampleUPNP_Delete(upnp UPNP.Instance, port int) {
	upnp.DeletePortMapping(port)
}
