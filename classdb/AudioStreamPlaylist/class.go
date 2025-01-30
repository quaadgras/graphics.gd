// Package AudioStreamPlaylist provides methods for working with AudioStreamPlaylist object instances.
package AudioStreamPlaylist

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
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

type Instance [1]gdclass.AudioStreamPlaylist

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlaylist() Instance
}

/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
func (self Instance) GetBpm() Float.X { //gd:AudioStreamPlaylist.get_bpm
	return Float.X(Float.X(class(self).GetBpm()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlaylist

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaylist"))
	casted := Instance{*(*gdclass.AudioStreamPlaylist)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Shuffle() bool {
	return bool(class(self).GetShuffle())
}

func (self Instance) SetShuffle(value bool) {
	class(self).SetShuffle(value)
}

func (self Instance) Loop() bool {
	return bool(class(self).HasLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

func (self Instance) FadeTime() Float.X {
	return Float.X(Float.X(class(self).GetFadeTime()))
}

func (self Instance) SetFadeTime(value Float.X) {
	class(self).SetFadeTime(float64(value))
}

func (self Instance) StreamCount() int {
	return int(int(class(self).GetStreamCount()))
}

func (self Instance) SetStreamCount(value int) {
	class(self).SetStreamCount(int64(value))
}

func (self Instance) Stream0() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(0))
}

func (self Instance) SetStream0(value [1]gdclass.AudioStream) {
	class(self).SetListStream(0, value)
}

func (self Instance) Stream1() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(1))
}

func (self Instance) SetStream1(value [1]gdclass.AudioStream) {
	class(self).SetListStream(1, value)
}

func (self Instance) Stream2() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(2))
}

func (self Instance) SetStream2(value [1]gdclass.AudioStream) {
	class(self).SetListStream(2, value)
}

func (self Instance) Stream3() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(3))
}

func (self Instance) SetStream3(value [1]gdclass.AudioStream) {
	class(self).SetListStream(3, value)
}

func (self Instance) Stream4() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(4))
}

func (self Instance) SetStream4(value [1]gdclass.AudioStream) {
	class(self).SetListStream(4, value)
}

func (self Instance) Stream5() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(5))
}

func (self Instance) SetStream5(value [1]gdclass.AudioStream) {
	class(self).SetListStream(5, value)
}

func (self Instance) Stream6() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(6))
}

func (self Instance) SetStream6(value [1]gdclass.AudioStream) {
	class(self).SetListStream(6, value)
}

func (self Instance) Stream7() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(7))
}

func (self Instance) SetStream7(value [1]gdclass.AudioStream) {
	class(self).SetListStream(7, value)
}

func (self Instance) Stream8() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(8))
}

func (self Instance) SetStream8(value [1]gdclass.AudioStream) {
	class(self).SetListStream(8, value)
}

func (self Instance) Stream9() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(9))
}

func (self Instance) SetStream9(value [1]gdclass.AudioStream) {
	class(self).SetListStream(9, value)
}

func (self Instance) Stream10() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(10))
}

func (self Instance) SetStream10(value [1]gdclass.AudioStream) {
	class(self).SetListStream(10, value)
}

func (self Instance) Stream11() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(11))
}

func (self Instance) SetStream11(value [1]gdclass.AudioStream) {
	class(self).SetListStream(11, value)
}

func (self Instance) Stream12() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(12))
}

func (self Instance) SetStream12(value [1]gdclass.AudioStream) {
	class(self).SetListStream(12, value)
}

func (self Instance) Stream13() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(13))
}

func (self Instance) SetStream13(value [1]gdclass.AudioStream) {
	class(self).SetListStream(13, value)
}

func (self Instance) Stream14() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(14))
}

func (self Instance) SetStream14(value [1]gdclass.AudioStream) {
	class(self).SetListStream(14, value)
}

func (self Instance) Stream15() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(15))
}

func (self Instance) SetStream15(value [1]gdclass.AudioStream) {
	class(self).SetListStream(15, value)
}

func (self Instance) Stream16() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(16))
}

func (self Instance) SetStream16(value [1]gdclass.AudioStream) {
	class(self).SetListStream(16, value)
}

func (self Instance) Stream17() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(17))
}

func (self Instance) SetStream17(value [1]gdclass.AudioStream) {
	class(self).SetListStream(17, value)
}

func (self Instance) Stream18() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(18))
}

func (self Instance) SetStream18(value [1]gdclass.AudioStream) {
	class(self).SetListStream(18, value)
}

func (self Instance) Stream19() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(19))
}

func (self Instance) SetStream19(value [1]gdclass.AudioStream) {
	class(self).SetListStream(19, value)
}

