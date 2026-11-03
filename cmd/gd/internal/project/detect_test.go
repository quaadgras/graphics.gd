package project

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestDetect(t *testing.T) {
	const (
		graphicsMod = "module game\n\ngo 1.27\n\nrequire graphics.gd v0.0.0\n"
		plainMod    = "module tool\n\ngo 1.27\n"
		godot       = "config_version=5\n"
		startupMain = "package main\n\nimport \"graphics.gd/startup\"\n\nfunc main() { startup.Scene() }\n"
		plainMain   = "package main\n\nfunc main() {}\n"
	)
	type files map[string]string
	for _, tt := range []struct {
		name  string
		files files
		cwd   string
		args  []string
		kind  Kind
		root  string
	}{
		{"graphics.gd module", files{"go.mod": graphicsMod}, "", nil, Graphics, ""},
		{"inside its graphics directory", files{"go.mod": graphicsMod, "graphics/project.godot": godot}, "graphics", nil, Graphics, ""},
		{"one of many in a module", files{"go.mod": graphicsMod, "2d/pong/graphics/project.godot": godot, "2d/pong/main.go": startupMain}, "2d/pong", nil, Graphics, "2d/pong"},
		{"inside one of many", files{"go.mod": graphicsMod, "2d/pong/graphics/project.godot": godot}, "2d/pong/graphics", nil, Graphics, "2d/pong"},
		{"plain Go repo with a graphics package", files{"go.mod": plainMod, "graphics/project.godot": godot, "graphics/draw.go": plainMain}, "", nil, Go, ""},
		{"graphics.gd itself", files{"go.mod": "module graphics.gd\n"}, "", nil, Graphics, ""},
		{"plain Go module", files{"go.mod": plainMod, "main.go": plainMain}, "", nil, Go, ""},
		{"plain Go package below", files{"go.mod": plainMod, "cmd/x/main.go": plainMain}, "cmd/x", nil, Go, ""},
		{"Godot project", files{"project.godot": godot}, "", nil, Godot, ""},
		{"inside a Godot project", files{"project.godot": godot, "scenes/a.tscn": ""}, "scenes", nil, Godot, ""},
		{"Godot project beside a plain go.mod", files{"go.mod": plainMod, "project.godot": godot}, "", nil, Godot, ""},
		{"Godot project in a Go repo", files{"go.mod": plainMod, "game/project.godot": godot}, "game", nil, Godot, "game"},
		{"Godot project in a graphics.gd repo", files{"go.mod": graphicsMod, "demo/project.godot": godot}, "demo", nil, Godot, "demo"},
		{"Go tool in a Godot project", files{"project.godot": godot, "tools/go.mod": plainMod}, "tools", nil, Go, "tools"},
		{"no go.mod, imports startup", files{"main.go": startupMain}, "", nil, Graphics, ""},
		{"no go.mod, plain Go", files{"main.go": plainMain}, "", nil, Go, ""},
		{"no go.mod, run a file that imports startup", files{"main.go": plainMain, "game.go": startupMain}, "", []string{"run", "game.go"}, Graphics, ""},
		{"no go.mod, run a plain file", files{"main.go": startupMain, "hello.go": plainMain}, "", []string{"run", "hello.go"}, Go, ""},
		{"no go.mod, run an absolute path", files{"main.go": startupMain, "sub/x": ""}, "sub", []string{"run", "/abs/main.go"}, Graphics, "sub"},
		{"nothing", files{"README": ""}, "", nil, Nothing, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			for name, content := range tt.files {
				write(t, dir, name, content, time.Now())
			}
			args := tt.args
			for i, arg := range args {
				if strings.HasPrefix(arg, "/abs/") {
					args = append([]string(nil), args...)
					args[i] = filepath.Join(dir, strings.TrimPrefix(arg, "/abs/"))
				}
			}
			kind, root := Detect(filepath.Join(dir, tt.cwd), args)
			if kind != tt.kind {
				t.Errorf("kind = %v, want %v", kind, tt.kind)
			}
			if want := filepath.Join(dir, tt.root); root != want && tt.kind != Nothing {
				t.Errorf("root = %v, want %v", root, want)
			}
		})
	}
}

