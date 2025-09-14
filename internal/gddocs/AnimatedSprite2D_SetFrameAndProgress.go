/*
[gdscript]
var current_frame = animated_sprite.get_frame()
var current_progress = animated_sprite.get_frame_progress()
animated_sprite.play("walk_another_skin")
animated_sprite.set_frame_and_progress(current_frame, current_progress)
[/gdscript]
*/

package main

import "graphics.gd/classdb/AnimatedSprite2D"

var animated_sprite AnimatedSprite2D.Instance

func AnimatedSprite2D_SetFrameAndProgress() {
	var current_frame = animated_sprite.Frame()
	var current_progress = animated_sprite.FrameProgress()
	animated_sprite.PlayNamed("walk_another_skin")
	animated_sprite.SetFrameAndProgress(current_frame, current_progress)
}
