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
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Viewport"
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

type Tester struct {
	WorldEnvironment.Extension[Tester]

	DirectionalLight3D *DayNightCycle
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

func NewTester() *Tester {
	return &Tester{
		rotX: Angle.InRadians(Angle.Degrees(-22.5)),
		rotY: Angle.InRadians(Angle.Degrees(90)),
		zoom: 2.5,
	}
}

func (t *Tester) Ready() {
	t.updateCameraRotation()
	t.UpdateGui()
}

func (t *Tester) updateCameraRotation() {
	var holder = t.CameraHolder.Transform()
	holder.Basis = Basis.FromEuler(Euler.Radians{Y: t.rotY}, Angle.OrderYXZ)
	t.CameraHolder.SetTransform(holder)
	var rotationX = t.RotationX.Transform()
	rotationX.Basis = Basis.FromEuler(Euler.Radians{X: t.rotX}, Angle.OrderYXZ)
	t.RotationX.SetTransform(rotationX)
}

func (t *Tester) UnhandledInput(event InputEvent.Instance) {
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
		var relativeMotion = motion.ScreenRelative()
		t.rotY -= Angle.Radians(relativeMotion.X * RotSpeed)
		t.rotX -= Angle.Radians(relativeMotion.Y * RotSpeed)
		t.rotX = Float.Clamp(t.rotX, Angle.InRadians(Angle.Degrees(-90)), 0)
		t.updateCameraRotation()
	}
}

func (t *Tester) Process(delta Float.X) {
	currentTester, _ := Object.As[Node3D.Instance](t.Testers.AsNode().GetChild(t.testerIndex))
	// This code assumes CameraHolder's X and Y coordinates are already correct.
	var currentPosition = t.CameraHolder.GlobalTransform().Origin.Z
	var targetPosition = currentTester.GlobalTransform().Origin.Z
	var holder = t.CameraHolder.GlobalTransform()
	holder.Origin.Z = Float.Lerp(currentPosition, targetPosition, 3*delta)
	t.CameraHolder.SetGlobalTransform(holder)
	var cameraPosition = t.Camera.AsNode3D().Position()
	cameraPosition.Z = Float.Lerp(cameraPosition.Z, t.zoom, 10*delta)
	t.Camera.AsNode3D().SetPosition(cameraPosition)
}

func (t *Tester) OnPreviousPressed() {
	t.testerIndex = max(0, t.testerIndex-1)
	t.UpdateGui()
}

func (t *Tester) OnNextPressed() {
	t.testerIndex = min(t.testerIndex+1, t.Testers.AsNode().GetChildCount()-1)
	t.UpdateGui()
}

func (t *Tester) UpdateGui() {
	t.TestName.SetText(String.Capitalize(t.Testers.AsNode().GetChild(t.testerIndex).Name()))
	t.Previous.AsBaseButton().SetDisabled(t.testerIndex == 0)
	t.Next.AsBaseButton().SetDisabled(t.testerIndex == t.Testers.AsNode().GetChildCount()-1)
}

func (t *Tester) OnEnableSunToggled(buttonPressed bool) {
	t.DirectionalLight3D.AsNode3D().SetVisible(buttonPressed)
}

func (t *Tester) OnAnimateLightsToggled(buttonPressed bool) {
	for _, animatableNode := range SceneTree.Get(t.AsNode()).GetNodesInGroup("animatable") {
		animatableNode.SetProcess(buttonPressed)
	}
}

func (t *Tester) OnShadowResolutionItemSelected(index int) {
	var size = 4096
	switch index {
	case 0:
		size = 1024
	case 1:
		size = 2048
	case 2:
		size = 4096
	case 3:
		size = 8192
	case 4:
		size = 16384
	}
	RenderingServer.DirectionalShadowAtlasSetSize(size, true)
	Viewport.Get(t.AsNode()).SetPositionalShadowAtlasSize(size)
}

func (t *Tester) OnShadowFilterQualityItemSelected(index int) {
	// Values are numbered in the OptionButton to match the RenderingServer.ShadowQuality enum.
	RenderingServer.DirectionalSoftShadowFilterSetQuality(RenderingServer.ShadowQuality(index))
	RenderingServer.PositionalSoftShadowFilterSetQuality(RenderingServer.ShadowQuality(index))
}

func (t *Tester) OnProjectorFilterModeItemSelected(index int) {
	// Values are numbered in the OptionButton to match the RenderingServer.LightProjectorFilter enum.
	RenderingServer.LightProjectorsSetFilter(RenderingServer.LightProjectorFilter(index))
}

func main() {
	classdb.Register[Tester](NewTester)
	classdb.Register[DayNightCycle]()
	classdb.Register[SpinOmniLight]()
	classdb.Register[SpinSpotLight]()
	startup.Scene()
}
