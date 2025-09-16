/*
[gdscript]
func _process(delta):
	if Input.is_action_just_pressed("animate"):
		state_machine.travel("Animate")
	set_quaternion(get_quaternion() * animation_tree.get_root_motion_rotation())
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Quaternion"
)

func AnimationMixer_GetRootMotionRotation() {
	if Input.IsActionJustPressed("animate", false) {
		animationNodeStateMachinePlayback.Travel("Animate")
	}
	characterBody3D.AsNode3D().SetQuaternion(Quaternion.Mul(characterBody3D.AsNode3D().Quaternion(), animationTree.AsAnimationMixer().GetRootMotionRotation()))
}
