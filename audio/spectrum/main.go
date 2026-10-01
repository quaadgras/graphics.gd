package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioEffectSpectrumAnalyzerInstance"
	"graphics.gd/classdb/AudioServer"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

const (
	VuCount        = 16
	FreqMax        = 11050.0
	Width          = 800
	Height         = 250
	HeightScale    = 8.0
	MinDB          = 60
	AnimationSpeed = 0.1
)

// ShowSpectrum draws a VU-style bar graph of the master bus spectrum analyzer.
type ShowSpectrum struct {
	Node2D.Extension[ShowSpectrum] `gd:"ShowSpectrum"`

	spectrum  AudioEffectSpectrumAnalyzerInstance.Instance
	minValues [VuCount]Float.X
	maxValues [VuCount]Float.X
}

func (s *ShowSpectrum) Draw() {
	canvas := s.AsCanvasItem()
	const w = Width / VuCount
	for i := range VuCount {
		height := Float.Lerp(s.minValues[i], s.maxValues[i], AnimationSpeed)
		hue := Float.X(VuCount*0.6+Float.X(i)*0.5) / VuCount
		fill := Color.HSV(hue, 0.5, 0.6)
		line := Color.HSV(hue, 0.5, 1.0)
		faded := Color.RGBA{R: 1, G: 1, B: 1, A: 0.125}

		canvas.DrawRect(Rect2.New(Float.X(w*i), Height-height, w-2, height), fill)
		canvas.MoreArgs().DrawLine(
			Vector2.New(Float.X(w*i), Height-height),
			Vector2.New(Float.X(w*i+w-2), Height-height),
			line, 2.0, true,
		)

		// Draw a reflection of the bars with lower opacity.
		canvas.DrawRect(Rect2.New(Float.X(w*i), Height, w-2, height), Color.Mul(fill, faded))
		canvas.MoreArgs().DrawLine(
			Vector2.New(Float.X(w*i), Height+height),
			Vector2.New(Float.X(w*i+w-2), Height+height),
			Color.Mul(line, faded), 2.0, true,
		)
	}
}

func (s *ShowSpectrum) Process(_ Float.X) {
	var data [VuCount]Float.X
	var prevHz Float.X
	for i := 1; i <= VuCount; i++ {
		hz := Float.X(i) * FreqMax / VuCount
		magnitude := Vector2.Length(s.spectrum.GetMagnitudeForFrequencyRange(prevHz, hz))
		energy := Float.Clamp((MinDB+Float.LinearToDecibels(magnitude))/MinDB, 0, 1)
		data[i-1] = energy * Height * HeightScale
		prevHz = hz
	}
	for i := range VuCount {
		if data[i] > s.maxValues[i] {
			s.maxValues[i] = data[i]
		} else {
			s.maxValues[i] = Float.Lerp(s.maxValues[i], data[i], AnimationSpeed)
		}
		if data[i] <= 0 {
			s.minValues[i] = Float.Lerp(s.minValues[i], 0, AnimationSpeed)
		}
	}
	// Sound plays back continuously, so the graph needs to be updated every frame.
	s.AsCanvasItem().QueueRedraw()
}

func (s *ShowSpectrum) Ready() {
	s.spectrum, _ = Object.As[AudioEffectSpectrumAnalyzerInstance.Instance](AudioServer.GetBusEffectInstance(0, 0, 0))
}

func main() {
	classdb.Register[ShowSpectrum]()
	startup.Scene()
}
