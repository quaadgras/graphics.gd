package GLTFBufferView

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
GLTFBufferView is a data structure representing GLTF a [code]bufferView[/code] that would be found in the [code]"bufferViews"[/code] array. A buffer is a blob of binary data. A buffer view is a slice of a buffer that can be used to identify and extract data from the buffer.
Most custom uses of buffers only need to use the [member buffer], [member byte_length], and [member byte_offset]. The [member byte_stride] and [member indices] properties are for more advanced use cases such as interleaved mesh data encoded for the GPU.

*/
type Go [1]classdb.GLTFBufferView

/*
Loads the buffer view data from the buffer referenced by this buffer view in the given [GLTFState]. Interleaved data with a byte stride is not yet supported by this method. The data is returned as a [PackedByteArray].
*/
func (self Go) LoadBufferViewData(state gdclass.GLTFState) []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).LoadBufferViewData(gc, state).Bytes())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFBufferView
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("GLTFBufferView"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Buffer() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBuffer()))
}

func (self Go) SetBuffer(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBuffer(gd.Int(value))
}

func (self Go) ByteOffset() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetByteOffset()))
}

func (self Go) SetByteOffset(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetByteOffset(gd.Int(value))
}

func (self Go) ByteLength() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetByteLength()))
}

func (self Go) SetByteLength(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetByteLength(gd.Int(value))
}

func (self Go) ByteStride() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetByteStride()))
}

func (self Go) SetByteStride(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetByteStride(gd.Int(value))
}

func (self Go) Indices() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetIndices())
}

func (self Go) SetIndices(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndices(value)
}

func (self Go) VertexAttributes() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetVertexAttributes())
}

func (self Go) SetVertexAttributes(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVertexAttributes(value)
}

/*
Loads the buffer view data from the buffer referenced by this buffer view in the given [GLTFState]. Interleaved data with a byte stride is not yet supported by this method. The data is returned as a [PackedByteArray].
*/
//go:nosplit
func (self class) LoadBufferViewData(ctx gd.Lifetime, state gdclass.GLTFState) gd.PackedByteArray {
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
func (self class) AsGLTFBufferView() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFBufferView() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("GLTFBufferView", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
