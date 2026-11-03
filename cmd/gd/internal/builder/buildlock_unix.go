//go:build unix

package builder

import (
	"fmt"
	"os"
	"syscall"
)

// lockBuilds waits for any other gd that is building into dir to finish,
// and keeps the rest out until the returned function is called: two builds
// at once write the same archive and link the same executable. It is a lock
// on the directory itself, which the kernel drops if gd dies, and is only
// held to build: what was built is run without it, so any number of runs of
// a project go on side by side. Where the directory cannot be locked the
// build goes ahead without.
func lockBuilds(dir string) (unlock func(), err error) {
	handle, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	fd := int(handle.Fd())
	err = syscall.Flock(fd, syscall.LOCK_EX|syscall.LOCK_NB)
	if err == syscall.EWOULDBLOCK {
		fmt.Fprintln(os.Stderr, "gd: waiting for another gd to finish building in", dir)
	}
	for err == syscall.EWOULDBLOCK || err == syscall.EINTR {
		err = syscall.Flock(fd, syscall.LOCK_EX)
	}
	// a filesystem with no locks to give builds as it always did.
	return func() { handle.Close() }, nil
}
