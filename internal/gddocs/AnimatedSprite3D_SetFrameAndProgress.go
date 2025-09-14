/*
[gdscript]
var current_frame = animated_sprite.get_frame()
var current_progress = animated_sprite.get_frame_progress()
animated_sprite.play("walk_another_skin")
animated_sprite.set_frame_and_progress(current_frame, current_progress)
[/gdscript]
*/

package main

import "graphics.gd/classdb/AnimatedSprite3D"

var animated_sprite3D AnimatedSprite3D.Instance

func AnimatedSprite3D_SetFrameAndProgress() {
	var current_frame = animated_sprite3D.Frame()
	var current_progress = animated_sprite3D.FrameProgress()
	animated_sprite3D.PlayNamed("walk_another_skin")
	animated_sprite3D.SetFrameAndProgress(current_frame, current_progress)
}
