package main

import (
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/StyleBoxFlat"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"
)

type Rectangles struct {
	Panel.Extension[Rectangles] `gd:"CustomRectangles"`

	antialiasable
}

func (r *Rectangles) Draw() {
	var margin = Vector2.New(200, 40)
	var line_width_thin = r.LineWidth()
	var antialiasing_width_offset = r.WidthOffset()
	var offset = Vector2.Zero
	var canvas = r.AsCanvasItem()

	/*
		 * draw_rect(
					Rect2(margin + offset, Vector2(100, 50)),
					Color.PURPLE,
					false,
					line_width_thin,
					use_antialiasing
			)

			offset += Vector2(120, 0)
			draw_rect(
					Rect2(margin + offset, Vector2(100, 50)),
					Color.PURPLE,
					false,
					2.0 - antialiasing_width_offset,
					use_antialiasing
			)

			offset += Vector2(120, 0)
			draw_rect(
					Rect2(margin + offset, Vector2(100, 50)),
					Color.PURPLE,
					false,
					6.0 - antialiasing_width_offset,
					use_antialiasing
			)

			# Draw a filled rectangle. The width parameter is ignored for filled rectangles (it's set to `-1.0` to avoid warnings).
			# We also reduce the rectangle's size by half the antialiasing width offset.
			# Otherwise, the rectangle becomes very slightly larger when draw antialiasing is enabled.
			offset += Vector2(120, 0)
			draw_rect(
					Rect2(margin + offset, Vector2(100, 50)).grow(-antialiasing_width_offset * 0.5),
					Color.PURPLE,
					true,
					-1.0,
					use_antialiasing
			)

			# `draw_set_transform()` is a stateful command: it affects *all* `draw_` methods within this
			# `_draw()` function after it. This can be used to translate, rotate or scale `draw_` methods
			# that don't offer dedicated parameters for this (such as `draw_rect()` not having a rotation parameter).
			# To reset back to the initial transform, call `draw_set_transform(Vector2())`.
			offset += Vector2(170, 0)
			draw_set_transform(margin + offset, deg_to_rad(22.5))
			draw_rect(
					Rect2(Vector2(), Vector2(100, 50)),
					Color.PURPLE,
					false,
					line_width_thin,
					use_antialiasing
			)
			offset += Vector2(120, 0)
			draw_set_transform(margin + offset, deg_to_rad(22.5))
			draw_rect(
					Rect2(Vector2(), Vector2(100, 50)),
					Color.PURPLE,
					false,
					2.0 - antialiasing_width_offset,
					use_antialiasing
			)
			offset += Vector2(120, 0)
			draw_set_transform(margin + offset, deg_to_rad(22.5))
			draw_rect(
					Rect2(Vector2(), Vector2(100, 50)),
					Color.PURPLE,
					false,
					6.0 - antialiasing_width_offset,
					use_antialiasing
			)

			# `draw_set_transform_matrix()` is a more advanced counterpart of `draw_set_transform()`.
			# It can be used to apply transforms that are not supported by `draw_set_transform()`, such as
			# skewing.
			offset = Vector2(20, 60)
			var custom_transform := get_transform().translated(margin + offset)
			# Perform horizontal skewing (the rectangle will appear "slanted").
			custom_transform.y.x -= 0.5
			draw_set_transform_matrix(custom_transform)
			draw_rect(
				Rect2(Vector2(), Vector2(100, 50)),
				Color.PURPLE,
				false,
				6.0 - antialiasing_width_offset,
				use_antialiasing
			)
			draw_set_transform(Vector2())

			offset = Vector2(0, 250)
			var style_box_flat := StyleBoxFlat.new()
			style_box_flat.set_border_width_all(4)
			style_box_flat.set_corner_radius_all(8)
			style_box_flat.shadow_size = 1
			style_box_flat.shadow_offset = Vector2(4, 4)
			style_box_flat.shadow_color = Color.RED
			style_box_flat.anti_aliasing = use_antialiasing
			draw_style_box(style_box_flat, Rect2(margin + offset, Vector2(100, 50)))

			offset += Vector2(130, 0)
			var style_box_flat_2 := StyleBoxFlat.new()
			style_box_flat_2.draw_center = false
			style_box_flat_2.set_border_width_all(4)
			style_box_flat_2.set_corner_radius_all(8)
			style_box_flat_2.corner_detail = 1
			style_box_flat_2.border_color = Color.GREEN
			style_box_flat_2.anti_aliasing = use_antialiasing
			draw_style_box(style_box_flat_2, Rect2(margin + offset, Vector2(100, 50)))

			offset += Vector2(160, 0)
			var style_box_flat_3 := StyleBoxFlat.new()
			style_box_flat_3.draw_center = false
			style_box_flat_3.set_border_width_all(4)
			style_box_flat_3.set_corner_radius_all(8)
			style_box_flat_3.border_color = Color.CYAN
			style_box_flat_3.shadow_size = 40
			style_box_flat_3.shadow_offset = Vector2()
			style_box_flat_3.shadow_color = Color.CORNFLOWER_BLUE
			style_box_flat_3.anti_aliasing = use_antialiasing
			custom_transform = get_transform().translated(margin + offset)
			# Perform vertical skewing (the rectangle will appear "slanted").
			custom_transform.x.y -= 0.5
			draw_set_transform_matrix(custom_transform)
			draw_style_box(style_box_flat_3, Rect2(Vector2(), Vector2(100, 50)))

			draw_set_transform(Vector2())
	*/
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		line_width_thin,
		r.UseAntialiasing,
	)
	offset = Vector2.Add(offset, Vector2.New(120, 0))
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		2.0-antialiasing_width_offset,
		r.UseAntialiasing,
	)
	offset = Vector2.Add(offset, Vector2.New(120, 0))
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		6.0-antialiasing_width_offset,
		r.UseAntialiasing,
	)
	// Draw a filled rectangle. The width parameter is ignored for filled rectangles (it's set to `-1.0` to avoid warnings).
	// We also reduce the rectangle's size by half the antialiasing width offset.
	// Otherwise, the rectangle becomes very slightly larger when draw antialiasing is enabled.
	offset = Vector2.Add(offset, Vector2.New(120, 0))
	canvas.MoreArgs().DrawRect(
		Rect2.Expand(Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(100, 50)}, -antialiasing_width_offset*0.5),
		Color.W3C.Purple,
		true,
		-1.0,
		r.UseAntialiasing,
	)
	// DrawSetTransform() is a stateful command: it affects *all* `draw_` methods within this
	// `Draw()` function after it. This can be used to translate, rotate or scale `Draw()` methods
	// that don't offer dedicated parameters for this (such as `DrawRect()` not having a rotation parameter).
	// To reset back to the initial transform, call `DrawSetTransform(Vector2.Zero)`.
	offset = Vector2.Add(offset, Vector2.New(170, 0))
	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), Angle.InRadians(22.5), Vector2.One)
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Zero, Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		line_width_thin,
		r.UseAntialiasing,
	)
	offset = Vector2.Add(offset, Vector2.New(120, 0))
	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), Angle.InRadians(22.5), Vector2.One)
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Zero, Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		2.0-antialiasing_width_offset,
		r.UseAntialiasing,
	)
	offset = Vector2.Add(offset, Vector2.New(120, 0))
	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), Angle.InRadians(22.5), Vector2.One)
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Zero, Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		6.0-antialiasing_width_offset,
		r.UseAntialiasing,
	)

	// DrawSetTransformMatrix() is a more advanced counterpart of DrawSetTransform().
	// It can be used to apply transforms that are not supported by DrawSetTransform(), such as
	// skewing.
	offset = Vector2.New(20, 60)
	var custom_transform = Transform2D.Translated(r.AsCanvasItem().GetTransform(), Vector2.Add(margin, offset))
	// Perform horizontal skewing (the rectangle will appear "slanted").
	custom_transform.Y.X -= 0.5
	canvas.DrawSetTransformMatrix(custom_transform)
	canvas.MoreArgs().DrawRect(
		Rect2.PositionSize{Vector2.Zero, Vector2.New(100, 50)},
		Color.W3C.Purple,
		false,
		6.0-antialiasing_width_offset,
		r.UseAntialiasing,
	)
	canvas.DrawSetTransform(Vector2.Zero)

	offset = Vector2.New(0, 250)
	var style_box_flat = StyleBoxFlat.New()
	style_box_flat.SetBorderWidthAll(4)
	style_box_flat.SetCornerRadiusAll(8)
	style_box_flat.SetShadowSize(1)
	style_box_flat.SetShadowOffset(Vector2.New(4, 4))
	style_box_flat.SetShadowColor(Color.W3C.Red)
	style_box_flat.SetAntiAliasing(r.UseAntialiasing)
	style_box_flat.AsStyleBox().Draw(canvas, Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(100, 50)})

	offset = Vector2.Add(offset, Vector2.New(130, 0))
	var style_box_flat_2 = StyleBoxFlat.New()
	style_box_flat_2.SetDrawCenter(false)
	style_box_flat_2.SetBorderWidthAll(4)
	style_box_flat_2.SetCornerRadiusAll(8)
	style_box_flat_2.SetCornerDetail(1)
	style_box_flat_2.SetBorderColor(Color.W3C.Green)
	style_box_flat_2.SetAntiAliasing(r.UseAntialiasing)
	style_box_flat_2.AsStyleBox().Draw(canvas, Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(100, 50)})

	offset = Vector2.Add(offset, Vector2.New(160, 0))
	var style_box_flat_3 = StyleBoxFlat.New()
	style_box_flat_3.SetDrawCenter(false)
	style_box_flat_3.SetBorderWidthAll(4)
	style_box_flat_3.SetCornerRadiusAll(8)
	style_box_flat_3.SetBorderColor(Color.W3C.Cyan)
	style_box_flat_3.SetShadowSize(40)
	style_box_flat_3.SetShadowOffset(Vector2.Zero)
	style_box_flat_3.SetShadowColor(Color.W3C.CornflowerBlue)
	style_box_flat_3.SetAntiAliasing(r.UseAntialiasing)
	custom_transform = Transform2D.Translated(r.AsCanvasItem().GetTransform(), Vector2.Add(margin, offset))
	// Perform vertical skewing (the rectangle will appear "slanted").
	custom_transform.X.Y -= 0.5
	canvas.DrawSetTransformMatrix(custom_transform)
	style_box_flat_3.AsStyleBox().Draw(canvas, Rect2.PositionSize{Vector2.Zero, Vector2.New(100, 50)})

	canvas.DrawSetTransform(Vector2.Zero)
}
