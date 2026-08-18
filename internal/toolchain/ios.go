//go:build ios

package toolchain

import (
	"errors"
	"time"
)

// iOS forbids process creation, so this tier must run everything
// in-process: compiler.gd for Go compilation, lld's library entry point
// for the Mach-O link, an ios_system-style dispatcher for shell
// commands, and a loopback HTTP handoff to SideStore for install.
// See the roadmap in the Readme. Stubbed until those land.

var errNoExec = errors.New("in-process toolchain is not implemented on iOS yet (see Readme roadmap)")

func (t *Toolchain) Shell(command string, timeout time.Duration) (string, error) {
	return "", errNoExec
}

func (t *Toolchain) GD(verb string, goos string) error {
	return errNoExec
}
