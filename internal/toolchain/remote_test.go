package toolchain

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestURLScanner(t *testing.T) {
	var out bytes.Buffer
	var got string
	w := &urlScanner{out: &out, onURL: func(u string) { got = u }}
	fmt.Fprintf(w, "serving on http://192.168.1.2:8080\nscan to install: side")
	fmt.Fprintf(w, "store://install?url=http%%3A%%2F%%2F192.168.1.2%%3A8080%%2Fapp.ipa (or add via URL)\n")
	want := "sidestore://install?url=http%3A%2F%2F192.168.1.2%3A8080%2Fapp.ipa"
	if got != want {
		t.Fatalf("scanned %q, want %q", got, want)
	}
	if !strings.Contains(out.String(), "serving on") {
		t.Fatal("scanner did not tee output through")
	}
}

// TestRemoteRoundTrip drives the remote kit through the local sshd.
// Opt-in: it edits ~/.ssh/authorized_keys (and restores it), so it only
// runs with GD_HARNESS_SSH_TEST=1.
func TestRemoteRoundTrip(t *testing.T) {
	if os.Getenv("GD_HARNESS_SSH_TEST") != "1" {
		t.Skip("set GD_HARNESS_SSH_TEST=1 to run against the local sshd")
	}
	state, project := t.TempDir(), t.TempDir()
	var console bytes.Buffer
	remote, err := NewRemote(fmt.Sprintf("%s@127.0.0.1:22 %s", os.Getenv("USER"), project), state, &console)
	if err != nil {
		t.Fatal(err)
	}

	home, _ := os.UserHomeDir()
	authorized := filepath.Join(home, ".ssh", "authorized_keys")
	previous, hadPrevious := os.ReadFile(authorized)
	os.MkdirAll(filepath.Dir(authorized), 0o700)
	if err := os.WriteFile(authorized, append(previous, []byte(remote.PublicKey()+"\n")...), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if hadPrevious == nil {
			os.WriteFile(authorized, previous, 0o600)
		} else {
			os.Remove(authorized)
		}
	})

	if err := remote.WriteFile("nested/dir/hello.txt", []byte("hello over ssh\n")); err != nil {
		t.Fatal(err)
	}
	data, err := remote.ReadFile("nested/dir/hello.txt")
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "hello over ssh\n" {
		t.Fatalf("read back %q", data)
	}
	listing, err := remote.List("nested")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(listing, "dir/") {
		t.Fatalf("listing %q missing dir/", listing)
	}
	out, err := remote.Shell("echo shell-ok && pwd", time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "shell-ok") || !strings.Contains(out, project) {
		t.Fatalf("shell output %q", out)
	}
	if _, err := remote.ReadFile("../escape"); err == nil {
		t.Fatal("path escape was not refused")
	}
	if !strings.Contains(console.String(), "trusting") {
		t.Fatalf("expected first-connection trust message, got %q", console.String())
	}
}
