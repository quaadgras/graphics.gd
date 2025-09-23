package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/TabContainer"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type Antialiaser interface {
	CanvasItem.Any

	SetUseAntialiasing(bool)
	LineWidth() Float.X
}

type antialiasable struct {
	UseAntialiasing bool
}

func (a *antialiasable) SetUseAntialiasing(v bool) {
	a.UseAntialiasing = v
}

func (a *antialiasable) LineWidth() Float.X {
	// Line width of `-1.0` is only usable with draw antialiasing disabled,
	// as it uses hardware line drawing as opposed to polygon-based line drawing.
	// Automatically use polygon-based line drawing when needed to avoid runtime warnings.
	// We also use a line width of `0.5` instead of `1.0` to better match the appearance
	// of non-antialiased line drawing, as draw antialiasing tends to make lines look thicker.
	if a.UseAntialiasing {
		return 0.5
	}
	return -1
}

func (a *antialiasable) WidthOffset() Float.X {
	if a.UseAntialiasing {
		return 1.0
	}
	return 0.0
}

type CustomDrawing struct {
	Control.Extension[CustomDrawing]

	TabContainer   TabContainer.Instance `gd:"%TabContainer"`
	AnimationSlice *AnimationSlice       `gd:"%AnimationSlice"`
}

func (c *CustomDrawing) OnMsaa2dItemSelected(index int) {
	Viewport.Get(c.AsNode()).SetMsaa2d(Viewport.MSAA(index))
}

func (c *CustomDrawing) OnDrawAntialiasingToggled(pressed bool) {
	var nodes = c.TabContainer.AsNode().GetChildren()
	nodes = append(nodes, c.AnimationSlice.AsNode())
	for _, tab := range nodes {
		if tab, ok := Object.As[Antialiaser](tab); ok {
			tab.SetUseAntialiasing(pressed)
			tab.AsCanvasItem().QueueRedraw()
		}
	}

}

func main() {
	classdb.Register[Animation]()
	classdb.Register[AnimationSlice]()
	classdb.Register[CustomDrawing]()
	classdb.Register[Lines]()
	classdb.Register[Polygons]()
	classdb.Register[Meshes]()
	classdb.Register[Rectangles]()
	classdb.Register[Text]()
	classdb.Register[Textures]()
	startup.Scene()
}
