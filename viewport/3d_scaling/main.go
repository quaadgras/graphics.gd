package main

import (
	"fmt"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/ClassDB"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
	"graphics.gd/variant/String"
)

// HUD scales the 3D viewport rendering without affecting 2D elements such as the HUD.
type HUD struct {
	Control.Extension[HUD] `gd:"HUD"`

	Scale  Label.Instance `gd:"VBoxContainer/Scale"`
	Filter Label.Instance `gd:"VBoxContainer/Filter"`

	// scaleFactor is the 3D viewport's shrink factor. For instance, 1 is full resolution,
	// 2 is half resolution and 4 is quarter resolution. Lower values look
	// sharper but are slower to render.
	scaleFactor int
	filterMode  Viewport.Scaling3DMode

	viewport Viewport.Instance
}

func NewHUD() *HUD {
	return &HUD{
		scaleFactor: 1,
		filterMode:  Viewport.Scaling3dModeBilinear,
	}
}

func (h *HUD) Ready() {
	h.viewport = SceneTree.Get(h.AsNode()).Root().AsViewport()
	h.viewport.SetScaling3dMode(h.filterMode)
}

func (h *HUD) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("cycle_viewport_resolution") {
		h.scaleFactor = Int.Wrap(h.scaleFactor+1, 1, 5)
		h.viewport.SetScaling3dScale(1.0 / Float.X(h.scaleFactor))
		h.Scale.SetText(fmt.Sprintf("Scale: %3.0f%%", 100.0/float64(h.scaleFactor)))
	}

	if event.IsActionPressed("toggle_filtering") {
		h.filterMode = Int.Wrap(h.filterMode+1, Viewport.Scaling3dModeBilinear, Viewport.Scaling3dModeMax)
		h.viewport.SetScaling3dMode(h.filterMode)
		text := String.Capitalize(ClassDB.ClassGetEnumConstants("Viewport", "Scaling3DMode", false)[h.filterMode])
		text = strings.ReplaceAll(text, "3d", "3D")
		text = strings.ReplaceAll(text, "Mode", "Mode:")
		text = strings.ReplaceAll(text, "Fsr", "FSR")
		h.Filter.SetText(text)
	}
}

func main() {
	classdb.Register[HUD](NewHUD)
	startup.Scene()
}
