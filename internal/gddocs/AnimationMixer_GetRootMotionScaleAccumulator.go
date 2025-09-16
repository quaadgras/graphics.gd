/*
[gdscript]
var prev_root_motion_scale_accumulator

func _process(delta):
	if Input.is_action_just_pressed("animate"):
		state_machine.travel("Animate")
	var current_root_motion_scale_accumulator = animation_tree.get_root_motion_scale_accumulator()
	var difference = current_root_motion_scale_accumulator - prev_root_motion_scale_accumulator
	prev_root_motion_scale_accumulator = current_root_motion_scale_accumulator
	transform.basis = transform.basis.scaled(difference)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Vector3"
)

func AnimationMixer_GetRootMotionScaleAccumulator() {
	var prev_root_motion_scale_accumulator Vector3.XYZ

	if Input.IsActionJustPressed("animate", false) {
		animationNodeStateMachinePlayback.Travel("Animate")
	}
	var current_root_motion_scale_accumulator = animationTree.AsAnimationMixer().GetRootMotionScaleAccumulator()
	var difference = Vector3.Sub(current_root_motion_scale_accumulator, prev_root_motion_scale_accumulator)
	prev_root_motion_scale_accumulator = current_root_motion_scale_accumulator
	transform := characterBody3D.AsNode3D().Transform()
	transform.Basis = Basis.Scaled(transform.Basis, difference)
	characterBody3D.AsNode3D().SetTransform(transform)
}
