package NavigationPathQueryParameters2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
By changing various properties of this object, such as the start and target position, you can configure path queries to the [NavigationServer2D].
*/
type Instance [1]gdclass.NavigationPathQueryParameters2D
type Any interface {
	gd.IsClass
	AsNavigationPathQueryParameters2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationPathQueryParameters2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationPathQueryParameters2D"))
	return Instance{*(*gdclass.NavigationPathQueryParameters2D)(unsafe.Pointer(&object))}
}

func (self Instance) Map() Resource.ID {
	return Resource.ID(class(self).GetMap())
}

func (self Instance) SetMap(value Resource.ID) {
	class(self).SetMap(value)
}

func (self Instance) StartPosition() Vector2.XY {
	return Vector2.XY(class(self).GetStartPosition())
}

func (self Instance) SetStartPosition(value Vector2.XY) {
	class(self).SetStartPosition(gd.Vector2(value))
}

func (self Instance) TargetPosition() Vector2.XY {
	return Vector2.XY(class(self).GetTargetPosition())
}

func (self Instance) SetTargetPosition(value Vector2.XY) {
	class(self).SetTargetPosition(gd.Vector2(value))
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Instance) PathfindingAlgorithm() gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm {
	return gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm(class(self).GetPathfindingAlgorithm())
}

func (self Instance) SetPathfindingAlgorithm(value gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm) {
	class(self).SetPathfindingAlgorithm(value)
}

func (self Instance) PathPostprocessing() gdclass.NavigationPathQueryParameters2DPathPostProcessing {
	return gdclass.NavigationPathQueryParameters2DPathPostProcessing(class(self).GetPathPostprocessing())
}

func (self Instance) SetPathPostprocessing(value gdclass.NavigationPathQueryParameters2DPathPostProcessing) {
	class(self).SetPathPostprocessing(value)
}

func (self Instance) MetadataFlags() gdclass.NavigationPathQueryParameters2DPathMetadataFlags {
	return gdclass.NavigationPathQueryParameters2DPathMetadataFlags(class(self).GetMetadataFlags())
}

func (self Instance) SetMetadataFlags(value gdclass.NavigationPathQueryParameters2DPathMetadataFlags) {
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
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm) {
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathfindingAlgorithm() gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathPostprocessing(path_postprocessing gdclass.NavigationPathQueryParameters2DPathPostProcessing) {
	var frame = callframe.New()
	callframe.Arg(frame, path_postprocessing)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathPostprocessing() gdclass.NavigationPathQueryParameters2DPathPostProcessing {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters2DPathPostProcessing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMap(mapping gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMap() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStartPosition(start_position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, start_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStartPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetPosition(target_position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, target_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetadataFlags(flags gdclass.NavigationPathQueryParameters2DPathMetadataFlags) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMetadataFlags() gdclass.NavigationPathQueryParameters2DPathMetadataFlags {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters2DPathMetadataFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyPath(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyPath() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyEpsilon(epsilon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, epsilon)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyEpsilon() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryParameters2D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationPathQueryParameters2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationPathQueryParameters2D() Instance {
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
	gdclass.Register("NavigationPathQueryParameters2D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationPathQueryParameters2D{*(*gdclass.NavigationPathQueryParameters2D)(unsafe.Pointer(&ptr))}
	})
}

type PathfindingAlgorithm = gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm

const (
	/*The path query uses the default A* pathfinding algorithm.*/
	PathfindingAlgorithmAstar PathfindingAlgorithm = 0
)

type PathPostProcessing = gdclass.NavigationPathQueryParameters2DPathPostProcessing

const (
	/*Applies a funnel algorithm to the raw path corridor found by the pathfinding algorithm. This will result in the shortest path possible inside the path corridor. This postprocessing very much depends on the navigation mesh polygon layout and the created corridor. Especially tile- or gridbased layouts can face artificial corners with diagonal movement due to a jagged path corridor imposed by the cell shapes.*/
	PathPostprocessingCorridorfunnel PathPostProcessing = 0
	/*Centers every path position in the middle of the traveled navigation mesh polygon edge. This creates better paths for tile- or gridbased layouts that restrict the movement to the cells center.*/
	PathPostprocessingEdgecentered PathPostProcessing = 1
)

type PathMetadataFlags = gdclass.NavigationPathQueryParameters2DPathMetadataFlags

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
