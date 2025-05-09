package gd_test

import (
	"testing"

	"graphics.gd/classdb/Node"
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
