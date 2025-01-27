// Package SceneReplicationConfig provides methods for working with SceneReplicationConfig object instances.
package SceneReplicationConfig

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

type Instance [1]gdclass.SceneReplicationConfig

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSceneReplicationConfig() Instance
}

/*
Returns a list of synchronized property [NodePath]s.
*/
func (self Instance) GetProperties() []string { //gd:SceneReplicationConfig.get_properties
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetProperties())))
}

/*
Adds the property identified by the given [param path] to the list of the properties being synchronized, optionally passing an [param index].
[b]Note:[/b] For details on restrictions and limitations on property synchronization, see [MultiplayerSynchronizer].
*/
func (self Instance) AddProperty(path string) { //gd:SceneReplicationConfig.add_property
	class(self).AddProperty(Path.ToNode(String.New(path)), gd.Int(-1))
}

/*
Returns [code]true[/code] if the given [param path] is configured for synchronization.
*/
func (self Instance) HasProperty(path string) bool { //gd:SceneReplicationConfig.has_property
	return bool(class(self).HasProperty(Path.ToNode(String.New(path))))
}

/*
Removes the property identified by the given [param path] from the configuration.
*/
func (self Instance) RemoveProperty(path string) { //gd:SceneReplicationConfig.remove_property
	class(self).RemoveProperty(Path.ToNode(String.New(path)))
}

/*
Finds the index of the given [param path].
*/
func (self Instance) PropertyGetIndex(path string) int { //gd:SceneReplicationConfig.property_get_index
	return int(int(class(self).PropertyGetIndex(Path.ToNode(String.New(path)))))
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on spawn.
*/
func (self Instance) PropertyGetSpawn(path string) bool { //gd:SceneReplicationConfig.property_get_spawn
	return bool(class(self).PropertyGetSpawn(Path.ToNode(String.New(path))))
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on spawn.
*/
func (self Instance) PropertySetSpawn(path string, enabled bool) { //gd:SceneReplicationConfig.property_set_spawn
	class(self).PropertySetSpawn(Path.ToNode(String.New(path)), enabled)
}

/*
Returns the replication mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
func (self Instance) PropertyGetReplicationMode(path string) gdclass.SceneReplicationConfigReplicationMode { //gd:SceneReplicationConfig.property_get_replication_mode
	return gdclass.SceneReplicationConfigReplicationMode(class(self).PropertyGetReplicationMode(Path.ToNode(String.New(path))))
}

/*
Sets the synchronization mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
func (self Instance) PropertySetReplicationMode(path string, mode gdclass.SceneReplicationConfigReplicationMode) { //gd:SceneReplicationConfig.property_set_replication_mode
	class(self).PropertySetReplicationMode(Path.ToNode(String.New(path)), mode)
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on process.
*/
func (self Instance) PropertyGetSync(path string) bool { //gd:SceneReplicationConfig.property_get_sync
	return bool(class(self).PropertyGetSync(Path.ToNode(String.New(path))))
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on process.
*/
func (self Instance) PropertySetSync(path string, enabled bool) { //gd:SceneReplicationConfig.property_set_sync
	class(self).PropertySetSync(Path.ToNode(String.New(path)), enabled)
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
func (self Instance) PropertyGetWatch(path string) bool { //gd:SceneReplicationConfig.property_get_watch
	return bool(class(self).PropertyGetWatch(Path.ToNode(String.New(path))))
}

/*
Sets whether the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
func (self Instance) PropertySetWatch(path string, enabled bool) { //gd:SceneReplicationConfig.property_set_watch
	class(self).PropertySetWatch(Path.ToNode(String.New(path)), enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SceneReplicationConfig

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SceneReplicationConfig"))
	casted := Instance{*(*gdclass.SceneReplicationConfig)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns a list of synchronized property [NodePath]s.
*/
//go:nosplit
func (self class) GetProperties() Array.Contains[Path.ToNode] { //gd:SceneReplicationConfig.get_properties
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_get_properties, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Path.ToNode]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Adds the property identified by the given [param path] to the list of the properties being synchronized, optionally passing an [param index].
[b]Note:[/b] For details on restrictions and limitations on property synchronization, see [MultiplayerSynchronizer].
*/
//go:nosplit
func (self class) AddProperty(path Path.ToNode, index gd.Int) { //gd:SceneReplicationConfig.add_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_add_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param path] is configured for synchronization.
*/
//go:nosplit
func (self class) HasProperty(path Path.ToNode) bool { //gd:SceneReplicationConfig.has_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_has_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the property identified by the given [param path] from the configuration.
*/
//go:nosplit
func (self class) RemoveProperty(path Path.ToNode) { //gd:SceneReplicationConfig.remove_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_remove_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Finds the index of the given [param path].
*/
//go:nosplit
func (self class) PropertyGetIndex(path Path.ToNode) gd.Int { //gd:SceneReplicationConfig.property_get_index
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on spawn.
*/
//go:nosplit
func (self class) PropertyGetSpawn(path Path.ToNode) bool { //gd:SceneReplicationConfig.property_get_spawn
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_spawn, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on spawn.
*/
//go:nosplit
func (self class) PropertySetSpawn(path Path.ToNode, enabled bool) { //gd:SceneReplicationConfig.property_set_spawn
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_spawn, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the replication mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
//go:nosplit
func (self class) PropertyGetReplicationMode(path Path.ToNode) gdclass.SceneReplicationConfigReplicationMode { //gd:SceneReplicationConfig.property_get_replication_mode
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[gdclass.SceneReplicationConfigReplicationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_replication_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the synchronization mode for the property identified by the given [param path]. See [enum ReplicationMode].
*/
//go:nosplit
func (self class) PropertySetReplicationMode(path Path.ToNode, mode gdclass.SceneReplicationConfigReplicationMode) { //gd:SceneReplicationConfig.property_set_replication_mode
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_replication_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be synchronized on process.
*/
//go:nosplit
func (self class) PropertyGetSync(path Path.ToNode) bool { //gd:SceneReplicationConfig.property_get_sync
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the property identified by the given [param path] is configured to be synchronized on process.
*/
//go:nosplit
func (self class) PropertySetSync(path Path.ToNode, enabled bool) { //gd:SceneReplicationConfig.property_set_sync
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
//go:nosplit
func (self class) PropertyGetWatch(path Path.ToNode) bool { //gd:SceneReplicationConfig.property_get_watch
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_get_watch, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the property identified by the given [param path] is configured to be reliably synchronized when changes are detected on process.
*/
//go:nosplit
func (self class) PropertySetWatch(path Path.ToNode, enabled bool) { //gd:SceneReplicationConfig.property_set_watch
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneReplicationConfig.Bind_property_set_watch, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("SceneReplicationConfig", func(ptr gd.Object) any {
		return [1]gdclass.SceneReplicationConfig{*(*gdclass.SceneReplicationConfig)(unsafe.Pointer(&ptr))}
	})
}

type ReplicationMode = gdclass.SceneReplicationConfigReplicationMode //gd:SceneReplicationConfig.ReplicationMode

const (
	/*Do not keep the given property synchronized.*/
	ReplicationModeNever ReplicationMode = 0
	/*Replicate the given property on process by constantly sending updates using unreliable transfer mode.*/
	ReplicationModeAlways ReplicationMode = 1
	/*Replicate the given property on process by sending updates using reliable transfer mode when its value changes.*/
	ReplicationModeOnChange ReplicationMode = 2
)
