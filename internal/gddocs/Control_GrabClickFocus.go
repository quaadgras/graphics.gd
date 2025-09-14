/*
[gdscript]
func _process(delta):
    grab_click_focus() # When clicking another Control node, this node will be clicked instead.
[/gdscript]
[csharp]
public override void _Process(double delta)
{
    GrabClickFocus(); // When clicking another Control node, this node will be clicked instead.
}
[/csharp]
*/

package main

func Control_GrabClickFocus() {
	control.GrabClickFocus() // When clicking another Control node, this node will be clicked instead.
}
