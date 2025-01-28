// Package NavigationServer2D provides methods for working with NavigationServer2D object instances.
package NavigationServer2D

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Transform2D"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
NavigationServer2D is the server that handles navigation maps, regions and agents. It does not handle A* navigation from [AStar2D] or [AStarGrid2D].
Maps are divided into regions, which are composed of navigation polygons. Together, they define the traversable areas in the 2D world.
[b]Note:[/b] Most [NavigationServer2D] changes take effect after the next physics frame and not immediately. This includes all changes made to maps, regions or agents by navigation-related nodes in the scene tree or made through scripts.
For two regions to be connected to each other, they must share a similar edge. An edge is considered connected to another if both of its two vertices are at a distance less than [code]edge_connection_margin[/code] to the respective other edge's vertex.
You may assign navigation layers to regions with [method NavigationServer2D.region_set_navigation_layers], which then can be checked upon when requesting a path with [method NavigationServer2D.map_get_path]. This can be used to allow or deny certain areas for some objects.
To use the collision avoidance system, you may use agents. You can set an agent's target velocity, then the servers will emit a callback with a modified velocity.
[b]Note:[/b] The collision avoidance system ignores regions. Using the modified velocity directly may move an agent outside of the traversable area. This is a limitation of the collision avoidance system, any more complex situation may require the use of the physics engine.
This server keeps tracks of any call and executes them during the sync phase. This means that you can request any change to the map, using any thread, without worrying.
*/
var self [1]gdclass.NavigationServer2D
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.NavigationServer2D)
	self = *(*[1]gdclass.NavigationServer2D)(unsafe.Pointer(&obj))
}

/*
Returns all created navigation map [RID]s on the NavigationServer. This returns both 2D and 3D created navigation maps as there is technically no distinction between them.
*/
func GetMaps() [][]RID.NavigationMap2D { //gd:NavigationServer2D.get_maps
	once.Do(singleton)
	return [][]RID.NavigationMap2D(gd.ArrayAs[[][]RID.NavigationMap2D](gd.InternalArray(class(self).GetMaps())))
}

/*
Create a new map.
*/
func MapCreate() RID.NavigationMap2D { //gd:NavigationServer2D.map_create
	once.Do(singleton)
	return RID.NavigationMap2D(class(self).MapCreate())
}

/*
Sets the map active.
*/
func MapSetActive(mapping RID.NavigationMap2D, active bool) { //gd:NavigationServer2D.map_set_active
	once.Do(singleton)
	class(self).MapSetActive(gd.RID(mapping), active)
}

/*
Returns true if the map is active.
*/
func MapIsActive(mapping RID.NavigationMap2D) bool { //gd:NavigationServer2D.map_is_active
	once.Do(singleton)
	return bool(class(self).MapIsActive(gd.RID(mapping)))
}

/*
Sets the map cell size used to rasterize the navigation mesh vertices. Must match with the cell size of the used navigation meshes.
*/
func MapSetCellSize(mapping RID.NavigationMap2D, cell_size Float.X) { //gd:NavigationServer2D.map_set_cell_size
	once.Do(singleton)
	class(self).MapSetCellSize(gd.RID(mapping), gd.Float(cell_size))
}

/*
Returns the map cell size used to rasterize the navigation mesh vertices.
*/
func MapGetCellSize(mapping RID.NavigationMap2D) Float.X { //gd:NavigationServer2D.map_get_cell_size
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetCellSize(gd.RID(mapping))))
}

/*
Set the navigation [param map] edge connection use. If [param enabled] is [code]true[/code], the navigation map allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func MapSetUseEdgeConnections(mapping RID.NavigationMap2D, enabled bool) { //gd:NavigationServer2D.map_set_use_edge_connections
	once.Do(singleton)
	class(self).MapSetUseEdgeConnections(gd.RID(mapping), enabled)
}

/*
Returns whether the navigation [param map] allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func MapGetUseEdgeConnections(mapping RID.NavigationMap2D) bool { //gd:NavigationServer2D.map_get_use_edge_connections
	once.Do(singleton)
	return bool(class(self).MapGetUseEdgeConnections(gd.RID(mapping)))
}

/*
Set the map edge connection margin used to weld the compatible region edges.
*/
func MapSetEdgeConnectionMargin(mapping RID.NavigationMap2D, margin Float.X) { //gd:NavigationServer2D.map_set_edge_connection_margin
	once.Do(singleton)
	class(self).MapSetEdgeConnectionMargin(gd.RID(mapping), gd.Float(margin))
}

/*
Returns the edge connection margin of the map. The edge connection margin is a distance used to connect two regions.
*/
func MapGetEdgeConnectionMargin(mapping RID.NavigationMap2D) Float.X { //gd:NavigationServer2D.map_get_edge_connection_margin
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetEdgeConnectionMargin(gd.RID(mapping))))
}

/*
Set the map's link connection radius used to connect links to navigation polygons.
*/
func MapSetLinkConnectionRadius(mapping RID.NavigationMap2D, radius Float.X) { //gd:NavigationServer2D.map_set_link_connection_radius
	once.Do(singleton)
	class(self).MapSetLinkConnectionRadius(gd.RID(mapping), gd.Float(radius))
}

/*
Returns the link connection radius of the map. This distance is the maximum range any link will search for navigation mesh polygons to connect to.
*/
func MapGetLinkConnectionRadius(mapping RID.NavigationMap2D) Float.X { //gd:NavigationServer2D.map_get_link_connection_radius
	once.Do(singleton)
	return Float.X(Float.X(class(self).MapGetLinkConnectionRadius(gd.RID(mapping))))
}

/*
Returns the navigation path to reach the destination from the origin. [param navigation_layers] is a bitmask of all region navigation layers that are allowed to be in the path.
*/
func MapGetPath(mapping RID.NavigationMap2D, origin Vector2.XY, destination Vector2.XY, optimize bool) []Vector2.XY { //gd:NavigationServer2D.map_get_path
	once.Do(singleton)
	return []Vector2.XY(slices.Collect(class(self).MapGetPath(gd.RID(mapping), gd.Vector2(origin), gd.Vector2(destination), optimize, gd.Int(1)).Values()))
}

/*
Returns the point closest to the provided [param to_point] on the navigation mesh surface.
*/
func MapGetClosestPoint(mapping RID.NavigationMap2D, to_point Vector2.XY) Vector2.XY { //gd:NavigationServer2D.map_get_closest_point
	once.Do(singleton)
	return Vector2.XY(class(self).MapGetClosestPoint(gd.RID(mapping), gd.Vector2(to_point)))
}

/*
Returns the owner region RID for the point returned by [method map_get_closest_point].
*/
func MapGetClosestPointOwner(mapping RID.NavigationMap2D, to_point Vector2.XY) RID.NavigationRegion2D { //gd:NavigationServer2D.map_get_closest_point_owner
	once.Do(singleton)
	return RID.NavigationRegion2D(class(self).MapGetClosestPointOwner(gd.RID(mapping), gd.Vector2(to_point)))
}

/*
Returns all navigation link [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetLinks(mapping RID.NavigationMap2D) [][]RID.NavigationLink2D { //gd:NavigationServer2D.map_get_links
	once.Do(singleton)
	return [][]RID.NavigationLink2D(gd.ArrayAs[[][]RID.NavigationLink2D](gd.InternalArray(class(self).MapGetLinks(gd.RID(mapping)))))
}

/*
Returns all navigation regions [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetRegions(mapping RID.NavigationMap2D) [][]RID.NavigationRegion2D { //gd:NavigationServer2D.map_get_regions
	once.Do(singleton)
	return [][]RID.NavigationRegion2D(gd.ArrayAs[[][]RID.NavigationRegion2D](gd.InternalArray(class(self).MapGetRegions(gd.RID(mapping)))))
}

/*
Returns all navigation agents [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetAgents(mapping RID.NavigationMap2D) [][]RID.NavigationAgent2D { //gd:NavigationServer2D.map_get_agents
	once.Do(singleton)
	return [][]RID.NavigationAgent2D(gd.ArrayAs[[][]RID.NavigationAgent2D](gd.InternalArray(class(self).MapGetAgents(gd.RID(mapping)))))
}

/*
Returns all navigation obstacle [RID]s that are currently assigned to the requested navigation [param map].
*/
func MapGetObstacles(mapping RID.NavigationMap2D) [][]RID.NavigationObstacle2D { //gd:NavigationServer2D.map_get_obstacles
	once.Do(singleton)
	return [][]RID.NavigationObstacle2D(gd.ArrayAs[[][]RID.NavigationObstacle2D](gd.InternalArray(class(self).MapGetObstacles(gd.RID(mapping)))))
}

/*
This function immediately forces synchronization of the specified navigation [param map] [RID]. By default navigation maps are only synchronized at the end of each physics frame. This function can be used to immediately (re)calculate all the navigation meshes and region connections of the navigation map. This makes it possible to query a navigation path for a changed map immediately and in the same frame (multiple times if needed).
Due to technical restrictions the current NavigationServer command queue will be flushed. This means all already queued update commands for this physics frame will be executed, even those intended for other maps, regions and agents not part of the specified map. The expensive computation of the navigation meshes and region connections of a map will only be done for the specified map. Other maps will receive the normal synchronization at the end of the physics frame. Should the specified map receive changes after the forced update it will update again as well when the other maps receive their update.
Avoidance processing and dispatch of the [code]safe_velocity[/code] signals is unaffected by this function and continues to happen for all maps and agents at the end of the physics frame.
[b]Note:[/b] With great power comes great responsibility. This function should only be used by users that really know what they are doing and have a good reason for it. Forcing an immediate update of a navigation map requires locking the NavigationServer and flushing the entire NavigationServer command queue. Not only can this severely impact the performance of a game but it can also introduce bugs if used inappropriately without much foresight.
*/
func MapForceUpdate(mapping RID.NavigationMap2D) { //gd:NavigationServer2D.map_force_update
	once.Do(singleton)
	class(self).MapForceUpdate(gd.RID(mapping))
}

