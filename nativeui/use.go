package nativeui

import (
	"sync"

	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Theme"
	"graphics.gd/classdb/ThemeDB"
	"graphics.gd/variant/Object"
)

var using struct {
	sync.Mutex
	pristine  Theme.Instance // Godot's own default theme, as it was
	customize func(*options)
	hooked    bool
}

// Use makes the native look the default theme of every Control in the
// program (popups and dialogs included), and keeps it current as the
// system's theme changes: dark mode, accent colour, contrast. Controls with
// a theme of their own, or a project theme, still use that.
//
// Call Use once the engine has started (in main, or from any node's Ready).
func Use() { useWith(nil) }

// useWith is Use, with customize adjusting the detected options each time
// (to force a look, for the gallery and reference renders).
func useWith(customize func(*options)) {
	using.Lock()
	using.customize = customize
	if using.pristine == Theme.Nil {
		// The default theme also styles the controls nativeui doesn't
		// (Tree, ItemList, GraphEdit...): keep a copy to rebuild on.
		pristine := Object.To[Theme.Instance](ThemeDB.GetDefaultTheme().AsResource().Duplicate())
		// Its fonts and font sizes are set on each control type, which
		// would win over the native look's default font: drop them.
		for _, typ := range pristine.GetFontTypeList() {
			for _, name := range pristine.GetFontList(typ) {
				pristine.ClearFont(name, typ)
			}
		}
		for _, typ := range pristine.GetFontSizeTypeList() {
			for _, name := range pristine.GetFontSizeList(typ) {
				pristine.ClearFontSize(name, typ)
			}
		}
		using.pristine = pristine
	}
	hook := !using.hooked
	using.hooked = true
	using.Unlock()
	apply()
	if hook {
		DisplayServer.SetSystemThemeChangeCallback(apply)
	}
}

// apply refills Godot's default theme in place, which every control using
// it is notified of: Godot's own items, then the native look's on top.
func apply() {
	using.Lock()
	customize, pristine := using.customize, using.pristine
	using.Unlock()
	opts := detect()
	if customize != nil {
		customize(&opts)
	}
	native := newTheme(opts)
	theme := ThemeDB.GetDefaultTheme()
	theme.Clear()
	theme.MergeWith(pristine)
	theme.MergeWith(native)
}
