/*
# Add a custom shortcut for a plugin action.
var my_shortcut = Shortcut.new()
var input_event = InputEventKey.new()
input_event.keycode = KEY_F5
input_event.ctrl_pressed = true
my_shortcut.events.append(input_event)

# This will appear under the "My Plugin" category as "Reload Data".
EditorInterface.get_editor_settings().add_shortcut("my_plugin/reload_data", my_shortcut)

# This will appear under the "Test Action" category as "Test Action".
EditorInterface.get_editor_settings().add_shortcut("test_action", my_shortcut)
*/

package main

import (
	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Shortcut"
)

func ExampleEditorSettingsAddShortcut() {
	// Add a custom shortcut for a plugin action.
	var myShortcut = Shortcut.New()
	var inputEvent = InputEventKey.New()
	inputEvent.SetKeycode(Input.KeyF5)
	inputEvent.AsInputEventWithModifiers().SetCtrlPressed(true)
	myShortcut.SetEvents([]InputEvent.Instance{inputEvent.AsInputEvent()})

	// This will appear under the "My Plugin" category as "Reload Data".
	EditorInterface.GetEditorSettings().AddShortcut("my_plugin/reload_data", myShortcut)

	// This will appear under the "Test Action" category as "Test Action".
	EditorInterface.GetEditorSettings().AddShortcut("test_action", myShortcut)
}
