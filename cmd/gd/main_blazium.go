//go:build blazium

package main

import (
	"fmt"
	"runtime"
)

const (
	version   = "0.5.68"
	engineCmd = "blazium"
)

func installEngine(gobin string) (string, error) {
	switch runtime.GOOS {
	case "android":
		return "echo", nil
	case "linux":
		fmt.Println("gd: downloading Blazium v" + version + " stable for linux")
		return download(gobin, "BlaziumEditor_v"+version+"_linux.x86_64",
			"https://cdn.blazium.app/release/"+version+"/BlaziumEditor_v"+version+"_linux.x86_64.zip")
	default:
		return "", fmt.Errorf("gd: installing "+engineCmd+" automatically for GOOS %v is not supported (please install it yourself)", runtime.GOOS)
	}
}
