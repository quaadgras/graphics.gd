package StreamPeerBuffer

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/StreamPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A data buffer stream peer that uses a byte array as the stream. This object can be used to handle binary data from network sessions. To handle binary data stored in files, [FileAccess] can be used directly.
A [StreamPeerBuffer] object keeps an internal cursor which is the offset in bytes to the start of the buffer. Get and put operations are performed at the cursor position and will move the cursor accordingly.

*/
type Simple [1]classdb.StreamPeerBuffer
func (self Simple) SeekTo(position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SeekTo(gd.Int(position))
}
func (self Simple) GetSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSize()))
}
func (self Simple) GetPosition() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPosition()))
}
func (self Simple) Resize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Resize(gd.Int(size))
}
func (self Simple) SetDataArray(data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDataArray(gc.PackedByteSlice(data))
}
func (self Simple) GetDataArray() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetDataArray(gc).Bytes())
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) Duplicate() [1]classdb.StreamPeerBuffer {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StreamPeerBuffer(Expert(self).Duplicate(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StreamPeerBuffer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) Duplicate(ctx gd.Lifetime) object.StreamPeerBuffer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerBuffer.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StreamPeerBuffer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsStreamPeerBuffer() Expert { return self[0].AsStreamPeerBuffer() }


//go:nosplit
func (self Simple) AsStreamPeerBuffer() Simple { return self[0].AsStreamPeerBuffer() }


//go:nosplit
func (self class) AsStreamPeer() StreamPeer.Expert { return self[0].AsStreamPeer() }


//go:nosplit
func (self Simple) AsStreamPeer() StreamPeer.Simple { return self[0].AsStreamPeer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("StreamPeerBuffer", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
