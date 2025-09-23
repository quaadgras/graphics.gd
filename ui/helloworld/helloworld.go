package main

import (
	"graphics.gd/startup"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
)

func main() {
	startup.LoadingScene()
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect)
	hello.SetHorizontalAlignment(GUI.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(GUI.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
	startup.Scene()
}
