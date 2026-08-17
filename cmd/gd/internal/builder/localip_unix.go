//go:build linux

package builder

import (
	"net"

	"golang.org/x/sys/unix"
)

// interfaceIPv4 returns the IPv4 address assigned to the named interface, or
// nil. It asks through the SIOCGIFADDR ioctl rather than a netlink dump,
// which android permits to unprivileged processes where the dumps are not.
func interfaceIPv4(name string) net.IP {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM|unix.SOCK_CLOEXEC, 0)
	if err != nil {
		return nil
	}
	defer unix.Close(fd)
	ifr, err := unix.NewIfreq(name)
	if err != nil {
		return nil
	}
	if err := unix.IoctlIfreq(fd, unix.SIOCGIFADDR, ifr); err != nil {
		return nil
	}
	addr, err := ifr.Inet4Addr()
	if err != nil {
		return nil
	}
	return net.IP(addr)
}
