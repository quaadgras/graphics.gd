/*
var upnp = UPNP.new()
upnp.discover()
upnp.add_port_mapping(7777)
*/

package main

import "graphics.gd/classdb/UPNP"

func ExampleUPNP() {
	var upnp = UPNP.New()
	upnp.Discover()
	upnp.AddPortMapping(7777)
}
