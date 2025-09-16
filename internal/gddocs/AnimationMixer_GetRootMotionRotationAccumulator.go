/*
[gdscript]
var prev_root_motion_rotation_accumulator

func _process(delta):
	if Input.is_action_just_pressed("animate"):
		state_machine.travel("Animate")
	var current_root_motion_rotation_accumulator = animation_tree.get_root_motion_rotation_accumulator()
	var difference = prev_root_motion_rotation_accumulator.inverse() * current_root_motion_rotation_accumulator
	prev_root_motion_rotation_accumulator = current_root_motion_rotation_accumulator
	transform.basis *=  Basis(difference)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Quaternion"
)

func AnimationMixer_GetRootMotionRotationAccumulator() {
	var prev_root_motion_rotation_accumulator Quaternion.IJKX

	if Input.IsActionJustPressed("animate", false) {
		animationNodeStateMachinePlayback.Travel("Animate")
	}
	var current_root_motion_rotation_accumulator = animationTree.AsAnimationMixer().GetRootMotionRotationAccumulator()
	var difference = Quaternion.Mul(Quaternion.Inverse(prev_root_motion_rotation_accumulator), current_root_motion_rotation_accumulator)
	prev_root_motion_rotation_accumulator = current_root_motion_rotation_accumulator
	transform := characterBody3D.AsNode3D().Transform()
	transform.Basis = Basis.Mul(transform.Basis, Quaternion.AsBasis(difference))
	characterBody3D.AsNode3D().SetTransform(transform)
}