/*
Returns the current iteration id of the navigation map. Every time the navigation map changes and synchronizes the iteration id increases. An iteration id of 0 means the navigation map has never synchronized.
[b]Note:[/b] The iteration id will wrap back to 1 after reaching its range limit.
*/
func MapGetIterationId(mapping RID.NavigationMap2D) int { //gd:NavigationServer2D.map_get_iteration_id
	once.Do(singleton)
	return int(int(class(self).MapGetIterationId(gd.RID(mapping))))
}

/*
Returns a random position picked from all map region polygons with matching [param navigation_layers].
If [param uniformly] is [code]true[/code], all map regions, polygons, and faces are weighted by their surface area (slower).
If [param uniformly] is [code]false[/code], just a random region and a random polygon are picked (faster).
*/
func MapGetRandomPoint(mapping RID.NavigationMap2D, navigation_layers int, uniformly bool) Vector2.XY { //gd:NavigationServer2D.map_get_random_point
	once.Do(singleton)
	return Vector2.XY(class(self).MapGetRandomPoint(gd.RID(mapping), gd.Int(navigation_layers), uniformly))
}

/*
Queries a path in a given navigation map. Start and target position and other parameters are defined through [NavigationPathQueryParameters2D]. Updates the provided [NavigationPathQueryResult2D] result object with the path among other results requested by the query.
*/
func QueryPath(parameters [1]gdclass.NavigationPathQueryParameters2D, result [1]gdclass.NavigationPathQueryResult2D) { //gd:NavigationServer2D.query_path
	once.Do(singleton)
	class(self).QueryPath(parameters, result)
}

/*
Creates a new region.
*/
func RegionCreate() RID.NavigationRegion2D { //gd:NavigationServer2D.region_create
	once.Do(singleton)
	return RID.NavigationRegion2D(class(self).RegionCreate())
}

/*
If [param enabled] is [code]true[/code] the specified [param region] will contribute to its current navigation map.
*/
func RegionSetEnabled(region RID.NavigationRegion2D, enabled bool) { //gd:NavigationServer2D.region_set_enabled
	once.Do(singleton)
	class(self).RegionSetEnabled(gd.RID(region), enabled)
}

/*
Returns [code]true[/code] if the specified [param region] is enabled.
*/
func RegionGetEnabled(region RID.NavigationRegion2D) bool { //gd:NavigationServer2D.region_get_enabled
	once.Do(singleton)
	return bool(class(self).RegionGetEnabled(gd.RID(region)))
}

/*
If [param enabled] is [code]true[/code], the navigation [param region] will use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func RegionSetUseEdgeConnections(region RID.NavigationRegion2D, enabled bool) { //gd:NavigationServer2D.region_set_use_edge_connections
	once.Do(singleton)
	class(self).RegionSetUseEdgeConnections(gd.RID(region), enabled)
}

/*
Returns whether the navigation [param region] is set to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
func RegionGetUseEdgeConnections(region RID.NavigationRegion2D) bool { //gd:NavigationServer2D.region_get_use_edge_connections
	once.Do(singleton)
	return bool(class(self).RegionGetUseEdgeConnections(gd.RID(region)))
}

/*
Sets the [param enter_cost] for this [param region].
*/
func RegionSetEnterCost(region RID.NavigationRegion2D, enter_cost Float.X) { //gd:NavigationServer2D.region_set_enter_cost
	once.Do(singleton)
	class(self).RegionSetEnterCost(gd.RID(region), gd.Float(enter_cost))
}

/*
Returns the enter cost of this [param region].
*/
func RegionGetEnterCost(region RID.NavigationRegion2D) Float.X { //gd:NavigationServer2D.region_get_enter_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).RegionGetEnterCost(gd.RID(region))))
}

/*
Sets the [param travel_cost] for this [param region].
*/
func RegionSetTravelCost(region RID.NavigationRegion2D, travel_cost Float.X) { //gd:NavigationServer2D.region_set_travel_cost
	once.Do(singleton)
	class(self).RegionSetTravelCost(gd.RID(region), gd.Float(travel_cost))
}

/*
Returns the travel cost of this [param region].
*/
func RegionGetTravelCost(region RID.NavigationRegion2D) Float.X { //gd:NavigationServer2D.region_get_travel_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).RegionGetTravelCost(gd.RID(region))))
}

/*
Set the [code]ObjectID[/code] of the object which manages this region.
*/
func RegionSetOwnerId(region RID.NavigationRegion2D, owner_id int) { //gd:NavigationServer2D.region_set_owner_id
	once.Do(singleton)
	class(self).RegionSetOwnerId(gd.RID(region), gd.Int(owner_id))
}

/*
Returns the [code]ObjectID[/code] of the object which manages this region.
*/
func RegionGetOwnerId(region RID.NavigationRegion2D) int { //gd:NavigationServer2D.region_get_owner_id
	once.Do(singleton)
	return int(int(class(self).RegionGetOwnerId(gd.RID(region))))
}

/*
Returns [code]true[/code] if the provided [param point] in world space is currently owned by the provided navigation [param region]. Owned in this context means that one of the region's navigation mesh polygon faces has a possible position at the closest distance to this point compared to all other navigation meshes from other navigation regions that are also registered on the navigation map of the provided region.
If multiple navigation meshes have positions at equal distance the navigation region whose polygons are processed first wins the ownership. Polygons are processed in the same order that navigation regions were registered on the NavigationServer.
[b]Note:[/b] If navigation meshes from different navigation regions overlap (which should be avoided in general) the result might not be what is expected.
*/
func RegionOwnsPoint(region RID.NavigationRegion2D, point Vector2.XY) bool { //gd:NavigationServer2D.region_owns_point
	once.Do(singleton)
	return bool(class(self).RegionOwnsPoint(gd.RID(region), gd.Vector2(point)))
}

/*
Sets the map for the region.
*/
func RegionSetMap(region RID.NavigationRegion2D, mapping RID.NavigationMap2D) { //gd:NavigationServer2D.region_set_map
	once.Do(singleton)
	class(self).RegionSetMap(gd.RID(region), gd.RID(mapping))
}

/*
Returns the navigation map [RID] the requested [param region] is currently assigned to.
*/
func RegionGetMap(region RID.NavigationRegion2D) RID.NavigationMap2D { //gd:NavigationServer2D.region_get_map
	once.Do(singleton)
	return RID.NavigationMap2D(class(self).RegionGetMap(gd.RID(region)))
}

/*
Set the region's navigation layers. This allows selecting regions from a path request (when using [method NavigationServer2D.map_get_path]).
*/
func RegionSetNavigationLayers(region RID.NavigationRegion2D, navigation_layers int) { //gd:NavigationServer2D.region_set_navigation_layers
	once.Do(singleton)
	class(self).RegionSetNavigationLayers(gd.RID(region), gd.Int(navigation_layers))
}

/*
Returns the region's navigation layers.
*/
func RegionGetNavigationLayers(region RID.NavigationRegion2D) int { //gd:NavigationServer2D.region_get_navigation_layers
	once.Do(singleton)
	return int(int(class(self).RegionGetNavigationLayers(gd.RID(region))))
}

/*
Sets the global transformation for the region.
*/
func RegionSetTransform(region RID.NavigationRegion2D, transform Transform2D.OriginXY) { //gd:NavigationServer2D.region_set_transform
	once.Do(singleton)
	class(self).RegionSetTransform(gd.RID(region), gd.Transform2D(transform))
}

/*
Returns the global transformation of this [param region].
*/
func RegionGetTransform(region RID.NavigationRegion2D) Transform2D.OriginXY { //gd:NavigationServer2D.region_get_transform
	once.Do(singleton)
	return Transform2D.OriginXY(class(self).RegionGetTransform(gd.RID(region)))
}

/*
Sets the [param navigation_polygon] for the region.
*/
func RegionSetNavigationPolygon(region RID.NavigationRegion2D, navigation_polygon [1]gdclass.NavigationPolygon) { //gd:NavigationServer2D.region_set_navigation_polygon
	once.Do(singleton)
	class(self).RegionSetNavigationPolygon(gd.RID(region), navigation_polygon)
}

/*
Returns how many connections this [param region] has with other regions in the map.
*/
func RegionGetConnectionsCount(region RID.NavigationRegion2D) int { //gd:NavigationServer2D.region_get_connections_count
	once.Do(singleton)
	return int(int(class(self).RegionGetConnectionsCount(gd.RID(region))))
}

/*
Returns the starting point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
func RegionGetConnectionPathwayStart(region RID.NavigationRegion2D, connection int) Vector2.XY { //gd:NavigationServer2D.region_get_connection_pathway_start
	once.Do(singleton)
	return Vector2.XY(class(self).RegionGetConnectionPathwayStart(gd.RID(region), gd.Int(connection)))
}

/*
Returns the ending point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
func RegionGetConnectionPathwayEnd(region RID.NavigationRegion2D, connection int) Vector2.XY { //gd:NavigationServer2D.region_get_connection_pathway_end
	once.Do(singleton)
	return Vector2.XY(class(self).RegionGetConnectionPathwayEnd(gd.RID(region), gd.Int(connection)))
}

/*
Returns a random position picked from all region polygons with matching [param navigation_layers].
If [param uniformly] is [code]true[/code], all region polygons and faces are weighted by their surface area (slower).
If [param uniformly] is [code]false[/code], just a random polygon and face is picked (faster).
*/
func RegionGetRandomPoint(region RID.NavigationRegion2D, navigation_layers int, uniformly bool) Vector2.XY { //gd:NavigationServer2D.region_get_random_point
	once.Do(singleton)
	return Vector2.XY(class(self).RegionGetRandomPoint(gd.RID(region), gd.Int(navigation_layers), uniformly))
}

