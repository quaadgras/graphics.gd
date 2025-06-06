//go:build redot

package main

import (
	"fmt"
	"runtime"
)

const (
	version   = "4.3"
	engineCmd = "redot"
)

func installEngine(gobin string) (string, error) {
	switch runtime.GOOS {
	case "android":
		return "echo", nil
	case "linux":
		fmt.Println("gd: downloading Redot v" + version + " stable for linux")
		return download(gobin, "Redot_v"+version+"-stable_linux.x86_64",
			"https://github.com/Redot-Engine/redot-engine/releases/download/redot-"+version+"-stable/Redot_v"+version+"-stable_linux.x86_64.zip")
	default:
		return "", fmt.Errorf("gd: installing "+engineCmd+" automatically for GOOS %v is not supported (please install it yourself)", runtime.GOOS)
	}
}
