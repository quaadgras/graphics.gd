package GLTFAccessor

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
GLTFAccessor is a data structure representing GLTF a [code]accessor[/code] that would be found in the [code]"accessors"[/code] array. A buffer is a blob of binary data. A buffer view is a slice of a buffer. An accessor is a typed interpretation of the data in a buffer view.
Most custom data stored in GLTF does not need accessors, only buffer views (see [GLTFBufferView]). Accessors are for more advanced use cases such as interleaved mesh data encoded for the GPU.

*/
type Go [1]classdb.GLTFAccessor
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFAccessor
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFAccessor"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BufferView() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBufferView()))
}

func (self Go) SetBufferView(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBufferView(gd.Int(value))
}

func (self Go) ByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetByteOffset()))
}

func (self Go) SetByteOffset(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetByteOffset(gd.Int(value))
}

func (self Go) ComponentType() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetComponentType()))
}

func (self Go) SetComponentType(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetComponentType(gd.Int(value))
}

func (self Go) Normalized() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetNormalized())
}

func (self Go) SetNormalized(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetNormalized(value)
}

func (self Go) Count() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCount()))
}

func (self Go) SetCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCount(gd.Int(value))
}

func (self Go) AccessorType() classdb.GLTFAccessorGLTFAccessorType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GLTFAccessorGLTFAccessorType(class(self).GetAccessorType())
}

func (self Go) SetAccessorType(value classdb.GLTFAccessorGLTFAccessorType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAccessorType(value)
}

func (self Go) Type() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetType()))
}

func (self Go) SetType(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetType(gd.Int(value))
}

func (self Go) Min() []float64 {
	gc := gd.GarbageCollector(); _ = gc
		return []float64(class(self).GetMin(gc).AsSlice())
}

func (self Go) SetMin(value []float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMin(gc.PackedFloat64Slice(value))
}

func (self Go) Max() []float64 {
	gc := gd.GarbageCollector(); _ = gc
		return []float64(class(self).GetMax(gc).AsSlice())
}

func (self Go) SetMax(value []float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMax(gc.PackedFloat64Slice(value))
}

func (self Go) SparseCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSparseCount()))
}

func (self Go) SetSparseCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSparseCount(gd.Int(value))
}

func (self Go) SparseIndicesBufferView() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSparseIndicesBufferView()))
}

func (self Go) SetSparseIndicesBufferView(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSparseIndicesBufferView(gd.Int(value))
}

func (self Go) SparseIndicesByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSparseIndicesByteOffset()))
}

func (self Go) SetSparseIndicesByteOffset(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSparseIndicesByteOffset(gd.Int(value))
}

func (self Go) SparseIndicesComponentType() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSparseIndicesComponentType()))
}

func (self Go) SetSparseIndicesComponentType(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSparseIndicesComponentType(gd.Int(value))
}

func (self Go) SparseValuesBufferView() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSparseValuesBufferView()))
}

func (self Go) SetSparseValuesBufferView(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSparseValuesBufferView(gd.Int(value))
}

func (self Go) SparseValuesByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSparseValuesByteOffset()))
}

func (self Go) SetSparseValuesByteOffset(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSparseValuesByteOffset(gd.Int(value))
}

//go:nosplit
func (self class) GetBufferView() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBufferView(buffer_view gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer_view)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetByteOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetByteOffset(byte_offset gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, byte_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetComponentType() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetComponentType(component_type gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, component_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNormalized() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNormalized(normalized bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, normalized)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAccessorType() classdb.GLTFAccessorGLTFAccessorType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GLTFAccessorGLTFAccessorType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_accessor_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAccessorType(accessor_type classdb.GLTFAccessorGLTFAccessorType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, accessor_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_accessor_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetType() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetType(atype gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMin(ctx gd.Lifetime) gd.PackedFloat64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMin(min gd.PackedFloat64Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(min))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMax(ctx gd.Lifetime) gd.PackedFloat64Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat64Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMax(max gd.PackedFloat64Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(max))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSparseCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_sparse_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSparseCount(sparse_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sparse_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_sparse_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSparseIndicesBufferView() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_sparse_indices_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSparseIndicesBufferView(sparse_indices_buffer_view gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_buffer_view)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_sparse_indices_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSparseIndicesByteOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_sparse_indices_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSparseIndicesByteOffset(sparse_indices_byte_offset gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_byte_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_sparse_indices_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSparseIndicesComponentType() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_sparse_indices_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSparseIndicesComponentType(sparse_indices_component_type gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sparse_indices_component_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_sparse_indices_component_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSparseValuesBufferView() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_sparse_values_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSparseValuesBufferView(sparse_values_buffer_view gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sparse_values_buffer_view)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_sparse_values_buffer_view, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSparseValuesByteOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_get_sparse_values_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSparseValuesByteOffset(sparse_values_byte_offset gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sparse_values_byte_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFAccessor.Bind_set_sparse_values_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFAccessor() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFAccessor() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFAccessor", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type GLTFAccessorType = classdb.GLTFAccessorGLTFAccessorType

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
