package main

import (
	"fmt"

	"graphics.gd/classdb/Label"
)

type ScoreLabel struct {
	Label.Extension[ScoreLabel]

	score int
}

func (s *ScoreLabel) OnMobSquashed() {
	s.score++
	s.AsLabel().SetText(fmt.Sprintf("Score: %d", s.score))
}
