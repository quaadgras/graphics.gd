/*
[gdscript]
var command_palette = EditorInterface.get_command_palette()
# external_command is a function that will be called with the command is executed.
var command_callable = Callable(self, "external_command").bind(arguments)
command_palette.add_command("command", "test/command",command_callable)
[/gdscript]
[csharp]
EditorCommandPalette commandPalette = EditorInterface.Singleton.GetCommandPalette();
// ExternalCommand is a function that will be called with the command is executed.
Callable commandCallable = new Callable(this, MethodName.ExternalCommand);
commandPalette.AddCommand("command", "test/command", commandCallable)
[/csharp]
*/

package main

import "graphics.gd/classdb/EditorInterface"

func ExampleEditorCommandPalette() {
	var command_palette = EditorInterface.GetCommandPalette()
	// external_command is a function that will be called with the command is executed.
	command_palette.AddCommand("command", "test/command", func() {
		// do something
	})
}
