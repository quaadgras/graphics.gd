package main

import (
	"fmt"
	"runtime"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

var godot gd.API

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	gd.Class[HelloWorld, gd.Object]
	/*gd.Scripting[struct {
		Print gd.Method `gd:"print()"`
		Echo  gd.Method `gd:"echo(s)"`
		Arch  gd.Method `gd:"arch() GOARCH"`
	}]*/
}

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
	gd.Class[ExtendedNode, gd.Node2D]

	hello HelloWorld

	engine gd.Engine
}

func (e *ExtendedNode) Ready() {
	fmt.Println("Ready!")

	node := e.Super()

	fmt.Println("class:", node.Object().GetClass().String())

	//var obj = gd.New[gd.Object](node)
	//obj.GetClass()

	fmt.Println(e.engine.GetSingletonList().Slice())
	fmt.Println("Scene is ready!")

	e.hello.Print()

	//fmt.Println("sin=", godot.Utility_sin(1.5))

	fmt.Println("rotation=", node.GetRotation())
	node.SetRotation(3.14)
	fmt.Println("rotation=", node.GetRotation())

	pos := node.GetPosition()

	fmt.Println("position=", pos)

	pos.X = 100

	node.SetPosition(pos)
	fmt.Println("position=", pos)
}

// main init function, where the extensions are exported so that
// they are available to the engine.
func main() {
	extension, classdb, ok := gdextension.Link()
	if !ok {
		return
	}
	gdextension.RegisterStruct[HelloWorld, gd.Object](extension, classdb)
	gdextension.RegisterStruct[ExtendedNode, gd.Node2D](extension, classdb)
}
