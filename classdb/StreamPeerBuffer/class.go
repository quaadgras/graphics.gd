// Package StreamPeerBuffer provides methods for working with StreamPeerBuffer object instances.
package StreamPeerBuffer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/StreamPeer"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A data buffer stream peer that uses a byte array as the stream. This object can be used to handle binary data from network sessions. To handle binary data stored in files, [FileAccess] can be used directly.
A [StreamPeerBuffer] object keeps an internal cursor which is the offset in bytes to the start of the buffer. Get and put operations are performed at the cursor position and will move the cursor accordingly.
*/
type Instance [1]gdclass.StreamPeerBuffer
type Any interface {
	gd.IsClass
	AsStreamPeerBuffer() Instance
}

/*
Moves the cursor to the specified position. [param position] must be a valid index of [member data_array].
*/
func (self Instance) SeekTo(position int) {
	class(self).SeekTo(gd.Int(position))
}

/*
Returns the size of [member data_array].
*/
func (self Instance) GetSize() int {
	return int(int(class(self).GetSize()))
}

/*
Returns the current cursor position.
*/
func (self Instance) GetPosition() int {
	return int(int(class(self).GetPosition()))
}

/*
Resizes the [member data_array]. This [i]doesn't[/i] update the cursor.
*/
func (self Instance) Resize(size int) {
	class(self).Resize(gd.Int(size))
}

/*
Clears the [member data_array] and resets the cursor.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns a new [StreamPeerBuffer] with the same [member data_array] content.
*/
func (self Instance) Duplicate() [1]gdclass.StreamPeerBuffer {
	return [1]gdclass.StreamPeerBuffer(class(self).Duplicate())
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
	class(self).SetDataArray(gd.NewPackedByteSlice(value))
}

/*
Moves the cursor to the specified position. [param position] must be a valid index of [member data_array].
*/
//go:nosplit
func (self class) SeekTo(position gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the size of [member data_array].
*/
//go:nosplit
func (self class) GetSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current cursor position.
*/
//go:nosplit
func (self class) GetPosition() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resizes the [member data_array]. This [i]doesn't[/i] update the cursor.
*/
//go:nosplit
func (self class) Resize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDataArray(data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_set_data_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDataArray() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_get_data_array, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the [member data_array] and resets the cursor.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a new [StreamPeerBuffer] with the same [member data_array] content.
*/
//go:nosplit
func (self class) Duplicate() [1]gdclass.StreamPeerBuffer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerBuffer.Bind_duplicate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
