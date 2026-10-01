package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
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

type SpritesTester struct {
	WorldEnvironment.Extension[SpritesTester]

	Testers      Node3D.Instance
	CameraHolder Node3D.Instance   // Has a position and rotates on Y.
	RotationX    Node3D.Instance   `gd:"CameraHolder/RotationX"`
	Camera       Camera3D.Instance `gd:"CameraHolder/RotationX/Camera3D"`
	TestName     Label.Instance
	Previous     Button.Instance
	Next         Button.Instance

	testerIndex    int
	rotX           Angle.Radians // This must be kept in sync with RotationX.
	rotY           Angle.Radians // This must be kept in sync with CameraHolder.
	cameraDistance Float.X
}

func NewSpritesTester() *SpritesTester {
	return &SpritesTester{
		rotX:           -Angle.Tau / 16,
		rotY:           Angle.Tau / 6,
		cameraDistance: 3.4,
	}
}

func (t *SpritesTester) Ready() {
	t.updateCameraRotation()
	t.UpdateGui()
}

func (t *SpritesTester) updateCameraRotation() {
	holder := t.CameraHolder.Transform()
	holder.Basis = Basis.FromEuler(Euler.Radians{Y: t.rotY}, Angle.OrderYXZ)
	t.CameraHolder.SetTransform(holder)
	rotationX := t.RotationX.Transform()
	rotationX.Basis = Basis.FromEuler(Euler.Radians{X: t.rotX}, Angle.OrderYXZ)
	t.RotationX.SetTransform(rotationX)
}

func (t *SpritesTester) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("ui_left") {
		t.OnPreviousPressed()
	}
	if event.IsActionPressed("ui_right") {
		t.OnNextPressed()
	}

	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			t.cameraDistance -= ZoomSpeed
		}
		if button.ButtonIndex() == Input.MouseButtonWheelDown {
			t.cameraDistance += ZoomSpeed
		}
		t.cameraDistance = Float.Clamp(t.cameraDistance, 1.5, 6)
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok && motion.AsInputEventMouse().ButtonMask()&MainButtons != 0 {
		// Use `screen_relative` to make mouse sensitivity independent of viewport resolution.
		relative := motion.ScreenRelative()
		t.rotY -= Angle.Radians(relative.X * RotSpeed)
		t.rotX -= Angle.Radians(relative.Y * RotSpeed)
		t.rotX = Float.Clamp(t.rotX, -1.57, 0)
		t.updateCameraRotation()
	}
}

func (t *SpritesTester) Process(delta Float.X) {
	currentTester, _ := Object.As[Node3D.Instance](t.Testers.AsNode().GetChild(t.testerIndex))
	// This code assumes CameraHolder's X and Y coordinates are already correct.
	holder := t.CameraHolder.GlobalTransform()
	holder.Origin.Z = Float.Lerp(holder.Origin.Z, currentTester.GlobalTransform().Origin.Z, 3*delta)
	t.CameraHolder.SetGlobalTransform(holder)
	cameraPosition := t.Camera.AsNode3D().Position()
	cameraPosition.Z = Float.Lerp(cameraPosition.Z, t.cameraDistance, 10*delta)
	t.Camera.AsNode3D().SetPosition(cameraPosition)
}

func (t *SpritesTester) OnPreviousPressed() {
	t.testerIndex = max(0, t.testerIndex-1)
	t.UpdateGui()
}

func (t *SpritesTester) OnNextPressed() {
	t.testerIndex = min(t.testerIndex+1, t.Testers.AsNode().GetChildCount()-1)
	t.UpdateGui()
}

func (t *SpritesTester) UpdateGui() {
	t.TestName.SetText(String.Capitalize(t.Testers.AsNode().GetChild(t.testerIndex).Name()))
	t.Previous.AsBaseButton().SetDisabled(t.testerIndex == 0)
	t.Next.AsBaseButton().SetDisabled(t.testerIndex == t.Testers.AsNode().GetChildCount()-1)
}

func main() {
	classdb.Register[SpritesTester](NewSpritesTester)
	classdb.Register[SpritePlayer](NewSpritePlayer)
	classdb.Register[SpriteRotate](NewSpriteRotate)
	startup.Scene()
}
