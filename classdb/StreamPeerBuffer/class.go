// Package StreamPeerBuffer provides methods for working with StreamPeerBuffer object instances.
package StreamPeerBuffer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/StreamPeer"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
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
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A data buffer stream peer that uses a byte array as the stream. This object can be used to handle binary data from network sessions. To handle binary data stored in files, [FileAccess] can be used directly.
A [StreamPeerBuffer] object keeps an internal cursor which is the offset in bytes to the start of the buffer. Get and put operations are performed at the cursor position and will move the cursor accordingly.
*/
type Instance [1]gdclass.StreamPeerBuffer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStreamPeerBuffer() Instance
}

/*
Moves the cursor to the specified position. [param position] must be a valid index of [member data_array].
*/
func (self Instance) SeekTo(position int) { //gd:StreamPeerBuffer.seek
	Advanced(self).SeekTo(int64(position))
}

/*
Returns the size of [member data_array].
*/
func (self Instance) GetSize() int { //gd:StreamPeerBuffer.get_size
	return int(int(Advanced(self).GetSize()))
}

/*
Returns the current cursor position.
*/
func (self Instance) GetPosition() int { //gd:StreamPeerBuffer.get_position
	return int(int(Advanced(self).GetPosition()))
}

/*
Resizes the [member data_array]. This [i]doesn't[/i] update the cursor.
*/
func (self Instance) Resize(size int) { //gd:StreamPeerBuffer.resize
	Advanced(self).Resize(int64(size))
}

/*
Clears the [member data_array] and resets the cursor.
*/
func (self Instance) Clear() { //gd:StreamPeerBuffer.clear
	Advanced(self).Clear()
}

/*
Returns a new [StreamPeerBuffer] with the same [member data_array] content.
*/
func (self Instance) Duplicate() [1]gdclass.StreamPeerBuffer { //gd:StreamPeerBuffer.duplicate
	return [1]gdclass.StreamPeerBuffer(Advanced(self).Duplicate())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StreamPeerBuffer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerBuffer"))
	casted := Instance{*(*gdclass.StreamPeerBuffer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DataArray() []byte {
	return []byte(class(self).GetDataArray().Bytes())
}

func (self Instance) SetDataArray(value []byte) {
	class(self).SetDataArray(Packed.Bytes(Packed.New(value...)))
}

/*
Moves the cursor to the specified position. [param position] must be a valid index of [member data_array].
*/
//go:nosplit
func (self class) SeekTo(position int64) { //gd:StreamPeerBuffer.seek
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of [member data_array].
*/
//go:nosplit
func (self class) GetSize() int64 { //gd:StreamPeerBuffer.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current cursor position.
*/
//go:nosplit
func (self class) GetPosition() int64 { //gd:StreamPeerBuffer.get_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resizes the [member data_array]. This [i]doesn't[/i] update the cursor.
*/
//go:nosplit
func (self class) Resize(size int64) { //gd:StreamPeerBuffer.resize
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetDataArray(data Packed.Bytes) { //gd:StreamPeerBuffer.set_data_array
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_set_data_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDataArray() Packed.Bytes { //gd:StreamPeerBuffer.get_data_array
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_get_data_array, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Clears the [member data_array] and resets the cursor.
*/
//go:nosplit
func (self class) Clear() { //gd:StreamPeerBuffer.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a new [StreamPeerBuffer] with the same [member data_array] content.
*/
//go:nosplit
func (self class) Duplicate() [1]gdclass.StreamPeerBuffer { //gd:StreamPeerBuffer.duplicate
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StreamPeerBuffer{gd.PointerWithOwnershipTransferredToGo[gdclass.StreamPeerBuffer](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsStreamPeerBuffer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerBuffer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(StreamPeer.Advanced(self.AsStreamPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StreamPeer.Instance(self.AsStreamPeer()), name)
	}
}
func init() {
	gdclass.Register("StreamPeerBuffer", func(ptr gd.Object) any {
		return [1]gdclass.StreamPeerBuffer{*(*gdclass.StreamPeerBuffer)(unsafe.Pointer(&ptr))}
	})
}
