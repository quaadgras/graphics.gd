//go:build !generate

package gd_test

import (
	"testing"

	gd "grow.graphics/gd/internal"
)

func TestStrings(t *testing.T) {
	godot := gd.NewContext(API)
	defer godot.End()

	var str = godot.String("Hello, World!")
	if str.String() != "Hello, World!" {
		t.Fail()
	}

	str.Append(godot, godot.String(" from Go!"))
	if str.String() != "Hello, World! from Go!" {
		t.Fail()
	}
}
