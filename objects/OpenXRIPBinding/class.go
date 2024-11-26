package OpenXRIPBinding

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This binding resource binds an [OpenXRAction] to inputs or outputs. As most controllers have left hand and right versions that are handled by the same interaction profile we can specify multiple bindings. For instance an action "Fire" could be bound to both "/user/hand/left/input/trigger" and "/user/hand/right/input/trigger".
*/
type Instance [1]classdb.OpenXRIPBinding

/*
Get the number of input/output paths in this binding.
*/
func (self Instance) GetPathCount() int {
	return int(int(class(self).GetPathCount()))
}

/*
Returns [code]true[/code] if this input/output path is part of this binding.
*/
func (self Instance) HasPath(path string) bool {
	return bool(class(self).HasPath(gd.NewString(path)))
}

/*
Add an input/output path to this binding.
*/
func (self Instance) AddPath(path string) {
	class(self).AddPath(gd.NewString(path))
}

/*
Removes this input/output path from this binding.
*/
func (self Instance) RemovePath(path string) {
	class(self).RemovePath(gd.NewString(path))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRIPBinding

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRIPBinding"))
	return Instance{classdb.OpenXRIPBinding(object)}
}

func (self Instance) Action() objects.OpenXRAction {
	return objects.OpenXRAction(class(self).GetAction())
}

func (self Instance) SetAction(value objects.OpenXRAction) {
	class(self).SetAction(value)
}

func (self Instance) Paths() []string {
	return []string(class(self).GetPaths().Strings())
}

func (self Instance) SetPaths(value []string) {
	class(self).SetPaths(gd.NewPackedStringSlice(value))
}

//go:nosplit
func (self class) SetAction(action objects.OpenXRAction) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_set_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAction() objects.OpenXRAction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_get_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.OpenXRAction{classdb.OpenXRAction(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Get the number of input/output paths in this binding.
*/
//go:nosplit
func (self class) GetPathCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_get_path_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPaths(paths gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(paths))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_set_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPaths() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_get_paths, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this input/output path is part of this binding.
*/
//go:nosplit
func (self class) HasPath(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_has_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Add an input/output path to this binding.
*/
//go:nosplit
func (self class) AddPath(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_add_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes this input/output path from this binding.
*/
//go:nosplit
func (self class) RemovePath(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_remove_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRIPBinding() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRIPBinding() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("OpenXRIPBinding", func(ptr gd.Object) any { return classdb.OpenXRIPBinding(ptr) })
}
