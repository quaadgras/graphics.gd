package MultiplayerSpawner

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Path"
import "grow.graphics/gd/variant/Callable"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Spawnable scenes can be configured in the editor or through code (see [method add_spawnable_scene]).
Also supports custom node spawns through [method spawn], calling [member spawn_function] on all peers.
Internally, [MultiplayerSpawner] uses [method MultiplayerAPI.object_configuration_add] to notify spawns passing the spawned node as the [code]object[/code] and itself as the [code]configuration[/code], and [method MultiplayerAPI.object_configuration_remove] to notify despawns in a similar way.
*/
type Instance [1]classdb.MultiplayerSpawner

/*
Adds a scene path to spawnable scenes, making it automatically replicated from the multiplayer authority to other peers when added as children of the node pointed by [member spawn_path].
*/
func (self Instance) AddSpawnableScene(path string) {
	class(self).AddSpawnableScene(gd.NewString(path))
}

/*
Returns the count of spawnable scene paths.
*/
func (self Instance) GetSpawnableSceneCount() int {
	return int(int(class(self).GetSpawnableSceneCount()))
}

/*
Returns the spawnable scene path by index.
*/
func (self Instance) GetSpawnableScene(index int) string {
	return string(class(self).GetSpawnableScene(gd.Int(index)).String())
}

/*
Clears all spawnable scenes. Does not despawn existing instances on remote peers.
*/
func (self Instance) ClearSpawnableScenes() {
	class(self).ClearSpawnableScenes()
}

/*
Requests a custom spawn, with [param data] passed to [member spawn_function] on all peers. Returns the locally spawned node instance already inside the scene tree, and added as a child of the node pointed by [member spawn_path].
[b]Note:[/b] Spawnable scenes are spawned automatically. [method spawn] is only needed for custom spawns.
*/
func (self Instance) Spawn() objects.Node {
	return objects.Node(class(self).Spawn(gd.NewVariant(gd.NewVariant(([1]any{}[0])))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.MultiplayerSpawner

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerSpawner"))
	return Instance{classdb.MultiplayerSpawner(object)}
}

func (self Instance) SpawnPath() Path.String {
	return Path.String(class(self).GetSpawnPath().String())
}

func (self Instance) SetSpawnPath(value Path.String) {
	class(self).SetSpawnPath(gd.NewString(string(value)).NodePath())
}

func (self Instance) SpawnLimit() int {
	return int(int(class(self).GetSpawnLimit()))
}

func (self Instance) SetSpawnLimit(value int) {
	class(self).SetSpawnLimit(gd.Int(value))
}

func (self Instance) SpawnFunction() Callable.Any {
	return Callable.Any(class(self).GetSpawnFunction())
}

func (self Instance) SetSpawnFunction(value Callable.Any) {
	class(self).SetSpawnFunction(gd.NewCallable(value))
}

/*
Adds a scene path to spawnable scenes, making it automatically replicated from the multiplayer authority to other peers when added as children of the node pointed by [member spawn_path].
*/
//go:nosplit
func (self class) AddSpawnableScene(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_add_spawnable_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the count of spawnable scene paths.
*/
//go:nosplit
func (self class) GetSpawnableSceneCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_get_spawnable_scene_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the spawnable scene path by index.
*/
//go:nosplit
func (self class) GetSpawnableScene(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_get_spawnable_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears all spawnable scenes. Does not despawn existing instances on remote peers.
*/
//go:nosplit
func (self class) ClearSpawnableScenes() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_clear_spawnable_scenes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Requests a custom spawn, with [param data] passed to [member spawn_function] on all peers. Returns the locally spawned node instance already inside the scene tree, and added as a child of the node pointed by [member spawn_path].
[b]Note:[/b] Spawnable scenes are spawned automatically. [method spawn] is only needed for custom spawns.
*/
//go:nosplit
func (self class) Spawn(data gd.Variant) objects.Node {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_spawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Node{classdb.Node(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSpawnPath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_get_spawn_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpawnPath(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_set_spawn_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpawnLimit() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_get_spawn_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpawnLimit(limit gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_set_spawn_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpawnFunction() gd.Callable {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_get_spawn_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpawnFunction(spawn_function gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(spawn_function))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSpawner.Bind_set_spawn_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnDespawned(cb func(node objects.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("despawned"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSpawned(cb func(node objects.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("spawned"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiplayerSpawner() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerSpawner() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	classdb.Register("MultiplayerSpawner", func(ptr gd.Object) any { return classdb.MultiplayerSpawner(ptr) })
}
