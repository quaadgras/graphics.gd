package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/StyleBoxFlat"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Object"
)

// ThemingOverride demonstrates how to alter StyleBoxes at runtime.
// Custom theme item properties aren't considered Object properties per se.
// This means that you should use AddThemeStyleboxOverride("normal", ...)
// instead of Set("custom_styles/normal", ...).
type ThemingOverride struct {
	Control.Extension[ThemingOverride] `gd:"ThemingOverride"`

	Label          Label.Instance  `gd:"Panel/MarginContainer/VBoxContainer/Label"`
	Button         Button.Instance `gd:"Panel/MarginContainer/VBoxContainer/Button"`
	Button2        Button.Instance `gd:"Panel/MarginContainer/VBoxContainer/Button2"`
	ResetAllButton Button.Instance `gd:"Panel/MarginContainer/VBoxContainer/ResetAllButton"`
}

func (t *ThemingOverride) Ready() {
	// Focus the first button automatically for keyboard/controller-friendly navigation.
	t.Button.AsControl().GrabFocus()
}

// overrideBorder duplicates the button's normal/hover/pressed styleboxes and
// gives them the requested border color. We have to modify all three at once
// to get a correct appearance when the button is hovered or pressed, and we
// can't share a single StyleBox as they have different background colors.
func overrideBorder(button Button.Instance, border Color.RGBA) {
	control := button.AsControl()
	for _, name := range []string{"normal", "hover", "pressed"} {
		stylebox, ok := Object.As[StyleBoxFlat.Instance](control.GetThemeStylebox(name).AsResource().Duplicate())
		if !ok {
			continue
		}
		stylebox.SetBorderColor(border)
		control.AddThemeStyleboxOverride(name, stylebox.AsStyleBox())
	}
}

func (t *ThemingOverride) OnButtonPressed() {
	overrideBorder(t.Button, Color.RGBA{R: 1, G: 1, B: 0, A: 1})
	t.Label.AsControl().AddThemeColorOverride("font_color", Color.RGBA{R: 1, G: 1, B: 0.375, A: 1})
}

func (t *ThemingOverride) OnButton2Pressed() {
	overrideBorder(t.Button2, Color.RGBA{R: 0, G: 1, B: 0.5, A: 1})
	t.Label.AsControl().AddThemeColorOverride("font_color", Color.RGBA{R: 0.375, G: 1, B: 0.75, A: 1})
}

func (t *ThemingOverride) OnResetAllButtonPressed() {
	for _, button := range []Button.Instance{t.Button, t.Button2} {
		for _, name := range []string{"normal", "hover", "pressed"} {
			button.AsControl().RemoveThemeStyleboxOverride(name)
		}
	}
	t.Label.AsControl().RemoveThemeColorOverride("font_color")
}

func main() {
	classdb.Register[ThemingOverride]()
	startup.Scene()
}
