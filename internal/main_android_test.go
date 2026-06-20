//go:build android

package gd_test

import (
	"os"
	"testing"
)

// TestMain emits a clean completion sentinel through the Go runtime's logd
// writer (tag "Go") after the suite finishes. On android the normal go test
// output goes through a pipe and can interleave, and the app is relaunched by
// the system after os.Exit (the suite re-runs), so `gd test` needs an
// unambiguous, non-interleaved marker to read the verdict and stop. println
// writes straight to logd, so this line is never garbled by the pipe.
func TestMain(m *testing.M) {
	code := m.Run()
	println("GDTEST_DONE", code)
	os.Exit(code)
}
