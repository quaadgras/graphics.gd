/*
[gdscript]
extends Node

func _ready():
    EngineDebugger.register_message_capture("my_plugin", _capture)
    EngineDebugger.send_message("my_plugin:ping", ["test"])

func _capture(message, data):
    # Note that the "my_plugin:" prefix is not used here.
    if message == "echo":
        prints("Echo received:", data)
        return true
    return false
[/gdscript]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/EngineDebugger"
	"graphics.gd/classdb/Node"
)

type ExampleDebuggerPlugin struct {
	Node.Extension[ExampleDebuggerPlugin]
}

func (e *ExampleDebuggerPlugin) Ready() {
	EngineDebugger.RegisterMessageCapture("my_plugin", func(message string, data []any) bool {
		// Note that the "my_plugin:" prefix is not used here.
		if message == "echo" {
			fmt.Println("Echo received:", data)
			return true
		}
		return false
	})
}
