package main

import (
	"fmt"
	"runtime"

	"github.com/readykit/gd"
)

// main init function, where the extensions are exported so that
// they are available to the engine.
func main() {}
func init() {
	gd.Export(
		&HelloWorld{},
		&ExtendedNode{},
	)
}

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	gd.Object
	gd.Extension
	gd.Scripting[struct {
		Print gd.Method `gd:"print()"`
		Echo  gd.Method `gd:"echo(s)"`
		Arch  gd.Method `gd:"arch() GOARCH"`
	}]
}

// Print prints "Hello World"
func (h HelloWorld) Print() {
	fmt.Println("Hello World")
}

// Echo prints the given string, signalling that it
// was printed by Go code.
func (h HelloWorld) Echo(s string) {
	fmt.Println(s + " from Go!")
}

// Arch returns the current GOARCH value.
func (h HelloWorld) Arch() string {
	return runtime.GOARCH
}

/*
ExtendedNode demonstrates how to call the methods of builtin objects.
*/
type ExtendedNode struct {
	gd.Node2D
	gd.Extension

	hello HelloWorld
}

func (e *ExtendedNode) Ready() {
	if gd.Engine.IsEditorHint() {
		fmt.Println(gd.Engine.GetLicenseText())
		return
	}

	fmt.Println(e.Node2D)

	fmt.Println("class:", e.CanvasItem().Node().Object().GetClass())

	// var obj gd.Object
	// e.Attach(&obj)
	// fmt.Println(obj.GetClass())
	// defer e.Detach(&obj)

	fmt.Println(gd.Engine.GetSingletonList())
	fmt.Println("Scene is ready!")

	e.hello.Print()

	fmt.Println("sin=", gd.Sin(1.5))

	fmt.Println("rotation=", e.GetRotation())
	e.SetRotation(3.14)
	fmt.Println("rotation=", e.GetRotation())

	pos := e.GetPosition()

	fmt.Println("position=", pos)

	pos.X = 100

	e.SetPosition(pos)
	fmt.Println("position=", pos)
}
