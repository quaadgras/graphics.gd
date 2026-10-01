package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/classdb/NavigationServer3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

type NavMesh struct {
	Node3D.Extension[NavMesh]

	Camera     Camera3D.Instance `gd:"CameraBase/Camera3D"`
	CameraBase Node3D.Instance
	Robot      *Character `gd:"RobotBase"`

	camRotation Angle.Radians
}

func (n *NavMesh) UnhandledInput(event InputEvent.Instance) {
	if button, ok := Object.As[InputEventMouseButton.Instance](event); ok && button.ButtonIndex() == Input.MouseButtonLeft && button.AsInputEvent().IsPressed() {
		// Get closest point on navmesh for the current mouse cursor position.
		var mouseCursorPosition = button.AsInputEventMouse().Position()
		const cameraRayLength = 1000.0
		var cameraRayStart = n.Camera.ProjectRayOrigin(mouseCursorPosition)
		var cameraRayEnd = Vector3.Add(cameraRayStart, Vector3.MulX(n.Camera.ProjectRayNormal(mouseCursorPosition), cameraRayLength))

		var closestPointOnNavmesh = NavigationServer3D.MapGetClosestPointToSegment(
			n.AsNode3D().GetWorld3d().NavigationMap(),
			cameraRayStart,
			cameraRayEnd,
			false,
		)
		n.Robot.SetTargetPosition(closestPointOnNavmesh)
	} else if motion, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		if motion.AsInputEventMouse().ButtonMask()&(Input.MouseButtonMaskMiddle|Input.MouseButtonMaskRight) != 0 {
			n.camRotation -= Angle.Radians(motion.ScreenRelative().X * 0.005)
			n.CameraBase.SetRotation(Euler.Radians{Y: n.camRotation})
		}
	}
}

func main() {
	classdb.Register[NavMesh]()
	classdb.Register[Character](NewCharacter)
	classdb.Register[Line3D]()
	startup.Scene()
}
