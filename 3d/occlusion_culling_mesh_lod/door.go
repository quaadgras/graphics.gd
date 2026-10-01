package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/OccluderInstance3D"
)

type Door struct {
	Node3D.Extension[Door]

	AnimationPlayer    AnimationPlayer.Instance
	OccluderInstance3D OccluderInstance3D.Instance

	open bool
}

func (d *Door) Input(input_event InputEvent.Instance) {
	if !input_event.IsActionPressed("toggle_doors") {
		return
	}
	if d.open {
		// Close the door.
		// The occluder will be re-enabled when the animation ends
		// using `OnAnimationPlayerAnimationFinished()`.
		d.AnimationPlayer.MoreArgs().PlayBackwards("open", -1)
		d.open = false
	} else {
		// Open the door.
		d.AnimationPlayer.MoreArgs().Play("open", -1, 1.0, false)
		d.open = true
		// Disable the occluder as soon as the door starts opening.
		// The occluder is not part of the pivot to prevent it from having its
		// position changed every frame, which causes the occlusion culling BVH
		// to be rebuilt each frame. This causes a CPU performance penalty.
		d.OccluderInstance3D.AsNode3D().SetVisible(false)
	}
}

func (d *Door) OnAnimationPlayerAnimationFinished(_ string) {
	if !d.open {
		// Re-enable the occluder when the door is done closing.
		// To prevent overocclusion, the door must be fully closed before the occluder can be re-enabled.
		d.OccluderInstance3D.AsNode3D().SetVisible(true)
	}
}
