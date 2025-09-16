/*
[gdscript]
var expression = Expression.new()

func _ready():
	$LineEdit.text_submitted.connect(self._on_text_submitted)

func _on_text_submitted(command):
	var error = expression.parse(command)
	if error != OK:
		print(expression.get_error_text())
		return
	var result = expression.execute()
	if not expression.has_execute_failed():
		$LineEdit.text = str(result)
[/gdscript]
[csharp]
private Expression _expression = new Expression();

public override void _Ready()
{
	GetNode<LineEdit>("LineEdit").TextSubmitted += OnTextEntered;
}

private void OnTextEntered(string command)
{
	Error error = _expression.Parse(command);
	if (error != Error.Ok)
	{
		GD.Print(_expression.GetErrorText());
		return;
	}
	Variant result = _expression.Execute();
	if (!_expression.HasExecuteFailed())
	{
		GetNode<LineEdit>("LineEdit").Text = result.ToString();
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Expression"
	"graphics.gd/classdb/LineEdit"
)

type ExampleExpression struct {
	LineEdit LineEdit.Instance

	expression Expression.Instance
}

func (e *ExampleExpression) Ready() {
	e.expression = Expression.New()
	e.LineEdit.OnTextSubmitted(func(new_text string) {
		var err = e.expression.Parse(new_text)
		if err != nil {
			fmt.Println(e.expression.GetErrorText())
			return
		}
		var result = e.expression.Execute()
		if !e.expression.HasExecuteFailed() {
			e.LineEdit.SetText(fmt.Sprint(result))
		}
	})
}
