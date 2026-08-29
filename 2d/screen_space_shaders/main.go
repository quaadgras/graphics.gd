package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/startup"
	"graphics.gd/variant/Object"
)

type ScreenShaders struct {
	Control.Extension[ScreenShaders]

	Effect   OptionButton.Instance
	Effects  Control.Instance
	Picture  OptionButton.Instance
	Pictures Control.Instance
}

func (s *ScreenShaders) Ready() {
	for _, c := range s.Pictures.AsNode().GetChildren() {
		s.Picture.AddItem("PIC: " + c.Name())
	}
	for _, c := range s.Effects.AsNode().GetChildren() {
		s.Effect.AddItem("FX: " + c.Name())
	}
}

func showOnly(parent Control.Instance, id int) {
	for i, c := range parent.AsNode().GetChildren() {
		if item, ok := Object.As[CanvasItem.Instance](c); ok {
			item.SetVisible(i == id)
		}
	}
}

func (s *ScreenShaders) OnPictureItemSelected(id int) { showOnly(s.Pictures, id) }
func (s *ScreenShaders) OnEffectItemSelected(id int)  { showOnly(s.Effects, id) }

func main() {
	classdb.Register[ScreenShaders]()
	startup.Scene()
}
