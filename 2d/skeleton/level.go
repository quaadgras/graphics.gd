package main

import (
	"graphics.gd/classdb/Camera2D"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type Level struct {
	Node2D.Extension[Level]

	CameraLimitMin Marker2D.Instance `gd:"CameraLimit_min"`
	CameraLimitMax Marker2D.Instance `gd:"CameraLimit_max"`
}

func (l *Level) Ready() {
	camera, ok := Object.As[Camera2D.Instance](l.AsNode().FindChild("Camera2D"))
	if !ok {
		return
	}
	var min_pos = l.CameraLimitMin.AsNode2D().GlobalPosition()
	var max_pos = l.CameraLimitMax.AsNode2D().GlobalPosition()
	camera.SetLimitLeft(int(Float.Round(min_pos.X)))
	camera.SetLimitTop(int(Float.Round(min_pos.Y)))
	camera.SetLimitRight(int(Float.Round(max_pos.X)))
	camera.SetLimitBottom(int(Float.Round(max_pos.Y)))
}
