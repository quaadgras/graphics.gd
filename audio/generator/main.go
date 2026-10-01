package main

import (
	"fmt"
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioStreamGenerator"
	"graphics.gd/classdb/AudioStreamGeneratorPlayback"
	"graphics.gd/classdb/AudioStreamPlayer"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

// GeneratorDemo fills an AudioStreamGenerator with a sine wave every frame.
type GeneratorDemo struct {
	Node.Extension[GeneratorDemo] `gd:"GeneratorDemo"`

	Player         AudioStreamPlayer.Instance
	FrequencyLabel Label.Instance `gd:"%FrequencyLabel"`
	VolumeLabel    Label.Instance `gd:"%VolumeLabel"`

	sampleHz Float.X
	pulseHz  Float.X
	phase    Float.X

	// Actual playback stream, assigned in Ready.
	playback AudioStreamGeneratorPlayback.Instance
}

func NewGeneratorDemo() *GeneratorDemo {
	return &GeneratorDemo{sampleHz: 22050, pulseHz: 440}
}

func (g *GeneratorDemo) fillBuffer() {
	increment := g.pulseHz / g.sampleHz
	toFill := g.playback.GetFramesAvailable()
	frames := make([]Vector2.XY, 0, toFill)
	for range toFill {
		frames = append(frames, Vector2.MulX(Vector2.One, Float.X(math.Sin(float64(g.phase*2*math.Pi))))) // Audio frames are stereo.
		g.phase = Float.Mod(g.phase+increment, 1.0)
	}
	g.playback.PushBuffer(frames)
}

func (g *GeneratorDemo) Process(_ Float.X) {
	g.fillBuffer()
}

func (g *GeneratorDemo) Ready() {
	// Setting mix rate is only possible before Play().
	if stream, ok := Object.As[AudioStreamGenerator.Instance](g.Player.Stream()); ok {
		stream.SetMixRate(g.sampleHz)
	}
	g.Player.Play()
	g.playback, _ = Object.As[AudioStreamGeneratorPlayback.Instance](g.Player.GetStreamPlayback())
	// fillBuffer must be called *after* setting playback.
	g.fillBuffer()
}

func (g *GeneratorDemo) OnFrequencyHSliderValueChanged(value Float.X) {
	g.FrequencyLabel.SetText(fmt.Sprintf("%d Hz", int(value)))
	g.pulseHz = value
}

func (g *GeneratorDemo) OnVolumeHSliderValueChanged(value Float.X) {
	// Use LinearToDecibels to get a volume slider that matches perceptual human hearing.
	g.VolumeLabel.SetText(fmt.Sprintf("%.2f dB", Float.LinearToDecibels(value)))
	g.Player.SetVolumeDb(Float.LinearToDecibels(value))
}

func main() {
	classdb.Register[GeneratorDemo](NewGeneratorDemo)
	startup.Scene()
}
