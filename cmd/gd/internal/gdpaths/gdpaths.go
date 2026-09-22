package gdpaths

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	Bin string
	Lib string
)

func init() {
	GDPATH := os.Getenv("GDPATH")
	whoami, err := user.Current()
	if GDPATH == "" && err == nil {
		GDPATH = filepath.Join(whoami.HomeDir, "gd")
	}
	Bin = filepath.Join(GDPATH, "bin")
	Lib = filepath.Join(GDPATH, "lib")
}

// ExportTemplates returns the directory where Godot keeps the export
// templates for the given version of the engine.
func ExportTemplates(version string) (string, bool) {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "export_templates", version+".stable"), true
	case "android":
		// On-device (Termux) the exporting editor is the static musl
		// (linuxbsd) build, which reads the XDG data dir under Termux's HOME.
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "export_templates", version+".stable"), true
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Godot", "export_templates", version+".stable"), true
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "export_templates", version+".stable"), true
	}
	return "", false
}
