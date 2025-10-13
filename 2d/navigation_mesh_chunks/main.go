package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/NavigationAgent2D"
	"graphics.gd/classdb/NavigationMeshSourceGeometryData2D"
	"graphics.gd/classdb/NavigationPolygon"
	"graphics.gd/classdb/NavigationRegion2D"
	"graphics.gd/classdb/NavigationServer2D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

var (
	map_cell_size Float.X = 1.0
	chunk_size    int     = 256
	cell_size     Float.X = 1.0
	agent_radius  Float.X = 10.0

	chunk_id_to_region = make(map[Vector2i.XY]NavigationRegion2D.Instance)
)

type NavMeshChunks struct {
	Node2D.Extension[NavMeshChunks]

	path_start_position Vector2.XY

	DebugPaths      Node2D.Instance `gd:"%DebugPaths"`
	ParseRootNode   Node2D.Instance `gd:"%ParseRootNode"`
	ChunksContainer Node2D.Instance `gd:"%ChunksContainer"`

	PathDebugCorridorFunnel   NavigationAgent2D.Instance `gd:"%PathDebugCorridorFunnel"`
	PathDebugEdgeCentered     NavigationAgent2D.Instance `gd:"%PathDebugEdgeCentered"`
	PathDebugNoPostProcessing NavigationAgent2D.Instance `gd:"%PathDebugNoPostProcessing"`
}

func (n *NavMeshChunks) Ready() {
	NavigationServer2D.SetDebugEnabled(true)
	n.path_start_position = n.DebugPaths.GlobalPosition()

	var nav_map = n.AsCanvasItem().GetWorld2d().NavigationMap()
	NavigationServer2D.MapSetCellSize(nav_map, map_cell_size)

	// Disable performance costly edge connection margin feature.
	// This feature is not needed to merge navigation mesh edges.
	// If edges are well aligned they will merge just fine by edge key.
	NavigationServer2D.MapSetUseEdgeConnections(nav_map, false)

	var source_geometry = NavigationMeshSourceGeometryData2D.New()
	var parse_settings = NavigationPolygon.New()
	parse_settings.SetParsedGeometryType(NavigationPolygon.ParsedGeometryStaticColliders)
	NavigationServer2D.ParseSourceGeometryData(parse_settings, source_geometry, n.ParseRootNode.AsNode(), nil)
	source_geometry.AddTraversableOutline([]Vector2.XY{
		{0, 0},
		{1920, 0},
		{1920, 1080},
		{0, 1080},
	})
	create_region_chunks(n.ChunksContainer.AsNode(), source_geometry, chunk_size*int(cell_size), agent_radius)
}

func (n *NavMeshChunks) Process(delta Float.X) {
	var mouse_cursor_position = n.AsCanvasItem().GetGlobalMousePosition()
	var nav_map = n.AsCanvasItem().GetWorld2d().NavigationMap()
	if NavigationServer2D.MapGetIterationId(nav_map) == 0 {
		return
	}
	var closest_point_on_navmesh = NavigationServer2D.MapGetClosestPoint(nav_map, mouse_cursor_position)
	if Input.IsMouseButtonPressed(Input.MouseButtonLeft) {
		n.path_start_position = closest_point_on_navmesh
	}
	n.DebugPaths.SetGlobalPosition(n.path_start_position)
	n.PathDebugCorridorFunnel.SetTargetPosition(closest_point_on_navmesh)
	n.PathDebugEdgeCentered.SetTargetPosition(closest_point_on_navmesh)
	n.PathDebugNoPostProcessing.SetTargetPosition(closest_point_on_navmesh)

	n.PathDebugCorridorFunnel.GetNextPathPosition()
	n.PathDebugEdgeCentered.GetNextPathPosition()
	n.PathDebugNoPostProcessing.GetNextPathPosition()
}

func create_region_chunks(chunks_root Node.Instance, source_geometry NavigationMeshSourceGeometryData2D.Instance, chunk_size int, agent_radius Float.X) {
	// We need to know how many chunks are required for the input geometry.
	// So first get an axis aligned bounding box that covers all vertices.
	var input_geometry_bounds = source_geometry.GetBounds()

	// Rasterize bounding box into chunk grid to know range of required chunks.
	var start_chunk = Vector2.Floor(Vector2.DivX(input_geometry_bounds.Position, Float.X(chunk_size)))
	var end_chunk = Vector2.Floor(Vector2.DivX(Rect2.End(input_geometry_bounds), Float.X(chunk_size)))

	for chunk_y := start_chunk.Y; chunk_y <= end_chunk.Y+1; chunk_y++ {
		for chunk_x := start_chunk.X; chunk_x <= end_chunk.X+1; chunk_x++ {
			var chunk_id = Vector2i.New(chunk_x, chunk_y)
			var chunk_bounding_box = Rect2.PositionSize{
				Vector2.MulX(Vector2.New(chunk_x, chunk_y), chunk_size),
				Vector2.New(chunk_size, chunk_size),
			}
			// We grow the chunk bounding box to include geometry
			// from all the neighbor chunks so edges can align.
			// The border size is the same value as our grow amount so
			// the final navigation mesh ends up with the intended chunk size.
			var baking_bounds = Rect2.Expand(chunk_bounding_box, chunk_size)

			var chunk_navmesh = NavigationPolygon.New()
			chunk_navmesh.SetParsedGeometryType(NavigationPolygon.ParsedGeometryStaticColliders)
			chunk_navmesh.SetBakingRect(baking_bounds)
			chunk_navmesh.SetBorderSize(Float.X(chunk_size))
			chunk_navmesh.SetAgentRadius(agent_radius)
			NavigationServer2D.BakeFromSourceGeometryData(chunk_navmesh, source_geometry, nil)

			// The only reason we reset the baking bounds here is to not render its debug.
			chunk_navmesh.SetBakingRect(Rect2.PositionSize{})

			// Snap vertex positions to avoid most rasterization issues with float precision.
			var navmesh_vertices = chunk_navmesh.Vertices()
			for i, vertex := range navmesh_vertices {
				navmesh_vertices[i] = Vector2.Snappedf(vertex, map_cell_size*0.1)
			}
			chunk_navmesh.SetVertices(navmesh_vertices)

			var chunk_region = NavigationRegion2D.New()
			chunk_region.SetNavigationPolygon(chunk_navmesh)
			chunks_root.AddChild(chunk_region.AsNode())

			chunk_id_to_region[chunk_id] = chunk_region
		}
	}
}

func main() {
	classdb.Register[NavMeshChunks]()
	startup.LoadingScene()
	Object.To[SceneTree.Instance](Engine.GetMainLoop()).SetDebugCollisionsHint(true)
	startup.Scene()
}
