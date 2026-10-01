package main

import (
	"math"

	"graphics.gd/classdb"
	"graphics.gd/classdb/BaseMaterial3D"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/SubViewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

// CameraIdleScale is the camera idle scale effect intensity.
const CameraIdleScale = 0.005

// TwoDInThreeD displays a 2D game (Pong) rendered into a SubViewport on a 3D quad.
type TwoDInThreeD struct {
	Node3D.Extension[TwoDInThreeD] `gd:"TwoDInThreeD"`

	Camera3D     Camera3D.Instance
	SubViewport  SubViewport.Instance
	ViewportQuad MeshInstance3D.Instance

	counter            Float.X
	cameraBaseRotation Euler.Radians
}

func (t *TwoDInThreeD) Ready() {
	t.cameraBaseRotation = t.Camera3D.AsNode3D().Rotation()

	// Clear the viewport.
	t.SubViewport.SetRenderTargetClearMode(SubViewport.ClearModeOnce)

	// Retrieve the texture and set it to the viewport quad.
	material, _ := Object.As[BaseMaterial3D.Instance](t.ViewportQuad.AsGeometryInstance3D().MaterialOverride())
	material.SetAlbedoTexture(t.SubViewport.AsViewport().GetTexture().AsTexture2D())
}

func (t *TwoDInThreeD) Process(delta Float.X) {
	// Animate the camera with an "idle" animation.
	t.counter += delta
	base := t.cameraBaseRotation.Y
	t.Camera3D.AsNode3D().SetRotation(Euler.Radians{
		X: base + Angle.Radians(math.Cos(float64(t.counter))*CameraIdleScale),
		Y: base + Angle.Radians(math.Sin(float64(t.counter))*CameraIdleScale),
		Z: base + Angle.Radians(math.Sin(float64(t.counter))*CameraIdleScale),
	})
}

func main() {
	classdb.Register[TwoDInThreeD]()
	classdb.Register[Pong](NewPong)
	startup.Scene()
}
