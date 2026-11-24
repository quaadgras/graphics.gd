package tooling

import (
	"archive/zip"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// zipOf writes a zip of files (path → content) below a top-level directory,
// as a tool's release archive is laid out.
func zipOf(t *testing.T, path string, files map[string]string) {
	t.Helper()
	out, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	w := zip.NewWriter(out)
	for name, content := range files {
		f, err := w.Create("zig-x86_64-linux-0.17.1/" + name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := f.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	if err := out.Close(); err != nil {
		t.Fatal(err)
	}
}

// tree returns every file below dir (relative path → content).
func tree(t *testing.T, dir string) map[string]string {
	t.Helper()
	files := map[string]string{}
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := os.ReadFile(path)
		rel, _ := filepath.Rel(dir, path)
		files[filepath.ToSlash(rel)] = string(data)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func write(t *testing.T, dir string, files map[string]string) {
	t.Helper()
	for name, content := range files {
		path := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
}

func same(t *testing.T, got, want map[string]string) {
	t.Helper()
	var keys []string
	for k := range got {
		keys = append(keys, k)
	}
	for k := range want {
		if _, ok := got[k]; !ok {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		if got[k] != want[k] {
			t.Errorf("%s = %q, want %q", k, got[k], want[k])
		}
	}
}

var (
	oldZig = map[string]string{"zig": "0.15.2", "lib/std.zig": "old std", "lib/libc/musl/malloc.c": "mallocng", "gd": "gd itself"}
	newZig = map[string]string{"zig": "0.17.1", "lib/std.zig": "new std", "lib/c/malloc.zig": "smp"}
)

func setup(t *testing.T) (dir, archive, previous string) {
	dir = t.TempDir()
	write(t, dir, oldZig)
	archive = filepath.Join(dir, "zig.0.17.1.download")
	zipOf(t, archive, newZig)
	return dir, archive, filepath.Join(dir, "zig.previous")
}

func TestInstallArchiveReplacesOlderVersion(t *testing.T) {
	dir, archive, previous := setup(t)
	if err := installArchive(archive, dir, "zip", previous); err != nil {
		t.Fatal(err)
	}
	os.Remove(archive)
	// the older lib directory goes as a whole, other tools stay.
	same(t, tree(t, dir), map[string]string{"zig": "0.17.1", "lib/std.zig": "new std", "lib/c/malloc.zig": "smp", "gd": "gd itself"})
}

func TestInstallArchiveRestoresOlderVersionWhenAMoveFails(t *testing.T) {
	dir, archive, previous := setup(t)
	defer func() { rename = os.Rename }()
	moves := 0
	rename = func(from, to string) error {
		// fail once the first new entry is in place and the second is
		// about to be.
		if moves++; moves == 4 {
			return errors.New("injected failure")
		}
		return os.Rename(from, to)
	}
	err := installArchive(archive, dir, "zip", previous)
	if err == nil || !strings.Contains(err.Error(), "injected failure") {
		t.Fatalf("err = %v, want the injected failure", err)
	}
	os.Remove(archive)
	same(t, tree(t, dir), oldZig)
}

func TestInstallArchiveKeepsOlderVersionWhenTheArchiveIsBroken(t *testing.T) {
	dir, archive, previous := setup(t)
	if err := os.WriteFile(archive, []byte("not a zip"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := installArchive(archive, dir, "zip", previous); err == nil {
		t.Fatal("a broken archive installed")
	}
	if _, err := os.Stat(archive); !os.IsNotExist(err) {
		t.Errorf("the broken archive was kept, so it would never be downloaded again")
	}
	same(t, tree(t, dir), oldZig)
}

func TestInstallArchiveRemovesWhatAnEarlierUpdateLeft(t *testing.T) {
	dir, archive, previous := setup(t)
	write(t, dir, map[string]string{
		"zig.previous-123/zig":    "0.14.0",
		"zig.0.17.1.staging/half": "extracted",
	})
	if err := installArchive(archive, dir, "zip", previous); err != nil {
		t.Fatal(err)
	}
	os.Remove(archive)
	same(t, tree(t, dir), map[string]string{"zig": "0.17.1", "lib/std.zig": "new std", "lib/c/malloc.zig": "smp", "gd": "gd itself"})
}
