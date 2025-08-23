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
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

func TestRegister(t *testing.T) {
	type TestingSimpleClass struct {
		Node2D.Extension[TestingSimpleClass]
	}
	classdb.Register[TestingSimpleClass]()

	if tag := gdextension.Host.Objects.Type(pointers.Get(gd.NewStringName("Node2D"))); tag == 0 {
		t.Fail()
	}
	if tag := gdextension.Host.Objects.Type(pointers.Get(gd.NewStringName("TestingSimpleClass"))); tag == 0 {
		t.Fail()
	}

	class := new(TestingSimpleClass)
	class_name := class.AsObject()[0].GetClass()
	if name := class_name.String(); name != "TestingSimpleClass" {
		t.Fatal(name)
	}
	class.AsNode().SetName("SimpleClass")
}

func TestEmbedding(t *testing.T) {
	type TestingEmbeddedClass struct {
		Node2D.Extension[TestingEmbeddedClass]
	}
	classdb.Register[TestingEmbeddedClass]()

	var node = Node.New()

	type Embeds struct {
		TestingEmbeddedClass
	}
	embeds := new(Embeds)
	node.AddChild(embeds.AsNode())
}

type TestingSingleton struct {
	Node.Extension[TestingSingleton]
}

func (TestingSingleton) Ready() {
	fmt.Println("Singleton Ready!")
}

func TestSingleton(t *testing.T) {
	classdb.Register[TestingSingleton]()
	Engine.RegisterSingleton("HelloWorld", new(TestingSingleton).AsObject())
}
