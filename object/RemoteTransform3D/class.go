package RemoteTransform3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
RemoteTransform3D pushes its own [Transform3D] to another [Node3D] derived Node (called the remote node) in the scene.
It can be set to update another Node's position, rotation and/or scale. It can use either global or local coordinates.

*/
type Simple [1]classdb.RemoteTransform3D
func (self Simple) SetRemoteNode(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRemoteNode(path)
}
func (self Simple) GetRemoteNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetRemoteNode(gc))
}
func (self Simple) ForceUpdateCache() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceUpdateCache()
}
func (self Simple) SetUseGlobalCoordinates(use_global_coordinates bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseGlobalCoordinates(use_global_coordinates)
}
func (self Simple) GetUseGlobalCoordinates() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseGlobalCoordinates())
}
func (self Simple) SetUpdatePosition(update_remote_position bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUpdatePosition(update_remote_position)
}
func (self Simple) GetUpdatePosition() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUpdatePosition())
}
func (self Simple) SetUpdateRotation(update_remote_rotation bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUpdateRotation(update_remote_rotation)
}
func (self Simple) GetUpdateRotation() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUpdateRotation())
}
func (self Simple) SetUpdateScale(update_remote_scale bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUpdateScale(update_remote_scale)
}
func (self Simple) GetUpdateScale() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUpdateScale())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RemoteTransform3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetRemoteNode(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_set_remote_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRemoteNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_get_remote_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
[RemoteTransform3D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
//go:nosplit
func (self class) ForceUpdateCache()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_force_update_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseGlobalCoordinates(use_global_coordinates bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_global_coordinates)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_set_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseGlobalCoordinates() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_get_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdatePosition(update_remote_position bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_set_update_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdatePosition() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_get_update_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateRotation(update_remote_rotation bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_rotation)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_set_update_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateRotation() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_get_update_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateScale(update_remote_scale bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_set_update_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateScale() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RemoteTransform3D.Bind_get_update_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsRemoteTransform3D() Expert { return self[0].AsRemoteTransform3D() }


//go:nosplit
func (self Simple) AsRemoteTransform3D() Simple { return self[0].AsRemoteTransform3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("RemoteTransform3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