func (self Instance) Stream20() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(20))
}

func (self Instance) SetStream20(value [1]gdclass.AudioStream) {
	class(self).SetListStream(20, value)
}

func (self Instance) Stream21() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(21))
}

func (self Instance) SetStream21(value [1]gdclass.AudioStream) {
	class(self).SetListStream(21, value)
}

func (self Instance) Stream22() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(22))
}

func (self Instance) SetStream22(value [1]gdclass.AudioStream) {
	class(self).SetListStream(22, value)
}

func (self Instance) Stream23() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(23))
}

func (self Instance) SetStream23(value [1]gdclass.AudioStream) {
	class(self).SetListStream(23, value)
}

func (self Instance) Stream24() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(24))
}

func (self Instance) SetStream24(value [1]gdclass.AudioStream) {
	class(self).SetListStream(24, value)
}

func (self Instance) Stream25() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(25))
}

func (self Instance) SetStream25(value [1]gdclass.AudioStream) {
	class(self).SetListStream(25, value)
}

func (self Instance) Stream26() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(26))
}

func (self Instance) SetStream26(value [1]gdclass.AudioStream) {
	class(self).SetListStream(26, value)
}

func (self Instance) Stream27() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(27))
}

func (self Instance) SetStream27(value [1]gdclass.AudioStream) {
	class(self).SetListStream(27, value)
}

func (self Instance) Stream28() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(28))
}

func (self Instance) SetStream28(value [1]gdclass.AudioStream) {
	class(self).SetListStream(28, value)
}

func (self Instance) Stream29() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(29))
}

func (self Instance) SetStream29(value [1]gdclass.AudioStream) {
	class(self).SetListStream(29, value)
}

func (self Instance) Stream30() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(30))
}

func (self Instance) SetStream30(value [1]gdclass.AudioStream) {
	class(self).SetListStream(30, value)
}

func (self Instance) Stream31() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(31))
}

func (self Instance) SetStream31(value [1]gdclass.AudioStream) {
	class(self).SetListStream(31, value)
}

func (self Instance) Stream32() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(32))
}

func (self Instance) SetStream32(value [1]gdclass.AudioStream) {
	class(self).SetListStream(32, value)
}

func (self Instance) Stream33() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(33))
}

func (self Instance) SetStream33(value [1]gdclass.AudioStream) {
	class(self).SetListStream(33, value)
}

func (self Instance) Stream34() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(34))
}

func (self Instance) SetStream34(value [1]gdclass.AudioStream) {
	class(self).SetListStream(34, value)
}

func (self Instance) Stream35() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(35))
}

func (self Instance) SetStream35(value [1]gdclass.AudioStream) {
	class(self).SetListStream(35, value)
}

func (self Instance) Stream36() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(36))
}

func (self Instance) SetStream36(value [1]gdclass.AudioStream) {
	class(self).SetListStream(36, value)
}

func (self Instance) Stream37() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(37))
}

func (self Instance) SetStream37(value [1]gdclass.AudioStream) {
	class(self).SetListStream(37, value)
}

func (self Instance) Stream38() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(38))
}

func (self Instance) SetStream38(value [1]gdclass.AudioStream) {
	class(self).SetListStream(38, value)
}

func (self Instance) Stream39() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(39))
}

func (self Instance) SetStream39(value [1]gdclass.AudioStream) {
	class(self).SetListStream(39, value)
}

func (self Instance) Stream40() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(40))
}

func (self Instance) SetStream40(value [1]gdclass.AudioStream) {
	class(self).SetListStream(40, value)
}

func (self Instance) Stream41() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(41))
}

func (self Instance) SetStream41(value [1]gdclass.AudioStream) {
	class(self).SetListStream(41, value)
}

func (self Instance) Stream42() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(42))
}

func (self Instance) SetStream42(value [1]gdclass.AudioStream) {
	class(self).SetListStream(42, value)
}

func (self Instance) Stream43() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(43))
}

func (self Instance) SetStream43(value [1]gdclass.AudioStream) {
	class(self).SetListStream(43, value)
}

func (self Instance) Stream44() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(44))
}

func (self Instance) SetStream44(value [1]gdclass.AudioStream) {
	class(self).SetListStream(44, value)
}

func (self Instance) Stream45() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(45))
}

func (self Instance) SetStream45(value [1]gdclass.AudioStream) {
	class(self).SetListStream(45, value)
}

func (self Instance) Stream46() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(46))
}

func (self Instance) SetStream46(value [1]gdclass.AudioStream) {
	class(self).SetListStream(46, value)
}

func (self Instance) Stream47() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(47))
}

