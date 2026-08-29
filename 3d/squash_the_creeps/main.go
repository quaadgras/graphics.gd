package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/ColorRect"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/PathFollow3D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Timer"
	"graphics.gd/startup"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type Main struct {
	Node.Extension[Main]

	MobScene PackedScene.Instance

	DirectionalLight3D DirectionalLight3D.Instance
	Player             *Player
	SpawnLocation      PathFollow3D.Instance `gd:"SpawnPath/SpawnLocation"`
	MobTimer           Timer.Instance
	ScoreLabel         *ScoreLabel        `gd:"UserInterface/ScoreLabel"`
	Retry              ColorRect.Instance `gd:"UserInterface/Retry"`
}

func (m *Main) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		// Use PCF13 shadow filtering to improve quality (Medium maps to PCF5 instead).
		RenderingServer.DirectionalSoftShadowFilterSetQuality(RenderingServer.ShadowQualitySoftHigh)

		// Darken the light's energy to compensate for sRGB blending (without affecting sky rendering).
		m.DirectionalLight3D.SetSkyMode(DirectionalLight3D.SkyModeSkyOnly)
		newLight, _ := Object.As[DirectionalLight3D.Instance](m.DirectionalLight3D.AsNode().Duplicate())
		newLight.AsLight3D().SetLightEnergy(0.35)
		newLight.SetSkyMode(DirectionalLight3D.SkyModeLightOnly)
		m.AsNode().AddChild(newLight.AsNode())
	}
	m.Retry.AsCanvasItem().Hide()
}

func (m *Main) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("ui_accept") && m.Retry.AsCanvasItem().Visible() {
		SceneTree.Get(m.AsNode()).ReloadCurrentScene()
	}
}

func (m *Main) OnMobTimerTimeout() {
	// Create a new instance of the Mob scene.
	mob, ok := Object.As[*Mob](m.MobScene.Instantiate())
	if !ok {
		return
	}
	// Choose a random location on the SpawnPath.
	m.SpawnLocation.SetProgressRatio(Float.Random())

	// Communicate the spawn location and the player's location to the mob.
	mob.Initialize(m.SpawnLocation.AsNode3D().Position(), m.Player.AsNode3D().Position())

	// Spawn the mob by adding it to the Main scene.
	m.AsNode().AddChild(mob.AsNode())
	// We connect the mob to the score label to update the score upon squashing a mob.
	mob.Squashed.Attach(Callable.New(m.ScoreLabel.OnMobSquashed))
}

func (m *Main) OnPlayerHit() {
	m.MobTimer.Stop()
	m.Retry.AsCanvasItem().Show()
}

func main() {
	classdb.Register[Main]()
	classdb.Register[Mob](NewMob)
	classdb.Register[Player](NewPlayer)
	classdb.Register[ScoreLabel]()
	startup.Scene()
}
