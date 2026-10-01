package main

import (
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/RenderingServer"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

type OcclusionDemo struct {
	Node3D.Extension[OcclusionDemo]

	OcclusionCulling Label.Instance
	MeshLOD          Label.Instance
	DrawMode         Label.Instance
	Performance      Label.Instance
}

func (d *OcclusionDemo) Input(input_event InputEvent.Instance) {
	var viewport = Viewport.Get(d.AsNode())
	if input_event.IsActionPressed("toggle_occlusion_culling") {
		viewport.SetUseOcclusionCulling(!viewport.UseOcclusionCulling())
		d.UpdateLabels()
	}
	if input_event.IsActionPressed("toggle_mesh_lod") {
		if Float.IsApproximatelyZero(viewport.MeshLodThreshold()) {
			viewport.SetMeshLodThreshold(1.0)
		} else {
			viewport.SetMeshLodThreshold(0.0)
		}
		d.UpdateLabels()
	}
	if input_event.IsActionPressed("cycle_draw_mode") {
		viewport.SetDebugDraw(Viewport.DebugDraw((int(viewport.DebugDraw()) + 1) % 5))
		d.UpdateLabels()
	}
	if input_event.IsActionPressed("toggle_vsync") {
		if DisplayServer.WindowGetVsyncMode(0) == DisplayServer.VsyncDisabled {
			DisplayServer.WindowSetVsyncMode(DisplayServer.VsyncEnabled, 0)
		} else {
			DisplayServer.WindowSetVsyncMode(DisplayServer.VsyncDisabled, 0)
		}
	}
}

func (d *OcclusionDemo) Process(_ Float.X) {
	fps := Engine.GetFramesPerSecond()
	d.Performance.SetText(fmt.Sprintf(`%d FPS (%.2f mspf)

Currently rendering:
%d objects
%dK primitive indices
%d draw calls
`,
		int(fps),
		1000.0/fps,
		RenderingServer.GetRenderingInfo(RenderingServer.RenderingInfoTotalObjectsInFrame),
		int(Float.Round(Float.X(RenderingServer.GetRenderingInfo(RenderingServer.RenderingInfoTotalPrimitivesInFrame))*0.001)),
		RenderingServer.GetRenderingInfo(RenderingServer.RenderingInfoTotalDrawCallsInFrame),
	))
}

func (d *OcclusionDemo) UpdateLabels() {
	var viewport = Viewport.Get(d.AsNode())
	if viewport.UseOcclusionCulling() {
		d.OcclusionCulling.SetText("Occlusion culling: Enabled")
	} else {
		d.OcclusionCulling.SetText("Occlusion culling: Disabled")
	}
	if !Float.IsApproximatelyZero(viewport.MeshLodThreshold()) {
		d.MeshLOD.SetText("Mesh LOD: Enabled")
	} else {
		d.MeshLOD.SetText("Mesh LOD: Disabled")
	}
	d.DrawMode.SetText("Draw mode: " + draw_mode_string(int(viewport.DebugDraw())))
}

func draw_mode_string(draw_mode int) string {
	switch draw_mode {
	case 0:
		return "Normal"
	case 1:
		return "Unshaded"
	case 2:
		return "Lighting"
	case 3:
		return "Overdraw"
	case 4:
		return "Wireframe"
	default:
		return "(unknown)"
	}
}

func main() {
	classdb.Register[OcclusionDemo]()
	classdb.Register[FlyCamera]()
	classdb.Register[Door]()
	startup.Scene()
}
