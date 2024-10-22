package NavigationAgent3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 3D agent used to pathfind to a position while avoiding static and dynamic obstacles. The calculation can be used by the parent node to dynamically move it along the path. Requires navigation data to work correctly.
Dynamic obstacles are avoided using RVO collision avoidance. Avoidance is computed before physics, so the pathfinding information can be used safely in the physics step.
[b]Note:[/b] After setting the [member target_position] property, the [method get_next_path_position] method must be used once every physics frame to update the internal path logic of the navigation agent. The vector position it returns should be used as the next movement position for the agent's parent node.

*/
type Go [1]classdb.NavigationAgent3D

/*
Returns the [RID] of this agent on the [NavigationServer3D].
*/
func (self Go) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Go) SetNavigationLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetNavigationLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Sets the [RID] of the navigation map this NavigationAgent node should use and also updates the [code]agent[/code] on the NavigationServer.
*/
func (self Go) SetNavigationMap(navigation_map gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationMap(navigation_map)
}

/*
Returns the [RID] of the navigation map for this NavigationAgent node. This function returns always the map set on the NavigationAgent node and not the map of the abstract agent on the NavigationServer. If the agent map is changed directly with the NavigationServer API the NavigationAgent node will not be aware of the map change. Use [method set_navigation_map] to change the navigation map for the NavigationAgent and also update the agent on the NavigationServer.
*/
func (self Go) GetNavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(class(self).GetNavigationMap())
}

/*
Returns the next position in global coordinates that can be moved to, making sure that there are no static objects in the way. If the agent does not have a navigation path, it will return the position of the agent's parent. The use of this function once every physics frame is required to update the internal path logic of the NavigationAgent.
*/
func (self Go) GetNextPathPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetNextPathPosition())
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
func (self Go) SetVelocityForced(velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVelocityForced(velocity)
}

/*
Returns the distance to the target position, using the agent's global position. The user must set [member target_position] in order for this to be accurate.
*/
func (self Go) DistanceToTarget() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).DistanceToTarget()))
}

/*
Returns the path query result for the path the agent is currently following.
*/
func (self Go) GetCurrentNavigationResult() gdclass.NavigationPathQueryResult3D {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.NavigationPathQueryResult3D(class(self).GetCurrentNavigationResult(gc))
}

/*
Returns this agent's current path from start to finish in global coordinates. The path only updates when the target position is changed or the agent requires a repath. The path array is not intended to be used in direct path movement as the agent has its own internal path logic that would get corrupted by changing the path array manually. Use the intended [method get_next_path_position] once every physics frame to receive the next path point for the agents movement as this function also updates the internal path logic.
*/
func (self Go) GetCurrentNavigationPath() []gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return []gd.Vector3(class(self).GetCurrentNavigationPath(gc).AsSlice())
}

/*
Returns which index the agent is currently on in the navigation path's [PackedVector3Array].
*/
func (self Go) GetCurrentNavigationPathIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCurrentNavigationPathIndex()))
}

/*
Returns [code]true[/code] if the agent reached the target, i.e. the agent moved within [member target_desired_distance] of the [member target_position]. It may not always be possible to reach the target but it should always be possible to reach the final position. See [method get_final_position].
*/
func (self Go) IsTargetReached() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTargetReached())
}

/*
Returns [code]true[/code] if [method get_final_position] is within [member target_desired_distance] of the [member target_position].
*/
func (self Go) IsTargetReachable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsTargetReachable())
}

/*
Returns [code]true[/code] if the agent's navigation has finished. If the target is reachable, navigation ends when the target is reached. If the target is unreachable, navigation ends when the last waypoint of the path is reached.
[b]Note:[/b] While [code]true[/code] prefer to stop calling update functions like [method get_next_path_position]. This avoids jittering the standing agent due to calling repeated path updates.
*/
func (self Go) IsNavigationFinished() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsNavigationFinished())
}

/*
Returns the reachable final position of the current navigation path in global coordinates. This position can change if the agent needs to update the navigation path which makes the agent emit the [signal path_changed] signal.
*/
func (self Go) GetFinalPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetFinalPosition())
}

/*
Based on [param value], enables or disables the specified layer in the [member avoidance_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Go) SetAvoidanceLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAvoidanceLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member avoidance_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetAvoidanceLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetAvoidanceLayerValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified mask in the [member avoidance_mask] bitmask, given a [param mask_number] between 1 and 32.
*/
func (self Go) SetAvoidanceMaskValue(mask_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAvoidanceMaskValue(gd.Int(mask_number), value)
}

/*
Returns whether or not the specified mask of the [member avoidance_mask] bitmask is enabled, given a [param mask_number] between 1 and 32.
*/
func (self Go) GetAvoidanceMaskValue(mask_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetAvoidanceMaskValue(gd.Int(mask_number)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationAgent3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("NavigationAgent3D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) TargetPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetTargetPosition())
}

func (self Go) SetTargetPosition(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetPosition(value)
}

func (self Go) PathDesiredDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPathDesiredDistance()))
}

func (self Go) SetPathDesiredDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathDesiredDistance(gd.Float(value))
}

func (self Go) TargetDesiredDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTargetDesiredDistance()))
}

func (self Go) SetTargetDesiredDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetDesiredDistance(gd.Float(value))
}

func (self Go) PathHeightOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPathHeightOffset()))
}

func (self Go) SetPathHeightOffset(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathHeightOffset(gd.Float(value))
}

