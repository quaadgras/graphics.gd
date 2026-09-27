// Package bakedenv sets the environment variables `gd run` baked into the
// library when it built it on an Android device (Termux), where the app is
// launched by the system and inherits nothing from the shell that ran gd.
//
// It is imported by internal/gdextension, which every package touching the
// engine depends on, so the variables are set before any package-level
// variable of the program reads them.
package bakedenv

import (
	"encoding/base64"
	"os"
	"strings"
)

// env is set by the linker (-X graphics.gd/internal/bakedenv.env=...): the
// variables, NAME=value, separated by NUL and base64 encoded.
var env string

func init() {
	if env == "" {
		return
	}
	raw, err := base64.StdEncoding.DecodeString(env)
	if err != nil {
		return
	}
	for _, kv := range strings.Split(string(raw), "\x00") {
		if name, value, ok := strings.Cut(kv, "="); ok && name != "" {
			os.Setenv(name, value)
		}
	}
}
