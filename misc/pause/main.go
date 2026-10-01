package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
)

// PauseButton pauses and unpauses the scene tree when toggled.
type PauseButton struct {
	Button.Extension[PauseButton]
}

func (b *PauseButton) Ready() {
	// This ensures that this Node won't be paused, allowing it to
	// process even when the SceneTree is paused. Without that it would
	// not be able to unpause the game. Note that you can set this through
	// the inspector as well.
	b.AsNode().SetProcessMode(Node.ProcessModeAlways)
}

func (b *PauseButton) Toggled(isButtonPressed bool) {
	// Pause or unpause the SceneTree based on whether the button is
	// toggled on or off.
	SceneTree.Get(b.AsNode()).SetPaused(isButtonPressed)
	if isButtonPressed {
		b.AsButton().SetText("Unpause")
	} else {
		b.AsButton().SetText("Pause")
	}
}

// ProcessModeButton selects the process mode of the cube's AnimationPlayer.
type ProcessModeButton struct {
	OptionButton.Extension[ProcessModeButton]

	CubeAnimation AnimationPlayer.Instance `gd:"../../AnimationPlayer"`
}

func (o *ProcessModeButton) OnOptionButtonItemSelected(index int) {
	o.CubeAnimation.AsNode().SetProcessMode(Node.ProcessMode(index))
}

func main() {
	classdb.Register[PauseButton]()
	classdb.Register[ProcessModeButton]()
	startup.Scene()
}
