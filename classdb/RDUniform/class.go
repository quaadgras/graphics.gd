// Package RDUniform provides methods for working with RDUniform object instances.
package RDUniform

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
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"

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

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDUniform

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDUniform() Instance
}

/*
Binds the given id to the uniform. The data associated with the id is then used when the uniform is passed to a shader.
*/
func (self Instance) AddId(id RID.Any) { //gd:RDUniform.add_id
	class(self).AddId(gd.RID(id))
}

/*
Unbinds all ids currently bound to the uniform.
*/
func (self Instance) ClearIds() { //gd:RDUniform.clear_ids
	class(self).ClearIds()
}

/*
Returns an array of all ids currently bound to the uniform.
*/
func (self Instance) GetIds() [][]RID.Any { //gd:RDUniform.get_ids
	return [][]RID.Any(gd.ArrayAs[[][]RID.Any](gd.InternalArray(class(self).GetIds())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDUniform

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDUniform"))
	casted := Instance{*(*gdclass.RDUniform)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
func (self class) SetUniformType(p_member gdclass.RenderingDeviceUniformType) { //gd:RDUniform.set_uniform_type
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_set_uniform_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUniformType() gdclass.RenderingDeviceUniformType { //gd:RDUniform.get_uniform_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceUniformType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_uniform_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBinding(p_member gd.Int) { //gd:RDUniform.set_binding
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_set_binding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBinding() gd.Int { //gd:RDUniform.get_binding
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Binds the given id to the uniform. The data associated with the id is then used when the uniform is passed to a shader.
*/
//go:nosplit
func (self class) AddId(id gd.RID) { //gd:RDUniform.add_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_add_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unbinds all ids currently bound to the uniform.
*/
//go:nosplit
func (self class) ClearIds() { //gd:RDUniform.clear_ids
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_clear_ids, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all ids currently bound to the uniform.
*/
//go:nosplit
func (self class) GetIds() Array.Contains[gd.RID] { //gd:RDUniform.get_ids
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDUniform.Bind_get_ids, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsRDUniform() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRDUniform() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDUniform", func(ptr gd.Object) any { return [1]gdclass.RDUniform{*(*gdclass.RDUniform)(unsafe.Pointer(&ptr))} })
}
