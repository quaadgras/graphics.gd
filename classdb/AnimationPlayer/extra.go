package AnimationPlayer

import "graphics.gd/variant/StringName"

func (self Instance) PlayNamed(anim_name string) {
	class(self).Play(StringName.New(anim_name), float64(-1), float64(1.0), false)
}
