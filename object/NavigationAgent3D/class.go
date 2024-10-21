package NavigationAgent3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 3D agent used to pathfind to a position while avoiding static and dynamic obstacles. The calculation can be used by the parent node to dynamically move it along the path. Requires navigation data to work correctly.
Dynamic obstacles are avoided using RVO collision avoidance. Avoidance is computed before physics, so the pathfinding information can be used safely in the physics step.
[b]Note:[/b] After setting the [member target_position] property, the [method get_next_path_position] method must be used once every physics frame to update the internal path logic of the navigation agent. The vector position it returns should be used as the next movement position for the agent's parent node.

*/
type Simple [1]classdb.NavigationAgent3D
func (self Simple) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetRid())
}
func (self Simple) SetAvoidanceEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAvoidanceEnabled(enabled)
}
func (self Simple) GetAvoidanceEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAvoidanceEnabled())
}
func (self Simple) SetPathDesiredDistance(desired_distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathDesiredDistance(gd.Float(desired_distance))
}
func (self Simple) GetPathDesiredDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPathDesiredDistance()))
}
func (self Simple) SetTargetDesiredDistance(desired_distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetDesiredDistance(gd.Float(desired_distance))
}
func (self Simple) GetTargetDesiredDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTargetDesiredDistance()))
}
func (self Simple) SetRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRadius(gd.Float(radius))
}
func (self Simple) GetRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRadius()))
}
func (self Simple) SetHeight(height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeight(gd.Float(height))
}
func (self Simple) GetHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHeight()))
}
func (self Simple) SetPathHeightOffset(path_height_offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathHeightOffset(gd.Float(path_height_offset))
}
func (self Simple) GetPathHeightOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPathHeightOffset()))
}
func (self Simple) SetUse3dAvoidance(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUse3dAvoidance(enabled)
}
func (self Simple) GetUse3dAvoidance() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUse3dAvoidance())
}
func (self Simple) SetKeepYVelocity(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeepYVelocity(enabled)
}
func (self Simple) GetKeepYVelocity() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetKeepYVelocity())
}
func (self Simple) SetNeighborDistance(neighbor_distance float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNeighborDistance(gd.Float(neighbor_distance))
}
func (self Simple) GetNeighborDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetNeighborDistance()))
}
func (self Simple) SetMaxNeighbors(max_neighbors int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxNeighbors(gd.Int(max_neighbors))
}
func (self Simple) GetMaxNeighbors() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxNeighbors()))
}
func (self Simple) SetTimeHorizonAgents(time_horizon float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTimeHorizonAgents(gd.Float(time_horizon))
}
func (self Simple) GetTimeHorizonAgents() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTimeHorizonAgents()))
}
func (self Simple) SetTimeHorizonObstacles(time_horizon float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTimeHorizonObstacles(gd.Float(time_horizon))
}
func (self Simple) GetTimeHorizonObstacles() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTimeHorizonObstacles()))
}
func (self Simple) SetMaxSpeed(max_speed float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxSpeed(gd.Float(max_speed))
}
func (self Simple) GetMaxSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMaxSpeed()))
}
func (self Simple) SetPathMaxDistance(max_speed float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathMaxDistance(gd.Float(max_speed))
}
func (self Simple) GetPathMaxDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPathMaxDistance()))
}
func (self Simple) SetNavigationLayers(navigation_layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationLayers(gd.Int(navigation_layers))
}
func (self Simple) GetNavigationLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNavigationLayers()))
}
func (self Simple) SetNavigationLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}
func (self Simple) GetNavigationLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNavigationLayerValue(gd.Int(layer_number)))
}
func (self Simple) SetPathfindingAlgorithm(pathfinding_algorithm classdb.NavigationPathQueryParameters3DPathfindingAlgorithm) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathfindingAlgorithm(pathfinding_algorithm)
}
func (self Simple) GetPathfindingAlgorithm() classdb.NavigationPathQueryParameters3DPathfindingAlgorithm {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NavigationPathQueryParameters3DPathfindingAlgorithm(Expert(self).GetPathfindingAlgorithm())
}
func (self Simple) SetPathPostprocessing(path_postprocessing classdb.NavigationPathQueryParameters3DPathPostProcessing) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathPostprocessing(path_postprocessing)
}
func (self Simple) GetPathPostprocessing() classdb.NavigationPathQueryParameters3DPathPostProcessing {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NavigationPathQueryParameters3DPathPostProcessing(Expert(self).GetPathPostprocessing())
}
func (self Simple) SetPathMetadataFlags(flags classdb.NavigationPathQueryParameters3DPathMetadataFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPathMetadataFlags(flags)
}
func (self Simple) GetPathMetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NavigationPathQueryParameters3DPathMetadataFlags(Expert(self).GetPathMetadataFlags())
}
func (self Simple) SetNavigationMap(navigation_map gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationMap(navigation_map)
}
func (self Simple) GetNavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetNavigationMap())
}
func (self Simple) SetTargetPosition(position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetPosition(position)
}
func (self Simple) GetTargetPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetTargetPosition())
}
func (self Simple) SetSimplifyPath(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSimplifyPath(enabled)
}
func (self Simple) GetSimplifyPath() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSimplifyPath())
}
func (self Simple) SetSimplifyEpsilon(epsilon float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSimplifyEpsilon(gd.Float(epsilon))
}
func (self Simple) GetSimplifyEpsilon() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSimplifyEpsilon()))
}
func (self Simple) GetNextPathPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetNextPathPosition())
}
func (self Simple) SetVelocityForced(velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVelocityForced(velocity)
}
func (self Simple) SetVelocity(velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVelocity(velocity)
}
func (self Simple) GetVelocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetVelocity())
}
func (self Simple) DistanceToTarget() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).DistanceToTarget()))
}
func (self Simple) GetCurrentNavigationResult() [1]classdb.NavigationPathQueryResult3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.NavigationPathQueryResult3D(Expert(self).GetCurrentNavigationResult(gc))
}
func (self Simple) GetCurrentNavigationPath() gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).GetCurrentNavigationPath(gc))
}
func (self Simple) GetCurrentNavigationPathIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCurrentNavigationPathIndex()))
}
func (self Simple) IsTargetReached() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTargetReached())
}
func (self Simple) IsTargetReachable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsTargetReachable())
}
func (self Simple) IsNavigationFinished() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNavigationFinished())
}
func (self Simple) GetFinalPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetFinalPosition())
}
func (self Simple) SetAvoidanceLayers(layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAvoidanceLayers(gd.Int(layers))
}
func (self Simple) GetAvoidanceLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAvoidanceLayers()))
}
func (self Simple) SetAvoidanceMask(mask int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAvoidanceMask(gd.Int(mask))
}
func (self Simple) GetAvoidanceMask() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAvoidanceMask()))
}
func (self Simple) SetAvoidanceLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAvoidanceLayerValue(gd.Int(layer_number), value)
}
func (self Simple) GetAvoidanceLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAvoidanceLayerValue(gd.Int(layer_number)))
}
func (self Simple) SetAvoidanceMaskValue(mask_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAvoidanceMaskValue(gd.Int(mask_number), value)
}
func (self Simple) GetAvoidanceMaskValue(mask_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAvoidanceMaskValue(gd.Int(mask_number)))
}
func (self Simple) SetAvoidancePriority(priority float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAvoidancePriority(gd.Float(priority))
}
func (self Simple) GetAvoidancePriority() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAvoidancePriority()))
}
func (self Simple) SetDebugEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugEnabled(enabled)
}
func (self Simple) GetDebugEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDebugEnabled())
}
func (self Simple) SetDebugUseCustom(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugUseCustom(enabled)
}
func (self Simple) GetDebugUseCustom() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDebugUseCustom())
}
func (self Simple) SetDebugPathCustomColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugPathCustomColor(color)
}
func (self Simple) GetDebugPathCustomColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetDebugPathCustomColor())
}
func (self Simple) SetDebugPathCustomPointSize(point_size float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDebugPathCustomPointSize(gd.Float(point_size))
}
func (self Simple) GetDebugPathCustomPointSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDebugPathCustomPointSize()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationAgent3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [RID] of this agent on the [NavigationServer3D].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAvoidanceEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAvoidanceEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathDesiredDistance(desired_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, desired_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_path_desired_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathDesiredDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_path_desired_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetDesiredDistance(desired_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, desired_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_target_desired_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetDesiredDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_target_desired_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathHeightOffset(path_height_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, path_height_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_path_height_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathHeightOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_path_height_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUse3dAvoidance(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUse3dAvoidance() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_use_3d_avoidance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepYVelocity(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_keep_y_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetKeepYVelocity() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_keep_y_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNeighborDistance(neighbor_distance gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, neighbor_distance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNeighborDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxNeighbors(max_neighbors gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_neighbors)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxNeighbors() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTimeHorizonAgents(time_horizon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_horizon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimeHorizonAgents() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTimeHorizonObstacles(time_horizon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, time_horizon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTimeHorizonObstacles() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSpeed(max_speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_max_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxSpeed() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_max_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathMaxDistance(max_speed gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_speed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_path_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathMaxDistance() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_path_max_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm classdb.NavigationPathQueryParameters3DPathfindingAlgorithm)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathfindingAlgorithm() classdb.NavigationPathQueryParameters3DPathfindingAlgorithm {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathfindingAlgorithm](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathPostprocessing(path_postprocessing classdb.NavigationPathQueryParameters3DPathPostProcessing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, path_postprocessing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathPostprocessing() classdb.NavigationPathQueryParameters3DPathPostProcessing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathPostProcessing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPathMetadataFlags(flags classdb.NavigationPathQueryParameters3DPathMetadataFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_path_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathMetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathMetadataFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_path_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [RID] of the navigation map this NavigationAgent node should use and also updates the [code]agent[/code] on the NavigationServer.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the navigation map for this NavigationAgent node. This function returns always the map set on the NavigationAgent node and not the map of the abstract agent on the NavigationServer. If the agent map is changed directly with the NavigationServer API the NavigationAgent node will not be aware of the map change. Use [method set_navigation_map] to change the navigation map for the NavigationAgent and also update the agent on the NavigationServer.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetPosition(position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSimplifyPath(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimplifyPath() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSimplifyEpsilon(epsilon gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, epsilon)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimplifyEpsilon() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next position in global coordinates that can be moved to, making sure that there are no static objects in the way. If the agent does not have a navigation path, it will return the position of the agent's parent. The use of this function once every physics frame is required to update the internal path logic of the NavigationAgent.
*/
//go:nosplit
func (self class) GetNextPathPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_next_path_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
//go:nosplit
func (self class) SetVelocityForced(velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_velocity_forced, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVelocity(velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocity() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the distance to the target position, using the agent's global position. The user must set [member target_position] in order for this to be accurate.
*/
//go:nosplit
func (self class) DistanceToTarget() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_distance_to_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the path query result for the path the agent is currently following.
*/
//go:nosplit
func (self class) GetCurrentNavigationResult(ctx gd.Lifetime) object.NavigationPathQueryResult3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_current_navigation_result, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.NavigationPathQueryResult3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns this agent's current path from start to finish in global coordinates. The path only updates when the target position is changed or the agent requires a repath. The path array is not intended to be used in direct path movement as the agent has its own internal path logic that would get corrupted by changing the path array manually. Use the intended [method get_next_path_position] once every physics frame to receive the next path point for the agents movement as this function also updates the internal path logic.
*/
//go:nosplit
func (self class) GetCurrentNavigationPath(ctx gd.Lifetime) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_current_navigation_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns which index the agent is currently on in the navigation path's [PackedVector3Array].
*/
//go:nosplit
func (self class) GetCurrentNavigationPathIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_current_navigation_path_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the agent reached the target, i.e. the agent moved within [member target_desired_distance] of the [member target_position]. It may not always be possible to reach the target but it should always be possible to reach the final position. See [method get_final_position].
*/
//go:nosplit
func (self class) IsTargetReached() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_is_target_reached, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [method get_final_position] is within [member target_desired_distance] of the [member target_position].
*/
//go:nosplit
func (self class) IsTargetReachable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_is_target_reachable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the agent's navigation has finished. If the target is reachable, navigation ends when the target is reached. If the target is unreachable, navigation ends when the last waypoint of the path is reached.
[b]Note:[/b] While [code]true[/code] prefer to stop calling update functions like [method get_next_path_position]. This avoids jittering the standing agent due to calling repeated path updates.
*/
//go:nosplit
func (self class) IsNavigationFinished() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_is_navigation_finished, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the reachable final position of the current navigation path in global coordinates. This position can change if the agent needs to update the navigation path which makes the agent emit the [signal path_changed] signal.
*/
//go:nosplit
func (self class) GetFinalPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_final_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAvoidanceLayers(layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAvoidanceLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAvoidanceMask(mask gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAvoidanceMask() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member avoidance_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetAvoidanceLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_avoidance_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member avoidance_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetAvoidanceLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_avoidance_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified mask in the [member avoidance_mask] bitmask, given a [param mask_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetAvoidanceMaskValue(mask_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_avoidance_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified mask of the [member avoidance_mask] bitmask is enabled, given a [param mask_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetAvoidanceMaskValue(mask_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_avoidance_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAvoidancePriority(priority gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAvoidancePriority() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugUseCustom(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_debug_use_custom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugUseCustom() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_debug_use_custom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugPathCustomColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_debug_path_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugPathCustomColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_debug_path_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugPathCustomPointSize(point_size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_set_debug_path_custom_point_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugPathCustomPointSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_debug_path_custom_point_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsNavigationAgent3D() Expert { return self[0].AsNavigationAgent3D() }


//go:nosplit
func (self Simple) AsNavigationAgent3D() Simple { return self[0].AsNavigationAgent3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("NavigationAgent3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
