//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
)

// Types for declarative children tests.

type FlatParent struct {
	Node.Extension[FlatParent]
	ChildA Node.Instance
}

type MultiParent struct {
	Node.Extension[MultiParent]
	First  Node.Instance
	Second Node.Instance
}

type NestedParent struct {
	Node.Extension[NestedParent]

	Container struct {
		Node.Instance

		Inner Node.Instance
	}
}

type DeepParent struct {
	Node.Extension[DeepParent]

	Level1 struct {
		Node.Instance

		Level2 struct {
			Node.Instance

			Leaf Node.Instance
		}
	}
}

func init() {
	classdb.Register[FlatParent]()
	classdb.Register[MultiParent]()
	classdb.Register[NestedParent]()
	classdb.Register[DeepParent]()
}

// TestDeclarativeChildFlat verifies that a simple exported Node.Instance
// field is automatically created and added as a child during Ready.
func TestDeclarativeChildFlat(t *testing.T) {
	parent := new(FlatParent)
	runOnMain(t, func(t testing.TB) {
		SceneTree.Add(parent.AsNode())
	})
	runOnMain(t, func(t testing.TB) {
		if parent.ChildA == (Node.Instance{}) {
			t.Fatal("expected ChildA to be populated after Ready, got zero value")
		}
		if parent.ChildA.Name() != "ChildA" {
			t.Fatalf("expected child name 'ChildA', got %q", parent.ChildA.Name())
		}
	})
}

// TestDeclarativeChildMultiple verifies that multiple exported Node fields
// are each created as separate children.
func TestDeclarativeChildMultiple(t *testing.T) {
	parent := new(MultiParent)
	runOnMain(t, func(t testing.TB) {
		SceneTree.Add(parent.AsNode())
	})
	runOnMain(t, func(t testing.TB) {
		if parent.First == (Node.Instance{}) {
			t.Fatal("expected First to be populated")
		}
		if parent.Second == (Node.Instance{}) {
			t.Fatal("expected Second to be populated")
		}
		if parent.AsNode().GetChildCount() < 2 {
			t.Fatalf("expected at least 2 children, got %d", parent.AsNode().GetChildCount())
		}
	})
}

// TestDeclarativeChildNested verifies that a named struct field with an
// embedded Node type and its own child fields are all created and nested
// correctly in the scene tree.
func TestDeclarativeChildNested(t *testing.T) {
	parent := new(NestedParent)
	runOnMain(t, func(t testing.TB) {
		SceneTree.Add(parent.AsNode())
	})
	runOnMain(t, func(t testing.TB) {
		if parent.Container.Instance == (Node.Instance{}) {
			t.Fatal("expected Container to be populated")
		}
		if parent.Container.Inner == (Node.Instance{}) {
			t.Fatal("expected Container.Inner to be populated")
		}
		if parent.Container.Inner.Name() != "Inner" {
			t.Fatalf("expected inner child name 'Inner', got %q", parent.Container.Inner.Name())
		}
	})
}

// TestDeclarativeChildDeeplyNested verifies multi-level nesting where
// nested structs contain further nested structs with their own children.
func TestDeclarativeChildDeeplyNested(t *testing.T) {
	parent := new(DeepParent)
	runOnMain(t, func(t testing.TB) {
		SceneTree.Add(parent.AsNode())
	})
	runOnMain(t, func(t testing.TB) {
		if parent.Level1.Instance == (Node.Instance{}) {
			t.Fatal("expected Level1 to be populated")
		}
		if parent.Level1.Level2.Instance == (Node.Instance{}) {
			t.Fatal("expected Level1.Level2 to be populated")
		}
		if parent.Level1.Level2.Leaf == (Node.Instance{}) {
			t.Fatal("expected Level1.Level2.Leaf to be populated")
		}
	})
}
