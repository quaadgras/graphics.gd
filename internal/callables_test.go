package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Object"
)

var callable_test string = `extends Object

var node: Node

func test_callable(c):
	node = c.call()
	return

func get_node() -> Node:
	return node
`

func TestCallable(t *testing.T) {
	var runner = Object.New()
	var script = GDScript.New().AsScript()
	script.SetSourceCode(callable_test)
	script.Reload()
	runner.SetScript(script)

	Object.Call(runner, "test_callable", Callable.New(func() Node.Instance {
		var node = Node.New()
		node.SetName("TestNode")
		return node
	}))
	//pointers.Cycle()
	//pointers.Cycle()

	var node Node.Instance = Object.Call(runner, "get_node").(classdb.Node)
	if node.Name() != "TestNode" {
		t.Fatalf("Expected 'TestNode', got '%s'", node.Name())
	}
}
