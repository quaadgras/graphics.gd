package SceneReplicationConfig

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Go [1]classdb.SceneReplicationConfig

/*
Returns a list of synchronized property [NodePath]s.
*/
func (self Go) GetProperties() gd.ArrayOf[gd.NodePath] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.NodePath](class(self).GetProperties(gc))
}

/*
Adds the property identified by the given [param path] to the list of the properties being synchronized, optionally passing an [param index].
[b]Note:[/b] For details on restrictions and limitations on property synchronization, see [MultiplayerSynchronizer].
*/
func (self Go) AddProperty(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddProperty(gc.String(path).NodePath(gc), gd.Int(-1))
}

/*
Returns [code]true[/code] if the given [param path] is configured for synchronization.
*/
func (self Go) HasProperty(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasProperty(gc.String(path).NodePath(gc)))
}

/*
Removes the property identified by the given [param path] from the configuration.
*/
func (self Go) RemoveProperty(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveProperty(gc.String(path).NodePath(gc))
}

/*
Finds the index of the given [param path].
*/
func (self Go) PropertyGetIndex(path string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).PropertyGetIndex(gc.String(path).NodePath(gc))))
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on spawn.
*/
func (self Go) PropertyGetSpawn(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).PropertyGetSpawn(gc.String(path).NodePath(gc)))
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on spawn.
*/
func (self Go) PropertySetSpawn(path string, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PropertySetSpawn(gc.String(path).NodePath(gc), enabled)
}

/*
Returns the replication mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
func (self Go) PropertyGetReplicationMode(path string) classdb.SceneReplicationConfigReplicationMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.SceneReplicationConfigReplicationMode(class(self).PropertyGetReplicationMode(gc.String(path).NodePath(gc)))
}

/*
Sets the synchronization mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
func (self Go) PropertySetReplicationMode(path string, mode classdb.SceneReplicationConfigReplicationMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PropertySetReplicationMode(gc.String(path).NodePath(gc), mode)
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on process.
*/
func (self Go) PropertyGetSync(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).PropertyGetSync(gc.String(path).NodePath(gc)))
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on process.
*/
func (self Go) PropertySetSync(path string, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PropertySetSync(gc.String(path).NodePath(gc), enabled)
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
func (self Go) PropertyGetWatch(path string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).PropertyGetWatch(gc.String(path).NodePath(gc)))
}

/*
Sets whether the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
func (self Go) PropertySetWatch(path string, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PropertySetWatch(gc.String(path).NodePath(gc), enabled)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SceneReplicationConfig
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SceneReplicationConfig"))
	return *(*Go)(unsafe.Pointer(&object))
}

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
func (self class) AsSceneReplicationConfig() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSceneReplicationConfig() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("SceneReplicationConfig", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ReplicationMode = classdb.SceneReplicationConfigReplicationMode

const (
/*Do not keep the given property synchronized.*/
	ReplicationModeNever ReplicationMode = 0
/*Replicate the given property on process by constantly sending updates using unreliable transfer mode.*/
	ReplicationModeAlways ReplicationMode = 1
/*Replicate the given property on process by sending updates using reliable transfer mode when its value changes.*/
	ReplicationModeOnChange ReplicationMode = 2
)
