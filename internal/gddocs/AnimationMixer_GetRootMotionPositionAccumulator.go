/*
[gdscript]
var prev_root_motion_position_accumulator

func _process(delta):
	if Input.is_action_just_pressed("animate"):
		state_machine.travel("Animate")
	var current_root_motion_position_accumulator = animation_tree.get_root_motion_position_accumulator()
	var difference = current_root_motion_position_accumulator - prev_root_motion_position_accumulator
	prev_root_motion_position_accumulator = current_root_motion_position_accumulator
	transform.origin += difference
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Vector3"
)

func AnimationMixer_GetRootMotionPositionAccumulator() {
	var prev_root_motion_position_accumulator Vector3.XYZ
	if Input.IsActionJustPressed("animate", false) {
		animationNodeStateMachinePlayback.Travel("Animate")
	}
	var current_root_motion_position_accumulator = animationTree.AsAnimationMixer().GetRootMotionPositionAccumulator()
	var difference = Vector3.Sub(current_root_motion_position_accumulator, prev_root_motion_position_accumulator)
	prev_root_motion_position_accumulator = current_root_motion_position_accumulator
	transform := characterBody3D.AsNode3D().Transform()
	transform.Origin = Vector3.Add(transform.Origin, difference)
	characterBody3D.AsNode3D().SetTransform(transform)
}
