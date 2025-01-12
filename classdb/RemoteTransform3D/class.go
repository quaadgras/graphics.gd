// Package RemoteTransform3D provides methods for working with RemoteTransform3D object instances.
package RemoteTransform3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
RemoteTransform3D pushes its own [Transform3D] to another [Node3D] derived Node (called the remote node) in the scene.
It can be set to update another Node's position, rotation and/or scale. It can use either global or local coordinates.
*/
type Instance [1]gdclass.RemoteTransform3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRemoteTransform3D() Instance
}

/*
[RemoteTransform3D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
func (self Instance) ForceUpdateCache() {
	class(self).ForceUpdateCache()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RemoteTransform3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RemoteTransform3D"))
	casted := Instance{*(*gdclass.RemoteTransform3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) RemotePath() NodePath.String {
	return NodePath.String(class(self).GetRemoteNode().String())
}

func (self Instance) SetRemotePath(value NodePath.String) {
	class(self).SetRemoteNode(gd.NewString(string(value)).NodePath())
}

func (self Instance) UseGlobalCoordinates() bool {
	return bool(class(self).GetUseGlobalCoordinates())
}

func (self Instance) SetUseGlobalCoordinates(value bool) {
	class(self).SetUseGlobalCoordinates(value)
}

func (self Instance) UpdatePosition() bool {
	return bool(class(self).GetUpdatePosition())
}

func (self Instance) SetUpdatePosition(value bool) {
	class(self).SetUpdatePosition(value)
}

func (self Instance) UpdateRotation() bool {
	return bool(class(self).GetUpdateRotation())
}

func (self Instance) SetUpdateRotation(value bool) {
	class(self).SetUpdateRotation(value)
}

func (self Instance) UpdateScale() bool {
	return bool(class(self).GetUpdateScale())
}

func (self Instance) SetUpdateScale(value bool) {
	class(self).SetUpdateScale(value)
}

//go:nosplit
func (self class) SetRemoteNode(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_remote_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRemoteNode() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_remote_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
[RemoteTransform3D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
//go:nosplit
func (self class) ForceUpdateCache() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_force_update_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseGlobalCoordinates(use_global_coordinates bool) {
	var frame = callframe.New()
	callframe.Arg(frame, use_global_coordinates)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseGlobalCoordinates() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdatePosition(update_remote_position bool) {
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_update_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdatePosition() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_update_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdateRotation(update_remote_rotation bool) {
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_update_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateRotation() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_update_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdateScale(update_remote_scale bool) {
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_update_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateScale() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_update_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRemoteTransform3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRemoteTransform3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced        { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance     { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced            { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance         { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("RemoteTransform3D", func(ptr gd.Object) any {
		return [1]gdclass.RemoteTransform3D{*(*gdclass.RemoteTransform3D)(unsafe.Pointer(&ptr))}
	})
}
