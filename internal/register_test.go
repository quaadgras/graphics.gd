//go:build !generate

package gd_test

import (
	"fmt"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	gd "graphics.gd/internal"
	internal "graphics.gd/internal"
	"graphics.gd/variant/String"
	"graphics.gd/variant/StringName"
)

func TestRegister(t *testing.T) {
	type TestingSimpleClass struct {
		classdb.Extension[TestingSimpleClass, Node2D.Advanced]
	}
	classdb.Register[TestingSimpleClass]()

	if tag := internal.Global.ClassDB.GetClassTag(gd.NewStringName("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := internal.Global.ClassDB.GetClassTag(gd.NewStringName("TestingSimpleClass")); tag == 0 {
		t.Fail()
	}

	class := new(TestingSimpleClass)
	class_name := class.AsObject()[0].GetClass()
	if name := class_name.String(); name != "TestingSimpleClass" {
		t.Fatal(name)
	}
	class.Super().AsNode().SetName(String.New("SimpleClass"))
}

type TestingMyClassWithConstants struct {
	classdb.Extension[TestingMyClassWithConstants, Node2D.Advanced]
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
	classdb.Register[TestingMyClassWithConstants]()
}

type TestingSingleton struct {
	classdb.Extension[TestingSingleton, Node.Advanced]
}

func (TestingSingleton) Ready() {
	fmt.Println("Singleton Ready!")
}

func TestSingleton(t *testing.T) {
	classdb.Register[TestingSingleton]()
	Engine.Advanced().RegisterSingleton(StringName.New("HelloWorld"), new(TestingSingleton).AsObject())
}
