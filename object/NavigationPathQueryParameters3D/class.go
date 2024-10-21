package NavigationPathQueryParameters3D

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
By changing various properties of this object, such as the start and target position, you can configure path queries to the [NavigationServer3D].

*/
type Simple [1]classdb.NavigationPathQueryParameters3D
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
func (self Simple) SetMap(mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMap(mapping)
}
func (self Simple) GetMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetMap())
}
func (self Simple) SetStartPosition(start_position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStartPosition(start_position)
}
func (self Simple) GetStartPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetStartPosition())
}
func (self Simple) SetTargetPosition(target_position gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTargetPosition(target_position)
}
func (self Simple) GetTargetPosition() gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetTargetPosition())
}
func (self Simple) SetNavigationLayers(navigation_layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationLayers(gd.Int(navigation_layers))
}
func (self Simple) GetNavigationLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNavigationLayers()))
}
func (self Simple) SetMetadataFlags(flags classdb.NavigationPathQueryParameters3DPathMetadataFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMetadataFlags(flags)
}
func (self Simple) GetMetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NavigationPathQueryParameters3DPathMetadataFlags(Expert(self).GetMetadataFlags())
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
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationPathQueryParameters3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm classdb.NavigationPathQueryParameters3DPathfindingAlgorithm)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathfindingAlgorithm() classdb.NavigationPathQueryParameters3DPathfindingAlgorithm {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathfindingAlgorithm](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPathPostprocessing() classdb.NavigationPathQueryParameters3DPathPostProcessing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathPostProcessing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMap(mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStartPosition(start_position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, start_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStartPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetPosition(target_position gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, target_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetPosition() gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetadataFlags(flags classdb.NavigationPathQueryParameters3DPathMetadataFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathMetadataFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimplifyPath() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSimplifyEpsilon() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationPathQueryParameters3D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsNavigationPathQueryParameters3D() Expert { return self[0].AsNavigationPathQueryParameters3D() }


//go:nosplit
func (self Simple) AsNavigationPathQueryParameters3D() Simple { return self[0].AsNavigationPathQueryParameters3D() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

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
func init() {classdb.Register("NavigationPathQueryParameters3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type PathfindingAlgorithm = classdb.NavigationPathQueryParameters3DPathfindingAlgorithm

const (
/*The path query uses the default A* pathfinding algorithm.*/
	PathfindingAlgorithmAstar PathfindingAlgorithm = 0
)
type PathPostProcessing = classdb.NavigationPathQueryParameters3DPathPostProcessing

const (
/*Applies a funnel algorithm to the raw path corridor found by the pathfinding algorithm. This will result in the shortest path possible inside the path corridor. This postprocessing very much depends on the navigation mesh polygon layout and the created corridor. Especially tile- or gridbased layouts can face artificial corners with diagonal movement due to a jagged path corridor imposed by the cell shapes.*/
	PathPostprocessingCorridorfunnel PathPostProcessing = 0
/*Centers every path position in the middle of the traveled navigation mesh polygon edge. This creates better paths for tile- or gridbased layouts that restrict the movement to the cells center.*/
	PathPostprocessingEdgecentered PathPostProcessing = 1
)
type PathMetadataFlags = classdb.NavigationPathQueryParameters3DPathMetadataFlags

const (
/*Don't include any additional metadata about the returned path.*/
	PathMetadataIncludeNone PathMetadataFlags = 0
/*Include the type of navigation primitive (region or link) that each point of the path goes through.*/
	PathMetadataIncludeTypes PathMetadataFlags = 1
/*Include the [RID]s of the regions and links that each point of the path goes through.*/
	PathMetadataIncludeRids PathMetadataFlags = 2
/*Include the [code]ObjectID[/code]s of the [Object]s which manage the regions and links each point of the path goes through.*/
	PathMetadataIncludeOwners PathMetadataFlags = 4
/*Include all available metadata about the returned path.*/
	PathMetadataIncludeAll PathMetadataFlags = 7
)
