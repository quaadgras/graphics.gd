// Package NavigationServer3D provides methods for working with NavigationServer3D object instances.
package NavigationServer3D

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
NavigationServer3D is the server that handles navigation maps, regions and agents. It does not handle A* navigation from [AStar3D].
Maps are divided into regions, which are composed of navigation meshes. Together, they define the navigable areas in the 3D world.
[b]Note:[/b] Most [NavigationServer3D] changes take effect after the next physics frame and not immediately. This includes all changes made to maps, regions or agents by navigation-related nodes in the scene tree or made through scripts.
For two regions to be connected to each other, they must share a similar edge. An edge is considered connected to another if both of its two vertices are at a distance less than [code]edge_connection_margin[/code] to the respective other edge's vertex.
You may assign navigation layers to regions with [method NavigationServer3D.region_set_navigation_layers], which then can be checked upon when requesting a path with [method NavigationServer3D.map_get_path]. This can be used to allow or deny certain areas for some objects.
To use the collision avoidance system, you may use agents. You can set an agent's target velocity, then the servers will emit a callback with a modified velocity.
[b]Note:[/b] The collision avoidance system ignores regions. Using the modified velocity directly may move an agent outside of the traversable area. This is a limitation of the collision avoidance system, any more complex situation may require the use of the physics engine.
This server keeps tracks of any call and executes them during the sync phase. This means that you can request any change to the map, using any thread, without worrying.
*/
var self [1]gdclass.NavigationServer3D
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.NavigationServer3D)
	self = *(*[1]gdclass.NavigationServer3D)(unsafe.Pointer(&obj))
}

/*
Returns all created navigation map [RID]s on the NavigationServer. This returns both 2D and 3D created navigation maps as there is technically no distinction between them.
*/
func GetMaps() [][]RID.NavigationMap3D { //gd:NavigationServer3D.get_maps
	once.Do(singleton)
	return [][]RID.NavigationMap3D(gd.ArrayAs[[][]RID.NavigationMap3D](gd.InternalArray(class(self).GetMaps())))
}

/*
Create a new map.
*/
func MapCreate() RID.NavigationMap3D { //gd:NavigationServer3D.map_create
	once.Do(singleton)
	return RID.NavigationMap3D(class(self).MapCreate())
}

/*
Sets the map active.
*/
func MapSetActive(mapping RID.NavigationMap3D, active bool) { //gd:NavigationServer3D.map_set_active
	once.Do(singleton)
	class(self).MapSetActive(RID.Any(mapping), active)
}

/*
Returns true if the map is active.
*/
func MapIsActive(mapping RID.NavigationMap3D) bool { //gd:NavigationServer3D.map_is_active
	once.Do(singleton)
	return bool(class(self).MapIsActive(RID.Any(mapping)))
}

/*
Sets the map up direction.
*/
func MapSetUp(mapping RID.NavigationMap3D, up Vector3.XYZ) { //gd:NavigationServer3D.map_set_up
	once.Do(singleton)
	class(self).MapSetUp(RID.Any(mapping), Vector3.XYZ(up))
}

/*
Returns the map's up direction.
*/
func MapGetUp(mapping RID.NavigationMap3D) Vector3.XYZ { //gd:NavigationServer3D.map_get_up
	once.Do(singleton)
	return Vector3.XYZ(class(self).MapGetUp(RID.Any(mapping)))
}

/*
Sets the map cell size used to rasterize the navigation mesh vertices on the XZ plane. Must match with the cell size of the used navigation meshes.
*/
func MapSetCellSize(mapping RID.NavigationMap3D, cell_size Float.X) { //gd:NavigationServer3D.map_set_cell_size
	once.Do(singleton)
	class(self).MapSetCellSize(RID.Any(mapping), float64(cell_size))
}

/*
Returns the map cell size used to rasterize the navigation mesh vertices on the XZ plane.
*/
func MapGetCellSize(mapping RID.NavigationMap3D) Float.X { //gd:NavigationServer3D.map_get_cell_size
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetCellSize(RID.Any(mapping))))
}

/*
Sets the map cell height used to rasterize the navigation mesh vertices on the Y axis. Must match with the cell height of the used navigation meshes.
*/
func MapSetCellHeight(mapping RID.NavigationMap3D, cell_height Float.X) { //gd:NavigationServer3D.map_set_cell_height
	once.Do(singleton)
	class(self).MapSetCellHeight(RID.Any(mapping), float64(cell_height))
}

/*
Returns the map cell height used to rasterize the navigation mesh vertices on the Y axis.
*/
func MapGetCellHeight(mapping RID.NavigationMap3D) Float.X { //gd:NavigationServer3D.map_get_cell_height
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetCellHeight(RID.Any(mapping))))
}

/*
Set the map's internal merge rasterizer cell scale used to control merging sensitivity.
*/
func MapSetMergeRasterizerCellScale(mapping RID.NavigationMap3D, scale Float.X) { //gd:NavigationServer3D.map_set_merge_rasterizer_cell_scale
	once.Do(singleton)
	class(self).MapSetMergeRasterizerCellScale(RID.Any(mapping), float64(scale))
}

/*
Returns map's internal merge rasterizer cell scale.
*/
func MapGetMergeRasterizerCellScale(mapping RID.NavigationMap3D) Float.X { //gd:NavigationServer3D.map_get_merge_rasterizer_cell_scale
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetMergeRasterizerCellScale(RID.Any(mapping))))
}

/*
Set the navigation [param map] edge connection use. If [param enabled] is [code]true[/code], the navigation map allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func MapSetUseEdgeConnections(mapping RID.NavigationMap3D, enabled bool) { //gd:NavigationServer3D.map_set_use_edge_connections
	once.Do(singleton)
	class(self).MapSetUseEdgeConnections(RID.Any(mapping), enabled)
}

/*
Returns true if the navigation [param map] allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func MapGetUseEdgeConnections(mapping RID.NavigationMap3D) bool { //gd:NavigationServer3D.map_get_use_edge_connections
	once.Do(singleton)
	return bool(class(self).MapGetUseEdgeConnections(RID.Any(mapping)))
}

/*
Set the map edge connection margin used to weld the compatible region edges.
*/
func MapSetEdgeConnectionMargin(mapping RID.NavigationMap3D, margin Float.X) { //gd:NavigationServer3D.map_set_edge_connection_margin
	once.Do(singleton)
	class(self).MapSetEdgeConnectionMargin(RID.Any(mapping), float64(margin))
}

/*
Returns the edge connection margin of the map. This distance is the minimum vertex distance needed to connect two edges from different regions.
*/
func MapGetEdgeConnectionMargin(mapping RID.NavigationMap3D) Float.X { //gd:NavigationServer3D.map_get_edge_connection_margin
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetEdgeConnectionMargin(RID.Any(mapping))))
}

/*
Set the map's link connection radius used to connect links to navigation polygons.
*/
func MapSetLinkConnectionRadius(mapping RID.NavigationMap3D, radius Float.X) { //gd:NavigationServer3D.map_set_link_connection_radius
	once.Do(singleton)
	class(self).MapSetLinkConnectionRadius(RID.Any(mapping), float64(radius))
}

/*
Returns the link connection radius of the map. This distance is the maximum range any link will search for navigation mesh polygons to connect to.
*/
func MapGetLinkConnectionRadius(mapping RID.NavigationMap3D) Float.X { //gd:NavigationServer3D.map_get_link_connection_radius
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetLinkConnectionRadius(RID.Any(mapping))))
}

/*
Returns the navigation path to reach the destination from the origin. [param navigation_layers] is a bitmask of all region navigation layers that are allowed to be in the path.
*/
func MapGetPath(mapping RID.NavigationMap3D, origin Vector3.XYZ, destination Vector3.XYZ, optimize bool) []Vector3.XYZ { //gd:NavigationServer3D.map_get_path
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).MapGetPath(RID.Any(mapping), Vector3.XYZ(origin), Vector3.XYZ(destination), optimize, int64(1)).Values()))
}

/*
Returns the closest point between the navigation surface and the segment.
*/
func MapGetClosestPointToSegment(mapping RID.NavigationMap3D, start Vector3.XYZ, end Vector3.XYZ) Vector3.XYZ { //gd:NavigationServer3D.map_get_closest_point_to_segment
	once.Do(singleton)
	return Vector3.XYZ(class(self).MapGetClosestPointToSegment(RID.Any(mapping), Vector3.XYZ(start), Vector3.XYZ(end), false))
}

/*
Returns the point closest to the provided [param to_point] on the navigation mesh surface.
*/
func MapGetClosestPoint(mapping RID.NavigationMap3D, to_point Vector3.XYZ) Vector3.XYZ { //gd:NavigationServer3D.map_get_closest_point
	once.Do(singleton)
	return Vector3.XYZ(class(self).MapGetClosestPoint(RID.Any(mapping), Vector3.XYZ(to_point)))
}

/*
Returns the normal for the point returned by [method map_get_closest_point].
*/
func MapGetClosestPointNormal(mapping RID.NavigationMap3D, to_point Vector3.XYZ) Vector3.XYZ { //gd:NavigationServer3D.map_get_closest_point_normal
	once.Do(singleton)
	return Vector3.XYZ(class(self).MapGetClosestPointNormal(RID.Any(mapping), Vector3.XYZ(to_point)))
}

/*
Returns the owner region RID for the point returned by [method map_get_closest_point].
*/
func MapGetClosestPointOwner(mapping RID.NavigationMap3D, to_point Vector3.XYZ) RID.NavigationRegion3D { //gd:NavigationServer3D.map_get_closest_point_owner
	once.Do(singleton)
	return RID.NavigationRegion3D(class(self).MapGetClosestPointOwner(RID.Any(mapping), Vector3.XYZ(to_point)))
}

/*
Returns all navigation link [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetLinks(mapping RID.NavigationMap3D) [][]RID.NavigationLink3D { //gd:NavigationServer3D.map_get_links
	once.Do(singleton)
	return [][]RID.NavigationLink3D(gd.ArrayAs[[][]RID.NavigationLink3D](gd.InternalArray(class(self).MapGetLinks(RID.Any(mapping)))))
}

/*
Returns all navigation regions [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetRegions(mapping RID.NavigationMap3D) [][]RID.NavigationRegion3D { //gd:NavigationServer3D.map_get_regions
	once.Do(singleton)
	return [][]RID.NavigationRegion3D(gd.ArrayAs[[][]RID.NavigationRegion3D](gd.InternalArray(class(self).MapGetRegions(RID.Any(mapping)))))
}

/*
Returns all navigation agents [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetAgents(mapping RID.NavigationMap3D) [][]RID.NavigationAgent3D { //gd:NavigationServer3D.map_get_agents
	once.Do(singleton)
	return [][]RID.NavigationAgent3D(gd.ArrayAs[[][]RID.NavigationAgent3D](gd.InternalArray(class(self).MapGetAgents(RID.Any(mapping)))))
}

/*
Returns all navigation obstacle [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetObstacles(mapping RID.NavigationMap3D) [][]RID.NavigationObstacle3D { //gd:NavigationServer3D.map_get_obstacles
	once.Do(singleton)
	return [][]RID.NavigationObstacle3D(gd.ArrayAs[[][]RID.NavigationObstacle3D](gd.InternalArray(class(self).MapGetObstacles(RID.Any(mapping)))))
}

/*
This function immediately forces synchronization of the specified navigation [param map] [RID]. By default navigation maps are only synchronized at the end of each physics frame. This function can be used to immediately (re)calculate all the navigation meshes and region connections of the navigation map. This makes it possible to query a navigation path for a changed map immediately and in the same frame (multiple times if needed).
Due to technical restrictions the current NavigationServer command queue will be flushed. This means all already queued update commands for this physics frame will be executed, even those intended for other maps, regions and agents not part of the specified map. The expensive computation of the navigation meshes and region connections of a map will only be done for the specified map. Other maps will receive the normal synchronization at the end of the physics frame. Should the specified map receive changes after the forced update it will update again as well when the other maps receive their update.
Avoidance processing and dispatch of the [code]safe_velocity[/code] signals is unaffected by this function and continues to happen for all maps and agents at the end of the physics frame.
[b]Note:[/b] With great power comes great responsibility. This function should only be used by users that really know what they are doing and have a good reason for it. Forcing an immediate update of a navigation map requires locking the NavigationServer and flushing the entire NavigationServer command queue. Not only can this severely impact the performance of a game but it can also introduce bugs if used inappropriately without much foresight.
*/
func MapForceUpdate(mapping RID.NavigationMap3D) { //gd:NavigationServer3D.map_force_update
	once.Do(singleton)
	class(self).MapForceUpdate(RID.Any(mapping))
}

