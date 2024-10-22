package SceneState

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Maintains a list of resources, nodes, exported and overridden properties, and built-in scripts associated with a scene. They cannot be modified from a [SceneState], only accessed. Useful for peeking into what a [PackedScene] contains without instantiating it.
This class cannot be instantiated directly, it is retrieved for a given scene as the result of [method PackedScene.get_state].

*/
type Go [1]classdb.SceneState

/*
Returns the number of nodes in the scene.
The [code]idx[/code] argument used to query node data in other [code]get_node_*[/code] methods in the interval [code][0, get_node_count() - 1][/code].
*/
func (self Go) GetNodeCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodeCount()))
}

/*
Returns the type of the node at [param idx].
*/
func (self Go) GetNodeType(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodeType(gc, gd.Int(idx)).String())
}

/*
Returns the name of the node at [param idx].
*/
func (self Go) GetNodeName(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodeName(gc, gd.Int(idx)).String())
}

/*
Returns the path to the node at [param idx].
If [param for_parent] is [code]true[/code], returns the path of the [param idx] node's parent instead.
*/
func (self Go) GetNodePath(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodePath(gc, gd.Int(idx), false).String())
}

/*
Returns the path to the owner of the node at [param idx], relative to the root node.
*/
func (self Go) GetNodeOwnerPath(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodeOwnerPath(gc, gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the node at [param idx] is an [InstancePlaceholder].
*/
func (self Go) IsNodeInstancePlaceholder(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsNodeInstancePlaceholder(gd.Int(idx)))
}

/*
Returns the path to the represented scene file if the node at [param idx] is an [InstancePlaceholder].
*/
func (self Go) GetNodeInstancePlaceholder(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodeInstancePlaceholder(gc, gd.Int(idx)).String())
}

/*
Returns a [PackedScene] for the node at [param idx] (i.e. the whole branch starting at this node, with its child nodes and resources), or [code]null[/code] if the node is not an instance.
*/
func (self Go) GetNodeInstance(idx int) gdclass.PackedScene {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PackedScene(class(self).GetNodeInstance(gc, gd.Int(idx)))
}

/*
Returns the list of group names associated with the node at [param idx].
*/
func (self Go) GetNodeGroups(idx int) []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetNodeGroups(gc, gd.Int(idx)).Strings(gc))
}

/*
Returns the node's index, which is its position relative to its siblings. This is only relevant and saved in scenes for cases where new nodes are added to an instantiated or inherited scene among siblings from the base scene. Despite the name, this index is not related to the [param idx] argument used here and in other methods.
*/
func (self Go) GetNodeIndex(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodeIndex(gd.Int(idx))))
}

/*
Returns the number of exported or overridden properties for the node at [param idx].
The [code]prop_idx[/code] argument used to query node property data in other [code]get_node_property_*[/code] methods in the interval [code][0, get_node_property_count() - 1][/code].
*/
func (self Go) GetNodePropertyCount(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetNodePropertyCount(gd.Int(idx))))
}

/*
Returns the name of the property at [param prop_idx] for the node at [param idx].
*/
func (self Go) GetNodePropertyName(idx int, prop_idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetNodePropertyName(gc, gd.Int(idx), gd.Int(prop_idx)).String())
}

/*
Returns the value of the property at [param prop_idx] for the node at [param idx].
*/
func (self Go) GetNodePropertyValue(idx int, prop_idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetNodePropertyValue(gc, gd.Int(idx), gd.Int(prop_idx)))
}

/*
Returns the number of signal connections in the scene.
The [code]idx[/code] argument used to query connection metadata in other [code]get_connection_*[/code] methods in the interval [code][0, get_connection_count() - 1][/code].
*/
func (self Go) GetConnectionCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetConnectionCount()))
}

/*
Returns the path to the node that owns the signal at [param idx], relative to the root node.
*/
func (self Go) GetConnectionSource(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetConnectionSource(gc, gd.Int(idx)).String())
}

/*
Returns the name of the signal at [param idx].
*/
func (self Go) GetConnectionSignal(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetConnectionSignal(gc, gd.Int(idx)).String())
}

/*
Returns the path to the node that owns the method connected to the signal at [param idx], relative to the root node.
*/
func (self Go) GetConnectionTarget(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetConnectionTarget(gc, gd.Int(idx)).String())
}

/*
Returns the method connected to the signal at [param idx].
*/
func (self Go) GetConnectionMethod(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetConnectionMethod(gc, gd.Int(idx)).String())
}

/*
Returns the connection flags for the signal at [param idx]. See [enum Object.ConnectFlags] constants.
*/
func (self Go) GetConnectionFlags(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetConnectionFlags(gd.Int(idx))))
}

/*
Returns the list of bound parameters for the signal at [param idx].
*/
func (self Go) GetConnectionBinds(idx int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(class(self).GetConnectionBinds(gc, gd.Int(idx)))
}

