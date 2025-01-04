package RDUniform

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDUniform
type Any interface {
	gd.IsClass
	AsRDUniform() Instance
}

/*
Binds the given id to the uniform. The data associated with the id is then used when the uniform is passed to a shader.
*/
func (self Instance) AddId(id Resource.ID) {
	class(self).AddId(id)
}

/*
Unbinds all ids currently bound to the uniform.
*/
func (self Instance) ClearIds() {
	class(self).ClearIds()
}

/*
Returns an array of all ids currently bound to the uniform.
*/
func (self Instance) GetIds() gd.Array {
	return gd.Array(class(self).GetIds())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDUniform

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDUniform"))
	return Instance{*(*gdclass.RDUniform)(unsafe.Pointer(&object))}
}

func (self Instance) UniformType() gdclass.RenderingDeviceUniformType {
	return gdclass.RenderingDeviceUniformType(class(self).GetUniformType())
}

func (self Instance) SetUniformType(value gdclass.RenderingDeviceUniformType) {
	class(self).SetUniformType(value)
}

func (self Instance) Binding() int {
	return int(int(class(self).GetBinding()))
}

func (self Instance) SetBinding(value int) {
	class(self).SetBinding(gd.Int(value))
}

//go:nosplit
func (self class) SetUniformType(p_member gdclass.RenderingDeviceUniformType) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_set_uniform_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUniformType() gdclass.RenderingDeviceUniformType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceUniformType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_uniform_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBinding(p_member gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_set_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBinding() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Binds the given id to the uniform. The data associated with the id is then used when the uniform is passed to a shader.
*/
//go:nosplit
func (self class) AddId(id gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_add_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Unbinds all ids currently bound to the uniform.
*/
//go:nosplit
func (self class) ClearIds() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_clear_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an array of all ids currently bound to the uniform.
*/
//go:nosplit
func (self class) GetIds() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_ids, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDUniform() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDUniform() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	gdclass.Register("RDUniform", func(ptr gd.Object) any { return [1]gdclass.RDUniform{*(*gdclass.RDUniform)(unsafe.Pointer(&ptr))} })
}
