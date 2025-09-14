/*
[gdscript]
var global_library = mixer.get_animation_library("")
global_library.add_animation("animation_name", animation_resource)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/Animation"
	"graphics.gd/classdb/AnimationMixer"
)

var mixer AnimationMixer.Instance
var animation_resource Animation.Instance

func AnimationMixer_AddAnimationLibrary() {
	var global_library = mixer.GetAnimationLibrary("")
	global_library.AddAnimation("animation_name", animation_resource)
}
