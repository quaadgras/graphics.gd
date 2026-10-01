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

type LabelsAndTexts struct {
	WorldEnvironment.Extension[LabelsAndTexts]

	Testers      Node3D.Instance
	CameraHolder Node3D.Instance   // Has a position and rotates on Y.
	RotationX    Node3D.Instance   `gd:"CameraHolder/RotationX"`
	Camera       Camera3D.Instance `gd:"CameraHolder/RotationX/Camera3D"`
	TestName     Label.Instance
	Previous     Button.Instance
	Next         Button.Instance
	HealthBar    *Label3DLayout `gd:"Testers/Label3DHealthBar"`

	testerIndex    int
	rotX           Angle.Radians // This must be kept in sync with RotationX.
	rotY           Angle.Radians // This must be kept in sync with CameraHolder.
	cameraDistance Float.X
}

func NewLabelsAndTexts() *LabelsAndTexts {
	return &LabelsAndTexts{
		rotX:           -Angle.Tau / 16,
		rotY:           Angle.Tau / 8,
		cameraDistance: 2.0,
	}
}

func (l *LabelsAndTexts) Ready() {
	l.updateCameraRotation()
	l.UpdateGui()
}

func (l *LabelsAndTexts) updateCameraRotation() {
	var holder = l.CameraHolder.Transform()
	holder.Basis = Basis.FromEuler(Euler.Radians{Y: l.rotY}, Angle.OrderYXZ)
	l.CameraHolder.SetTransform(holder)
	var rotationX = l.RotationX.Transform()
	rotationX.Basis = Basis.FromEuler(Euler.Radians{X: l.rotX}, Angle.OrderYXZ)
	l.RotationX.SetTransform(rotationX)
}

func (l *LabelsAndTexts) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("ui_left") {
		l.OnPreviousPressed()
	}
	if event.IsActionPressed("ui_right") {
		l.OnNextPressed()
	}

	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			l.cameraDistance -= ZoomSpeed
		}
		if button.ButtonIndex() == Input.MouseButtonWheelDown {
			l.cameraDistance += ZoomSpeed
		}
		l.cameraDistance = Float.Clamp(l.cameraDistance, 1.5, 6.0)
	}

	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok && motion.AsInputEventMouse().ButtonMask()&MainButtons != 0 {
		// Use `screen_relative` to make mouse sensitivity independent of viewport resolution.
		var relativeMotion = motion.ScreenRelative()
		l.rotY -= Angle.Radians(relativeMotion.X * RotSpeed)
		l.rotX -= Angle.Radians(relativeMotion.Y * RotSpeed)
		l.rotX = Float.Clamp(l.rotX, -1.57, 0.0)
		l.updateCameraRotation()
	}
}

func (l *LabelsAndTexts) Process(delta Float.X) {
	currentTester, _ := Object.As[Node3D.Instance](l.Testers.AsNode().GetChild(l.testerIndex))
	// This code assumes CameraHolder's X and Y coordinates are already correct.
	var currentPosition = l.CameraHolder.GlobalTransform().Origin.Z
	var targetPosition = currentTester.GlobalTransform().Origin.Z
	var holder = l.CameraHolder.GlobalTransform()
	holder.Origin.Z = Float.Lerp(currentPosition, targetPosition, 3*delta)
	l.CameraHolder.SetGlobalTransform(holder)
	var cameraPosition = l.Camera.AsNode3D().Position()
	cameraPosition.Z = Float.Lerp(cameraPosition.Z, l.cameraDistance, 10*delta)
	l.Camera.AsNode3D().SetPosition(cameraPosition)
}

func (l *LabelsAndTexts) OnPreviousPressed() {
	l.testerIndex = max(0, l.testerIndex-1)
	l.UpdateGui()
}

func (l *LabelsAndTexts) OnNextPressed() {
	l.testerIndex = min(l.testerIndex+1, l.Testers.AsNode().GetChildCount()-1)
	l.UpdateGui()
}

func (l *LabelsAndTexts) UpdateGui() {
	var name = l.Testers.AsNode().GetChild(l.testerIndex).Name()
	l.TestName.SetText(String.Capitalize(name))
	l.Previous.AsBaseButton().SetDisabled(l.testerIndex == 0)
	l.Next.AsBaseButton().SetDisabled(l.testerIndex == l.Testers.AsNode().GetChildCount()-1)

	// Only display player name field if relevant.
	l.HealthBar.Name2.AsCanvasItem().SetVisible(name == "Label3DHealthBar")
	l.HealthBar.LineEdit.AsCanvasItem().SetVisible(name == "Label3DHealthBar")
}

func (l *LabelsAndTexts) OnLineEditTextSubmitted(_ string) {
	l.HealthBar.LineEdit.AsControl().ReleaseFocus()
}

func main() {
	classdb.Register[LabelsAndTexts](NewLabelsAndTexts)
	classdb.Register[Label3DLayout]()
	startup.Scene()
}
