package AudioStreamPlaybackResampled

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AudioStreamPlayback"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]classdb.AudioStreamPlaybackResampled
type Any interface {
	gd.IsClass
	AsAudioStreamPlaybackResampled() Instance
}

func (Instance) _mix_resampled(impl func(ptr unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args, 0)
		var frame_count = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, dst_buffer, int(frame_count))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _get_stream_sampling_rate(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (self Instance) BeginResample() {
	class(self).BeginResample()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamPlaybackResampled

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackResampled"))
	return Instance{*(*classdb.AudioStreamPlaybackResampled)(unsafe.Pointer(&object))}
}

func (class) _mix_resampled(impl func(ptr unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args, 0)
		var frame_count = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, dst_buffer, frame_count)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_stream_sampling_rate(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) BeginResample() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackResampled.Bind_begin_resample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsAudioStreamPlaybackResampled() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlaybackResampled() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Advanced {
	return *((*AudioStreamPlayback.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlayback() AudioStreamPlayback.Instance {
	return *((*AudioStreamPlayback.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_mix_resampled":
		return reflect.ValueOf(self._mix_resampled)
	case "_get_stream_sampling_rate":
		return reflect.ValueOf(self._get_stream_sampling_rate)
	default:
		return gd.VirtualByName(self.AsAudioStreamPlayback(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_mix_resampled":
		return reflect.ValueOf(self._mix_resampled)
	case "_get_stream_sampling_rate":
		return reflect.ValueOf(self._get_stream_sampling_rate)
	default:
		return gd.VirtualByName(self.AsAudioStreamPlayback(), name)
	}
}
func init() {
	classdb.Register("AudioStreamPlaybackResampled", func(ptr gd.Object) any {
		return [1]classdb.AudioStreamPlaybackResampled{*(*classdb.AudioStreamPlaybackResampled)(unsafe.Pointer(&ptr))}
	})
}
