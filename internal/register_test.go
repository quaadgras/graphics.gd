//go:build !generate

package gd_test

import (
	"fmt"
	"testing"

	"grow.graphics/gd/gdextension"
	gd "grow.graphics/gd/internal"
)

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

	fmt.Println(e.engine.GetSingletonList(godot).AsSlice(godot))
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

	variant := godot.Variant(node)
	result, err := variant.Call(godot, godot.StringName("get_position"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result.Interface(godot))
	}
}

func TestRegister(t *testing.T) {
	godot := gd.NewContext(API)
	defer godot.End()

	gdextension.Register[ExtendedNode](godot)

	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("ExtendedNode")); tag == 0 {
		t.Fail()
	}
}
