package AudioStreamPlaylist

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStream"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

type Go [1]classdb.AudioStreamPlaylist

/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
func (self Go) GetBpm() float64 {
	return float64(float64(class(self).GetBpm()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlaylist
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaylist"))
	return Go{classdb.AudioStreamPlaylist(object)}
}

func (self Go) Shuffle() bool {
		return bool(class(self).GetShuffle())
}

func (self Go) SetShuffle(value bool) {
	class(self).SetShuffle(value)
}

func (self Go) Loop() bool {
		return bool(class(self).HasLoop())
}

func (self Go) SetLoop(value bool) {
	class(self).SetLoop(value)
}

func (self Go) FadeTime() float64 {
		return float64(float64(class(self).GetFadeTime()))
}

func (self Go) SetFadeTime(value float64) {
	class(self).SetFadeTime(gd.Float(value))
}

func (self Go) StreamCount() int {
		return int(int(class(self).GetStreamCount()))
}

func (self Go) SetStreamCount(value int) {
	class(self).SetStreamCount(gd.Int(value))
}

func (self Go) Stream0() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(0))
}

func (self Go) SetStream0(value gdclass.AudioStream) {
	class(self).SetListStream(0, value)
}

func (self Go) Stream1() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(1))
}

func (self Go) SetStream1(value gdclass.AudioStream) {
	class(self).SetListStream(1, value)
}

func (self Go) Stream2() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(2))
}

func (self Go) SetStream2(value gdclass.AudioStream) {
	class(self).SetListStream(2, value)
}

func (self Go) Stream3() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(3))
}

func (self Go) SetStream3(value gdclass.AudioStream) {
	class(self).SetListStream(3, value)
}

func (self Go) Stream4() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(4))
}

func (self Go) SetStream4(value gdclass.AudioStream) {
	class(self).SetListStream(4, value)
}

func (self Go) Stream5() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(5))
}

func (self Go) SetStream5(value gdclass.AudioStream) {
	class(self).SetListStream(5, value)
}

func (self Go) Stream6() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(6))
}

func (self Go) SetStream6(value gdclass.AudioStream) {
	class(self).SetListStream(6, value)
}

func (self Go) Stream7() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(7))
}

func (self Go) SetStream7(value gdclass.AudioStream) {
	class(self).SetListStream(7, value)
}

func (self Go) Stream8() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(8))
}

func (self Go) SetStream8(value gdclass.AudioStream) {
	class(self).SetListStream(8, value)
}

func (self Go) Stream9() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(9))
}

func (self Go) SetStream9(value gdclass.AudioStream) {
	class(self).SetListStream(9, value)
}

func (self Go) Stream10() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(10))
}

func (self Go) SetStream10(value gdclass.AudioStream) {
	class(self).SetListStream(10, value)
}

func (self Go) Stream11() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(11))
}

func (self Go) SetStream11(value gdclass.AudioStream) {
	class(self).SetListStream(11, value)
}

func (self Go) Stream12() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(12))
}

func (self Go) SetStream12(value gdclass.AudioStream) {
	class(self).SetListStream(12, value)
}

func (self Go) Stream13() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(13))
}

func (self Go) SetStream13(value gdclass.AudioStream) {
	class(self).SetListStream(13, value)
}

func (self Go) Stream14() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(14))
}

func (self Go) SetStream14(value gdclass.AudioStream) {
	class(self).SetListStream(14, value)
}

func (self Go) Stream15() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(15))
}

func (self Go) SetStream15(value gdclass.AudioStream) {
	class(self).SetListStream(15, value)
}

func (self Go) Stream16() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(16))
}

func (self Go) SetStream16(value gdclass.AudioStream) {
	class(self).SetListStream(16, value)
}

func (self Go) Stream17() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(17))
}

func (self Go) SetStream17(value gdclass.AudioStream) {
	class(self).SetListStream(17, value)
}

func (self Go) Stream18() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(18))
}

func (self Go) SetStream18(value gdclass.AudioStream) {
	class(self).SetListStream(18, value)
}

func (self Go) Stream19() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(19))
}

