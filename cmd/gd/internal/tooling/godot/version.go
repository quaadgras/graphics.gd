//go:build !redot && !blazium

package godot

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/tooling"
)

const (
	Version = "4.4.1"
	Name    = "godot"
)

func install(gobin string) (string, error) {
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
	case "windows":
		var arch = "win64"
		fmt.Println("gd: downloading Godot v" + Version + " stable for windows/" + arch)
		return tooling.Download(filepath.Join(gobin, Name+"-"+Version), "Godot_v"+Version+"-stable_"+arch+".exe",
			"https://github.com/godotengine/godot-builds/releases/download/"+Version+"-stable/Godot_v"+Version+"-stable_"+arch+".exe.zip")
	case "linux":
		var arch = "x86_64"
		switch runtime.GOARCH {
		case "arm64":
			arch = "arm64"
		}
		fmt.Println("gd: downloading Godot v" + Version + " stable for linux/" + arch)
		return tooling.Download(filepath.Join(gobin, Name+"-"+Version), "Godot_v"+Version+"-stable_linux."+arch,
			"https://github.com/godotengine/godot-builds/releases/download/"+Version+"-stable/Godot_v"+Version+"-stable_linux."+arch+".zip")
	default:
		return "", fmt.Errorf("gd: installing "+Name+" automatically for GOOS %v is not supported (please install it yourself)", runtime.GOOS)
	}
}