/*
Create a new link between two positions on a map.
*/
func LinkCreate() RID.NavigationLink2D { //gd:NavigationServer2D.link_create
	once.Do(singleton)
	return RID.NavigationLink2D(class(self).LinkCreate())
}

/*
Sets the navigation map [RID] for the link.
*/
func LinkSetMap(link RID.NavigationLink2D, mapping RID.NavigationMap2D) { //gd:NavigationServer2D.link_set_map
	once.Do(singleton)
	class(self).LinkSetMap(gd.RID(link), gd.RID(mapping))
}

/*
Returns the navigation map [RID] the requested [param link] is currently assigned to.
*/
func LinkGetMap(link RID.NavigationLink2D) RID.NavigationMap2D { //gd:NavigationServer2D.link_get_map
	once.Do(singleton)
	return RID.NavigationMap2D(class(self).LinkGetMap(gd.RID(link)))
}

/*
If [param enabled] is [code]true[/code], the specified [param link] will contribute to its current navigation map.
*/
func LinkSetEnabled(link RID.NavigationLink2D, enabled bool) { //gd:NavigationServer2D.link_set_enabled
	once.Do(singleton)
	class(self).LinkSetEnabled(gd.RID(link), enabled)
}

/*
Returns [code]true[/code] if the specified [param link] is enabled.
*/
func LinkGetEnabled(link RID.NavigationLink2D) bool { //gd:NavigationServer2D.link_get_enabled
	once.Do(singleton)
	return bool(class(self).LinkGetEnabled(gd.RID(link)))
}

/*
Sets whether this [param link] can be travelled in both directions.
*/
func LinkSetBidirectional(link RID.NavigationLink2D, bidirectional bool) { //gd:NavigationServer2D.link_set_bidirectional
	once.Do(singleton)
	class(self).LinkSetBidirectional(gd.RID(link), bidirectional)
}

/*
Returns whether this [param link] can be travelled in both directions.
*/
func LinkIsBidirectional(link RID.NavigationLink2D) bool { //gd:NavigationServer2D.link_is_bidirectional
	once.Do(singleton)
	return bool(class(self).LinkIsBidirectional(gd.RID(link)))
}

/*
Set the links's navigation layers. This allows selecting links from a path request (when using [method NavigationServer2D.map_get_path]).
*/
func LinkSetNavigationLayers(link RID.NavigationLink2D, navigation_layers int) { //gd:NavigationServer2D.link_set_navigation_layers
	once.Do(singleton)
	class(self).LinkSetNavigationLayers(gd.RID(link), gd.Int(navigation_layers))
}

/*
Returns the navigation layers for this [param link].
*/
func LinkGetNavigationLayers(link RID.NavigationLink2D) int { //gd:NavigationServer2D.link_get_navigation_layers
	once.Do(singleton)
	return int(int(class(self).LinkGetNavigationLayers(gd.RID(link))))
}

/*
Sets the entry position for this [param link].
*/
func LinkSetStartPosition(link RID.NavigationLink2D, position Vector2.XY) { //gd:NavigationServer2D.link_set_start_position
	once.Do(singleton)
	class(self).LinkSetStartPosition(gd.RID(link), gd.Vector2(position))
}

/*
Returns the starting position of this [param link].
*/
func LinkGetStartPosition(link RID.NavigationLink2D) Vector2.XY { //gd:NavigationServer2D.link_get_start_position
	once.Do(singleton)
	return Vector2.XY(class(self).LinkGetStartPosition(gd.RID(link)))
}

/*
Sets the exit position for the [param link].
*/
func LinkSetEndPosition(link RID.NavigationLink2D, position Vector2.XY) { //gd:NavigationServer2D.link_set_end_position
	once.Do(singleton)
	class(self).LinkSetEndPosition(gd.RID(link), gd.Vector2(position))
}

/*
Returns the ending position of this [param link].
*/
func LinkGetEndPosition(link RID.NavigationLink2D) Vector2.XY { //gd:NavigationServer2D.link_get_end_position
	once.Do(singleton)
	return Vector2.XY(class(self).LinkGetEndPosition(gd.RID(link)))
}

/*
Sets the [param enter_cost] for this [param link].
*/
func LinkSetEnterCost(link RID.NavigationLink2D, enter_cost Float.X) { //gd:NavigationServer2D.link_set_enter_cost
	once.Do(singleton)
	class(self).LinkSetEnterCost(gd.RID(link), gd.Float(enter_cost))
}

/*
Returns the enter cost of this [param link].
*/
func LinkGetEnterCost(link RID.NavigationLink2D) Float.X { //gd:NavigationServer2D.link_get_enter_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).LinkGetEnterCost(gd.RID(link))))
}

/*
Sets the [param travel_cost] for this [param link].
*/
func LinkSetTravelCost(link RID.NavigationLink2D, travel_cost Float.X) { //gd:NavigationServer2D.link_set_travel_cost
	once.Do(singleton)
	class(self).LinkSetTravelCost(gd.RID(link), gd.Float(travel_cost))
}

/*
Returns the travel cost of this [param link].
*/
func LinkGetTravelCost(link RID.NavigationLink2D) Float.X { //gd:NavigationServer2D.link_get_travel_cost
	once.Do(singleton)
	return Float.X(Float.X(class(self).LinkGetTravelCost(gd.RID(link))))
}

/*
Set the [code]ObjectID[/code] of the object which manages this link.
*/
func LinkSetOwnerId(link RID.NavigationLink2D, owner_id int) { //gd:NavigationServer2D.link_set_owner_id
	once.Do(singleton)
	class(self).LinkSetOwnerId(gd.RID(link), gd.Int(owner_id))
}

/*
Returns the [code]ObjectID[/code] of the object which manages this link.
*/
func LinkGetOwnerId(link RID.NavigationLink2D) int { //gd:NavigationServer2D.link_get_owner_id
	once.Do(singleton)
	return int(int(class(self).LinkGetOwnerId(gd.RID(link))))
}

/*
Creates the agent.
*/
func AgentCreate() RID.NavigationAgent2D { //gd:NavigationServer2D.agent_create
	once.Do(singleton)
	return RID.NavigationAgent2D(class(self).AgentCreate())
}

/*
If [param enabled] is [code]true[/code], the specified [param agent] uses avoidance.
*/
func AgentSetAvoidanceEnabled(agent RID.NavigationAgent2D, enabled bool) { //gd:NavigationServer2D.agent_set_avoidance_enabled
	once.Do(singleton)
	class(self).AgentSetAvoidanceEnabled(gd.RID(agent), enabled)
}

/*
Return [code]true[/code] if the specified [param agent] uses avoidance.
*/
func AgentGetAvoidanceEnabled(agent RID.NavigationAgent2D) bool { //gd:NavigationServer2D.agent_get_avoidance_enabled
	once.Do(singleton)
	return bool(class(self).AgentGetAvoidanceEnabled(gd.RID(agent)))
}

/*
Puts the agent in the map.
*/
func AgentSetMap(agent RID.NavigationAgent2D, mapping RID.NavigationMap2D) { //gd:NavigationServer2D.agent_set_map
	once.Do(singleton)
	class(self).AgentSetMap(gd.RID(agent), gd.RID(mapping))
}

/*
Returns the navigation map [RID] the requested [param agent] is currently assigned to.
*/
func AgentGetMap(agent RID.NavigationAgent2D) RID.NavigationMap2D { //gd:NavigationServer2D.agent_get_map
	once.Do(singleton)
	return RID.NavigationMap2D(class(self).AgentGetMap(gd.RID(agent)))
}

/*
If [param paused] is true the specified [param agent] will not be processed, e.g. calculate avoidance velocities or receive avoidance callbacks.
*/
func AgentSetPaused(agent RID.NavigationAgent2D, paused bool) { //gd:NavigationServer2D.agent_set_paused
	once.Do(singleton)
	class(self).AgentSetPaused(gd.RID(agent), paused)
}

/*
Returns [code]true[/code] if the specified [param agent] is paused.
*/
func AgentGetPaused(agent RID.NavigationAgent2D) bool { //gd:NavigationServer2D.agent_get_paused
	once.Do(singleton)
	return bool(class(self).AgentGetPaused(gd.RID(agent)))
}

/*
Sets the maximum distance to other agents this agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
func AgentSetNeighborDistance(agent RID.NavigationAgent2D, distance Float.X) { //gd:NavigationServer2D.agent_set_neighbor_distance
	once.Do(singleton)
	class(self).AgentSetNeighborDistance(gd.RID(agent), gd.Float(distance))
}

/*
Returns the maximum distance to other agents the specified [param agent] takes into account in the navigation.
*/
func AgentGetNeighborDistance(agent RID.NavigationAgent2D) Float.X { //gd:NavigationServer2D.agent_get_neighbor_distance
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetNeighborDistance(gd.RID(agent))))
}

