//go:build !generate

package gd_test

import (
	"fmt"
	"testing"

	gdunsafe "graphics.gd"
	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

func TestRegister(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		type TestingSimpleClass struct {
			Node2D.Extension[TestingSimpleClass]
		}
		classdb.Register[TestingSimpleClass]()

		Node2D_name := gdunsafe.UTF8.Intern("Node2D")
		defer gdunsafe.Free(Node2D_name)

		TestingSimpleClass_name := gdunsafe.UTF8.Intern("TestingSimpleClass")
		defer gdunsafe.Free(TestingSimpleClass_name)

		if tag := gdunsafe.Class(Node2D_name).Tag(); tag == (gdunsafe.ClassTag{}) {
			t.Fail()
		}
		if tag := gdunsafe.Class(TestingSimpleClass_name).Tag(); tag == (gdunsafe.ClassTag{}) {
			t.Fail()
		}

		class := new(TestingSimpleClass)
		class_name := gd.ObjectGetClass(class.AsObject()[0])
		if name := class_name.String(); name != "TestingSimpleClass" {
			t.Fatal(name)
		}
		class.AsNode().SetName("SimpleClass")
	})
}

func TestEmbedding(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
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
	})
}

type TestingSingleton struct {
	Node.Extension[TestingSingleton]
}

func (TestingSingleton) Ready() {
	fmt.Println("Singleton Ready!")
}

func TestSingleton(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		classdb.Register[TestingSingleton]()
		Engine.RegisterSingleton("HelloWorld", new(TestingSingleton).AsObject())
	})
}

type HealthResource struct {
	Resource.Extension[HealthResource]
	MaxHealth Float.X
}
type Health struct {
	Node.Extension[Health]
	Template      *HealthResource
	CurrentHealth Float.X
}

func (h *Health) Ready() {
	h.CurrentHealth = h.Template.MaxHealth
}

type NestedGame struct {
	Node3D.Extension[NestedGame]
	Health *Health
}

func (g *NestedGame) Ready() {

}

func TestNested(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		classdb.Register[HealthResource]()
		classdb.Register[Health]()
		classdb.Register[NestedGame]()

		game := &NestedGame{
			Health: &Health{Template: new(HealthResource)},
		}
		gd.ObjectNotification(game.AsObject()[0], gd.Int(Node.NotificationReady), false)
		game.Health.Ready()
	})
}

// TestExtensionInherits verifies that extension classes can inherit from other
// extension classes without causing the "set_instance_binding" error.
func TestExtensionInherits(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		type ParentExtension struct {
			Node2D.Extension[ParentExtension]
		}

		type ChildExtension struct {
			classdb.ExtensionInherits[ParentExtension, ChildExtension]
		}

		// Register parent first, then child
		classdb.Register[ParentExtension]()
		classdb.Register[ChildExtension]()

		ParentExtension_name := gdunsafe.UTF8.Intern("ParentExtension")
		defer gdunsafe.Free(ParentExtension_name)

		ChildExtension_name := gdunsafe.UTF8.Intern("ChildExtension")
		defer gdunsafe.Free(ChildExtension_name)

		// Verify both classes are registered
		if tag := gdunsafe.Class(ParentExtension_name).Tag(); tag == (gdunsafe.ClassTag{}) {
			t.Fatal("ParentExtension not registered")
		}
		if tag := gdunsafe.Class(ChildExtension_name).Tag(); tag == (gdunsafe.ClassTag{}) {
			t.Fatal("ChildExtension not registered")
		}

		// Create an instance of the child class - this would previously cause
		// "set_instance_binding" error because both parent and child tried to
		// set the instance binding on the same object.
		child := new(ChildExtension)
		className := gd.ObjectGetClass(child.AsObject()[0])
		if name := className.String(); name != "ChildExtension" {
			t.Fatalf("expected ChildExtension, got %s", name)
		}

		// Verify the object can receive notifications without crashing
		gd.ObjectNotification(child.AsObject()[0], gd.Int(Object.NotificationPostInitialize), false)
	})
}