/*
Returns the current iteration id of the navigation map. Every time the navigation map changes and synchronizes the iteration id increases. An iteration id of 0 means the navigation map has never synchronized.
[b]Note:[/b] The iteration id will wrap back to 1 after reaching its range limit.
*/
func MapGetIterationId(mapping RID.NavigationMap3D) int { //gd:NavigationServer3D.map_get_iteration_id
	once.Do(singleton)
	return int(int(class(self).MapGetIterationId(RID.Any(mapping))))
}

/*
Returns a random position picked from all map region polygons with matching [param navigation_layers].
If [param uniformly] is [code]true[/code], all map regions, polygons, and faces are weighted by their surface area (slower).
If [param uniformly] is [code]false[/code], just a random region and a random polygon are picked (faster).
*/
func MapGetRandomPoint(mapping RID.NavigationMap3D, navigation_layers int, uniformly bool) Vector3.XYZ { //gd:NavigationServer3D.map_get_random_point
	once.Do(singleton)
	return Vector3.XYZ(class(self).MapGetRandomPoint(RID.Any(mapping), int64(navigation_layers), uniformly))
}

/*
Queries a path in a given navigation map. Start and target position and other parameters are defined through [NavigationPathQueryParameters3D]. Updates the provided [NavigationPathQueryResult3D] result object with the path among other results requested by the query.
*/
func QueryPath(parameters [1]gdclass.NavigationPathQueryParameters3D, result [1]gdclass.NavigationPathQueryResult3D) { //gd:NavigationServer3D.query_path
	once.Do(singleton)
	class(self).QueryPath(parameters, result)
}

/*
Creates a new region.
*/
func RegionCreate() RID.NavigationRegion3D { //gd:NavigationServer3D.region_create
	once.Do(singleton)
	return RID.NavigationRegion3D(class(self).RegionCreate())
}

/*
If [param enabled] is [code]true[/code], the specified [param region] will contribute to its current navigation map.
*/
func RegionSetEnabled(region RID.NavigationRegion3D, enabled bool) { //gd:NavigationServer3D.region_set_enabled
	once.Do(singleton)
	class(self).RegionSetEnabled(RID.Any(region), enabled)
}

/*
Returns [code]true[/code] if the specified [param region] is enabled.
*/
func RegionGetEnabled(region RID.NavigationRegion3D) bool { //gd:NavigationServer3D.region_get_enabled
	once.Do(singleton)
	return bool(class(self).RegionGetEnabled(RID.Any(region)))
}

/*
If [param enabled] is [code]true[/code], the navigation [param region] will use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func RegionSetUseEdgeConnections(region RID.NavigationRegion3D, enabled bool) { //gd:NavigationServer3D.region_set_use_edge_connections
	once.Do(singleton)
	class(self).RegionSetUseEdgeConnections(RID.Any(region), enabled)
}

/*
Returns true if the navigation [param region] is set to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func RegionGetUseEdgeConnections(region RID.NavigationRegion3D) bool { //gd:NavigationServer3D.region_get_use_edge_connections
	once.Do(singleton)
	return bool(class(self).RegionGetUseEdgeConnections(RID.Any(region)))
}

/*
Sets the [param enter_cost] for this [param region].
*/
func RegionSetEnterCost(region RID.NavigationRegion3D, enter_cost Float.X) { //gd:NavigationServer3D.region_set_enter_cost
	once.Do(singleton)
	class(self).RegionSetEnterCost(RID.Any(region), float64(enter_cost))
}

/*
Returns the enter cost of this [param region].
*/
func RegionGetEnterCost(region RID.NavigationRegion3D) Float.X { //gd:NavigationServer3D.region_get_enter_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).RegionGetEnterCost(RID.Any(region))))
}

/*
Sets the [param travel_cost] for this [param region].
*/
func RegionSetTravelCost(region RID.NavigationRegion3D, travel_cost Float.X) { //gd:NavigationServer3D.region_set_travel_cost
	once.Do(singleton)
	class(self).RegionSetTravelCost(RID.Any(region), float64(travel_cost))
}

/*
Returns the travel cost of this [param region].
*/
func RegionGetTravelCost(region RID.NavigationRegion3D) Float.X { //gd:NavigationServer3D.region_get_travel_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).RegionGetTravelCost(RID.Any(region))))
}

/*
Set the [code]ObjectID[/code] of the object which manages this region.
*/
func RegionSetOwnerId(region RID.NavigationRegion3D, owner_id int) { //gd:NavigationServer3D.region_set_owner_id
	once.Do(singleton)
	class(self).RegionSetOwnerId(RID.Any(region), int64(owner_id))
}

/*
Returns the [code]ObjectID[/code] of the object which manages this region.
*/
func RegionGetOwnerId(region RID.NavigationRegion3D) int { //gd:NavigationServer3D.region_get_owner_id
	once.Do(singleton)
	return int(int(class(self).RegionGetOwnerId(RID.Any(region))))
}

/*
Returns [code]true[/code] if the provided [param point] in world space is currently owned by the provided navigation [param region]. Owned in this context means that one of the region's navigation mesh polygon faces has a possible position at the closest distance to this point compared to all other navigation meshes from other navigation regions that are also registered on the navigation map of the provided region.
If multiple navigation meshes have positions at equal distance the navigation region whose polygons are processed first wins the ownership. Polygons are processed in the same order that navigation regions were registered on the NavigationServer.
[b]Note:[/b] If navigation meshes from different navigation regions overlap (which should be avoided in general) the result might not be what is expected.
*/
func RegionOwnsPoint(region RID.NavigationRegion3D, point Vector3.XYZ) bool { //gd:NavigationServer3D.region_owns_point
	once.Do(singleton)
	return bool(class(self).RegionOwnsPoint(RID.Any(region), Vector3.XYZ(point)))
}

/*
Sets the map for the region.
*/
func RegionSetMap(region RID.NavigationRegion3D, mapping RID.NavigationMap3D) { //gd:NavigationServer3D.region_set_map
	once.Do(singleton)
	class(self).RegionSetMap(RID.Any(region), RID.Any(mapping))
}

/*
Returns the navigation map [RID] the requested [param region] is currently assigned to.
*/
func RegionGetMap(region RID.NavigationRegion3D) RID.NavigationMap3D { //gd:NavigationServer3D.region_get_map
	once.Do(singleton)
	return RID.NavigationMap3D(class(self).RegionGetMap(RID.Any(region)))
}

/*
Set the region's navigation layers. This allows selecting regions from a path request (when using [method NavigationServer3D.map_get_path]).
*/
func RegionSetNavigationLayers(region RID.NavigationRegion3D, navigation_layers int) { //gd:NavigationServer3D.region_set_navigation_layers
	once.Do(singleton)
	class(self).RegionSetNavigationLayers(RID.Any(region), int64(navigation_layers))
}

/*
Returns the region's navigation layers.
*/
func RegionGetNavigationLayers(region RID.NavigationRegion3D) int { //gd:NavigationServer3D.region_get_navigation_layers
	once.Do(singleton)
	return int(int(class(self).RegionGetNavigationLayers(RID.Any(region))))
}

/*
Sets the global transformation for the region.
*/
func RegionSetTransform(region RID.NavigationRegion3D, transform Transform3D.BasisOrigin) { //gd:NavigationServer3D.region_set_transform
	once.Do(singleton)
	class(self).RegionSetTransform(RID.Any(region), Transform3D.BasisOrigin(transform))
}

/*
Returns the global transformation of this [param region].
*/
func RegionGetTransform(region RID.NavigationRegion3D) Transform3D.BasisOrigin { //gd:NavigationServer3D.region_get_transform
	once.Do(singleton)
	return Transform3D.BasisOrigin(class(self).RegionGetTransform(RID.Any(region)))
}

/*
Sets the navigation mesh for the region.
*/
func RegionSetNavigationMesh(region RID.NavigationRegion3D, navigation_mesh [1]gdclass.NavigationMesh) { //gd:NavigationServer3D.region_set_navigation_mesh
	once.Do(singleton)
	class(self).RegionSetNavigationMesh(RID.Any(region), navigation_mesh)
}

/*
Bakes the [param navigation_mesh] with bake source geometry collected starting from the [param root_node].
*/
func RegionBakeNavigationMesh(navigation_mesh [1]gdclass.NavigationMesh, root_node [1]gdclass.Node) { //gd:NavigationServer3D.region_bake_navigation_mesh
	once.Do(singleton)
	class(self).RegionBakeNavigationMesh(navigation_mesh, root_node)
}

/*
Returns how many connections this [param region] has with other regions in the map.
*/
func RegionGetConnectionsCount(region RID.NavigationRegion3D) int { //gd:NavigationServer3D.region_get_connections_count
	once.Do(singleton)
	return int(int(class(self).RegionGetConnectionsCount(RID.Any(region))))
}

/*
Returns the starting point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
func RegionGetConnectionPathwayStart(region RID.NavigationRegion3D, connection int) Vector3.XYZ { //gd:NavigationServer3D.region_get_connection_pathway_start
	once.Do(singleton)
	return Vector3.XYZ(class(self).RegionGetConnectionPathwayStart(RID.Any(region), int64(connection)))
}

/*
Returns the ending point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
func RegionGetConnectionPathwayEnd(region RID.NavigationRegion3D, connection int) Vector3.XYZ { //gd:NavigationServer3D.region_get_connection_pathway_end
	once.Do(singleton)
	return Vector3.XYZ(class(self).RegionGetConnectionPathwayEnd(RID.Any(region), int64(connection)))
}

/*
Returns a random position picked from all region polygons with matching [param navigation_layers].
If [param uniformly] is [code]true[/code], all region polygons and faces are weighted by their surface area (slower).
If [param uniformly] is [code]false[/code], just a random polygon and face is picked (faster).
*/
func RegionGetRandomPoint(region RID.NavigationRegion3D, navigation_layers int, uniformly bool) Vector3.XYZ { //gd:NavigationServer3D.region_get_random_point
	once.Do(singleton)
	return Vector3.XYZ(class(self).RegionGetRandomPoint(RID.Any(region), int64(navigation_layers), uniformly))
}

/*
Create a new link between two positions on a map.
*/
func LinkCreate() RID.NavigationLink3D { //gd:NavigationServer3D.link_create
	once.Do(singleton)
	return RID.NavigationLink3D(class(self).LinkCreate())
}

