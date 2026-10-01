package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/String"
)

const (
	RotSpeed    = 0.003
	ZoomSpeed   = 0.125
	MainButtons = Input.MouseButtonMaskLeft | Input.MouseButtonMaskRight | Input.MouseButtonMaskMiddle
)

type MaterialTester struct {
	WorldEnvironment.Extension[MaterialTester]

	DirectionalLight3D DirectionalLight3D.Instance
	Testers            Node3D.Instance
	CameraHolder       Node3D.Instance   // Has a position and rotates on Y.
	RotationX          Node3D.Instance   `gd:"CameraHolder/RotationX"`
	Camera             Camera3D.Instance `gd:"CameraHolder/RotationX/Camera3D"`
	TestName           Label.Instance
	Previous           Button.Instance
	Next               Button.Instance

	testerIndex int
	rotX        Angle.Radians // This must be kept in sync with RotationX.
	rotY        Angle.Radians // This must be kept in sync with CameraHolder.
	zoom        Float.X
}

func NewMaterialTester() *MaterialTester {
	return &MaterialTester{
		rotX: Angle.InRadians(Angle.Degrees(-22.5)),
		rotY: Angle.InRadians(Angle.Degrees(90)),
		zoom: 2.5,
	}
}

func (t *MaterialTester) Ready() {
	if RenderingServer.GetCurrentRenderingMethod() == "gl_compatibility" {
		// Use PCF13 shadow filtering to improve quality (Medium maps to PCF5 instead).
		RenderingServer.DirectionalSoftShadowFilterSetQuality(RenderingServer.ShadowQualitySoftHigh)

		// Darken the light's energy to compensate for sRGB blending (without affecting sky rendering).
		t.DirectionalLight3D.SetSkyMode(DirectionalLight3D.SkyModeSkyOnly)
		newLight, _ := Object.As[DirectionalLight3D.Instance](t.DirectionalLight3D.AsNode().Duplicate())
		newLight.AsLight3D().SetLightEnergy(0.25)
		newLight.SetSkyMode(DirectionalLight3D.SkyModeLightOnly)
		t.AsNode().AddChild(newLight.AsNode())
	}

	t.updateCameraRotation()
	t.UpdateGui()
}

func (t *MaterialTester) updateCameraRotation() {
	holder := t.CameraHolder.Transform()
	holder.Basis = Basis.FromEuler(Euler.Radians{Y: t.rotY}, Angle.OrderYXZ)
	t.CameraHolder.SetTransform(holder)
	rotationX := t.RotationX.Transform()
	rotationX.Basis = Basis.FromEuler(Euler.Radians{X: t.rotX}, Angle.OrderYXZ)
	t.RotationX.SetTransform(rotationX)
}

func (t *MaterialTester) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("ui_left") {
		t.OnPreviousPressed()
	}
	if event.IsActionPressed("ui_right") {
		t.OnNextPressed()
	}

	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			t.zoom -= ZoomSpeed
		}
		if button.ButtonIndex() == Input.MouseButtonWheelDown {
			t.zoom += ZoomSpeed
		}
		t.zoom = Float.Clamp(t.zoom, 1.5, 4)
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok && motion.AsInputEventMouse().ButtonMask()&MainButtons != 0 {
		// Use `screen_relative` to make mouse sensitivity independent of viewport resolution.
		relative := motion.ScreenRelative()
		t.rotY -= Angle.Radians(relative.X * RotSpeed)
		t.rotX -= Angle.Radians(relative.Y * RotSpeed)
		t.rotX = Float.Clamp(t.rotX, Angle.InRadians(Angle.Degrees(-90)), 0)
		t.updateCameraRotation()
	}
}

func (t *MaterialTester) Process(delta Float.X) {
	currentTester, _ := Object.As[Node3D.Instance](t.Testers.AsNode().GetChild(t.testerIndex))
	// This code assumes CameraHolder's X and Y coordinates are already correct.
	holder := t.CameraHolder.GlobalTransform()
	holder.Origin.Z = Float.Lerp(holder.Origin.Z, currentTester.GlobalTransform().Origin.Z, 3*delta)
	t.CameraHolder.SetGlobalTransform(holder)
	cameraPosition := t.Camera.AsNode3D().Position()
	cameraPosition.Z = Float.Lerp(cameraPosition.Z, t.zoom, 10*delta)
	t.Camera.AsNode3D().SetPosition(cameraPosition)
}

func (t *MaterialTester) OnPreviousPressed() {
	t.testerIndex = max(0, t.testerIndex-1)
	t.UpdateGui()
}

func (t *MaterialTester) OnNextPressed() {
	t.testerIndex = min(t.testerIndex+1, t.Testers.AsNode().GetChildCount()-1)
	t.UpdateGui()
}

func (t *MaterialTester) UpdateGui() {
	t.TestName.SetText(String.Capitalize(t.Testers.AsNode().GetChild(t.testerIndex).Name()))
	t.Previous.AsBaseButton().SetDisabled(t.testerIndex == 0)
	t.Next.AsBaseButton().SetDisabled(t.testerIndex == t.Testers.AsNode().GetChildCount()-1)
}

func main() {
	classdb.Register[LoadingScreen]()
	classdb.Register[MaterialTester](NewMaterialTester)
	classdb.Register[GridMaterial]()
	startup.Scene()
}
