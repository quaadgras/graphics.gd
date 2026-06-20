//go:build android

package gd_test

import (
	"testing"

	"graphics.gd/classdb/JavaClassWrapper"
	"graphics.gd/variant/Object"
)

// TestJavaClassWrapperCallNoArgs reproduces the exact scenario from issue #309:
// calling a no-argument Java method through Object.Call on a JavaClassWrapper
// instance used to panic with "index out of range [0] with length 0". This only
// works on android, where JavaClassWrapper.Wrap returns a real JNI-backed class.
func TestJavaClassWrapperCallNoArgs(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var LocalDateTime = JavaClassWrapper.Wrap("java.time.LocalDateTime")
		// The crashing line from the issue: a no-arg dynamic call.
		var datetime = Object.Call(LocalDateTime, "now")
		if datetime == nil {
			t.Fatal("expected a LocalDateTime instance, got nil")
		}
	})
}
