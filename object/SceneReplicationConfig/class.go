package SceneReplicationConfig

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Simple [1]classdb.SceneReplicationConfig
func (self Simple) GetProperties() gd.ArrayOf[gd.NodePath] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.NodePath](Expert(self).GetProperties(gc))
}
func (self Simple) AddProperty(path gd.NodePath, index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddProperty(path, gd.Int(index))
}
func (self Simple) HasProperty(path gd.NodePath) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasProperty(path))
}
func (self Simple) RemoveProperty(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveProperty(path)
}
func (self Simple) PropertyGetIndex(path gd.NodePath) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).PropertyGetIndex(path)))
}
func (self Simple) PropertyGetSpawn(path gd.NodePath) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).PropertyGetSpawn(path))
}
func (self Simple) PropertySetSpawn(path gd.NodePath, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PropertySetSpawn(path, enabled)
}
func (self Simple) PropertyGetReplicationMode(path gd.NodePath) classdb.SceneReplicationConfigReplicationMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SceneReplicationConfigReplicationMode(Expert(self).PropertyGetReplicationMode(path))
}
func (self Simple) PropertySetReplicationMode(path gd.NodePath, mode classdb.SceneReplicationConfigReplicationMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PropertySetReplicationMode(path, mode)
}
func (self Simple) PropertyGetSync(path gd.NodePath) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).PropertyGetSync(path))
}
func (self Simple) PropertySetSync(path gd.NodePath, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PropertySetSync(path, enabled)
}
func (self Simple) PropertyGetWatch(path gd.NodePath) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).PropertyGetWatch(path))
}
func (self Simple) PropertySetWatch(path gd.NodePath, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PropertySetWatch(path, enabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SceneReplicationConfig
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns a list of synchronized property [NodePath]s.
*/
//go:nosplit
func (self class) GetProperties(ctx gd.Lifetime) gd.ArrayOf[gd.NodePath] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_get_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.NodePath](ret)
}
/*
Adds the property identified by the given [param path] to the list of the properties being synchronized, optionally passing an [param index].
[b]Note:[/b] For details on restrictions and limitations on property synchronization, see [MultiplayerSynchronizer].
*/
//go:nosplit
func (self class) AddProperty(path gd.NodePath, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_add_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param path] is configured for synchronization.
*/
//go:nosplit
func (self class) HasProperty(path gd.NodePath) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_has_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the property identified by the given [param path] from the configuration.
*/
//go:nosplit
func (self class) RemoveProperty(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_remove_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Finds the index of the given [param path].
*/
//go:nosplit
func (self class) PropertyGetIndex(path gd.NodePath) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on spawn.
*/
//go:nosplit
func (self class) PropertyGetSpawn(path gd.NodePath) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_get_spawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the property identified by the given [param path] is configured to be synchronized on spawn.
*/
//go:nosplit
func (self class) PropertySetSpawn(path gd.NodePath, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_set_spawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the replication mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
//go:nosplit
func (self class) PropertyGetReplicationMode(path gd.NodePath) classdb.SceneReplicationConfigReplicationMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[classdb.SceneReplicationConfigReplicationMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_get_replication_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the synchronization mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
//go:nosplit
func (self class) PropertySetReplicationMode(path gd.NodePath, mode classdb.SceneReplicationConfigReplicationMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_set_replication_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on process.
*/
//go:nosplit
func (self class) PropertyGetSync(path gd.NodePath) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_get_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the property identified by the given [param path] is configured to be synchronized on process.
*/
//go:nosplit
func (self class) PropertySetSync(path gd.NodePath, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_set_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
//go:nosplit
func (self class) PropertyGetWatch(path gd.NodePath) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_get_watch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
//go:nosplit
func (self class) PropertySetWatch(path gd.NodePath, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneReplicationConfig.Bind_property_set_watch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsSceneReplicationConfig() Expert { return self[0].AsSceneReplicationConfig() }


//go:nosplit
func (self Simple) AsSceneReplicationConfig() Simple { return self[0].AsSceneReplicationConfig() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


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
func init() {classdb.Register("SceneReplicationConfig", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type ReplicationMode = classdb.SceneReplicationConfigReplicationMode

const (
/*Do not keep the given property synchronized.*/
	ReplicationModeNever ReplicationMode = 0
/*Replicate the given property on process by constantly sending updates using unreliable transfer mode.*/
	ReplicationModeAlways ReplicationMode = 1
/*Replicate the given property on process by sending updates using reliable transfer mode when its value changes.*/
	ReplicationModeOnChange ReplicationMode = 2
)