/*
Sets the navigation map [RID] for the link.
*/
func LinkSetMap(link RID.NavigationLink3D, mapping RID.NavigationMap3D) { //gd:NavigationServer3D.link_set_map
	once.Do(singleton)
	class(self).LinkSetMap(RID.Any(link), RID.Any(mapping))
}

/*
Returns the navigation map [RID] the requested [param link] is currently assigned to.
*/
func LinkGetMap(link RID.NavigationLink3D) RID.NavigationMap3D { //gd:NavigationServer3D.link_get_map
	once.Do(singleton)
	return RID.NavigationMap3D(class(self).LinkGetMap(RID.Any(link)))
}

/*
If [param enabled] is [code]true[/code], the specified [param link] will contribute to its current navigation map.
*/
func LinkSetEnabled(link RID.NavigationLink3D, enabled bool) { //gd:NavigationServer3D.link_set_enabled
	once.Do(singleton)
	class(self).LinkSetEnabled(RID.Any(link), enabled)
}

/*
Returns [code]true[/code] if the specified [param link] is enabled.
*/
func LinkGetEnabled(link RID.NavigationLink3D) bool { //gd:NavigationServer3D.link_get_enabled
	once.Do(singleton)
	return bool(class(self).LinkGetEnabled(RID.Any(link)))
}

/*
Sets whether this [param link] can be travelled in both directions.
*/
func LinkSetBidirectional(link RID.NavigationLink3D, bidirectional bool) { //gd:NavigationServer3D.link_set_bidirectional
	once.Do(singleton)
	class(self).LinkSetBidirectional(RID.Any(link), bidirectional)
}

/*
Returns whether this [param link] can be travelled in both directions.
*/
func LinkIsBidirectional(link RID.NavigationLink3D) bool { //gd:NavigationServer3D.link_is_bidirectional
	once.Do(singleton)
	return bool(class(self).LinkIsBidirectional(RID.Any(link)))
}

/*
Set the links's navigation layers. This allows selecting links from a path request (when using [method NavigationServer3D.map_get_path]).
*/
func LinkSetNavigationLayers(link RID.NavigationLink3D, navigation_layers int) { //gd:NavigationServer3D.link_set_navigation_layers
	once.Do(singleton)
	class(self).LinkSetNavigationLayers(RID.Any(link), int64(navigation_layers))
}

/*
Returns the navigation layers for this [param link].
*/
func LinkGetNavigationLayers(link RID.NavigationLink3D) int { //gd:NavigationServer3D.link_get_navigation_layers
	once.Do(singleton)
	return int(int(class(self).LinkGetNavigationLayers(RID.Any(link))))
}

/*
Sets the entry position for this [param link].
*/
func LinkSetStartPosition(link RID.NavigationLink3D, position Vector3.XYZ) { //gd:NavigationServer3D.link_set_start_position
	once.Do(singleton)
	class(self).LinkSetStartPosition(RID.Any(link), Vector3.XYZ(position))
}

/*
Returns the starting position of this [param link].
*/
func LinkGetStartPosition(link RID.NavigationLink3D) Vector3.XYZ { //gd:NavigationServer3D.link_get_start_position
	once.Do(singleton)
	return Vector3.XYZ(class(self).LinkGetStartPosition(RID.Any(link)))
}

/*
Sets the exit position for the [param link].
*/
func LinkSetEndPosition(link RID.NavigationLink3D, position Vector3.XYZ) { //gd:NavigationServer3D.link_set_end_position
	once.Do(singleton)
	class(self).LinkSetEndPosition(RID.Any(link), Vector3.XYZ(position))
}

/*
Returns the ending position of this [param link].
*/
func LinkGetEndPosition(link RID.NavigationLink3D) Vector3.XYZ { //gd:NavigationServer3D.link_get_end_position
	once.Do(singleton)
	return Vector3.XYZ(class(self).LinkGetEndPosition(RID.Any(link)))
}

/*
Sets the [param enter_cost] for this [param link].
*/
func LinkSetEnterCost(link RID.NavigationLink3D, enter_cost Float.X) { //gd:NavigationServer3D.link_set_enter_cost
	once.Do(singleton)
	class(self).LinkSetEnterCost(RID.Any(link), float64(enter_cost))
}

/*
Returns the enter cost of this [param link].
*/
func LinkGetEnterCost(link RID.NavigationLink3D) Float.X { //gd:NavigationServer3D.link_get_enter_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).LinkGetEnterCost(RID.Any(link))))
}

/*
Sets the [param travel_cost] for this [param link].
*/
func LinkSetTravelCost(link RID.NavigationLink3D, travel_cost Float.X) { //gd:NavigationServer3D.link_set_travel_cost
	once.Do(singleton)
	class(self).LinkSetTravelCost(RID.Any(link), float64(travel_cost))
}

/*
Returns the travel cost of this [param link].
*/
func LinkGetTravelCost(link RID.NavigationLink3D) Float.X { //gd:NavigationServer3D.link_get_travel_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).LinkGetTravelCost(RID.Any(link))))
}

/*
Set the [code]ObjectID[/code] of the object which manages this link.
*/
func LinkSetOwnerId(link RID.NavigationLink3D, owner_id int) { //gd:NavigationServer3D.link_set_owner_id
	once.Do(singleton)
	class(self).LinkSetOwnerId(RID.Any(link), int64(owner_id))
}

/*
Returns the [code]ObjectID[/code] of the object which manages this link.
*/
func LinkGetOwnerId(link RID.NavigationLink3D) int { //gd:NavigationServer3D.link_get_owner_id
	once.Do(singleton)
	return int(int(class(self).LinkGetOwnerId(RID.Any(link))))
}

/*
Creates the agent.
*/
func AgentCreate() RID.NavigationAgent3D { //gd:NavigationServer3D.agent_create
	once.Do(singleton)
	return RID.NavigationAgent3D(class(self).AgentCreate())
}

/*
If [param enabled] is [code]true[/code], the provided [param agent] calculates avoidance.
*/
func AgentSetAvoidanceEnabled(agent RID.NavigationAgent3D, enabled bool) { //gd:NavigationServer3D.agent_set_avoidance_enabled
	once.Do(singleton)
	class(self).AgentSetAvoidanceEnabled(RID.Any(agent), enabled)
}

/*
Returns [code]true[/code] if the provided [param agent] has avoidance enabled.
*/
func AgentGetAvoidanceEnabled(agent RID.NavigationAgent3D) bool { //gd:NavigationServer3D.agent_get_avoidance_enabled
	once.Do(singleton)
	return bool(class(self).AgentGetAvoidanceEnabled(RID.Any(agent)))
}

/*
Sets if the agent uses the 2D avoidance or the 3D avoidance while avoidance is enabled.
If [code]true[/code] the agent calculates avoidance velocities in 3D for the xyz-axis, e.g. for games that take place in air, underwater or space. The 3D using agent only avoids other 3D avoidance using agent's. The 3D using agent only reacts to radius based avoidance obstacles. The 3D using agent ignores any vertices based obstacles. The 3D using agent only avoids other 3D using agent's.
If [code]false[/code] the agent calculates avoidance velocities in 2D along the xz-axis ignoring the y-axis. The 2D using agent only avoids other 2D avoidance using agent's. The 2D using agent reacts to radius avoidance obstacles. The 2D using agent reacts to vertices based avoidance obstacles. The 2D using agent only avoids other 2D using agent's. 2D using agents will ignore other 2D using agents or obstacles that are below their current position or above their current position including the agents height in 2D avoidance.
*/
func AgentSetUse3dAvoidance(agent RID.NavigationAgent3D, enabled bool) { //gd:NavigationServer3D.agent_set_use_3d_avoidance
	once.Do(singleton)
	class(self).AgentSetUse3dAvoidance(RID.Any(agent), enabled)
}

/*
Returns [code]true[/code] if the provided [param agent] uses avoidance in 3D space Vector3(x,y,z) instead of horizontal 2D Vector2(x,y) / Vector3(x,0.0,z).
*/
func AgentGetUse3dAvoidance(agent RID.NavigationAgent3D) bool { //gd:NavigationServer3D.agent_get_use_3d_avoidance
	once.Do(singleton)
	return bool(class(self).AgentGetUse3dAvoidance(RID.Any(agent)))
}

/*
Puts the agent in the map.
*/
func AgentSetMap(agent RID.NavigationAgent3D, mapping RID.NavigationMap3D) { //gd:NavigationServer3D.agent_set_map
	once.Do(singleton)
	class(self).AgentSetMap(RID.Any(agent), RID.Any(mapping))
}

/*
Returns the navigation map [RID] the requested [param agent] is currently assigned to.
*/
func AgentGetMap(agent RID.NavigationAgent3D) RID.NavigationMap3D { //gd:NavigationServer3D.agent_get_map
	once.Do(singleton)
	return RID.NavigationMap3D(class(self).AgentGetMap(RID.Any(agent)))
}

/*
If [param paused] is true the specified [param agent] will not be processed, e.g. calculate avoidance velocities or receive avoidance callbacks.
*/
func AgentSetPaused(agent RID.NavigationAgent3D, paused bool) { //gd:NavigationServer3D.agent_set_paused
	once.Do(singleton)
	class(self).AgentSetPaused(RID.Any(agent), paused)
}

/*
Returns [code]true[/code] if the specified [param agent] is paused.
*/
func AgentGetPaused(agent RID.NavigationAgent3D) bool { //gd:NavigationServer3D.agent_get_paused
	once.Do(singleton)
	return bool(class(self).AgentGetPaused(RID.Any(agent)))
}

/*
Sets the maximum distance to other agents this agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
func AgentSetNeighborDistance(agent RID.NavigationAgent3D, distance Float.X) { //gd:NavigationServer3D.agent_set_neighbor_distance
	once.Do(singleton)
	class(self).AgentSetNeighborDistance(RID.Any(agent), float64(distance))
}

/*
Returns the maximum distance to other agents the specified [param agent] takes into account in the navigation.
*/
func AgentGetNeighborDistance(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_neighbor_distance
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetNeighborDistance(RID.Any(agent))))
}

/*
Sets the maximum number of other agents the agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
func AgentSetMaxNeighbors(agent RID.NavigationAgent3D, count int) { //gd:NavigationServer3D.agent_set_max_neighbors
	once.Do(singleton)
	class(self).AgentSetMaxNeighbors(RID.Any(agent), int64(count))
}

/*
Returns the maximum number of other agents the specified [param agent] takes into account in the navigation.
*/
func AgentGetMaxNeighbors(agent RID.NavigationAgent3D) int { //gd:NavigationServer3D.agent_get_max_neighbors
	once.Do(singleton)
	return int(int(class(self).AgentGetMaxNeighbors(RID.Any(agent))))
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to other agents. The larger this number, the sooner this agent will respond to the presence of other agents, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
func AgentSetTimeHorizonAgents(agent RID.NavigationAgent3D, time_horizon Float.X) { //gd:NavigationServer3D.agent_set_time_horizon_agents
	once.Do(singleton)
	class(self).AgentSetTimeHorizonAgents(RID.Any(agent), float64(time_horizon))
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to other agents.
*/
func AgentGetTimeHorizonAgents(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_time_horizon_agents
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetTimeHorizonAgents(RID.Any(agent))))
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to static avoidance obstacles. The larger this number, the sooner this agent will respond to the presence of static avoidance obstacles, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
func AgentSetTimeHorizonObstacles(agent RID.NavigationAgent3D, time_horizon Float.X) { //gd:NavigationServer3D.agent_set_time_horizon_obstacles
	once.Do(singleton)
	class(self).AgentSetTimeHorizonObstacles(RID.Any(agent), float64(time_horizon))
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to static avoidance obstacles.
*/
func AgentGetTimeHorizonObstacles(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_time_horizon_obstacles
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetTimeHorizonObstacles(RID.Any(agent))))
}

