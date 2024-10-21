package GLTFBufferView

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
GLTFBufferView is a data structure representing GLTF a [code]bufferView[/code] that would be found in the [code]"bufferViews"[/code] array. A buffer is a blob of binary data. A buffer view is a slice of a buffer that can be used to identify and extract data from the buffer.
Most custom uses of buffers only need to use the [member buffer], [member byte_length], and [member byte_offset]. The [member byte_stride] and [member indices] properties are for more advanced use cases such as interleaved mesh data encoded for the GPU.

*/
type Simple [1]classdb.GLTFBufferView
func (self Simple) LoadBufferViewData(state [1]classdb.GLTFState) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).LoadBufferViewData(gc, state).Bytes())
}
func (self Simple) GetBuffer() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBuffer()))
}
func (self Simple) SetBuffer(buffer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBuffer(gd.Int(buffer))
}
func (self Simple) GetByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetByteOffset()))
}
func (self Simple) SetByteOffset(byte_offset int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetByteOffset(gd.Int(byte_offset))
}
func (self Simple) GetByteLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetByteLength()))
}
func (self Simple) SetByteLength(byte_length int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetByteLength(gd.Int(byte_length))
}
func (self Simple) GetByteStride() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetByteStride()))
}
func (self Simple) SetByteStride(byte_stride int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetByteStride(gd.Int(byte_stride))
}
func (self Simple) GetIndices() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetIndices())
}
func (self Simple) SetIndices(indices bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndices(indices)
}
func (self Simple) GetVertexAttributes() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetVertexAttributes())
}
func (self Simple) SetVertexAttributes(is_attributes bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVertexAttributes(is_attributes)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.GLTFBufferView
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Loads the buffer view data from the buffer referenced by this buffer view in the given [GLTFState]. Interleaved data with a byte stride is not yet supported by this method. The data is returned as a [PackedByteArray].
*/
//go:nosplit
func (self class) LoadBufferViewData(ctx gd.Lifetime, state object.GLTFState) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(state[0].AsPointer())[0])
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_load_buffer_view_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBuffer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBuffer(buffer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, buffer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_set_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetByteOffset() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_get_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_set_byte_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetByteLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_get_byte_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetByteLength(byte_length gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, byte_length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_set_byte_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetByteStride() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_get_byte_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetByteStride(byte_stride gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, byte_stride)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_set_byte_stride, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIndices() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_get_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIndices(indices bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, indices)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVertexAttributes() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_get_vertex_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVertexAttributes(is_attributes bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, is_attributes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.GLTFBufferView.Bind_set_vertex_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsGLTFBufferView() Expert { return self[0].AsGLTFBufferView() }


//go:nosplit
func (self Simple) AsGLTFBufferView() Simple { return self[0].AsGLTFBufferView() }


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
func init() {classdb.Register("GLTFBufferView", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
