//go:build !generate

package gd_test

import (
	"fmt"
	"testing"

	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
	"grow.graphics/gd/object/Engine"
	"grow.graphics/gd/object/Node"
	"grow.graphics/gd/object/Node2D"
)

func TestRegister(t *testing.T) {
	godot := internal.NewLifetime(API)
	defer godot.End()

	type SimpleClass struct {
		gd.Class[SimpleClass, Node2D.Expert]
	}

	gd.Register[SimpleClass](godot)

	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("SimpleClass")); tag == 0 {
		t.Fail()
	}

	class := gd.New[SimpleClass](godot)
	if class.AsObject().GetClass(godot).String() != "SimpleClass" {
		t.Fail()
	}
	class.Super().AsNode().SetName(godot.String("SimpleClass"))
}

type MyClassWithConstants struct {
	gd.Class[MyClassWithConstants, Node2D.Expert]
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
	godot := internal.NewLifetime(API)
	defer godot.End()

	gd.Register[MyClassWithConstants](godot)
}

type Singleton struct {
	gd.Class[Singleton, Node.Expert]
}

func (Singleton) Ready() {
	fmt.Println("Singleton Ready!")
}

func TestSingleton(t *testing.T) {
	godot := internal.NewLifetime(API)
	defer godot.End()

	gd.Register[Singleton](godot)

	Engine.Expert(gd.Engine(godot)).RegisterSingleton(godot.StringName("HelloWorld"), gd.New[Singleton](godot).AsObject())
}
