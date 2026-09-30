package builder

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestDedupEditsApply checks that the hooks of the generic instantiation
// sharing patch still apply to the toolchain running the test, so that a Go
// point release that moves them fails here rather than quietly leaving
// users with the stock compiler.
func TestDedupEditsApply(t *testing.T) {
	if !strings.HasPrefix(runtime.Version(), "go1.27") {
		t.Skipf("the patch tracks go1.27, this is %v", runtime.Version())
	}
	src := make(map[string]string)
	for _, edit := range dedupEdits {
		if _, ok := src[edit.file]; !ok {
			data, err := os.ReadFile(filepath.Join(runtime.GOROOT(), "src", filepath.FromSlash(edit.file)))
			if err != nil {
				t.Skipf("no compiler sources: %v", err)
			}
			src[edit.file] = string(data)
		}
		if n := strings.Count(src[edit.file], edit.old); n != 1 {
			t.Errorf("%v: hook found %d times, want once:\n%s", edit.file, n, edit.old)
		}
		src[edit.file] = strings.Replace(src[edit.file], edit.old, edit.new, 1)
	}
}
