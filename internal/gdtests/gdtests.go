package gdtests

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"grow.graphics/gd/gdmaths/Float"
)

func That(t *testing.T, value, expectation any) {
	t.Helper()
	if !reflect.DeepEqual(value, expectation) {
		t.Fatalf("expected %v, got %v", expectation, value)
	}
}

func Print(t *testing.T, output string, value any) {
	t.Helper()
	printed := fmt.Sprint(value)
	if slice, ok := value.([]string); ok {
		printed = "["
		for i, v := range slice {
			if i > 0 {
				printed += ", "
			}
			printed += fmt.Sprintf("%q", v)
		}
		printed += "]"
	}
	if _, ok := value.(Float.X); ok {
		if !strings.Contains(printed, ".") {
			printed += ".0"
		}
	}
	if printed != output {
		t.Fatalf("expected %q, got %q", output, printed)
	}
}
