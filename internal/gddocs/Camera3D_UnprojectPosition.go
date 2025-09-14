/*
# This code block is part of a script that inherits from Node3D.
# `control` is a reference to a node inheriting from Control.
control.visible = not get_viewport().get_camera_3d().is_position_behind(global_transform.origin)
control.position = get_viewport().get_camera_3d().unproject_position(global_transform.origin)
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/Viewport"
)

var node3d Node3D.Instance
var control Control.Instance

func Camera3D_UnprojectPosition() {
	// This code block is part of a script that inherits from Node3D.
	// `control` is a reference to a node inheriting from Control.
	control.AsCanvasItem().SetVisible(!Viewport.Get(node3d.AsNode()).GetCamera3d().IsPositionBehind(node3d.GlobalTransform().Origin))
	control.SetPosition(Viewport.Get(node3d.AsNode()).GetCamera3d().UnprojectPosition(node3d.GlobalTransform().Origin))
}
