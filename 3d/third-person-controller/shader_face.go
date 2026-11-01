package main

import (
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

// FaceShader shows a face (or eyes) on a pixelated screen, with the color
// channels offset from each other. Ported from face_mat.gdshader and the
// identical screen_shader.gdshader.
type FaceShader struct {
	Spatial.Shader[FaceShader]

	FaceTexture       texture.Sampler2D[vec4.XYZW] `gd:",filter_nearest,repeat_disable"`
	Intensity         float.X
	ScreenColor       vec4.RGBA `gd:",source_color"`
	ScreenRedOffset   vec2.XY
	ScreenGreenOffset vec2.XY
	ScreenBlueOffset  vec2.XY
	PixelSize         float.X
}

func NewFaceShader() *FaceShader {
	return &FaceShader{
		Intensity:         float.New(1.5),
		ScreenColor:       rgba.New(1.0, 1.0, 1.0, 1.0),
		ScreenRedOffset:   vec2.New(0.0, 0.0),
		ScreenGreenOffset: vec2.New(0.0, 0.0),
		ScreenBlueOffset:  vec2.New(0.0, 0.0),
		PixelSize:         float.New(20.0),
	}
}

func (s *FaceShader) Material(frag Spatial.Fragment) Spatial.Material {
	scaledUV := vec2.Div(vec2.Round(vec2.Mul(frag.UV, s.PixelSize)), s.PixelSize)
	channel := func(offset vec2.XY) float.X {
		return float.Sub(1.0, s.FaceTexture.Sample(vec2.Add(offset, scaledUV)).X)
	}
	grid2D := vec2.Sin(vec2.Mul(vec2.Fract(vec2.Add(vec2.Mul(frag.UV, s.PixelSize), 0.5)), 3.14))
	grid := float.Add(float.Mul(float.Mul(grid2D.X, grid2D.Y), 0.6), 0.4)
	iris := vec3.RGB{
		R: float.Div(float.Mul(s.ScreenColor.R, channel(s.ScreenRedOffset)), 3.0),
		G: float.Div(float.Mul(s.ScreenColor.G, channel(s.ScreenGreenOffset)), 3.0),
		B: float.Div(float.Mul(s.ScreenColor.B, channel(s.ScreenBlueOffset)), 3.0),
	}
	iris = rgb.Mul(iris, float.Sub(1.0, s.ScreenColor.A))
	fill := rgb.Mul(swizzle.RGB(s.ScreenColor), s.ScreenColor.A)
	return Spatial.Material{
		Albedo:    rgb.New(0.0, 0.0, 0.0),
		Specular:  float.New(0.25),
		Roughness: float.New(0.45),
		Emission:  rgb.Mul(rgb.Mul(rgb.Add(iris, fill), grid), s.Intensity),
	}
}
