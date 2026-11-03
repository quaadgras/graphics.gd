package builder

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"graphics.gd/cmd/gd/internal/project"
)

// handshake_patch lets a program that embeds the engine (libgodot) with an
// extension of its own linked in name that extension's .gdextension file, in
// the LIBGODOT_MAIN_EXTENSION environment variable (graphics.gd's startup
// does). The engine then registers the program's extension under that path,
// and loads the project's extensions after it, so that the file is taken as
// loaded instead of having its libraries loaded a second time. Without the
// variable the engine is unchanged.
//
//go:embed bundled/libgodot_main_extension.patch
var handshake_patch []byte

// applyHandshake gives the engine source at src the handshake, unless it has it
// already (graphics.gd's fork does), and reports whether it has it.
func applyHandshake(src string) bool {
	if hasHandshake(src) {
		return true
	}
	if err := gitApply(src, "--check"); err != nil {
		fmt.Fprintf(os.Stderr, "gd: warning: the engine at %s does not take graphics.gd's patch for loading the Go program's own extension (%v), extension files will be moved aside while it is built with\n", src, err)
		return false
	}
	if err := gitApply(src); err != nil {
		fmt.Fprintf(os.Stderr, "gd: warning: patching the engine at %s: %v\n", src, err)
		return false
	}
	return true
}

// hasHandshake reports whether the engine source at src has the handshake.
func hasHandshake(src string) bool {
	return gitApply(src, "--reverse", "--check") == nil
}

func gitApply(src string, args ...string) error {
	cmd := exec.Command("git", append([]string{"-C", src, "apply"}, append(args, "-")...)...)
	cmd.Stdin = bytes.NewReader(handshake_patch)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}
	return nil
}

// excludeFromExport adds the file to the exclude_filter of the named export
// preset, for the length of an export: the returned func puts the filter back
// as it was. It is how graphics.gd's library.gdextension is left out of an
// export whose template has the Go program (and its extension) linked in, with
// any editor: the export then has no libraries of it to look for.
func excludeFromExport(presetName, file string) (restore func(), err error) {
	path := filepath.Join(project.GraphicsDirectory, "export_presets.cfg")
	set := func(update func(was string) string) (was string, err error) {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		lines := strings.Split(string(data), "\n")
		inPreset := false
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			switch {
			case strings.HasPrefix(trimmed, "["):
				inPreset = false
			case trimmed == "name="+strconv.Quote(presetName):
				inPreset = true
			case inPreset && strings.HasPrefix(trimmed, "exclude_filter="):
				was, err := strconv.Unquote(strings.TrimPrefix(trimmed, "exclude_filter="))
				if err != nil {
					return "", fmt.Errorf("exclude_filter of preset %q: %w", presetName, err)
				}
				now := update(was)
				if now == was {
					return was, nil
				}
				lines[i] = "exclude_filter=" + strconv.Quote(now)
				return was, os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0o644)
			}
		}
		return "", fmt.Errorf("no exclude_filter for preset %q in %s", presetName, path)
	}
	was, err := set(func(was string) string {
		for _, filter := range strings.Split(was, ",") {
			if strings.TrimSpace(filter) == file {
				return was
			}
		}
		if strings.TrimSpace(was) == "" {
			return file
		}
		return was + "," + file
	})
	if err != nil {
		return func() {}, err
	}
	return func() { set(func(string) string { return was }) }, nil
}
