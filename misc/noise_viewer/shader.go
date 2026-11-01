package main

import (
	"graphics.gd/shaders/bool"
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/CanvasItem"
	"graphics.gd/shaders/rgba"
)

// NoiseViewerShader shows values below MinValue as black and above MaxValue as
// white, ported from noise_viewer.gdshader.
type NoiseViewerShader struct {
	CanvasItem.Shader[NoiseViewerShader]

	MinValue float.X
	MaxValue float.X
}

func NewNoiseViewerShader() *NoiseViewerShader {
	return &NoiseViewerShader{MinValue: float.New(-1.0), MaxValue: float.New(1.0)}
}

func (s *NoiseViewerShader) Material(frag CanvasItem.Fragment) CanvasItem.Material {
	color := frag.Texture().Sample(frag.UV)
	gray := color.R
	return CanvasItem.Material{
		Color: bool.Select(float.Lt(gray, s.MinValue), rgba.New(0.0, 0.0, 0.0, 1.0),
			bool.Select(float.Gt(gray, s.MaxValue), rgba.New(1.0, 1.0, 1.0, 1.0), color)),
	}
}

