package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/HBoxContainer"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
)

type PolygonsLines struct {
	Node2D.Extension[PolygonsLines]

	MSAA             HBoxContainer.Instance
	UnsupportedLabel Label.Instance
}

func (p *PolygonsLines) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		p.MSAA.AsCanvasItem().SetVisible(false)
		p.UnsupportedLabel.AsCanvasItem().SetVisible(true)
	}
}

func (p *PolygonsLines) OnMsaaOptionButtonItemSelected(index int) {
	Viewport.Get(p.AsNode()).SetMsaa2d(Viewport.MSAA(index))
}

func main() {
	classdb.Register[PolygonsLines]()
	startup.Scene()
}
