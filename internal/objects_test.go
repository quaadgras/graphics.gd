package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
)

func TestObjectIDs(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
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

		gd.ObjectFree(node.AsObject()[0])

		if _, ok := nodeID.Instance(); ok {
			t.Error("expected invalid instance after free")
		}
	})
}

func TestAliasFreed(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
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

		gdreference.GC(gd.Free)
		gdreference.GC(gd.Free)

		if alias.Name() == "Hello" {
			t.Error("access alias after free")
		} else {
			t.Error("corrupted name")
		}
	})
}

type MyObject struct {
	Object.Extension[MyObject]

	Field1 string
	Field2 int
}

func init() {
	classdb.Register[MyObject]()
}

type MyTool struct {
	Object.Extension[MyTool]
	classdb.Tool
}

func init() {
	classdb.Register[MyTool]()
}

func TestGetSet(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
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
	})
}

func TestObjectAsGoClass(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var object = new(MyObject)
		ptr, ok := Object.As[*MyObject](Object.Instance(object.AsObject()))
		if !ok {
			t.Error("Expected to convert Object to *MyObject")
		}
		if ptr != object {
			t.Error("Expected to get the same pointer back")
		}
	})
}

func TestObjectAsGoTool(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var object = new(MyTool)
		ptr, ok := Object.As[*MyTool](Object.Instance(object.AsObject()))
		if !ok {
			t.Error("Expected to convert Object to *MyTool")
		}
		if ptr != object {
			t.Error("Expected to get the same pointer back")
		}
	})
}

type MyNode struct {
	Node.Extension[MyNode]
}

func init() {
	classdb.Register[MyNode]()
}

func TestExtensionClassAliasCastThenAddedToScene(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var m = new(MyNode)

		// simulate the scenario where the engine returns MyNode as an 'owned' Object, ie. PackedScene.Instantiate
		var another_ref_from_the_engine = gdreference.OwnObject(gdreference.GetObject(m.AsObject()[0]), gd.Free)
		var obj = Node.Instance{gdclass.NewNode(another_ref_from_the_engine)}

		m = Object.To[*MyNode](obj)
		var node = Node.New()
		node.AddChild(obj)
	})
}

var call_callable string = `extends Object

var obj: MyObject

func test_basis(get_obj: Callable):
	obj = get_obj.call()
`

func TestExtensionClassReturnedToTheEngineFromCallable(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var script = GDScript.New().AsScript()
		script.SetSourceCode(call_callable)
		script.Reload()

		var myobj = new(MyObject)
		myobj.Field1 = "Hello from callable"
		myobj.Field2 = 123

		var alias = Node.Instance{gdclass.NewNode(gdreference.RawObject(gdreference.GetObject(myobj.AsObject()[0])))}

		var runner = Object.New()
		runner.SetScript(script)
		Object.Call(runner, "test_basis", Callable.New(func() *MyObject {
			return myobj
		}))
		Object.Call(runner, "test_basis", Callable.New(func() Object.Instance {
			return alias.AsObject()
		}))

		keep_reachable_instances_alive()
		keep_reachable_instances_alive()

		if !Object.InstanceIsValid(Object.Instance(myobj.AsObject())) {
			t.Error("Expected MyObject to still be valid after GC")
		}
	})
	// Also test the case when the object is passed back from the engine as an aliased object.
	runOnMain(t, func(t testing.TB) {
		var script = GDScript.New().AsScript()
		script.SetSourceCode(call_callable)
		script.Reload()

		var myobj = new(MyObject)
		myobj.Field1 = "Hello from callable"
		myobj.Field2 = 123

		var alias = Node.Instance{gdclass.NewNode(gdreference.RawObject(gdreference.GetObject(myobj.AsObject()[0])))}

		var runner = Object.New()
		runner.SetScript(script)
		Object.Call(runner, "test_basis", Callable.New(func() Object.Instance {
			return alias.AsObject()
		}))

		keep_reachable_instances_alive()
		keep_reachable_instances_alive()

		if !Object.InstanceIsValid(Object.Instance(myobj.AsObject())) {
			t.Error("Expected MyObject to still be valid after GC")
		}
	})
}

func TestJumpOnlyCall(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var node = Node.New()
		node.SetProcessInput(true)
		if !node.IsProcessingInput() {
			t.Error("Expected node to be processing input")
		}
		node.SetProcessInput(false)
		if node.IsProcessingInput() {
			t.Error("Expected node to not be processing input")
		}
	})
}

type KeepAliveNode struct {
	Node.Extension[KeepAliveNode]

	Object *MyObject
}

func init() {
	classdb.Register[KeepAliveNode]()
}

func TestFree(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var obj = new(MyObject)
		Object.Free(obj)
		if Object.InstanceIsValid(Object.Instance(obj.AsObject())) {
			t.Error("Expected object to be invalid after free")
		}
	})
}

func TestAutomaticKeepAlive(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		var node = Node.New()
		var child = new(KeepAliveNode)
		child.Object = new(MyObject)
		child.Object.AsObject() // trigger the reference
		gd.ExtensionInstanceGoOnly(gdreference.GetObject(child.AsObject()[0]), true)
		node.AddChild(child.AsNode())

		keep_reachable_instances_alive()
		keep_reachable_instances_alive()

		if !Object.InstanceIsValid(Object.Instance(child.Object.AsObject())) {
			t.Error("Expected child node to still be valid after GC")
		}
	})
}

func TestNoInheritance(t *testing.T) {
	type Common struct {
		Node.Extension[Common]
	}
	type Player struct {
		Common
	}
	runOnMain(t, func(t testing.TB) {
		defer func() {
			if recover() == nil {
				t.Error("Expected panic when trying to cast to an unrelated type")
			}
		}()
		classdb.Register[Player]()
	})
}
