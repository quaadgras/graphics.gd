package gdpaths

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
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

// EditorSettings returns the file Godot keeps the editor's settings in,
// for the given version of the engine (which names it by its minor version).
func EditorSettings(version string) (string, bool) {
	if parts := strings.SplitN(version, ".", 3); len(parts) >= 2 {
		version = parts[0] + "." + parts[1]
	}
	name := "editor_settings-" + version + ".tres"
	switch runtime.GOOS {
	case "linux", "android": // on-device (Termux) the editor is the linuxbsd build.
		config := os.Getenv("XDG_CONFIG_HOME")
		if config == "" {
			config = filepath.Join(os.Getenv("HOME"), ".config")
		}
		return filepath.Join(config, "godot", name), true
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Godot", name), true
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", name), true
	}
	return "", false
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
