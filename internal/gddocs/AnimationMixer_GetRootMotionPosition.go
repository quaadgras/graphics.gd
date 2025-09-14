/*
[gdscript]
var current_rotation

func _process(delta):
    if Input.is_action_just_pressed("animate"):
        current_rotation = get_quaternion()
        state_machine.travel("Animate")
    var velocity = current_rotation * animation_tree.get_root_motion_position() / delta
    set_velocity(velocity)
    move_and_slide()
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/AnimationNodeStateMachinePlayback"
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Vector3"
)

var delta Float.X
var current_rotation Quaternion.IJKX

var animationTree AnimationTree.Instance
var animationNodeStateMachinePlayback AnimationNodeStateMachinePlayback.Instance
var characterBody3D CharacterBody3D.Instance

func AnimationMixer_GetRootMotionPosition() {
	if Input.IsActionJustPressed("animate", false) {
		current_rotation = characterBody3D.AsNode3D().Quaternion()
		animationNodeStateMachinePlayback.Travel("Animate")
	}
	var velocity = Vector3.DivX(Quaternion.Rotate(animationTree.AsAnimationMixer().GetRootMotionPosition(), current_rotation), delta)
	characterBody3D.SetVelocity(velocity)
	characterBody3D.MoveAndSlide()
}
