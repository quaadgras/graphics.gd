package main

import (
	"graphics.gd/shaders"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgba"
	"graphics.gd/shaders/texture"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec4"
)

// Plasma scrolls two noise textures over each other, ported from plasma.gdshader.
type Plasma struct {
	CanvasItem.Shader[Plasma]

	Noise1 texture.Sampler2D[vec4.RGBA] `gd:"noise1,repeat_enable"`
	Noise2 texture.Sampler2D[vec4.RGBA] `gd:"noise2,repeat_enable"`
}

const scrollSpeed = 0.04

func (s *Plasma) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	uv1 := vec2.Add(frag.UV, float.Mul(scrollSpeed, shaders.Time))
	color := s.Noise1.Sample(uv1)
	warp := s.Noise1.Sample(vec2.Mul(uv1, 0.3))
	drift := vec2.Mul(vec2.Mul(vec2.New(warp.R, warp.G), 0.015), float.Mul(shaders.Time, 2.0))
	glow := float.Mul(0.3, s.Noise2.Sample(vec2.Sub(vec2.Mul(frag.UV, 4.0), drift)).R)
	return CanvasItem.Material{
		Color: rgba.Add(color, rgba.New(glow, glow, glow, 0.0)),
	}
}
