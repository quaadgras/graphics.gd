package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Area3D"
	"graphics.gd/classdb/BaseMaterial3D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouse"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/InputEventScreenDrag"
	"graphics.gd/classdb/InputEventScreenTouch"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PlaneMesh"
	"graphics.gd/classdb/SubViewport"
	"graphics.gd/classdb/Time"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
)

// GUIPanel3D displays a GUI rendered into a SubViewport on a 3D quad,
// forwarding mouse and keyboard input to the GUI.
type GUIPanel3D struct {
	Node3D.Extension[GUIPanel3D] `gd:"GUIPanel3D"`

	NodeViewport SubViewport.Instance    `gd:"SubViewport"`
	NodeQuad     MeshInstance3D.Instance `gd:"Quad"`
	NodeArea     Area3D.Instance         `gd:"Quad/Area3D"`

	isMouseInside bool // Used for checking if the mouse is inside the Area3D.

	// The last processed input touch/mouse event. Used to calculate relative movement.
	lastEventPos2D    Vector2.XY
	hasLastEventPos2D bool

	lastEventTime Float.X // The time of the last event in seconds since engine start.
}

func (g *GUIPanel3D) Ready() {
	g.NodeArea.AsCollisionObject3D().OnMouseEntered(g.mouseEnteredArea)
	g.NodeArea.AsCollisionObject3D().OnMouseExited(g.mouseExitedArea)
	g.NodeArea.AsCollisionObject3D().OnInputEvent(g.mouseInputEvent)

	// If the material is NOT set to use billboard settings, then avoid running billboard specific code.
	material, _ := Object.As[BaseMaterial3D.Instance](g.NodeQuad.GetSurfaceOverrideMaterial(0))
	if material.BillboardMode() == BaseMaterial3D.BillboardDisabled {
		g.AsNode().SetProcess(false)
	}
}

func (g *GUIPanel3D) Process(_ Float.X) {
	// NOTE: Remove this function if you don't plan on using billboard settings.
	g.rotateAreaToBillboard()
}

func (g *GUIPanel3D) mouseEnteredArea() {
	g.isMouseInside = true
	// Notify the viewport that the mouse is now hovering it.
	g.NodeViewport.AsNode().PropagateNotification(int(Node.NotificationVpMouseEnter))
}

func (g *GUIPanel3D) mouseExitedArea() {
	// Notify the viewport that the mouse is no longer hovering it.
	g.NodeViewport.AsNode().PropagateNotification(int(Node.NotificationVpMouseExit))
	g.isMouseInside = false
}

func (g *GUIPanel3D) UnhandledInput(event InputEvent.Instance) {
	// If the event is a mouse/touch event, then we can ignore it here, because it will be
	// handled via Physics Picking.
	if _, ok := Object.As[InputEventMouseButton.Instance](event); ok {
		return
	}
	if _, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		return
	}
	if _, ok := Object.As[InputEventScreenDrag.Instance](event); ok {
		return
	}
	if _, ok := Object.As[InputEventScreenTouch.Instance](event); ok {
		return
	}
	g.NodeViewport.AsViewport().PushInput(event)
}

