package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioServer"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/PhysicsRayQueryParameters3D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const (
	MouseSensitivity        = 0.01
	InitialVelocityStrength = 0.5

	// Margin to add to the automatically computed shadow maximum distance.
	// This value was empirically chosen to cover the whole scene when zoomed
	// all the way in.
	DirectionalShadowMaxDistanceMargin = 9.0
)

var RagdollScene = Resource.Load[PackedScene.Instance]("res://characters/mannequiny_ragdoll.tscn")

type RagdollPhysics struct {
	Node3D.Extension[RagdollPhysics]

	CameraPivot      Node3D.Instance
	Camera           Camera3D.Instance           `gd:"CameraPivot/Camera3D"`
	DirectionalLight DirectionalLight3D.Instance `gd:"DirectionalLight3D"`
}

func (r *RagdollPhysics) UnhandledInput(event InputEvent.Instance) {
	if event.IsActionPressed("reset_simulation") {
		SceneTree.Get(r.AsNode()).ReloadCurrentScene()
	}

	if event.IsActionPressed("place_ragdoll") {
		var origin = r.Camera.AsNode3D().GlobalPosition()
		var target = r.Camera.ProjectPosition(Viewport.Get(r.AsNode()).GetMousePosition(), 100)

		var query = PhysicsRayQueryParameters3D.Create(origin, target, nil)
		var result = r.Camera.AsNode3D().GetWorld3d().DirectSpaceState().IntersectRay(query)

		if result.Collider != Object.Nil {
			ragdoll, ok := Object.As[*MannequinyRagdoll](RagdollScene.Instantiate())
			if !ok {
				return
			}
			var position = Vector3.XYZ{result.Position.X, result.Position.Y, result.Position.Z}
			ragdoll.AsNode3D().SetPosition(Vector3.Add(position, Vector3.XYZ{0, 0.5, 0}))
			// Make newly spawned ragdolls face the camera.
			var rotation = ragdoll.AsNode3D().Rotation()
			rotation.Y = r.CameraPivot.Rotation().Y
			ragdoll.AsNode3D().SetRotation(rotation)
			// Give some initial velocity in a random horizontal direction.
			ragdoll.InitialVelocity = Vector3.MulX(Vector3.Rotated(Vector3.Forward, Vector3.Up, Angle.Radians(Float.RandomBetween(0, Float.X(Angle.Tau)))), InitialVelocityStrength)
			r.AsNode().AddChild(ragdoll.AsNode())
		}
	}

	if event.IsActionPressed("slow_motion") {
		Engine.SetTimeScale(0.25)
		// Don't set pitch scale too low as it sounds strange.
		// `0.5` is the square root of `0.25` and gives a good result here.
		AudioServer.SetPlaybackSpeedScale(0.5)
	}

	if event.IsActionReleased("slow_motion") {
		Engine.SetTimeScale(1.0)
		AudioServer.SetPlaybackSpeedScale(1.0)
	}

	// Pan the camera with right mouse button.
	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		if motion.AsInputEventMouse().ButtonMask()&Input.MouseButtonMaskRight != 0 {
			var rotation = r.CameraPivot.GlobalRotation()
			rotation.X = Float.Clamp(rotation.X-Angle.Radians(motion.ScreenRelative().Y*MouseSensitivity), -Angle.Tau*0.249, Angle.Tau*0.021)
			rotation.Y -= Angle.Radians(motion.ScreenRelative().X * MouseSensitivity)
			r.CameraPivot.SetGlobalRotation(rotation)
		}
	}

	// Zoom with mouse wheel.
	// This also adjusts shadow maximum distance to always cover the scene regardless of zoom level.
	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok {
		if button.ButtonIndex() == Input.MouseButtonWheelUp {
			r.Camera.AsNode3D().TranslateObjectLocal(Vector3.MulX(Vector3.Forward, 0.5))
			r.setShadowMaxDistance()
		} else if button.ButtonIndex() == Input.MouseButtonWheelDown {
			r.Camera.AsNode3D().TranslateObjectLocal(Vector3.MulX(Vector3.Back, 0.5))
			r.setShadowMaxDistance()
		}
	}
}

// setShadowMaxDistance adjusts the shadow maximum distance to cover the scene at the current zoom level.
func (r *RagdollPhysics) setShadowMaxDistance() {
	Object.Set(r.DirectionalLight, "directional_shadow_max_distance", Vector3.Length(r.Camera.AsNode3D().Position())+DirectionalShadowMaxDistanceMargin)
}

func main() {
	classdb.Register[RagdollPhysics]()
	classdb.Register[MannequinyRagdoll]()
	startup.Scene()
}
