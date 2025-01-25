// Package OpenXRIPBinding provides methods for working with OpenXRIPBinding object instances.
package OpenXRIPBinding

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
This binding resource binds an [OpenXRAction] to inputs or outputs. As most controllers have left hand and right versions that are handled by the same interaction profile we can specify multiple bindings. For instance an action "Fire" could be bound to both "/user/hand/left/input/trigger" and "/user/hand/right/input/trigger".
*/
type Instance [1]gdclass.OpenXRIPBinding

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRIPBinding() Instance
}

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
type class [1]gdclass.OpenXRIPBinding

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRIPBinding"))
	casted := Instance{*(*gdclass.OpenXRIPBinding)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Action() [1]gdclass.OpenXRAction {
	return [1]gdclass.OpenXRAction(class(self).GetAction())
}

func (self Instance) SetAction(value [1]gdclass.OpenXRAction) {
	class(self).SetAction(value)
}

func (self Instance) Paths() []string {
	return []string(class(self).GetPaths().Strings())
}

func (self Instance) SetPaths(value []string) {
	class(self).SetPaths(gd.NewPackedStringSlice(value))
}

//go:nosplit
func (self class) SetAction(action [1]gdclass.OpenXRAction) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_set_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAction() [1]gdclass.OpenXRAction {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_get_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRAction{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRAction](r_ret.Get())}
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_get_path_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPaths(paths gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(paths))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_set_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPaths() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_get_paths, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_has_path, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_add_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes this input/output path from this binding.
*/
//go:nosplit
func (self class) RemovePath(path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRIPBinding.Bind_remove_path, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("OpenXRIPBinding", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRIPBinding{*(*gdclass.OpenXRIPBinding)(unsafe.Pointer(&ptr))}
	})
}
