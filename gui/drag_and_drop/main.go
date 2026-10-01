package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/ColorPickerButton"
	"graphics.gd/classdb/Control"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

// DragDropColorPicker is a color picker button whose color can be dragged
// onto other pickers.
type DragDropColorPicker struct {
	ColorPickerButton.Extension[DragDropColorPicker] `gd:"DragDropColorPicker"`
}

// GetDragData returns the data to pass from an object when you click and drag
// away from this object. Also calls SetDragPreview to show the mouse dragging
// something so the user knows that the operation is working.
func (d *DragDropColorPicker) GetDragData(_ Vector2.XY) any {
	color := d.AsColorPickerButton().Color()

	// Use another colorpicker as drag preview.
	cpb := ColorPickerButton.New()
	cpb.SetColor(color)
	cpb.AsControl().SetSize(Vector2.XY{X: 80, Y: 50})

	// Allows us to center the color picker on the mouse.
	preview := Control.New()
	preview.AsNode().AddChild(cpb.AsNode())
	cpb.AsControl().SetPosition(Vector2.MulX(cpb.AsControl().Size(), -0.5))

	// Sets what the user will see they are dragging.
	d.AsControl().SetDragPreview(preview)

	// Return color as drag data.
	return color
}

// CanDropData returns a boolean by examining the data being dragged to see if
// it's valid to drop here.
func (d *DragDropColorPicker) CanDropData(_ Vector2.XY, data any) bool {
	_, ok := data.(Color.RGBA)
	return ok
}

// DropData takes the data being dragged and processes it. In this case, we are
// assigning a new color to the target color picker button.
func (d *DragDropColorPicker) DropData(_ Vector2.XY, data any) {
	if color, ok := data.(Color.RGBA); ok {
		d.AsColorPickerButton().SetColor(color)
	}
}

func main() {
	classdb.Register[DragDropColorPicker]()
	startup.Scene()
}
