package NavigationPathQueryParameters3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
By changing various properties of this object, such as the start and target position, you can configure path queries to the [NavigationServer3D].
*/
type Instance [1]classdb.NavigationPathQueryParameters3D
type Any interface {
	gd.IsClass
	AsNavigationPathQueryParameters3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.NavigationPathQueryParameters3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationPathQueryParameters3D"))
	return Instance{*(*classdb.NavigationPathQueryParameters3D)(unsafe.Pointer(&object))}
}

func (self Instance) Map() Resource.ID {
	return Resource.ID(class(self).GetMap())
}

func (self Instance) SetMap(value Resource.ID) {
	class(self).SetMap(value)
}

func (self Instance) StartPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetStartPosition())
}

func (self Instance) SetStartPosition(value Vector3.XYZ) {
	class(self).SetStartPosition(gd.Vector3(value))
}

func (self Instance) TargetPosition() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetTargetPosition())
}

func (self Instance) SetTargetPosition(value Vector3.XYZ) {
	class(self).SetTargetPosition(gd.Vector3(value))
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Instance) PathfindingAlgorithm() classdb.NavigationPathQueryParameters3DPathfindingAlgorithm {
	return classdb.NavigationPathQueryParameters3DPathfindingAlgorithm(class(self).GetPathfindingAlgorithm())
}

func (self Instance) SetPathfindingAlgorithm(value classdb.NavigationPathQueryParameters3DPathfindingAlgorithm) {
	class(self).SetPathfindingAlgorithm(value)
}

func (self Instance) PathPostprocessing() classdb.NavigationPathQueryParameters3DPathPostProcessing {
	return classdb.NavigationPathQueryParameters3DPathPostProcessing(class(self).GetPathPostprocessing())
}

func (self Instance) SetPathPostprocessing(value classdb.NavigationPathQueryParameters3DPathPostProcessing) {
	class(self).SetPathPostprocessing(value)
}

func (self Instance) MetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	return classdb.NavigationPathQueryParameters3DPathMetadataFlags(class(self).GetMetadataFlags())
}

func (self Instance) SetMetadataFlags(value classdb.NavigationPathQueryParameters3DPathMetadataFlags) {
	class(self).SetMetadataFlags(value)
}

func (self Instance) SimplifyPath() bool {
	return bool(class(self).GetSimplifyPath())
}

func (self Instance) SetSimplifyPath(value bool) {
	class(self).SetSimplifyPath(value)
}

func (self Instance) SimplifyEpsilon() Float.X {
	return Float.X(Float.X(class(self).GetSimplifyEpsilon()))
}

func (self Instance) SetSimplifyEpsilon(value Float.X) {
	class(self).SetSimplifyEpsilon(gd.Float(value))
}

//go:nosplit
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm classdb.NavigationPathQueryParameters3DPathfindingAlgorithm) {
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathfindingAlgorithm() classdb.NavigationPathQueryParameters3DPathfindingAlgorithm {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathfindingAlgorithm](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathPostprocessing(path_postprocessing classdb.NavigationPathQueryParameters3DPathPostProcessing) {
	var frame = callframe.New()
	callframe.Arg(frame, path_postprocessing)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathPostprocessing() classdb.NavigationPathQueryParameters3DPathPostProcessing {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathPostProcessing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMap(mapping gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMap() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartPosition(start_position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, start_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartPosition() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetPosition(target_position gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, target_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetPosition() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetadataFlags(flags classdb.NavigationPathQueryParameters3DPathMetadataFlags) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMetadataFlags() classdb.NavigationPathQueryParameters3DPathMetadataFlags {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NavigationPathQueryParameters3DPathMetadataFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyPath(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyPath() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyEpsilon(epsilon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, epsilon)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyEpsilon() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters3D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationPathQueryParameters3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationPathQueryParameters3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("NavigationPathQueryParameters3D", func(ptr gd.Object) any {
		return [1]classdb.NavigationPathQueryParameters3D{*(*classdb.NavigationPathQueryParameters3D)(unsafe.Pointer(&ptr))}
	})
}

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
