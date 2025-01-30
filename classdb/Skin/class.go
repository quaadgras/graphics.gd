// Package Skin provides methods for working with Skin object instances.
package Skin

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

type Instance [1]gdclass.Skin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkin() Instance
}

func (self Instance) SetBindCount(bind_count int) { //gd:Skin.set_bind_count
	class(self).SetBindCount(int64(bind_count))
}
func (self Instance) GetBindCount() int { //gd:Skin.get_bind_count
	return int(int(class(self).GetBindCount()))
}
func (self Instance) AddBind(bone int, pose Transform3D.BasisOrigin) { //gd:Skin.add_bind
	class(self).AddBind(int64(bone), Transform3D.BasisOrigin(pose))
}
func (self Instance) AddNamedBind(name string, pose Transform3D.BasisOrigin) { //gd:Skin.add_named_bind
	class(self).AddNamedBind(String.New(name), Transform3D.BasisOrigin(pose))
}
func (self Instance) SetBindPose(bind_index int, pose Transform3D.BasisOrigin) { //gd:Skin.set_bind_pose
	class(self).SetBindPose(int64(bind_index), Transform3D.BasisOrigin(pose))
}
func (self Instance) GetBindPose(bind_index int) Transform3D.BasisOrigin { //gd:Skin.get_bind_pose
	return Transform3D.BasisOrigin(class(self).GetBindPose(int64(bind_index)))
}
func (self Instance) SetBindName(bind_index int, name string) { //gd:Skin.set_bind_name
	class(self).SetBindName(int64(bind_index), String.Name(String.New(name)))
}
func (self Instance) GetBindName(bind_index int) string { //gd:Skin.get_bind_name
	return string(class(self).GetBindName(int64(bind_index)).String())
}
func (self Instance) SetBindBone(bind_index int, bone int) { //gd:Skin.set_bind_bone
	class(self).SetBindBone(int64(bind_index), int64(bone))
}
func (self Instance) GetBindBone(bind_index int) int { //gd:Skin.get_bind_bone
	return int(int(class(self).GetBindBone(int64(bind_index))))
}
func (self Instance) ClearBinds() { //gd:Skin.clear_binds
	class(self).ClearBinds()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Skin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Skin"))
	casted := Instance{*(*gdclass.Skin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

//go:nosplit
func (self class) SetBindCount(bind_count int64) { //gd:Skin.set_bind_count
	var frame = callframe.New()
	callframe.Arg(frame, bind_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindCount() int64 { //gd:Skin.get_bind_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AddBind(bone int64, pose Transform3D.BasisOrigin) { //gd:Skin.add_bind
	var frame = callframe.New()
	callframe.Arg(frame, bone)
	callframe.Arg(frame, pose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_add_bind, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) AddNamedBind(name String.Readable, pose Transform3D.BasisOrigin) { //gd:Skin.add_named_bind
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_add_named_bind, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetBindPose(bind_index int64, pose Transform3D.BasisOrigin) { //gd:Skin.set_bind_pose
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	callframe.Arg(frame, pose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindPose(bind_index int64) Transform3D.BasisOrigin { //gd:Skin.get_bind_pose
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_pose, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBindName(bind_index int64, name String.Name) { //gd:Skin.set_bind_name
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindName(bind_index int64) String.Name { //gd:Skin.get_bind_name
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBindBone(bind_index int64, bone int64) { //gd:Skin.set_bind_bone
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	callframe.Arg(frame, bone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindBone(bind_index int64) int64 { //gd:Skin.get_bind_bone
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) ClearBinds() { //gd:Skin.clear_binds
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_clear_binds, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsSkin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Skin", func(ptr gd.Object) any { return [1]gdclass.Skin{*(*gdclass.Skin)(unsafe.Pointer(&ptr))} })
}
