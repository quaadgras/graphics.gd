package MultiplayerSynchronizer

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
By default, [MultiplayerSynchronizer] synchronizes configured properties to all peers.
Visibility can be handled directly with [method set_visibility_for] or as-needed with [method add_visibility_filter] and [method update_visibility].
[MultiplayerSpawner]s will handle nodes according to visibility of synchronizers as long as the node at [member root_path] was spawned by one.
Internally, [MultiplayerSynchronizer] uses [method MultiplayerAPI.object_configuration_add] to notify synchronization start passing the [Node] at [member root_path] as the [code]object[/code] and itself as the [code]configuration[/code], and uses [method MultiplayerAPI.object_configuration_remove] to notify synchronization end in a similar way.
[b]Note:[/b] Synchronization is not supported for [Object] type properties, like [Resource]. Properties that are unique to each peer, like the instance IDs of [Object]s (see [method Object.get_instance_id]) or [RID]s, will also not work in synchronization.

*/
type Go [1]classdb.MultiplayerSynchronizer

/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
func (self Go) UpdateVisibility() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UpdateVisibility(gd.Int(0))
}

/*
Adds a peer visibility filter for this synchronizer.
[param filter] should take a peer ID [int] and return a [bool].
*/
func (self Go) AddVisibilityFilter(filter gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddVisibilityFilter(filter)
}

/*
Removes a peer visibility filter from this synchronizer.
*/
func (self Go) RemoveVisibilityFilter(filter gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveVisibilityFilter(filter)
}

/*
Sets the visibility of [param peer] to [param visible]. If [param peer] is [code]0[/code], the value of [member public_visibility] will be updated instead.
*/
func (self Go) SetVisibilityFor(peer int, visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityFor(gd.Int(peer), visible)
}

/*
Queries the current visibility for peer [param peer].
*/
func (self Go) GetVisibilityFor(peer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetVisibilityFor(gd.Int(peer)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MultiplayerSynchronizer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MultiplayerSynchronizer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RootPath() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetRootPath(gc).String())
}

func (self Go) SetRootPath(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRootPath(gc.String(value).NodePath(gc))
}

func (self Go) ReplicationInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetReplicationInterval()))
}

func (self Go) SetReplicationInterval(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReplicationInterval(gd.Float(value))
}

func (self Go) DeltaInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetDeltaInterval()))
}

func (self Go) SetDeltaInterval(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDeltaInterval(gd.Float(value))
}

func (self Go) ReplicationConfig() gdclass.SceneReplicationConfig {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.SceneReplicationConfig(class(self).GetReplicationConfig(gc))
}

func (self Go) SetReplicationConfig(value gdclass.SceneReplicationConfig) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReplicationConfig(value)
}

func (self Go) VisibilityUpdateMode() classdb.MultiplayerSynchronizerVisibilityUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.MultiplayerSynchronizerVisibilityUpdateMode(class(self).GetVisibilityUpdateMode())
}

func (self Go) SetVisibilityUpdateMode(value classdb.MultiplayerSynchronizerVisibilityUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityUpdateMode(value)
}

func (self Go) PublicVisibility() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVisibilityPublic())
}

func (self Go) SetPublicVisibility(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibilityPublic(value)
}

//go:nosplit
func (self class) SetRootPath(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootPath(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReplicationInterval(milliseconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, milliseconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_replication_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReplicationInterval() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_replication_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeltaInterval(milliseconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, milliseconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_delta_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDeltaInterval() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_delta_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReplicationConfig(config gdclass.SceneReplicationConfig)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(config[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_replication_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReplicationConfig(ctx gd.Lifetime) gdclass.SceneReplicationConfig {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_replication_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.SceneReplicationConfig
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibilityUpdateMode(mode classdb.MultiplayerSynchronizerVisibilityUpdateMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_visibility_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibilityUpdateMode() classdb.MultiplayerSynchronizerVisibilityUpdateMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerSynchronizerVisibilityUpdateMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_visibility_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
//go:nosplit
func (self class) UpdateVisibility(for_peer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, for_peer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_update_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVisibilityPublic(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_visibility_public, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisibilityPublic() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_is_visibility_public, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a peer visibility filter for this synchronizer.
[param filter] should take a peer ID [int] and return a [bool].
*/
//go:nosplit
func (self class) AddVisibilityFilter(filter gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filter))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_add_visibility_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a peer visibility filter from this synchronizer.
*/
//go:nosplit
func (self class) RemoveVisibilityFilter(filter gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(filter))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_remove_visibility_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the visibility of [param peer] to [param visible]. If [param peer] is [code]0[/code], the value of [member public_visibility] will be updated instead.
*/
//go:nosplit
func (self class) SetVisibilityFor(peer gd.Int, visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_visibility_for, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Queries the current visibility for peer [param peer].
*/
//go:nosplit
func (self class) GetVisibilityFor(peer gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_visibility_for, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnSynchronized(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("synchronized"), gc.Callable(cb), 0)
}


func (self Go) OnDeltaSynchronized(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("delta_synchronized"), gc.Callable(cb), 0)
}


func (self Go) OnVisibilityChanged(cb func(for_peer int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("visibility_changed"), gc.Callable(cb), 0)
}


func (self class) AsMultiplayerSynchronizer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerSynchronizer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("MultiplayerSynchronizer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type VisibilityUpdateMode = classdb.MultiplayerSynchronizerVisibilityUpdateMode

const (
/*Visibility filters are updated during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	VisibilityProcessIdle VisibilityUpdateMode = 0
/*Visibility filters are updated during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	VisibilityProcessPhysics VisibilityUpdateMode = 1
/*Visibility filters are not updated automatically, and must be updated manually by calling [method update_visibility].*/
	VisibilityProcessNone VisibilityUpdateMode = 2
)