/*
Sets the radius of the agent.
*/
func AgentSetRadius(agent RID.NavigationAgent3D, radius Float.X) { //gd:NavigationServer3D.agent_set_radius
	once.Do(singleton)
	class(self).AgentSetRadius(RID.Any(agent), float64(radius))
}

/*
Returns the radius of the specified [param agent].
*/
func AgentGetRadius(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_radius
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetRadius(RID.Any(agent))))
}

/*
Updates the provided [param agent] [param height].
*/
func AgentSetHeight(agent RID.NavigationAgent3D, height Float.X) { //gd:NavigationServer3D.agent_set_height
	once.Do(singleton)
	class(self).AgentSetHeight(RID.Any(agent), float64(height))
}

/*
Returns the [code]height[/code] of the specified [param agent].
*/
func AgentGetHeight(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_height
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetHeight(RID.Any(agent))))
}

/*
Sets the maximum speed of the agent. Must be positive.
*/
func AgentSetMaxSpeed(agent RID.NavigationAgent3D, max_speed Float.X) { //gd:NavigationServer3D.agent_set_max_speed
	once.Do(singleton)
	class(self).AgentSetMaxSpeed(RID.Any(agent), float64(max_speed))
}

/*
Returns the maximum speed of the specified [param agent].
*/
func AgentGetMaxSpeed(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_max_speed
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetMaxSpeed(RID.Any(agent))))
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity] for the specified [param agent]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
func AgentSetVelocityForced(agent RID.NavigationAgent3D, velocity Vector3.XYZ) { //gd:NavigationServer3D.agent_set_velocity_forced
	once.Do(singleton)
	class(self).AgentSetVelocityForced(RID.Any(agent), Vector3.XYZ(velocity))
}

/*
Sets [param velocity] as the new wanted velocity for the specified [param agent]. The avoidance simulation will try to fulfill this velocity if possible but will modify it to avoid collision with other agent's and obstacles. When an agent is teleported to a new position use [method agent_set_velocity_forced] as well to reset the internal simulation velocity.
*/
func AgentSetVelocity(agent RID.NavigationAgent3D, velocity Vector3.XYZ) { //gd:NavigationServer3D.agent_set_velocity
	once.Do(singleton)
	class(self).AgentSetVelocity(RID.Any(agent), Vector3.XYZ(velocity))
}

/*
Returns the velocity of the specified [param agent].
*/
func AgentGetVelocity(agent RID.NavigationAgent3D) Vector3.XYZ { //gd:NavigationServer3D.agent_get_velocity
	once.Do(singleton)
	return Vector3.XYZ(class(self).AgentGetVelocity(RID.Any(agent)))
}

/*
Sets the position of the agent in world space.
*/
func AgentSetPosition(agent RID.NavigationAgent3D, position Vector3.XYZ) { //gd:NavigationServer3D.agent_set_position
	once.Do(singleton)
	class(self).AgentSetPosition(RID.Any(agent), Vector3.XYZ(position))
}

/*
Returns the position of the specified [param agent] in world space.
*/
func AgentGetPosition(agent RID.NavigationAgent3D) Vector3.XYZ { //gd:NavigationServer3D.agent_get_position
	once.Do(singleton)
	return Vector3.XYZ(class(self).AgentGetPosition(RID.Any(agent)))
}

/*
Returns true if the map got changed the previous frame.
*/
func AgentIsMapChanged(agent RID.NavigationAgent3D) bool { //gd:NavigationServer3D.agent_is_map_changed
	once.Do(singleton)
	return bool(class(self).AgentIsMapChanged(RID.Any(agent)))
}

/*
Sets the callback [Callable] that gets called after each avoidance processing step for the [param agent]. The calculated [code]safe_velocity[/code] will be dispatched with a signal to the object just before the physics calculations.
[b]Note:[/b] Created callbacks are always processed independently of the SceneTree state as long as the agent is on a navigation map and not freed. To disable the dispatch of a callback from an agent use [method agent_set_avoidance_callback] again with an empty [Callable].
*/
func AgentSetAvoidanceCallback(agent RID.NavigationAgent3D, callback func(velocity Vector3.XYZ)) { //gd:NavigationServer3D.agent_set_avoidance_callback
	once.Do(singleton)
	class(self).AgentSetAvoidanceCallback(RID.Any(agent), Callable.New(callback))
}

/*
Return [code]true[/code] if the specified [param agent] has an avoidance callback.
*/
func AgentHasAvoidanceCallback(agent RID.NavigationAgent3D) bool { //gd:NavigationServer3D.agent_has_avoidance_callback
	once.Do(singleton)
	return bool(class(self).AgentHasAvoidanceCallback(RID.Any(agent)))
}

/*
Set the agent's [code]avoidance_layers[/code] bitmask.
*/
func AgentSetAvoidanceLayers(agent RID.NavigationAgent3D, layers int) { //gd:NavigationServer3D.agent_set_avoidance_layers
	once.Do(singleton)
	class(self).AgentSetAvoidanceLayers(RID.Any(agent), int64(layers))
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param agent].
*/
func AgentGetAvoidanceLayers(agent RID.NavigationAgent3D) int { //gd:NavigationServer3D.agent_get_avoidance_layers
	once.Do(singleton)
	return int(int(class(self).AgentGetAvoidanceLayers(RID.Any(agent))))
}

/*
Set the agent's [code]avoidance_mask[/code] bitmask.
*/
func AgentSetAvoidanceMask(agent RID.NavigationAgent3D, mask int) { //gd:NavigationServer3D.agent_set_avoidance_mask
	once.Do(singleton)
	class(self).AgentSetAvoidanceMask(RID.Any(agent), int64(mask))
}

/*
Returns the [code]avoidance_mask[/code] bitmask of the specified [param agent].
*/
func AgentGetAvoidanceMask(agent RID.NavigationAgent3D) int { //gd:NavigationServer3D.agent_get_avoidance_mask
	once.Do(singleton)
	return int(int(class(self).AgentGetAvoidanceMask(RID.Any(agent))))
}

/*
Set the agent's [code]avoidance_priority[/code] with a [param priority] between 0.0 (lowest priority) to 1.0 (highest priority).
The specified [param agent] does not adjust the velocity for other agents that would match the [code]avoidance_mask[/code] but have a lower [code]avoidance_priority[/code]. This in turn makes the other agents with lower priority adjust their velocities even more to avoid collision with this agent.
*/
func AgentSetAvoidancePriority(agent RID.NavigationAgent3D, priority Float.X) { //gd:NavigationServer3D.agent_set_avoidance_priority
	once.Do(singleton)
	class(self).AgentSetAvoidancePriority(RID.Any(agent), float64(priority))
}

/*
Returns the [code]avoidance_priority[/code] of the specified [param agent].
*/
func AgentGetAvoidancePriority(agent RID.NavigationAgent3D) Float.X { //gd:NavigationServer3D.agent_get_avoidance_priority
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetAvoidancePriority(RID.Any(agent))))
}

/*
Creates a new obstacle.
*/
func ObstacleCreate() RID.NavigationObstacle3D { //gd:NavigationServer3D.obstacle_create
	once.Do(singleton)
	return RID.NavigationObstacle3D(class(self).ObstacleCreate())
}

/*
If [param enabled] is [code]true[/code], the provided [param obstacle] affects avoidance using agents.
*/
func ObstacleSetAvoidanceEnabled(obstacle RID.NavigationObstacle3D, enabled bool) { //gd:NavigationServer3D.obstacle_set_avoidance_enabled
	once.Do(singleton)
	class(self).ObstacleSetAvoidanceEnabled(RID.Any(obstacle), enabled)
}

/*
Returns [code]true[/code] if the provided [param obstacle] has avoidance enabled.
*/
func ObstacleGetAvoidanceEnabled(obstacle RID.NavigationObstacle3D) bool { //gd:NavigationServer3D.obstacle_get_avoidance_enabled
	once.Do(singleton)
	return bool(class(self).ObstacleGetAvoidanceEnabled(RID.Any(obstacle)))
}

/*
Sets if the [param obstacle] uses the 2D avoidance or the 3D avoidance while avoidance is enabled.
*/
func ObstacleSetUse3dAvoidance(obstacle RID.NavigationObstacle3D, enabled bool) { //gd:NavigationServer3D.obstacle_set_use_3d_avoidance
	once.Do(singleton)
	class(self).ObstacleSetUse3dAvoidance(RID.Any(obstacle), enabled)
}

/*
Returns [code]true[/code] if the provided [param obstacle] uses avoidance in 3D space Vector3(x,y,z) instead of horizontal 2D Vector2(x,y) / Vector3(x,0.0,z).
*/
func ObstacleGetUse3dAvoidance(obstacle RID.NavigationObstacle3D) bool { //gd:NavigationServer3D.obstacle_get_use_3d_avoidance
	once.Do(singleton)
	return bool(class(self).ObstacleGetUse3dAvoidance(RID.Any(obstacle)))
}

/*
Assigns the [param obstacle] to a navigation map.
*/
func ObstacleSetMap(obstacle RID.NavigationObstacle3D, mapping RID.NavigationMap3D) { //gd:NavigationServer3D.obstacle_set_map
	once.Do(singleton)
	class(self).ObstacleSetMap(RID.Any(obstacle), RID.Any(mapping))
}

/*
Returns the navigation map [RID] the requested [param obstacle] is currently assigned to.
*/
func ObstacleGetMap(obstacle RID.NavigationObstacle3D) RID.NavigationMap3D { //gd:NavigationServer3D.obstacle_get_map
	once.Do(singleton)
	return RID.NavigationMap3D(class(self).ObstacleGetMap(RID.Any(obstacle)))
}

/*
If [param paused] is true the specified [param obstacle] will not be processed, e.g. affect avoidance velocities.
*/
func ObstacleSetPaused(obstacle RID.NavigationObstacle3D, paused bool) { //gd:NavigationServer3D.obstacle_set_paused
	once.Do(singleton)
	class(self).ObstacleSetPaused(RID.Any(obstacle), paused)
}

/*
Returns [code]true[/code] if the specified [param obstacle] is paused.
*/
func ObstacleGetPaused(obstacle RID.NavigationObstacle3D) bool { //gd:NavigationServer3D.obstacle_get_paused
	once.Do(singleton)
	return bool(class(self).ObstacleGetPaused(RID.Any(obstacle)))
}

/*
Sets the radius of the dynamic obstacle.
*/
func ObstacleSetRadius(obstacle RID.NavigationObstacle3D, radius Float.X) { //gd:NavigationServer3D.obstacle_set_radius
	once.Do(singleton)
	class(self).ObstacleSetRadius(RID.Any(obstacle), float64(radius))
}

/*
Returns the radius of the specified dynamic [param obstacle].
*/
func ObstacleGetRadius(obstacle RID.NavigationObstacle3D) Float.X { //gd:NavigationServer3D.obstacle_get_radius
	once.Do(singleton)
	return Float.X(Float.X(class(self).ObstacleGetRadius(RID.Any(obstacle))))
}

/*
Sets the [param height] for the [param obstacle]. In 3D agents will ignore obstacles that are above or below them while using 2D avoidance.
*/
func ObstacleSetHeight(obstacle RID.NavigationObstacle3D, height Float.X) { //gd:NavigationServer3D.obstacle_set_height
	once.Do(singleton)
	class(self).ObstacleSetHeight(RID.Any(obstacle), float64(height))
}

