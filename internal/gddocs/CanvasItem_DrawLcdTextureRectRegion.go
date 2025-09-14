/*
dst.r = texture.r * modulate.r * modulate.a + dst.r * (1.0 - texture.r * modulate.a);
dst.g = texture.g * modulate.g * modulate.a + dst.g * (1.0 - texture.g * modulate.a);
dst.b = texture.b * modulate.b * modulate.a + dst.b * (1.0 - texture.b * modulate.a);
dst.a = modulate.a + dst.a * (1.0 - modulate.a);
*/

package main

import (
	"graphics.gd/variant/Color"
)

func CanvasItem_DrawLcdTextureRectRegion() {
	var dst, texture, modulate Color.RGBA
	dst.R = texture.R*modulate.R*modulate.A + dst.R*(1.0-texture.R*modulate.A)
	dst.G = texture.G*modulate.G*modulate.A + dst.G*(1.0-texture.G*modulate.A)
	dst.B = texture.B*modulate.B*modulate.A + dst.B*(1.0-texture.B*modulate.A)
	dst.A = modulate.A + dst.A*(1.0-modulate.A)
}