func (g *GUIPanel3D) mouseInputEvent(_ Node.Instance, event InputEvent.Instance, eventPosition Vector3.XYZ, _ Vector3.XYZ, _ int) {
	// Get mesh size to detect edges and make conversions. This code only supports PlaneMesh and QuadMesh.
	mesh, _ := Object.As[PlaneMesh.Instance](g.NodeQuad.Mesh())
	quadMeshSize := mesh.Size()

	// Current time in seconds since engine start.
	now := Float.X(Time.GetTicksMsec()) / 1000.0

	// Convert position to a coordinate space relative to the Area3D node.
	// NOTE: affine_inverse accounts for the Area3D node's scale, rotation, and position in the scene!
	eventPos3D := Transform3D.Transform(eventPosition, Transform3D.AffineInverse(g.NodeQuad.AsNode3D().GlobalTransform()))

	var eventPos2D Vector2.XY

	if g.isMouseInside {
		// Convert the relative event position from 3D to 2D.
		eventPos2D = Vector2.XY{X: eventPos3D.X, Y: -eventPos3D.Y}

		// Right now the event position's range is the following: (-quad_size/2) -> (quad_size/2)
		// We need to convert it into the following range: -0.5 -> 0.5
		eventPos2D.X = eventPos2D.X / quadMeshSize.X
		eventPos2D.Y = eventPos2D.Y / quadMeshSize.Y
		// Then we need to convert it into the following range: 0 -> 1
		eventPos2D.X += 0.5
		eventPos2D.Y += 0.5

		// Finally, we convert the position to the following range: 0 -> viewport.size
		viewportSize := g.NodeViewport.Size()
		eventPos2D.X *= Float.X(viewportSize.X)
		eventPos2D.Y *= Float.X(viewportSize.Y)
		// We need to do these conversions so the event's position is in the viewport's coordinate system.
	} else if g.hasLastEventPos2D {
		// Fall back to the last known event position.
		eventPos2D = g.lastEventPos2D
	}

	// Set the event's position and global position.
	if mouse, ok := Object.As[InputEventMouse.Instance](event); ok {
		mouse.SetPosition(eventPos2D)
		mouse.SetGlobalPosition(eventPos2D)
	}
	if touch, ok := Object.As[InputEventScreenTouch.Instance](event); ok {
		touch.SetPosition(eventPos2D)
	}

	// Calculate the relative event distance.
	if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		if !g.hasLastEventPos2D {
			// If there is not a stored previous position, then we'll assume there is no relative motion.
			motion.SetRelative(Vector2.Zero)
		} else {
			// If there is a stored previous position, then we'll calculate the relative position by subtracting
			// the previous position from the new position. This will give us the distance the event traveled from prev_pos.
			relative := Vector2.Sub(eventPos2D, g.lastEventPos2D)
			motion.SetRelative(relative)
			motion.SetVelocity(Vector2.DivX(relative, now-g.lastEventTime))
		}
	}
	if drag, ok := Object.As[InputEventScreenDrag.Instance](event); ok {
		drag.SetPosition(eventPos2D)
		if !g.hasLastEventPos2D {
			drag.SetRelative(Vector2.Zero)
		} else {
			relative := Vector2.Sub(eventPos2D, g.lastEventPos2D)
			drag.SetRelative(relative)
			drag.SetVelocity(Vector2.DivX(relative, now-g.lastEventTime))
		}
	}

	// Update lastEventPos2D with the position we just calculated.
	g.lastEventPos2D = eventPos2D
	g.hasLastEventPos2D = true

	// Update lastEventTime to current time.
	g.lastEventTime = now

	// Finally, send the processed input event to the viewport.
	g.NodeViewport.AsViewport().PushInput(event)
}

func (g *GUIPanel3D) rotateAreaToBillboard() {
	material, _ := Object.As[BaseMaterial3D.Instance](g.NodeQuad.GetSurfaceOverrideMaterial(0))
	billboardMode := material.BillboardMode()

	// Try to match the area with the material's billboard setting, if enabled.
	if billboardMode > 0 {
		// Get the camera.
		camera := Viewport.Get(g.AsNode()).GetCamera3d()
		// Look in the same direction as the camera.
		look := Vector3.Sub(camera.AsNode3D().ToGlobal(Vector3.XYZ{0, 0, -100}), camera.AsNode3D().GlobalTransform().Origin)
		look = Vector3.Add(g.NodeArea.AsNode3D().Position(), look)

		// Y-Billboard: Lock Y rotation, but gives bad results if the camera is tilted.
		if billboardMode == BaseMaterial3D.BillboardFixedY {
			look = Vector3.XYZ{X: look.X, Y: 0, Z: look.Z}
		}

		g.NodeArea.AsNode3D().LookAt(look)

		// Rotate in the Z axis to compensate camera tilt.
		g.NodeArea.AsNode3D().RotateObjectLocal(Vector3.Back, camera.AsNode3D().Rotation().Z)
	}
}

func main() {
	classdb.Register[GUIPanel3D]()
	startup.Scene()
}
