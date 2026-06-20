//go:build android

package startup

import (
	"bufio"
	"os"
	"syscall"
	"testing"
)

// prepareTestOutput routes the test binary's stdout/stderr to android logcat via
// the Go runtime's own logging (tag "Go"), which writes straight to logd and so
// is reliable regardless of the (link-time stub) liblog — calling
// __android_log_write directly is unreliable on some images, e.g. Waydroid.
// `gd test` reads the result back with `adb logcat -s Go:E -v raw`.
//
// Called once, before the test main is spawned, with the engine initialised.
func prepareTestOutput() {
	if !testing.Testing() {
		return
	}
	// The app's process args are android's, not test flags; reset to a clean
	// verbose invocation so the whole suite runs.
	os.Args = []string{"gdtest", "-test.v"}

	r, w, err := os.Pipe()
	if err != nil {
		return
	}
	// arm64 linux has no dup2 syscall; Dup3 covers both arches. Redirecting the
	// fds catches C-side writes too; the Go handles cover the testing package.
	_ = syscall.Dup3(int(w.Fd()), 1, 0)
	_ = syscall.Dup3(int(w.Fd()), 2, 0)
	os.Stdout = w
	os.Stderr = w

	go func() {
		sc := bufio.NewScanner(r)
		sc.Buffer(make([]byte, 64*1024), 4*1024*1024)
		for sc.Scan() {
			// println goes through the runtime's logd writer (tag "Go"), not the
			// redirected fds, so this does not loop back into the pipe.
			println(sc.Text())
		}
	}()
}
