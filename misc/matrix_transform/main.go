package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Line2D"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/startup"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector3"
)

// AxisMarker2D draws its local axes and a line from its parent's origin to its own.
type AxisMarker2D struct {
	Node2D.Extension[AxisMarker2D] `gd:"AxisMarker2D"`
	classdb.Tool
}

func (m *AxisMarker2D) Process(_ Float.X) {
	line := Object.To[Line2D.Instance](m.AsNode().GetChild(0).GetChild(0))
	line.SetPointPosition(1, m.AsNode2D().Position())
	if parent, ok := Object.As[Node2D.Instance](m.AsNode().GetParent()); ok {
		line.AsNode2D().SetTransform(parent.AsCanvasItem().GetGlobalTransform())
	}
}

// AxisMarker3D draws its local axes and a bar from its parent's origin to its own.
type AxisMarker3D struct {
	Node3D.Extension[AxisMarker3D] `gd:"AxisMarker3D"`
	classdb.Tool
}

func (m *AxisMarker3D) Process(_ Float.X) {
	holder := Object.To[Node3D.Instance](m.AsNode().GetChild(0).GetChild(0))
	cube := Object.To[Node3D.Instance](holder.AsNode().GetChild(0))
	position := m.AsNode3D().Position()
	// "Hide" the origin vector if the AxisMarker is at (0, 0, 0)
	if position == Vector3.Zero {
		holder.SetTransform(Transform3D.New())
		cube.SetTransform(Transform3D.Scaled(Transform3D.New(), Vector3.MulX(Vector3.One, 0.0001)))
		return
	}
	transform := Transform3D.BasisOrigin{Basis: Basis.New(), Origin: Vector3.DivX(position, 2)}
	transform = Transform3D.LookingAt(transform, position, Vector3.Up)
	parent := Object.To[Node3D.Instance](m.AsNode().GetParent())
	holder.SetTransform(Transform3D.Mul(parent.GlobalTransform(), transform))
	cube.SetTransform(Transform3D.BasisOrigin{
		Basis: Basis.Scaled(Basis.New(), Vector3.XYZ{X: 0.1, Y: 0.1, Z: Vector3.Length(position)}),
	})
}

func main() {
	classdb.Register[AxisMarker2D]()
	classdb.Register[AxisMarker3D]()
	startup.Scene()
}
