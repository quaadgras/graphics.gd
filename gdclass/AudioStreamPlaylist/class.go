package AudioStreamPlaylist

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

type Go [1]classdb.AudioStreamPlaylist

/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
func (self Go) GetBpm() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetBpm()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioStreamPlaylist
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioStreamPlaylist"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Shuffle() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetShuffle())
}

func (self Go) SetShuffle(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShuffle(value)
}

func (self Go) Loop() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasLoop())
}

func (self Go) SetLoop(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLoop(value)
}

func (self Go) FadeTime() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetFadeTime()))
}

func (self Go) SetFadeTime(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFadeTime(gd.Float(value))
}

func (self Go) StreamCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetStreamCount()))
}

func (self Go) SetStreamCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStreamCount(gd.Int(value))
}

func (self Go) Stream0() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,0))
}

func (self Go) SetStream0(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(0, value)
}

func (self Go) Stream1() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,1))
}

func (self Go) SetStream1(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(1, value)
}

func (self Go) Stream2() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,2))
}

func (self Go) SetStream2(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(2, value)
}

func (self Go) Stream3() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,3))
}

func (self Go) SetStream3(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(3, value)
}

func (self Go) Stream4() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,4))
}

func (self Go) SetStream4(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(4, value)
}

func (self Go) Stream5() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,5))
}

func (self Go) SetStream5(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(5, value)
}

func (self Go) Stream6() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,6))
}

func (self Go) SetStream6(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(6, value)
}

func (self Go) Stream7() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,7))
}

func (self Go) SetStream7(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(7, value)
}

func (self Go) Stream8() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,8))
}

func (self Go) SetStream8(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(8, value)
}

func (self Go) Stream9() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,9))
}

func (self Go) SetStream9(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(9, value)
}

func (self Go) Stream10() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,10))
}

func (self Go) SetStream10(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(10, value)
}

func (self Go) Stream11() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,11))
}

func (self Go) SetStream11(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(11, value)
}

func (self Go) Stream12() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,12))
}

func (self Go) SetStream12(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(12, value)
}

func (self Go) Stream13() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,13))
}

func (self Go) SetStream13(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(13, value)
}

func (self Go) Stream14() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,14))
}

func (self Go) SetStream14(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(14, value)
}

func (self Go) Stream15() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,15))
}

func (self Go) SetStream15(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(15, value)
}

func (self Go) Stream16() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,16))
}

func (self Go) SetStream16(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(16, value)
}

func (self Go) Stream17() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,17))
}

func (self Go) SetStream17(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(17, value)
}

func (self Go) Stream18() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,18))
}

func (self Go) SetStream18(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(18, value)
}

func (self Go) Stream19() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,19))
}

func (self Go) SetStream19(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(19, value)
}

func (self Go) Stream20() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,20))
}

func (self Go) SetStream20(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(20, value)
}

func (self Go) Stream21() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,21))
}

func (self Go) SetStream21(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(21, value)
}

func (self Go) Stream22() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,22))
}

func (self Go) SetStream22(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(22, value)
}

func (self Go) Stream23() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,23))
}

func (self Go) SetStream23(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(23, value)
}

func (self Go) Stream24() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,24))
}

func (self Go) SetStream24(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(24, value)
}

func (self Go) Stream25() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,25))
}

func (self Go) SetStream25(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(25, value)
}

func (self Go) Stream26() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,26))
}

func (self Go) SetStream26(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(26, value)
}

func (self Go) Stream27() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,27))
}

func (self Go) SetStream27(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(27, value)
}

func (self Go) Stream28() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,28))
}

func (self Go) SetStream28(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(28, value)
}

func (self Go) Stream29() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,29))
}

func (self Go) SetStream29(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(29, value)
}

func (self Go) Stream30() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,30))
}

func (self Go) SetStream30(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(30, value)
}

func (self Go) Stream31() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,31))
}

func (self Go) SetStream31(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(31, value)
}

func (self Go) Stream32() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,32))
}

func (self Go) SetStream32(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(32, value)
}

