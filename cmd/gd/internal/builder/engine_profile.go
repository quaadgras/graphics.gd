package builder

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"
)

// engineClass as the engine's documentation describes it.
type engineClass struct {
	Inherits string
	Core     bool // as opposed to editor only.
}

// engineClasses reads the classes of the engine at src, along with the
// dependencies its ClassDB declares between them (ADD_CLASS_DEPENDENCY).
func engineClasses(src string) (classes map[string]engineClass, dependencies map[string][]string, err error) {
	classes, dependencies = map[string]engineClass{}, map[string][]string{}
	header := regexp.MustCompile(`<class\s+name="([^"]+)"(?:\s+inherits="([^"]+)")?[^>]*?(?:\s+api_type="([^"]+)")?[^>]*>`)
	docs, err := filepath.Glob(filepath.Join(src, "doc", "classes", "*.xml"))
	if err != nil {
		return nil, nil, err
	}
	modules, _ := filepath.Glob(filepath.Join(src, "modules", "*", "doc_classes", "*.xml"))
	platforms, _ := filepath.Glob(filepath.Join(src, "platform", "*", "doc_classes", "*.xml"))
	for _, path := range slices.Concat(docs, modules, platforms) {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, nil, err
		}
		match := header.FindSubmatch(data)
		if match == nil {
			continue
		}
		classes[string(match[1])] = engineClass{
			Inherits: string(match[2]),
			Core:     string(match[3]) != "editor",
		}
	}
	if len(classes) == 0 {
		return nil, nil, fmt.Errorf("gd: no class documentation in %s, is this a Godot source tree?", src)
	}
	binding := regexp.MustCompile(`(?m)^\s*void\s+([A-Za-z0-9_]+)::_bind_methods\(\)|ADD_CLASS_DEPENDENCY\("([^"]+)"\)`)
	err = filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if name := d.Name(); name == "thirdparty" || name == "bin" || name == ".git" || name == "editor" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".cpp" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if !strings.Contains(string(data), "ADD_CLASS_DEPENDENCY") {
			return nil
		}
		var class string
		for _, match := range binding.FindAllSubmatch(data, -1) {
			if match[1] != nil {
				class = string(match[1])
			} else if class != "" {
				dependencies[class] = append(dependencies[class], string(match[2]))
			}
		}
		return nil
	})
	return classes, dependencies, err
}

// classesUsed by the project: whatever its Go code imports from classdb,
// along with every name of a class that appears in its files (scenes,
// resources, imports, scripts, and Go sources for classes named by their
// string). Over-inclusion only keeps a class that could have been dropped.
func classesUsed(classes map[string]engineClass) (map[string]bool, error) {
	used := map[string]bool{}
	var dirs []string
	if project.IncludesGo {
		goexe, err := tooling.Go.Lookup()
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(goexe, "list", "-deps", "-f", "{{.ImportPath}} {{.Dir}} {{if .Module}}{{.Module.Main}}{{end}}", "./...")
		cmd.Dir = project.Directory
		cmd.Stderr = os.Stderr
		list, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("gd: listing the project's dependencies: %w", err)
		}
		for line := range strings.SplitSeq(string(list), "\n") {
			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue
			}
			if class, ok := strings.CutPrefix(fields[0], "graphics.gd/classdb/"); ok {
				if _, ok := classes[class]; ok {
					used[class] = true
				}
			}
			if len(fields) == 3 && fields[2] == "true" {
				dirs = append(dirs, fields[1])
			}
		}
	}
	// build outputs (which land in the graphics directory) name whatever
	// they happen to be linked with, not what the project uses.
	outputs := map[string]bool{".so": true, ".dll": true, ".dylib": true, ".exe": true, ".a": true, ".o": true,
		".wasm": true, ".pck": true, ".zip": true, ".apk": true, ".aab": true, ".ipa": true, ".editor": true}
	identifier := regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)
	scan := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if name := d.Name(); (strings.HasPrefix(name, ".") && path != project.GraphicsDirectory) || strings.HasSuffix(name, ".xcframework") {
				return filepath.SkipDir
			}
			return nil
		}
		if outputs[filepath.Ext(path)] || d.Name() == "library_documentation.xml" {
			return nil
		}
		info, err := d.Info()
		if err != nil || info.Size() > 64<<20 {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, name := range identifier.FindAll(data, -1) {
			if _, ok := classes[string(name)]; ok {
				used[string(name)] = true
			}
		}
		return nil
	}
	if err := filepath.WalkDir(project.GraphicsDirectory, scan); err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".go" {
				if err := scan(filepath.Join(dir, entry.Name()), entry, nil); err != nil {
					return nil, err
				}
			}
		}
	}
	return used, nil
}

