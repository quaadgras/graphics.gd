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

	class := gd.Create(godot, new(SimpleClass))
	if class.AsObject().GetClass(godot).String() != "SimpleClass" {
		t.Fail()
	}
	class.Super().AsNode().SetName(godot.String("SimpleClass"))
}

type MyClassWithConstants struct {
	gd.Class[MyClassWithConstants, gd.Node2D]
}

func (*MyClassWithConstants) OnRegister(godot gd.Lifetime) {
	godot.Register(gd.Enum[MyClassWithConstants, int]{
		Name: "MyEnum",
		Values: map[string]int{
			"Value1": 1,
			"Value2": 2,
		},
	})
}

func TestRegisterConstants(t *testing.T) {
	godot := internal.NewContext(API)
	defer godot.End()

	gd.Register[MyClassWithConstants](godot)
}
