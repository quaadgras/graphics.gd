package MultiplayerSpawner

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
Spawnable scenes can be configured in the editor or through code (see [method add_spawnable_scene]).
Also supports custom node spawns through [method spawn], calling [member spawn_function] on all peers.
Internally, [MultiplayerSpawner] uses [method MultiplayerAPI.object_configuration_add] to notify spawns passing the spawned node as the [code]object[/code] and itself as the [code]configuration[/code], and [method MultiplayerAPI.object_configuration_remove] to notify despawns in a similar way.

*/
type Go [1]classdb.MultiplayerSpawner

/*
Adds a scene path to spawnable scenes, making it automatically replicated from the multiplayer authority to other peers when added as children of the node pointed by [member spawn_path].
*/
func (self Go) AddSpawnableScene(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddSpawnableScene(gc.String(path))
}

/*
Returns the count of spawnable scene paths.
*/
func (self Go) GetSpawnableSceneCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSpawnableSceneCount()))
}

/*
Returns the spawnable scene path by index.
*/
func (self Go) GetSpawnableScene(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSpawnableScene(gc, gd.Int(index)).String())
}

/*
Clears all spawnable scenes. Does not despawn existing instances on remote peers.
*/
func (self Go) ClearSpawnableScenes() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearSpawnableScenes()
}

/*
Requests a custom spawn, with [param data] passed to [member spawn_function] on all peers. Returns the locally spawned node instance already inside the scene tree, and added as a child of the node pointed by [member spawn_path].
[b]Note:[/b] Spawnable scenes are spawned automatically. [method spawn] is only needed for custom spawns.
*/
func (self Go) Spawn() gdclass.Node {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Node(class(self).Spawn(gc, gc.Variant(([1]gd.Variant{}[0]))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MultiplayerSpawner
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MultiplayerSpawner"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SpawnPath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSpawnPath(gc).String())
}

func (self Go) SetSpawnPath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpawnPath(gc.String(value).NodePath(gc))
}

func (self Go) SpawnLimit() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSpawnLimit()))
}

func (self Go) SetSpawnLimit(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpawnLimit(gd.Int(value))
}

func (self Go) SpawnFunction() gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Callable(class(self).GetSpawnFunction(gc))
}

func (self Go) SetSpawnFunction(value gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSpawnFunction(value)
}

/*
Adds a scene path to spawnable scenes, making it automatically replicated from the multiplayer authority to other peers when added as children of the node pointed by [member spawn_path].
*/
//go:nosplit
func (self class) AddSpawnableScene(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_add_spawnable_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the count of spawnable scene paths.
*/
//go:nosplit
func (self class) GetSpawnableSceneCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_get_spawnable_scene_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the spawnable scene path by index.
*/
//go:nosplit
func (self class) GetSpawnableScene(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_get_spawnable_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears all spawnable scenes. Does not despawn existing instances on remote peers.
*/
//go:nosplit
func (self class) ClearSpawnableScenes()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_clear_spawnable_scenes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Requests a custom spawn, with [param data] passed to [member spawn_function] on all peers. Returns the locally spawned node instance already inside the scene tree, and added as a child of the node pointed by [member spawn_path].
[b]Note:[/b] Spawnable scenes are spawned automatically. [method spawn] is only needed for custom spawns.
*/
//go:nosplit
func (self class) Spawn(ctx gd.Lifetime, data gd.Variant) gdclass.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_spawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Node
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSpawnPath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_get_spawn_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpawnPath(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_set_spawn_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpawnLimit() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_get_spawn_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpawnLimit(limit gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, limit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_set_spawn_limit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpawnFunction(ctx gd.Lifetime) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_get_spawn_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSpawnFunction(spawn_function gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(spawn_function))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSpawner.Bind_set_spawn_function, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnDespawned(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("despawned"), gc.Callable(cb), 0)
}


func (self Go) OnSpawned(cb func(node gdclass.Node)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("spawned"), gc.Callable(cb), 0)
}


func (self class) AsMultiplayerSpawner() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerSpawner() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("MultiplayerSpawner", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