// disabledClasses returns the classes of the engine that a build for the
// project can leave out, following the editor's build profile manager
// (EditorBuildProfileManager::_detect_from_project).
func disabledClasses(classes map[string]engineClass, dependencies map[string][]string, used map[string]bool, forced []string) (disabled []string, dropped int) {
	// HACK: Some classes are included due to creating clashes with unrelated
	// when disabled. Until that is fixed, they need to always be enabled.
	hardcoded := []string{"Font", "InputEvent", "ShaderInclude", "StyleBox", "Window",
		// gd's own: the Go code is loaded through a GDExtension resource,
		// its use is not something the project's files reveal.
		"GDExtension"}
	descends := func(class, ancestor string) bool {
		for class != "" {
			if class == ancestor {
				return true
			}
			class = classes[class].Inherits
		}
		return false
	}
	keep := map[string]bool{}
	for class := range used {
		keep[class] = true
	}
	for _, class := range forced {
		if class = strings.TrimSpace(class); class != "" {
			keep[class] = true
		}
	}
	for _, class := range hardcoded {
		keep[class] = true
		for other := range classes {
			if descends(other, class) {
				keep[other] = true
			}
		}
	}
	// a kept class needs its dependencies and its ancestors.
	for {
		var more []string
		for class := range keep {
			for _, other := range slices.Concat(dependencies[class], []string{classes[class].Inherits}) {
				if other != "" && !keep[other] {
					more = append(more, other)
				}
			}
		}
		if len(more) == 0 {
			break
		}
		for _, class := range more {
			keep[class] = true
		}
	}
	for name, class := range classes {
		if !class.Core || keep[name] || !(descends(name, "Node") || descends(name, "Resource")) {
			continue
		}
		dropped++
		// only the root of a disabled subtree is listed, the way the editor
		// does it: with a parent that is disabled too, this one is redundant.
		if class.Inherits == "" || keep[class.Inherits] {
			disabled = append(disabled, name)
		}
	}
	slices.Sort(disabled)
	return disabled, dropped
}

// engineProfiles already written this run, by engine source directory.
var engineProfiles = map[string]string{}

// profile writes the build profile that leaves out the classes of the
// engine at src that the project does not use, returning its path.
func (custom engine) profile(src string) (string, error) {
	if path, ok := engineProfiles[src]; ok {
		return path, nil
	}
	classes, dependencies, err := engineClasses(src)
	if err != nil {
		return "", err
	}
	used, err := classesUsed(classes)
	if err != nil {
		return "", err
	}
	disabled, dropped := disabledClasses(classes, dependencies, used, custom.Classes)
	data, err := json.MarshalIndent(map[string]any{
		"type":                   "build_profile",
		"disabled_classes":       disabled,
		"disabled_build_options": map[string]any{},
	}, "", "\t")
	if err != nil {
		return "", err
	}
	path := filepath.Join(project.GraphicsDirectory, ".godot", "engine.build")
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return "", err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", err
	}
	fmt.Printf("gd: leaving %d of the engine's %d classes out of the build, as unused by the project (%s)\n", dropped, len(classes), path)
	engineProfiles[src] = path
	return path, nil
}