func (self Go) SetStream19(value gdclass.AudioStream) {
	class(self).SetListStream(19, value)
}

func (self Go) Stream20() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(20))
}

func (self Go) SetStream20(value gdclass.AudioStream) {
	class(self).SetListStream(20, value)
}

func (self Go) Stream21() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(21))
}

func (self Go) SetStream21(value gdclass.AudioStream) {
	class(self).SetListStream(21, value)
}

func (self Go) Stream22() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(22))
}

func (self Go) SetStream22(value gdclass.AudioStream) {
	class(self).SetListStream(22, value)
}

func (self Go) Stream23() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(23))
}

func (self Go) SetStream23(value gdclass.AudioStream) {
	class(self).SetListStream(23, value)
}

func (self Go) Stream24() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(24))
}

func (self Go) SetStream24(value gdclass.AudioStream) {
	class(self).SetListStream(24, value)
}

func (self Go) Stream25() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(25))
}

func (self Go) SetStream25(value gdclass.AudioStream) {
	class(self).SetListStream(25, value)
}

func (self Go) Stream26() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(26))
}

func (self Go) SetStream26(value gdclass.AudioStream) {
	class(self).SetListStream(26, value)
}

func (self Go) Stream27() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(27))
}

func (self Go) SetStream27(value gdclass.AudioStream) {
	class(self).SetListStream(27, value)
}

func (self Go) Stream28() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(28))
}

func (self Go) SetStream28(value gdclass.AudioStream) {
	class(self).SetListStream(28, value)
}

func (self Go) Stream29() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(29))
}

func (self Go) SetStream29(value gdclass.AudioStream) {
	class(self).SetListStream(29, value)
}

func (self Go) Stream30() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(30))
}

func (self Go) SetStream30(value gdclass.AudioStream) {
	class(self).SetListStream(30, value)
}

func (self Go) Stream31() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(31))
}

func (self Go) SetStream31(value gdclass.AudioStream) {
	class(self).SetListStream(31, value)
}

func (self Go) Stream32() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(32))
}

func (self Go) SetStream32(value gdclass.AudioStream) {
	class(self).SetListStream(32, value)
}

func (self Go) Stream33() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(33))
}

func (self Go) SetStream33(value gdclass.AudioStream) {
	class(self).SetListStream(33, value)
}

func (self Go) Stream34() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(34))
}

func (self Go) SetStream34(value gdclass.AudioStream) {
	class(self).SetListStream(34, value)
}

func (self Go) Stream35() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(35))
}

func (self Go) SetStream35(value gdclass.AudioStream) {
	class(self).SetListStream(35, value)
}

func (self Go) Stream36() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(36))
}

func (self Go) SetStream36(value gdclass.AudioStream) {
	class(self).SetListStream(36, value)
}

func (self Go) Stream37() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(37))
}

func (self Go) SetStream37(value gdclass.AudioStream) {
	class(self).SetListStream(37, value)
}

func (self Go) Stream38() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(38))
}

func (self Go) SetStream38(value gdclass.AudioStream) {
	class(self).SetListStream(38, value)
}

func (self Go) Stream39() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(39))
}

func (self Go) SetStream39(value gdclass.AudioStream) {
	class(self).SetListStream(39, value)
}

func (self Go) Stream40() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(40))
}

func (self Go) SetStream40(value gdclass.AudioStream) {
	class(self).SetListStream(40, value)
}

func (self Go) Stream41() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(41))
}

func (self Go) SetStream41(value gdclass.AudioStream) {
	class(self).SetListStream(41, value)
}

func (self Go) Stream42() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(42))
}

func (self Go) SetStream42(value gdclass.AudioStream) {
	class(self).SetListStream(42, value)
}

func (self Go) Stream43() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(43))
}

func (self Go) SetStream43(value gdclass.AudioStream) {
	class(self).SetListStream(43, value)
}

func (self Go) Stream44() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(44))
}

func (self Go) SetStream44(value gdclass.AudioStream) {
	class(self).SetListStream(44, value)
}

func (self Go) Stream45() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(45))
}

