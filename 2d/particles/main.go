package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/GPUParticles2D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type PauseLabel struct {
	Label.Extension[PauseLabel]

	UnsupportedLabel Label.Instance            `gd:"../UnsupportedLabel"`
	World            WorldEnvironment.Instance `gd:"../.."`

	isCompatibility bool
}

func (l *PauseLabel) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		l.isCompatibility = true
		l.AsLabel().SetText("Space: Pause/Resume\nG: Toggle glow\n\n\n")
		l.UnsupportedLabel.AsCanvasItem().SetVisible(true)
		// Increase glow intensity to compensate for lower dynamic range.
		l.World.Environment().SetGlowIntensity(4.0)
	}
}

func (l *PauseLabel) trailableParticles() []GPUParticles2D.Instance {
	var result []GPUParticles2D.Instance
	for _, node := range SceneTree.Get(l.AsNode()).GetNodesInGroup("trailable_particles") {
		if particles, ok := Object.As[GPUParticles2D.Instance](node); ok {
			result = append(result, particles)
		}
	}
	return result
}

func (l *PauseLabel) Input(event InputEvent.Instance) {
	tree := SceneTree.Get(l.AsNode())
	if event.IsActionPressed("toggle_pause") {
		tree.SetPaused(!tree.Paused())
	}
	// Particles disappear if trail type is changed while paused.
	// Prevent changing particle type while paused to avoid confusion.
	if !l.isCompatibility && event.IsActionPressed("toggle_trails") {
		for _, particles := range l.trailableParticles() {
			particles.SetTrailEnabled(!particles.TrailEnabled())
		}
	}
	if !l.isCompatibility && event.IsActionPressed("increase_trail_length") {
		for _, particles := range l.trailableParticles() {
			particles.SetTrailLifetime(Float.Clamp(particles.TrailLifetime()+0.05, 0.1, 1.0))
		}
	}
	if !l.isCompatibility && event.IsActionPressed("decrease_trail_length") {
		for _, particles := range l.trailableParticles() {
			particles.SetTrailLifetime(Float.Clamp(particles.TrailLifetime()-0.05, 0.1, 1.0))
		}
	}
	if event.IsActionPressed("toggle_glow") {
		env := l.World.Environment()
		env.SetGlowEnabled(!env.GlowEnabled())
	}
}

func main() {
	classdb.Register[PauseLabel]()
	startup.Scene()
}
