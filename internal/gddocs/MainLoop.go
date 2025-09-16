/*
[gdscript]
class_name CustomMainLoop
extends MainLoop

var time_elapsed = 0

func _initialize():
	print("Initialized:")
	print("  Starting time: %s" % str(time_elapsed))

func _process(delta):
	time_elapsed += delta
	# Return true to end the main loop.
	return Input.get_mouse_button_mask() != 0 || Input.is_key_pressed(KEY_ESCAPE)

func _finalize():
	print("Finalized:")
	print("  End time: %s" % str(time_elapsed))
[/gdscript]
[csharp]
using Godot;

[GlobalClass]
public partial class CustomMainLoop : MainLoop
{
	private double _timeElapsed = 0;

	public override void _Initialize()
	{
		GD.Print("Initialized:");
		GD.Print($"  Starting Time: {_timeElapsed}");
	}

	public override bool _Process(double delta)
	{
		_timeElapsed += delta;
		// Return true to end the main loop.
		return Input.GetMouseButtonMask() != 0 || Input.IsKeyPressed(Key.Escape);
	}

	private void _Finalize()
	{
		GD.Print("Finalized:");
		GD.Print($"  End Time: {_timeElapsed}");
	}
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/variant/Float"
)

type MyMainLoop struct {
	MainLoop.Extension[MyMainLoop]

	timeElapsed Float.X
}

func (m *MyMainLoop) Initialize() {
	println("Initialized:")
	println("  Starting time: ", m.timeElapsed)
}

func (m *MyMainLoop) Process(delta Float.X) bool {
	m.timeElapsed += delta
	return Input.GetMouseButtonMask() != 0 || Input.IsKeyPressed(Input.KeyEscape)
}

func (m *MyMainLoop) Finalize() {
	println("Finalized:")
	println("  End time: ", m.timeElapsed)
}
