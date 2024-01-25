//go:build !generate

package gd_test

import (
	"testing"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
)

func TestRegister(t *testing.T) {
	godot := internal.NewContext(API)
	defer godot.End()

	type SimpleClass struct {
		gd.Class[SimpleClass, gd.Node2D]
	}

	gd.Register[SimpleClass](godot)

	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("SimpleClass")); tag == 0 {
		t.Fail()
	}
}
