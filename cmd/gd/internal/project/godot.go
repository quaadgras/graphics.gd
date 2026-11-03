package project

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

// setupGodot gets an ordinary Godot project (with no Go code of its own) at
// dir ready to export: the project directory is the graphics directory, the
// exports go to its releases directory and every export preset that gd
// builds with is added to the ones it has.
func setupGodot(dir string, build_godot func() error) error {
	Name = filepath.Base(dir)
	Directory = dir
	GraphicsDirectory = dir
	ReleasesDirectory = filepath.Join(dir, "releases")
	if err := os.MkdirAll(ReleasesDirectory, 0o755); err != nil {
		return xray.New(err)
	}
	if err := SetupFile(false, filepath.Join(ReleasesDirectory, ".gdignore"), ""); err != nil {
		return xray.New(err)
	}
	if err := SetupPresets(filepath.Join(GraphicsDirectory, "export_presets.cfg"), gdscript_export_presets_cfg, Name, AndroidSafePackageName(Name), AppleSafePackageName(Name)); err != nil {
		return xray.New(err)
	}
	// the musl build puts its files in .godot, so whether the project has
	// ever been imported is told before it does.
	_, err := os.Stat(filepath.Join(GraphicsDirectory, ".godot"))
	neverImported := os.IsNotExist(err)
	if err := build_godot(); err != nil {
		return xray.New(err)
	}
	// see Setup: on-device there is no godot to import with, unless
	// build_godot made one.
	if runtime.GOOS != "android" || tooling.Godot.Path != "" {
		if err := importProject(neverImported); err != nil {
			return xray.New(err)
		}
	}
	return nil
}

// RequireETC2ASTC turns on the import of ETC2/ASTC compressed textures, which
// the Android and iOS exports refuse to go without, in a project.godot that
// has not. It is what the editor offers to do when such a preset is added.
func RequireETC2ASTC() error {
	const (
		section = "[rendering]"
		key     = "textures/vram_compression/import_etc2_astc"
		setting = key + "=true"
	)
	path := filepath.Join(GraphicsDirectory, "project.godot")
	data, err := os.ReadFile(path)
	if err != nil {
		return xray.New(err)
	}
	lines := strings.Split(string(data), "\n")
	inSection := false
	insertAt := -1
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "[") {
			inSection = trimmed == section
			if inSection {
				insertAt = i + 1
			}
			continue
		}
		if !inSection {
			continue
		}
		if value, ok := strings.CutPrefix(trimmed, key+"="); ok {
			if value == "true" {
				return nil
			}
			lines[i] = setting
			insertAt = -2
			break
		}
	}
	switch {
	case insertAt == -1:
		for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
			lines = lines[:len(lines)-1]
		}
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, section, "", setting, "")
	case insertAt >= 0:
		// after the blank line that follows a section header, as the
		// editor writes it.
		if insertAt < len(lines) && strings.TrimSpace(lines[insertAt]) == "" {
			insertAt++
		}
		lines = append(lines[:insertAt], append([]string{setting}, lines[insertAt:]...)...)
	}
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644); err != nil {
		return xray.New(err)
	}
	fmt.Fprintf(os.Stderr, "gd: turned on %s in %s, which Android and iOS exports require\n", key, path)
	return nil
}
