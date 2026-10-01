package main

import (
	"graphics.gd/classdb/BaseMaterial3D"
	"graphics.gd/classdb/ImmediateMesh"
	"graphics.gd/classdb/Mesh"
	"graphics.gd/classdb/MeshInstance3D"
	"graphics.gd/classdb/StandardMaterial3D"
	"graphics.gd/variant/Vector3"
)

type Line3D struct {
	MeshInstance3D.Extension[Line3D]

	mesh ImmediateMesh.Instance
}

func (l *Line3D) Ready() {
	l.mesh = ImmediateMesh.New()
	l.AsMeshInstance3D().SetMesh(l.mesh.AsMesh())
	var material = StandardMaterial3D.New()
	material.AsBaseMaterial3D().SetShadingMode(BaseMaterial3D.ShadingModeUnshaded)
	l.AsGeometryInstance3D().SetMaterialOverride(material.AsMaterial())
}

func (l *Line3D) DrawPath(path []Vector3.XYZ) {
	if len(path) == 0 {
		return
	}
	im := l.mesh
	im.ClearSurfaces()
	im.SurfaceBegin(Mesh.PrimitivePoints)
	im.SurfaceAddVertex(path[0])
	im.SurfaceAddVertex(path[len(path)-1])
	im.SurfaceEnd()
	im.SurfaceBegin(Mesh.PrimitiveLineStrip)
	for _, currentVector := range path {
		im.SurfaceAddVertex(currentVector)
	}
	im.SurfaceEnd()
}
