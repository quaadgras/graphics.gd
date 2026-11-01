package main

import (
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec4"
)

// Shaders ported from the .gdshader files in the original demo. Many of them
// sample the texture at a ring of offsets around each pixel, so the shared
// helpers below replace what was copy-pasted lines in the shading language.

// Offsets of the neighbouring pixels, in units of the effect's radius.
var (
	cross = [][2]float64{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	ring  = [][2]float64{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
)

// neighbour samples the sprite's texture at offset*radius pixels away from the current pixel.
func neighbour(frag CanvasItem.Fragment, offset [2]float64, radius float.X) vec4.RGBA {
	shift := vec2.Mul(vec2.Mul(vec2.New(offset[0], offset[1]), radius), frag.TexturePixelSize())
	return frag.Texture().Sample(vec2.Add(frag.UV, shift))
}

// sum adds up the neighbouring samples at each offset.
func sum(frag CanvasItem.Fragment, offsets [][2]float64, radius float.X) vec4.RGBA {
	total := neighbour(frag, offsets[0], radius)
	for _, offset := range offsets[1:] {
		total = rgba.Add(total, neighbour(frag, offset, radius))
	}
	return total
}

// edge returns how much the alpha changes across the cross of neighbouring
// pixels, which is 1 on the boundary of the sprite and 0 inside or outside of it.
func edge(frag CanvasItem.Fragment, col vec4.RGBA, width float.X) float.X {
	maxa, mina := float.X(col.A), float.X(col.A)
	for _, offset := range cross {
		a := neighbour(frag, offset, width).A
		maxa = float.Max(a, maxa)
		mina = float.Min(a, mina)
	}
	return float.Sub(maxa, mina)
}

// premultiply multiplies the color channels by the alpha channel.
func premultiply(col vec4.RGBA) vec4.RGBA {
	return rgba.Mul(col, rgba.New(col.A, col.A, col.A, 1.0))
}

type Outline struct {
	CanvasItem.Shader[Outline]

	OutlineWidth float.X
	OutlineColor vec4.RGBA `gd:",source_color"`
}

func (s *Outline) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	col := frag.Texture().Sample(frag.UV)
	return CanvasItem.Material{
		Color: rgba.Mix(col, s.OutlineColor, edge(frag, col, s.OutlineWidth)),
	}
}

// Aura only works properly with premultiplied alpha blend mode.
type Aura struct {
	CanvasItem.Shader[Aura]

	AuraWidth float.X
	AuraColor vec4.RGBA `gd:",source_color"`
}

func (*Aura) RenderMode() []CanvasItem.RenderMode {
	return []CanvasItem.RenderMode{CanvasItem.BlendPremulAlpha}
}

func (s *Aura) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	col := frag.Texture().Sample(frag.UV)
	glow := float.X(edge(frag, col, s.AuraWidth))
	return CanvasItem.Material{
		Color: rgba.Add(premultiply(col), rgba.New(
			float.Mul(s.AuraColor.R, glow),
			float.Mul(s.AuraColor.G, glow),
			float.Mul(s.AuraColor.B, glow),
			0.0,
		)),
	}
}

type Blur struct {
	CanvasItem.Shader[Blur]

	Radius float.X
}

func (s *Blur) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	col := rgba.Add(frag.Texture().Sample(frag.UV), sum(frag, cross, s.Radius))
	return CanvasItem.Material{
		Color: rgba.Div(col, 5.0),
	}
}

type Fatty struct {
	CanvasItem.Shader[Fatty]

	Fattyness float.X
}

func (s *Fatty) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	ruv := vec2.Sub(frag.UV, 0.5)
	dir := vec2.Normalize(ruv)
	length := float.Mul(float.Pow(float.Mul(vec2.Length(ruv), 2.0), s.Fattyness), 0.5)
	return CanvasItem.Material{
		Color: frag.Texture().Sample(vec2.Add(vec2.Mul(dir, length), 0.5)),
	}
}

type DropShadow struct {
	CanvasItem.Shader[DropShadow]

	Radius   float.X
	Modulate vec4.RGBA `gd:",source_color"`
}

func (s *DropShadow) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	shadow := rgba.Mul(rgba.Div(sum(frag, ring, s.Radius), 8.0), s.Modulate)
	col := frag.Texture().Sample(frag.UV)
	return CanvasItem.Material{
		Color: rgba.Mix(shadow, col, col.A),
	}
}

type OffsetShadow struct {
	CanvasItem.Shader[OffsetShadow]

	Offset   vec2.XY
	Modulate vec4.RGBA `gd:",source_color"`
}

func (s *OffsetShadow) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	behind := frag.Texture().Sample(vec2.Sub(frag.UV, vec2.Mul(s.Offset, frag.TexturePixelSize())))
	shadow := rgba.New(s.Modulate.R, s.Modulate.G, s.Modulate.B, float.Mul(behind.A, s.Modulate.A))
	col := frag.Texture().Sample(frag.UV)
	return CanvasItem.Material{
		Color: rgba.Mix(shadow, col, col.A),
	}
}

type Silhouette struct {
	CanvasItem.Shader[Silhouette]

	Modulate vec4.RGBA `gd:",source_color"`
}

func (s *Silhouette) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	a := frag.Texture().Sample(frag.UV).A
	return CanvasItem.Material{
		Color: rgba.New(s.Modulate.R, s.Modulate.G, s.Modulate.B, float.Mul(a, s.Modulate.A)),
	}
}

type Glow struct {
	CanvasItem.Shader[Glow]

	Radius float.X
	Amount float.X
}

func (*Glow) RenderMode() []CanvasItem.RenderMode {
	return []CanvasItem.RenderMode{CanvasItem.BlendPremulAlpha}
}

func (s *Glow) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	col := frag.Texture().Sample(frag.UV)
	glow := rgba.Add(col, sum(frag, ring, s.Radius))
	glow = rgba.Add(glow, sum(frag, ring, float.Mul(s.Radius, 2.0)))
	glow = rgba.Mul(rgba.Div(glow, 17.0), s.Amount)
	return CanvasItem.Material{
		Color: rgba.Add(glow, premultiply(col)),
	}
}

type Disintegrate struct {
	CanvasItem.Shader[Disintegrate]

	Amount float.X
}

func (s *Disintegrate) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	noise := float.Fract(float.Mul(float.Sin(vec2.Dot(frag.UV, vec2.New(12.9898, 78.233))), 438.5453))
	col := frag.Texture().Sample(frag.UV)
	return CanvasItem.Material{
		Color: rgba.New(col.R, col.G, col.B, float.Mul(col.A, float.Pow(noise, s.Amount))),
	}
}
