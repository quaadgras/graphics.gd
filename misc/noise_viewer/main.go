package main

import (
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/FastNoiseLite"
	"graphics.gd/classdb/NoiseTexture2D"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/ShaderMaterial"
	"graphics.gd/classdb/SpinBox"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/VBoxContainer"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

// NoiseViewer lets the user tweak the parameters of a FastNoiseLite texture.
type NoiseViewer struct {
	Control.Extension[NoiseViewer]

	SeamlessNoiseTexture TextureRect.Instance

	ParameterContainer struct {
		VBoxContainer.Instance

		SeedSpinBox              SpinBox.Instance
		FrequencySpinBox         SpinBox.Instance
		FractalOctavesSpinBox    SpinBox.Instance
		FractalGainSpinBox       SpinBox.Instance
		FractalLacunaritySpinBox SpinBox.Instance
	}

	noise FastNoiseLite.Instance // The FastNoiseLite object.

	// Various noise parameters.
	minNoise Float.X
	maxNoise Float.X
}

func NewNoiseViewer() *NoiseViewer {
	return &NoiseViewer{minNoise: -1, maxNoise: 1}
}

func (v *NoiseViewer) Ready() {
	texture := Object.To[NoiseTexture2D.Instance](v.SeamlessNoiseTexture.Texture())
	v.noise = Object.To[FastNoiseLite.Instance](texture.Noise())

	// Set up noise with basic info.
	p := &v.ParameterContainer
	p.SeedSpinBox.AsRange().SetValue(Float.X(v.noise.Seed()))
	p.FrequencySpinBox.AsRange().SetValue(v.noise.Frequency())
	p.FractalOctavesSpinBox.AsRange().SetValue(Float.X(v.noise.FractalOctaves()))
	p.FractalGainSpinBox.AsRange().SetValue(v.noise.FractalGain())
	p.FractalLacunaritySpinBox.AsRange().SetValue(v.noise.FractalLacunarity())

	// Render the noise.
	v.refreshShaderParams()
}

func (v *NoiseViewer) refreshShaderParams() {
	// Adjust min/max for shader.
	min := (v.minNoise + 1) / 2
	max := (v.maxNoise + 1) / 2
	material := Object.To[ShaderMaterial.Instance](v.SeamlessNoiseTexture.AsCanvasItem().Material())
	material.SetShaderParameter("min_value", min)
	material.SetShaderParameter("max_value", max)
}

func (v *NoiseViewer) OnDocumentationButtonPressed() {
	OS.ShellOpen("https://docs.godotengine.org/en/latest/classes/class_fastnoiselite.html")
}

func (v *NoiseViewer) OnRandomSeedButtonPressed() {
	v.ParameterContainer.SeedSpinBox.AsRange().SetValue(Float.X(math.Floor(float64(Float.RandomBetween(-2147483648, 2147483648)))))
}

func (v *NoiseViewer) OnSeedSpinBoxValueChanged(value Float.X) { v.noise.SetSeed(int(value)) }

func (v *NoiseViewer) OnFrequencySpinBoxValueChanged(value Float.X) { v.noise.SetFrequency(value) }

func (v *NoiseViewer) OnFractalOctavesSpinBoxValueChanged(value Float.X) {
	v.noise.SetFractalOctaves(int(value))
}

func (v *NoiseViewer) OnFractalGainSpinBoxValueChanged(value Float.X) { v.noise.SetFractalGain(value) }

func (v *NoiseViewer) OnFractalLacunaritySpinBoxValueChanged(value Float.X) {
	v.noise.SetFractalLacunarity(value)
}

func (v *NoiseViewer) OnMinClipSpinBoxValueChanged(value Float.X) {
	v.minNoise = value
	v.refreshShaderParams()
}

func (v *NoiseViewer) OnMaxClipSpinBoxValueChanged(value Float.X) {
	v.maxNoise = value
	v.refreshShaderParams()
}

func main() {
	classdb.Register[NoiseViewer](NewNoiseViewer)
	startup.Scene()
}