/*
Returns the [code]height[/code] of the specified [param obstacle].
*/
func ObstacleGetHeight(obstacle RID.NavigationObstacle3D) Float.X { //gd:NavigationServer3D.obstacle_get_height
	once.Do(singleton)
	return Float.X(Float.X(class(self).ObstacleGetHeight(RID.Any(obstacle))))
}

/*
Sets [param velocity] of the dynamic [param obstacle]. Allows other agents to better predict the movement of the dynamic obstacle. Only works in combination with the radius of the obstacle.
*/
func ObstacleSetVelocity(obstacle RID.NavigationObstacle3D, velocity Vector3.XYZ) { //gd:NavigationServer3D.obstacle_set_velocity
	once.Do(singleton)
	class(self).ObstacleSetVelocity(RID.Any(obstacle), Vector3.XYZ(velocity))
}

/*
Returns the velocity of the specified dynamic [param obstacle].
*/
func ObstacleGetVelocity(obstacle RID.NavigationObstacle3D) Vector3.XYZ { //gd:NavigationServer3D.obstacle_get_velocity
	once.Do(singleton)
	return Vector3.XYZ(class(self).ObstacleGetVelocity(RID.Any(obstacle)))
}

/*
Updates the [param position] in world space for the [param obstacle].
*/
func ObstacleSetPosition(obstacle RID.NavigationObstacle3D, position Vector3.XYZ) { //gd:NavigationServer3D.obstacle_set_position
	once.Do(singleton)
	class(self).ObstacleSetPosition(RID.Any(obstacle), Vector3.XYZ(position))
}

/*
Returns the position of the specified [param obstacle] in world space.
*/
func ObstacleGetPosition(obstacle RID.NavigationObstacle3D) Vector3.XYZ { //gd:NavigationServer3D.obstacle_get_position
	once.Do(singleton)
	return Vector3.XYZ(class(self).ObstacleGetPosition(RID.Any(obstacle)))
}

/*
Sets the outline vertices for the obstacle. If the vertices are winded in clockwise order agents will be pushed in by the obstacle, else they will be pushed out.
*/
func ObstacleSetVertices(obstacle RID.NavigationObstacle3D, vertices []Vector3.XYZ) { //gd:NavigationServer3D.obstacle_set_vertices
	once.Do(singleton)
	class(self).ObstacleSetVertices(RID.Any(obstacle), Packed.New(vertices...))
}

/*
Returns the outline vertices for the specified [param obstacle].
*/
func ObstacleGetVertices(obstacle RID.NavigationObstacle3D) []Vector3.XYZ { //gd:NavigationServer3D.obstacle_get_vertices
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).ObstacleGetVertices(RID.Any(obstacle)).Values()))
}

/*
Set the obstacles's [code]avoidance_layers[/code] bitmask.
*/
func ObstacleSetAvoidanceLayers(obstacle RID.NavigationObstacle3D, layers int) { //gd:NavigationServer3D.obstacle_set_avoidance_layers
	once.Do(singleton)
	class(self).ObstacleSetAvoidanceLayers(RID.Any(obstacle), int64(layers))
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param obstacle].
*/
func ObstacleGetAvoidanceLayers(obstacle RID.NavigationObstacle3D) int { //gd:NavigationServer3D.obstacle_get_avoidance_layers
	once.Do(singleton)
	return int(int(class(self).ObstacleGetAvoidanceLayers(RID.Any(obstacle))))
}

/*
Parses the [SceneTree] for source geometry according to the properties of [param navigation_mesh]. Updates the provided [param source_geometry_data] resource with the resulting data. The resource can then be used to bake a navigation mesh with [method bake_from_source_geometry_data]. After the process is finished the optional [param callback] will be called.
[b]Note:[/b] This function needs to run on the main thread or with a deferred call as the SceneTree is not thread-safe.
[b]Performance:[/b] While convenient, reading data arrays from [Mesh] resources can affect the frame rate negatively. The data needs to be received from the GPU, stalling the [RenderingServer] in the process. For performance prefer the use of e.g. collision shapes or creating the data arrays entirely in code.
*/
func ParseSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, root_node [1]gdclass.Node) { //gd:NavigationServer3D.parse_source_geometry_data
	once.Do(singleton)
	class(self).ParseSourceGeometryData(navigation_mesh, source_geometry_data, root_node, Callable.New(Callable.Nil))
}

/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
func BakeFromSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D) { //gd:NavigationServer3D.bake_from_source_geometry_data
	once.Do(singleton)
	class(self).BakeFromSourceGeometryData(navigation_mesh, source_geometry_data, Callable.New(Callable.Nil))
}

/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data] as an async task running on a background thread. After the process is finished the optional [param callback] will be called.
*/
func BakeFromSourceGeometryDataAsync(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D) { //gd:NavigationServer3D.bake_from_source_geometry_data_async
	once.Do(singleton)
	class(self).BakeFromSourceGeometryDataAsync(navigation_mesh, source_geometry_data, Callable.New(Callable.Nil))
}

/*
Returns [code]true[/code] when the provided navigation mesh is being baked on a background thread.
*/
func IsBakingNavigationMesh(navigation_mesh [1]gdclass.NavigationMesh) bool { //gd:NavigationServer3D.is_baking_navigation_mesh
	once.Do(singleton)
	return bool(class(self).IsBakingNavigationMesh(navigation_mesh))
}

/*
Creates a new source geometry parser. If a [Callable] is set for the parser with [method source_geometry_parser_set_callback] the callback will be called for every single node that gets parsed whenever [method parse_source_geometry_data] is used.
*/
func SourceGeometryParserCreate() RID.NavigationSourceGeometryParser3D { //gd:NavigationServer3D.source_geometry_parser_create
	once.Do(singleton)
	return RID.NavigationSourceGeometryParser3D(class(self).SourceGeometryParserCreate())
}

/*
Sets the [param callback] [Callable] for the specific source geometry [param parser]. The [Callable] will receive a call with the following parameters:
- [code]navigation_mesh[/code] - The [NavigationMesh] reference used to define the parse settings. Do NOT edit or add directly to the navigation mesh.
- [code]source_geometry_data[/code] - The [NavigationMeshSourceGeometryData3D] reference. Add custom source geometry for navigation mesh baking to this object.
- [code]node[/code] - The [Node] that is parsed.
*/
func SourceGeometryParserSetCallback(parser RID.NavigationSourceGeometryParser3D, callback func(navigation_mesh [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, node [1]gdclass.Node)) { //gd:NavigationServer3D.source_geometry_parser_set_callback
	once.Do(singleton)
	class(self).SourceGeometryParserSetCallback(RID.Any(parser), Callable.New(callback))
}

/*
Returns a simplified version of [param path] with less critical path points removed. The simplification amount is in worlds units and controlled by [param epsilon]. The simplification uses a variant of Ramer-Douglas-Peucker algorithm for curve point decimation.
Path simplification can be helpful to mitigate various path following issues that can arise with certain agent types and script behaviors. E.g. "steering" agents or avoidance in "open fields".
*/
func SimplifyPath(path []Vector3.XYZ, epsilon Float.X) []Vector3.XYZ { //gd:NavigationServer3D.simplify_path
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).SimplifyPath(Packed.New(path...), float64(epsilon)).Values()))
}

/*
Destroys the given RID.
*/
func FreeRid(rid RID.Any) { //gd:NavigationServer3D.free_rid
	once.Do(singleton)
	class(self).FreeRid(RID.Any(rid))
}

/*
Control activation of this server.
*/
func SetActive(active bool) { //gd:NavigationServer3D.set_active
	once.Do(singleton)
	class(self).SetActive(active)
}

/*
If [code]true[/code] enables debug mode on the NavigationServer.
*/
func SetDebugEnabled(enabled bool) { //gd:NavigationServer3D.set_debug_enabled
	once.Do(singleton)
	class(self).SetDebugEnabled(enabled)
}

/*
Returns [code]true[/code] when the NavigationServer has debug enabled.
*/
func GetDebugEnabled() bool { //gd:NavigationServer3D.get_debug_enabled
	once.Do(singleton)
	return bool(class(self).GetDebugEnabled())
}

