package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/Window"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
)

// Margin to keep the marker away from the screen's corners.
const Margin = 8

type Waypoint struct {
	Control.Extension[Waypoint]

	// Text is the waypoint's text.
	Text string
	// Sticky: if true, the waypoint sticks to the viewport's edges when moving off-screen.
	Sticky bool

	Label  Label.Instance
	Marker TextureRect.Instance

	camera Camera3D.Instance
	parent Node3D.Instance
}

func NewWaypoint() *Waypoint {
	return &Waypoint{Text: "Waypoint", Sticky: true}
}

func (w *Waypoint) Ready() {
	w.camera = Viewport.Get(w.AsNode()).GetCamera3d()
	parent, ok := Object.As[Node3D.Instance](w.AsNode().GetParent())
	if !ok {
		panic("The waypoint's parent node must inherit from Node3D.")
	}
	w.parent = parent
	w.Label.SetText(w.Text)
}

func (w *Waypoint) Process(_ Float.X) {
	viewport := Viewport.Get(w.AsNode())
	if !w.camera.Current() {
		// If the camera we have isn't the current one, get the current camera.
		w.camera = viewport.GetCamera3d()
	}

	parent_position := w.parent.GlobalTransform().Origin
	camera_transform := w.camera.AsNode3D().GlobalTransform()
	camera_position := camera_transform.Origin

	// We would use "camera.IsPositionBehind(parent_position)", except
	// that it also accounts for the near clip plane, which we don't want.
	is_behind := Vector3.Dot(camera_transform.Basis.Z, Vector3.Sub(parent_position, camera_position)) > 0

	// Fade the waypoint when the camera gets close.
	distance := Vector3.Distance(camera_position, parent_position)
	modulate := w.AsCanvasItem().Modulate()
	modulate.A = Float.Clamp(Float.Remap(distance, 0, 2, 0, 1), 0, 1)
	w.AsCanvasItem().SetModulate(modulate)

	unprojected_position := w.camera.UnprojectPosition(parent_position)
	// The content scale size is only valid if the stretch mode is `2d`.
	// Otherwise, the viewport size is used directly.
	var viewport_base_size Vector2.XY
	if window, ok := Object.As[Window.Instance](viewport); ok && Vector2i.Length(window.ContentScaleSize()) > 0 {
		viewport_base_size = Vector2.From(window.ContentScaleSize())
	} else {
		viewport_base_size = viewport.GetVisibleRect().Size
	}

	control := w.AsControl()
	if !w.Sticky {
		// For non-sticky waypoints, we don't need to clamp and calculate
		// the position if the waypoint goes off screen.
		control.SetPosition(unprojected_position)
		w.AsCanvasItem().SetVisible(!is_behind)
		return
	}

	// We need to handle the axes differently.
	// For the screen's X axis, the projected position is useful to us,
	// but we need to force it to the side if it's also behind.
	if is_behind {
		if unprojected_position.X < viewport_base_size.X/2 {
			unprojected_position.X = viewport_base_size.X - Margin
		} else {
			unprojected_position.X = Margin
		}
	}

	// For the screen's Y axis, the projected position is NOT useful to us
	// because we don't want to indicate to the user that they need to look
	// up or down to see something behind them. Instead, here we approximate
	// the correct position using difference of the X axis Euler angles
	// (up/down rotation) and the ratio of that with the camera's FOV.
	// This will be slightly off from the theoretical "ideal" position.
	if is_behind || unprojected_position.X < Margin || unprojected_position.X > viewport_base_size.X-Margin {
		look := Transform3D.LookingAt(camera_transform, parent_position, Vector3.Up)
		diff := Angle.Difference(Basis.AsEulerAngles(look.Basis, Angle.OrderYXZ).X, Basis.AsEulerAngles(camera_transform.Basis, Angle.OrderYXZ).X)
		unprojected_position.Y = viewport_base_size.Y * (0.5 + Float.X(diff)/Float.X(Angle.InRadians(Angle.Degrees(w.camera.Fov()))))
	}

	position := Vector2.XY{
		X: Float.Clamp(unprojected_position.X, Margin, viewport_base_size.X-Margin),
		Y: Float.Clamp(unprojected_position.Y, Margin, viewport_base_size.Y-Margin),
	}
	control.SetPosition(position)

	w.Label.AsCanvasItem().SetVisible(true)
	control.SetRotation(0)
	// Used to display a diagonal arrow when the waypoint is displayed in
	// one of the screen corners.
	var overflow Angle.Radians

	if position.X <= Margin {
		// Left overflow.
		overflow = -Angle.Tau / 8.0
		w.Label.AsCanvasItem().SetVisible(false)
		control.SetRotation(Float.X(Angle.Tau / 4.0))
	} else if position.X >= viewport_base_size.X-Margin {
		// Right overflow.
		overflow = Angle.Tau / 8.0
		w.Label.AsCanvasItem().SetVisible(false)
		control.SetRotation(Float.X(Angle.Tau * 3.0 / 4.0))
	}

	if position.Y <= Margin {
		// Top overflow.
		w.Label.AsCanvasItem().SetVisible(false)
		control.SetRotation(Float.X(Angle.Tau/2.0 + overflow))
	} else if position.Y >= viewport_base_size.Y-Margin {
		// Bottom overflow.
		w.Label.AsCanvasItem().SetVisible(false)
		control.SetRotation(Float.X(-overflow))
	}
}