/*
Returns the number of unbound parameters for the signal at [param idx].
*/
func (self Go) GetConnectionUnbinds(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetConnectionUnbinds(gd.Int(idx))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SceneState
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SceneState"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the number of nodes in the scene.
The [code]idx[/code] argument used to query node data in other [code]get_node_*[/code] methods in the interval [code][0, get_node_count() - 1][/code].
*/
//go:nosplit
func (self class) GetNodeCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the type of the node at [param idx].
*/
//go:nosplit
func (self class) GetNodeType(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the node at [param idx].
*/
//go:nosplit
func (self class) GetNodeName(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to the node at [param idx].
If [param for_parent] is [code]true[/code], returns the path of the [param idx] node's parent instead.
*/
//go:nosplit
func (self class) GetNodePath(ctx gd.Lifetime, idx gd.Int, for_parent bool) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, for_parent)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to the owner of the node at [param idx], relative to the root node.
*/
//go:nosplit
func (self class) GetNodeOwnerPath(ctx gd.Lifetime, idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_owner_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the node at [param idx] is an [InstancePlaceholder].
*/
//go:nosplit
func (self class) IsNodeInstancePlaceholder(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_is_node_instance_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the path to the represented scene file if the node at [param idx] is an [InstancePlaceholder].
*/
//go:nosplit
func (self class) GetNodeInstancePlaceholder(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_instance_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a [PackedScene] for the node at [param idx] (i.e. the whole branch starting at this node, with its child nodes and resources), or [code]null[/code] if the node is not an instance.
*/
//go:nosplit
func (self class) GetNodeInstance(ctx gd.Lifetime, idx gd.Int) gdclass.PackedScene {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PackedScene
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the list of group names associated with the node at [param idx].
*/
//go:nosplit
func (self class) GetNodeGroups(ctx gd.Lifetime, idx gd.Int) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_groups, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the node's index, which is its position relative to its siblings. This is only relevant and saved in scenes for cases where new nodes are added to an instantiated or inherited scene among siblings from the base scene. Despite the name, this index is not related to the [param idx] argument used here and in other methods.
*/
//go:nosplit
func (self class) GetNodeIndex(idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of exported or overridden properties for the node at [param idx].
The [code]prop_idx[/code] argument used to query node property data in other [code]get_node_property_*[/code] methods in the interval [code][0, get_node_property_count() - 1][/code].
*/
//go:nosplit
func (self class) GetNodePropertyCount(idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_property_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the property at [param prop_idx] for the node at [param idx].
*/
//go:nosplit
func (self class) GetNodePropertyName(ctx gd.Lifetime, idx gd.Int, prop_idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, prop_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_property_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the value of the property at [param prop_idx] for the node at [param idx].
*/
//go:nosplit
func (self class) GetNodePropertyValue(ctx gd.Lifetime, idx gd.Int, prop_idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, prop_idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_node_property_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of signal connections in the scene.
The [code]idx[/code] argument used to query connection metadata in other [code]get_connection_*[/code] methods in the interval [code][0, get_connection_count() - 1][/code].
*/
//go:nosplit
func (self class) GetConnectionCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the path to the node that owns the signal at [param idx], relative to the root node.
*/
//go:nosplit
func (self class) GetConnectionSource(ctx gd.Lifetime, idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionSignal(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to the node that owns the method connected to the signal at [param idx], relative to the root node.
*/
//go:nosplit
func (self class) GetConnectionTarget(ctx gd.Lifetime, idx gd.Int) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_target, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the method connected to the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionMethod(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the connection flags for the signal at [param idx]. See [enum Object.ConnectFlags] constants.
*/
//go:nosplit
func (self class) GetConnectionFlags(idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of bound parameters for the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionBinds(ctx gd.Lifetime, idx gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the number of unbound parameters for the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionUnbinds(idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SceneState.Bind_get_connection_unbinds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSceneState() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSceneState() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("SceneState", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type GenEditState = classdb.SceneStateGenEditState

const (
/*If passed to [method PackedScene.instantiate], blocks edits to the scene state.*/
	GenEditStateDisabled GenEditState = 0
/*If passed to [method PackedScene.instantiate], provides inherited scene resources to the local scene.
[b]Note:[/b] Only available in editor builds.*/
	GenEditStateInstance GenEditState = 1
/*If passed to [method PackedScene.instantiate], provides local scene resources to the local scene. Only the main scene should receive the main edit state.
[b]Note:[/b] Only available in editor builds.*/
	GenEditStateMain GenEditState = 2
/*If passed to [method PackedScene.instantiate], it's similar to [constant GEN_EDIT_STATE_MAIN], but for the case where the scene is being instantiated to be the base of another one.
[b]Note:[/b] Only available in editor builds.*/
	GenEditStateMainInherited GenEditState = 3
)
