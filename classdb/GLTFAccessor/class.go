// Package GLTFAccessor provides methods for working with GLTFAccessor object instances.
package GLTFAccessor

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
GLTFAccessor is a data structure representing GLTF a [code]accessor[/code] that would be found in the [code]"accessors"[/code] array. A buffer is a blob of binary data. A buffer view is a slice of a buffer. An accessor is a typed interpretation of the data in a buffer view.
Most custom data stored in GLTF does not need accessors, only buffer views (see [GLTFBufferView]). Accessors are for more advanced use cases such as interleaved mesh data encoded for the GPU.
*/
type Instance [1]gdclass.GLTFAccessor
type Any interface {
	gd.IsClass
	AsGLTFAccessor() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFAccessor

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFAccessor"))
	return Instance{*(*gdclass.GLTFAccessor)(unsafe.Pointer(&object))}
}

func (self Instance) BufferView() int {
	return int(int(class(self).GetBufferView()))
}

func (self Instance) SetBufferView(value int) {
	class(self).SetBufferView(gd.Int(value))
}

func (self Instance) ByteOffset() int {
	return int(int(class(self).GetByteOffset()))
}

func (self Instance) SetByteOffset(value int) {
	class(self).SetByteOffset(gd.Int(value))
}

func (self Instance) ComponentType() int {
	return int(int(class(self).GetComponentType()))
}

func (self Instance) SetComponentType(value int) {
	class(self).SetComponentType(gd.Int(value))
}

func (self Instance) Normalized() bool {
	return bool(class(self).GetNormalized())
}

func (self Instance) SetNormalized(value bool) {
	class(self).SetNormalized(value)
}

func (self Instance) Count() int {
	return int(int(class(self).GetCount()))
}

func (self Instance) SetCount(value int) {
	class(self).SetCount(gd.Int(value))
}

func (self Instance) AccessorType() gdclass.GLTFAccessorGLTFAccessorType {
	return gdclass.GLTFAccessorGLTFAccessorType(class(self).GetAccessorType())
}

func (self Instance) SetAccessorType(value gdclass.GLTFAccessorGLTFAccessorType) {
	class(self).SetAccessorType(value)
}

func (self Instance) Type() int {
	return int(int(class(self).GetType()))
}

func (self Instance) SetType(value int) {
	class(self).SetType(gd.Int(value))
}

func (self Instance) Min() []float64 {
	return []float64(class(self).GetMin().AsSlice())
}

func (self Instance) SetMin(value []float64) {
	class(self).SetMin(gd.NewPackedFloat64Slice(value))
}

func (self Instance) Max() []float64 {
	return []float64(class(self).GetMax().AsSlice())
}

func (self Instance) SetMax(value []float64) {
	class(self).SetMax(gd.NewPackedFloat64Slice(value))
}

func (self Instance) SparseCount() int {
	return int(int(class(self).GetSparseCount()))
}

func (self Instance) SetSparseCount(value int) {
	class(self).SetSparseCount(gd.Int(value))
}

func (self Instance) SparseIndicesBufferView() int {
	return int(int(class(self).GetSparseIndicesBufferView()))
}

func (self Instance) SetSparseIndicesBufferView(value int) {
	class(self).SetSparseIndicesBufferView(gd.Int(value))
}

func (self Instance) SparseIndicesByteOffset() int {
	return int(int(class(self).GetSparseIndicesByteOffset()))
}

func (self Instance) SetSparseIndicesByteOffset(value int) {
	class(self).SetSparseIndicesByteOffset(gd.Int(value))
}

func (self Instance) SparseIndicesComponentType() int {
	return int(int(class(self).GetSparseIndicesComponentType()))
}

func (self Instance) SetSparseIndicesComponentType(value int) {
	class(self).SetSparseIndicesComponentType(gd.Int(value))
}

func (self Instance) SparseValuesBufferView() int {
	return int(int(class(self).GetSparseValuesBufferView()))
}

func (self Instance) SetSparseValuesBufferView(value int) {
	class(self).SetSparseValuesBufferView(gd.Int(value))
}

func (self Instance) SparseValuesByteOffset() int {
	return int(int(class(self).GetSparseValuesByteOffset()))
}

func (self Instance) SetSparseValuesByteOffset(value int) {
	class(self).SetSparseValuesByteOffset(gd.Int(value))
}

//go:nosplit
func (self class) GetBufferView() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBufferView(buffer_view gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, buffer_view)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetByteOffset() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetByteOffset(byte_offset gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, byte_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetComponentType() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetComponentType(component_type gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, component_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalized() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalized(normalized bool) {
	var frame = callframe.New()
	callframe.Arg(frame, normalized)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessorType() gdclass.GLTFAccessorGLTFAccessorType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GLTFAccessorGLTFAccessorType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_accessor_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessorType(accessor_type gdclass.GLTFAccessorGLTFAccessorType) {
	var frame = callframe.New()
	callframe.Arg(frame, accessor_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_accessor_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetType() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetType(atype gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMin() gd.PackedFloat64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat64Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMin(min gd.PackedFloat64Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(min))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMax() gd.PackedFloat64Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat64Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMax(max gd.PackedFloat64Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(max))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseCount(sparse_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sparse_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseIndicesBufferView() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_indices_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseIndicesBufferView(sparse_indices_buffer_view gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_buffer_view)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_indices_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseIndicesByteOffset() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_indices_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseIndicesByteOffset(sparse_indices_byte_offset gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_byte_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_indices_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseIndicesComponentType() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_indices_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseIndicesComponentType(sparse_indices_component_type gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_component_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_indices_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseValuesBufferView() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_values_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseValuesBufferView(sparse_values_buffer_view gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sparse_values_buffer_view)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_values_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseValuesByteOffset() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_values_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseValuesByteOffset(sparse_values_byte_offset gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, sparse_values_byte_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_values_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFAccessor() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFAccessor() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFAccessor", func(ptr gd.Object) any {
		return [1]gdclass.GLTFAccessor{*(*gdclass.GLTFAccessor)(unsafe.Pointer(&ptr))}
	})
}

type GLTFAccessorType = gdclass.GLTFAccessorGLTFAccessorType

const (
	/*Accessor type "SCALAR". For the glTF object model, this can be used to map to a single float, int, or bool value, or a float array.*/
	TypeScalar GLTFAccessorType = 0
	/*Accessor type "VEC2". For the glTF object model, this maps to "float2", represented in the glTF JSON as an array of two floats.*/
	TypeVec2 GLTFAccessorType = 1
	/*Accessor type "VEC3". For the glTF object model, this maps to "float3", represented in the glTF JSON as an array of three floats.*/
	TypeVec3 GLTFAccessorType = 2
	/*Accessor type "VEC4". For the glTF object model, this maps to "float4", represented in the glTF JSON as an array of four floats.*/
	TypeVec4 GLTFAccessorType = 3
	/*Accessor type "MAT2". For the glTF object model, this maps to "float2x2", represented in the glTF JSON as an array of four floats.*/
	TypeMat2 GLTFAccessorType = 4
	/*Accessor type "MAT3". For the glTF object model, this maps to "float3x3", represented in the glTF JSON as an array of nine floats.*/
	TypeMat3 GLTFAccessorType = 5
	/*Accessor type "MAT4". For the glTF object model, this maps to "float4x4", represented in the glTF JSON as an array of sixteen floats.*/
	TypeMat4 GLTFAccessorType = 6
)
