//go:build cgo && !android

package startup

// prepareTestOutput is a no-op on platforms whose stdout/stderr already reach
// somewhere the `gd test` harness can observe (a terminal, or logcat via the Go
// runtime). The android build overrides it — see startup_android.go.
func prepareTestOutput() {}
