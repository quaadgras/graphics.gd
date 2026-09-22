package shim

import (
	"errors"
	"os"
	"os/exec"
)

// execute runs the program at path and exits with its status, windows has
// no way to replace a process.
func execute(path string, args, env []string) error {
	cmd := exec.Command(path, args[1:]...)
	cmd.Env = env
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err := cmd.Run()
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		os.Exit(exit.ExitCode())
	}
	return err
}
