//go:build !generate

package gd_test

import (
	"fmt"
	"testing"

	"graphics.gd/defined"
	internal "graphics.gd/internal"
	"graphics.gd/objects/Engine"
	"graphics.gd/objects/Node"
	"graphics.gd/objects/Node2D"
	"graphics.gd/variant/String"
	"graphics.gd/variant/StringName"
)

func TestRegister(t *testing.T) {
	type TestingSimpleClass struct {
		defined.Object[TestingSimpleClass, Node2D.Advanced]
	}
	defined.InEditor[TestingSimpleClass]()

	if tag := internal.Global.ClassDB.GetClassTag(StringName.New("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := internal.Global.ClassDB.GetClassTag(StringName.New("TestingSimpleClass")); tag == 0 {
		t.Fail()
	}

	class := new(TestingSimpleClass)
	if name := class.AsObject().GetClass().String(); name != "TestingSimpleClass" {
		t.Fatal(name)
	}
	class.Super().AsNode().SetName(String.New("SimpleClass"))
}

type TestingMyClassWithConstants struct {
	defined.Object[TestingMyClassWithConstants, Node2D.Advanced]
}

func (*TestingMyClassWithConstants) OnRegister() {
	/*godot.Register(gd.Enum[MyClassWithConstants, int]{
	Name: "MyEnum",
	Values: map[string]int{
		"Value1": 1,
		"Value2": 2,
	},
	})*/
}

func TestRegisterConstants(t *testing.T) {
	defined.InEditor[TestingMyClassWithConstants]()
}

type TestingSingleton struct {
	defined.Object[TestingSingleton, Node.Advanced]
}

func (TestingSingleton) Ready() {
	fmt.Println("Singleton Ready!")
}

func TestSingleton(t *testing.T) {
	defined.InEditor[TestingSingleton]()
	Engine.Advanced().RegisterSingleton(StringName.New("HelloWorld"), new(TestingSingleton).AsObject())
}
