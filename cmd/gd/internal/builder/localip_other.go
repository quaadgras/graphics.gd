//go:build !linux

package builder

import "net"

// interfaceIPv4 is the android escape hatch for a denied netlink interface
// dump (see localip_unix.go); on other hosts net.InterfaceAddrs works, so
// the fallback path that calls this is never reached.
func interfaceIPv4(string) net.IP { return nil }
