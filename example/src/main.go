package main

import (
	"fmt"

	"github.com/readykit/gd"
)

func main() {}

type HelloWorld struct {
	gd.Extension

	object gd.Object

	Message string
}

var NewHelloWorld, NewHelloWorldAt = gd.Register(func(hello *HelloWorld) (*gd.Extension, *gd.Object) {
	hello.Message = "Hello World!"
	return &hello.Extension, gd.NewObjectAt(&hello.object, gd.BelongsTo(hello))
})

func (h HelloWorld) Print() {
	fmt.Println(h.Message)
}

func (h HelloWorld) Echo(s string) {
	fmt.Println(s + " from Go!")
}

type ExtendedNode struct {
	gd.Extension

	node  gd.Node2D
	hello HelloWorld
}

var NewExtendedNode, NewExtendedNodeAt = gd.Register(func(enode *ExtendedNode) (*gd.Extension, *gd.Node2D) {

	NewHelloWorldAt(&enode.hello, gd.BelongsTo(enode))

	return &enode.Extension, gd.NewNode2DAt(&enode.node, gd.BelongsTo(enode))
})

func (e ExtendedNode) Free() {
	e.hello.Free()
	e.Extension.Free()
}

func (e *ExtendedNode) Ready() {
	if gd.Engine.IsEditorHint() {
		fmt.Println(gd.Engine.GetLicenseText())
		return
	}

	fmt.Println(e.node.CanvasItem().Node().Object().GetClass())

	var obj = gd.NewObject(gd.BelongsTo(e))
	fmt.Println(obj.GetClass())
	defer obj.Free()

	//gd.LoadSingletons()

	fmt.Println(gd.Engine.GetSingletonList())

	fmt.Println("Scene is ready!")

	e.hello.Print()

	fmt.Println("sin=", gd.Sin(1.5))

	fmt.Println("rotation=", e.node.GetRotation())
	e.node.SetRotation(3.14)
	fmt.Println("rotation=", e.node.GetRotation())

	pos := e.node.GetPosition()

	fmt.Println("position=", pos)

	pos.X = 100

	e.node.SetPosition(pos)
	fmt.Println("position=", pos)

	//gd.DisplayServerSingleton(gd.Engine.GetSingleton("DisplayServer"))

	//fmt.Println(godot.Engine.GetSingletonList())

	fmt.Println(gd.DisplayServer)

	gd.DisplayServer.WindowSetCurrentScreen(1, 0)
}