/*
Sets the maximum number of other agents the agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
func AgentSetMaxNeighbors(agent RID.NavigationAgent2D, count int) { //gd:NavigationServer2D.agent_set_max_neighbors
	once.Do(singleton)
	class(self).AgentSetMaxNeighbors(gd.RID(agent), gd.Int(count))
}

/*
Returns the maximum number of other agents the specified [param agent] takes into account in the navigation.
*/
func AgentGetMaxNeighbors(agent RID.NavigationAgent2D) int { //gd:NavigationServer2D.agent_get_max_neighbors
	once.Do(singleton)
	return int(int(class(self).AgentGetMaxNeighbors(gd.RID(agent))))
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to other agents. The larger this number, the sooner this agent will respond to the presence of other agents, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
func AgentSetTimeHorizonAgents(agent RID.NavigationAgent2D, time_horizon Float.X) { //gd:NavigationServer2D.agent_set_time_horizon_agents
	once.Do(singleton)
	class(self).AgentSetTimeHorizonAgents(gd.RID(agent), gd.Float(time_horizon))
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to other agents.
*/
func AgentGetTimeHorizonAgents(agent RID.NavigationAgent2D) Float.X { //gd:NavigationServer2D.agent_get_time_horizon_agents
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetTimeHorizonAgents(gd.RID(agent))))
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to static avoidance obstacles. The larger this number, the sooner this agent will respond to the presence of static avoidance obstacles, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
func AgentSetTimeHorizonObstacles(agent RID.NavigationAgent2D, time_horizon Float.X) { //gd:NavigationServer2D.agent_set_time_horizon_obstacles
	once.Do(singleton)
	class(self).AgentSetTimeHorizonObstacles(gd.RID(agent), gd.Float(time_horizon))
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to static avoidance obstacles.
*/
func AgentGetTimeHorizonObstacles(agent RID.NavigationAgent2D) Float.X { //gd:NavigationServer2D.agent_get_time_horizon_obstacles
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetTimeHorizonObstacles(gd.RID(agent))))
}

/*
Sets the radius of the agent.
*/
func AgentSetRadius(agent RID.NavigationAgent2D, radius Float.X) { //gd:NavigationServer2D.agent_set_radius
	once.Do(singleton)
	class(self).AgentSetRadius(gd.RID(agent), gd.Float(radius))
}

/*
Returns the radius of the specified [param agent].
*/
func AgentGetRadius(agent RID.NavigationAgent2D) Float.X { //gd:NavigationServer2D.agent_get_radius
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetRadius(gd.RID(agent))))
}

/*
Sets the maximum speed of the agent. Must be positive.
*/
func AgentSetMaxSpeed(agent RID.NavigationAgent2D, max_speed Float.X) { //gd:NavigationServer2D.agent_set_max_speed
	once.Do(singleton)
	class(self).AgentSetMaxSpeed(gd.RID(agent), gd.Float(max_speed))
}

/*
Returns the maximum speed of the specified [param agent].
*/
func AgentGetMaxSpeed(agent RID.NavigationAgent2D) Float.X { //gd:NavigationServer2D.agent_get_max_speed
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetMaxSpeed(gd.RID(agent))))
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity] for the specified [param agent]. When an agent is teleported to a new position far away this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
func AgentSetVelocityForced(agent RID.NavigationAgent2D, velocity Vector2.XY) { //gd:NavigationServer2D.agent_set_velocity_forced
	once.Do(singleton)
	class(self).AgentSetVelocityForced(gd.RID(agent), gd.Vector2(velocity))
}

/*
Sets [param velocity] as the new wanted velocity for the specified [param agent]. The avoidance simulation will try to fulfill this velocity if possible but will modify it to avoid collision with other agent's and obstacles. When an agent is teleported to a new position far away use [method agent_set_velocity_forced] instead to reset the internal velocity state.
*/
func AgentSetVelocity(agent RID.NavigationAgent2D, velocity Vector2.XY) { //gd:NavigationServer2D.agent_set_velocity
	once.Do(singleton)
	class(self).AgentSetVelocity(gd.RID(agent), gd.Vector2(velocity))
}

/*
Returns the velocity of the specified [param agent].
*/
func AgentGetVelocity(agent RID.NavigationAgent2D) Vector2.XY { //gd:NavigationServer2D.agent_get_velocity
	once.Do(singleton)
	return Vector2.XY(class(self).AgentGetVelocity(gd.RID(agent)))
}

/*
Sets the position of the agent in world space.
*/
func AgentSetPosition(agent RID.NavigationAgent2D, position Vector2.XY) { //gd:NavigationServer2D.agent_set_position
	once.Do(singleton)
	class(self).AgentSetPosition(gd.RID(agent), gd.Vector2(position))
}

/*
Returns the position of the specified [param agent] in world space.
*/
func AgentGetPosition(agent RID.NavigationAgent2D) Vector2.XY { //gd:NavigationServer2D.agent_get_position
	once.Do(singleton)
	return Vector2.XY(class(self).AgentGetPosition(gd.RID(agent)))
}

/*
Returns true if the map got changed the previous frame.
*/
func AgentIsMapChanged(agent RID.NavigationAgent2D) bool { //gd:NavigationServer2D.agent_is_map_changed
	once.Do(singleton)
	return bool(class(self).AgentIsMapChanged(gd.RID(agent)))
}

/*
Sets the callback [Callable] that gets called after each avoidance processing step for the [param agent]. The calculated [code]safe_velocity[/code] will be dispatched with a signal to the object just before the physics calculations.
[b]Note:[/b] Created callbacks are always processed independently of the SceneTree state as long as the agent is on a navigation map and not freed. To disable the dispatch of a callback from an agent use [method agent_set_avoidance_callback] again with an empty [Callable].
*/
func AgentSetAvoidanceCallback(agent RID.NavigationAgent2D, callback func(velocity Vector2.XY)) { //gd:NavigationServer2D.agent_set_avoidance_callback
	once.Do(singleton)
	class(self).AgentSetAvoidanceCallback(gd.RID(agent), Callable.New(callback))
}

/*
Return [code]true[/code] if the specified [param agent] has an avoidance callback.
*/
func AgentHasAvoidanceCallback(agent RID.NavigationAgent2D) bool { //gd:NavigationServer2D.agent_has_avoidance_callback
	once.Do(singleton)
	return bool(class(self).AgentHasAvoidanceCallback(gd.RID(agent)))
}

/*
Set the agent's [code]avoidance_layers[/code] bitmask.
*/
func AgentSetAvoidanceLayers(agent RID.NavigationAgent2D, layers int) { //gd:NavigationServer2D.agent_set_avoidance_layers
	once.Do(singleton)
	class(self).AgentSetAvoidanceLayers(gd.RID(agent), gd.Int(layers))
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param agent].
*/
func AgentGetAvoidanceLayers(agent RID.NavigationAgent2D) int { //gd:NavigationServer2D.agent_get_avoidance_layers
	once.Do(singleton)
	return int(int(class(self).AgentGetAvoidanceLayers(gd.RID(agent))))
}

/*
Set the agent's [code]avoidance_mask[/code] bitmask.
*/
func AgentSetAvoidanceMask(agent RID.NavigationAgent2D, mask int) { //gd:NavigationServer2D.agent_set_avoidance_mask
	once.Do(singleton)
	class(self).AgentSetAvoidanceMask(gd.RID(agent), gd.Int(mask))
}

/*
Returns the [code]avoidance_mask[/code] bitmask of the specified [param agent].
*/
func AgentGetAvoidanceMask(agent RID.NavigationAgent2D) int { //gd:NavigationServer2D.agent_get_avoidance_mask
	once.Do(singleton)
	return int(int(class(self).AgentGetAvoidanceMask(gd.RID(agent))))
}

/*
Set the agent's [code]avoidance_priority[/code] with a [param priority] between 0.0 (lowest priority) to 1.0 (highest priority).
The specified [param agent] does not adjust the velocity for other agents that would match the [code]avoidance_mask[/code] but have a lower [code]avoidance_priority[/code]. This in turn makes the other agents with lower priority adjust their velocities even more to avoid collision with this agent.
*/
func AgentSetAvoidancePriority(agent RID.NavigationAgent2D, priority Float.X) { //gd:NavigationServer2D.agent_set_avoidance_priority
	once.Do(singleton)
	class(self).AgentSetAvoidancePriority(gd.RID(agent), gd.Float(priority))
}

/*
Returns the [code]avoidance_priority[/code] of the specified [param agent].
*/
func AgentGetAvoidancePriority(agent RID.NavigationAgent2D) Float.X { //gd:NavigationServer2D.agent_get_avoidance_priority
	once.Do(singleton)
	return Float.X(Float.X(class(self).AgentGetAvoidancePriority(gd.RID(agent))))
}

/*
Creates a new navigation obstacle.
*/
func ObstacleCreate() RID.NavigationObstacle2D { //gd:NavigationServer2D.obstacle_create
	once.Do(singleton)
	return RID.NavigationObstacle2D(class(self).ObstacleCreate())
}

/*
If [param enabled] is [code]true[/code], the provided [param obstacle] affects avoidance using agents.
*/
func ObstacleSetAvoidanceEnabled(obstacle RID.NavigationObstacle2D, enabled bool) { //gd:NavigationServer2D.obstacle_set_avoidance_enabled
	once.Do(singleton)
	class(self).ObstacleSetAvoidanceEnabled(gd.RID(obstacle), enabled)
}

/*
Returns [code]true[/code] if the provided [param obstacle] has avoidance enabled.
*/
func ObstacleGetAvoidanceEnabled(obstacle RID.NavigationObstacle2D) bool { //gd:NavigationServer2D.obstacle_get_avoidance_enabled
	once.Do(singleton)
	return bool(class(self).ObstacleGetAvoidanceEnabled(gd.RID(obstacle)))
}

/*
Sets the navigation map [RID] for the obstacle.
*/
func ObstacleSetMap(obstacle RID.NavigationObstacle2D, mapping RID.NavigationMap2D) { //gd:NavigationServer2D.obstacle_set_map
	once.Do(singleton)
	class(self).ObstacleSetMap(gd.RID(obstacle), gd.RID(mapping))
}

