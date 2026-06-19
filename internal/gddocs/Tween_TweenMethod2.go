/*
[gdscript]
func _ready():
	var tween = create_tween()
	tween.tween_method(set_label_text, 0, 10, 1.0).set_delay(1.0)

func set_label_text(value: int):
	$Label.text = "Counting " + str(value)
[/gdscript]
[csharp]
public override void _Ready()
{
	base._Ready();

	Tween tween = CreateTween();
	tween.TweenMethod(Callable.From<int>(SetLabelText), 0.0f, 10.0f, 1.0f).SetDelay(1.0f);
}

private void SetLabelText(int value)
{
	GetNode<Label>("Label").Text = $"Counting {value}";
}
[/csharp]
*/

package main

import (
	"strconv"

	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/MethodTweener"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Callable"
)

type tweenMethodLabel struct {
	Node.Extension[tweenMethodLabel]

	Label Label.Instance
}

func (n tweenMethodLabel) Ready() {
	var tween = n.AsNode().CreateTween()
	MethodTweener.Make(tween, Callable.New(n.SetLabelText), 0, 10, 1).SetDelay(1.0)
}

func (n tweenMethodLabel) SetLabelText(value int) {
	n.Label.SetText("Counting " + strconv.Itoa(value))
}
