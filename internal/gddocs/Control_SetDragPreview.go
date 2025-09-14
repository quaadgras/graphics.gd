/*
[gdscript]
@export var color = Color(1, 0, 0, 1)

func _get_drag_data(position):
    # Use a control that is not in the tree
    var cpb = ColorPickerButton.new()
    cpb.color = color
    cpb.size = Vector2(50, 50)
    set_drag_preview(cpb)
    return color
[/gdscript]
[csharp]
[Export]
private Color _color = new Color(1, 0, 0, 1);

public override Variant _GetDragData(Vector2 atPosition)
{
    // Use a control that is not in the tree
    var cpb = new ColorPickerButton();
    cpb.Color = _color;
    cpb.Size = new Vector2(50, 50);
    SetDragPreview(cpb);
    return _color;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/ColorPickerButton"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

func Control_SetDragPreview() {
	GetDragData := func(position Vector2.XY) any {
		// Use a control that is not in the tree
		var cpb = ColorPickerButton.New()
		cpb.SetColor(Color.RGBA{1, 0, 0, 1})
		cpb.AsControl().SetSize(Vector2.XY{50, 50})
		control.SetDragPreview(cpb.AsControl())
		return cpb.Color()
	}
	_ = GetDragData
}
