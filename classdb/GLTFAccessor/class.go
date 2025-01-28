// Package GLTFAccessor provides methods for working with GLTFAccessor object instances.
package GLTFAccessor

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Resource"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
GLTFAccessor is a data structure representing GLTF a [code]accessor[/code] that would be found in the [code]"accessors"[/code] array. A buffer is a blob of binary data. A buffer view is a slice of a buffer. An accessor is a typed interpretation of the data in a buffer view.
Most custom data stored in GLTF does not need accessors, only buffer views (see [GLTFBufferView]). Accessors are for more advanced use cases such as interleaved mesh data encoded for the GPU.
*/
type Instance [1]gdclass.GLTFAccessor

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

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
	casted := Instance{*(*gdclass.GLTFAccessor)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
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
	return []float64(slices.Collect(class(self).GetMin().Values()))
}

func (self Instance) SetMin(value []float64) {
	class(self).SetMin(Packed.New(value...))
}

func (self Instance) Max() []float64 {
	return []float64(slices.Collect(class(self).GetMax().Values()))
}

func (self Instance) SetMax(value []float64) {
	class(self).SetMax(Packed.New(value...))
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
func (self class) GetBufferView() gd.Int { //gd:GLTFAccessor.get_buffer_view
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_buffer_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBufferView(buffer_view gd.Int) { //gd:GLTFAccessor.set_buffer_view
	var frame = callframe.New()
	callframe.Arg(frame, buffer_view)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_buffer_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetByteOffset() gd.Int { //gd:GLTFAccessor.get_byte_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_byte_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetByteOffset(byte_offset gd.Int) { //gd:GLTFAccessor.set_byte_offset
	var frame = callframe.New()
	callframe.Arg(frame, byte_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_byte_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetComponentType() gd.Int { //gd:GLTFAccessor.get_component_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_component_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetComponentType(component_type gd.Int) { //gd:GLTFAccessor.set_component_type
	var frame = callframe.New()
	callframe.Arg(frame, component_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_component_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalized() bool { //gd:GLTFAccessor.get_normalized
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_normalized, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalized(normalized bool) { //gd:GLTFAccessor.set_normalized
	var frame = callframe.New()
	callframe.Arg(frame, normalized)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_normalized, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCount() gd.Int { //gd:GLTFAccessor.get_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCount(count gd.Int) { //gd:GLTFAccessor.set_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAccessorType() gdclass.GLTFAccessorGLTFAccessorType { //gd:GLTFAccessor.get_accessor_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GLTFAccessorGLTFAccessorType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_accessor_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAccessorType(accessor_type gdclass.GLTFAccessorGLTFAccessorType) { //gd:GLTFAccessor.set_accessor_type
	var frame = callframe.New()
	callframe.Arg(frame, accessor_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_accessor_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetType() gd.Int { //gd:GLTFAccessor.get_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetType(atype gd.Int) { //gd:GLTFAccessor.set_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMin() Packed.Array[float64] { //gd:GLTFAccessor.get_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float64](Array.Through(gd.PackedProxy[gd.PackedFloat64Array, float64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMin(min Packed.Array[float64]) { //gd:GLTFAccessor.set_min
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat64Array, float64](min)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMax() Packed.Array[float64] { //gd:GLTFAccessor.get_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float64](Array.Through(gd.PackedProxy[gd.PackedFloat64Array, float64]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMax(max Packed.Array[float64]) { //gd:GLTFAccessor.set_max
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat64Array, float64](max)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseCount() gd.Int { //gd:GLTFAccessor.get_sparse_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseCount(sparse_count gd.Int) { //gd:GLTFAccessor.set_sparse_count
	var frame = callframe.New()
	callframe.Arg(frame, sparse_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseIndicesBufferView() gd.Int { //gd:GLTFAccessor.get_sparse_indices_buffer_view
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_indices_buffer_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseIndicesBufferView(sparse_indices_buffer_view gd.Int) { //gd:GLTFAccessor.set_sparse_indices_buffer_view
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_buffer_view)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_indices_buffer_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseIndicesByteOffset() gd.Int { //gd:GLTFAccessor.get_sparse_indices_byte_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_indices_byte_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseIndicesByteOffset(sparse_indices_byte_offset gd.Int) { //gd:GLTFAccessor.set_sparse_indices_byte_offset
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_byte_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_indices_byte_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseIndicesComponentType() gd.Int { //gd:GLTFAccessor.get_sparse_indices_component_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_indices_component_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseIndicesComponentType(sparse_indices_component_type gd.Int) { //gd:GLTFAccessor.set_sparse_indices_component_type
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_component_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_indices_component_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseValuesBufferView() gd.Int { //gd:GLTFAccessor.get_sparse_values_buffer_view
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_values_buffer_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseValuesBufferView(sparse_values_buffer_view gd.Int) { //gd:GLTFAccessor.set_sparse_values_buffer_view
	var frame = callframe.New()
	callframe.Arg(frame, sparse_values_buffer_view)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_values_buffer_view, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSparseValuesByteOffset() gd.Int { //gd:GLTFAccessor.get_sparse_values_byte_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_get_sparse_values_byte_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSparseValuesByteOffset(sparse_values_byte_offset gd.Int) { //gd:GLTFAccessor.set_sparse_values_byte_offset
	var frame = callframe.New()
	callframe.Arg(frame, sparse_values_byte_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFAccessor.Bind_set_sparse_values_byte_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
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

type GLTFAccessorType = gdclass.GLTFAccessorGLTFAccessorType //gd:GLTFAccessor.GLTFAccessorType

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
