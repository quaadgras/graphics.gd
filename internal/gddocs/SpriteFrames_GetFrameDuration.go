/*
absolute_duration = relative_duration / (animation_fps * abs(playing_speed))
*/

package main

import "graphics.gd/variant/Float"

var (
	relative_duration, animation_fps, playing_speed Float.X
)

func SpriteFrames_GetFrameDuration() {
	absolute_duration := relative_duration / (animation_fps * Float.Abs(playing_speed))
	_ = absolute_duration
}
