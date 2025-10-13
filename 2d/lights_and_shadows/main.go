package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/DirectionalLight2D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Light2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
	"graphics.gd/variant/Int"
	"graphics.gd/variant/Object"
)

type LightShadows struct {
	Node2D.Extension[LightShadows]

	DirectionalLight2D DirectionalLight2D.Instance
}

func (l *LightShadows) Input(event InputEvent.Instance) {
	if event.IsActionPressed("toggle_directional_light") {
		l.DirectionalLight2D.AsCanvasItem().SetVisible(!l.DirectionalLight2D.AsCanvasItem().Visible())
	}
	if event.IsActionPressed("toggle_point_lights") {
		for _, point_light := range SceneTree.Get(l.AsNode()).GetNodesInGroup("point_light") {
			point_light := Object.To[CanvasItem.Instance](point_light)
			point_light.SetVisible(!point_light.Visible())
		}
	}
	if event.IsActionPressed("cycle_directional_light_shadows_quality") {
		filter := l.DirectionalLight2D.AsLight2D().ShadowFilter()
		l.DirectionalLight2D.AsLight2D().SetShadowFilter(Int.Wrap(filter+1, 0, 3))
	}
	if event.IsActionPressed("cycle_point_light_shadows_quality") {
		for _, point_light := range SceneTree.Get(l.AsNode()).GetNodesInGroup("point_light") {
			point_light := Object.To[Light2D.Instance](point_light)
			filter := point_light.AsLight2D().ShadowFilter()
			point_light.AsLight2D().SetShadowFilter(Int.Wrap(filter+1, 0, 3))
		}
	}
}

func main() {
	classdb.Register[LightShadows]()
	startup.Scene()
}
