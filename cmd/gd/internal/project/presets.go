package project

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// presetHeader matches the [preset.N] and [preset.N.options] section headers
// of an export_presets.cfg.
var presetHeader = regexp.MustCompile(`^\[preset\.(\d+)(\.options)?\]\s*$`)

// preset is one export preset of an export_presets.cfg: both of its sections.
type preset struct {
	index int
	name  string
	text  string
}

// parsePresets splits an export_presets.cfg into its presets, in the order
// they appear.
func parsePresets(cfg string) []preset {
	var presets []preset
	current := -1
	for _, line := range strings.SplitAfter(cfg, "\n") {
		if m := presetHeader.FindStringSubmatch(strings.TrimRight(line, "\r\n")); m != nil {
			index, _ := strconv.Atoi(m[1])
			if current < 0 || presets[current].index != index {
				presets = append(presets, preset{index: index})
				current = len(presets) - 1
			}
		}
		if current < 0 {
			continue
		}
		if name, ok := strings.CutPrefix(strings.TrimSpace(line), "name="); ok && presets[current].name == "" {
			if unquoted, err := strconv.Unquote(name); err == nil {
				presets[current].name = unquoted
			}
		}
		presets[current].text += line
	}
	return presets
}

// value returns the value of the key in the preset's sections, unquoted.
func (p preset) value(key string) string {
	for _, line := range strings.Split(p.text, "\n") {
		if v, ok := strings.CutPrefix(strings.TrimSpace(line), key+"="); ok {
			if unquoted, err := strconv.Unquote(v); err == nil {
				return unquoted
			}
			return v
		}
	}
	return ""
}

// androidCovers reports whether the project's own preset p already does
// what gd's Android preset t is for: gd exports android with the preset of
// that architecture when none has gd's name (see builder.androidPresetFor),
// which is how a project's own Android preset, with its own package name and
// keystore, has always been used.
func androidCovers(p, t preset) bool {
	xrMode := func(p preset) string { // 0, no XR, when it is not set
		if mode := p.value("xr_features/xr_mode"); mode != "" {
			return mode
		}
		return "0"
	}
	if p.value("platform") != "Android" || t.value("platform") != "Android" || xrMode(p) != xrMode(t) {
		return false
	}
	for _, abi := range []string{"arm64-v8a", "x86_64", "armeabi-v7a", "x86"} {
		if t.value("architectures/"+abi) == "true" && p.value("architectures/"+abi) == "true" {
			return true
		}
	}
	return false
}

// renumber gives the preset's sections the index.
func (p preset) renumber(index int) string {
	lines := strings.SplitAfter(p.text, "\n")
	for i, line := range lines {
		if m := presetHeader.FindStringSubmatch(strings.TrimRight(line, "\r\n")); m != nil {
			lines[i] = fmt.Sprintf("[preset.%d%s]\n", index, m[2])
		}
	}
	return strings.Join(lines, "")
}

// SetupPresets writes the template (formatted with args) to the
// export_presets.cfg at path, or when there is one already, adds to it each
// preset of the template that it has none of by that name. The presets that
// are there are the project's own, and are left exactly as they are: gd
// exports with the preset named for the platform, so a project can change
// any of its options.
func SetupPresets(path, template string, args ...any) error {
	if len(args) > 0 {
		template = fmt.Sprintf(template, args...)
	}
	existing, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return os.WriteFile(path, []byte(template), 0o644)
	}
	if err != nil {
		return err
	}
	cfg := string(existing)
	own := parsePresets(cfg)
	have := map[string]bool{}
	used := map[int]bool{}
	for _, p := range own {
		have[p.name] = true
		used[p.index] = true
	}
	// Godot reads preset.0, preset.1... up to the first one missing, so
	// gd's go in the first free numbers.
	next := 0
	var added []string
	for _, p := range parsePresets(template) {
		if have[p.name] || slices.ContainsFunc(own, func(o preset) bool { return androidCovers(o, p) }) {
			continue
		}
		for used[next] {
			next++
		}
		used[next] = true
		if cfg != "" && !strings.HasSuffix(cfg, "\n\n") {
			if !strings.HasSuffix(cfg, "\n") {
				cfg += "\n"
			}
			cfg += "\n"
		}
		cfg += strings.TrimRight(p.renumber(next), "\n") + "\n"
		added = append(added, p.name)
	}
	if len(added) == 0 {
		return nil
	}
	if err := os.WriteFile(path, []byte(cfg), 0o644); err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "gd: added export presets %q to %s\n", added, path)
	return nil
}
