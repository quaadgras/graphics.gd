package StreamPeerBuffer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/StreamPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A data buffer stream peer that uses a byte array as the stream. This object can be used to handle binary data from network sessions. To handle binary data stored in files, [FileAccess] can be used directly.
A [StreamPeerBuffer] object keeps an internal cursor which is the offset in bytes to the start of the buffer. Get and put operations are performed at the cursor position and will move the cursor accordingly.

*/
type Go [1]classdb.StreamPeerBuffer

/*
Moves the cursor to the specified position. [param position] must be a valid index of [member data_array].
*/
func (self Go) SeekTo(position int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SeekTo(gd.Int(position))
}

/*
Returns the size of [member data_array].
*/
func (self Go) GetSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSize()))
}

/*
Returns the current cursor position.
*/
func (self Go) GetPosition() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPosition()))
}

/*
Resizes the [member data_array]. This [i]doesn't[/i] update the cursor.
*/
func (self Go) Resize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Resize(gd.Int(size))
}

/*
Clears the [member data_array] and resets the cursor.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Returns a new [StreamPeerBuffer] with the same [member data_array] content.
*/
func (self Go) Duplicate() gdclass.StreamPeerBuffer {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.StreamPeerBuffer(class(self).Duplicate(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StreamPeerBuffer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StreamPeerBuffer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) DataArray() []byte {
	gc := gd.GarbageCollector(); _ = gc
		return []byte(class(self).GetDataArray(gc).Bytes())
}

func (self Go) SetDataArray(value []byte) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDataArray(gc.PackedByteSlice(value))
}

/*
Moves the cursor to the specified position. [param position] must be a valid index of [member data_array].
*/
//go:nosplit
func (self class) SeekTo(position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of [member data_array].
*/
//go:nosplit
func (self class) GetSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current cursor position.
*/
//go:nosplit
func (self class) GetPosition() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resizes the [member data_array]. This [i]doesn't[/i] update the cursor.
*/
//go:nosplit
func (self class) Resize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDataArray(data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_set_data_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDataArray(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_get_data_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears the [member data_array] and resets the cursor.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a new [StreamPeerBuffer] with the same [member data_array] content.
*/
//go:nosplit
func (self class) Duplicate(ctx gd.Lifetime) gdclass.StreamPeerBuffer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.StreamPeerBuffer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsStreamPeerBuffer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStreamPeerBuffer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.GD { return *((*StreamPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsStreamPeer() StreamPeer.Go { return *((*StreamPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}
func init() {classdb.Register("StreamPeerBuffer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