func (self Instance) SetStream47(value [1]gdclass.AudioStream) {
	class(self).SetListStream(47, value)
}

func (self Instance) Stream48() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(48))
}

func (self Instance) SetStream48(value [1]gdclass.AudioStream) {
	class(self).SetListStream(48, value)
}

func (self Instance) Stream49() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(49))
}

func (self Instance) SetStream49(value [1]gdclass.AudioStream) {
	class(self).SetListStream(49, value)
}

func (self Instance) Stream50() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(50))
}

func (self Instance) SetStream50(value [1]gdclass.AudioStream) {
	class(self).SetListStream(50, value)
}

func (self Instance) Stream51() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(51))
}

func (self Instance) SetStream51(value [1]gdclass.AudioStream) {
	class(self).SetListStream(51, value)
}

func (self Instance) Stream52() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(52))
}

func (self Instance) SetStream52(value [1]gdclass.AudioStream) {
	class(self).SetListStream(52, value)
}

func (self Instance) Stream53() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(53))
}

func (self Instance) SetStream53(value [1]gdclass.AudioStream) {
	class(self).SetListStream(53, value)
}

func (self Instance) Stream54() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(54))
}

func (self Instance) SetStream54(value [1]gdclass.AudioStream) {
	class(self).SetListStream(54, value)
}

func (self Instance) Stream55() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(55))
}

func (self Instance) SetStream55(value [1]gdclass.AudioStream) {
	class(self).SetListStream(55, value)
}

func (self Instance) Stream56() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(56))
}

func (self Instance) SetStream56(value [1]gdclass.AudioStream) {
	class(self).SetListStream(56, value)
}

func (self Instance) Stream57() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(57))
}

func (self Instance) SetStream57(value [1]gdclass.AudioStream) {
	class(self).SetListStream(57, value)
}

func (self Instance) Stream58() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(58))
}

func (self Instance) SetStream58(value [1]gdclass.AudioStream) {
	class(self).SetListStream(58, value)
}

func (self Instance) Stream59() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(59))
}

func (self Instance) SetStream59(value [1]gdclass.AudioStream) {
	class(self).SetListStream(59, value)
}

func (self Instance) Stream60() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(60))
}

func (self Instance) SetStream60(value [1]gdclass.AudioStream) {
	class(self).SetListStream(60, value)
}

func (self Instance) Stream61() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(61))
}

func (self Instance) SetStream61(value [1]gdclass.AudioStream) {
	class(self).SetListStream(61, value)
}

func (self Instance) Stream62() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(62))
}

func (self Instance) SetStream62(value [1]gdclass.AudioStream) {
	class(self).SetListStream(62, value)
}

func (self Instance) Stream63() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetListStream(63))
}

func (self Instance) SetStream63(value [1]gdclass.AudioStream) {
	class(self).SetListStream(63, value)
}

//go:nosplit
func (self class) SetStreamCount(stream_count int64) { //gd:AudioStreamPlaylist.set_stream_count
	var frame = callframe.New()
	callframe.Arg(frame, stream_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamCount() int64 { //gd:AudioStreamPlaylist.get_stream_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
//go:nosplit
func (self class) GetBpm() float64 { //gd:AudioStreamPlaylist.get_bpm
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the stream at playback position index.
*/
//go:nosplit
func (self class) SetListStream(stream_index int64, audio_stream [1]gdclass.AudioStream) { //gd:AudioStreamPlaylist.set_list_stream
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, pointers.Get(audio_stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_list_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the stream at playback position index.
*/
//go:nosplit
func (self class) GetListStream(stream_index int64) [1]gdclass.AudioStream { //gd:AudioStreamPlaylist.get_list_stream
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_list_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShuffle(shuffle bool) { //gd:AudioStreamPlaylist.set_shuffle
	var frame = callframe.New()
	callframe.Arg(frame, shuffle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_shuffle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShuffle() bool { //gd:AudioStreamPlaylist.get_shuffle
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_shuffle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeTime(dec float64) { //gd:AudioStreamPlaylist.set_fade_time
	var frame = callframe.New()
	callframe.Arg(frame, dec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_fade_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeTime() float64 { //gd:AudioStreamPlaylist.get_fade_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_fade_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(loop bool) { //gd:AudioStreamPlaylist.set_loop
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool { //gd:AudioStreamPlaylist.has_loop
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPlaylist() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlaylist() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.Advanced {
	return *((*AudioStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStream() AudioStream.Instance {
	return *((*AudioStream.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioStream.Advanced(self.AsAudioStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Instance(self.AsAudioStream()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlaylist", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlaylist{*(*gdclass.AudioStreamPlaylist)(unsafe.Pointer(&ptr))}
	})
}
