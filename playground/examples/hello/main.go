// A label in the middle of the window.
package main

import (
	"graphics.gd/startup"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
)

func main() {
	startup.LoadingScene() // set up the SceneTree and wait for the engine.
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect) // fill the window.
	hello.SetHorizontalAlignment(GUI.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(GUI.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
	startup.Scene() // run the scene until the engine shuts down.
}
