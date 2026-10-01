package main

// This acts as a staging scene shown until the main scene is fully loaded.

import (
	"graphics.gd/classdb/ColorRect"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Float"
)

type LoadingScreen struct {
	ColorRect.Extension[LoadingScreen]

	frames int
}

func (l *LoadingScreen) Process(_ Float.X) {
	l.frames++
	// Wait 2 frames before starting to change to the main scene,
	// so that the loading text can be shown instead of the splash screen.
	if l.frames == 2 {
		// Do not load the scene earlier to avoid incurring the loading time
		// before the loading text can be shown.
		SceneTree.Get(l.AsNode()).ChangeSceneToPacked(Resource.Load[PackedScene.Instance]("res://test.tscn"))
	}
}
