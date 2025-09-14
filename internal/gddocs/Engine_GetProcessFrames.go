/*
[gdscript]
func _process(_delta):
    if Engine.get_process_frames() % 5 == 0:
        pass # Run expensive logic only once every 5 process (render) frames here.
[/gdscript]
[csharp]
public override void _Process(double delta)
{
    base._Process(delta);

    if (Engine.GetProcessFrames() % 5 == 0)
    {
        // Run expensive logic only once every 5 process (render) frames here.
    }
}
[/csharp]
*/

package main

import "graphics.gd/classdb/Engine"

func Engine_GetProcessFrames() {
	if Engine.GetProcessFrames()%5 == 0 {
		// Run expensive logic only once every 5 process (render) frames here.
	}
}
