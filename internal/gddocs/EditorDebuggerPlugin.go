/*
[gdscript]
@tool
extends EditorPlugin

class ExampleEditorDebugger extends EditorDebuggerPlugin:

	func _has_capture(capture):
		# Return true if you wish to handle messages with the prefix "my_plugin:".
		return capture == "my_plugin"

	func _capture(message, data, session_id):
		if message == "my_plugin:ping":
			get_session(session_id).send_message("my_plugin:echo", data)
			return true
		return false

	func _setup_session(session_id):
		# Add a new tab in the debugger session UI containing a label.
		var label = Label.new()
		label.name = "Example plugin" # Will be used as the tab title.
		label.text = "Example plugin"
		var session = get_session(session_id)
		# Listens to the session started and stopped signals.
		session.started.connect(func (): print("Session started"))
		session.stopped.connect(func (): print("Session stopped"))
		session.add_session_tab(label)

var debugger = ExampleEditorDebugger.new()

func _enter_tree():
	add_debugger_plugin(debugger)

func _exit_tree():
	remove_debugger_plugin(debugger)
[/gdscript]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/EditorDebuggerPlugin"
	"graphics.gd/classdb/EditorDebuggerSession"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/Label"
)

type ExampleEditorDebuggerPlugin struct {
	EditorPlugin.Extension[ExampleEditorDebuggerPlugin]

	debugger *ExampleEditorDebugger
}

func NewExampleEditorDebuggerPlugin() *ExampleEditorDebuggerPlugin {
	return &ExampleEditorDebuggerPlugin{
		debugger: new(ExampleEditorDebugger),
	}
}

func (p *ExampleEditorDebuggerPlugin) EnterTree() {
	p.AsEditorPlugin().AddDebuggerPlugin(p.debugger.AsEditorDebuggerPlugin())
}

func (p *ExampleEditorDebuggerPlugin) ExitTree() {
	p.AsEditorPlugin().RemoveDebuggerPlugin(p.debugger.AsEditorDebuggerPlugin())
}

type ExampleEditorDebugger struct {
	EditorDebuggerPlugin.Extension[ExampleEditorDebugger]
}

func (d *ExampleEditorDebugger) HasCapture(capture string) bool {
	// Return true if you wish to handle messages with the prefix "my_plugin:".
	return capture == "my_plugin"
}

func (d *ExampleEditorDebugger) Capture(message string, data []any, sessionID int) bool {
	if message == "my_plugin:ping" {
		EditorDebuggerSession.Expanded(d.AsEditorDebuggerPlugin().GetSession(sessionID)).SendMessage("my_plugin:echo", data)
		return true
	}
	return false
}

func (d *ExampleEditorDebugger) SetupSession(sessionID int) {
	// Add a new tab in the debugger session UI containing a label.
	var label = Label.New()
	label.AsNode().SetName("Example plugin") // Will be used as the tab title.
	label.SetText("Example plugin")
	var session = d.AsEditorDebuggerPlugin().GetSession(sessionID)
	// Listens to the session started and stopped signals.
	session.OnStarted(func() { fmt.Println("Session started") })
	session.OnStopped(func() { fmt.Println("Session stopped") })
	session.AddSessionTab(label.AsControl())
}
