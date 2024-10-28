package RDUniform

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This object is used by [RenderingDevice].

*/
type Go [1]classdb.RDUniform

/*
Binds the given id to the uniform. The data associated with the id is then used when the uniform is passed to a shader.
*/
func (self Go) AddId(id gd.RID) {
	class(self).AddId(id)
}

/*
Unbinds all ids currently bound to the uniform.
*/
func (self Go) ClearIds() {
	class(self).ClearIds()
}

/*
Returns an array of all ids currently bound to the uniform.
*/
func (self Go) GetIds() gd.Array {
	return gd.Array(class(self).GetIds())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDUniform
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDUniform"))
	return Go{classdb.RDUniform(object)}
}

func (self Go) UniformType() classdb.RenderingDeviceUniformType {
		return classdb.RenderingDeviceUniformType(class(self).GetUniformType())
}

func (self Go) SetUniformType(value classdb.RenderingDeviceUniformType) {
	class(self).SetUniformType(value)
}

func (self Go) Binding() int {
		return int(int(class(self).GetBinding()))
}

func (self Go) SetBinding(value int) {
	class(self).SetBinding(gd.Int(value))
}

//go:nosplit
func (self class) SetUniformType(p_member classdb.RenderingDeviceUniformType)  {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_set_uniform_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUniformType() classdb.RenderingDeviceUniformType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RenderingDeviceUniformType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_uniform_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBinding(p_member gd.Int)  {
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
func (self class) AddId(id gd.RID)  {
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
func (self class) ClearIds()  {
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
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsRDUniform() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDUniform() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RDUniform", func(ptr gd.Object) any { return classdb.RDUniform(ptr) })}
