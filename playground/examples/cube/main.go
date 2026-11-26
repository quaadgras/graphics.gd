// A lit, spinning cube.
package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/BoxMesh"
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/DirectionalLight3D"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

type Spinner struct {
	Node3D.Extension[Spinner]

	Speed Angle.Radians // per second
}

func (s *Spinner) Process(delta Float.X) {
	node := s.AsNode3D()
	node.RotateY(s.Speed * Angle.Radians(delta))
	node.RotateX(s.Speed * Angle.Radians(delta) / 3)
}

func main() {
	classdb.Register[Spinner]()
	startup.LoadingScene()

	cube := MeshInstance3D.New()
	cube.SetMesh(BoxMesh.New().AsMesh())
	spinner := new(Spinner)
	spinner.Speed = 1.2
	spinner.AsNode().AddChild(cube.AsNode())
	SceneTree.Add(spinner)

	camera := Camera3D.New()
	camera.AsNode3D().SetPosition(Vector3.New(0, 0.6, 2.5))
	SceneTree.Add(camera)
	camera.AsNode3D().LookAt(Vector3.Zero) // (a node looks at a point once it is in the tree.)

	light := DirectionalLight3D.New()
	light.AsNode3D().SetRotation(Euler.Radians{X: -0.8, Y: 0.6})
	SceneTree.Add(light)

	startup.Scene()
}
