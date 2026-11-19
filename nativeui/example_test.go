package nativeui_test

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/nativeui"
)

// Themes a user interface like the platform's own, following the system's
// light/dark mode and accent colour as they change.
func ExampleFollow() {
	ui := Control.New()
	nativeui.Follow(ui)

	save := Button.New()
	save.SetText("Save")
	save.AsControl().SetThemeTypeVariation(nativeui.AccentButton) // the default action
	ui.AsNode().AddChild(save.AsNode())
}

// Builds a theme for a given look, whatever the platform (with the
// nativeui_all build tag, every look is compiled in).
func ExampleNew() {
	ui := Control.New()
	ui.SetTheme(nativeui.New(nativeui.Options{Look: nativeui.Windows, Dark: true, Scale: 1}))
}
