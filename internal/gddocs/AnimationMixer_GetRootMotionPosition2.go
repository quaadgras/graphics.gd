/*
[gdscript]
func _process(delta):
	if Input.is_action_just_pressed("animate"):
		state_machine.travel("Animate")
	set_quaternion(get_quaternion() * animation_tree.get_root_motion_rotation())
	var velocity = (animation_tree.get_root_motion_rotation_accumulator().inverse() * get_quaternion()) * animation_tree.get_root_motion_position() / delta
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

type rootMotionExample2 struct {
	CharacterBody3D.Extension[rootMotionExample2]

	StateMachine  AnimationNodeStateMachinePlayback.Instance
	AnimationTree AnimationTree.Instance
}

func (n rootMotionExample2) Process(delta Float.X) {
	if Input.IsActionJustPressed("animate", false) {
		n.StateMachine.Travel("Animate")
	}
	n.AsNode3D().SetQuaternion(Quaternion.Mul(n.AsNode3D().Quaternion(), n.AnimationTree.AsAnimationMixer().GetRootMotionRotation()))
	var velocity = Vector3.Div(
		Quaternion.Rotate(
			n.AnimationTree.AsAnimationMixer().GetRootMotionPosition(),
			Quaternion.Mul(Quaternion.Inverse(n.AnimationTree.AsAnimationMixer().GetRootMotionRotationAccumulator()), n.AsNode3D().Quaternion()),
		),
		Vector3.New(delta, delta, delta),
	)
	n.AsCharacterBody3D().SetVelocity(velocity)
	n.AsCharacterBody3D().MoveAndSlide()
}
