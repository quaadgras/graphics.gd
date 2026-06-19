//go:build android

package startup

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"

	OSClass "graphics.gd/classdb/OS"
)

// TestOutputFile is the basename, inside the app's user-data directory, that
// `gd test` writes its output to on android. The CI harness reads it back with
//
//	adb shell run-as <pkg> cat files/<TestOutputFile>
//
// A file is used rather than logcat because android discards a process's fd 1/2
// by default, and the amd64 build links liblog as a no-op stub (resolved to the
// device's real liblog only at runtime) — so logcat is an unreliable channel
// for the test result. A file in the app's private dir is not.
const TestOutputFile = "gdtest.log"

// prepareTestOutput redirects Go's stdout and stderr to a file under the app's
// user-data directory so `go test` output (the PASS/FAIL lines) survives on
// android. It is called once, just before the test main is spawned, with the
// engine already initialised — so OS.GetUserDataDir() is valid here.
func prepareTestOutput() {
	if !testing.Testing() {
		return
	}
	f, err := os.Create(filepath.Join(OSClass.GetUserDataDir(), TestOutputFile))
	if err != nil {
		return
	}
	// The app's process args are android's, not test flags, and `testing`'s flag
	// parser would reject unknown ones. Reset to a clean, verbose invocation so
	// the whole suite runs. (Per-test -run/-v passthrough is a follow-up — the
	// args would have to be injected here from an intent extra or a pushed file.)
	if len(os.Args) > 0 {
		os.Args = []string{os.Args[0], "-test.v"}
	} else {
		os.Args = []string{"gdtest", "-test.v"}
	}
	// Redirect at the fd level so writes to fd 1/2 from any source (cgo, the C
	// side) also land in the file, and reassign the Go handles so the testing
	// package — which references os.Stdout/os.Stderr — writes there too.
	// arm64 linux has no dup2 syscall; Dup3 covers both arm64 and amd64.
	_ = syscall.Dup3(int(f.Fd()), 1, 0)
	_ = syscall.Dup3(int(f.Fd()), 2, 0)
	os.Stdout = f
	os.Stderr = f
}
