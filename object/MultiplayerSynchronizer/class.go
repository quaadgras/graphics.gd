package MultiplayerSynchronizer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
By default, [MultiplayerSynchronizer] synchronizes configured properties to all peers.
Visibility can be handled directly with [method set_visibility_for] or as-needed with [method add_visibility_filter] and [method update_visibility].
[MultiplayerSpawner]s will handle nodes according to visibility of synchronizers as long as the node at [member root_path] was spawned by one.
Internally, [MultiplayerSynchronizer] uses [method MultiplayerAPI.object_configuration_add] to notify synchronization start passing the [Node] at [member root_path] as the [code]object[/code] and itself as the [code]configuration[/code], and uses [method MultiplayerAPI.object_configuration_remove] to notify synchronization end in a similar way.
[b]Note:[/b] Synchronization is not supported for [Object] type properties, like [Resource]. Properties that are unique to each peer, like the instance IDs of [Object]s (see [method Object.get_instance_id]) or [RID]s, will also not work in synchronization.

*/
type Simple [1]classdb.MultiplayerSynchronizer
func (self Simple) SetRootPath(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRootPath(path)
}
func (self Simple) GetRootPath() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetRootPath(gc))
}
func (self Simple) SetReplicationInterval(milliseconds float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReplicationInterval(gd.Float(milliseconds))
}
func (self Simple) GetReplicationInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetReplicationInterval()))
}
func (self Simple) SetDeltaInterval(milliseconds float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeltaInterval(gd.Float(milliseconds))
}
func (self Simple) GetDeltaInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetDeltaInterval()))
}
func (self Simple) SetReplicationConfig(config [1]classdb.SceneReplicationConfig) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetReplicationConfig(config)
}
func (self Simple) GetReplicationConfig() [1]classdb.SceneReplicationConfig {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SceneReplicationConfig(Expert(self).GetReplicationConfig(gc))
}
func (self Simple) SetVisibilityUpdateMode(mode classdb.MultiplayerSynchronizerVisibilityUpdateMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibilityUpdateMode(mode)
}
func (self Simple) GetVisibilityUpdateMode() classdb.MultiplayerSynchronizerVisibilityUpdateMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiplayerSynchronizerVisibilityUpdateMode(Expert(self).GetVisibilityUpdateMode())
}
func (self Simple) UpdateVisibility(for_peer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateVisibility(gd.Int(for_peer))
}
func (self Simple) SetVisibilityPublic(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibilityPublic(visible)
}
func (self Simple) IsVisibilityPublic() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVisibilityPublic())
}
func (self Simple) AddVisibilityFilter(filter gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddVisibilityFilter(filter)
}
func (self Simple) RemoveVisibilityFilter(filter gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveVisibilityFilter(filter)
}
func (self Simple) SetVisibilityFor(peer int, visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibilityFor(gd.Int(peer), visible)
}
func (self Simple) GetVisibilityFor(peer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetVisibilityFor(gd.Int(peer)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MultiplayerSynchronizer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) SetReplicationConfig(config object.SceneReplicationConfig)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(config[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_set_replication_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReplicationConfig(ctx gd.Lifetime) object.SceneReplicationConfig {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerSynchronizer.Bind_get_replication_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SceneReplicationConfig
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

//go:nosplit
func (self class) AsMultiplayerSynchronizer() Expert { return self[0].AsMultiplayerSynchronizer() }


//go:nosplit
func (self Simple) AsMultiplayerSynchronizer() Simple { return self[0].AsMultiplayerSynchronizer() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MultiplayerSynchronizer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type VisibilityUpdateMode = classdb.MultiplayerSynchronizerVisibilityUpdateMode

const (
/*Visibility filters are updated during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	VisibilityProcessIdle VisibilityUpdateMode = 0
/*Visibility filters are updated during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	VisibilityProcessPhysics VisibilityUpdateMode = 1
/*Visibility filters are not updated automatically, and must be updated manually by calling [method update_visibility].*/
	VisibilityProcessNone VisibilityUpdateMode = 2
)
