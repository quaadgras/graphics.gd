package main

import (
	"fmt"

	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Label"
	"graphics.gd/variant/Float"
)

type FPSLabel struct {
	Label.Extension[FPSLabel]
}

func (l *FPSLabel) Process(_ Float.X) {
	fps := Engine.GetFramesPerSecond()
	l.AsLabel().SetText(fmt.Sprintf("%d FPS (%.2f mspf)", int(fps), 1000.0/fps))
}
