// Package AudioStreamPlaybackResampled provides methods for working with AudioStreamPlaybackResampled object instances.
package AudioStreamPlaybackResampled

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioStreamPlayback"
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

type Instance [1]gdclass.AudioStreamPlaybackResampled

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlaybackResampled() Instance
}
type Interface interface {
	MixResampled(dst_buffer *AudioFrame, frame_count int) int
	GetStreamSamplingRate() Float.X
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) MixResampled(dst_buffer *AudioFrame, frame_count int) (_ int) { return }
func (self implementation) GetStreamSamplingRate() (_ Float.X)                           { return }
func (Instance) _mix_resampled(impl func(ptr unsafe.Pointer, dst_buffer *AudioFrame, frame_count int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var dst_buffer = gd.UnsafeGet[*AudioFrame](p_args, 0)

		var frame_count = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, dst_buffer, int(frame_count))
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _get_stream_sampling_rate(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (self Instance) BeginResample() { //gd:AudioStreamPlaybackResampled.begin_resample
	Advanced(self).BeginResample()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlaybackResampled

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackResampled"))
	casted := Instance{*(*gdclass.AudioStreamPlaybackResampled)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _mix_resampled(impl func(ptr unsafe.Pointer, dst_buffer *AudioFrame, frame_count int64) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var dst_buffer = gd.UnsafeGet[*AudioFrame](p_args, 0)

		var frame_count = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, dst_buffer, frame_count)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_stream_sampling_rate(impl func(ptr unsafe.Pointer) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) BeginResample() { //gd:AudioStreamPlaybackResampled.begin_resample
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaybackResampled.Bind_begin_resample, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_mix_resampled":
		return reflect.ValueOf(self._mix_resampled)
	case "_get_stream_sampling_rate":
		return reflect.ValueOf(self._get_stream_sampling_rate)
	default:
		return gd.VirtualByName(AudioStreamPlayback.Advanced(self.AsAudioStreamPlayback()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_mix_resampled":
		return reflect.ValueOf(self._mix_resampled)
	case "_get_stream_sampling_rate":
		return reflect.ValueOf(self._get_stream_sampling_rate)
	default:
		return gd.VirtualByName(AudioStreamPlayback.Instance(self.AsAudioStreamPlayback()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlaybackResampled", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlaybackResampled{*(*gdclass.AudioStreamPlaybackResampled)(unsafe.Pointer(&ptr))}
	})
}
