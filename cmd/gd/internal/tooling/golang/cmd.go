package golang

import (
	"os"
	"os/exec"

	"runtime.link/api"
	"runtime.link/api/cmdl"
)

var CMD = api.Import[struct {
	Env struct {
		GOROOT func() string `cmdl:"env GOROOT"`
	}
}](cmdl.API, "go", nil)

// Forwarded forwards arguments to the `go` command.
func Forwarded(args ...string) error {
	golang := exec.Command("go", args...)
	golang.Stderr = os.Stderr
	golang.Stdout = os.Stdout
	golang.Stdin = os.Stdin
	return golang.Run()
}

// Build arguments.
func Build(suffix []string, args ...string) error {
	return Forwarded(append(append([]string{"build", "-ldflags=-s -w"}, args...), suffix...)...)
}

// Test arguments
func Test(suffix []string, args ...string) error {
	return Forwarded(append(append([]string{"test"}, args...), suffix...)...)
}

// Run arguments.
func Run(args ...string) error {
	return Forwarded(append([]string{"run"}, args...)...)
}
