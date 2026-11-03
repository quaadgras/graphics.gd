//go:build linux

package builder

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// A second build in a directory waits for the first to let go of it.
func TestLockBuilds(t *testing.T) {
	dir := t.TempDir()
	unlock, err := lockBuilds(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := make(chan struct{})
	go func() {
		second, err := lockBuilds(dir)
		if err != nil {
			t.Error(err)
		}
		close(got)
		second()
	}()
	select {
	case <-got:
		t.Fatal("two builds held the same directory at once")
	case <-time.After(200 * time.Millisecond):
	}
	unlock()
	select {
	case <-got:
	case <-time.After(5 * time.Second):
		t.Fatal("the directory was not let go of")
	}
}

// What is held goes on being what runs, whatever is put under its name.
func TestHoldExecutable(t *testing.T) {
	self, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	program, err := os.ReadFile(self)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "editor")
	if err := os.WriteFile(path, program, 0o755); err != nil {
		t.Fatal(err)
	}
	held := holdExecutable(path)
	if held == path {
		t.Skip("no /proc to hold an executable through")
	}
	if err := os.WriteFile(path+".tmp", []byte("not a program"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(path+".tmp", path); err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command(held, "-test.run=^$").CombinedOutput(); err != nil {
		t.Fatalf("the held executable did not run: %v\n%s", err, out)
	}
	if err := exec.Command(path).Run(); err == nil {
		t.Fatal("what replaced the executable ran")
	}
}
