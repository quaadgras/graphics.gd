package main

import (
	"graphics.gd/classdb/Panel"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

type Lines struct {
	Panel.Extension[Lines] `gd:"CustomLines"`

	antialiasable
}

func (l *Lines) Draw() {
	var margin = Vector2.New(200, 50)
	var line_width_thin = l.LineWidth()
	var antialiasing_width_offset = l.WidthOffset()
	var offset = Vector2.Zero
	var line_length = Vector2.New(140, 35)
	var canvas = l.AsCanvasItem()

	canvas.MoreArgs().DrawLine(Vector2.Add(margin, offset), Vector2.Add(Vector2.Add(margin, offset), line_length), Color.W3C.Green, line_width_thin, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(line_length.X+15, 0))
	canvas.MoreArgs().DrawLine(Vector2.Add(margin, offset), Vector2.Add(Vector2.Add(margin, offset), line_length), Color.W3C.Green, 2.0-antialiasing_width_offset, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(line_length.X+15, 0))
	canvas.MoreArgs().DrawLine(Vector2.Add(margin, offset), Vector2.Add(Vector2.Add(margin, offset), line_length), Color.W3C.Green, 6.0-antialiasing_width_offset, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(line_length.X+15, 0))
	canvas.MoreArgs().DrawDashedLine(Vector2.Add(margin, offset), Vector2.Add(Vector2.Add(margin, offset), line_length), Color.W3C.Cyan, line_width_thin, 5.0, true, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(line_length.X+15, 0))
	canvas.MoreArgs().DrawDashedLine(Vector2.Add(margin, offset), Vector2.Add(Vector2.Add(margin, offset), line_length), Color.W3C.Cyan, 2.0-antialiasing_width_offset, 10.0, true, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(line_length.X+15, 0))
	canvas.MoreArgs().DrawDashedLine(Vector2.Add(margin, offset), Vector2.Add(Vector2.Add(margin, offset), line_length), Color.W3C.Cyan, 6.0-antialiasing_width_offset, 15.0, true, l.UseAntialiasing)

	offset = Vector2.New(40, 120)
	canvas.MoreArgs().DrawCircle(Vector2.Add(margin, offset), 40, Color.W3C.Orange, false, line_width_thin, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawCircle(Vector2.Add(margin, offset), 40, Color.W3C.Orange, false, 2.0-antialiasing_width_offset, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawCircle(Vector2.Add(margin, offset), 40, Color.W3C.Orange, false, 6.0-antialiasing_width_offset, l.UseAntialiasing)
	// Draw a filled circle. The width parameter is ignored for filled circles (it's set to `-1.0` to avoid warnings).
	// We also reduce the radius by half the antialiasing width offset.
	// Otherwise, the circle becomes very slightly larger when draw antialiasing is enabled.
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawCircle(Vector2.Add(margin, offset), 40-antialiasing_width_offset*0.5, Color.W3C.Orange, true, -1.0, l.UseAntialiasing)

	// `draw_set_transform()` is a stateful command: it affects *all* `draw_` methods within this
	// `_draw()` function after it. This can be used to translate, rotate or scale `draw_` methods
	// that don't offer dedicated parameters for this (such as `draw_primitive()` not having a position parameter).
	// To reset back to the initial transform, call `draw_set_transform(Vector2())`.
	//
	// Draw an horizontally stretched circle.
	offset = Vector2.Add(offset, Vector2.New(200, 0))
	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), 0.0, Vector2.New(3.0, 1.0))
	canvas.MoreArgs().DrawCircle(Vector2.Zero, 40, Color.W3C.Orange, false, line_width_thin, l.UseAntialiasing)
	canvas.DrawSetTransform(Vector2.Zero)

	// Draw a quarter circle (`TAU` represents a full turn in radians).
	const PointCountHigh = 24
	offset = Vector2.New(0, 200)
	canvas.MoreArgs().DrawArc(Vector2.Add(margin, offset), 60, 0, 0.25*Angle.Tau, PointCountHigh, Color.W3C.Yellow, line_width_thin, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawArc(Vector2.Add(margin, offset), 60, 0, 0.25*Angle.Tau, PointCountHigh, Color.W3C.Yellow, 2.0-antialiasing_width_offset, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawArc(Vector2.Add(margin, offset), 60, 0, 0.25*Angle.Tau, PointCountHigh, Color.W3C.Yellow, 6.0-antialiasing_width_offset, l.UseAntialiasing)

	// Draw a three quarters of a circle with a low point count, which gives it an angular look.
	const PointCountLow = 7
	offset = Vector2.Add(offset, Vector2.New(125, 30))
	canvas.MoreArgs().DrawArc(Vector2.Add(margin, offset), 40, -0.25*Angle.Tau, 0.5*Angle.Tau, PointCountLow, Color.W3C.Yellow, line_width_thin, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawArc(Vector2.Add(margin, offset), 40, -0.25*Angle.Tau, 0.5*Angle.Tau, PointCountLow, Color.W3C.Yellow, 2.0-antialiasing_width_offset, l.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(100, 0))
	canvas.MoreArgs().DrawArc(Vector2.Add(margin, offset), 40, -0.25*Angle.Tau, 0.5*Angle.Tau, PointCountLow, Color.W3C.Yellow, 6.0-antialiasing_width_offset, l.UseAntialiasing)

	// Draw an horizontally stretched arc.
	offset = Vector2.Add(offset, Vector2.New(200, 0))
	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), 0.0, Vector2.New(3.0, 1.0))
	canvas.MoreArgs().DrawArc(Vector2.Zero, 40, -0.25*Angle.Tau, 0.5*Angle.Tau, PointCountLow, Color.W3C.Yellow, line_width_thin, l.UseAntialiasing)
	canvas.DrawSetTransform(Vector2.Zero)
}
