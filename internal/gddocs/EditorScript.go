/*
[gdscript]
@tool
extends EditorScript

func _run():
	print("Hello from the Godot Editor!")
[/gdscript]
[csharp]
using Godot;

[Tool]
public partial class HelloEditor : EditorScript
{
	public override void _Run()
	{
		GD.Print("Hello from the Godot Editor!");
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/EditorScript"
)

type HelloEditor struct {
	EditorScript.Extension[HelloEditor]
}

func (h *HelloEditor) Run() {
	fmt.Println("Hello from the Godot Editor!")
}
