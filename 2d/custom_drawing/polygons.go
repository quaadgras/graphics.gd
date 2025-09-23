package main

import (
	"graphics.gd/classdb/Panel"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

type Polygons struct {
	Panel.Extension[Polygons] `gd:"CustomPolygons"`

	antialiasable
}

func (p *Polygons) Draw() {
	var margin = Vector2.New(240, 40)
	var line_width_thin = p.LineWidth()
	var antialiasing_width_offset = p.WidthOffset()
	var points = []Vector2.XY{
		{0, 0}, {0, 60}, {60, 90}, {60, 0}, {40, 25}, {10, 40},
	}
	var colors = []Color.RGBA{
		Color.W3C.White, Color.W3C.Red, Color.W3C.Green, Color.W3C.Blue, Color.W3C.Magenta, Color.W3C.Magenta,
	}
	var offset = Vector2.Zero
	var canvas = p.AsCanvasItem()
	// `DrawSetTransform()` is a stateful command: it affects *all* `Draw` methods within this
	// `Draw()` function after it. This can be used to translate, rotate or scale `Draw` methods
	// that don't offer dedicated parameters for this (such as `DrawPrimitive()` not having a position parameter).
	// To reset back to the initial transform, call `DrawSetTransform(Vector2.Zero)`.
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawPrimitive(points[0:1], colors[0:1], nil)

	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawPrimitive(points[0:2], colors[0:2], nil)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawPrimitive(points[0:3], colors[0:3], nil)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawPrimitive(points[0:4], colors[0:4], nil)

	// Draws a polygon with multiple colors that are interpolated between each point.
	// Colors are specified in the same order as points' positions, but in a different array.
	offset = Vector2.New(0, 120)
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawPolygon(points, colors)

	// Draw a polygon with a single color. Only a points array is needed in this case.
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawColoredPolygon(points, Color.W3C.Yellow)

	// Draw a polygon-based line. Each segment is connected to the previous one, which means
	// `DrawPolyline()` always draws a contiguous line.
	offset = Vector2.New(0, 240)
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawPolyline(points, Color.W3C.SkyBlue, line_width_thin, p.UseAntialiasing)

	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawPolyline(points, Color.W3C.SkyBlue, 2.0-antialiasing_width_offset, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawPolyline(points, Color.W3C.SkyBlue, 6.0-antialiasing_width_offset, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawPolylineColors(points, colors, line_width_thin, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawPolylineColors(points, colors, 2.0-antialiasing_width_offset, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawPolylineColors(points, colors, 6.0-antialiasing_width_offset, p.UseAntialiasing)

	// Draw multiple lines in a single draw command. Unlike `DrawPolyline()`,
	// lines are not connected to the last segment.
	// This is faster than calling `DrawLine()` several times and should be preferred
	// when drawing dozens of lines or more at once.
	offset = Vector2.New(0, 360)
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawMultiline(points, Color.W3C.SkyBlue, line_width_thin, p.UseAntialiasing)

	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawMultiline(points, Color.W3C.SkyBlue, 2.0-antialiasing_width_offset, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawMultiline(points, Color.W3C.SkyBlue, 6.0-antialiasing_width_offset, p.UseAntialiasing)

	// `DrawMultilineColors()` makes it possible to draw lines of different colors in a single
	// draw command, although gradients are not possible this way (unlike `DrawPolygon()` and `DrawPolyline()`).
	// This means the number of supplied colors must be equal to the number of segments, which means
	// we have to only pass a subset of the colors array in this example.
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawMultilineColors(points, colors[0:3], line_width_thin, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawMultilineColors(points, colors[0:3], 2.0-antialiasing_width_offset, p.UseAntialiasing)
	offset = Vector2.Add(offset, Vector2.New(90, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.MoreArgs().DrawMultilineColors(points, colors[0:3], 6.0-antialiasing_width_offset, p.UseAntialiasing)
}
