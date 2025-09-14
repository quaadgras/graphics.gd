/*
[gdscript]
if Engine.is_editor_hint():
    draw_gizmos()
else:
    simulate_physics()
[/gdscript]
[csharp]
if (Engine.IsEditorHint())
    DrawGizmos();
else
    SimulatePhysics();
[/csharp]
*/

package main

import "graphics.gd/classdb/Engine"

func DrawGizmos()      {}
func SimulatePhysics() {}

func Engine_IsEditorHint() {
	if Engine.IsEditorHint() {
		DrawGizmos()
	} else {
		SimulatePhysics()
	}
}
