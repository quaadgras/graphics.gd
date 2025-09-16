/*
[gdscript]
func _physics_process(_delta):
	if Engine.get_physics_frames() % 2 == 0:
		pass # Run expensive logic only once every 2 physics frames here.
[/gdscript]
[csharp]
public override void _PhysicsProcess(double delta)
{
	base._PhysicsProcess(delta);

	if (Engine.GetPhysicsFrames() % 2 == 0)
	{
		// Run expensive logic only once every 2 physics frames here.
	}
}
[/csharp]
*/

package main

import "graphics.gd/classdb/Engine"

func Engine_GetPhysicsFrames() {
	if Engine.GetPhysicsFrames()%2 == 0 {
		// Run expensive logic only once every 2 physics frames here.
	}
}