/*
Returns information about the current state of the NavigationServer. See [enum ProcessInfo] for a list of available states.
*/
func GetProcessInfo(process_info gdclass.NavigationServer3DProcessInfo) int { //gd:NavigationServer3D.get_process_info
	once.Do(singleton)
	return int(int(class(self).GetProcessInfo(process_info)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.NavigationServer3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns all created navigation map [RID]s on the NavigationServer. This returns both 2D and 3D created navigation maps as there is technically no distinction between them.
*/
//go:nosplit
func (self class) GetMaps() Array.Contains[RID.Any] { //gd:NavigationServer3D.get_maps
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_get_maps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[RID.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Create a new map.
*/
//go:nosplit
func (self class) MapCreate() RID.Any { //gd:NavigationServer3D.map_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map active.
*/
//go:nosplit
func (self class) MapSetActive(mapping RID.Any, active bool) { //gd:NavigationServer3D.map_set_active
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns true if the map is active.
*/
//go:nosplit
func (self class) MapIsActive(mapping RID.Any) bool { //gd:NavigationServer3D.map_is_active
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map up direction.
*/
//go:nosplit
func (self class) MapSetUp(mapping RID.Any, up Vector3.XYZ) { //gd:NavigationServer3D.map_set_up
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, up)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_up, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the map's up direction.
*/
//go:nosplit
func (self class) MapGetUp(mapping RID.Any) Vector3.XYZ { //gd:NavigationServer3D.map_get_up
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_up, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map cell size used to rasterize the navigation mesh vertices on the XZ plane. Must match with the cell size of the used navigation meshes.
*/
//go:nosplit
func (self class) MapSetCellSize(mapping RID.Any, cell_size float64) { //gd:NavigationServer3D.map_set_cell_size
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, cell_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the map cell size used to rasterize the navigation mesh vertices on the XZ plane.
*/
//go:nosplit
func (self class) MapGetCellSize(mapping RID.Any) float64 { //gd:NavigationServer3D.map_get_cell_size
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map cell height used to rasterize the navigation mesh vertices on the Y axis. Must match with the cell height of the used navigation meshes.
*/
//go:nosplit
func (self class) MapSetCellHeight(mapping RID.Any, cell_height float64) { //gd:NavigationServer3D.map_set_cell_height
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, cell_height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_cell_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the map cell height used to rasterize the navigation mesh vertices on the Y axis.
*/
//go:nosplit
func (self class) MapGetCellHeight(mapping RID.Any) float64 { //gd:NavigationServer3D.map_get_cell_height
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_cell_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the map's internal merge rasterizer cell scale used to control merging sensitivity.
*/
//go:nosplit
func (self class) MapSetMergeRasterizerCellScale(mapping RID.Any, scale float64) { //gd:NavigationServer3D.map_set_merge_rasterizer_cell_scale
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_merge_rasterizer_cell_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns map's internal merge rasterizer cell scale.
*/
//go:nosplit
func (self class) MapGetMergeRasterizerCellScale(mapping RID.Any) float64 { //gd:NavigationServer3D.map_get_merge_rasterizer_cell_scale
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_merge_rasterizer_cell_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the navigation [param map] edge connection use. If [param enabled] is [code]true[/code], the navigation map allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) MapSetUseEdgeConnections(mapping RID.Any, enabled bool) { //gd:NavigationServer3D.map_set_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns true if the navigation [param map] allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) MapGetUseEdgeConnections(mapping RID.Any) bool { //gd:NavigationServer3D.map_get_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the map edge connection margin used to weld the compatible region edges.
*/
//go:nosplit
func (self class) MapSetEdgeConnectionMargin(mapping RID.Any, margin float64) { //gd:NavigationServer3D.map_set_edge_connection_margin
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_edge_connection_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the edge connection margin of the map. This distance is the minimum vertex distance needed to connect two edges from different regions.
*/
//go:nosplit
func (self class) MapGetEdgeConnectionMargin(mapping RID.Any) float64 { //gd:NavigationServer3D.map_get_edge_connection_margin
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_edge_connection_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the map's link connection radius used to connect links to navigation polygons.
*/
//go:nosplit
func (self class) MapSetLinkConnectionRadius(mapping RID.Any, radius float64) { //gd:NavigationServer3D.map_set_link_connection_radius
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_set_link_connection_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the link connection radius of the map. This distance is the maximum range any link will search for navigation mesh polygons to connect to.
*/
//go:nosplit
func (self class) MapGetLinkConnectionRadius(mapping RID.Any) float64 { //gd:NavigationServer3D.map_get_link_connection_radius
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_link_connection_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the navigation path to reach the destination from the origin. [param navigation_layers] is a bitmask of all region navigation layers that are allowed to be in the path.
*/
//go:nosplit
func (self class) MapGetPath(mapping RID.Any, origin Vector3.XYZ, destination Vector3.XYZ, optimize bool, navigation_layers int64) Packed.Array[Vector3.XYZ] { //gd:NavigationServer3D.map_get_path
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, origin)
	callframe.Arg(frame, destination)
	callframe.Arg(frame, optimize)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the closest point between the navigation surface and the segment.
*/
//go:nosplit
func (self class) MapGetClosestPointToSegment(mapping RID.Any, start Vector3.XYZ, end Vector3.XYZ, use_collision bool) Vector3.XYZ { //gd:NavigationServer3D.map_get_closest_point_to_segment
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	callframe.Arg(frame, use_collision)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_closest_point_to_segment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the point closest to the provided [param to_point] on the navigation mesh surface.
*/
//go:nosplit
func (self class) MapGetClosestPoint(mapping RID.Any, to_point Vector3.XYZ) Vector3.XYZ { //gd:NavigationServer3D.map_get_closest_point
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the normal for the point returned by [method map_get_closest_point].
*/
//go:nosplit
func (self class) MapGetClosestPointNormal(mapping RID.Any, to_point Vector3.XYZ) Vector3.XYZ { //gd:NavigationServer3D.map_get_closest_point_normal
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_closest_point_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the owner region RID for the point returned by [method map_get_closest_point].
*/
//go:nosplit
func (self class) MapGetClosestPointOwner(mapping RID.Any, to_point Vector3.XYZ) RID.Any { //gd:NavigationServer3D.map_get_closest_point_owner
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_closest_point_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns all navigation link [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetLinks(mapping RID.Any) Array.Contains[RID.Any] { //gd:NavigationServer3D.map_get_links
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_links, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[RID.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns all navigation regions [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetRegions(mapping RID.Any) Array.Contains[RID.Any] { //gd:NavigationServer3D.map_get_regions
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_regions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[RID.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns all navigation agents [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetAgents(mapping RID.Any) Array.Contains[RID.Any] { //gd:NavigationServer3D.map_get_agents
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[RID.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns all navigation obstacle [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetObstacles(mapping RID.Any) Array.Contains[RID.Any] { //gd:NavigationServer3D.map_get_obstacles
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[RID.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
This function immediately forces synchronization of the specified navigation [param map] [RID]. By default navigation maps are only synchronized at the end of each physics frame. This function can be used to immediately (re)calculate all the navigation meshes and region connections of the navigation map. This makes it possible to query a navigation path for a changed map immediately and in the same frame (multiple times if needed).
Due to technical restrictions the current NavigationServer command queue will be flushed. This means all already queued update commands for this physics frame will be executed, even those intended for other maps, regions and agents not part of the specified map. The expensive computation of the navigation meshes and region connections of a map will only be done for the specified map. Other maps will receive the normal synchronization at the end of the physics frame. Should the specified map receive changes after the forced update it will update again as well when the other maps receive their update.
Avoidance processing and dispatch of the [code]safe_velocity[/code] signals is unaffected by this function and continues to happen for all maps and agents at the end of the physics frame.
[b]Note:[/b] With great power comes great responsibility. This function should only be used by users that really know what they are doing and have a good reason for it. Forcing an immediate update of a navigation map requires locking the NavigationServer and flushing the entire NavigationServer command queue. Not only can this severely impact the performance of a game but it can also introduce bugs if used inappropriately without much foresight.
*/
//go:nosplit
func (self class) MapForceUpdate(mapping RID.Any) { //gd:NavigationServer3D.map_force_update
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_force_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current iteration id of the navigation map. Every time the navigation map changes and synchronizes the iteration id increases. An iteration id of 0 means the navigation map has never synchronized.
[b]Note:[/b] The iteration id will wrap back to 1 after reaching its range limit.
*/
//go:nosplit
func (self class) MapGetIterationId(mapping RID.Any) int64 { //gd:NavigationServer3D.map_get_iteration_id
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_iteration_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a random position picked from all map region polygons with matching [param navigation_layers].
If [param uniformly] is [code]true[/code], all map regions, polygons, and faces are weighted by their surface area (slower).
If [param uniformly] is [code]false[/code], just a random region and a random polygon are picked (faster).
*/
//go:nosplit
func (self class) MapGetRandomPoint(mapping RID.Any, navigation_layers int64, uniformly bool) Vector3.XYZ { //gd:NavigationServer3D.map_get_random_point
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, navigation_layers)
	callframe.Arg(frame, uniformly)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_map_get_random_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Queries a path in a given navigation map. Start and target position and other parameters are defined through [NavigationPathQueryParameters3D]. Updates the provided [NavigationPathQueryResult3D] result object with the path among other results requested by the query.
*/
//go:nosplit
func (self class) QueryPath(parameters [1]gdclass.NavigationPathQueryParameters3D, result [1]gdclass.NavigationPathQueryResult3D) { //gd:NavigationServer3D.query_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	callframe.Arg(frame, pointers.Get(result[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_query_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a new region.
*/
//go:nosplit
func (self class) RegionCreate() RID.Any { //gd:NavigationServer3D.region_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the specified [param region] will contribute to its current navigation map.
*/
//go:nosplit
func (self class) RegionSetEnabled(region RID.Any, enabled bool) { //gd:NavigationServer3D.region_set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param region] is enabled.
*/
//go:nosplit
func (self class) RegionGetEnabled(region RID.Any) bool { //gd:NavigationServer3D.region_get_enabled
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the navigation [param region] will use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) RegionSetUseEdgeConnections(region RID.Any, enabled bool) { //gd:NavigationServer3D.region_set_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns true if the navigation [param region] is set to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) RegionGetUseEdgeConnections(region RID.Any) bool { //gd:NavigationServer3D.region_get_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param enter_cost] for this [param region].
*/
//go:nosplit
func (self class) RegionSetEnterCost(region RID.Any, enter_cost float64) { //gd:NavigationServer3D.region_set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the enter cost of this [param region].
*/
//go:nosplit
func (self class) RegionGetEnterCost(region RID.Any) float64 { //gd:NavigationServer3D.region_get_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param travel_cost] for this [param region].
*/
//go:nosplit
func (self class) RegionSetTravelCost(region RID.Any, travel_cost float64) { //gd:NavigationServer3D.region_set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the travel cost of this [param region].
*/
//go:nosplit
func (self class) RegionGetTravelCost(region RID.Any) float64 { //gd:NavigationServer3D.region_get_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the [code]ObjectID[/code] of the object which manages this region.
*/
//go:nosplit
func (self class) RegionSetOwnerId(region RID.Any, owner_id int64) { //gd:NavigationServer3D.region_set_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]ObjectID[/code] of the object which manages this region.
*/
//go:nosplit
func (self class) RegionGetOwnerId(region RID.Any) int64 { //gd:NavigationServer3D.region_get_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the provided [param point] in world space is currently owned by the provided navigation [param region]. Owned in this context means that one of the region's navigation mesh polygon faces has a possible position at the closest distance to this point compared to all other navigation meshes from other navigation regions that are also registered on the navigation map of the provided region.
If multiple navigation meshes have positions at equal distance the navigation region whose polygons are processed first wins the ownership. Polygons are processed in the same order that navigation regions were registered on the NavigationServer.
[b]Note:[/b] If navigation meshes from different navigation regions overlap (which should be avoided in general) the result might not be what is expected.
*/
//go:nosplit
func (self class) RegionOwnsPoint(region RID.Any, point Vector3.XYZ) bool { //gd:NavigationServer3D.region_owns_point
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_owns_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map for the region.
*/
//go:nosplit
func (self class) RegionSetMap(region RID.Any, mapping RID.Any) { //gd:NavigationServer3D.region_set_map
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param region] is currently assigned to.
*/
//go:nosplit
func (self class) RegionGetMap(region RID.Any) RID.Any { //gd:NavigationServer3D.region_get_map
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the region's navigation layers. This allows selecting regions from a path request (when using [method NavigationServer3D.map_get_path]).
*/
//go:nosplit
func (self class) RegionSetNavigationLayers(region RID.Any, navigation_layers int64) { //gd:NavigationServer3D.region_set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the region's navigation layers.
*/
//go:nosplit
func (self class) RegionGetNavigationLayers(region RID.Any) int64 { //gd:NavigationServer3D.region_get_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the global transformation for the region.
*/
//go:nosplit
func (self class) RegionSetTransform(region RID.Any, transform Transform3D.BasisOrigin) { //gd:NavigationServer3D.region_set_transform
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the global transformation of this [param region].
*/
//go:nosplit
func (self class) RegionGetTransform(region RID.Any) Transform3D.BasisOrigin { //gd:NavigationServer3D.region_get_transform
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the navigation mesh for the region.
*/
//go:nosplit
func (self class) RegionSetNavigationMesh(region RID.Any, navigation_mesh [1]gdclass.NavigationMesh) { //gd:NavigationServer3D.region_set_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_set_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes the [param navigation_mesh] with bake source geometry collected starting from the [param root_node].
*/
//go:nosplit
func (self class) RegionBakeNavigationMesh(navigation_mesh [1]gdclass.NavigationMesh, root_node [1]gdclass.Node) { //gd:NavigationServer3D.region_bake_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(root_node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_bake_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns how many connections this [param region] has with other regions in the map.
*/
//go:nosplit
func (self class) RegionGetConnectionsCount(region RID.Any) int64 { //gd:NavigationServer3D.region_get_connections_count
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_connections_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the starting point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
//go:nosplit
func (self class) RegionGetConnectionPathwayStart(region RID.Any, connection int64) Vector3.XYZ { //gd:NavigationServer3D.region_get_connection_pathway_start
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, connection)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_connection_pathway_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ending point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
//go:nosplit
func (self class) RegionGetConnectionPathwayEnd(region RID.Any, connection int64) Vector3.XYZ { //gd:NavigationServer3D.region_get_connection_pathway_end
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, connection)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_connection_pathway_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a random position picked from all region polygons with matching [param navigation_layers].
If [param uniformly] is [code]true[/code], all region polygons and faces are weighted by their surface area (slower).
If [param uniformly] is [code]false[/code], just a random polygon and face is picked (faster).
*/
//go:nosplit
func (self class) RegionGetRandomPoint(region RID.Any, navigation_layers int64, uniformly bool) Vector3.XYZ { //gd:NavigationServer3D.region_get_random_point
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, navigation_layers)
	callframe.Arg(frame, uniformly)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_region_get_random_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new link between two positions on a map.
*/
//go:nosplit
func (self class) LinkCreate() RID.Any { //gd:NavigationServer3D.link_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the navigation map [RID] for the link.
*/
//go:nosplit
func (self class) LinkSetMap(link RID.Any, mapping RID.Any) { //gd:NavigationServer3D.link_set_map
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param link] is currently assigned to.
*/
//go:nosplit
func (self class) LinkGetMap(link RID.Any) RID.Any { //gd:NavigationServer3D.link_get_map
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the specified [param link] will contribute to its current navigation map.
*/
//go:nosplit
func (self class) LinkSetEnabled(link RID.Any, enabled bool) { //gd:NavigationServer3D.link_set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param link] is enabled.
*/
//go:nosplit
func (self class) LinkGetEnabled(link RID.Any) bool { //gd:NavigationServer3D.link_get_enabled
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether this [param link] can be travelled in both directions.
*/
//go:nosplit
func (self class) LinkSetBidirectional(link RID.Any, bidirectional bool) { //gd:NavigationServer3D.link_set_bidirectional
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, bidirectional)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether this [param link] can be travelled in both directions.
*/
//go:nosplit
func (self class) LinkIsBidirectional(link RID.Any) bool { //gd:NavigationServer3D.link_is_bidirectional
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the links's navigation layers. This allows selecting links from a path request (when using [method NavigationServer3D.map_get_path]).
*/
//go:nosplit
func (self class) LinkSetNavigationLayers(link RID.Any, navigation_layers int64) { //gd:NavigationServer3D.link_set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation layers for this [param link].
*/
//go:nosplit
func (self class) LinkGetNavigationLayers(link RID.Any) int64 { //gd:NavigationServer3D.link_get_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the entry position for this [param link].
*/
//go:nosplit
func (self class) LinkSetStartPosition(link RID.Any, position Vector3.XYZ) { //gd:NavigationServer3D.link_set_start_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the starting position of this [param link].
*/
//go:nosplit
func (self class) LinkGetStartPosition(link RID.Any) Vector3.XYZ { //gd:NavigationServer3D.link_get_start_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the exit position for the [param link].
*/
//go:nosplit
func (self class) LinkSetEndPosition(link RID.Any, position Vector3.XYZ) { //gd:NavigationServer3D.link_set_end_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the ending position of this [param link].
*/
//go:nosplit
func (self class) LinkGetEndPosition(link RID.Any) Vector3.XYZ { //gd:NavigationServer3D.link_get_end_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param enter_cost] for this [param link].
*/
//go:nosplit
func (self class) LinkSetEnterCost(link RID.Any, enter_cost float64) { //gd:NavigationServer3D.link_set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the enter cost of this [param link].
*/
//go:nosplit
func (self class) LinkGetEnterCost(link RID.Any) float64 { //gd:NavigationServer3D.link_get_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param travel_cost] for this [param link].
*/
//go:nosplit
func (self class) LinkSetTravelCost(link RID.Any, travel_cost float64) { //gd:NavigationServer3D.link_set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the travel cost of this [param link].
*/
//go:nosplit
func (self class) LinkGetTravelCost(link RID.Any) float64 { //gd:NavigationServer3D.link_get_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the [code]ObjectID[/code] of the object which manages this link.
*/
//go:nosplit
func (self class) LinkSetOwnerId(link RID.Any, owner_id int64) { //gd:NavigationServer3D.link_set_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_set_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]ObjectID[/code] of the object which manages this link.
*/
//go:nosplit
func (self class) LinkGetOwnerId(link RID.Any) int64 { //gd:NavigationServer3D.link_get_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_link_get_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates the agent.
*/
//go:nosplit
func (self class) AgentCreate() RID.Any { //gd:NavigationServer3D.agent_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the provided [param agent] calculates avoidance.
*/
//go:nosplit
func (self class) AgentSetAvoidanceEnabled(agent RID.Any, enabled bool) { //gd:NavigationServer3D.agent_set_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the provided [param agent] has avoidance enabled.
*/
//go:nosplit
func (self class) AgentGetAvoidanceEnabled(agent RID.Any) bool { //gd:NavigationServer3D.agent_get_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets if the agent uses the 2D avoidance or the 3D avoidance while avoidance is enabled.
If [code]true[/code] the agent calculates avoidance velocities in 3D for the xyz-axis, e.g. for games that take place in air, underwater or space. The 3D using agent only avoids other 3D avoidance using agent's. The 3D using agent only reacts to radius based avoidance obstacles. The 3D using agent ignores any vertices based obstacles. The 3D using agent only avoids other 3D using agent's.
If [code]false[/code] the agent calculates avoidance velocities in 2D along the xz-axis ignoring the y-axis. The 2D using agent only avoids other 2D avoidance using agent's. The 2D using agent reacts to radius avoidance obstacles. The 2D using agent reacts to vertices based avoidance obstacles. The 2D using agent only avoids other 2D using agent's. 2D using agents will ignore other 2D using agents or obstacles that are below their current position or above their current position including the agents height in 2D avoidance.
*/
//go:nosplit
func (self class) AgentSetUse3dAvoidance(agent RID.Any, enabled bool) { //gd:NavigationServer3D.agent_set_use_3d_avoidance
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the provided [param agent] uses avoidance in 3D space Vector3(x,y,z) instead of horizontal 2D Vector2(x,y) / Vector3(x,0.0,z).
*/
//go:nosplit
func (self class) AgentGetUse3dAvoidance(agent RID.Any) bool { //gd:NavigationServer3D.agent_get_use_3d_avoidance
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Puts the agent in the map.
*/
//go:nosplit
func (self class) AgentSetMap(agent RID.Any, mapping RID.Any) { //gd:NavigationServer3D.agent_set_map
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param agent] is currently assigned to.
*/
//go:nosplit
func (self class) AgentGetMap(agent RID.Any) RID.Any { //gd:NavigationServer3D.agent_get_map
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param paused] is true the specified [param agent] will not be processed, e.g. calculate avoidance velocities or receive avoidance callbacks.
*/
//go:nosplit
func (self class) AgentSetPaused(agent RID.Any, paused bool) { //gd:NavigationServer3D.agent_set_paused
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, paused)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param agent] is paused.
*/
//go:nosplit
func (self class) AgentGetPaused(agent RID.Any) bool { //gd:NavigationServer3D.agent_get_paused
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum distance to other agents this agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
//go:nosplit
func (self class) AgentSetNeighborDistance(agent RID.Any, distance float64) { //gd:NavigationServer3D.agent_set_neighbor_distance
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum distance to other agents the specified [param agent] takes into account in the navigation.
*/
//go:nosplit
func (self class) AgentGetNeighborDistance(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_neighbor_distance
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum number of other agents the agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
//go:nosplit
func (self class) AgentSetMaxNeighbors(agent RID.Any, count int64) { //gd:NavigationServer3D.agent_set_max_neighbors
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum number of other agents the specified [param agent] takes into account in the navigation.
*/
//go:nosplit
func (self class) AgentGetMaxNeighbors(agent RID.Any) int64 { //gd:NavigationServer3D.agent_get_max_neighbors
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to other agents. The larger this number, the sooner this agent will respond to the presence of other agents, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
//go:nosplit
func (self class) AgentSetTimeHorizonAgents(agent RID.Any, time_horizon float64) { //gd:NavigationServer3D.agent_set_time_horizon_agents
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, time_horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to other agents.
*/
//go:nosplit
func (self class) AgentGetTimeHorizonAgents(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_time_horizon_agents
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to static avoidance obstacles. The larger this number, the sooner this agent will respond to the presence of static avoidance obstacles, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
//go:nosplit
func (self class) AgentSetTimeHorizonObstacles(agent RID.Any, time_horizon float64) { //gd:NavigationServer3D.agent_set_time_horizon_obstacles
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, time_horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to static avoidance obstacles.
*/
//go:nosplit
func (self class) AgentGetTimeHorizonObstacles(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_time_horizon_obstacles
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the radius of the agent.
*/
//go:nosplit
func (self class) AgentSetRadius(agent RID.Any, radius float64) { //gd:NavigationServer3D.agent_set_radius
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the radius of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetRadius(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_radius
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the provided [param agent] [param height].
*/
//go:nosplit
func (self class) AgentSetHeight(agent RID.Any, height float64) { //gd:NavigationServer3D.agent_set_height
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]height[/code] of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetHeight(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_height
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum speed of the agent. Must be positive.
*/
//go:nosplit
func (self class) AgentSetMaxSpeed(agent RID.Any, max_speed float64) { //gd:NavigationServer3D.agent_set_max_speed
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, max_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_max_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum speed of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetMaxSpeed(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_max_speed
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_max_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity] for the specified [param agent]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
//go:nosplit
func (self class) AgentSetVelocityForced(agent RID.Any, velocity Vector3.XYZ) { //gd:NavigationServer3D.agent_set_velocity_forced
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_velocity_forced, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets [param velocity] as the new wanted velocity for the specified [param agent]. The avoidance simulation will try to fulfill this velocity if possible but will modify it to avoid collision with other agent's and obstacles. When an agent is teleported to a new position use [method agent_set_velocity_forced] as well to reset the internal simulation velocity.
*/
//go:nosplit
func (self class) AgentSetVelocity(agent RID.Any, velocity Vector3.XYZ) { //gd:NavigationServer3D.agent_set_velocity
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the velocity of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetVelocity(agent RID.Any) Vector3.XYZ { //gd:NavigationServer3D.agent_get_velocity
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the agent in world space.
*/
//go:nosplit
func (self class) AgentSetPosition(agent RID.Any, position Vector3.XYZ) { //gd:NavigationServer3D.agent_set_position
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the specified [param agent] in world space.
*/
//go:nosplit
func (self class) AgentGetPosition(agent RID.Any) Vector3.XYZ { //gd:NavigationServer3D.agent_get_position
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns true if the map got changed the previous frame.
*/
//go:nosplit
func (self class) AgentIsMapChanged(agent RID.Any) bool { //gd:NavigationServer3D.agent_is_map_changed
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_is_map_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the callback [Callable] that gets called after each avoidance processing step for the [param agent]. The calculated [code]safe_velocity[/code] will be dispatched with a signal to the object just before the physics calculations.
[b]Note:[/b] Created callbacks are always processed independently of the SceneTree state as long as the agent is on a navigation map and not freed. To disable the dispatch of a callback from an agent use [method agent_set_avoidance_callback] again with an empty [Callable].
*/
//go:nosplit
func (self class) AgentSetAvoidanceCallback(agent RID.Any, callback Callable.Function) { //gd:NavigationServer3D.agent_set_avoidance_callback
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_avoidance_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return [code]true[/code] if the specified [param agent] has an avoidance callback.
*/
//go:nosplit
func (self class) AgentHasAvoidanceCallback(agent RID.Any) bool { //gd:NavigationServer3D.agent_has_avoidance_callback
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_has_avoidance_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the agent's [code]avoidance_layers[/code] bitmask.
*/
//go:nosplit
func (self class) AgentSetAvoidanceLayers(agent RID.Any, layers int64) { //gd:NavigationServer3D.agent_set_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidanceLayers(agent RID.Any) int64 { //gd:NavigationServer3D.agent_get_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the agent's [code]avoidance_mask[/code] bitmask.
*/
//go:nosplit
func (self class) AgentSetAvoidanceMask(agent RID.Any, mask int64) { //gd:NavigationServer3D.agent_set_avoidance_mask
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_mask[/code] bitmask of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidanceMask(agent RID.Any) int64 { //gd:NavigationServer3D.agent_get_avoidance_mask
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the agent's [code]avoidance_priority[/code] with a [param priority] between 0.0 (lowest priority) to 1.0 (highest priority).
The specified [param agent] does not adjust the velocity for other agents that would match the [code]avoidance_mask[/code] but have a lower [code]avoidance_priority[/code]. This in turn makes the other agents with lower priority adjust their velocities even more to avoid collision with this agent.
*/
//go:nosplit
func (self class) AgentSetAvoidancePriority(agent RID.Any, priority float64) { //gd:NavigationServer3D.agent_set_avoidance_priority
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_set_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_priority[/code] of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidancePriority(agent RID.Any) float64 { //gd:NavigationServer3D.agent_get_avoidance_priority
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_agent_get_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new obstacle.
*/
//go:nosplit
func (self class) ObstacleCreate() RID.Any { //gd:NavigationServer3D.obstacle_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the provided [param obstacle] affects avoidance using agents.
*/
//go:nosplit
func (self class) ObstacleSetAvoidanceEnabled(obstacle RID.Any, enabled bool) { //gd:NavigationServer3D.obstacle_set_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the provided [param obstacle] has avoidance enabled.
*/
//go:nosplit
func (self class) ObstacleGetAvoidanceEnabled(obstacle RID.Any) bool { //gd:NavigationServer3D.obstacle_get_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets if the [param obstacle] uses the 2D avoidance or the 3D avoidance while avoidance is enabled.
*/
//go:nosplit
func (self class) ObstacleSetUse3dAvoidance(obstacle RID.Any, enabled bool) { //gd:NavigationServer3D.obstacle_set_use_3d_avoidance
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the provided [param obstacle] uses avoidance in 3D space Vector3(x,y,z) instead of horizontal 2D Vector2(x,y) / Vector3(x,0.0,z).
*/
//go:nosplit
func (self class) ObstacleGetUse3dAvoidance(obstacle RID.Any) bool { //gd:NavigationServer3D.obstacle_get_use_3d_avoidance
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Assigns the [param obstacle] to a navigation map.
*/
//go:nosplit
func (self class) ObstacleSetMap(obstacle RID.Any, mapping RID.Any) { //gd:NavigationServer3D.obstacle_set_map
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param obstacle] is currently assigned to.
*/
//go:nosplit
func (self class) ObstacleGetMap(obstacle RID.Any) RID.Any { //gd:NavigationServer3D.obstacle_get_map
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param paused] is true the specified [param obstacle] will not be processed, e.g. affect avoidance velocities.
*/
//go:nosplit
func (self class) ObstacleSetPaused(obstacle RID.Any, paused bool) { //gd:NavigationServer3D.obstacle_set_paused
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, paused)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param obstacle] is paused.
*/
//go:nosplit
func (self class) ObstacleGetPaused(obstacle RID.Any) bool { //gd:NavigationServer3D.obstacle_get_paused
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the radius of the dynamic obstacle.
*/
//go:nosplit
func (self class) ObstacleSetRadius(obstacle RID.Any, radius float64) { //gd:NavigationServer3D.obstacle_set_radius
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the radius of the specified dynamic [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetRadius(obstacle RID.Any) float64 { //gd:NavigationServer3D.obstacle_get_radius
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param height] for the [param obstacle]. In 3D agents will ignore obstacles that are above or below them while using 2D avoidance.
*/
//go:nosplit
func (self class) ObstacleSetHeight(obstacle RID.Any, height float64) { //gd:NavigationServer3D.obstacle_set_height
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]height[/code] of the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetHeight(obstacle RID.Any) float64 { //gd:NavigationServer3D.obstacle_get_height
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets [param velocity] of the dynamic [param obstacle]. Allows other agents to better predict the movement of the dynamic obstacle. Only works in combination with the radius of the obstacle.
*/
//go:nosplit
func (self class) ObstacleSetVelocity(obstacle RID.Any, velocity Vector3.XYZ) { //gd:NavigationServer3D.obstacle_set_velocity
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the velocity of the specified dynamic [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetVelocity(obstacle RID.Any) Vector3.XYZ { //gd:NavigationServer3D.obstacle_get_velocity
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the [param position] in world space for the [param obstacle].
*/
//go:nosplit
func (self class) ObstacleSetPosition(obstacle RID.Any, position Vector3.XYZ) { //gd:NavigationServer3D.obstacle_set_position
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the specified [param obstacle] in world space.
*/
//go:nosplit
func (self class) ObstacleGetPosition(obstacle RID.Any) Vector3.XYZ { //gd:NavigationServer3D.obstacle_get_position
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the outline vertices for the obstacle. If the vertices are winded in clockwise order agents will be pushed in by the obstacle, else they will be pushed out.
*/
//go:nosplit
func (self class) ObstacleSetVertices(obstacle RID.Any, vertices Packed.Array[Vector3.XYZ]) { //gd:NavigationServer3D.obstacle_set_vertices
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](vertices)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the outline vertices for the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetVertices(obstacle RID.Any) Packed.Array[Vector3.XYZ] { //gd:NavigationServer3D.obstacle_get_vertices
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Set the obstacles's [code]avoidance_layers[/code] bitmask.
*/
//go:nosplit
func (self class) ObstacleSetAvoidanceLayers(obstacle RID.Any, layers int64) { //gd:NavigationServer3D.obstacle_set_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetAvoidanceLayers(obstacle RID.Any) int64 { //gd:NavigationServer3D.obstacle_get_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_obstacle_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Parses the [SceneTree] for source geometry according to the properties of [param navigation_mesh]. Updates the provided [param source_geometry_data] resource with the resulting data. The resource can then be used to bake a navigation mesh with [method bake_from_source_geometry_data]. After the process is finished the optional [param callback] will be called.
[b]Note:[/b] This function needs to run on the main thread or with a deferred call as the SceneTree is not thread-safe.
[b]Performance:[/b] While convenient, reading data arrays from [Mesh] resources can affect the frame rate negatively. The data needs to be received from the GPU, stalling the [RenderingServer] in the process. For performance prefer the use of e.g. collision shapes or creating the data arrays entirely in code.
*/
//go:nosplit
func (self class) ParseSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, root_node [1]gdclass.Node, callback Callable.Function) { //gd:NavigationServer3D.parse_source_geometry_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(root_node[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_parse_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, callback Callable.Function) { //gd:NavigationServer3D.bake_from_source_geometry_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_bake_from_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data] as an async task running on a background thread. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryDataAsync(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, callback Callable.Function) { //gd:NavigationServer3D.bake_from_source_geometry_data_async
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_bake_from_source_geometry_data_async, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] when the provided navigation mesh is being baked on a background thread.
*/
//go:nosplit
func (self class) IsBakingNavigationMesh(navigation_mesh [1]gdclass.NavigationMesh) bool { //gd:NavigationServer3D.is_baking_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_is_baking_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new source geometry parser. If a [Callable] is set for the parser with [method source_geometry_parser_set_callback] the callback will be called for every single node that gets parsed whenever [method parse_source_geometry_data] is used.
*/
//go:nosplit
func (self class) SourceGeometryParserCreate() RID.Any { //gd:NavigationServer3D.source_geometry_parser_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_source_geometry_parser_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param callback] [Callable] for the specific source geometry [param parser]. The [Callable] will receive a call with the following parameters:
- [code]navigation_mesh[/code] - The [NavigationMesh] reference used to define the parse settings. Do NOT edit or add directly to the navigation mesh.
- [code]source_geometry_data[/code] - The [NavigationMeshSourceGeometryData3D] reference. Add custom source geometry for navigation mesh baking to this object.
- [code]node[/code] - The [Node] that is parsed.
*/
//go:nosplit
func (self class) SourceGeometryParserSetCallback(parser RID.Any, callback Callable.Function) { //gd:NavigationServer3D.source_geometry_parser_set_callback
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_source_geometry_parser_set_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a simplified version of [param path] with less critical path points removed. The simplification amount is in worlds units and controlled by [param epsilon]. The simplification uses a variant of Ramer-Douglas-Peucker algorithm for curve point decimation.
Path simplification can be helpful to mitigate various path following issues that can arise with certain agent types and script behaviors. E.g. "steering" agents or avoidance in "open fields".
*/
//go:nosplit
func (self class) SimplifyPath(path Packed.Array[Vector3.XYZ], epsilon float64) Packed.Array[Vector3.XYZ] { //gd:NavigationServer3D.simplify_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](path)))
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_simplify_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Destroys the given RID.
*/
//go:nosplit
func (self class) FreeRid(rid RID.Any) { //gd:NavigationServer3D.free_rid
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Control activation of this server.
*/
//go:nosplit
func (self class) SetActive(active bool) { //gd:NavigationServer3D.set_active
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [code]true[/code] enables debug mode on the NavigationServer.
*/
//go:nosplit
func (self class) SetDebugEnabled(enabled bool) { //gd:NavigationServer3D.set_debug_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_set_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] when the NavigationServer has debug enabled.
*/
//go:nosplit
func (self class) GetDebugEnabled() bool { //gd:NavigationServer3D.get_debug_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_get_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns information about the current state of the NavigationServer. See [enum ProcessInfo] for a list of available states.
*/
//go:nosplit
func (self class) GetProcessInfo(process_info gdclass.NavigationServer3DProcessInfo) int64 { //gd:NavigationServer3D.get_process_info
	var frame = callframe.New()
	callframe.Arg(frame, process_info)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer3D.Bind_get_process_info, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func OnMapChanged(cb func(mapping RID.Any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("map_changed"), gd.NewCallable(cb), 0)
}

func OnNavigationDebugChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("navigation_debug_changed"), gd.NewCallable(cb), 0)
}

func OnAvoidanceDebugChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("avoidance_debug_changed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("NavigationServer3D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationServer3D{*(*gdclass.NavigationServer3D)(unsafe.Pointer(&ptr))}
	})
}

type ProcessInfo = gdclass.NavigationServer3DProcessInfo //gd:NavigationServer3D.ProcessInfo

const (
	/*Constant to get the number of active navigation maps.*/
	InfoActiveMaps ProcessInfo = 0
	/*Constant to get the number of active navigation regions.*/
	InfoRegionCount ProcessInfo = 1
	/*Constant to get the number of active navigation agents processing avoidance.*/
	InfoAgentCount ProcessInfo = 2
	/*Constant to get the number of active navigation links.*/
	InfoLinkCount ProcessInfo = 3
	/*Constant to get the number of navigation mesh polygons.*/
	InfoPolygonCount ProcessInfo = 4
	/*Constant to get the number of navigation mesh polygon edges.*/
	InfoEdgeCount ProcessInfo = 5
	/*Constant to get the number of navigation mesh polygon edges that were merged due to edge key overlap.*/
	InfoEdgeMergeCount ProcessInfo = 6
	/*Constant to get the number of navigation mesh polygon edges that are considered connected by edge proximity.*/
	InfoEdgeConnectionCount ProcessInfo = 7
	/*Constant to get the number of navigation mesh polygon edges that could not be merged but may be still connected by edge proximity or with links.*/
	InfoEdgeFreeCount ProcessInfo = 8
)
