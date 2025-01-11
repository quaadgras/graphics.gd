// Package NavigationMeshGenerator provides methods for working with NavigationMeshGenerator object instances.
package NavigationMeshGenerator

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
This class is responsible for creating and clearing 3D navigation meshes used as [NavigationMesh] resources inside [NavigationRegion3D]. The [NavigationMeshGenerator] has very limited to no use for 2D as the navigation mesh baking process expects 3D node types and 3D source geometry to parse.
The entire navigation mesh baking is best done in a separate thread as the voxelization, collision tests and mesh optimization steps involved are very slow and performance-intensive operations.
Navigation mesh baking happens in multiple steps and the result depends on 3D source geometry and properties of the [NavigationMesh] resource. In the first step, starting from a root node and depending on [NavigationMesh] properties all valid 3D source geometry nodes are collected from the [SceneTree]. Second, all collected nodes are parsed for their relevant 3D geometry data and a combined 3D mesh is build. Due to the many different types of parsable objects, from normal [MeshInstance3D]s to [CSGShape3D]s or various [CollisionObject3D]s, some operations to collect geometry data can trigger [RenderingServer] and [PhysicsServer3D] synchronizations. Server synchronization can have a negative effect on baking time or framerate as it often involves [Mutex] locking for thread security. Many parsable objects and the continuous synchronization with other threaded Servers can increase the baking time significantly. On the other hand only a few but very large and complex objects will take some time to prepare for the Servers which can noticeably stall the next frame render. As a general rule the total number of parsable objects and their individual size and complexity should be balanced to avoid framerate issues or very long baking times. The combined mesh is then passed to the Recast Navigation Object to test the source geometry for walkable terrain suitable to [NavigationMesh] agent properties by creating a voxel world around the meshes bounding area.
The finalized navigation mesh is then returned and stored inside the [NavigationMesh] for use as a resource inside [NavigationRegion3D] nodes.
[b]Note:[/b] Using meshes to not only define walkable surfaces but also obstruct navigation baking does not always work. The navigation baking has no concept of what is a geometry "inside" when dealing with mesh source geometry and this is intentional. Depending on current baking parameters, as soon as the obstructing mesh is large enough to fit a navigation mesh area inside, the baking will generate navigation mesh areas that are inside the obstructing source geometry mesh.
*/
var self [1]gdclass.NavigationMeshGenerator
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.NavigationMeshGenerator)
	self = *(*[1]gdclass.NavigationMeshGenerator)(unsafe.Pointer(&obj))
}

/*
Bakes the [param navigation_mesh] with source geometry collected starting from the [param root_node].
*/
func Bake(navigation_mesh [1]gdclass.NavigationMesh, root_node [1]gdclass.Node) {
	once.Do(singleton)
	class(self).Bake(navigation_mesh, root_node)
}

/*
Removes all polygons and vertices from the provided [param navigation_mesh] resource.
*/
func Clear(navigation_mesh [1]gdclass.NavigationMesh) {
	once.Do(singleton)
	class(self).Clear(navigation_mesh)
}

/*
Parses the [SceneTree] for source geometry according to the properties of [param navigation_mesh]. Updates the provided [param source_geometry_data] resource with the resulting data. The resource can then be used to bake a navigation mesh with [method bake_from_source_geometry_data]. After the process is finished the optional [param callback] will be called.
[b]Note:[/b] This function needs to run on the main thread or with a deferred call as the SceneTree is not thread-safe.
[b]Performance:[/b] While convenient, reading data arrays from [Mesh] resources can affect the frame rate negatively. The data needs to be received from the GPU, stalling the [RenderingServer] in the process. For performance prefer the use of e.g. collision shapes or creating the data arrays entirely in code.
*/
func ParseSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, root_node [1]gdclass.Node) {
	once.Do(singleton)
	class(self).ParseSourceGeometryData(navigation_mesh, source_geometry_data, root_node, gd.NewCallable(nil))
}

/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
func BakeFromSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D) {
	once.Do(singleton)
	class(self).BakeFromSourceGeometryData(navigation_mesh, source_geometry_data, gd.NewCallable(nil))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.NavigationMeshGenerator

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Bakes the [param navigation_mesh] with source geometry collected starting from the [param root_node].
*/
//go:nosplit
func (self class) Bake(navigation_mesh [1]gdclass.NavigationMesh, root_node [1]gdclass.Node) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(root_node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshGenerator.Bind_bake, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all polygons and vertices from the provided [param navigation_mesh] resource.
*/
//go:nosplit
func (self class) Clear(navigation_mesh [1]gdclass.NavigationMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshGenerator.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Parses the [SceneTree] for source geometry according to the properties of [param navigation_mesh]. Updates the provided [param source_geometry_data] resource with the resulting data. The resource can then be used to bake a navigation mesh with [method bake_from_source_geometry_data]. After the process is finished the optional [param callback] will be called.
[b]Note:[/b] This function needs to run on the main thread or with a deferred call as the SceneTree is not thread-safe.
[b]Performance:[/b] While convenient, reading data arrays from [Mesh] resources can affect the frame rate negatively. The data needs to be received from the GPU, stalling the [RenderingServer] in the process. For performance prefer the use of e.g. collision shapes or creating the data arrays entirely in code.
*/
//go:nosplit
func (self class) ParseSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, root_node [1]gdclass.Node, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(root_node[0])[0])
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshGenerator.Bind_parse_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Bakes the provided [param navigation_mesh] with the data from the provided [param source_geometry_data]. After the process is finished the optional [param callback] will be called.
*/
//go:nosplit
func (self class) BakeFromSourceGeometryData(navigation_mesh [1]gdclass.NavigationMesh, source_geometry_data [1]gdclass.NavigationMeshSourceGeometryData3D, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	callframe.Arg(frame, pointers.Get(source_geometry_data[0])[0])
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshGenerator.Bind_bake_from_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("NavigationMeshGenerator", func(ptr gd.Object) any {
		return [1]gdclass.NavigationMeshGenerator{*(*gdclass.NavigationMeshGenerator)(unsafe.Pointer(&ptr))}
	})
}
