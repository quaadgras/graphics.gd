package AudioStreamPlaybackResampled

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AudioStreamPlayback"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

type Simple [1]classdb.AudioStreamPlaybackResampled
func (Simple) _mix_resampled(impl func(ptr unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,0)
		var frame_count = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, dst_buffer, int(frame_count))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _get_stream_sampling_rate(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (self Simple) BeginResample() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BeginResample()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AudioStreamPlaybackResampled
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _mix_resampled(impl func(ptr unsafe.Pointer, dst_buffer *classdb.AudioFrame, frame_count gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var dst_buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args,0)
		var frame_count = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, dst_buffer, frame_count)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_stream_sampling_rate(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

//go:nosplit
func (self class) BeginResample()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaybackResampled.Bind_begin_resample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsAudioStreamPlaybackResampled() Expert { return self[0].AsAudioStreamPlaybackResampled() }


//go:nosplit
func (self Simple) AsAudioStreamPlaybackResampled() Simple { return self[0].AsAudioStreamPlaybackResampled() }


//go:nosplit
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Expert { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self Simple) AsAudioStreamPlayback() AudioStreamPlayback.Simple { return self[0].AsAudioStreamPlayback() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_mix_resampled": return reflect.ValueOf(self._mix_resampled);
	case "_get_stream_sampling_rate": return reflect.ValueOf(self._get_stream_sampling_rate);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_mix_resampled": return reflect.ValueOf(self._mix_resampled);
	case "_get_stream_sampling_rate": return reflect.ValueOf(self._get_stream_sampling_rate);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AudioStreamPlaybackResampled", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
