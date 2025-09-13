/*
[gdscript]
var state_machine = $AnimationTree.get("parameters/playback")
state_machine.travel("some_state")
[/gdscript]
[csharp]
var stateMachine = GetNode<AnimationTree>("AnimationTree").Get("parameters/playback") as AnimationNodeStateMachinePlayback;
stateMachine.Travel("some_state");
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AnimationNodeStateMachinePlayback"
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/variant/Object"
)

func ExampleAnimationNodeStateMachine(tree AnimationTree.Instance) {
	var stateMachine = Object.Get(tree, "parameters/playback").(AnimationNodeStateMachinePlayback.Instance)
	stateMachine.Travel("some_state")
}
