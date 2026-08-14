package project

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func write(t *testing.T, dir, name, content string, mtime time.Time) {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, mtime, mtime); err != nil {
		t.Fatal(err)
	}
}

func read(t *testing.T, dir, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestSyncAndroidProject(t *testing.T) {
	local, staged := t.TempDir(), t.TempDir()
	old := time.Now().Add(-time.Hour)
	new := time.Now()

	write(t, local, "project.godot", "local project", old)          // only local
	write(t, staged, "graphics.gen.go", "generated on device", old) // only staged
	write(t, local, "main.tscn", "edited in editor", old)           // both: staged newer
	write(t, staged, "main.tscn", "edited in editor v2", new)
	write(t, local, "bullet.png", "new art", new) // both: local newer
	write(t, staged, "bullet.png", "old art", old)
	write(t, local, "icon.svg", "same", old) // both: same time
	write(t, staged, "icon.svg", "same", old)
	write(t, local, "sub/asset.png", "nested", old)             // nested, only local
	write(t, local, ".godot/uid_cache.bin", "local cache", old) // caches never sync
	write(t, staged, ".godot/imported/a.ctex", "device cache", old)
	write(t, local, "linux_amd64.so", "host build", old) // libraries never sync
	write(t, staged, "libandroid_arm64.so", "device build", old)

	if err := syncAndroidProject(local, staged); err != nil {
		t.Fatal(err)
	}

	if got := read(t, staged, "project.godot"); got != "local project" {
		t.Errorf("project.godot not copied to staged: %q", got)
	}
	if got := read(t, local, "graphics.gen.go"); got != "generated on device" {
		t.Errorf("graphics.gen.go not copied back: %q", got)
	}
	if got := read(t, local, "main.tscn"); got != "edited in editor v2" {
		t.Errorf("newer staged main.tscn should win: %q", got)
	}
	if got := read(t, staged, "bullet.png"); got != "new art" {
		t.Errorf("newer local bullet.png should win: %q", got)
	}
	if got := read(t, staged, "sub/asset.png"); got != "nested" {
		t.Errorf("nested file not copied: %q", got)
	}
	if _, err := os.Stat(filepath.Join(staged, ".godot", "uid_cache.bin")); !os.IsNotExist(err) {
		t.Error("local .godot cache leaked into staged")
	}
	if _, err := os.Stat(filepath.Join(local, ".godot", "imported", "a.ctex")); !os.IsNotExist(err) {
		t.Error("staged .godot cache leaked into local")
	}
	if _, err := os.Stat(filepath.Join(staged, "linux_amd64.so")); !os.IsNotExist(err) {
		t.Error("host library leaked into staged")
	}
	if _, err := os.Stat(filepath.Join(local, "libandroid_arm64.so")); !os.IsNotExist(err) {
		t.Error("device library leaked into local")
	}

	// A second sync must be a no-op: carried-over mtimes compare up to date.
	before, err := os.Stat(filepath.Join(local, "main.tscn"))
	if err != nil {
		t.Fatal(err)
	}
	if err := syncAndroidProject(local, staged); err != nil {
		t.Fatal(err)
	}
	after, err := os.Stat(filepath.Join(local, "main.tscn"))
	if err != nil {
		t.Fatal(err)
	}
	if !after.ModTime().Equal(before.ModTime()) {
		t.Error("second sync was not a no-op")
	}
}
