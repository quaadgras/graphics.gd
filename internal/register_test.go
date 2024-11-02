//go:build !generate

package gd_test

import (
	"fmt"
	"testing"

	"grow.graphics/gd/gdextension"
	internal "grow.graphics/gd/internal"
	"grow.graphics/gd/objects/Engine"
	"grow.graphics/gd/objects/Node"
	"grow.graphics/gd/objects/Node2D"
	"grow.graphics/gd/variant/String"
)

func TestRegister(t *testing.T) {
	type SimpleClass struct {
		gdextension.Class[SimpleClass, Node2D.Advanced]
	}
	gdextension.Register[SimpleClass]()

	if tag := internal.Global.ClassDB.GetClassTag(internal.NewStringName("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := internal.Global.ClassDB.GetClassTag(internal.NewStringName("SimpleClass")); tag == 0 {
		t.Fail()
	}

	class := new(SimpleClass)
	if class.AsObject().GetClass().String() != "SimpleClass" {
		t.Fail()
	}
	class.Super().AsNode().SetName(String.New("SimpleClass"))
}

type MyClassWithConstants struct {
	gdextension.Class[MyClassWithConstants, Node2D.Advanced]
}

func (*MyClassWithConstants) OnRegister() {
	/*godot.Register(gd.Enum[MyClassWithConstants, int]{
	Name: "MyEnum",
	Values: map[string]int{
		"Value1": 1,
		"Value2": 2,
	},
	})*/
}

func TestRegisterConstants(t *testing.T) {
	gdextension.Register[MyClassWithConstants]()
}

type Singleton struct {
	gdextension.Class[Singleton, Node.Advanced]
}

func (Singleton) Ready() {
	fmt.Println("Singleton Ready!")
}

func TestSingleton(t *testing.T) {
	gdextension.Register[Singleton]()
	Engine.Advanced().RegisterSingleton(internal.NewStringName("HelloWorld"), new(Singleton).AsObject())
}
