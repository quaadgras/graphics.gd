package main

import (
	"math/rand"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GeometryInstance3D"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector3"
)

const (
	NumTreeClusters = 2000
	Spread          = 1250
)

var TreeClusterScene = Resource.Load[PackedScene.Instance]("res://tree_cluster.tscn")

type TreeClusters struct {
	Node3D.Extension[TreeClusters]

	Loading          Control.Instance
	VisibilityRanges Label.Instance
	FadeMode         Label.Instance

	// If `false`, highest detail is always used (slower).
	visibility_ranges_enabled bool
	// `true` = use transparency fade, `false` = use hysteresis.
	fade_mode_enabled bool

	frames  int
	spawned bool
}

func NewTreeClusters() *TreeClusters {
	return &TreeClusters{
		visibility_ranges_enabled: true,
		fade_mode_enabled:         true,
	}
}

func (t *TreeClusters) Process(_ Float.X) {
	if t.spawned {
		return
	}
	t.frames++
	// Draw two frames to let the loading screen be visible.
	if t.frames < 2 {
		return
	}
	t.spawned = true

	// Use a predefined random seed for better reproducibility of results.
	random := rand.New(rand.NewSource(0x60d07))

	for range NumTreeClusters {
		tree_cluster, _ := Object.As[Node3D.Instance](TreeClusterScene.Instantiate())
		tree_cluster.SetPosition(Vector3.XYZ{
			X: Float.X(random.Float64())*(Spread*2) - Spread,
			Y: 0,
			Z: Float.X(random.Float64())*(Spread*2) - Spread,
		})
		t.AsNode().AddChild(tree_cluster.AsNode())
	}

	t.Loading.AsCanvasItem().SetVisible(false)
}

func (t *TreeClusters) group(name string) []GeometryInstance3D.Instance {
	var nodes []GeometryInstance3D.Instance
	for _, node := range SceneTree.Get(t.AsNode()).GetNodesInGroup(name) {
		if geometry, ok := Object.As[GeometryInstance3D.Instance](node); ok {
			nodes = append(nodes, geometry)
		}
	}
	return nodes
}

func (t *TreeClusters) Input(event InputEvent.Instance) {
	if event.IsActionPressed("toggle_visibility_ranges") {
		t.visibility_ranges_enabled = !t.visibility_ranges_enabled
		if t.visibility_ranges_enabled {
			t.VisibilityRanges.SetText("Visibility ranges: Enabled")
			t.VisibilityRanges.AsCanvasItem().SetModulate(Color.RGBA{R: 1, G: 1, B: 1, A: 1})
		} else {
			t.VisibilityRanges.SetText("Visibility ranges: Disabled")
			t.VisibilityRanges.AsCanvasItem().SetModulate(Color.RGBA{R: 1, G: 1, B: 0, A: 1})
		}
		t.FadeMode.AsCanvasItem().SetVisible(t.visibility_ranges_enabled)

		// When disabling visibility ranges, display the high-detail trees at any range.
		for _, node := range t.group("tree_high_detail") {
			if t.visibility_ranges_enabled {
				node.SetVisibilityRangeBegin(0)
				node.SetVisibilityRangeEnd(20)
			} else {
				node.SetVisibilityRangeBegin(0)
				node.SetVisibilityRangeEnd(0)
			}
		}
		for _, node := range t.group("tree_low_detail") {
			node.AsNode3D().SetVisible(t.visibility_ranges_enabled)
		}
		for _, node := range t.group("tree_cluster_high_detail") {
			node.AsNode3D().SetVisible(t.visibility_ranges_enabled)
		}
		for _, node := range t.group("tree_cluster_low_detail") {
			node.AsNode3D().SetVisible(t.visibility_ranges_enabled)
		}
	}

	if event.IsActionPressed("toggle_fade_mode") {
		t.fade_mode_enabled = !t.fade_mode_enabled
		if t.fade_mode_enabled {
			t.FadeMode.SetText("Fade mode: Enabled (Transparency)")
		} else {
			t.FadeMode.SetText("Fade mode: Disabled (Hysteresis)")
		}

		for _, node := range t.group("tree_high_detail") {
			if t.fade_mode_enabled {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeSelf)
			} else {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeDisabled)
			}
		}

		for _, node := range t.group("tree_low_detail") {
			if t.fade_mode_enabled {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeSelf)
				node.SetVisibilityRangeEndMargin(50)
			} else {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeDisabled)
				node.SetVisibilityRangeEndMargin(0)
			}
		}

		for _, node := range t.group("tree_cluster_high_detail") {
			if t.fade_mode_enabled {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeSelf)
				node.SetVisibilityRangeBeginMargin(50)
				node.SetVisibilityRangeEndMargin(50)
			} else {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeDisabled)
				node.SetVisibilityRangeBeginMargin(0)
				node.SetVisibilityRangeEndMargin(0)
			}
		}

		for _, node := range t.group("tree_cluster_low_detail") {
			if t.fade_mode_enabled {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeSelf)
				node.SetVisibilityRangeEndMargin(100)
			} else {
				node.SetVisibilityRangeFadeMode(GeometryInstance3D.VisibilityRangeFadeDisabled)
				node.SetVisibilityRangeEndMargin(0)
			}
		}
	}
}

func main() {
	classdb.Register[TreeClusters](NewTreeClusters)
	classdb.Register[FreeCamera]()
	classdb.Register[FPSLabel]()
	startup.Scene()
}
