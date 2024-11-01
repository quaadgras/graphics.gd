package Skin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

type Instance [1]classdb.Skin

func (self Instance) SetBindCount(bind_count int) {
	class(self).SetBindCount(gd.Int(bind_count))
}
func (self Instance) GetBindCount() int {
	return int(int(class(self).GetBindCount()))
}
func (self Instance) AddBind(bone int, pose gd.Transform3D) {
	class(self).AddBind(gd.Int(bone), pose)
}
func (self Instance) AddNamedBind(name string, pose gd.Transform3D) {
	class(self).AddNamedBind(gd.NewString(name), pose)
}
func (self Instance) SetBindPose(bind_index int, pose gd.Transform3D) {
	class(self).SetBindPose(gd.Int(bind_index), pose)
}
func (self Instance) GetBindPose(bind_index int) gd.Transform3D {
	return gd.Transform3D(class(self).GetBindPose(gd.Int(bind_index)))
}
func (self Instance) SetBindName(bind_index int, name string) {
	class(self).SetBindName(gd.Int(bind_index), gd.NewStringName(name))
}
func (self Instance) GetBindName(bind_index int) string {
	return string(class(self).GetBindName(gd.Int(bind_index)).String())
}
func (self Instance) SetBindBone(bind_index int, bone int) {
	class(self).SetBindBone(gd.Int(bind_index), gd.Int(bone))
}
func (self Instance) GetBindBone(bind_index int) int {
	return int(int(class(self).GetBindBone(gd.Int(bind_index))))
}
func (self Instance) ClearBinds() {
	class(self).ClearBinds()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Skin

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Skin"))
	return Instance{classdb.Skin(object)}
}

//go:nosplit
func (self class) SetBindCount(bind_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bind_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AddBind(bone gd.Int, pose gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, bone)
	callframe.Arg(frame, pose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_add_bind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AddNamedBind(name gd.String, pose gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_add_named_bind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetBindPose(bind_index gd.Int, pose gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	callframe.Arg(frame, pose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindPose(bind_index gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBindName(bind_index gd.Int, name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindName(bind_index gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBindBone(bind_index gd.Int, bone gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	callframe.Arg(frame, bone)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_set_bind_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindBone(bind_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, bind_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_get_bind_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) ClearBinds() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skin.Bind_clear_binds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() { classdb.Register("Skin", func(ptr gd.Object) any { return classdb.Skin(ptr) }) }
