package main

import (
	"graphics.gd/classdb/FastNoiseLite"
	"graphics.gd/classdb/Gradient"
	"graphics.gd/classdb/GradientTexture2D"
	"graphics.gd/classdb/MultiMesh"
	"graphics.gd/classdb/NoiseTexture2D"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/SphereMesh"
	"graphics.gd/classdb/TextMesh"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"
)

type Meshes struct {
	Panel.Extension[Meshes] `gd:"CustomMeshes"`

	TextMesh        TextMesh.Instance
	NoiseTexture    NoiseTexture2D.Instance
	GradientTexture GradientTexture2D.Instance
	SphereMesh      SphereMesh.Instance
	MultiMesh       MultiMesh.Instance
}

func (m *Meshes) Ready() {
	m.TextMesh = TextMesh.New()
	m.NoiseTexture = NoiseTexture2D.New()
	m.GradientTexture = GradientTexture2D.New()
	m.SphereMesh = SphereMesh.New()
	m.MultiMesh = MultiMesh.New()

	m.TextMesh.SetText("TextMesh")
	// In 2D, 1 unit equals 1 pixel, so the default size at which PrimitiveMeshes are displayed is tiny.
	// Use much larger mesh size to compensate, or use `DrawSetTransform()` before using `DrawMesh()`
	// to scale the draw command.
	m.TextMesh.SetPixelSize(2.5)

	m.NoiseTexture.SetSeamless(true)
	m.NoiseTexture.SetAsNormalMap(true)
	m.NoiseTexture.SetNoise(FastNoiseLite.New().AsNoise())

	m.GradientTexture.SetGradient(Gradient.New())

	m.SphereMesh.SetHeight(80)
	m.SphereMesh.SetRadius(40)

	m.MultiMesh.SetUseColors(true)
	m.MultiMesh.SetInstanceCount(5)
	m.MultiMesh.SetInstanceTransform2d(0, Transform2D.New())
	m.MultiMesh.SetInstanceColor(0, Color.RGBA{1, 0.7, 0.7, 1})
	m.MultiMesh.SetInstanceTransform2d(1, Transform2D.Translates(Vector2.New(0, 100)))
	m.MultiMesh.SetInstanceColor(1, Color.RGBA{0.7, 1, 0.7, 1})
	m.MultiMesh.SetInstanceTransform2d(2, Transform2D.Translates(Vector2.New(100, 100)))
	m.MultiMesh.SetInstanceColor(2, Color.RGBA{0.7, 0.7, 1, 1})
	m.MultiMesh.SetInstanceTransform2d(3, Transform2D.Translates(Vector2.New(100, 0)))
	m.MultiMesh.SetInstanceColor(3, Color.RGBA{1, 1, 0.7, 1})
	m.MultiMesh.SetInstanceTransform2d(4, Transform2D.Translates(Vector2.New(50, 50)))
	m.MultiMesh.SetInstanceColor(4, Color.RGBA{0.7, 1, 1, 1})
	m.MultiMesh.SetMesh(m.SphereMesh.AsMesh())
}

func (m *Meshes) Draw() {
	var margin = Vector2.New(300, 70)
	var offset = Vector2.Zero
	var canvas = m.AsCanvasItem()

	canvas.MoreArgs().DrawSetTransform(Vector2.Add(margin, offset), 0, Vector2.New(1, -1))
	canvas.DrawMesh(m.TextMesh.AsMesh(), m.NoiseTexture.AsTexture2D())

	offset = Vector2.Add(offset, Vector2.New(150, 0))
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawMesh(m.SphereMesh.AsMesh(), m.NoiseTexture.AsTexture2D())

	offset = Vector2.New(0, 120)
	canvas.DrawSetTransform(Vector2.Add(margin, offset))
	canvas.DrawMultimesh(m.MultiMesh, m.GradientTexture.AsTexture2D())
}