func (self Go) Stream33() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,33))
}

func (self Go) SetStream33(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(33, value)
}

func (self Go) Stream34() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,34))
}

func (self Go) SetStream34(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(34, value)
}

func (self Go) Stream35() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,35))
}

func (self Go) SetStream35(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(35, value)
}

func (self Go) Stream36() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,36))
}

func (self Go) SetStream36(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(36, value)
}

func (self Go) Stream37() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,37))
}

func (self Go) SetStream37(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(37, value)
}

func (self Go) Stream38() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,38))
}

func (self Go) SetStream38(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(38, value)
}

func (self Go) Stream39() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,39))
}

func (self Go) SetStream39(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(39, value)
}

func (self Go) Stream40() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,40))
}

func (self Go) SetStream40(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(40, value)
}

func (self Go) Stream41() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,41))
}

func (self Go) SetStream41(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(41, value)
}

func (self Go) Stream42() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,42))
}

func (self Go) SetStream42(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(42, value)
}

func (self Go) Stream43() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,43))
}

func (self Go) SetStream43(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(43, value)
}

func (self Go) Stream44() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,44))
}

func (self Go) SetStream44(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(44, value)
}

func (self Go) Stream45() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,45))
}

func (self Go) SetStream45(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(45, value)
}

func (self Go) Stream46() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,46))
}

func (self Go) SetStream46(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(46, value)
}

func (self Go) Stream47() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,47))
}

func (self Go) SetStream47(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(47, value)
}

func (self Go) Stream48() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,48))
}

func (self Go) SetStream48(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(48, value)
}

func (self Go) Stream49() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,49))
}

func (self Go) SetStream49(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(49, value)
}

func (self Go) Stream50() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,50))
}

func (self Go) SetStream50(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(50, value)
}

func (self Go) Stream51() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,51))
}

func (self Go) SetStream51(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(51, value)
}

func (self Go) Stream52() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,52))
}

func (self Go) SetStream52(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(52, value)
}

func (self Go) Stream53() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,53))
}

func (self Go) SetStream53(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(53, value)
}

func (self Go) Stream54() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,54))
}

func (self Go) SetStream54(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(54, value)
}

func (self Go) Stream55() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,55))
}

func (self Go) SetStream55(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(55, value)
}

func (self Go) Stream56() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,56))
}

func (self Go) SetStream56(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(56, value)
}

func (self Go) Stream57() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,57))
}

func (self Go) SetStream57(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(57, value)
}

func (self Go) Stream58() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,58))
}

func (self Go) SetStream58(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(58, value)
}

func (self Go) Stream59() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,59))
}

func (self Go) SetStream59(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(59, value)
}

func (self Go) Stream60() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,60))
}

func (self Go) SetStream60(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(60, value)
}

func (self Go) Stream61() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,61))
}

func (self Go) SetStream61(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(61, value)
}

func (self Go) Stream62() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,62))
}

func (self Go) SetStream62(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(62, value)
}

func (self Go) Stream63() gdclass.AudioStream {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.AudioStream(class(self).GetListStream(gc,63))
}

func (self Go) SetStream63(value gdclass.AudioStream) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetListStream(63, value)
}

//go:nosplit
func (self class) SetStreamCount(stream_count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStreamCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the BPM of the playlist, which can vary depending on the clip being played.
*/
//go:nosplit
func (self class) GetBpm() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the stream at playback position index.
*/
//go:nosplit
func (self class) SetListStream(stream_index gd.Int, audio_stream gdclass.AudioStream)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, mmm.Get(audio_stream[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_list_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the stream at playback position index.
*/
//go:nosplit
func (self class) GetListStream(ctx gd.Lifetime, stream_index gd.Int) gdclass.AudioStream {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_list_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.AudioStream
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShuffle(shuffle bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shuffle)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_shuffle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShuffle() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_shuffle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFadeTime(dec gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, dec)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_fade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFadeTime() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_get_fade_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLoop(loop bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasLoop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioStreamPlaylist.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func init() {classdb.Register("AudioStreamPlaylist", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
