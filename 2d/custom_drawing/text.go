package main

import (
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/TextServer"
	"graphics.gd/classdb/TextServerManager"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector2"
)

type Text struct {
	Panel.Extension[Text] `gd:"CustomText"`
}

func (t *Text) Draw() {
	var font = t.AsControl().GetThemeDefaultFont()
	const FontSize = 24
	const String = "Hello world!"
	var margin = Vector2.New(240, 60)

	/*
		 * var offset := Vector2()
			var advance := Vector2()
			for character in STRING:
				# Draw each character with a random pastel color.
				# Notice how the advance calculated on the loop's previous iteration is used as an offset here.
				draw_char(font, margin + offset + advance, character, FONT_SIZE, Color.from_hsv(randf(), 0.4, 1.0))

				# Get the glyph index of the character we've just drawn, so we can retrieve the glyph advance.
				# This determines the spacing between glyphs so the next character is positioned correctly.
				var glyph_idx := TextServerManager.get_primary_interface().font_get_glyph_index(
							get_theme_default_font().get_rids()[0],
							FONT_SIZE,
							character.unicode_at(0),
							0
				)
				advance.x += TextServerManager.get_primary_interface().font_get_glyph_advance(
						get_theme_default_font().get_rids()[0],
						FONT_SIZE,
						glyph_idx
				).x

			offset += Vector2(0, 32)
			# When drawing a font outline, it must be drawn *before* the main text.
			# This way, the outline appears behind the main text.
			draw_string_outline(
					font,
					margin + offset,
					STRING,
					HORIZONTAL_ALIGNMENT_LEFT,
					-1,
					FONT_SIZE,
					12,
					Color.ORANGE.darkened(0.6)
			)
			# NOTE: Use `draw_multiline_string()` to draw strings that contain line breaks (`\n`) or with
			# automatic line wrapping based on the specified width.
			# A width of `-1` is used here, which means "no limit". If width is limited, the end of the string
			# will be cut off if it doesn't fit within the specified width.
			draw_string(
					font,
					margin + offset,
					STRING,
					HORIZONTAL_ALIGNMENT_LEFT,
					-1,
					FONT_SIZE,
					Color.YELLOW
			)
	*/
	var canvas = t.AsCanvasItem()
	var offset, advance Vector2.XY
	for _, character := range String {
		// Draw each character with a random pastel color.
		// Notice how the advance calculated on the loop's previous iteration is used as an offset here.
		canvas.MoreArgs().DrawChar(font, Vector2.Add(Vector2.Add(margin, offset), advance), string(character), FontSize, Color.HSV(Float.Random(), 0.4, 1.0), 0)
		// Get the glyph index of the character we've just drawn, so we can retrieve the glyph advance.
		// This determines the spacing between glyphs so the next character is positioned correctly.
		var glyph_idx = TextServer.Advanced(TextServerManager.GetPrimaryInterface()).FontGetGlyphIndex(
			RID.Any(font.GetRids()[0]),
			FontSize,
			int64(character),
			0,
		)
		advance.X += TextServer.Advanced(TextServerManager.GetPrimaryInterface()).FontGetGlyphAdvance(
			RID.Any(font.GetRids()[0]),
			FontSize,
			glyph_idx,
		).X
	}
	offset = Vector2.Add(offset, Vector2.New(0, 32))
	// When drawing a font outline, it must be drawn *before* the main text.
	// This way, the outline appears behind the main text.
	canvas.MoreArgs().DrawStringOutline(
		font,
		Vector2.Add(margin, offset),
		String,
		GUI.HorizontalAlignmentLeft,
		-1,
		FontSize,
		12,
		Color.Darkened(Color.W3C.Orange, 0.6),
		TextServer.JustificationKashida|TextServer.JustificationWordBound,
		0, 0, 0,
	)
	// NOTE: Use `DrawMultilineString()` to draw strings that contain line breaks (`\n`) or with
	// automatic line wrapping based on the specified width.
	// A width of `-1` is used here, which means "no limit". If width is limited, the end of the string
	// will be cut off if it doesn't fit within the specified width.
	canvas.MoreArgs().DrawString(
		font,
		Vector2.Add(margin, offset),
		String,
		GUI.HorizontalAlignmentLeft,
		-1,
		FontSize,
		Color.W3C.Yellow,
		TextServer.JustificationKashida|TextServer.JustificationWordBound,
		0, 0, 0,
	)
}
