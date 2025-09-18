/*
[gdscript]
extends Node

var save_shortcut = Shortcut.new()
func _ready():
	var key_event = InputEventKey.new()
	key_event.keycode = KEY_S
	key_event.ctrl_pressed = true
	key_event.command_or_control_autoremap = true # Swaps Ctrl for Command on Mac.
	save_shortcut.events = [key_event]

func _input(event):
	if save_shortcut.matches_event(event) and event.is_pressed() and not event.is_echo():
		print("Save shortcut pressed!")
		get_viewport().set_input_as_handled()
[/gdscript]
[csharp]
using Godot;

public partial class MyNode : Node
{
	private readonly Shortcut _saveShortcut = new Shortcut();

	public override void _Ready()
	{
		InputEventKey keyEvent = new InputEventKey
		{
			Keycode = Key.S,
			CtrlPressed = true,
			CommandOrControlAutoremap = true, // Swaps Ctrl for Command on Mac.
		};

		_saveShortcut.Events = [keyEvent];
	}

	public override void _Input(InputEvent @event)
	{
		if (@event is InputEventKey keyEvent &&
			_saveShortcut.MatchesEvent(@event) &&
			keyEvent.Pressed && !keyEvent.Echo)
		{
			GD.Print("Save shortcut pressed!");
			GetViewport().SetInputAsHandled();
		}
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Shortcut"
	"graphics.gd/classdb/Viewport"
)

type NodeWithShortcut struct {
	Node.Extension[NodeWithShortcut]

	save_shortcut Shortcut.Instance
}

func (n *NodeWithShortcut) Ready() {
	n.save_shortcut = Shortcut.New()
	var key_event = InputEventKey.New()
	key_event.SetKeycode(Input.KeyS)
	key_event.AsInputEventWithModifiers().SetCtrlPressed(true)
	key_event.AsInputEventWithModifiers().SetCommandOrControlAutoremap(true) // Swaps Ctrl for Command on Mac.
	n.save_shortcut.SetEvents([]InputEvent.Instance{key_event.AsInputEvent()})
}

func (n *NodeWithShortcut) Input(event InputEvent.Instance) {
	if n.save_shortcut.MatchesEvent(event) && event.IsPressed() && !event.IsEcho() {
		fmt.Println("Save shortcut pressed!")
		Viewport.Get(n.AsNode()).SetInputAsHandled()
	}
}
