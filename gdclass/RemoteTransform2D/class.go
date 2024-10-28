package RemoteTransform2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
RemoteTransform2D pushes its own [Transform2D] to another [Node2D] derived node (called the remote node) in the scene.
It can be set to update another node's position, rotation and/or scale. It can use either global or local coordinates.

*/
type Go [1]classdb.RemoteTransform2D

/*
[RemoteTransform2D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
func (self Go) ForceUpdateCache() {
	class(self).ForceUpdateCache()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RemoteTransform2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RemoteTransform2D"))
	return Go{classdb.RemoteTransform2D(object)}
}

func (self Go) RemotePath() string {
		return string(class(self).GetRemoteNode().String())
}

func (self Go) SetRemotePath(value string) {
	class(self).SetRemoteNode(gd.NewString(value).NodePath())
}

func (self Go) UseGlobalCoordinates() bool {
		return bool(class(self).GetUseGlobalCoordinates())
}

func (self Go) SetUseGlobalCoordinates(value bool) {
	class(self).SetUseGlobalCoordinates(value)
}

func (self Go) UpdatePosition() bool {
		return bool(class(self).GetUpdatePosition())
}

func (self Go) SetUpdatePosition(value bool) {
	class(self).SetUpdatePosition(value)
}

func (self Go) UpdateRotation() bool {
		return bool(class(self).GetUpdateRotation())
}

func (self Go) SetUpdateRotation(value bool) {
	class(self).SetUpdateRotation(value)
}

func (self Go) UpdateScale() bool {
		return bool(class(self).GetUpdateScale())
}

func (self Go) SetUpdateScale(value bool) {
	class(self).SetUpdateScale(value)
}

//go:nosplit
func (self class) SetRemoteNode(path gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_set_remote_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRemoteNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_get_remote_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
/*
[RemoteTransform2D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
//go:nosplit
func (self class) ForceUpdateCache()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_force_update_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseGlobalCoordinates(use_global_coordinates bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, use_global_coordinates)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_set_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseGlobalCoordinates() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_get_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdatePosition(update_remote_position bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_set_update_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdatePosition() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_get_update_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateRotation(update_remote_rotation bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_rotation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_set_update_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateRotation() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_get_update_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUpdateScale(update_remote_scale bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_set_update_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUpdateScale() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform2D.Bind_get_update_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRemoteTransform2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRemoteTransform2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("RemoteTransform2D", func(ptr gd.Object) any { return classdb.RemoteTransform2D(ptr) })}
