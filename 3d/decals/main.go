package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Decal"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/PhysicsRayQueryParameters3D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/WorldEnvironment"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/String"
	"graphics.gd/variant/Vector3"
)

const (
	RotSpeed    = 0.003
	ZoomSpeed   = 0.125
	MainButtons = Input.MouseButtonMaskLeft | Input.MouseButtonMaskRight | Input.MouseButtonMaskMiddle
)

var DecalScene = Resource.Load[PackedScene.Is[Node3D.Instance]]("res://decal.tscn")

type DecalsTester struct {
	WorldEnvironment.Extension[DecalsTester]

	Testers      Node3D.Instance
	CameraHolder Node3D.Instance   // Has a position and rotates on Y.
	RotationX    Node3D.Instance   `gd:"CameraHolder/RotationX"`
	Camera       Camera3D.Instance `gd:"CameraHolder/RotationX/Camera3D"`
	TestName     Label.Instance
	Previous     Button.Instance
	Next         Button.Instance

	testerIndex int
	rotX        Angle.Radians // This must be kept in sync with RotationX.
	rotY        Angle.Radians // This must be kept in sync with CameraHolder.
	zoom        Float.X
}

func NewDecalsTester() *DecalsTester {
	return &DecalsTester{
		rotX: Angle.InRadians(Angle.Degrees(-22.5)),
		rotY: Angle.InRadians(Angle.Degrees(90)),
		zoom: 1.5,
	}
}

func (t *DecalsTester) Ready() {
	t.updateCameraRotation()
	t.UpdateGui()
}

func (t *DecalsTester) updateCameraRotation() {
	holder := t.CameraHolder.Transform()
	holder.Basis = Basis.FromEuler(Euler.Radians{Y: t.rotY}, Angle.OrderYXZ)
	t.CameraHolder.SetTransform(holder)
	rotationX := t.RotationX.Transform()
	rotationX.Basis = Basis.FromEuler(Euler.Radians{X: t.rotX}, Angle.OrderYXZ)
	t.RotationX.SetTransform(rotationX)
}

func (t *DecalsTester) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("ui_left") {
		t.OnPreviousPressed()
	}
	if event.IsActionPressed("ui_right") {
		t.OnNextPressed()
	}

	if event.IsActionPressed("place_decal") {
		origin := t.Camera.AsNode3D().GlobalPosition()
		target := t.Camera.ProjectPosition(Viewport.Get(t.AsNode()).GetMousePosition(), 100)

		query := PhysicsRayQueryParameters3D.Create(origin, target, nil)
		result := t.Camera.AsNode3D().GetWorld3d().DirectSpaceState().IntersectRay(query)

		if result.Collider != Object.Nil {
			decal := DecalScene.Instantiate()
			t.AsNode().AddChild(decal.AsNode())
			Object.To[Decal.Instance](decal.AsNode().GetNode("Decal")).SetModulate(Color.RGBA{R: 1, G: 0, B: 0, A: 1})
			decal.SetPosition(Vector3.XYZ{X: result.Position.X, Y: result.Position.Y, Z: result.Position.Z})
			transform := decal.Transform()
			transform.Basis = t.Camera.AsNode3D().GlobalTransform().Basis
			decal.SetTransform(transform)
		}
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

func (t *DecalsTester) Process(delta Float.X) {
	currentTester, _ := Object.As[Node3D.Instance](t.Testers.AsNode().GetChild(t.testerIndex))
	// This code assumes CameraHolder's X and Y coordinates are already correct.
	holder := t.CameraHolder.GlobalTransform()
	holder.Origin.Z = Float.Lerp(holder.Origin.Z, currentTester.GlobalTransform().Origin.Z, 3*delta)
	t.CameraHolder.SetGlobalTransform(holder)
	cameraPosition := t.Camera.AsNode3D().Position()
	cameraPosition.Z = Float.Lerp(cameraPosition.Z, t.zoom, 10*delta)
	t.Camera.AsNode3D().SetPosition(cameraPosition)
}

func (t *DecalsTester) OnPreviousPressed() {
	t.testerIndex = max(0, t.testerIndex-1)
	t.UpdateGui()
}

func (t *DecalsTester) OnNextPressed() {
	t.testerIndex = min(t.testerIndex+1, t.Testers.AsNode().GetChildCount()-1)
	t.UpdateGui()
}

func (t *DecalsTester) UpdateGui() {
	t.TestName.SetText(String.Capitalize(t.Testers.AsNode().GetChild(t.testerIndex).Name()))
	t.Previous.AsBaseButton().SetDisabled(t.testerIndex == 0)
	t.Next.AsBaseButton().SetDisabled(t.testerIndex == t.Testers.AsNode().GetChildCount()-1)
}

// OnDecalFilterModeItemSelected is connected to the "item_selected" signal of the OptionButton.
func (t *DecalsTester) OnDecalFilterModeItemSelected(index int) {
	// Indices in the OptionButton match RenderingServer decal filtering constants.
	RenderingServer.DecalsSetFilter(RenderingServer.DecalFilter(index))
}

func main() {
	classdb.Register[DecalsTester](NewDecalsTester)
	startup.Scene()
}
