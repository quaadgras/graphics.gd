/*
{
	"index": "1", # Interface index.
	"name": "eth0", # Interface name.
	"friendly": "Ethernet One", # A friendly name (might be empty).
	"addresses": ["192.168.1.101"], # An array of IP addresses associated to this interface.
}
*/

package main

import (
	"net/netip"

	"graphics.gd/classdb/IP"
)

func IP_GetLocalInterfaces() {
	example := IP.LocalInterface{
		Index:     "1",
		Name:      "eth0",
		Friendly:  "Ethernet One",
		Addresses: []netip.Addr{netip.MustParseAddr("192.168.1.101")},
	}
	_ = example
}
