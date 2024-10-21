package NavigationServer3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Simple [1]classdb.NavigationServer3D
func (self Simple) GetMaps() gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).GetMaps(gc))
}
func (self Simple) MapCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MapCreate())
}
func (self Simple) MapSetActive(mapping gd.RID, active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetActive(mapping, active)
}
func (self Simple) MapIsActive(mapping gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).MapIsActive(mapping))
}
func (self Simple) MapSetUp(mapping gd.RID, up gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetUp(mapping, up)
}
func (self Simple) MapGetUp(mapping gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).MapGetUp(mapping))
}
func (self Simple) MapSetCellSize(mapping gd.RID, cell_size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetCellSize(mapping, gd.Float(cell_size))
}
func (self Simple) MapGetCellSize(mapping gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).MapGetCellSize(mapping)))
}
func (self Simple) MapSetCellHeight(mapping gd.RID, cell_height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetCellHeight(mapping, gd.Float(cell_height))
}
func (self Simple) MapGetCellHeight(mapping gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).MapGetCellHeight(mapping)))
}
func (self Simple) MapSetMergeRasterizerCellScale(mapping gd.RID, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetMergeRasterizerCellScale(mapping, gd.Float(scale))
}
func (self Simple) MapGetMergeRasterizerCellScale(mapping gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).MapGetMergeRasterizerCellScale(mapping)))
}
func (self Simple) MapSetUseEdgeConnections(mapping gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetUseEdgeConnections(mapping, enabled)
}
func (self Simple) MapGetUseEdgeConnections(mapping gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).MapGetUseEdgeConnections(mapping))
}
func (self Simple) MapSetEdgeConnectionMargin(mapping gd.RID, margin float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetEdgeConnectionMargin(mapping, gd.Float(margin))
}
func (self Simple) MapGetEdgeConnectionMargin(mapping gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).MapGetEdgeConnectionMargin(mapping)))
}
func (self Simple) MapSetLinkConnectionRadius(mapping gd.RID, radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapSetLinkConnectionRadius(mapping, gd.Float(radius))
}
func (self Simple) MapGetLinkConnectionRadius(mapping gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).MapGetLinkConnectionRadius(mapping)))
}
func (self Simple) MapGetPath(mapping gd.RID, origin gd.Vector3, destination gd.Vector3, optimize bool, navigation_layers int) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).MapGetPath(gc, mapping, origin, destination, optimize, gd.Int(navigation_layers)))
}
func (self Simple) MapGetClosestPointToSegment(mapping gd.RID, start gd.Vector3, end gd.Vector3, use_collision bool) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).MapGetClosestPointToSegment(mapping, start, end, use_collision))
}
func (self Simple) MapGetClosestPoint(mapping gd.RID, to_point gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).MapGetClosestPoint(mapping, to_point))
}
func (self Simple) MapGetClosestPointNormal(mapping gd.RID, to_point gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).MapGetClosestPointNormal(mapping, to_point))
}
func (self Simple) MapGetClosestPointOwner(mapping gd.RID, to_point gd.Vector3) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).MapGetClosestPointOwner(mapping, to_point))
}
func (self Simple) MapGetLinks(mapping gd.RID) gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).MapGetLinks(gc, mapping))
}
func (self Simple) MapGetRegions(mapping gd.RID) gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).MapGetRegions(gc, mapping))
}
func (self Simple) MapGetAgents(mapping gd.RID) gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).MapGetAgents(gc, mapping))
}
func (self Simple) MapGetObstacles(mapping gd.RID) gd.ArrayOf[gd.RID] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.RID](Expert(self).MapGetObstacles(gc, mapping))
}
func (self Simple) MapForceUpdate(mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MapForceUpdate(mapping)
}
func (self Simple) MapGetIterationId(mapping gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).MapGetIterationId(mapping)))
}
func (self Simple) MapGetRandomPoint(mapping gd.RID, navigation_layers int, uniformly bool) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).MapGetRandomPoint(mapping, gd.Int(navigation_layers), uniformly))
}
func (self Simple) QueryPath(parameters [1]classdb.NavigationPathQueryParameters3D, result [1]classdb.NavigationPathQueryResult3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).QueryPath(parameters, result)
}
func (self Simple) RegionCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).RegionCreate())
}
func (self Simple) RegionSetEnabled(region gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetEnabled(region, enabled)
}
func (self Simple) RegionGetEnabled(region gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).RegionGetEnabled(region))
}
func (self Simple) RegionSetUseEdgeConnections(region gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetUseEdgeConnections(region, enabled)
}
func (self Simple) RegionGetUseEdgeConnections(region gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).RegionGetUseEdgeConnections(region))
}
func (self Simple) RegionSetEnterCost(region gd.RID, enter_cost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetEnterCost(region, gd.Float(enter_cost))
}
func (self Simple) RegionGetEnterCost(region gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).RegionGetEnterCost(region)))
}
func (self Simple) RegionSetTravelCost(region gd.RID, travel_cost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetTravelCost(region, gd.Float(travel_cost))
}
func (self Simple) RegionGetTravelCost(region gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).RegionGetTravelCost(region)))
}
func (self Simple) RegionSetOwnerId(region gd.RID, owner_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetOwnerId(region, gd.Int(owner_id))
}
func (self Simple) RegionGetOwnerId(region gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).RegionGetOwnerId(region)))
}
func (self Simple) RegionOwnsPoint(region gd.RID, point gd.Vector3) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).RegionOwnsPoint(region, point))
}
func (self Simple) RegionSetMap(region gd.RID, mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetMap(region, mapping)
}
func (self Simple) RegionGetMap(region gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).RegionGetMap(region))
}
func (self Simple) RegionSetNavigationLayers(region gd.RID, navigation_layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetNavigationLayers(region, gd.Int(navigation_layers))
}
func (self Simple) RegionGetNavigationLayers(region gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).RegionGetNavigationLayers(region)))
}
func (self Simple) RegionSetTransform(region gd.RID, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetTransform(region, transform)
}
func (self Simple) RegionGetTransform(region gd.RID) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).RegionGetTransform(region))
}
func (self Simple) RegionSetNavigationMesh(region gd.RID, navigation_mesh [1]classdb.NavigationMesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionSetNavigationMesh(region, navigation_mesh)
}
func (self Simple) RegionBakeNavigationMesh(navigation_mesh [1]classdb.NavigationMesh, root_node [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegionBakeNavigationMesh(navigation_mesh, root_node)
}
func (self Simple) RegionGetConnectionsCount(region gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).RegionGetConnectionsCount(region)))
}
func (self Simple) RegionGetConnectionPathwayStart(region gd.RID, connection int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).RegionGetConnectionPathwayStart(region, gd.Int(connection)))
}
func (self Simple) RegionGetConnectionPathwayEnd(region gd.RID, connection int) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).RegionGetConnectionPathwayEnd(region, gd.Int(connection)))
}
func (self Simple) RegionGetRandomPoint(region gd.RID, navigation_layers int, uniformly bool) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).RegionGetRandomPoint(region, gd.Int(navigation_layers), uniformly))
}
func (self Simple) LinkCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).LinkCreate())
}
func (self Simple) LinkSetMap(link gd.RID, mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetMap(link, mapping)
}
func (self Simple) LinkGetMap(link gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).LinkGetMap(link))
}
func (self Simple) LinkSetEnabled(link gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetEnabled(link, enabled)
}
func (self Simple) LinkGetEnabled(link gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).LinkGetEnabled(link))
}
func (self Simple) LinkSetBidirectional(link gd.RID, bidirectional bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetBidirectional(link, bidirectional)
}
func (self Simple) LinkIsBidirectional(link gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).LinkIsBidirectional(link))
}
func (self Simple) LinkSetNavigationLayers(link gd.RID, navigation_layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetNavigationLayers(link, gd.Int(navigation_layers))
}
func (self Simple) LinkGetNavigationLayers(link gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).LinkGetNavigationLayers(link)))
}
func (self Simple) LinkSetStartPosition(link gd.RID, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetStartPosition(link, position)
}
func (self Simple) LinkGetStartPosition(link gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).LinkGetStartPosition(link))
}
func (self Simple) LinkSetEndPosition(link gd.RID, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetEndPosition(link, position)
}
func (self Simple) LinkGetEndPosition(link gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).LinkGetEndPosition(link))
}
func (self Simple) LinkSetEnterCost(link gd.RID, enter_cost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetEnterCost(link, gd.Float(enter_cost))
}
func (self Simple) LinkGetEnterCost(link gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).LinkGetEnterCost(link)))
}
func (self Simple) LinkSetTravelCost(link gd.RID, travel_cost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetTravelCost(link, gd.Float(travel_cost))
}
func (self Simple) LinkGetTravelCost(link gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).LinkGetTravelCost(link)))
}
func (self Simple) LinkSetOwnerId(link gd.RID, owner_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinkSetOwnerId(link, gd.Int(owner_id))
}
func (self Simple) LinkGetOwnerId(link gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).LinkGetOwnerId(link)))
}
func (self Simple) AgentCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).AgentCreate())
}
func (self Simple) AgentSetAvoidanceEnabled(agent gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetAvoidanceEnabled(agent, enabled)
}
func (self Simple) AgentGetAvoidanceEnabled(agent gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AgentGetAvoidanceEnabled(agent))
}
func (self Simple) AgentSetUse3dAvoidance(agent gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetUse3dAvoidance(agent, enabled)
}
func (self Simple) AgentGetUse3dAvoidance(agent gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AgentGetUse3dAvoidance(agent))
}
func (self Simple) AgentSetMap(agent gd.RID, mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetMap(agent, mapping)
}
func (self Simple) AgentGetMap(agent gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).AgentGetMap(agent))
}
func (self Simple) AgentSetPaused(agent gd.RID, paused bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetPaused(agent, paused)
}
func (self Simple) AgentGetPaused(agent gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AgentGetPaused(agent))
}
func (self Simple) AgentSetNeighborDistance(agent gd.RID, distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetNeighborDistance(agent, gd.Float(distance))
}
func (self Simple) AgentGetNeighborDistance(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetNeighborDistance(agent)))
}
func (self Simple) AgentSetMaxNeighbors(agent gd.RID, count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetMaxNeighbors(agent, gd.Int(count))
}
func (self Simple) AgentGetMaxNeighbors(agent gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AgentGetMaxNeighbors(agent)))
}
func (self Simple) AgentSetTimeHorizonAgents(agent gd.RID, time_horizon float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetTimeHorizonAgents(agent, gd.Float(time_horizon))
}
func (self Simple) AgentGetTimeHorizonAgents(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetTimeHorizonAgents(agent)))
}
func (self Simple) AgentSetTimeHorizonObstacles(agent gd.RID, time_horizon float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetTimeHorizonObstacles(agent, gd.Float(time_horizon))
}
func (self Simple) AgentGetTimeHorizonObstacles(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetTimeHorizonObstacles(agent)))
}
func (self Simple) AgentSetRadius(agent gd.RID, radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetRadius(agent, gd.Float(radius))
}
func (self Simple) AgentGetRadius(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetRadius(agent)))
}
func (self Simple) AgentSetHeight(agent gd.RID, height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetHeight(agent, gd.Float(height))
}
func (self Simple) AgentGetHeight(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetHeight(agent)))
}
func (self Simple) AgentSetMaxSpeed(agent gd.RID, max_speed float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetMaxSpeed(agent, gd.Float(max_speed))
}
func (self Simple) AgentGetMaxSpeed(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetMaxSpeed(agent)))
}
func (self Simple) AgentSetVelocityForced(agent gd.RID, velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetVelocityForced(agent, velocity)
}
func (self Simple) AgentSetVelocity(agent gd.RID, velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetVelocity(agent, velocity)
}
func (self Simple) AgentGetVelocity(agent gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).AgentGetVelocity(agent))
}
func (self Simple) AgentSetPosition(agent gd.RID, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetPosition(agent, position)
}
func (self Simple) AgentGetPosition(agent gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).AgentGetPosition(agent))
}
func (self Simple) AgentIsMapChanged(agent gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AgentIsMapChanged(agent))
}
func (self Simple) AgentSetAvoidanceCallback(agent gd.RID, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetAvoidanceCallback(agent, callback)
}
func (self Simple) AgentHasAvoidanceCallback(agent gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AgentHasAvoidanceCallback(agent))
}
func (self Simple) AgentSetAvoidanceLayers(agent gd.RID, layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetAvoidanceLayers(agent, gd.Int(layers))
}
func (self Simple) AgentGetAvoidanceLayers(agent gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AgentGetAvoidanceLayers(agent)))
}
func (self Simple) AgentSetAvoidanceMask(agent gd.RID, mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetAvoidanceMask(agent, gd.Int(mask))
}
func (self Simple) AgentGetAvoidanceMask(agent gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AgentGetAvoidanceMask(agent)))
}
func (self Simple) AgentSetAvoidancePriority(agent gd.RID, priority float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AgentSetAvoidancePriority(agent, gd.Float(priority))
}
func (self Simple) AgentGetAvoidancePriority(agent gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).AgentGetAvoidancePriority(agent)))
}
func (self Simple) ObstacleCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ObstacleCreate())
}
func (self Simple) ObstacleSetAvoidanceEnabled(obstacle gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetAvoidanceEnabled(obstacle, enabled)
}
func (self Simple) ObstacleGetAvoidanceEnabled(obstacle gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ObstacleGetAvoidanceEnabled(obstacle))
}
func (self Simple) ObstacleSetUse3dAvoidance(obstacle gd.RID, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetUse3dAvoidance(obstacle, enabled)
}
func (self Simple) ObstacleGetUse3dAvoidance(obstacle gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ObstacleGetUse3dAvoidance(obstacle))
}
func (self Simple) ObstacleSetMap(obstacle gd.RID, mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetMap(obstacle, mapping)
}
func (self Simple) ObstacleGetMap(obstacle gd.RID) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).ObstacleGetMap(obstacle))
}
func (self Simple) ObstacleSetPaused(obstacle gd.RID, paused bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetPaused(obstacle, paused)
}
func (self Simple) ObstacleGetPaused(obstacle gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ObstacleGetPaused(obstacle))
}
func (self Simple) ObstacleSetRadius(obstacle gd.RID, radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetRadius(obstacle, gd.Float(radius))
}
func (self Simple) ObstacleGetRadius(obstacle gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ObstacleGetRadius(obstacle)))
}
func (self Simple) ObstacleSetHeight(obstacle gd.RID, height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetHeight(obstacle, gd.Float(height))
}
func (self Simple) ObstacleGetHeight(obstacle gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).ObstacleGetHeight(obstacle)))
}
func (self Simple) ObstacleSetVelocity(obstacle gd.RID, velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetVelocity(obstacle, velocity)
}
func (self Simple) ObstacleGetVelocity(obstacle gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).ObstacleGetVelocity(obstacle))
}
func (self Simple) ObstacleSetPosition(obstacle gd.RID, position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetPosition(obstacle, position)
}
func (self Simple) ObstacleGetPosition(obstacle gd.RID) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).ObstacleGetPosition(obstacle))
}
func (self Simple) ObstacleSetVertices(obstacle gd.RID, vertices gd.PackedVector3Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetVertices(obstacle, vertices)
}
func (self Simple) ObstacleGetVertices(obstacle gd.RID) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).ObstacleGetVertices(gc, obstacle))
}
func (self Simple) ObstacleSetAvoidanceLayers(obstacle gd.RID, layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ObstacleSetAvoidanceLayers(obstacle, gd.Int(layers))
}
func (self Simple) ObstacleGetAvoidanceLayers(obstacle gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).ObstacleGetAvoidanceLayers(obstacle)))
}
func (self Simple) ParseSourceGeometryData(navigation_mesh [1]classdb.NavigationMesh, source_geometry_data [1]classdb.NavigationMeshSourceGeometryData3D, root_node [1]classdb.Node, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParseSourceGeometryData(navigation_mesh, source_geometry_data, root_node, callback)
}
func (self Simple) BakeFromSourceGeometryData(navigation_mesh [1]classdb.NavigationMesh, source_geometry_data [1]classdb.NavigationMeshSourceGeometryData3D, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BakeFromSourceGeometryData(navigation_mesh, source_geometry_data, callback)
}
func (self Simple) BakeFromSourceGeometryDataAsync(navigation_mesh [1]classdb.NavigationMesh, source_geometry_data [1]classdb.NavigationMeshSourceGeometryData3D, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BakeFromSourceGeometryDataAsync(navigation_mesh, source_geometry_data, callback)
}
func (self Simple) IsBakingNavigationMesh(navigation_mesh [1]classdb.NavigationMesh) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBakingNavigationMesh(navigation_mesh))
}
func (self Simple) SourceGeometryParserCreate() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).SourceGeometryParserCreate())
}
func (self Simple) SourceGeometryParserSetCallback(parser gd.RID, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SourceGeometryParserSetCallback(parser, callback)
}
func (self Simple) SimplifyPath(path gd.PackedVector3Array, epsilon float64) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).SimplifyPath(gc, path, gd.Float(epsilon)))
}
func (self Simple) FreeRid(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FreeRid(rid)
}
func (self Simple) SetActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActive(active)
}
func (self Simple) SetDebugEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugEnabled(enabled)
}
func (self Simple) GetDebugEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDebugEnabled())
}
func (self Simple) GetProcessInfo(process_info classdb.NavigationServer3DProcessInfo) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProcessInfo(process_info)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationServer3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns all created navigation map [RID]s on the NavigationServer. This returns both 2D and 3D created navigation maps as there is technically no distinction between them.
*/
//go:nosplit
func (self class) GetMaps(ctx gd.Lifetime) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_get_maps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
/*
Create a new map.
*/
//go:nosplit
func (self class) MapCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the map active.
*/
//go:nosplit
func (self class) MapSetActive(mapping gd.RID, active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns true if the map is active.
*/
//go:nosplit
func (self class) MapIsActive(mapping gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the map up direction.
*/
//go:nosplit
func (self class) MapSetUp(mapping gd.RID, up gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, up)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_up, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the map's up direction.
*/
//go:nosplit
func (self class) MapGetUp(mapping gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_up, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the map cell size used to rasterize the navigation mesh vertices on the XZ plane. Must match with the cell size of the used navigation meshes.
*/
//go:nosplit
func (self class) MapSetCellSize(mapping gd.RID, cell_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, cell_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the map cell size used to rasterize the navigation mesh vertices on the XZ plane.
*/
//go:nosplit
func (self class) MapGetCellSize(mapping gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the map cell height used to rasterize the navigation mesh vertices on the Y axis. Must match with the cell height of the used navigation meshes.
*/
//go:nosplit
func (self class) MapSetCellHeight(mapping gd.RID, cell_height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, cell_height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_cell_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the map cell height used to rasterize the navigation mesh vertices on the Y axis.
*/
//go:nosplit
func (self class) MapGetCellHeight(mapping gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_cell_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the map's internal merge rasterizer cell scale used to control merging sensitivity.
*/
//go:nosplit
func (self class) MapSetMergeRasterizerCellScale(mapping gd.RID, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_merge_rasterizer_cell_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns map's internal merge rasterizer cell scale.
*/
//go:nosplit
func (self class) MapGetMergeRasterizerCellScale(mapping gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_merge_rasterizer_cell_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the navigation [param map] edge connection use. If [param enabled] is [code]true[/code], the navigation map allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) MapSetUseEdgeConnections(mapping gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns true if the navigation [param map] allows navigation regions to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) MapGetUseEdgeConnections(mapping gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the map edge connection margin used to weld the compatible region edges.
*/
//go:nosplit
func (self class) MapSetEdgeConnectionMargin(mapping gd.RID, margin gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_edge_connection_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the edge connection margin of the map. This distance is the minimum vertex distance needed to connect two edges from different regions.
*/
//go:nosplit
func (self class) MapGetEdgeConnectionMargin(mapping gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_edge_connection_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the map's link connection radius used to connect links to navigation polygons.
*/
//go:nosplit
func (self class) MapSetLinkConnectionRadius(mapping gd.RID, radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_set_link_connection_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the link connection radius of the map. This distance is the maximum range any link will search for navigation mesh polygons to connect to.
*/
//go:nosplit
func (self class) MapGetLinkConnectionRadius(mapping gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_link_connection_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the navigation path to reach the destination from the origin. [param navigation_layers] is a bitmask of all region navigation layers that are allowed to be in the path.
*/
//go:nosplit
func (self class) MapGetPath(ctx gd.Lifetime, mapping gd.RID, origin gd.Vector3, destination gd.Vector3, optimize bool, navigation_layers gd.Int) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, origin)
	callframe.Arg(frame, destination)
	callframe.Arg(frame, optimize)
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the closest point between the navigation surface and the segment.
*/
//go:nosplit
func (self class) MapGetClosestPointToSegment(mapping gd.RID, start gd.Vector3, end gd.Vector3, use_collision bool) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	callframe.Arg(frame, use_collision)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_closest_point_to_segment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the point closest to the provided [param to_point] on the navigation mesh surface.
*/
//go:nosplit
func (self class) MapGetClosestPoint(mapping gd.RID, to_point gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the normal for the point returned by [method map_get_closest_point].
*/
//go:nosplit
func (self class) MapGetClosestPointNormal(mapping gd.RID, to_point gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_closest_point_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the owner region RID for the point returned by [method map_get_closest_point].
*/
//go:nosplit
func (self class) MapGetClosestPointOwner(mapping gd.RID, to_point gd.Vector3) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_closest_point_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns all navigation link [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetLinks(ctx gd.Lifetime, mapping gd.RID) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_links, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
/*
Returns all navigation regions [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetRegions(ctx gd.Lifetime, mapping gd.RID) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_regions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
/*
Returns all navigation agents [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetAgents(ctx gd.Lifetime, mapping gd.RID) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_agents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
/*
Returns all navigation obstacle [RID]s that are currently assigned to the requested navigation [param map].
*/
//go:nosplit
func (self class) MapGetObstacles(ctx gd.Lifetime, mapping gd.RID) gd.ArrayOf[gd.RID] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.RID](ret)
}
/*
This function immediately forces synchronization of the specified navigation [param map] [RID]. By default navigation maps are only synchronized at the end of each physics frame. This function can be used to immediately (re)calculate all the navigation meshes and region connections of the navigation map. This makes it possible to query a navigation path for a changed map immediately and in the same frame (multiple times if needed).
Due to technical restrictions the current NavigationServer command queue will be flushed. This means all already queued update commands for this physics frame will be executed, even those intended for other maps, regions and agents not part of the specified map. The expensive computation of the navigation meshes and region connections of a map will only be done for the specified map. Other maps will receive the normal synchronization at the end of the physics frame. Should the specified map receive changes after the forced update it will update again as well when the other maps receive their update.
Avoidance processing and dispatch of the [code]safe_velocity[/code] signals is unaffected by this function and continues to happen for all maps and agents at the end of the physics frame.
[b]Note:[/b] With great power comes great responsibility. This function should only be used by users that really know what they are doing and have a good reason for it. Forcing an immediate update of a navigation map requires locking the NavigationServer and flushing the entire NavigationServer command queue. Not only can this severely impact the performance of a game but it can also introduce bugs if used inappropriately without much foresight.
*/
//go:nosplit
func (self class) MapForceUpdate(mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_force_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current iteration id of the navigation map. Every time the navigation map changes and synchronizes the iteration id increases. An iteration id of 0 means the navigation map has never synchronized.
[b]Note:[/b] The iteration id will wrap back to 1 after reaching its range limit.
*/
//go:nosplit
func (self class) MapGetIterationId(mapping gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_iteration_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) MapGetRandomPoint(mapping gd.RID, navigation_layers gd.Int, uniformly bool) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	callframe.Arg(frame, navigation_layers)
	callframe.Arg(frame, uniformly)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_map_get_random_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Queries a path in a given navigation map. Start and target position and other parameters are defined through [NavigationPathQueryParameters3D]. Updates the provided [NavigationPathQueryResult3D] result object with the path among other results requested by the query.
*/
//go:nosplit
func (self class) QueryPath(parameters object.NavigationPathQueryParameters3D, result object.NavigationPathQueryResult3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(parameters[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(result[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_query_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a new region.
*/
//go:nosplit
func (self class) RegionCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enabled] is [code]true[/code], the specified [param region] will contribute to its current navigation map.
*/
//go:nosplit
func (self class) RegionSetEnabled(region gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified [param region] is enabled.
*/
//go:nosplit
func (self class) RegionGetEnabled(region gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enabled] is [code]true[/code], the navigation [param region] will use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) RegionSetUseEdgeConnections(region gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns true if the navigation [param region] is set to use edge connections to connect with other navigation regions within proximity of the navigation map edge connection margin.
*/
//go:nosplit
func (self class) RegionGetUseEdgeConnections(region gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param enter_cost] for this [param region].
*/
//go:nosplit
func (self class) RegionSetEnterCost(region gd.RID, enter_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, enter_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the enter cost of this [param region].
*/
//go:nosplit
func (self class) RegionGetEnterCost(region gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param travel_cost] for this [param region].
*/
//go:nosplit
func (self class) RegionSetTravelCost(region gd.RID, travel_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, travel_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the travel cost of this [param region].
*/
//go:nosplit
func (self class) RegionGetTravelCost(region gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the [code]ObjectID[/code] of the object which manages this region.
*/
//go:nosplit
func (self class) RegionSetOwnerId(region gd.RID, owner_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, owner_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_owner_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]ObjectID[/code] of the object which manages this region.
*/
//go:nosplit
func (self class) RegionGetOwnerId(region gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_owner_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) RegionOwnsPoint(region gd.RID, point gd.Vector3) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_owns_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the map for the region.
*/
//go:nosplit
func (self class) RegionSetMap(region gd.RID, mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation map [RID] the requested [param region] is currently assigned to.
*/
//go:nosplit
func (self class) RegionGetMap(region gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the region's navigation layers. This allows selecting regions from a path request (when using [method NavigationServer3D.map_get_path]).
*/
//go:nosplit
func (self class) RegionSetNavigationLayers(region gd.RID, navigation_layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the region's navigation layers.
*/
//go:nosplit
func (self class) RegionGetNavigationLayers(region gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the global transformation for the region.
*/
//go:nosplit
func (self class) RegionSetTransform(region gd.RID, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the global transformation of this [param region].
*/
//go:nosplit
func (self class) RegionGetTransform(region gd.RID) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the navigation mesh for the region.
*/
//go:nosplit
func (self class) RegionSetNavigationMesh(region gd.RID, navigation_mesh object.NavigationMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_set_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Bakes the [param navigation_mesh] with bake source geometry collected starting from the [param root_node].
*/
//go:nosplit
func (self class) RegionBakeNavigationMesh(navigation_mesh object.NavigationMesh, root_node object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(root_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_bake_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns how many connections this [param region] has with other regions in the map.
*/
//go:nosplit
func (self class) RegionGetConnectionsCount(region gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_connections_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the starting point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
//go:nosplit
func (self class) RegionGetConnectionPathwayStart(region gd.RID, connection gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, connection)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_connection_pathway_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ending point of a connection door. [param connection] is an index between 0 and the return value of [method region_get_connections_count].
*/
//go:nosplit
func (self class) RegionGetConnectionPathwayEnd(region gd.RID, connection gd.Int) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, connection)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_connection_pathway_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) RegionGetRandomPoint(region gd.RID, navigation_layers gd.Int, uniformly bool) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	callframe.Arg(frame, navigation_layers)
	callframe.Arg(frame, uniformly)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_region_get_random_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Create a new link between two positions on a map.
*/
//go:nosplit
func (self class) LinkCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the navigation map [RID] for the link.
*/
//go:nosplit
func (self class) LinkSetMap(link gd.RID, mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation map [RID] the requested [param link] is currently assigned to.
*/
//go:nosplit
func (self class) LinkGetMap(link gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enabled] is [code]true[/code], the specified [param link] will contribute to its current navigation map.
*/
//go:nosplit
func (self class) LinkSetEnabled(link gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified [param link] is enabled.
*/
//go:nosplit
func (self class) LinkGetEnabled(link gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether this [param link] can be travelled in both directions.
*/
//go:nosplit
func (self class) LinkSetBidirectional(link gd.RID, bidirectional bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether this [param link] can be travelled in both directions.
*/
//go:nosplit
func (self class) LinkIsBidirectional(link gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the links's navigation layers. This allows selecting links from a path request (when using [method NavigationServer3D.map_get_path]).
*/
//go:nosplit
func (self class) LinkSetNavigationLayers(link gd.RID, navigation_layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation layers for this [param link].
*/
//go:nosplit
func (self class) LinkGetNavigationLayers(link gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the entry position for this [param link].
*/
//go:nosplit
func (self class) LinkSetStartPosition(link gd.RID, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the starting position of this [param link].
*/
//go:nosplit
func (self class) LinkGetStartPosition(link gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the exit position for the [param link].
*/
//go:nosplit
func (self class) LinkSetEndPosition(link gd.RID, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the ending position of this [param link].
*/
//go:nosplit
func (self class) LinkGetEndPosition(link gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param enter_cost] for this [param link].
*/
//go:nosplit
func (self class) LinkSetEnterCost(link gd.RID, enter_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, enter_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the enter cost of this [param link].
*/
//go:nosplit
func (self class) LinkGetEnterCost(link gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param travel_cost] for this [param link].
*/
//go:nosplit
func (self class) LinkSetTravelCost(link gd.RID, travel_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, travel_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the travel cost of this [param link].
*/
//go:nosplit
func (self class) LinkGetTravelCost(link gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the [code]ObjectID[/code] of the object which manages this link.
*/
//go:nosplit
func (self class) LinkSetOwnerId(link gd.RID, owner_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	callframe.Arg(frame, owner_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_set_owner_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]ObjectID[/code] of the object which manages this link.
*/
//go:nosplit
func (self class) LinkGetOwnerId(link gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, link)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_link_get_owner_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates the agent.
*/
//go:nosplit
func (self class) AgentCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enabled] is [code]true[/code], the provided [param agent] calculates avoidance.
*/
//go:nosplit
func (self class) AgentSetAvoidanceEnabled(agent gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the provided [param agent] has avoidance enabled.
*/
//go:nosplit
func (self class) AgentGetAvoidanceEnabled(agent gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AgentSetUse3dAvoidance(agent gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the provided [param agent] uses avoidance in 3D space Vector3(x,y,z) instead of horizontal 2D Vector2(x,y) / Vector3(x,0.0,z).
*/
//go:nosplit
func (self class) AgentGetUse3dAvoidance(agent gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Puts the agent in the map.
*/
//go:nosplit
func (self class) AgentSetMap(agent gd.RID, mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation map [RID] the requested [param agent] is currently assigned to.
*/
//go:nosplit
func (self class) AgentGetMap(agent gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param paused] is true the specified [param agent] will not be processed, e.g. calculate avoidance velocities or receive avoidance callbacks.
*/
//go:nosplit
func (self class) AgentSetPaused(agent gd.RID, paused bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, paused)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified [param agent] is paused.
*/
//go:nosplit
func (self class) AgentGetPaused(agent gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum distance to other agents this agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
//go:nosplit
func (self class) AgentSetNeighborDistance(agent gd.RID, distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum distance to other agents the specified [param agent] takes into account in the navigation.
*/
//go:nosplit
func (self class) AgentGetNeighborDistance(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum number of other agents the agent takes into account in the navigation. The larger this number, the longer the running time of the simulation. If the number is too low, the simulation will not be safe.
*/
//go:nosplit
func (self class) AgentSetMaxNeighbors(agent gd.RID, count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum number of other agents the specified [param agent] takes into account in the navigation.
*/
//go:nosplit
func (self class) AgentGetMaxNeighbors(agent gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to other agents. The larger this number, the sooner this agent will respond to the presence of other agents, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
//go:nosplit
func (self class) AgentSetTimeHorizonAgents(agent gd.RID, time_horizon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, time_horizon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to other agents.
*/
//go:nosplit
func (self class) AgentGetTimeHorizonAgents(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The minimal amount of time for which the agent's velocities that are computed by the simulation are safe with respect to static avoidance obstacles. The larger this number, the sooner this agent will respond to the presence of static avoidance obstacles, but the less freedom this agent has in choosing its velocities. A too high value will slow down agents movement considerably. Must be positive.
*/
//go:nosplit
func (self class) AgentSetTimeHorizonObstacles(agent gd.RID, time_horizon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, time_horizon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimal amount of time for which the specified [param agent]'s velocities that are computed by the simulation are safe with respect to static avoidance obstacles.
*/
//go:nosplit
func (self class) AgentGetTimeHorizonObstacles(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the radius of the agent.
*/
//go:nosplit
func (self class) AgentSetRadius(agent gd.RID, radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the radius of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetRadius(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the provided [param agent] [param height].
*/
//go:nosplit
func (self class) AgentSetHeight(agent gd.RID, height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]height[/code] of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetHeight(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum speed of the agent. Must be positive.
*/
//go:nosplit
func (self class) AgentSetMaxSpeed(agent gd.RID, max_speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, max_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_max_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum speed of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetMaxSpeed(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_max_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity] for the specified [param agent]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
//go:nosplit
func (self class) AgentSetVelocityForced(agent gd.RID, velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_velocity_forced, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets [param velocity] as the new wanted velocity for the specified [param agent]. The avoidance simulation will try to fulfill this velocity if possible but will modify it to avoid collision with other agent's and obstacles. When an agent is teleported to a new position use [method agent_set_velocity_forced] as well to reset the internal simulation velocity.
*/
//go:nosplit
func (self class) AgentSetVelocity(agent gd.RID, velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the velocity of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetVelocity(agent gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the position of the agent in world space.
*/
//go:nosplit
func (self class) AgentSetPosition(agent gd.RID, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the specified [param agent] in world space.
*/
//go:nosplit
func (self class) AgentGetPosition(agent gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns true if the map got changed the previous frame.
*/
//go:nosplit
func (self class) AgentIsMapChanged(agent gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_is_map_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the callback [Callable] that gets called after each avoidance processing step for the [param agent]. The calculated [code]safe_velocity[/code] will be dispatched with a signal to the object just before the physics calculations.
[b]Note:[/b] Created callbacks are always processed independently of the SceneTree state as long as the agent is on a navigation map and not freed. To disable the dispatch of a callback from an agent use [method agent_set_avoidance_callback] again with an empty [Callable].
*/
//go:nosplit
func (self class) AgentSetAvoidanceCallback(agent gd.RID, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_avoidance_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Return [code]true[/code] if the specified [param agent] has an avoidance callback.
*/
//go:nosplit
func (self class) AgentHasAvoidanceCallback(agent gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_has_avoidance_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the agent's [code]avoidance_layers[/code] bitmask.
*/
//go:nosplit
func (self class) AgentSetAvoidanceLayers(agent gd.RID, layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidanceLayers(agent gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the agent's [code]avoidance_mask[/code] bitmask.
*/
//go:nosplit
func (self class) AgentSetAvoidanceMask(agent gd.RID, mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]avoidance_mask[/code] bitmask of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidanceMask(agent gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set the agent's [code]avoidance_priority[/code] with a [param priority] between 0.0 (lowest priority) to 1.0 (highest priority).
The specified [param agent] does not adjust the velocity for other agents that would match the [code]avoidance_mask[/code] but have a lower [code]avoidance_priority[/code]. This in turn makes the other agents with lower priority adjust their velocities even more to avoid collision with this agent.
*/
//go:nosplit
func (self class) AgentSetAvoidancePriority(agent gd.RID, priority gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_set_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]avoidance_priority[/code] of the specified [param agent].
*/
//go:nosplit
func (self class) AgentGetAvoidancePriority(agent gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, agent)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_agent_get_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new obstacle.
*/
//go:nosplit
func (self class) ObstacleCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enabled] is [code]true[/code], the provided [param obstacle] affects avoidance using agents.
*/
//go:nosplit
func (self class) ObstacleSetAvoidanceEnabled(obstacle gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the provided [param obstacle] has avoidance enabled.
*/
//go:nosplit
func (self class) ObstacleGetAvoidanceEnabled(obstacle gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets if the [param obstacle] uses the 2D avoidance or the 3D avoidance while avoidance is enabled.
*/
//go:nosplit
func (self class) ObstacleSetUse3dAvoidance(obstacle gd.RID, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the provided [param obstacle] uses avoidance in 3D space Vector3(x,y,z) instead of horizontal 2D Vector2(x,y) / Vector3(x,0.0,z).
*/
//go:nosplit
func (self class) ObstacleGetUse3dAvoidance(obstacle gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns the [param obstacle] to a navigation map.
*/
//go:nosplit
func (self class) ObstacleSetMap(obstacle gd.RID, mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation map [RID] the requested [param obstacle] is currently assigned to.
*/
//go:nosplit
func (self class) ObstacleGetMap(obstacle gd.RID) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param paused] is true the specified [param obstacle] will not be processed, e.g. affect avoidance velocities.
*/
//go:nosplit
func (self class) ObstacleSetPaused(obstacle gd.RID, paused bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, paused)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the specified [param obstacle] is paused.
*/
//go:nosplit
func (self class) ObstacleGetPaused(obstacle gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_paused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the radius of the dynamic obstacle.
*/
//go:nosplit
func (self class) ObstacleSetRadius(obstacle gd.RID, radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the radius of the specified dynamic [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetRadius(obstacle gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param height] for the [param obstacle]. In 3D agents will ignore obstacles that are above or below them while using 2D avoidance.
*/
//go:nosplit
func (self class) ObstacleSetHeight(obstacle gd.RID, height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]height[/code] of the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetHeight(obstacle gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets [param velocity] of the dynamic [param obstacle]. Allows other agents to better predict the movement of the dynamic obstacle. Only works in combination with the radius of the obstacle.
*/
//go:nosplit
func (self class) ObstacleSetVelocity(obstacle gd.RID, velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the velocity of the specified dynamic [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetVelocity(obstacle gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the [param position] in world space for the [param obstacle].
*/
//go:nosplit
func (self class) ObstacleSetPosition(obstacle gd.RID, position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the position of the specified [param obstacle] in world space.
*/
//go:nosplit
func (self class) ObstacleGetPosition(obstacle gd.RID) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the outline vertices for the obstacle. If the vertices are winded in clockwise order agents will be pushed in by the obstacle, else they will be pushed out.
*/
//go:nosplit
func (self class) ObstacleSetVertices(obstacle gd.RID, vertices gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, mmm.Get(vertices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the outline vertices for the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetVertices(ctx gd.Lifetime, obstacle gd.RID) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Set the obstacles's [code]avoidance_layers[/code] bitmask.
*/
//go:nosplit
func (self class) ObstacleSetAvoidanceLayers(obstacle gd.RID, layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [code]avoidance_layers[/code] bitmask of the specified [param obstacle].
*/
//go:nosplit
func (self class) ObstacleGetAvoidanceLayers(obstacle gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, obstacle)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_obstacle_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) ParseSourceGeometryData(navigation_mesh object.NavigationMesh, source_geometry_data object.NavigationMeshSourceGeometryData3D, root_node object.Node, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(source_geometry_data[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(root_node[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_parse_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryData(navigation_mesh object.NavigationMesh, source_geometry_data object.NavigationMeshSourceGeometryData3D, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(source_geometry_data[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_bake_from_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data] as an async task running on a background thread. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryDataAsync(navigation_mesh object.NavigationMesh, source_geometry_data object.NavigationMeshSourceGeometryData3D, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(source_geometry_data[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_bake_from_source_geometry_data_async, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] when the provided navigation mesh is being baked on a background thread.
*/
//go:nosplit
func (self class) IsBakingNavigationMesh(navigation_mesh object.NavigationMesh) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_is_baking_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new source geometry parser. If a [Callable] is set for the parser with [method source_geometry_parser_set_callback] the callback will be called for every single node that gets parsed whenever [method parse_source_geometry_data] is used.
*/
//go:nosplit
func (self class) SourceGeometryParserCreate() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_source_geometry_parser_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SourceGeometryParserSetCallback(parser gd.RID, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_source_geometry_parser_set_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a simplified version of [param path] with less critical path points removed. The simplification amount is in worlds units and controlled by [param epsilon]. The simplification uses a variant of Ramer-Douglas-Peucker algorithm for curve point decimation.
Path simplification can be helpful to mitigate various path following issues that can arise with certain agent types and script behaviors. E.g. "steering" agents or avoidance in "open fields".
*/
//go:nosplit
func (self class) SimplifyPath(ctx gd.Lifetime, path gd.PackedVector3Array, epsilon gd.Float) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Destroys the given RID.
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Control activation of this server.
*/
//go:nosplit
func (self class) SetActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code] enables debug mode on the NavigationServer.
*/
//go:nosplit
func (self class) SetDebugEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_set_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] when the NavigationServer has debug enabled.
*/
//go:nosplit
func (self class) GetDebugEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_get_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns information about the current state of the NavigationServer. See [enum ProcessInfo] for a list of available states.
*/
//go:nosplit
func (self class) GetProcessInfo(process_info classdb.NavigationServer3DProcessInfo) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, process_info)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationServer3D.Bind_get_process_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsNavigationServer3D() Expert { return self[0].AsNavigationServer3D() }


//go:nosplit
func (self Simple) AsNavigationServer3D() Simple { return self[0].AsNavigationServer3D() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("NavigationServer3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ProcessInfo = classdb.NavigationServer3DProcessInfo

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
