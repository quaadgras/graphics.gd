package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/NavigationAgent3D"
	"graphics.gd/classdb/NavigationMesh"
	"graphics.gd/classdb/NavigationMeshSourceGeometryData3D"
	"graphics.gd/classdb/NavigationRegion3D"
	"graphics.gd/classdb/NavigationServer3D"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/startup"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
)

var (
	map_cell_size Float.X = 0.25
	chunk_size    int     = 16
	cell_size     Float.X = 0.25
	agent_radius  Float.X = 0.5

	chunk_id_to_region = make(map[Vector3i.XYZ]NavigationRegion3D.Instance)
)

type NavMeshChunks struct {
	Node3D.Extension[NavMeshChunks]

	path_start_position Vector3.XYZ

	DebugPaths      Node3D.Instance `gd:"%DebugPaths"`
	ParseRootNode   Node3D.Instance `gd:"%ParseRootNode"`
	ChunksContainer Node3D.Instance `gd:"%ChunksContainer"`

	PathDebugCorridorFunnel   NavigationAgent3D.Instance `gd:"%PathDebugCorridorFunnel"`
	PathDebugEdgeCentered     NavigationAgent3D.Instance `gd:"%PathDebugEdgeCentered"`
	PathDebugNoPostProcessing NavigationAgent3D.Instance `gd:"%PathDebugNoPostProcessing"`
}

func (n *NavMeshChunks) Ready() {
	NavigationServer3D.SetDebugEnabled(true)

	n.path_start_position = n.DebugPaths.GlobalPosition()

	var nav_map = n.AsNode3D().GetWorld3d().NavigationMap()
	NavigationServer3D.MapSetCellSize(nav_map, map_cell_size)

	// Disable performance costly edge connection margin feature.
	// This feature is not needed to merge navigation mesh edges.
	// If edges are well aligned they will merge just fine by edge key.
	NavigationServer3D.MapSetUseEdgeConnections(nav_map, false)

	// Parse the collision shapes below our parse root node.
	var source_geometry = NavigationMeshSourceGeometryData3D.New()
	var parse_settings = NavigationMesh.New()
	parse_settings.SetGeometryParsedGeometryType(NavigationMesh.ParsedGeometryStaticColliders)
	NavigationServer3D.ParseSourceGeometryData(parse_settings, source_geometry, n.ParseRootNode.AsNode(), nil)

	create_region_chunks(n.ChunksContainer.AsNode(), source_geometry, Float.X(chunk_size)*cell_size, agent_radius)
}

func create_region_chunks(chunks_root_node Node.Instance, p_source_geometry NavigationMeshSourceGeometryData3D.Instance, p_chunk_size Float.X, p_agent_radius Float.X) {
	// We need to know how many chunks are required for the input geometry.
	// So first get an axis aligned bounding box that covers all vertices.
	var input_geometry_bounds = p_source_geometry.GetBounds()

	// Rasterize bounding box into chunk grid to know range of required chunks.
	var start_chunk = Vector3.XYZ{
		Float.Floor(input_geometry_bounds.Position.X / p_chunk_size),
		Float.Floor(input_geometry_bounds.Position.Y / p_chunk_size),
		Float.Floor(input_geometry_bounds.Position.Z / p_chunk_size),
	}
	var input_geometry_end = AABB.End(input_geometry_bounds)
	var end_chunk = Vector3.XYZ{
		Float.Floor(input_geometry_end.X / p_chunk_size),
		Float.Floor(input_geometry_end.Y / p_chunk_size),
		Float.Floor(input_geometry_end.Z / p_chunk_size),
	}

	// NavigationMesh.border_size is limited to the xz-axis.
	// So we can only bake one chunk for the y-axis and also
	// need to span the bake bounds over the entire y-axis.
	// If we dont do this we would create duplicated polygons
	// and stack them on top of each other causing merge errors.
	var bounds_min_height = start_chunk.Y
	var bounds_max_height = end_chunk.Y + p_chunk_size
	var chunk_y = 0

	for chunk_z := int(start_chunk.Z); chunk_z <= int(end_chunk.Z); chunk_z++ {
		for chunk_x := int(start_chunk.X); chunk_x <= int(end_chunk.X); chunk_x++ {
			var chunk_id = Vector3i.New(chunk_x, chunk_y, chunk_z)

			var chunk_bounding_box = AABB.PositionSize{
				Position: Vector3.MulX(Vector3.XYZ{Float.X(chunk_x), bounds_min_height, Float.X(chunk_z)}, p_chunk_size),
				Size:     Vector3.XYZ{p_chunk_size, bounds_max_height, p_chunk_size},
			}
			// We grow the chunk bounding box to include geometry
			// from all the neighbor chunks so edges can align.
			// The border size is the same value as our grow amount so
			// the final navigation mesh ends up with the intended chunk size.
			var baking_bounds = AABB.Expand(chunk_bounding_box, p_chunk_size)

			var chunk_navmesh = NavigationMesh.New()
			chunk_navmesh.SetGeometryParsedGeometryType(NavigationMesh.ParsedGeometryStaticColliders)
			chunk_navmesh.SetCellSize(cell_size)
			chunk_navmesh.SetCellHeight(cell_size)
			chunk_navmesh.SetFilterBakingAabb(baking_bounds)
			chunk_navmesh.SetBorderSize(p_chunk_size)
			chunk_navmesh.SetAgentRadius(p_agent_radius)
			NavigationServer3D.BakeFromSourceGeometryData(chunk_navmesh, p_source_geometry, nil)

			// The only reason we reset the baking bounds here is to not render its debug.
			chunk_navmesh.SetFilterBakingAabb(AABB.PositionSize{})

			// Snap vertex positions to avoid most rasterization issues with float precision.
			var navmesh_vertices = chunk_navmesh.Vertices()
			for i, vertex := range navmesh_vertices {
				navmesh_vertices[i] = Vector3.Snappedf(vertex, map_cell_size*0.1)
			}
			chunk_navmesh.SetVertices(navmesh_vertices)

			var chunk_region = NavigationRegion3D.New()
			chunk_region.SetNavigationMesh(chunk_navmesh)
			chunks_root_node.AddChild(chunk_region.AsNode())

			chunk_id_to_region[chunk_id] = chunk_region
		}
	}
}

func (n *NavMeshChunks) Process(_ Float.X) {
	var viewport = Viewport.Get(n.AsNode())
	var mouse_cursor_position = viewport.GetMousePosition()

	var nav_map = n.AsNode3D().GetWorld3d().NavigationMap()
	// Do not query when the map has never synchronized and is empty.
	if NavigationServer3D.MapGetIterationId(nav_map) == 0 {
		return
	}

	var camera = viewport.GetCamera3d()
	var camera_ray_length Float.X = 1000.0
	var camera_ray_start = camera.ProjectRayOrigin(mouse_cursor_position)
	var camera_ray_end = Vector3.Add(camera_ray_start, Vector3.MulX(camera.ProjectRayNormal(mouse_cursor_position), camera_ray_length))
	var closest_point_on_navmesh = NavigationServer3D.MapGetClosestPointToSegment(nav_map, camera_ray_start, camera_ray_end, false)

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

func main() {
	classdb.Register[NavMeshChunks]()
	startup.Scene()
}
