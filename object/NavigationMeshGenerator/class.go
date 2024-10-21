package NavigationMeshGenerator

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
This class is responsible for creating and clearing 3D navigation meshes used as [NavigationMesh] resources inside [NavigationRegion3D]. The [NavigationMeshGenerator] has very limited to no use for 2D as the navigation mesh baking process expects 3D node types and 3D source geometry to parse.
The entire navigation mesh baking is best done in a separate thread as the voxelization, collision tests and mesh optimization steps involved are very slow and performance-intensive operations.
Navigation mesh baking happens in multiple steps and the result depends on 3D source geometry and properties of the [NavigationMesh] resource. In the first step, starting from a root node and depending on [NavigationMesh] properties all valid 3D source geometry nodes are collected from the [SceneTree]. Second, all collected nodes are parsed for their relevant 3D geometry data and a combined 3D mesh is build. Due to the many different types of parsable objects, from normal [MeshInstance3D]s to [CSGShape3D]s or various [CollisionObject3D]s, some operations to collect geometry data can trigger [RenderingServer] and [PhysicsServer3D] synchronizations. Server synchronization can have a negative effect on baking time or framerate as it often involves [Mutex] locking for thread security. Many parsable objects and the continuous synchronization with other threaded Servers can increase the baking time significantly. On the other hand only a few but very large and complex objects will take some time to prepare for the Servers which can noticeably stall the next frame render. As a general rule the total number of parsable objects and their individual size and complexity should be balanced to avoid framerate issues or very long baking times. The combined mesh is then passed to the Recast Navigation Object to test the source geometry for walkable terrain suitable to [NavigationMesh] agent properties by creating a voxel world around the meshes bounding area.
The finalized navigation mesh is then returned and stored inside the [NavigationMesh] for use as a resource inside [NavigationRegion3D] nodes.
[b]Note:[/b] Using meshes to not only define walkable surfaces but also obstruct navigation baking does not always work. The navigation baking has no concept of what is a geometry "inside" when dealing with mesh source geometry and this is intentional. Depending on current baking parameters, as soon as the obstructing mesh is large enough to fit a navigation mesh area inside, the baking will generate navigation mesh areas that are inside the obstructing source geometry mesh.

*/
type Simple [1]classdb.NavigationMeshGenerator
func (self Simple) Bake(navigation_mesh [1]classdb.NavigationMesh, root_node [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Bake(navigation_mesh, root_node)
}
func (self Simple) Clear(navigation_mesh [1]classdb.NavigationMesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear(navigation_mesh)
}
func (self Simple) ParseSourceGeometryData(navigation_mesh [1]classdb.NavigationMesh, source_geometry_data [1]classdb.NavigationMeshSourceGeometryData3D, root_node [1]classdb.Node, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParseSourceGeometryData(navigation_mesh, source_geometry_data, root_node, callback)
}
func (self Simple) BakeFromSourceGeometryData(navigation_mesh [1]classdb.NavigationMesh, source_geometry_data [1]classdb.NavigationMeshSourceGeometryData3D, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BakeFromSourceGeometryData(navigation_mesh, source_geometry_data, callback)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationMeshGenerator
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Bakes the [param navigation_mesh] with source geometry collected starting from the [param root_node].
*/
//go:nosplit
func (self class) Bake(navigation_mesh object.NavigationMesh, root_node object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(root_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshGenerator.Bind_bake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all polygons and vertices from the provided [param navigation_mesh] resource.
*/
//go:nosplit
func (self class) Clear(navigation_mesh object.NavigationMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshGenerator.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshGenerator.Bind_parse_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshGenerator.Bind_bake_from_source_geometry_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsNavigationMeshGenerator() Expert { return self[0].AsNavigationMeshGenerator() }


//go:nosplit
func (self Simple) AsNavigationMeshGenerator() Simple { return self[0].AsNavigationMeshGenerator() }

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
func init() {classdb.Register("NavigationMeshGenerator", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