func (self Go) SetStream45(value gdclass.AudioStream) {
	class(self).SetListStream(45, value)
}

func (self Go) Stream46() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(46))
}

func (self Go) SetStream46(value gdclass.AudioStream) {
	class(self).SetListStream(46, value)
}

func (self Go) Stream47() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(47))
}

func (self Go) SetStream47(value gdclass.AudioStream) {
	class(self).SetListStream(47, value)
}

func (self Go) Stream48() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(48))
}

func (self Go) SetStream48(value gdclass.AudioStream) {
	class(self).SetListStream(48, value)
}

func (self Go) Stream49() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(49))
}

func (self Go) SetStream49(value gdclass.AudioStream) {
	class(self).SetListStream(49, value)
}

func (self Go) Stream50() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(50))
}

func (self Go) SetStream50(value gdclass.AudioStream) {
	class(self).SetListStream(50, value)
}

func (self Go) Stream51() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(51))
}

func (self Go) SetStream51(value gdclass.AudioStream) {
	class(self).SetListStream(51, value)
}

func (self Go) Stream52() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(52))
}

func (self Go) SetStream52(value gdclass.AudioStream) {
	class(self).SetListStream(52, value)
}

func (self Go) Stream53() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(53))
}

func (self Go) SetStream53(value gdclass.AudioStream) {
	class(self).SetListStream(53, value)
}

func (self Go) Stream54() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(54))
}

func (self Go) SetStream54(value gdclass.AudioStream) {
	class(self).SetListStream(54, value)
}

func (self Go) Stream55() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(55))
}

func (self Go) SetStream55(value gdclass.AudioStream) {
	class(self).SetListStream(55, value)
}

func (self Go) Stream56() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(56))
}

func (self Go) SetStream56(value gdclass.AudioStream) {
	class(self).SetListStream(56, value)
}

func (self Go) Stream57() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(57))
}

func (self Go) SetStream57(value gdclass.AudioStream) {
	class(self).SetListStream(57, value)
}

func (self Go) Stream58() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(58))
}

func (self Go) SetStream58(value gdclass.AudioStream) {
	class(self).SetListStream(58, value)
}

func (self Go) Stream59() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(59))
}

func (self Go) SetStream59(value gdclass.AudioStream) {
	class(self).SetListStream(59, value)
}

func (self Go) Stream60() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(60))
}

func (self Go) SetStream60(value gdclass.AudioStream) {
	class(self).SetListStream(60, value)
}

func (self Go) Stream61() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(61))
}

func (self Go) SetStream61(value gdclass.AudioStream) {
	class(self).SetListStream(61, value)
}

func (self Go) Stream62() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(62))
}

func (self Go) SetStream62(value gdclass.AudioStream) {
	class(self).SetListStream(62, value)
}

func (self Go) Stream63() gdclass.AudioStream {
		return gdclass.AudioStream(class(self).GetListStream(63))
}

func (self Go) SetStream63(value gdclass.AudioStream) {
	class(self).SetListStream(63, value)
}

//go:nosplit
func (self class) SetStreamCount(stream_count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, stream_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
//go:nosplit
func (self class) GetBpm() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the stream at playback position index.
*/
//go:nosplit
func (self class) SetListStream(stream_index gd.Int, audio_stream gdclass.AudioStream)  {
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, discreet.Get(audio_stream[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_list_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stream at playback position index.
*/
//go:nosplit
func (self class) GetListStream(stream_index gd.Int) gdclass.AudioStream {
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_list_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.AudioStream{classdb.AudioStream(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShuffle(shuffle bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, shuffle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_shuffle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShuffle() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_shuffle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFadeTime(dec gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, dec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_fade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFadeTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_get_fade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoop(loop bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasLoop() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlaylist.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPlaylist() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStreamPlaylist() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioStream() AudioStream.GD { return *((*AudioStream.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioStream() AudioStream.Go { return *((*AudioStream.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {classdb.Register("AudioStreamPlaylist", func(ptr gd.Object) any { return classdb.AudioStreamPlaylist(ptr) })}
