package gd_test

import (
	"testing"

	"graphics.gd/classdb/Node"
	"graphics.gd/internal/pointers"
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
