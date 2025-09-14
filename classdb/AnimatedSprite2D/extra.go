package AnimatedSprite2D

// PlayNamed plays the named animation from the beginning.
func (self Instance) PlayNamed(name string) {
	Expanded(self).Play(name, 1.0, false)
}