func (self Go) PathMaxDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetPathMaxDistance()))
}

func (self Go) SetPathMaxDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathMaxDistance(gd.Float(value))
}

func (self Go) NavigationLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetNavigationLayers()))
}

func (self Go) SetNavigationLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Go) PathfindingAlgorithm() classdb.NavigationPathQueryParameters3DPathfindingAlgorithm {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPathQueryParameters3DPathfindingAlgorithm(class(self).GetPathfindingAlgorithm())
}

func (self Go) SetPathfindingAlgorithm(value classdb.NavigationPathQueryParameters3DPathfindingAlgorithm) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathfindingAlgorithm(value)
}

func (self Go) PathPostprocessing() classdb.NavigationPathQueryParameters3DPathPostProcessing {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPathQueryParameters3DPathPostProcessing(class(self).GetPathPostprocessing())
}

func (self Go) SetPathPostprocessing(value classdb.NavigationPathQueryParameters3DPathPostProcessing) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathPostprocessing(value)
}

func (self Go) PathMetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.NavigationPathQueryParameters3DPathMetadataFlags(class(self).GetPathMetadataFlags())
}

func (self Go) SetPathMetadataFlags(value classdb.NavigationPathQueryParameters3DPathMetadataFlags) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPathMetadataFlags(value)
}

func (self Go) SimplifyPath() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetSimplifyPath())
}

func (self Go) SetSimplifyPath(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSimplifyPath(value)
}

func (self Go) SimplifyEpsilon() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSimplifyEpsilon()))
}

func (self Go) SetSimplifyEpsilon(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSimplifyEpsilon(gd.Float(value))
}

func (self Go) AvoidanceEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetAvoidanceEnabled())
}

func (self Go) SetAvoidanceEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAvoidanceEnabled(value)
}

func (self Go) Velocity() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector3(class(self).GetVelocity())
}

func (self Go) SetVelocity(value gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVelocity(value)
}

func (self Go) Height() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetHeight()))
}

func (self Go) SetHeight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHeight(gd.Float(value))
}

func (self Go) Radius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadius(gd.Float(value))
}

func (self Go) NeighborDistance() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetNeighborDistance()))
}

func (self Go) SetNeighborDistance(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNeighborDistance(gd.Float(value))
}

func (self Go) MaxNeighbors() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxNeighbors()))
}

func (self Go) SetMaxNeighbors(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxNeighbors(gd.Int(value))
}

func (self Go) TimeHorizonAgents() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTimeHorizonAgents()))
}

func (self Go) SetTimeHorizonAgents(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTimeHorizonAgents(gd.Float(value))
}

func (self Go) TimeHorizonObstacles() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTimeHorizonObstacles()))
}

func (self Go) SetTimeHorizonObstacles(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTimeHorizonObstacles(gd.Float(value))
}

func (self Go) MaxSpeed() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMaxSpeed()))
}

func (self Go) SetMaxSpeed(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxSpeed(gd.Float(value))
}

func (self Go) Use3dAvoidance() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetUse3dAvoidance())
}

func (self Go) SetUse3dAvoidance(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUse3dAvoidance(value)
}

func (self Go) KeepYVelocity() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetKeepYVelocity())
}

func (self Go) SetKeepYVelocity(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeepYVelocity(value)
}

func (self Go) AvoidanceLayers() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetAvoidanceLayers()))
}

func (self Go) SetAvoidanceLayers(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAvoidanceLayers(gd.Int(value))
}

func (self Go) AvoidanceMask() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetAvoidanceMask()))
}

func (self Go) SetAvoidanceMask(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAvoidanceMask(gd.Int(value))
}

func (self Go) AvoidancePriority() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAvoidancePriority()))
}

func (self Go) SetAvoidancePriority(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAvoidancePriority(gd.Float(value))
}

func (self Go) DebugEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDebugEnabled())
}

func (self Go) SetDebugEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugEnabled(value)
}

func (self Go) DebugUseCustom() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDebugUseCustom())
}

func (self Go) SetDebugUseCustom(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugUseCustom(value)
}

func (self Go) DebugPathCustomColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetDebugPathCustomColor())
}

func (self Go) SetDebugPathCustomColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugPathCustomColor(value)
}

func (self Go) DebugPathCustomPointSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDebugPathCustomPointSize()))
}

func (self Go) SetDebugPathCustomPointSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDebugPathCustomPointSize(gd.Float(value))
}

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
func (self class) GetCurrentNavigationResult(ctx gd.Lifetime) gdclass.NavigationPathQueryResult3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationAgent3D.Bind_get_current_navigation_result, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.NavigationPathQueryResult3D
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
func (self Go) OnPathChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("path_changed"), gc.Callable(cb), 0)
}


func (self Go) OnTargetReached(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("target_reached"), gc.Callable(cb), 0)
}


func (self Go) OnWaypointReached(cb func(details gd.Dictionary)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("waypoint_reached"), gc.Callable(cb), 0)
}


func (self Go) OnLinkReached(cb func(details gd.Dictionary)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("link_reached"), gc.Callable(cb), 0)
}


func (self Go) OnNavigationFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("navigation_finished"), gc.Callable(cb), 0)
}


func (self Go) OnVelocityComputed(cb func(safe_velocity gd.Vector3)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("velocity_computed"), gc.Callable(cb), 0)
}


func (self class) AsNavigationAgent3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationAgent3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("NavigationAgent3D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
