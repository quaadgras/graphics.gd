/*
[gdscript]
var current_scale = Vector3(1, 1, 1)
var scale_accum = Vector3(1, 1, 1)

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        current_scale = get_scale()
        scale_accum = Vector3(1, 1, 1)
        state_machine.travel("Animate")
    scale_accum += animation_tree.get_root_motion_scale()
    set_scale(current_scale * scale_accum)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Vector3"
)

func AnimationMixer_GetRootMotionScale() {
	var current_scale = Vector3.New(1, 1, 1)
	var scale_accum = Vector3.New(1, 1, 1)

	if Input.IsActionJustPressed("animate", false) {
		current_scale = characterBody3D.AsNode3D().Scale()
		scale_accum = Vector3.New(1, 1, 1)
		animationNodeStateMachinePlayback.Travel("Animate")
	}
	scale_accum = Vector3.Add(scale_accum, animationTree.AsAnimationMixer().GetRootMotionScale())
	characterBody3D.AsNode3D().SetScale(Vector3.Mul(current_scale, scale_accum))
}
