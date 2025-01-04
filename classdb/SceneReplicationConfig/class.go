// Package SceneReplicationConfig provides methods for working with SceneReplicationConfig object instances.
package SceneReplicationConfig

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Path"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.SceneReplicationConfig
type Any interface {
	gd.IsClass
	AsSceneReplicationConfig() Instance
}

/*
Returns a list of synchronized property [NodePath]s.
*/
func (self Instance) GetProperties() gd.Array {
	return gd.Array(class(self).GetProperties())
}

/*
Adds the property identified by the given [param path] to the list of the properties being synchronized, optionally passing an [param index].
[b]Note:[/b] For details on restrictions and limitations on property synchronization, see [MultiplayerSynchronizer].
*/
func (self Instance) AddProperty(path Path.String) {
	class(self).AddProperty(gd.NewString(string(path)).NodePath(), gd.Int(-1))
}

/*
Returns [code]true[/code] if the given [param path] is configured for synchronization.
*/
func (self Instance) HasProperty(path Path.String) bool {
	return bool(class(self).HasProperty(gd.NewString(string(path)).NodePath()))
}

/*
Removes the property identified by the given [param path] from the configuration.
*/
func (self Instance) RemoveProperty(path Path.String) {
	class(self).RemoveProperty(gd.NewString(string(path)).NodePath())
}

/*
Finds the index of the given [param path].
*/
func (self Instance) PropertyGetIndex(path Path.String) int {
	return int(int(class(self).PropertyGetIndex(gd.NewString(string(path)).NodePath())))
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on spawn.
*/
func (self Instance) PropertyGetSpawn(path Path.String) bool {
	return bool(class(self).PropertyGetSpawn(gd.NewString(string(path)).NodePath()))
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on spawn.
*/
func (self Instance) PropertySetSpawn(path Path.String, enabled bool) {
	class(self).PropertySetSpawn(gd.NewString(string(path)).NodePath(), enabled)
}

/*
Returns the replication mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
func (self Instance) PropertyGetReplicationMode(path Path.String) gdclass.SceneReplicationConfigReplicationMode {
	return gdclass.SceneReplicationConfigReplicationMode(class(self).PropertyGetReplicationMode(gd.NewString(string(path)).NodePath()))
}

/*
Sets the synchronization mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
func (self Instance) PropertySetReplicationMode(path Path.String, mode gdclass.SceneReplicationConfigReplicationMode) {
	class(self).PropertySetReplicationMode(gd.NewString(string(path)).NodePath(), mode)
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on process.
*/
func (self Instance) PropertyGetSync(path Path.String) bool {
	return bool(class(self).PropertyGetSync(gd.NewString(string(path)).NodePath()))
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on process.
*/
func (self Instance) PropertySetSync(path Path.String, enabled bool) {
	class(self).PropertySetSync(gd.NewString(string(path)).NodePath(), enabled)
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
func (self Instance) PropertyGetWatch(path Path.String) bool {
	return bool(class(self).PropertyGetWatch(gd.NewString(string(path)).NodePath()))
}

/*
Sets whether the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
func (self Instance) PropertySetWatch(path Path.String, enabled bool) {
	class(self).PropertySetWatch(gd.NewString(string(path)).NodePath(), enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SceneReplicationConfig

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SceneReplicationConfig"))
	return Instance{*(*gdclass.SceneReplicationConfig)(unsafe.Pointer(&object))}
}

/*
Returns a list of synchronized property [NodePath]s.
*/
//go:nosplit
func (self class) GetProperties() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_get_properties, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds the property identified by the given [param path] to the list of the properties being synchronized, optionally passing an [param index].
[b]Note:[/b] For details on restrictions and limitations on property synchronization, see [MultiplayerSynchronizer].
*/
//go:nosplit
func (self class) AddProperty(path gd.NodePath, index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_add_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param path] is configured for synchronization.
*/
//go:nosplit
func (self class) HasProperty(path gd.NodePath) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_has_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the property identified by the given [param path] from the configuration.
*/
//go:nosplit
func (self class) RemoveProperty(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_remove_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Finds the index of the given [param path].
*/
//go:nosplit
func (self class) PropertyGetIndex(path gd.NodePath) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on spawn.
*/
//go:nosplit
func (self class) PropertyGetSpawn(path gd.NodePath) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_spawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on spawn.
*/
//go:nosplit
func (self class) PropertySetSpawn(path gd.NodePath, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_spawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the replication mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
//go:nosplit
func (self class) PropertyGetReplicationMode(path gd.NodePath) gdclass.SceneReplicationConfigReplicationMode {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gdclass.SceneReplicationConfigReplicationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_replication_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the synchronization mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
//go:nosplit
func (self class) PropertySetReplicationMode(path gd.NodePath, mode gdclass.SceneReplicationConfigReplicationMode) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_replication_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on process.
*/
//go:nosplit
func (self class) PropertyGetSync(path gd.NodePath) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on process.
*/
//go:nosplit
func (self class) PropertySetSync(path gd.NodePath, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_sync, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
//go:nosplit
func (self class) PropertyGetWatch(path gd.NodePath) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_watch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
//go:nosplit
func (self class) PropertySetWatch(path gd.NodePath, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_watch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsSceneReplicationConfig() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSceneReplicationConfig() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	gdclass.Register("SceneReplicationConfig", func(ptr gd.Object) any {
		return [1]gdclass.SceneReplicationConfig{*(*gdclass.SceneReplicationConfig)(unsafe.Pointer(&ptr))}
	})
}

type ReplicationMode = gdclass.SceneReplicationConfigReplicationMode

const (
	/*Do not keep the given property synchronized.*/
	ReplicationModeNever ReplicationMode = 0
	/*Replicate the given property on process by constantly sending updates using unreliable transfer mode.*/
	ReplicationModeAlways ReplicationMode = 1
	/*Replicate the given property on process by sending updates using reliable transfer mode when its value changes.*/
	ReplicationModeOnChange ReplicationMode = 2
)