func TestSetupPresets(t *testing.T) {
	path := filepath.Join(t.TempDir(), "export_presets.cfg")
	template := "[preset.0]\n\nname=\"Linux x86_64\"\nexport_path=\"./releases/%[1]v\"\n\n[preset.0.options]\n\na=1\n\n" +
		"[preset.1]\n\nname=\"Web\"\n\n[preset.1.options]\n\nb=2\n"
	// written as it is when the project has none.
	if err := SetupPresets(path, template, "game"); err != nil {
		t.Fatal(err)
	}
	if got := read(t, filepath.Dir(path), "export_presets.cfg"); !strings.Contains(got, `"./releases/game"`) || len(parsePresets(got)) != 2 {
		t.Fatalf("new file:\n%s", got)
	}
	// the project's own presets are kept as they are, gd's missing ones added after them.
	own := "[preset.0]\n\nname=\"Web\"\ncustom=\"mine\"\n\n[preset.0.options]\n\nb=3\n\n" +
		"[preset.1]\n\nname=\"Windows Desktop\"\n\n[preset.1.options]\n\nc=4\n"
	if err := os.WriteFile(path, []byte(own), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := SetupPresets(path, template, "game"); err != nil {
		t.Fatal(err)
	}
	got := read(t, filepath.Dir(path), "export_presets.cfg")
	if !strings.HasPrefix(got, own) {
		t.Fatalf("own presets changed:\n%s", got)
	}
	presets := parsePresets(got)
	if len(presets) != 3 || presets[2].name != "Linux x86_64" || presets[2].index != 2 ||
		!strings.Contains(presets[2].text, "[preset.2.options]") || !strings.Contains(presets[2].text, "a=1") {
		t.Fatalf("merged:\n%s", got)
	}
	// and nothing is written when there is nothing to add.
	before, _ := os.Stat(path)
	time.Sleep(10 * time.Millisecond)
	if err := SetupPresets(path, template, "game"); err != nil {
		t.Fatal(err)
	}
	if after, _ := os.Stat(path); !after.ModTime().Equal(before.ModTime()) {
		t.Fatal("rewrote a complete export_presets.cfg")
	}
}

// The presets gd builds with must be in both templates, and each template
// numbered as Godot reads it: from 0, with no gaps.
func TestSetupPresetsNumbering(t *testing.T) {
	path := filepath.Join(t.TempDir(), "export_presets.cfg")
	template := "[preset.0]\n\nname=\"A\"\n\n[preset.0.options]\n\n[preset.1]\n\nname=\"B\"\n\n[preset.1.options]\n"
	// Godot stops at the first missing number, so gd's fill the gaps.
	own := "[preset.0]\n\nname=\"Mine\"\n\n[preset.0.options]\n\n[preset.2]\n\nname=\"Other\"\n\n[preset.2.options]\n"
	if err := os.WriteFile(path, []byte(own), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := SetupPresets(path, template); err != nil {
		t.Fatal(err)
	}
	got := map[string]int{}
	for _, p := range parsePresets(read(t, filepath.Dir(path), "export_presets.cfg")) {
		got[p.name] = p.index
	}
	if got["Mine"] != 0 || got["A"] != 1 || got["Other"] != 2 || got["B"] != 3 {
		t.Fatalf("numbered %v", got)
	}
}

// A project's own Android preset is the one gd has always exported android
// with (for its package name and keystore), so gd's is not added beside it.
func TestSetupPresetsAndroid(t *testing.T) {
	path := filepath.Join(t.TempDir(), "export_presets.cfg")
	own := "[preset.0]\n\nname=\"Android\"\nplatform=\"Android\"\n\n[preset.0.options]\n\narchitectures/arm64-v8a=true\npackage/unique_name=\"com.mine.game\"\n"
	if err := os.WriteFile(path, []byte(own), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := SetupPresets(path, export_presets_cfg, "game", "game", "game"); err != nil {
		t.Fatal(err)
	}
	names := map[string]bool{}
	for _, p := range parsePresets(read(t, filepath.Dir(path), "export_presets.cfg")) {
		names[p.name] = true
	}
	if names["Android arm64-v8a"] {
		t.Error("added Android arm64-v8a beside the project's own arm64 Android preset")
	}
	for _, name := range []string{"Android x86_64", "Meta Quest", "Linux x86_64"} {
		if !names[name] {
			t.Errorf("did not add %q", name)
		}
	}
}

func TestPresetTemplates(t *testing.T) {
	names := func(cfg string) map[string]bool {
		seen := map[string]bool{}
		for i, p := range parsePresets(cfg) {
			if p.index != i {
				t.Errorf("preset %q is number %d, want %d", p.name, p.index, i)
			}
			seen[p.name] = true
		}
		return seen
	}
	gd, godot := names(export_presets_cfg), names(gdscript_export_presets_cfg)
	for _, name := range []string{"Windows x86_64", "Windows arm64", "macOS", "iOS", "Android arm64-v8a", "Android x86_64", "Linux x86_64", "Linux arm64", "Web", "Musl x86_64", "Musl arm64"} {
		if !gd[name] || !godot[name] {
			t.Errorf("preset %q missing (graphics.gd: %v, Godot: %v)", name, gd[name], godot[name])
		}
	}
}

func TestRequireETC2ASTC(t *testing.T) {
	for _, tt := range []struct{ name, in, want string }{
		{"no rendering section", "config_version=5\n\n[application]\n\nconfig/name=\"x\"\n",
			"config_version=5\n\n[application]\n\nconfig/name=\"x\"\n\n[rendering]\n\ntextures/vram_compression/import_etc2_astc=true\n"},
		{"rendering section", "[rendering]\n\nrenderer/rendering_method=\"mobile\"\n",
			"[rendering]\n\ntextures/vram_compression/import_etc2_astc=true\nrenderer/rendering_method=\"mobile\"\n"},
		{"turned off", "[rendering]\n\ntextures/vram_compression/import_etc2_astc=false\n",
			"[rendering]\n\ntextures/vram_compression/import_etc2_astc=true\n"},
		{"already on", "[rendering]\n\ntextures/vram_compression/import_etc2_astc=true\n",
			"[rendering]\n\ntextures/vram_compression/import_etc2_astc=true\n"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			GraphicsDirectory = t.TempDir()
			write(t, GraphicsDirectory, "project.godot", tt.in, time.Now())
			if err := RequireETC2ASTC(); err != nil {
				t.Fatal(err)
			}
			if got := read(t, GraphicsDirectory, "project.godot"); got != tt.want {
				t.Errorf("got\n%s\nwant\n%s", got, tt.want)
			}
		})
	}
}
