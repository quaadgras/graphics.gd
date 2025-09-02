//go:build !redot && !blazium

package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
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
		var arch = "x86_64"
		switch runtime.GOARCH {
		case "arm64":
			arch = "arm64"
		}
		fmt.Println("gd: downloading Godot v" + version + " stable for linux/" + arch)
		return download(filepath.Join(gobin, engineCmd+"-"+version), "Godot_v"+version+"-stable_linux."+arch,
			"https://github.com/godotengine/godot-builds/releases/download/"+version+"-stable/Godot_v"+version+"-stable_linux."+arch+".zip")
	default:
		return "", fmt.Errorf("gd: installing "+engineCmd+" automatically for GOOS %v is not supported (please install it yourself)", runtime.GOOS)
	}
}
