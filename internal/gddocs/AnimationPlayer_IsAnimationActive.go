/*
var is_paused = not is_playing() and is_animation_active()
var is_stopped = not is_playing() and not is_animation_active()
*/

package main

func AnimationPlayer_IsAnimationActive() {
	var is_paused = !animationPlayer.IsPlaying() && animationPlayer.IsAnimationActive()
	var is_stopped = !animationPlayer.IsPlaying() && !animationPlayer.IsAnimationActive()
	_ = is_paused
	_ = is_stopped
}
