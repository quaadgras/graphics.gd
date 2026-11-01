package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/Spatial"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/swizzle"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
)

// Shaders ported from the .gdshader files in the original demo.

type Negative struct {
	Spatial.Shader[Negative]

	TextureAlbedo texture.Sampler2D[vec4.RGBA] `gd:",filter_nearest"`
}

func (*Negative) RenderMode() []Spatial.RenderMode {
	return []Spatial.RenderMode{Spatial.DepthPrepassAlpha, Spatial.Unshaded}
}

func (s *Negative) Material(frag Spatial.Fragment) Spatial.Material {
	albedo := s.TextureAlbedo.Sample(frag.UV)
	return Spatial.Material{
		Albedo: rgb.Sub(rgb.New(1.0, 1.0, 1.0), swizzle.RGB(albedo)),
		Alpha:  float.X(albedo.A),
	}
}

// Paper holds the uniforms shared by the paper shaders, which draw a sprite
// with an outline around its opaque pixels.
type Paper struct {
	TextureAlbedo    texture.Sampler2D[vec4.RGBA] `gd:",source_color,filter_nearest"`
	OutlineThickness float.X                      `gd:",hint_range(0.001, 0.05)"`
	OutlineColor     vec4.RGBA                    `gd:",source_color"`
}

func newPaper() Paper {
	return Paper{
		OutlineThickness: float.New(0.1),
		OutlineColor:     rgba.New(1.0, 1.0, 1.0, 1.0),
	}
}

// PaperRenderMode is the render mode of all of the paper shaders.
var PaperRenderMode = []Spatial.RenderMode{Spatial.DepthDrawOpaque, Spatial.CullDisabled, Spatial.DepthPrepassAlpha}

const threshold = 0.1

// sprite returns the color of the sprite at the current pixel, and whether the
// pixel is part of the outline: transparent, but next to an opaque pixel.
func (p *Paper) sprite(frag Spatial.Fragment) (center vec4.RGBA, outline bool.X) {
	center = p.TextureAlbedo.Sample(frag.UV)
	opaque := func(offset vec2.XY) float.X {
		return float.Step(threshold, p.TextureAlbedo.Sample(vec2.Add(frag.UV, offset)).A)
	}
	width := p.OutlineThickness
	neighbours := float.Add(
		float.Add(opaque(vec2.New(width, 0.0)), opaque(vec2.New(float.Neg(width), 0.0))),
		float.Add(opaque(vec2.New(0.0, width)), opaque(vec2.New(0.0, float.Neg(width)))),
	)
	return center, bool.And(float.Lt(center.A, threshold), float.Gt(neighbours, 0.0))
}

// normal returns the surface normal perturbed by the normal map.
func normal(frag Spatial.Fragment, normalMap texture.Sampler2D[vec4.XYZW]) vec3.XYZ {
	n := normalMap.Sample(frag.UV)
	n3 := vec3.Sub(vec3.Mul(vec3.New(n.X, n.Y, n.Z), 2.0), 1.0)
	return vec3.Normalize(vec3.Add(vec3.Add(
		vec3.Mul(frag.Tangent, n3.X),
		vec3.Mul(frag.Binormal, n3.Y)),
		vec3.Mul(frag.Normal, n3.Z)))
}

type PaperShader struct {
	Spatial.Shader[PaperShader]
	Paper
}

func NewPaperShader() *PaperShader { return &PaperShader{Paper: newPaper()} }

func (*PaperShader) RenderMode() []Spatial.RenderMode { return PaperRenderMode }

func (s *PaperShader) Material(frag Spatial.Fragment) Spatial.Material {
	center, outline := s.sprite(frag)
	return Spatial.Material{
		Albedo: bool.Select(outline, swizzle.RGB(s.OutlineColor), swizzle.RGB(center)),
		Alpha:  bool.Select(outline, float.New(1.0), float.X(center.A)),
	}
}

type PaperMaterial struct {
	Spatial.Shader[PaperMaterial]
	Paper

	NormalMap texture.Sampler2D[vec4.XYZW] `gd:",hint_normal"`
}

func NewPaperMaterial() *PaperMaterial { return &PaperMaterial{Paper: newPaper()} }

func (*PaperMaterial) RenderMode() []Spatial.RenderMode { return PaperRenderMode }

func (s *PaperMaterial) Material(frag Spatial.Fragment) Spatial.Material {
	center, outline := s.sprite(frag)
	return Spatial.Material{
		Normal: normal(frag, s.NormalMap),
		Albedo: bool.Select(outline, swizzle.RGB(s.OutlineColor), swizzle.RGB(center)),
		Alpha:  bool.Select(outline, float.New(1.0), float.X(center.A)),
	}
}

// PaperEmission makes the outline glow.
type PaperEmission struct {
	Spatial.Shader[PaperEmission]
	Paper

	NormalMap texture.Sampler2D[vec4.XYZW] `gd:",hint_normal"`
}

func NewPaperEmission() *PaperEmission { return &PaperEmission{Paper: newPaper()} }

func (*PaperEmission) RenderMode() []Spatial.RenderMode { return PaperRenderMode }

func (s *PaperEmission) Material(frag Spatial.Fragment) Spatial.Material {
	center, outline := s.sprite(frag)
	color := swizzle.RGB(s.OutlineColor)
	return Spatial.Material{
		Normal:   bool.Select(outline, frag.Normal, normal(frag, s.NormalMap)),
		Albedo:   bool.Select(outline, color, swizzle.RGB(center)),
		Alpha:    bool.Select(outline, float.New(1.0), float.X(center.A)),
		Emission: bool.Select(outline, rgb.Mul(color, 4.0), rgb.New(0.0, 0.0, 0.0)),
	}
}

// PaperRainbow cycles the outline through the colors of the rainbow.
type PaperRainbow struct {
	Spatial.Shader[PaperRainbow]
	Paper

	NormalMap texture.Sampler2D[vec4.XYZW] `gd:",hint_normal"`
}

func NewPaperRainbow() *PaperRainbow { return &PaperRainbow{Paper: newPaper()} }

func (*PaperRainbow) RenderMode() []Spatial.RenderMode { return PaperRenderMode }

func (s *PaperRainbow) Material(frag Spatial.Fragment) Spatial.Material {
	center, outline := s.sprite(frag)
	t := float.Add(float.Mul(shaders.Time, 2.0), float.Add(float.Mul(frag.UV.X, 10.0), float.Mul(frag.UV.Y, 10.0)))
	wave := func(phase float64) float.X {
		return float.Add(0.5, float.Mul(0.5, float.Sin(float.Add(t, phase))))
	}
	rainbow := rgb.New(wave(0.0), wave(2.0), wave(4.0))
	return Spatial.Material{
		Normal: bool.Select(outline, frag.Normal, normal(frag, s.NormalMap)),
		Albedo: bool.Select(outline, rainbow, swizzle.RGB(center)),
		Alpha:  bool.Select(outline, float.New(1.0), float.X(center.A)),
	}
}