/*
Returns the navigation map [RID] the requested [param obstacle] is currently assigned to.
*/
func ObstacleGetMap(obstacle RID.NavigationObstacle2D) RID.NavigationMap2D { //gd:NavigationServer2D.obstacle_get_map
	once.Do(singleton)
	return RID.NavigationMap2D(class(self).ObstacleGetMap(gd.RID(obstacle)))
}

/*
If [param paused] is true the specified [param obstacle] will not be processed, e.g. affect avoidance velocities.
*/
func ObstacleSetPaused(obstacle RID.NavigationObstacle2D, paused bool) { //gd:NavigationServer2D.obstacle_set_paused
	once.Do(singleton)
	class(self).ObstacleSetPaused(gd.RID(obstacle), paused)
}

/*
Returns [code]true[/code] if the specified [param obstacle] is paused.
*/
func ObstacleGetPaused(obstacle RID.NavigationObstacle2D) bool { //gd:NavigationServer2D.obstacle_get_paused
	once.Do(singleton)
	return bool(class(self).ObstacleGetPaused(gd.RID(obstacle)))
}

/*
Sets the radius of the dynamic obstacle.
*/
func ObstacleSetRadius(obstacle RID.NavigationObstacle2D, radius Float.X) { //gd:NavigationServer2D.obstacle_set_radius
	once.Do(singleton)
	class(self).ObstacleSetRadius(gd.RID(obstacle), gd.Float(radius))
}

/*
Returns the radius of the specified dynamic [param obstacle].
*/
func ObstacleGetRadius(obstacle RID.NavigationObstacle2D) Float.X { //gd:NavigationServer2D.obstacle_get_radius
	once.Do(singleton)
	return Float.X(Float.X(class(self).ObstacleGetRadius(gd.RID(obstacle))))
}

/*
Sets [param velocity] of the dynamic [param obstacle]. Allows other agents to better predict the movement of the dynamic obstacle. Only works in combination with the radius of the obstacle.
*/
func ObstacleSetVelocity(obstacle RID.NavigationObstacle2D, velocity Vector2.XY) { //gd:NavigationServer2D.obstacle_set_velocity
	once.Do(singleton)
	class(self).ObstacleSetVelocity(gd.RID(obstacle), gd.Vector2(velocity))
}

/*
Returns the velocity of the specified dynamic [param obstacle].
*/
func ObstacleGetVelocity(obstacle RID.NavigationObstacle2D) Vector2.XY { //gd:NavigationServer2D.obstacle_get_velocity
	once.Do(singleton)
	return Vector2.XY(class(self).ObstacleGetVelocity(gd.RID(obstacle)))
}

/*
Sets the position of the obstacle in world space.
*/
func ObstacleSetPosition(obstacle RID.NavigationObstacle2D, position Vector2.XY) { //gd:NavigationServer2D.obstacle_set_position
	once.Do(singleton)
	class(self).ObstacleSetPosition(gd.RID(obstacle), gd.Vector2(position))
}

/*
Returns the position of the specified [param obstacle] in world space.
*/
func ObstacleGetPosition(obstacle RID.NavigationObstacle2D) Vector2.XY { //gd:NavigationServer2D.obstacle_get_position
	once.Do(singleton)
	return Vector2.XY(class(self).ObstacleGetPosition(gd.RID(obstacle)))
}

/*
Sets the outline vertices for the obstacle. If the vertices are winded in clockwise order agents will be pushed in by the obstacle, else they will be pushed out.
*/
func ObstacleSetVertices(obstacle RID.NavigationObstacle2D, vertices []Vector2.XY) { //gd:NavigationServer2D.obstacle_set_vertices
	once.Do(singleton)
	class(self).ObstacleSetVertices(gd.RID(obstacle), Packed.New(vertices...))
}

/*
Returns the outline vertices for the specified [param obstacle].
*/
func ObstacleGetVertices(obstacle RID.NavigationObstacle2D) []Vector2.XY { //gd:NavigationServer2D.obstacle_get_vertices
	once.Do(singleton)
	return []Vector2.XY(slices.Collect(class(self).ObstacleGetVertices(gd.RID(obstacle)).Values()))
}

/*
Set the obstacles's [code]avoidance_layers[/code] bitmask.
*/
func ObstacleSetAvoidanceLayers(obstacle RID.NavigationObstacle2D, layers int) { //gd:NavigationServer2D.obstacle_set_avoidance_layers
	once.Do(singleton)
	class(self).ObstacleSetAvoidanceLayers(gd.RID(obstacle), gd.Int(layers))
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param obstacle].
*/
func ObstacleGetAvoidanceLayers(obstacle RID.NavigationObstacle2D) int { //gd:NavigationServer2D.obstacle_get_avoidance_layers
	once.Do(singleton)
	return int(int(class(self).ObstacleGetAvoidanceLayers(gd.RID(obstacle))))
}

