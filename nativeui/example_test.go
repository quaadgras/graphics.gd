package nativeui_test

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/nativeui"
)

// Every control in the program looks like the platform's own, and follows
// the system's light/dark mode and accent colour as they change.
func ExampleUse() {
	nativeui.Use()

	save := Button.New()
	save.SetText("Save")
	save.AsControl().SetThemeTypeVariation("AccentButton") // the default action
}

// A theme with the system's native look, for one part of the interface.
func ExampleNew() {
	panel := Control.New()
	panel.SetTheme(nativeui.New())
}
