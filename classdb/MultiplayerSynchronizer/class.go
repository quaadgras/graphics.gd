// Package MultiplayerSynchronizer provides methods for working with MultiplayerSynchronizer object instances.
package MultiplayerSynchronizer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Callable"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
By default, [MultiplayerSynchronizer] synchronizes configured properties to all peers.
Visibility can be handled directly with [method set_visibility_for] or as-needed with [method add_visibility_filter] and [method update_visibility].
[MultiplayerSpawner]s will handle nodes according to visibility of synchronizers as long as the node at [member root_path] was spawned by one.
Internally, [MultiplayerSynchronizer] uses [method MultiplayerAPI.object_configuration_add] to notify synchronization start passing the [Node] at [member root_path] as the [code]object[/code] and itself as the [code]configuration[/code], and uses [method MultiplayerAPI.object_configuration_remove] to notify synchronization end in a similar way.
[b]Note:[/b] Synchronization is not supported for [Object] type properties, like [Resource]. Properties that are unique to each peer, like the instance IDs of [Object]s (see [method Object.get_instance_id]) or [RID]s, will also not work in synchronization.
*/
type Instance [1]gdclass.MultiplayerSynchronizer
type Any interface {
	gd.IsClass
	AsMultiplayerSynchronizer() Instance
}

/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
func (self Instance) UpdateVisibility() {
	class(self).UpdateVisibility(gd.Int(0))
}

/*
Adds a peer visibility filter for this synchronizer.
[param filter] should take a peer ID [int] and return a [bool].
*/
func (self Instance) AddVisibilityFilter(filter func(peer_id int) bool) {
	class(self).AddVisibilityFilter(gd.NewCallable(filter))
}

/*
Removes a peer visibility filter from this synchronizer.
*/
func (self Instance) RemoveVisibilityFilter(filter Callable.Any) {
	class(self).RemoveVisibilityFilter(gd.NewCallable(filter))
}

/*
Sets the visibility of [param peer] to [param visible]. If [param peer] is [code]0[/code], the value of [member public_visibility] will be updated instead.
*/
func (self Instance) SetVisibilityFor(peer int, visible bool) {
	class(self).SetVisibilityFor(gd.Int(peer), visible)
}

/*
Queries the current visibility for peer [param peer].
*/
func (self Instance) GetVisibilityFor(peer int) bool {
	return bool(class(self).GetVisibilityFor(gd.Int(peer)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerSynchronizer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerSynchronizer"))
	casted := Instance{*(*gdclass.MultiplayerSynchronizer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) RootPath() NodePath.String {
	return NodePath.String(class(self).GetRootPath().String())
}

func (self Instance) SetRootPath(value NodePath.String) {
	class(self).SetRootPath(gd.NewString(string(value)).NodePath())
}

func (self Instance) ReplicationInterval() Float.X {
	return Float.X(Float.X(class(self).GetReplicationInterval()))
}

func (self Instance) SetReplicationInterval(value Float.X) {
	class(self).SetReplicationInterval(gd.Float(value))
}

func (self Instance) DeltaInterval() Float.X {
	return Float.X(Float.X(class(self).GetDeltaInterval()))
}

func (self Instance) SetDeltaInterval(value Float.X) {
	class(self).SetDeltaInterval(gd.Float(value))
}

func (self Instance) ReplicationConfig() [1]gdclass.SceneReplicationConfig {
	return [1]gdclass.SceneReplicationConfig(class(self).GetReplicationConfig())
}

func (self Instance) SetReplicationConfig(value [1]gdclass.SceneReplicationConfig) {
	class(self).SetReplicationConfig(value)
}

func (self Instance) VisibilityUpdateMode() gdclass.MultiplayerSynchronizerVisibilityUpdateMode {
	return gdclass.MultiplayerSynchronizerVisibilityUpdateMode(class(self).GetVisibilityUpdateMode())
}

func (self Instance) SetVisibilityUpdateMode(value gdclass.MultiplayerSynchronizerVisibilityUpdateMode) {
	class(self).SetVisibilityUpdateMode(value)
}

func (self Instance) PublicVisibility() bool {
	return bool(class(self).IsVisibilityPublic())
}

func (self Instance) SetPublicVisibility(value bool) {
	class(self).SetVisibilityPublic(value)
}

//go:nosplit
func (self class) SetRootPath(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootPath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_root_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReplicationInterval(milliseconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, milliseconds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_replication_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetReplicationInterval() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_replication_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeltaInterval(milliseconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, milliseconds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_delta_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeltaInterval() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_delta_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReplicationConfig(config [1]gdclass.SceneReplicationConfig) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(config[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_replication_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetReplicationConfig() [1]gdclass.SceneReplicationConfig {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_replication_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.SceneReplicationConfig{gd.PointerWithOwnershipTransferredToGo[gdclass.SceneReplicationConfig](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityUpdateMode(mode gdclass.MultiplayerSynchronizerVisibilityUpdateMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_visibility_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityUpdateMode() gdclass.MultiplayerSynchronizerVisibilityUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.MultiplayerSynchronizerVisibilityUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_visibility_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
//go:nosplit
func (self class) UpdateVisibility(for_peer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, for_peer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_update_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityPublic(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_visibility_public, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisibilityPublic() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_is_visibility_public, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a peer visibility filter for this synchronizer.
[param filter] should take a peer ID [int] and return a [bool].
*/
//go:nosplit
func (self class) AddVisibilityFilter(filter gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filter))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_add_visibility_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a peer visibility filter from this synchronizer.
*/
//go:nosplit
func (self class) RemoveVisibilityFilter(filter gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filter))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_remove_visibility_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the visibility of [param peer] to [param visible]. If [param peer] is [code]0[/code], the value of [member public_visibility] will be updated instead.
*/
//go:nosplit
func (self class) SetVisibilityFor(peer gd.Int, visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_visibility_for, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Queries the current visibility for peer [param peer].
*/
//go:nosplit
func (self class) GetVisibilityFor(peer gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_visibility_for, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnSynchronized(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("synchronized"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDeltaSynchronized(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("delta_synchronized"), gd.NewCallable(cb), 0)
}

func (self Instance) OnVisibilityChanged(cb func(for_peer int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("visibility_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiplayerSynchronizer() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerSynchronizer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("MultiplayerSynchronizer", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerSynchronizer{*(*gdclass.MultiplayerSynchronizer)(unsafe.Pointer(&ptr))}
	})
}

type VisibilityUpdateMode = gdclass.MultiplayerSynchronizerVisibilityUpdateMode

const (
	/*Visibility filters are updated during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	VisibilityProcessIdle VisibilityUpdateMode = 0
	/*Visibility filters are updated during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	VisibilityProcessPhysics VisibilityUpdateMode = 1
	/*Visibility filters are not updated automatically, and must be updated manually by calling [method update_visibility].*/
	VisibilityProcessNone VisibilityUpdateMode = 2
)
