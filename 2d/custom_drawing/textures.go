package main

import (
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

type Textures struct {
	Panel.Extension[Textures] `gd:"CustomTextures"`
}

var Icon = Resource.Load[Texture2D.Instance]("res://icon.svg")

func (t *Textures) Draw() {
	var margin = Vector2.New(260, 40)
	var offset = Vector2.Zero
	var canvas = t.AsCanvasItem()
	canvas.MoreArgs().DrawTexture(Icon, Vector2.Add(margin, offset), Color.W3C.White)

	// Draw a rotated texture at half the scale of its original pixel size.
	offset = Vector2.Add(offset, Vector2.New(200, 20))
	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), Angle.InRadians(45), Vector2.New(0.5, 0.5))
	canvas.MoreArgs().DrawTexture(Icon, Vector2.Zero, Color.W3C.White)
	canvas.DrawSetTransform(Vector2.Zero)

	// Draw a stretched texture. In this example, the icon is 128×128 so it will be drawn at 2× scale.
	offset = Vector2.Add(offset, Vector2.New(70, -20))
	canvas.MoreArgs().DrawTextureRect(Icon,
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(256, 256)}, false, Color.X11.Green, false,
	)

	// Draw a tiled texture. In this example, the icon is 128×128 so it will be drawn twice on each axis.
	offset = Vector2.Add(offset, Vector2.New(270, 0))
	canvas.MoreArgs().DrawTextureRect(Icon,
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(256, 256)}, true, Color.X11.Green, false,
	)

	offset = Vector2.New(0, 300)
	canvas.MoreArgs().DrawTextureRectRegion(Icon,
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(128, 128)},
		Rect2.PositionSize{Vector2.New(32, 32), Vector2.New(64, 64)}, Color.W3C.Violet, false, false,
	)

	// Draw a tiled texture from a region that is larger than the original texture size (128×128).
	// Transposing is enabled, which will rotate the image by 90 degrees counter-clockwise.
	// (For more advanced transforms, use `DrawSetTransform()` before calling `DrawTextureRectRegion()`.)
	//
	// For tiling to work with this approach, the CanvasItem in which this `Draw()` method is called
	// must have its Repeat property set to Enabled in the inspector.
	offset = Vector2.Add(offset, Vector2.New(140, 0))
	canvas.MoreArgs().DrawTextureRectRegion(Icon,
		Rect2.PositionSize{Vector2.Add(margin, offset), Vector2.New(128, 128)},
		Rect2.PositionSize{Vector2.Zero, Vector2.New(512, 512)}, Color.W3C.Violet, true, false,
	)
}
