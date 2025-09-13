/*
[gdscript]
func _ready():
    var button = Button.new()
    button.text = "Click me"
    button.pressed.connect(_button_pressed)
    add_child(button)

func _button_pressed():
    print("Hello world!")
[/gdscript]
[csharp]
public override void _Ready()
{
    var button = new Button();
    button.Text = "Click me";
    button.Pressed += ButtonPressed;
    AddChild(button);
}

private void ButtonPressed()
{
    GD.Print("Hello world!");
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Node"
)

type ExampleForButton struct {
	Node.Extension[ExampleForButton]
}

func (eg *ExampleForButton) Ready() {
	button := Button.New()
	button.SetText("Click me")
	button.AsBaseButton().OnPressed(func() {
		print("Hello world!")
	})
	eg.AsNode().AddChild(button.AsNode())
}
