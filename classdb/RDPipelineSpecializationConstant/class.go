// Package RDPipelineSpecializationConstant provides methods for working with RDPipelineSpecializationConstant object instances.
package RDPipelineSpecializationConstant

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
import "graphics.gd/variant/String"

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

/*
A [i]specialization constant[/i] is a way to create additional variants of shaders without actually increasing the number of shader versions that are compiled. This allows improving performance by reducing the number of shader versions and reducing [code]if[/code] branching, while still allowing shaders to be flexible for different use cases.
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDPipelineSpecializationConstant

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDPipelineSpecializationConstant() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDPipelineSpecializationConstant

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineSpecializationConstant"))
	casted := Instance{*(*gdclass.RDPipelineSpecializationConstant)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Value() any {
	return any(class(self).GetValue().Interface())
}

func (self Instance) SetValue(value any) {
	class(self).SetValue(gd.NewVariant(value))
}

func (self Instance) ConstantId() int {
	return int(int(class(self).GetConstantId()))
}

func (self Instance) SetConstantId(value int) {
	class(self).SetConstantId(gd.Int(value))
}

//go:nosplit
func (self class) SetValue(value gd.Variant) { //gd:RDPipelineSpecializationConstant.set_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetValue() gd.Variant { //gd:RDPipelineSpecializationConstant.get_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConstantId(constant_id gd.Int) { //gd:RDPipelineSpecializationConstant.set_constant_id
	var frame = callframe.New()
	callframe.Arg(frame, constant_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_set_constant_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantId() gd.Int { //gd:RDPipelineSpecializationConstant.get_constant_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_get_constant_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineSpecializationConstant() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineSpecializationConstant() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
	gdclass.Register("RDPipelineSpecializationConstant", func(ptr gd.Object) any {
		return [1]gdclass.RDPipelineSpecializationConstant{*(*gdclass.RDPipelineSpecializationConstant)(unsafe.Pointer(&ptr))}
	})
}
