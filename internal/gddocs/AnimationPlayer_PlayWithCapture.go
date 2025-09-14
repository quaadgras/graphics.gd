/*
capture(name, duration, trans_type, ease_type)
play(name, custom_blend, custom_speed, from_end)
*/

package main

import (
	"graphics.gd/classdb/AnimationMixer"
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Float"
)

var animationMixer AnimationMixer.Instance
var animationPlayer AnimationPlayer.Instance

type duration = Float.X
type name = string
type custom_speed = Float.X
type custom_blend = Float.X
type from_end = bool

func AnimationPlayer_PlayWithCapture() {
	AnimationMixer.Expanded(animationMixer).Capture(name("run"), duration(0.2), Tween.TransitionType(0), Tween.EaseType(0))
	AnimationPlayer.Expanded(animationPlayer).PlayWithCapture(name("run"), duration(0.5), custom_blend(1), custom_speed(1), from_end(false), Tween.TransitionType(0), Tween.EaseType(0))
}
