package zig

import (
	"fmt"
	"os/exec"
)

var Executable string

func Assert() error {
	zig, err := exec.LookPath("zig")
	if err != nil {
		return fmt.Errorf("gd: zig is required to cross-compile, please install zig from https://ziglang.org/download/ or via your package manager: %w", err)
	}
	Executable = zig
	return nil
}