/*
Parses the [SceneTree] for source geometry according to the properties of [param navigation_polygon]. Updates the provided [param source_geometry_data] resource with the resulting data. The resource can then be used to bake a navigation mesh with [method bake_from_source_geometry_data]. After the process is finished the optional [param callback] will be called.
[b]Note:[/b] This function needs to run on the main thread or with a deferred call as the SceneTree is not thread-safe.
[b]Performance:[/b] While convenient, reading data arrays from [Mesh] resources can affect the frame rate negatively. The data needs to be received from the GPU, stalling the [RenderingServer] in the process. For performance prefer the use of e.g. collision shapes or creating the data arrays entirely in code.
*/
func ParseSourceGeometryData(navigation_polygon [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D, root_node [1]gdclass.Node) { //gd:NavigationServer2D.parse_source_geometry_data
	once.Do(singleton)
	class(self).ParseSourceGeometryData(navigation_polygon, source_geometry_data, root_node, Callable.New(Callable.Nil))
}

/*
Bakes the provided [param navigation_polygon] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
func BakeFromSourceGeometryData(navigation_polygon [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D) { //gd:NavigationServer2D.bake_from_source_geometry_data
	once.Do(singleton)
	class(self).BakeFromSourceGeometryData(navigation_polygon, source_geometry_data, Callable.New(Callable.Nil))
}

/*
Bakes the provided [param navigation_polygon] with the data from the provided [param source_geometry_data] as an async task running on a background thread. After the process is finished the optional [param callback] will be called.
*/
func BakeFromSourceGeometryDataAsync(navigation_polygon [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D) { //gd:NavigationServer2D.bake_from_source_geometry_data_async
	once.Do(singleton)
	class(self).BakeFromSourceGeometryDataAsync(navigation_polygon, source_geometry_data, Callable.New(Callable.Nil))
}

/*
Returns [code]true[/code] when the provided navigation polygon is being baked on a background thread.
*/
func IsBakingNavigationPolygon(navigation_polygon [1]gdclass.NavigationPolygon) bool { //gd:NavigationServer2D.is_baking_navigation_polygon
	once.Do(singleton)
	return bool(class(self).IsBakingNavigationPolygon(navigation_polygon))
}

/*
Creates a new source geometry parser. If a [Callable] is set for the parser with [method source_geometry_parser_set_callback] the callback will be called for every single node that gets parsed whenever [method parse_source_geometry_data] is used.
*/
func SourceGeometryParserCreate() RID.NavigationSourceGeometryParser2D { //gd:NavigationServer2D.source_geometry_parser_create
	once.Do(singleton)
	return RID.NavigationSourceGeometryParser2D(class(self).SourceGeometryParserCreate())
}

/*
Sets the [param callback] [Callable] for the specific source geometry [param parser]. The [Callable] will receive a call with the following parameters:
- [code]navigation_mesh[/code] - The [NavigationPolygon] reference used to define the parse settings. Do NOT edit or add directly to the navigation mesh.
- [code]source_geometry_data[/code] - The [NavigationMeshSourceGeometryData2D] reference. Add custom source geometry for navigation mesh baking to this object.
- [code]node[/code] - The [Node] that is parsed.
*/
func SourceGeometryParserSetCallback(parser RID.NavigationSourceGeometryParser2D, callback func(navigation_mesh [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D, node [1]gdclass.Node)) { //gd:NavigationServer2D.source_geometry_parser_set_callback
	once.Do(singleton)
	class(self).SourceGeometryParserSetCallback(gd.RID(parser), Callable.New(callback))
}

/*
Returns a simplified version of [param path] with less critical path points removed. The simplification amount is in worlds units and controlled by [param epsilon]. The simplification uses a variant of Ramer-Douglas-Peucker algorithm for curve point decimation.
Path simplification can be helpful to mitigate various path following issues that can arise with certain agent types and script behaviors. E.g. "steering" agents or avoidance in "open fields".
*/
func SimplifyPath(path []Vector2.XY, epsilon Float.X) []Vector2.XY { //gd:NavigationServer2D.simplify_path
	once.Do(singleton)
	return []Vector2.XY(slices.Collect(class(self).SimplifyPath(Packed.New(path...), gd.Float(epsilon)).Values()))
}

/*
Destroys the given RID.
*/
func FreeRid(rid RID.Any) { //gd:NavigationServer2D.free_rid
	once.Do(singleton)
	class(self).FreeRid(gd.RID(rid))
}

/*
If [code]true[/code] enables debug mode on the NavigationServer.
*/
func SetDebugEnabled(enabled bool) { //gd:NavigationServer2D.set_debug_enabled
	once.Do(singleton)
	class(self).SetDebugEnabled(enabled)
}

/*
Returns [code]true[/code] when the NavigationServer has debug enabled.
*/
func GetDebugEnabled() bool { //gd:NavigationServer2D.get_debug_enabled
	once.Do(singleton)
	return bool(class(self).GetDebugEnabled())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.NavigationServer2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns all created navigation map [RID]s on the NavigationServer. This returns both 2D and 3D created navigation maps as there is technically no distinction between them.
*/
//go:nosplit
func (self class) GetMaps() Array.Contains[gd.RID] { //gd:NavigationServer2D.get_maps
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_get_maps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Create a new map.
*/
//go:nosplit
func (self class) MapCreate() gd.RID { //gd:NavigationServer2D.map_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map active.
*/
//go:nosplit
func (self class) MapSetActive(mapping gd.RID, active bool) { //gd:NavigationServer2D.map_set_active
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_set_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns true if the map is active.
*/
//go:nosplit
func (self class) MapIsActive(mapping gd.RID) bool { //gd:NavigationServer2D.map_is_active
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map cell size used to rasterize the navigation mesh vertices. Must match with the cell size of the used navigation meshes.
*/
//go:nosplit
func (self class) MapSetCellSize(mapping gd.RID, cell_size gd.Float) { //gd:NavigationServer2D.map_set_cell_size
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, cell_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the map cell size used to rasterize the navigation mesh vertices.
*/
//go:nosplit
func (self class) MapGetCellSize(mapping gd.RID) gd.Float { //gd:NavigationServer2D.map_get_cell_size
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the navigation [param map] edge connection use. If [param enabled] is [code]true[/code], the navigation map allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) MapSetUseEdgeConnections(mapping gd.RID, enabled bool) { //gd:NavigationServer2D.map_set_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the navigation [param map] allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) MapGetUseEdgeConnections(mapping gd.RID) bool { //gd:NavigationServer2D.map_get_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the map edge connection margin used to weld the compatible region edges.
*/
//go:nosplit
func (self class) MapSetEdgeConnectionMargin(mapping gd.RID, margin gd.Float) { //gd:NavigationServer2D.map_set_edge_connection_margin
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_set_edge_connection_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the edge connection margin of the map. The edge connection margin is a distance used to connect two regions.
*/
//go:nosplit
func (self class) MapGetEdgeConnectionMargin(mapping gd.RID) gd.Float { //gd:NavigationServer2D.map_get_edge_connection_margin
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_edge_connection_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the map's link connection radius used to connect links to navigation polygons.
*/
//go:nosplit
func (self class) MapSetLinkConnectionRadius(mapping gd.RID, radius gd.Float) { //gd:NavigationServer2D.map_set_link_connection_radius
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_set_link_connection_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the link connection radius of the map. This distance is the maximum range any link will search for navigation mesh polygons to connect to.
*/
//go:nosplit
func (self class) MapGetLinkConnectionRadius(mapping gd.RID) gd.Float { //gd:NavigationServer2D.map_get_link_connection_radius
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_link_connection_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the navigation path to reach the destination from the origin. [param navigation_layers] is a bitmask of all region navigation layers that are allowed to be in the path.
*/
//go:nosplit
func (self class) MapGetPath(mapping gd.RID, origin gd.Vector2, destination gd.Vector2, optimize bool, navigation_layers gd.Int) Packed.Array[Vector2.XY] { //gd:NavigationServer2D.map_get_path
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, origin)
	callframe.Arg(frame, destination)
	callframe.Arg(frame, optimize)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the point closest to the provided [param to_point] on the navigation mesh surface.
*/
//go:nosplit
func (self class) MapGetClosestPoint(mapping gd.RID, to_point gd.Vector2) gd.Vector2 { //gd:NavigationServer2D.map_get_closest_point
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the owner region RID for the point returned by [method map_get_closest_point].
*/
//go:nosplit
func (self class) MapGetClosestPointOwner(mapping gd.RID, to_point gd.Vector2) gd.RID { //gd:NavigationServer2D.map_get_closest_point_owner
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_closest_point_owner, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns all navigation link [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetLinks(mapping gd.RID) Array.Contains[gd.RID] { //gd:NavigationServer2D.map_get_links
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_links, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns all navigation regions [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetRegions(mapping gd.RID) Array.Contains[gd.RID] { //gd:NavigationServer2D.map_get_regions
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_regions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns all navigation agents [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetAgents(mapping gd.RID) Array.Contains[gd.RID] { //gd:NavigationServer2D.map_get_agents
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns all navigation obstacle [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetObstacles(mapping gd.RID) Array.Contains[gd.RID] { //gd:NavigationServer2D.map_get_obstacles
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
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
func (self class) MapForceUpdate(mapping gd.RID) { //gd:NavigationServer2D.map_force_update
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_force_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current iteration id of the navigation map. Every time the navigation map changes and synchronizes the iteration id increases. An iteration id of 0 means the navigation map has never synchronized.
[b]Note:[/b] The iteration id will wrap back to 1 after reaching its range limit.
*/
//go:nosplit
func (self class) MapGetIterationId(mapping gd.RID) gd.Int { //gd:NavigationServer2D.map_get_iteration_id
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_iteration_id, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) MapGetRandomPoint(mapping gd.RID, navigation_layers gd.Int, uniformly bool) gd.Vector2 { //gd:NavigationServer2D.map_get_random_point
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, navigation_layers)
	callframe.Arg(frame, uniformly)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_map_get_random_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Queries a path in a given navigation map. Start and target position and other parameters are defined through [NavigationPathQueryParameters2D]. Updates the provided [NavigationPathQueryResult2D] result object with the path among other results requested by the query.
*/
//go:nosplit
func (self class) QueryPath(parameters [1]gdclass.NavigationPathQueryParameters2D, result [1]gdclass.NavigationPathQueryResult2D) { //gd:NavigationServer2D.query_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parameters[0])[0])
	callframe.Arg(frame, pointers.Get(result[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_query_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a new region.
*/
//go:nosplit
func (self class) RegionCreate() gd.RID { //gd:NavigationServer2D.region_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code] the specified [param region] will contribute to its current navigation map.
*/
//go:nosplit
func (self class) RegionSetEnabled(region gd.RID, enabled bool) { //gd:NavigationServer2D.region_set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param region] is enabled.
*/
//go:nosplit
func (self class) RegionGetEnabled(region gd.RID) bool { //gd:NavigationServer2D.region_get_enabled
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the navigation [param region] will use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) RegionSetUseEdgeConnections(region gd.RID, enabled bool) { //gd:NavigationServer2D.region_set_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the navigation [param region] is set to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) RegionGetUseEdgeConnections(region gd.RID) bool { //gd:NavigationServer2D.region_get_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param enter_cost] for this [param region].
*/
//go:nosplit
func (self class) RegionSetEnterCost(region gd.RID, enter_cost gd.Float) { //gd:NavigationServer2D.region_set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the enter cost of this [param region].
*/
//go:nosplit
func (self class) RegionGetEnterCost(region gd.RID) gd.Float { //gd:NavigationServer2D.region_get_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param travel_cost] for this [param region].
*/
//go:nosplit
func (self class) RegionSetTravelCost(region gd.RID, travel_cost gd.Float) { //gd:NavigationServer2D.region_set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the travel cost of this [param region].
*/
//go:nosplit
func (self class) RegionGetTravelCost(region gd.RID) gd.Float { //gd:NavigationServer2D.region_get_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the [code]ObjectID[/code] of the object which manages this region.
*/
//go:nosplit
func (self class) RegionSetOwnerId(region gd.RID, owner_id gd.Int) { //gd:NavigationServer2D.region_set_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]ObjectID[/code] of the object which manages this region.
*/
//go:nosplit
func (self class) RegionGetOwnerId(region gd.RID) gd.Int { //gd:NavigationServer2D.region_get_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) RegionOwnsPoint(region gd.RID, point gd.Vector2) bool { //gd:NavigationServer2D.region_owns_point
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_owns_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the map for the region.
*/
//go:nosplit
func (self class) RegionSetMap(region gd.RID, mapping gd.RID) { //gd:NavigationServer2D.region_set_map
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param region] is currently assigned to.
*/
//go:nosplit
func (self class) RegionGetMap(region gd.RID) gd.RID { //gd:NavigationServer2D.region_get_map
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the region's navigation layers. This allows selecting regions from a path request (when using [method NavigationServer2D.map_get_path]).
*/
//go:nosplit
func (self class) RegionSetNavigationLayers(region gd.RID, navigation_layers gd.Int) { //gd:NavigationServer2D.region_set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the region's navigation layers.
*/
//go:nosplit
func (self class) RegionGetNavigationLayers(region gd.RID) gd.Int { //gd:NavigationServer2D.region_get_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the global transformation for the region.
*/
//go:nosplit
func (self class) RegionSetTransform(region gd.RID, transform gd.Transform2D) { //gd:NavigationServer2D.region_set_transform
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the global transformation of this [param region].
*/
//go:nosplit
func (self class) RegionGetTransform(region gd.RID) gd.Transform2D { //gd:NavigationServer2D.region_get_transform
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param navigation_polygon] for the region.
*/
//go:nosplit
func (self class) RegionSetNavigationPolygon(region gd.RID, navigation_polygon [1]gdclass.NavigationPolygon) { //gd:NavigationServer2D.region_set_navigation_polygon
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_set_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns how many connections this [param region] has with other regions in the map.
*/
//go:nosplit
func (self class) RegionGetConnectionsCount(region gd.RID) gd.Int { //gd:NavigationServer2D.region_get_connections_count
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_connections_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the starting point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
//go:nosplit
func (self class) RegionGetConnectionPathwayStart(region gd.RID, connection gd.Int) gd.Vector2 { //gd:NavigationServer2D.region_get_connection_pathway_start
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, connection)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_connection_pathway_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ending point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
//go:nosplit
func (self class) RegionGetConnectionPathwayEnd(region gd.RID, connection gd.Int) gd.Vector2 { //gd:NavigationServer2D.region_get_connection_pathway_end
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, connection)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_connection_pathway_end, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) RegionGetRandomPoint(region gd.RID, navigation_layers gd.Int, uniformly bool) gd.Vector2 { //gd:NavigationServer2D.region_get_random_point
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, navigation_layers)
	callframe.Arg(frame, uniformly)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_region_get_random_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Create a new link between two positions on a map.
*/
//go:nosplit
func (self class) LinkCreate() gd.RID { //gd:NavigationServer2D.link_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the navigation map [RID] for the link.
*/
//go:nosplit
func (self class) LinkSetMap(link gd.RID, mapping gd.RID) { //gd:NavigationServer2D.link_set_map
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param link] is currently assigned to.
*/
//go:nosplit
func (self class) LinkGetMap(link gd.RID) gd.RID { //gd:NavigationServer2D.link_get_map
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the specified [param link] will contribute to its current navigation map.
*/
//go:nosplit
func (self class) LinkSetEnabled(link gd.RID, enabled bool) { //gd:NavigationServer2D.link_set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param link] is enabled.
*/
//go:nosplit
func (self class) LinkGetEnabled(link gd.RID) bool { //gd:NavigationServer2D.link_get_enabled
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether this [param link] can be travelled in both directions.
*/
//go:nosplit
func (self class) LinkSetBidirectional(link gd.RID, bidirectional bool) { //gd:NavigationServer2D.link_set_bidirectional
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, bidirectional)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether this [param link] can be travelled in both directions.
*/
//go:nosplit
func (self class) LinkIsBidirectional(link gd.RID) bool { //gd:NavigationServer2D.link_is_bidirectional
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the links's navigation layers. This allows selecting links from a path request (when using [method NavigationServer2D.map_get_path]).
*/
//go:nosplit
func (self class) LinkSetNavigationLayers(link gd.RID, navigation_layers gd.Int) { //gd:NavigationServer2D.link_set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation layers for this [param link].
*/
//go:nosplit
func (self class) LinkGetNavigationLayers(link gd.RID) gd.Int { //gd:NavigationServer2D.link_get_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the entry position for this [param link].
*/
//go:nosplit
func (self class) LinkSetStartPosition(link gd.RID, position gd.Vector2) { //gd:NavigationServer2D.link_set_start_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the starting position of this [param link].
*/
//go:nosplit
func (self class) LinkGetStartPosition(link gd.RID) gd.Vector2 { //gd:NavigationServer2D.link_get_start_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the exit position for the [param link].
*/
//go:nosplit
func (self class) LinkSetEndPosition(link gd.RID, position gd.Vector2) { //gd:NavigationServer2D.link_set_end_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the ending position of this [param link].
*/
//go:nosplit
func (self class) LinkGetEndPosition(link gd.RID) gd.Vector2 { //gd:NavigationServer2D.link_get_end_position
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param enter_cost] for this [param link].
*/
//go:nosplit
func (self class) LinkSetEnterCost(link gd.RID, enter_cost gd.Float) { //gd:NavigationServer2D.link_set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the enter cost of this [param link].
*/
//go:nosplit
func (self class) LinkGetEnterCost(link gd.RID) gd.Float { //gd:NavigationServer2D.link_get_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param travel_cost] for this [param link].
*/
//go:nosplit
func (self class) LinkSetTravelCost(link gd.RID, travel_cost gd.Float) { //gd:NavigationServer2D.link_set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the travel cost of this [param link].
*/
//go:nosplit
func (self class) LinkGetTravelCost(link gd.RID) gd.Float { //gd:NavigationServer2D.link_get_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the [code]ObjectID[/code] of the object which manages this link.
*/
//go:nosplit
func (self class) LinkSetOwnerId(link gd.RID, owner_id gd.Int) { //gd:NavigationServer2D.link_set_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, owner_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_set_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]ObjectID[/code] of the object which manages this link.
*/
//go:nosplit
func (self class) LinkGetOwnerId(link gd.RID) gd.Int { //gd:NavigationServer2D.link_get_owner_id
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_link_get_owner_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates the agent.
*/
//go:nosplit
func (self class) AgentCreate() gd.RID { //gd:NavigationServer2D.agent_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the specified [param agent] uses avoidance.
*/
//go:nosplit
func (self class) AgentSetAvoidanceEnabled(agent gd.RID, enabled bool) { //gd:NavigationServer2D.agent_set_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return [code]true[/code] if the specified [param agent] uses avoidance.
*/
//go:nosplit
func (self class) AgentGetAvoidanceEnabled(agent gd.RID) bool { //gd:NavigationServer2D.agent_get_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Puts the agent in the map.
*/
//go:nosplit
func (self class) AgentSetMap(agent gd.RID, mapping gd.RID) { //gd:NavigationServer2D.agent_set_map
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param agent] is currently assigned to.
*/
//go:nosplit
func (self class) AgentGetMap(agent gd.RID) gd.RID { //gd:NavigationServer2D.agent_get_map
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param paused] is true the specified [param agent] will not be processed, e.g. calculate avoidance velocities or receive avoidance callbacks.
*/
//go:nosplit
func (self class) AgentSetPaused(agent gd.RID, paused bool) { //gd:NavigationServer2D.agent_set_paused
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, paused)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param agent] is paused.
*/
//go:nosplit
func (self class) AgentGetPaused(agent gd.RID) bool { //gd:NavigationServer2D.agent_get_paused
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum distance to other agents this agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
//go:nosplit
func (self class) AgentSetNeighborDistance(agent gd.RID, distance gd.Float) { //gd:NavigationServer2D.agent_set_neighbor_distance
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum distance to other agents the specified [param agent] takes into account in the navigation.
*/
//go:nosplit
func (self class) AgentGetNeighborDistance(agent gd.RID) gd.Float { //gd:NavigationServer2D.agent_get_neighbor_distance
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum number of other agents the agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
//go:nosplit
func (self class) AgentSetMaxNeighbors(agent gd.RID, count gd.Int) { //gd:NavigationServer2D.agent_set_max_neighbors
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum number of other agents the specified [param agent] takes into account in the navigation.
*/
//go:nosplit
func (self class) AgentGetMaxNeighbors(agent gd.RID) gd.Int { //gd:NavigationServer2D.agent_get_max_neighbors
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to other agents. The larger this number, the sooner this agent will respond to the presence of other agents, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
//go:nosplit
func (self class) AgentSetTimeHorizonAgents(agent gd.RID, time_horizon gd.Float) { //gd:NavigationServer2D.agent_set_time_horizon_agents
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, time_horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to other agents.
*/
//go:nosplit
func (self class) AgentGetTimeHorizonAgents(agent gd.RID) gd.Float { //gd:NavigationServer2D.agent_get_time_horizon_agents
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to static avoidance obstacles. The larger this number, the sooner this agent will respond to the presence of static avoidance obstacles, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
//go:nosplit
func (self class) AgentSetTimeHorizonObstacles(agent gd.RID, time_horizon gd.Float) { //gd:NavigationServer2D.agent_set_time_horizon_obstacles
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, time_horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to static avoidance obstacles.
*/
//go:nosplit
func (self class) AgentGetTimeHorizonObstacles(agent gd.RID) gd.Float { //gd:NavigationServer2D.agent_get_time_horizon_obstacles
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the radius of the agent.
*/
//go:nosplit
func (self class) AgentSetRadius(agent gd.RID, radius gd.Float) { //gd:NavigationServer2D.agent_set_radius
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the radius of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetRadius(agent gd.RID) gd.Float { //gd:NavigationServer2D.agent_get_radius
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum speed of the agent. Must be positive.
*/
//go:nosplit
func (self class) AgentSetMaxSpeed(agent gd.RID, max_speed gd.Float) { //gd:NavigationServer2D.agent_set_max_speed
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, max_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_max_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum speed of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetMaxSpeed(agent gd.RID) gd.Float { //gd:NavigationServer2D.agent_get_max_speed
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_max_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity] for the specified [param agent]. When an agent is teleported to a new position far away this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
//go:nosplit
func (self class) AgentSetVelocityForced(agent gd.RID, velocity gd.Vector2) { //gd:NavigationServer2D.agent_set_velocity_forced
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_velocity_forced, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets [param velocity] as the new wanted velocity for the specified [param agent]. The avoidance simulation will try to fulfill this velocity if possible but will modify it to avoid collision with other agent's and obstacles. When an agent is teleported to a new position far away use [method agent_set_velocity_forced] instead to reset the internal velocity state.
*/
//go:nosplit
func (self class) AgentSetVelocity(agent gd.RID, velocity gd.Vector2) { //gd:NavigationServer2D.agent_set_velocity
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the velocity of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetVelocity(agent gd.RID) gd.Vector2 { //gd:NavigationServer2D.agent_get_velocity
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the agent in world space.
*/
//go:nosplit
func (self class) AgentSetPosition(agent gd.RID, position gd.Vector2) { //gd:NavigationServer2D.agent_set_position
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the specified [param agent] in world space.
*/
//go:nosplit
func (self class) AgentGetPosition(agent gd.RID) gd.Vector2 { //gd:NavigationServer2D.agent_get_position
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns true if the map got changed the previous frame.
*/
//go:nosplit
func (self class) AgentIsMapChanged(agent gd.RID) bool { //gd:NavigationServer2D.agent_is_map_changed
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_is_map_changed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the callback [Callable] that gets called after each avoidance processing step for the [param agent]. The calculated [code]safe_velocity[/code] will be dispatched with a signal to the object just before the physics calculations.
[b]Note:[/b] Created callbacks are always processed independently of the SceneTree state as long as the agent is on a navigation map and not freed. To disable the dispatch of a callback from an agent use [method agent_set_avoidance_callback] again with an empty [Callable].
*/
//go:nosplit
func (self class) AgentSetAvoidanceCallback(agent gd.RID, callback Callable.Function) { //gd:NavigationServer2D.agent_set_avoidance_callback
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_avoidance_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Return [code]true[/code] if the specified [param agent] has an avoidance callback.
*/
//go:nosplit
func (self class) AgentHasAvoidanceCallback(agent gd.RID) bool { //gd:NavigationServer2D.agent_has_avoidance_callback
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_has_avoidance_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the agent's [code]avoidance_layers[/code] bitmask.
*/
//go:nosplit
func (self class) AgentSetAvoidanceLayers(agent gd.RID, layers gd.Int) { //gd:NavigationServer2D.agent_set_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidanceLayers(agent gd.RID) gd.Int { //gd:NavigationServer2D.agent_get_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the agent's [code]avoidance_mask[/code] bitmask.
*/
//go:nosplit
func (self class) AgentSetAvoidanceMask(agent gd.RID, mask gd.Int) { //gd:NavigationServer2D.agent_set_avoidance_mask
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_mask[/code] bitmask of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidanceMask(agent gd.RID) gd.Int { //gd:NavigationServer2D.agent_get_avoidance_mask
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the agent's [code]avoidance_priority[/code] with a [param priority] between 0.0 (lowest priority) to 1.0 (highest priority).
The specified [param agent] does not adjust the velocity for other agents that would match the [code]avoidance_mask[/code] but have a lower [code]avoidance_priority[/code]. This in turn makes the other agents with lower priority adjust their velocities even more to avoid collision with this agent.
*/
//go:nosplit
func (self class) AgentSetAvoidancePriority(agent gd.RID, priority gd.Float) { //gd:NavigationServer2D.agent_set_avoidance_priority
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_set_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_priority[/code] of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidancePriority(agent gd.RID) gd.Float { //gd:NavigationServer2D.agent_get_avoidance_priority
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_agent_get_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new navigation obstacle.
*/
//go:nosplit
func (self class) ObstacleCreate() gd.RID { //gd:NavigationServer2D.obstacle_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enabled] is [code]true[/code], the provided [param obstacle] affects avoidance using agents.
*/
//go:nosplit
func (self class) ObstacleSetAvoidanceEnabled(obstacle gd.RID, enabled bool) { //gd:NavigationServer2D.obstacle_set_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the provided [param obstacle] has avoidance enabled.
*/
//go:nosplit
func (self class) ObstacleGetAvoidanceEnabled(obstacle gd.RID) bool { //gd:NavigationServer2D.obstacle_get_avoidance_enabled
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the navigation map [RID] for the obstacle.
*/
//go:nosplit
func (self class) ObstacleSetMap(obstacle gd.RID, mapping gd.RID) { //gd:NavigationServer2D.obstacle_set_map
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation map [RID] the requested [param obstacle] is currently assigned to.
*/
//go:nosplit
func (self class) ObstacleGetMap(obstacle gd.RID) gd.RID { //gd:NavigationServer2D.obstacle_get_map
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param paused] is true the specified [param obstacle] will not be processed, e.g. affect avoidance velocities.
*/
//go:nosplit
func (self class) ObstacleSetPaused(obstacle gd.RID, paused bool) { //gd:NavigationServer2D.obstacle_set_paused
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, paused)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified [param obstacle] is paused.
*/
//go:nosplit
func (self class) ObstacleGetPaused(obstacle gd.RID) bool { //gd:NavigationServer2D.obstacle_get_paused
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the radius of the dynamic obstacle.
*/
//go:nosplit
func (self class) ObstacleSetRadius(obstacle gd.RID, radius gd.Float) { //gd:NavigationServer2D.obstacle_set_radius
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the radius of the specified dynamic [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetRadius(obstacle gd.RID) gd.Float { //gd:NavigationServer2D.obstacle_get_radius
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets [param velocity] of the dynamic [param obstacle]. Allows other agents to better predict the movement of the dynamic obstacle. Only works in combination with the radius of the obstacle.
*/
//go:nosplit
func (self class) ObstacleSetVelocity(obstacle gd.RID, velocity gd.Vector2) { //gd:NavigationServer2D.obstacle_set_velocity
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the velocity of the specified dynamic [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetVelocity(obstacle gd.RID) gd.Vector2 { //gd:NavigationServer2D.obstacle_get_velocity
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the obstacle in world space.
*/
//go:nosplit
func (self class) ObstacleSetPosition(obstacle gd.RID, position gd.Vector2) { //gd:NavigationServer2D.obstacle_set_position
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the specified [param obstacle] in world space.
*/
//go:nosplit
func (self class) ObstacleGetPosition(obstacle gd.RID) gd.Vector2 { //gd:NavigationServer2D.obstacle_get_position
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the outline vertices for the obstacle. If the vertices are winded in clockwise order agents will be pushed in by the obstacle, else they will be pushed out.
*/
//go:nosplit
func (self class) ObstacleSetVertices(obstacle gd.RID, vertices Packed.Array[Vector2.XY]) { //gd:NavigationServer2D.obstacle_set_vertices
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](vertices)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the outline vertices for the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetVertices(obstacle gd.RID) Packed.Array[Vector2.XY] { //gd:NavigationServer2D.obstacle_get_vertices
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_vertices, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Set the obstacles's [code]avoidance_layers[/code] bitmask.
*/
//go:nosplit
func (self class) ObstacleSetAvoidanceLayers(obstacle gd.RID, layers gd.Int) { //gd:NavigationServer2D.obstacle_set_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetAvoidanceLayers(obstacle gd.RID) gd.Int { //gd:NavigationServer2D.obstacle_get_avoidance_layers
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_obstacle_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Parses the [SceneTree] for source geometry according to the properties of [param navigation_polygon]. Updates the provided [param source_geometry_data] resource with the resulting data. The resource can then be used to bake a navigation mesh with [method bake_from_source_geometry_data]. After the process is finished the optional [param callback] will be called.
[b]Note:[/b] This function needs to run on the main thread or with a deferred call as the SceneTree is not thread-safe.
[b]Performance:[/b] While convenient, reading data arrays from [Mesh] resources can affect the frame rate negatively. The data needs to be received from the GPU, stalling the [RenderingServer] in the process. For performance prefer the use of e.g. collision shapes or creating the data arrays entirely in code.
*/
//go:nosplit
func (self class) ParseSourceGeometryData(navigation_polygon [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D, root_node [1]gdclass.Node, callback Callable.Function) { //gd:NavigationServer2D.parse_source_geometry_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(root_node[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_parse_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes the provided [param navigation_polygon] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryData(navigation_polygon [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D, callback Callable.Function) { //gd:NavigationServer2D.bake_from_source_geometry_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_bake_from_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes the provided [param navigation_polygon] with the data from the provided [param source_geometry_data] as an async task running on a background thread. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryDataAsync(navigation_polygon [1]gdclass.NavigationPolygon, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData2D, callback Callable.Function) { //gd:NavigationServer2D.bake_from_source_geometry_data_async
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_bake_from_source_geometry_data_async, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] when the provided navigation polygon is being baked on a background thread.
*/
//go:nosplit
func (self class) IsBakingNavigationPolygon(navigation_polygon [1]gdclass.NavigationPolygon) bool { //gd:NavigationServer2D.is_baking_navigation_polygon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_is_baking_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new source geometry parser. If a [Callable] is set for the parser with [method source_geometry_parser_set_callback] the callback will be called for every single node that gets parsed whenever [method parse_source_geometry_data] is used.
*/
//go:nosplit
func (self class) SourceGeometryParserCreate() gd.RID { //gd:NavigationServer2D.source_geometry_parser_create
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_source_geometry_parser_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param callback] [Callable] for the specific source geometry [param parser]. The [Callable] will receive a call with the following parameters:
- [code]navigation_mesh[/code] - The [NavigationPolygon] reference used to define the parse settings. Do NOT edit or add directly to the navigation mesh.
- [code]source_geometry_data[/code] - The [NavigationMeshSourceGeometryData2D] reference. Add custom source geometry for navigation mesh baking to this object.
- [code]node[/code] - The [Node] that is parsed.
*/
//go:nosplit
func (self class) SourceGeometryParserSetCallback(parser gd.RID, callback Callable.Function) { //gd:NavigationServer2D.source_geometry_parser_set_callback
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_source_geometry_parser_set_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a simplified version of [param path] with less critical path points removed. The simplification amount is in worlds units and controlled by [param epsilon]. The simplification uses a variant of Ramer-Douglas-Peucker algorithm for curve point decimation.
Path simplification can be helpful to mitigate various path following issues that can arise with certain agent types and script behaviors. E.g. "steering" agents or avoidance in "open fields".
*/
//go:nosplit
func (self class) SimplifyPath(path Packed.Array[Vector2.XY], epsilon gd.Float) Packed.Array[Vector2.XY] { //gd:NavigationServer2D.simplify_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](path)))
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_simplify_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Destroys the given RID.
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID) { //gd:NavigationServer2D.free_rid
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [code]true[/code] enables debug mode on the NavigationServer.
*/
//go:nosplit
func (self class) SetDebugEnabled(enabled bool) { //gd:NavigationServer2D.set_debug_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_set_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] when the NavigationServer has debug enabled.
*/
//go:nosplit
func (self class) GetDebugEnabled() bool { //gd:NavigationServer2D.get_debug_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationServer2D.Bind_get_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("NavigationServer2D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationServer2D{*(*gdclass.NavigationServer2D)(unsafe.Pointer(&ptr))}
	})
}
