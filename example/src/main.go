package main

import (
	"fmt"
	"runtime"

	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

var godot gd.API

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	obj gd.Object
	/*gd.Scripting[struct {
		Print gd.Method `gd:"print()"`
		Echo  gd.Method `gd:"echo(s)"`
		Arch  gd.Method `gd:"arch() GOARCH"`
	}]*/
}

func (h *HelloWorld) Object() gd.Object { return h.obj }

// Print prints "Hello World"
func (h *HelloWorld) Print() {
	fmt.Println("Hello World")
}

// Echo prints the given string, signalling that it
// was printed by Go code.
func (h *HelloWorld) Echo(s string) {
	fmt.Println(s + " from Go!")
}

// Arch returns the current GOARCH value.
func (h *HelloWorld) Arch() string {
	return runtime.GOARCH
}

/*
ExtendedNode demonstrates how to call the methods of builtin objects.
*/
type ExtendedNode struct {
	node gd.Node2D

	hello HelloWorld
}

func (e *ExtendedNode) Node2D() gd.Node2D { return e.node }

func (e *ExtendedNode) Ready() {
	if godot.Engine.IsEditorHint() {
		fmt.Println(godot.Engine.GetLicenseText())
		return
	}

	fmt.Println("class:", e.node.CanvasItem().Node().Object().GetClass())

	// var obj gd.Object
	// e.Attach(&obj)
	// fmt.Println(obj.GetClass())
	// defer e.Detach(&obj)

	fmt.Println(godot.Engine.GetSingletonList())
	fmt.Println("Scene is ready!")

	e.hello.Print()

	fmt.Println("sin=", godot.Utility_sin(1.5))

	fmt.Println("rotation=", e.node.GetRotation())
	e.node.SetRotation(3.14)
	fmt.Println("rotation=", e.node.GetRotation())

	pos := e.node.GetPosition()

	fmt.Println("position=", pos)

	pos.X = 100

	e.node.SetPosition(pos)
	fmt.Println("position=", pos)
}

// main init function, where the extensions are exported so that
// they are available to the engine.
func main() {
	//gdextension.Register[HelloWorld]("HelloWorld")
	//gdextension.Register[ExtendedNode]("ExtendedNode")
	fmt.Println("Hello World!")
}
