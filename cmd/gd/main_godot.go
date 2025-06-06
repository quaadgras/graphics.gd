//go:build !redot && !blazium

package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

const (
	version   = "4.4"
	engineCmd = "godot"
)

func installEngine(gobin string) (string, error) {
	switch runtime.GOOS {
	case "android":
		return "echo", nil
	case "darwin":
		if _, err := exec.LookPath("brew"); err != nil {
			return "", fmt.Errorf("gd: cannot install godot without homebrew")
		}
		fmt.Println("gd: installing Godot stable for macOS (via brew)")
		if err := exec.Command("brew", "install", "godot").Run(); err != nil {
			return "", fmt.Errorf("gd: failed to 'brew install godot': %w", err)
		}
		return "godot", nil
	case "linux":
		fmt.Println("gd: downloading Godot v" + version + " stable for linux")
		return download(gobin, "Godot_v"+version+"-stable_linux.x86_64",
			"https://github.com/godotengine/godot-builds/releases/download/"+version+"-stable/Godot_v"+version+"-stable_linux.x86_64.zip")
	default:
		return "", fmt.Errorf("gd: installing "+engineCmd+" automatically for GOOS %v is not supported (please install it yourself)", runtime.GOOS)
	}
}
