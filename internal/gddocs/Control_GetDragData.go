/*
[gdscript]
func _get_drag_data(position):
    var mydata = make_data() # This is your custom method generating the drag data.
    set_drag_preview(make_preview(mydata)) # This is your custom method generating the preview of the drag data.
    return mydata
[/gdscript]
[csharp]
public override Variant _GetDragData(Vector2 atPosition)
{
    var myData = MakeData(); // This is your custom method generating the drag data.
    SetDragPreview(MakePreview(myData)); // This is your custom method generating the preview of the drag data.
    return myData;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/variant/Vector2"
)

func make_data() any                    { return nil }
func make_preview(any) Control.Instance { return Control.Nil }

func Control_GetDragData() {
	GetDragData := func(position Vector2.XY) any {
		var mydata = make_data() // This is your custom method generating the drag data.
		control.SetDragPreview(make_preview(mydata))
		return mydata
	}
	_ = GetDragData
}
