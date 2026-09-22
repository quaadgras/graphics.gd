//go:build !windows

package shim

import "syscall"

// execute replaces gd with the program at path.
func execute(path string, args, env []string) error {
	return syscall.Exec(path, args, env)
}
