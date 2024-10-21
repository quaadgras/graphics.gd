package GLTFAccessor

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
GLTFAccessor is a data structure representing GLTF a [code]accessor[/code] that would be found in the [code]"accessors"[/code] array. A buffer is a blob of binary data. A buffer view is a slice of a buffer. An accessor is a typed interpretation of the data in a buffer view.
Most custom data stored in GLTF does not need accessors, only buffer views (see [GLTFBufferView]). Accessors are for more advanced use cases such as interleaved mesh data encoded for the GPU.

*/
type Simple [1]classdb.GLTFAccessor
func (self Simple) GetBufferView() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBufferView()))
}
func (self Simple) SetBufferView(buffer_view int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBufferView(gd.Int(buffer_view))
}
func (self Simple) GetByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetByteOffset()))
}
func (self Simple) SetByteOffset(byte_offset int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetByteOffset(gd.Int(byte_offset))
}
func (self Simple) GetComponentType() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetComponentType()))
}
func (self Simple) SetComponentType(component_type int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetComponentType(gd.Int(component_type))
}
func (self Simple) GetNormalized() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNormalized())
}
func (self Simple) SetNormalized(normalized bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNormalized(normalized)
}
func (self Simple) GetCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCount()))
}
func (self Simple) SetCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCount(gd.Int(count))
}
func (self Simple) GetAccessorType() classdb.GLTFAccessorGLTFAccessorType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GLTFAccessorGLTFAccessorType(Expert(self).GetAccessorType())
}
func (self Simple) SetAccessorType(accessor_type classdb.GLTFAccessorGLTFAccessorType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAccessorType(accessor_type)
}
func (self Simple) GetType() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetType()))
}
func (self Simple) SetType(atype int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetType(gd.Int(atype))
}
func (self Simple) GetMin() gd.PackedFloat64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat64Array(Expert(self).GetMin(gc))
}
func (self Simple) SetMin(min gd.PackedFloat64Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMin(min)
}
func (self Simple) GetMax() gd.PackedFloat64Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat64Array(Expert(self).GetMax(gc))
}
func (self Simple) SetMax(max gd.PackedFloat64Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMax(max)
}
func (self Simple) GetSparseCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSparseCount()))
}
func (self Simple) SetSparseCount(sparse_count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSparseCount(gd.Int(sparse_count))
}
func (self Simple) GetSparseIndicesBufferView() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSparseIndicesBufferView()))
}
func (self Simple) SetSparseIndicesBufferView(sparse_indices_buffer_view int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSparseIndicesBufferView(gd.Int(sparse_indices_buffer_view))
}
func (self Simple) GetSparseIndicesByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSparseIndicesByteOffset()))
}
func (self Simple) SetSparseIndicesByteOffset(sparse_indices_byte_offset int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSparseIndicesByteOffset(gd.Int(sparse_indices_byte_offset))
}
func (self Simple) GetSparseIndicesComponentType() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSparseIndicesComponentType()))
}
func (self Simple) SetSparseIndicesComponentType(sparse_indices_component_type int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSparseIndicesComponentType(gd.Int(sparse_indices_component_type))
}
func (self Simple) GetSparseValuesBufferView() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSparseValuesBufferView()))
}
func (self Simple) SetSparseValuesBufferView(sparse_values_buffer_view int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSparseValuesBufferView(gd.Int(sparse_values_buffer_view))
}
func (self Simple) GetSparseValuesByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSparseValuesByteOffset()))
}
func (self Simple) SetSparseValuesByteOffset(sparse_values_byte_offset int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSparseValuesByteOffset(gd.Int(sparse_values_byte_offset))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFAccessor
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsGLTFAccessor() Expert { return self[0].AsGLTFAccessor() }


//go:nosplit
func (self Simple) AsGLTFAccessor() Simple { return self[0].AsGLTFAccessor() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("GLTFAccessor", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
