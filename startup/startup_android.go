//go:build android

package startup

/*
#include <stdlib.h>
#include <android/log.h>
static void gd_logcat(const char *msg) { __android_log_write(ANDROID_LOG_INFO, "gdtest", msg); }
*/
import "C"

import (
	"bufio"
	"os"
	"syscall"
	"testing"
	"unsafe"
)

// prepareTestOutput routes the test binary's stdout/stderr to android logcat
// (tag "gdtest"), since android discards a process's fd 1/2 by default. `gd test`
// reads the result back with `adb logcat -s gdtest:I -v raw`. __android_log_write
// links against the (stub) liblog and resolves to the device's real liblog at
// runtime — the same path Godot's own logging uses — so this needs no
// debuggable/run-as access and works with a release export.
//
// Called once, before the test main is spawned, with the engine initialised.
func prepareTestOutput() {
	if !testing.Testing() {
		return
	}
	// The app's process args are android's, not test flags; reset to a clean
	// verbose invocation so the whole suite runs. (Per-test -run/-v passthrough
	// on android is a follow-up.)
	os.Args = []string{"gdtest", "-test.v"}

	r, w, err := os.Pipe()
	if err != nil {
		return
	}
	// arm64 linux has no dup2 syscall; Dup3 covers both arm64 and amd64.
	_ = syscall.Dup3(int(w.Fd()), 1, 0)
	_ = syscall.Dup3(int(w.Fd()), 2, 0)
	os.Stdout = w
	os.Stderr = w

	go func() {
		sc := bufio.NewScanner(r)
		sc.Buffer(make([]byte, 64*1024), 4*1024*1024)
		for sc.Scan() {
			cs := C.CString(sc.Text())
			C.gd_logcat(cs)
			C.free(unsafe.Pointer(cs))
		}
	}()
}
