package main

import (
	"fmt"
	"runtime"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

/*
HelloWorld is a simple extension to demonstrate how to export
Go methods so that they can be used in scripts.
*/
type HelloWorld struct {
	gd.Class[HelloWorld, gd.Object]
}

// Print prints "Hello World"
func (h *HelloWorld) Print(gd.Context) {
	fmt.Println("Hello World")
}

// Echo prints the given string, signalling that it
// was printed by Go code.
func (h *HelloWorld) Echo(_ gd.Context, s gd.String) {
	fmt.Println(s.String() + " from Go!")
}

// Arch returns the current GOARCH value.
func (h *HelloWorld) Arch(godot gd.Context) gd.String {
	return godot.String(runtime.GOARCH)
}

type Rotator struct {
	gd.Class[Rotator, gd.Sprite2D]
}

func (r *Rotator) Process(_ gd.Context, delta gd.Float) {
	node2D := r.Super().AsNode2D()
	node2D.SetRotation(node2D.GetRotation() + delta)
}

type StartedSignalEmitter struct {
	gd.Class[StartedSignalEmitter, gd.Node]

	started gd.Signal[func()]
}

func (r *StartedSignalEmitter) Ready(gd.Context) {
	if r.started.Emit != nil {
		r.started.Emit()
	}
}

/*
ExtendedNode demonstrates how to call the methods of builtin objects.
*/
type ExtendedNode struct {
	gd.Class[ExtendedNode, gd.Node2D]

	StringField gd.String

	engine gd.Engine
}

func (e *ExtendedNode) Ready(godot gd.Context) {
	fmt.Println("Ready!")

	node := e.Super()

	fmt.Println("class:", node.AsObject().GetClass(godot).String())

	var obj = gd.Create(godot, new(gd.Object))
	fmt.Println(obj.GetClass(godot).String())

	fmt.Println(e.engine.GetSingletonList(godot).AsSlice())
	fmt.Println("Scene is ready!")

	fmt.Println("sin=", godot.Sin(1.5))

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
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	fmt.Println("Godot Version is: ", godot.Version())
	fmt.Println("Extension: ", godot.GetLibraryPath())
	gdextension.Register[HelloWorld](godot)
	gdextension.Register[ExtendedNode](godot)
	gdextension.Register[Rotator](godot)
	gdextension.Register[StartedSignalEmitter](godot)
}
