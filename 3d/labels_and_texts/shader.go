package main

import (
	"graphics.gd/shaders/float"
	"graphics.gd/shaders/pipeline/Spatial"
	"graphics.gd/shaders/rgb"
	"graphics.gd/shaders/vec3"
)

// Curvature bends the mesh around the Y axis, ported from curvature.gdshader.
type Curvature struct {
	Spatial.Shader[Curvature]
}

func (*Curvature) Fragment(vertex Spatial.Vertex) Spatial.Fragment {
	pos := vertex.Position
	bend := float.Cos(pos.X)
	return Spatial.Fragment{
		VertexPosition: vec3.New(
			float.Mul(pos.X, float.Add(0.8, float.Mul(bend, 0.2))),
			pos.Y,
			float.Sub(pos.Z, float.Sub(0.6, float.Mul(bend, 0.6))),
		),
	}
}

func (*Curvature) Material(Spatial.Fragment) Spatial.Material {
	return Spatial.Material{
		Albedo: rgb.New(1.0, 0.6, 0.1),
	}
}
