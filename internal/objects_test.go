package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Object"
)

func TestObjectIDs(t *testing.T) {
	node := Node.New()
	node.SetName("test")

	nodeID := node.ID()

	if node, ok := nodeID.Instance(); ok {
		if node.Name() != "test" {
			t.Errorf("expected name 'test', got '%s'", node.Name())
		}
	} else {
		t.Error("expected valid instance")
	}

	node.AsObject()[0].Free()

	if _, ok := nodeID.Instance(); ok {
		t.Error("expected invalid instance after free")
	}
}

func TestAliasFreed(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected panic when accessing freed object")
		}
	}()
	node := Node.New()
	child := Node.New()
	child.SetName("Hello")
	node.AddChild(child)
	alias := node.GetChild(0)

	pointers.Cycle() // don't call this twice.

	child.AsObject()[0].Free()

	if alias.Name() == "Hello" {
		t.Error("access alias after free")
	} else {
		t.Error("corrupted name")
	}
}

type MyObject struct {
	Object.Extension[MyObject]

	Field1 string
	Field2 int
}

func init() {
	classdb.Register[MyObject]()
}

func TestGetSet(t *testing.T) {
	var basis_test string = `extends Object

func set_fields(testing: MyObject):
	testing.Field1 = "Hello"
	testing.Field2 = 42

`
	var runner = Object.New()
	var script = GDScript.New().AsScript()
	script.SetSourceCode(basis_test)
	script.Reload()
	runner.SetScript(script)

	var myobject = new(MyObject)
	Object.Call(runner, "set_fields", myobject)

	if myobject.Field1 != "Hello" || myobject.Field2 != 42 {
		t.Errorf("Expected Field1='Hello', Field2=42, got %v, %v", myobject.Field1, myobject.Field2)
	}
}

func TestObjectAsGoClass(t *testing.T) {
	var object = new(MyObject)
	ptr, ok := Object.As[*MyObject](Object.Instance(object.AsObject()))
	if !ok {
		t.Error("Expected to convert Object to *MyObject")
	}
	if ptr != object {
		t.Error("Expected to get the same pointer back")
	}
}
